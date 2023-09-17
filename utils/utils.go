package utils

import (
	"github.com/jaevor/go-nanoid"
)

const DefaultAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func CreateNanoID(length int) (string, error) {
	nanoID, err := nanoid.CustomASCII(DefaultAlphabet, length)
	if err != nil {
		return "", err
	}
	return nanoID(), nil
}
