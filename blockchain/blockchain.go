package blockchain

import (
	"fmt"
	"math"
	"time"

	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/go-blockchain/helpers"
)

type header struct {
	blockHash string
	nonce     int
}

type payload struct {
	Data         string
	PreviousHash string
	Sequence     int
	Timestamp    int64
}

type block struct {
	header  header
	Payload payload
}

type Chain []block

type minedBlock struct {
	MinedBlock block
	minedHash  string
	mineTime   int
	shortHash  string
}

type Blockchain struct {
	blockNumber int
	chain       Chain
	difficulty  int
	powPrefix   string
}

func (b Blockchain) createGenesisBlock() block {
	p := payload{
		Timestamp: time.Now().UnixNano(),
		Data:      string("Genesis Block"),
	}

	data := fmt.Sprintf("%#v", p)
	h := helpers.Hash(data)

	log.
		WithField("data", data).
		WithField("hash", h[0:b.difficulty]).
		Debug("Blockchain.createGenesisBlock: hash created from data")

	return block{
		header: header{
			blockHash: h,
		},
		Payload: p,
	}
}

func (b Blockchain) lastBlock() block {
	return b.chain[len(b.chain)-1]
}

func (b Blockchain) Chain() []block {
	return b.chain
}

func (b Blockchain) getPreviousBlockHash() string {
	return b.lastBlock().header.blockHash
}

func (b Blockchain) CreatePayload(data string) payload {
	p := payload{
		Sequence:     b.lastBlock().Payload.Sequence + 1,
		Timestamp:    time.Now().UnixNano(),
		Data:         data,
		PreviousHash: b.getPreviousBlockHash(),
	}

	log.
		WithField("data", p.Data).
		WithField("lastBlockPayloadSequence", b.lastBlock().Payload.Sequence).
		WithField("previousHash", p.PreviousHash[0:b.difficulty]).
		WithField("sequence", p.Sequence).
		Infof("Block #%d created", p.Sequence)

	return p
}

func (b Blockchain) Mine(p payload) minedBlock {
	nonce := 0
	start := time.Now()

	for {
		/******************* FOR SECURITY *********************/
		if float64(nonce) >= math.Pow((26+10), float64(b.difficulty)) {
			log.Fatal("it's enough")
		}
		/******************* FOR SECURITY *********************/

		data := fmt.Sprintf("%#v", p)
		blockHash := helpers.Hash(data)
		proofingData := fmt.Sprintf("%s%d", blockHash, nonce)
		proofingHash := helpers.Hash(proofingData)

		log.
			WithField("blockHash", blockHash[0:b.difficulty]).
			WithField("nonce", nonce).
			WithField("proofingData", "..."+
				proofingData[len(proofingData)-b.difficulty:]).
			WithField("proofingHash", proofingHash[0:b.difficulty]).
			Debug("Mine attempt")

		if helpers.IsHashProofed(
			proofingHash,
			b.difficulty,
			b.powPrefix) {
			elapsed := time.Since(start)

			log.
				WithField("attempts", nonce).
				WithField("blockHash", blockHash[0:b.difficulty]).
				WithField("duration", elapsed).
				WithField("sequence", p.Sequence).
				Infof("Block #%d mining finished", p.Sequence)

			return minedBlock{
				MinedBlock: block{
					Payload: p,
					header: header{
						nonce:     nonce,
						blockHash: blockHash,
					},
				},
				minedHash: proofingHash,
				mineTime:  int(elapsed),
			}
		}
		nonce++
	}
}

func (b Blockchain) verifyBlock(bl block) bool {
	lastChainBlockHash := b.getPreviousBlockHash()
	payload := fmt.Sprintf("%#v", bl.Payload)

	cph := bl.Payload.PreviousHash
	if len(cph) > 0 {
		cph = bl.Payload.PreviousHash[0:b.difficulty]
	}

	log.
		WithField("currentPreviousHash", cph).
		WithField("lastChainBlockHash", lastChainBlockHash[0:b.difficulty]).
		WithField("nonce", bl.header.nonce).
		WithField("sequence", bl.Payload.Sequence).
		Debugf("verifyBlock: Block #%d to check previous",
			bl.Payload.Sequence)

	if bl.Payload.PreviousHash != lastChainBlockHash {
		pp.Println(payload)

		log.
			WithField("nonce", bl.header.nonce).
			WithField("sequence", bl.Payload.Sequence).
			Errorf("verifyBlock: Invalid block #%d (previous not equal)",
				bl.Payload.Sequence)

		return false
	}

	hash := helpers.Hash(payload)
	dataToProof := fmt.Sprintf("%s%d", hash, bl.header.nonce)
	hashToProof := helpers.Hash(dataToProof)

	log.
		WithField("payload", payload).
		WithField("dataToProof", dataToProof).
		WithField("hash", hash[0:b.difficulty]).
		WithField("hashToProof", hashToProof[0:b.difficulty]).
		WithField("nonce", bl.header.nonce).
		Debugf("verifyBlock: Block #%d to proof hash",
			bl.Payload.Sequence)

	if !helpers.IsHashProofed(hashToProof, b.difficulty, b.powPrefix) {
		log.
			WithField("hashToProof", hashToProof[0:b.difficulty]).
			WithField("nonce", bl.header.nonce).
			WithField("sequence", bl.Payload.Sequence).
			Errorf("verifyBlock: Invalid block #%d (not proofed, nonce invalid)",
				bl.Payload.Sequence)

		return false
	}

	return true
}

func (b *Blockchain) PushBlock(bl block) Chain {
	bld := fmt.Sprintf("%#v", bl)

	log.
		WithField("block", bld).
		Debug("PushBlock: to verify")

	if b.verifyBlock(bl) {
		b.chain = append(b.chain, bl)
		log.
			WithField("blockHash", bl.header.blockHash[0:b.difficulty]).
			WithField("nonce", bl.header.nonce).
			WithField("sequence", bl.Payload.Sequence).
			Infof("Block #%d verified and pushed", bl.Payload.Sequence)
	}

	return b.chain
}

func New(difficulty, blockNumber int) Blockchain {
	blockchain := Blockchain{
		difficulty:  difficulty,
		blockNumber: blockNumber,
	}

	gb := blockchain.createGenesisBlock()
	log.
		WithField("block", fmt.Sprintf("%#v", gb)).
		WithField("hash", gb.header.blockHash[0:difficulty]).
		Debug("blockchain.New: created genesis block")

	blockchain.chain = append(blockchain.chain, gb)

	return blockchain
}
