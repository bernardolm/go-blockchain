package blockchain

import (
	"fmt"
	"time"

	"github.com/bernardolm/blockchain-poc/helpers"
)

type header struct {
	nonce     int
	blockHash string
}

type payload struct {
	sequence     int
	timestamp    int64
	data         interface{}
	previousHash string
}

type block struct {
	header  header
	Payload *payload
}

type minedBlock struct {
	MinedBlock *block
	minedHash  string
	shortHash  string
	mineTime   int
}

type Blockchain struct {
	chain      []block
	powPrefix  string
	difficulty int
}

func (b *Blockchain) createGenesisBlock() *block {
	p := payload{
		timestamp:    time.Now().UnixNano(),
		data:         string("Genesis Block"),
		previousHash: "",
	}

	return &block{
		header: header{
			blockHash: helpers.Hash(p),
		},
	}
}

func (b *Blockchain) lastBlock() *block {
	return &b.chain[len(b.chain)-1]
}

func (b *Blockchain) Chain() []block {
	return b.chain
}

func (b *Blockchain) getPreviousBlockHash() string {
	return b.lastBlock().header.blockHash
}

func (b *Blockchain) CreateBlock(data []byte) *payload {
	p := payload{
		sequence:     b.lastBlock().Payload.sequence + 1,
		timestamp:    time.Now().UnixNano(),
		data:         data,
		previousHash: b.getPreviousBlockHash(),
	}

	fmt.Printf("Created block %d: %v\n", p.sequence, p)

	return &p
}

func (b *Blockchain) MineBlock(p *payload) *minedBlock {
	nonce := 0
	startTime := time.Now().UnixNano()

	for {
		blockHash := helpers.Hash(*p)
		proofingHash := helpers.Hash(blockHash + fmt.Sprint(nonce))

		if helpers.IsHashProofed(proofingHash, b.difficulty, b.powPrefix) {
			endTime := time.Now().UnixNano()
			shortHash := blockHash[0:12]
			mineTime := (endTime - startTime) / 1000

			fmt.Printf("Mined block %d in %d seconds. "+
				"Hash: %v (%v attempts)\n",
				p.sequence, mineTime, &shortHash, nonce)

			minedBlock := minedBlock{
				MinedBlock: &block{
					Payload: p,
					header: header{
						nonce:     nonce,
						blockHash: blockHash,
					},
				},
				minedHash: proofingHash,
				mineTime:  int(mineTime),
			}

			return &minedBlock
		}
		nonce++
	}
}

func (b *Blockchain) verifyBlock(bl *block) bool {
	if bl.Payload.previousHash != b.getPreviousBlockHash() {
		fmt.Printf(
			"Invalid block #%d: Previous block hash is %s not %s\n",
			bl.Payload.sequence,
			b.getPreviousBlockHash()[0:12],
			bl.Payload.previousHash[0:12],
		)
		return false
	}

	h := fmt.Sprintf("%s%d", helpers.Hash(bl.Payload), bl.header.nonce)

	if !helpers.IsHashProofed(h, b.difficulty, b.powPrefix) {
		fmt.Printf(
			"Invalid block #%d: "+
				"Hash is not proofed, nonce %d is not valid\n",
			bl.Payload.sequence, bl.header.nonce)
		return false
	}

	return true
}

func (b *Blockchain) PushBlock(bl *block) []block {
	if b.verifyBlock(bl) {
		b.chain = append(b.chain, *bl)
		fmt.Printf("Pushed block %v\n", bl)
	}
	return b.chain
}

func New(difficulty int) *Blockchain {
	d := 4
	if difficulty > 0 {
		d = 4
	}

	blockchain := Blockchain{
		difficulty: d,
	}

	genesisBlock := blockchain.createGenesisBlock()
	blockchain.chain = append(blockchain.chain, *genesisBlock)

	return &blockchain
}
