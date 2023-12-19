package runner

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/go-blockchain/blockchain"
)

func Run(bc *blockchain.Blockchain, blockNumber int) (blockchain.Chain, time.Duration) {
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

	return chain, elapsed
}
