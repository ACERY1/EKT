package common

import (
	"bytes"

	"github.com/EducationEKT/EKT/io/ekt8/crypto"
)

type Account struct {
	Address    Address
	PublickKey crypto.PublicKey
}

func (account *Account) Validate(data []byte, sign []byte) bool {
	decryptData := crypto.Decrypt(sign, account.PublickKey)
	return bytes.Equal(data, decryptData)
}
