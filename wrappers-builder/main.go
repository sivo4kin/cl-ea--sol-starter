package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"text/template"
)

type JsonContract struct {
	Abi         interface{}
	Code        string `json:"bytecode"`
	RuntimeCode string `json:"deployedBytecode"`
}

var (
	app         *cli.App
	packageName string
	outputDir   string
	solFlag     = cli.StringFlag{
		Name:  "sol",
		Usage: "this flag defines that input files will be .sol's and should be compiled with build/wrappers/solc",
	}
	jsonFlag = cli.StringFlag{
		Name:  "json",
		Usage: "this flag defines that input files will be .json's compiled by Truffle",
	}
	outputFlag = cli.StringFlag{
		Name:  "output, out",
		Usage: "this flag defines the output directory for Go contract wrappers.",
		Value: "wrappers",
	}
	packageFlag = cli.StringFlag{
		Name:  "package, pkg",
		Usage: "this flag defines the package name that will be used for Go contract wrappers",
		Value: "wrappers",
	}
)

func init() {
	app = cli.NewApp()
	app.Flags = []cli.Flag{
		solFlag,
		jsonFlag,
		outputFlag,
		packageFlag,
	}
	app.Usage = "use for generate Go wrappers for smart contracts"
	app.Description = "CLI App for compile Go wrappers for smart contracts"
	app.Action = compile
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Println(err)
			os.Exit(-1)
		}
	}()

	app.Run(os.Args)
}

//findAllSourceFiles recursively find files in directory with root root path with given extensions
func findAllSourceFiles(root, extension string) ([]string, error) {
	var sourceFiles []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == extension {
			sourceFiles = append(sourceFiles, path)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return sourceFiles, nil
}

//dumpContracts generates go binding files from contract ABI and write them to files
func dumpContracts(contracts map[string]*compiler.Contract, packageName, outputDir string) error {
	for key, value := range contracts {
		//logrus.Printf("VALUE %x \n",value.Code)
		//logrus.Printf("KEY %s \n",key)
		var (
			types   []string
			fsigs   []map[string]string
			libs    = make(map[string]string)
			aliases = make(map[string]string)
		)
		logrus.Print("1")
		//logrus.Print(value)

		//if value.Code == "0x" {
		abi, err := json.Marshal(value.Info.AbiDefinition)
		if err != nil {
			logrus.Fatal(err)
			return err
		}
		logrus.Print("2")
		keyParts := strings.Split(key, ":")
		types = append(types, keyParts[len(keyParts)-1])
		fsigs = append(fsigs, value.Hashes)
		logrus.Print("3")
		code, err := bind.Bind(types, []string{string(abi)}, []string{value.Code}, fsigs, packageName, bind.LangGo, libs, aliases)
		if err != nil {
			logrus.Fatal(err)
			return err
		}
		logrus.Print("4")
		if err := os.MkdirAll(outputDir, 0700); err != nil {
			logrus.Fatal(err)
			return err
		}
		//logrus.Print(code)
		if err := ioutil.WriteFile(filepath.Join(outputDir, types[0]+".go"), []byte(code), 0600); err != nil {
			logrus.Fatal(err)
			return err
		}
	}
	//}
	return nil
}

func findContract(name string, contracts map[string]*compiler.Contract) *compiler.Contract {
	name = "_" + name + ".sol:" + name
	for k, v := range contracts {
		if strings.HasSuffix(k, name) {
			return v
		}
	}
	return nil
}

//updateFront generates JS binding file from contract ABI
func updateFront(contracts map[string]*compiler.Contract) {
	js, err := os.Create("./arm/src/contracts.json")
	if err != nil {
		panic(err)
	}
	n := reflect.ValueOf(contracts).MapKeys()
	sort.Slice(n, func(i, j int) bool { return n[i].String() < n[j].String() })
	fmt.Fprintf(js, "{\n")
	for _, fullName := range n {
		name := fullName.String()
		contract := contracts[name]
		if contract.Code != "0x" {
			name = name[strings.Index(name, ":")+1:]
			abi, err := json.Marshal(contract.Info.AbiDefinition)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(js, `"%sAbi": %s,%s`, name, string(abi), "\n")
			fmt.Fprintf(js, `"%sBytecode": "%s",%s`, name, contract.Code, "\n")
		}
	}
	fmt.Fprintf(js, "\"closer\":{}}")
	js.Close()
}

func createFromTemplate(templatePath, filePath string, data interface{}) {
	fileTemplate, err := ioutil.ReadFile(templatePath)
	if err != nil {
		panic(fmt.Errorf("can't read %s: %v", templatePath, err))
	}
	t := template.New("")
	_, err = t.Parse(string(fileTemplate))
	if err != nil {
		panic(fmt.Errorf("can't parse %s: %v", templatePath, err))
	}
	genesisWriter, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Errorf("can't create genesis.json: %v", err))
	}
	err = t.Execute(genesisWriter, data)
	if err != nil {
		panic(fmt.Errorf("can't prepare genesis.json: %v", err))
	}
}

func PackagePathFile(name string) string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Sprintf("Could not locate package dir with %s in it.", name))
	}
	return path.Join(filepath.Dir(thisFile), name)
}

//Method compiles go wrappers from sol files or json files
func compile(c *cli.Context) {
	outputDir = c.String("output")
	packageName = c.String("package")
	var (
		err         error
		sourceFiles []string
		contracts   map[string]*compiler.Contract
	)
	switch {
	case c.IsSet(solFlag.Name):
		root := c.String("sol")
		sourceFiles, err = findAllSourceFiles(root, ".sol")
		if err != nil {
			panic(err)
		}
		var solcPath string
		if runtime.GOOS == "windows" {
			solcPath = PackagePathFile("solc\\solc.exe")
		} else {
			solcPath = PackagePathFile("solc/solc-static-linux")
		}
		err = os.Chmod(solcPath, 0700)
		if err != nil {
			panic(err)
		}
		contracts, err = compiler.CompileSolidity(solcPath, sourceFiles...)
	case c.IsSet(jsonFlag.Name):
		root := c.String("json")
		sourceFiles, err = findAllSourceFiles(root, ".json")
		logrus.Printf("found %d files", reflect.ValueOf(sourceFiles).Len())
		if err != nil {
			panic(err)
		}
		contracts, err = CompileJson(sourceFiles...)
		//logrus.Printf("DUMPING CONTRACTS %t package %s dir %s",contracts, packageName, outputDir)
	default:
		panic("You should either set --sol or --json flag. Run command with --help or --h flag to see more information about command")
	}
	if err != nil {
		panic(err)
	}
	logrus.Printf("DUMPING CONTRACTS package %s dir %s", packageName, outputDir)

	err = dumpContracts(contracts, packageName, outputDir)
	if err != nil {
		logrus.Fatal(err)
		panic(err)
	}
	//updateFront(contracts)
}

func CompileJson(sourcefiles ...string) (contracts map[string]*compiler.Contract, err error) {
	contracts = make(map[string]*compiler.Contract)
	for num, file := range sourcefiles {
		logrus.Printf("number %d file %s", num, file)
		//TODO:find a proper way to find contract name
		o := strings.Split(file, "/")
		name := strings.Split(o[len(o)-1], ".")[0]
		jsonOutput, err := ioutil.ReadFile(file)
		//logrus.Printf("JSON", jsonOutput)
		if err != nil {
			return contracts, err
		}
		var output JsonContract
		err = json.Unmarshal(jsonOutput, &output)
		if err != nil {
			return contracts, nil
		}
		logrus.Printf("FILE %s NAME %s\nQQQQQQQQQQQQQQQQQQQQQQ", file, name)

		contracts[file+":"+name] = &compiler.Contract{
			Code: output.Code,
			Info: compiler.ContractInfo{AbiDefinition: output.Abi},
		}
	}
	logrus.Print("EEEEEEEEEEEEEEE")
	return contracts, nil
}
