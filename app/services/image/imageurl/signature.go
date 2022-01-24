package imageurl

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/mises-id/sns-storagesvc/config/env"
)

var (
	signatureSize = 32
	//signKey                     = env.Envs.SignKey
	//signSalt                    = env.Envs.SignSalt
	errInvalidSignature         = errors.New("Invalid signature")
	errInvalidSignatureEncoding = errors.New("Invalid signature encoding")
)

func verifySignature(signature, path, opPath string) (err error) {
	if !env.Envs.SignURL {
		return nil
	}
	var key, salt []byte
	if key, err = hex.DecodeString(env.Envs.SignKey); err != nil {
		return err
	}
	if salt, err = hex.DecodeString(env.Envs.SignSalt); err != nil {
		return err
	}
	messageMAC, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		return errInvalidSignature
	}
	opPath = strings.TrimPrefix(opPath, "/")
	opPath = strings.TrimSuffix(opPath, "/")
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(path))
	str := fmt.Sprintf("%s%s", encodedURL, opPath)
	str = strings.TrimSuffix(str, "/")
	if hmac.Equal(messageMAC, signatureFor(str, key, salt)) {
		return nil
	}
	return errInvalidSignature
}

func signatureFor(str string, key, salt []byte) []byte {

	mac := hmac.New(sha256.New, key)
	mac.Write(salt)
	mac.Write([]byte(str))
	expectedMAC := mac.Sum(nil)
	if signatureSize < 32 {
		return expectedMAC[:signatureSize]
	}
	return expectedMAC
}

func signature(path, opPath string) (signature string, err error) {
	var key, salt []byte
	if key, err = hex.DecodeString(env.Envs.SignKey); err != nil {

		return "", err
	}

	if salt, err = hex.DecodeString(env.Envs.SignSalt); err != nil {
		return "", err

	}
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(path))
	str := fmt.Sprintf("%s%s", encodedURL, opPath)
	str = strings.TrimSuffix(str, "/")
	mac := hmac.New(sha256.New, key)
	mac.Write(salt)
	mac.Write([]byte(str))
	signature = base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return signature, nil
}
