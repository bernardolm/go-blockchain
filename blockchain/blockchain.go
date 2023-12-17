package blockchain

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/go-blockchain/helpers"
)

type header struct {
	nonce     int
	blockHash string
}

type payload struct {
	Sequence     int
	Timestamp    int64
	Data         interface{}
	PreviousHash string
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
		Timestamp: time.Now().UnixNano(),
		Data:      string("Genesis Block"),
	}

	return &block{
		header: header{
			blockHash: helpers.Hash(p),
		},
		Payload: &p,
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
		Sequence:     b.lastBlock().Payload.Sequence + 1,
		Timestamp:    time.Now().UnixNano(),
		Data:         data,
		PreviousHash: b.getPreviousBlockHash(),
	}

	log.Infof("Created block %d: %#v", p.Sequence, p)

	return &p
}

func (b *Blockchain) MineBlock(p *payload) *minedBlock {
	nonce := 0
	startTime := time.Now().UnixNano()

	for {
		/******************* FOR SECURITY **********************/
		limit := 10000
		if nonce >= limit {
			log.Panicf("too many attemps: %d", limit)
		}
		/******************* FOR SECURITY **********************/

		blockHash := helpers.Hash(*p)
		proofingHash := helpers.Hash(blockHash + fmt.Sprint(nonce))

		if helpers.IsHashProofed(proofingHash, b.difficulty, b.powPrefix) {
			endTime := time.Now().UnixNano()
			shortHash := blockHash[0:12]
			mineTime := (endTime - startTime) / 1000

			log.Infof("Mined block %d in %d seconds. "+
				"Hash: %v (%v attempts)",
				p.Sequence, mineTime, &shortHash, nonce)

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
	if bl.Payload.PreviousHash != b.getPreviousBlockHash() {
		log.Errorf(
			"Invalid block #%d: Previous block hash is %s not %s",
			bl.Payload.Sequence,
			b.getPreviousBlockHash()[0:12],
			bl.Payload.PreviousHash[0:12],
		)
		return false
	}

	h := fmt.Sprintf("%s%d", helpers.Hash(bl.Payload), bl.header.nonce)

	if !helpers.IsHashProofed(h, b.difficulty, b.powPrefix) {
		log.Errorf(
			"Invalid block #%d: Hash is not proofed, nonce %d is not valid",
			bl.Payload.Sequence, bl.header.nonce)
		return false
	}

	return true
}

func (b *Blockchain) PushBlock(bl *block) []block {
	if b.verifyBlock(bl) {
		b.chain = append(b.chain, *bl)
		log.Infof("Pushed block %v", bl)
	}
	return b.chain
}

func New(difficulty int) *Blockchain {
	blockchain := Blockchain{
		difficulty: difficulty,
		chain:      []block{},
	}

	genesisBlock := blockchain.createGenesisBlock()
	log.Debugf("genesisBlock: %#v", genesisBlock)

	blockchain.chain = append(blockchain.chain, *genesisBlock)

	return &blockchain
}
