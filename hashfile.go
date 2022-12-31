package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func HashFile(path string) (string, error) {
	file, err := os.Open(path) // Open the file for reading
	if err != nil {
		return "", err
	}
	defer file.Close() // Be sure to close your file!

	hash := sha256.New() // Use the Hash in crypto/sha256

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	sum := fmt.Sprintf("%x", hash.Sum(nil)) // Get encoded hash sum
	return sum, nil
}
