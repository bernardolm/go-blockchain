package helpers

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Hash(data interface{}) string {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		log.Println(err)
	}

	byts := buf.Bytes()

	h := sha256.New()

	h.Write(byts)
	bs := h.Sum(nil)

	// log.Debugf("Hash.data: '%#v'", data)
	// log.Debugf("Hash.bytes: '%x'", bs)

	return string(bs)
}

func IsHashProofed(hash string, difficulty int, prefix string) bool {
	if difficulty == 0 {
		difficulty = 4
	}

	if prefix == "" {
		prefix = "0"
	}

	mask := "%" + prefix
	check := fmt.Sprintf(mask+"d", difficulty)

	return hash[0:difficulty] == check
}
