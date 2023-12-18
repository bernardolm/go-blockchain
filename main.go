package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/go-blockchain/blockchain"
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

	chain := bc.Chain()

	start := time.Now()

	for i := 1; i <= blockNumber; i++ {
		data := fmt.Sprintf("block %d", i)

		payload := bc.CreatePayload(data)
		mineInfo := bc.Mine(payload)
		chain = bc.PushBlock(mineInfo.MinedBlock)

		log.
			WithField("data", data).
			WithField("payload", fmt.Sprintf("%#v", payload)).
			WithField("mineInfo", fmt.Sprintf("%#v", mineInfo)).
			Debug("main: loop")
	}

	elapsed := time.Since(start)

	log.
		WithField("length", len(chain)).
		WithField("duration", elapsed).
		Info("main: generated chain")
}
