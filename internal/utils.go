package internal

import (
	"fmt"
	"math/rand"
	"os"
)

func generateID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func InitArchiveFolder() {
	os.MkdirAll("./archives", os.ModePerm)
}
