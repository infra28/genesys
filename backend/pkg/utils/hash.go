package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var p = &Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, p.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func ComparePassword(password, encodedHash string) (bool, error) {
	vals := strings.Split(encodedHash, "$")
	var mem, iter uint32
	var para uint8
	fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &mem, &iter, &para)

	salt, _ := base64.RawStdEncoding.DecodeString(vals[4])
	hash, _ := base64.RawStdEncoding.DecodeString(vals[5])

	comparisonHash := argon2.IDKey([]byte(password), salt, iter, mem, para, uint32(len(hash)))

	return subtle.ConstantTimeCompare(hash, comparisonHash) == 1, nil
}
