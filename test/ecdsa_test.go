package test

import (
	"github.com/sirupsen/logrus"
	common2 "github.com/sivo4kin/ea-starter/common"
	"testing"
)

func TestECDSA(t *testing.T) {
	pkey := common2.ToECDSAFromHex("0x469e5c05e289274dd8570c31f2f0f21236f2e071613ac9c565821985e7ae641e")
	logrus.Print("PUBLIC", pkey.PublicKey)
}
