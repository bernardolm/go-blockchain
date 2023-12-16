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
		fmt.Printf("args are required: difficulty and block number\n")
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) < 1 {
		fmt.Printf("difficulty required\n")
		os.Exit(1)
	}
	difficult, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("difficulty error: %\n", err)
		os.Exit(1)
	}
	bc := blockchain.New(difficult)

	if len(args) < 2 {
		fmt.Printf("block number required\n")
		os.Exit(1)
	}
	var blockNumber int
	blockNumber, err = strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("block number error: %\n", err)
		os.Exit(1)
	}

	chain := bc.Chain()

	for i := 1; i <= blockNumber; i++ {

		block := bc.CreateBlock([]byte(fmt.Sprintf("Block %d", i)))
		mineInfo := bc.MineBlock(block)

		chain = bc.PushBlock(mineInfo.MinedBlock)
	}

	fmt.Printf("--- GENERATED CHAIN ---\n")
	fmt.Println(chain)
}
