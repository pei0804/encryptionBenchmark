package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"crypto/sha256"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"
)

func toHashFromScrypt(str string) string {
	converted, _ := scrypt.Key([]byte(str), []byte("some salt"), 16384, 8, 1, 32)
	return hex.EncodeToString(converted[:])
}

func toHashFromBcrypt(str string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(str), 10)
	return hex.EncodeToString(converted[:])
}

func toHashFromSha256(str string) string {
	converted := sha256.Sum256([]byte(str))
	return hex.EncodeToString(converted[:])
}

func toHashFromSha512(str string) string {
	converted := sha512.Sum512([]byte(str))
	return hex.EncodeToString(converted[:])
}

func toHashFromSha3_256(str string) string {
	converted := sha3.Sum256([]byte(str))
	return hex.EncodeToString(converted[:])
}

func toHasgFromRipemd_160(str string) string {
	io.WriteString(ripemd160.New(), str)
	converted := sha3.Sum256([]byte(str))
	return hex.EncodeToString(converted[:])
}

func main() {

	str := "hash"

	// scrypt
	result := testing.Benchmark(func(b *testing.B) {
		toHashFromBcrypt(str)
	})
	fmt.Printf("scrypt: %s\n", result.T)

	// bcrypt
	result = testing.Benchmark(func(b *testing.B) {
		toHashFromScrypt(str)
	})
	fmt.Printf("bcrypt: %s\n", result.T)

	// sha256.Sum256
	result = testing.Benchmark(func(b *testing.B) {
		toHashFromSha256(str)
	})
	fmt.Printf("sha256: %s\n", result.T)

	// sha512
	result = testing.Benchmark(func(b *testing.B) {
		toHashFromSha512(str)
	})
	fmt.Printf("sha512: %s\n", result.T)

	// sha3-256
	result = testing.Benchmark(func(b *testing.B) {
		toHashFromSha3_256(str)
	})
	fmt.Printf("sha3-256: %s\n", result.T)

	// ripemd
	result = testing.Benchmark(func(b *testing.B) {
		toHasgFromRipemd_160(str)
	})
	fmt.Printf("Ripemd: %s\n", result.T)
}
