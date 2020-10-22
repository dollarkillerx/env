package env

import (
	"log"
	"testing"
)

type Name struct {
	Goos string `json:"goos"`
	GoPath string `json:"gopath"`
}

func TestEnv(t *testing.T) {
	var n Name
	err := FillBase(&n)
	log.Println(err)
	log.Println(n)
}
