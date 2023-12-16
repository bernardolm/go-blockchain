package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/bernardolm/blockchain-poc/blockchain"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()

	blockchain := blockchain.New(args[2])

	var blockNumber int
	var err error
	blockNumber, err = strconv.Atoi(args[3])
	if err != nil {
		panic(err)
	}

	chain := blockchain.Chain()

	for i := 1; i <= blockNumber; i++ {
		block := blockchain.Createblock(`Block ${i}`)
		mineInfo := blockchain.Mineblock(block)
		chain = blockchain.Pushblock(mineInfo.minedBlock)
	}

	fmt.Println("--- GENERATED CHAIN ---\n")
	fmt.Println(chain)
}
