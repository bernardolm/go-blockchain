package main

import (
	"flag"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/go-blockchain/blockchain"
)

func main() {
	log.SetLevel(log.DebugLevel)

	flag.Parse()
	args := flag.Args()
	difficulty := 1
	blockNumber := 1

	var err error

	if len(args) > 0 {
		difficulty, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("difficulty error: %v", err)
			// os.Exit(1)
		}
	}
	log.Debugf("difficulty: %d", difficulty)
	bc := blockchain.New(difficulty)

	if len(args) > 1 {
		blockNumber, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("block number error: %v", err)
			// os.Exit(1)
		}
	}
	log.Debugf("block number: %d", blockNumber)

	chain := bc.Chain()

	for i := 1; i <= blockNumber; i++ {
		data := []byte(fmt.Sprintf("Block %d", i))
		block := bc.CreateBlock(data)
		mineInfo := bc.MineBlock(block)
		chain = bc.PushBlock(mineInfo.MinedBlock)
	}

	log.Infof("--- GENERATED CHAIN ---\n%#v", chain)
}
