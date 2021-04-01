package test

import (
	"github.com/sirupsen/logrus"
	common2 "github.com/sivo4kin/ea-starter/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestECDSA(t *testing.T) {
	pkey, err := common2.ToECDSAFromHex("0x469e5c05e289274dd8570c31f2f0f21236f2e071613ac9c565821985e7ae641e")
	require.Nil(t, err)
	logrus.Print("PUBLIC", pkey.PublicKey)
}
