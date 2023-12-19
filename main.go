package main

import (
	"flag"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/go-blockchain/blockchain"
	"github.com/bernardolm/go-blockchain/runner"
)

func main() {
	// log.SetLevel(log.DebugLevel)

	flag.Parse()
	args := flag.Args()
	difficulty := 4
	blockNumber := 200

	var err error

	if len(args) > 1 {
		blockNumber, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("main: block number error: %v", err)
		}
	}
	log.Debugf("main: flag arg block number: %d", blockNumber)

	if len(args) > 0 {
		difficulty, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("main: difficulty error: %v", err)
		}
	}
	log.Debugf("main: flag arg difficulty: %d", difficulty)

	bc := blockchain.New(difficulty, blockNumber)

	log.
		WithField("blockchain", fmt.Sprintf("%#v", bc)).
		Debug("main: blockchain created")

	chain, elapsed := runner.Run(&bc, blockNumber)

	log.
		WithField("length", len(chain)).
		WithField("duration", elapsed).
		Info("main: generated chain")
}
