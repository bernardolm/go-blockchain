package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Hash(data string) string {
	// _, file, line, _ := runtime.Caller(1)

	s := sha256.New()
	s.Reset()
	s.Write([]byte(data))
	bs := s.Sum(nil)

	h := hex.EncodeToString(bs)

	// log.
	// 	WithField("caller", fmt.Sprintf("%s#%d", file, line)).
	// 	WithField("data", data).
	// 	WithField("hash", h[0:4]).
	// 	Debug("helpers.Hash: hash created from data")

	return h
}

func IsHashProofed(hash string, difficulty int, prefix string) bool {
	if prefix == "" {
		prefix = "0"
	}

	check := strings.Repeat(prefix, difficulty)
	hashToProof := hash[0:difficulty]

	log.
		WithField("check", check).
		WithField("hashToProof", hashToProof).
		Debug("helpers.IsHashProofed")

	return hashToProof == check
}
