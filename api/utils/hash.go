package utils

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// GenerateHash generates an unsigned hashed string of random uni code chars
func GenerateHash(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789" + "({[$.`!~-/&#%@]})" + os.Getenv("secret_key"))
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(characters[rand.Intn(len(characters))])
	}

	return b.String()
}

// GenerateTextHash generates an unsigned hashed string of letters
func GenerateTextHash(length int) string {
	rand.Seed(time.Now().UnixNano())
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(characters[rand.Intn(len(characters))])
	}

	return b.String()
}
