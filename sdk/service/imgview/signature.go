package imgview

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

var (
	keyInvalid  = "Key expected to be hex-encoded string"
	saltInvalid = "Salt expected to be hex-encoded string"
)

func signature(key, salt, path, opPath string) (string, error) {
	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal(keyInvalid)
		return "", errors.New(keyInvalid)
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal(saltInvalid)
		return "", errors.New(saltInvalid)
	}

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(path))

	str := fmt.Sprintf("%s%s", encodedURL, opPath)
	fmt.Println("sign str:", str)
	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(str))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return signature, nil
}
