package helpers

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

func Hash(data interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(data); err != nil {
		log.Println(err)
	}

	byts := buf.Bytes()

	h := sha256.New()

	h.Write(byts)
	bs := h.Sum(nil)

	fmt.Printf("%v\n", data)
	fmt.Printf("%x\n", bs)

	return bs
}

func IsHashProofed(hash string, difficulty int, prefix string) bool {
	mask := "%" + prefix
	check := fmt.Sprintf(mask+"d", difficulty)
	return hash[0:4] == check
}
