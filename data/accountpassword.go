package data

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

type AccountPassword struct {
	algorithm  string `bson:"algorithm"`
	salt       []byte `bson:"salt"`
	iteration  int    `bson:"iteration"`
	keyLen     int    `bson:"keylen"`
	derivedKey []byte `bson:"derivedkey"`
}

func NewAccountPassword(raw string) (AccountPassword, error) {
	accPass := AccountPassword{}
	accPass.salt = make([]byte, 32)
	accPass.algorithm = "SHA1"
	_, err := rand.Read(accPass.salt)
	if err != nil {
		return AccountPassword{}, err
	}
	accPass.iteration = 5678
	accPass.keyLen = 64
	accPass.derivedKey = pbkdf2.Key([]byte(raw), accPass.salt, accPass.iteration, accPass.keyLen, sha1.New)

	return accPass, nil
}

func (p AccountPassword) Match(raw string) bool {
	key := pbkdf2.Key([]byte(raw), p.salt, p.iteration, p.keyLen, sha1.New)
	return bytes.Equal(key, p.derivedKey)
}
