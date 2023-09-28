package main

import (
	"epay/internal/apiserver"
	"epay/pkg/log"
)

func main() {
	log.Init(nil)
	apiserver.NewApp("link2web3-background").Run()
}
