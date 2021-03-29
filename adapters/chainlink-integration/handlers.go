package chainlink_integration

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// RequestData makes the call to the external resources (in this case,
// that's simulated in chainlink.go).
func RequestData(w http.ResponseWriter, r *http.Request) Chainlink {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Log the headers
	log.Println("Headers:")
	log.Println(r.Header)

	// Log the input
	log.Println("Input:")
	log.Println(string(body))

	cl := Chainlink{}
	if err := json.Unmarshal(body, &cl); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	return cl

}

// WriteData returns the data to the requester
func WriteData(w http.ResponseWriter, result RunResult) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}

	// Log the output
	outString, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	log.Println("Output:")
	log.Println(string(outString))
}

// TaskRun reads the input JSON given from the core, calls the
// GetData method, and fulfills the request. It will log both the
// input and output JSON for troubleshooting.
func TaskRun(w http.ResponseWriter, r *http.Request) {
	cl := RequestData(w, r)
	WriteData(w, GetData(cl))
}

// PendingTaskRun reads the input JSON given from the core, calls the

// GetPending method, and returns a RunResult with the pending status
// set to true. It will log both the input and output JSON for troubleshooting.
func PendingTaskRun(w http.ResponseWriter, r *http.Request) {
	cl := RequestData(w, r)
	WriteData(w, GetPending(cl))
}

func GetRegisteredJobSpecs(w http.ResponseWriter, r *http.Request) {
	//  "<application>/runs?jobSpecId=:jobSpecId&size=1&page=2"
	url := "http://localhost:6688/v2/specs/"
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf(" HTTP STATUSNOT 200 %s", resp)
	}
	//WriteData(w, resp)
	log.Println(resp)
}

// ResumeFromPending allows you to resume a pending run, given its
// JobRunID.
func ResumeFromPending(w http.ResponseWriter, r *http.Request) {
	cl := RequestData(w, r)
	log.Printf("ID %s", cl.ID)
	result := GetData(cl)
	url := "http://localhost:6688/v2/runs/" + cl.ID
	outString, err := json.Marshal(result)
	log.Printf("outString %s", outString)

	//// trimmed := outString[1:len(outString)-1]
	//if err != nil {
	//	panic(err)
	//}

	log.Println("PATCHing to:")
	log.Println(url)

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(outString))
	req.Header.Set("Authorization", "Bearer 7a76ab15fd984d4ba6171c064f537c29")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Printf("respONCE %s code %d", resp.Status, resp.StatusCode)
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(responseBody))
	WriteData(w, result)
}

// ReturnBigInt reads the input JSON given from the core, calls the
// GetBigInt method, and fulfills the request. It will log both the
// input and output JSON for troubleshooting.
func ReturnBigInt(w http.ResponseWriter, r *http.Request) {
	cl := RequestData(w, r)
	WriteData(w, GetBigInt(cl))
}

// InputDataExample reads the input JSON given from the core, calls the
// GetRestData method, and fulfills the request. It will log both the
// input and output JSON for troubleshooting.
func InputDataExample(w http.ResponseWriter, r *http.Request) {
	cl := RequestData(w, r)
	WriteData(w, GetMockData(cl))
}

// ReturnError returns an error back to the caller
func ReturnError(w http.ResponseWriter, r *http.Request) {
	cl := RequestData(w, r)
	WriteData(w, GetError(cl))
}
