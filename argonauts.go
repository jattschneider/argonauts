package argonauts

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	stringPrefix = "ARG("
	stringSuffix = ")"
)

type Options struct {
	salt    []byte
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func DefaultOptions(salt []byte) *Options {
	return &Options{
		salt:    salt,
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}
}

func Salt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func Hash(opts *Options, password []byte) []byte {
	return argon2.IDKey(password, opts.salt, opts.time,
		opts.memory, opts.threads, opts.keyLen)
}

func Compare(opts *Options, password, hash []byte) (bool, error) {
	comparisonHash := argon2.IDKey([]byte(password), opts.salt, opts.time,
		opts.memory, opts.threads, opts.keyLen)
	return (subtle.ConstantTimeCompare(hash, comparisonHash) == 1), nil
}

func Sprint(hash []byte) string {
	return fmt.Sprintf("%v%v%v", stringPrefix, EncodeToString(hash), stringSuffix)
}

func EncodeToString(hash []byte) string {
	return base64.RawStdEncoding.EncodeToString(hash)
}

func InnerString(s string) string {
	return s[len(stringPrefix) : len(s)-len(stringSuffix)]
}

func ReadString(s string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(InnerString(s))
}
