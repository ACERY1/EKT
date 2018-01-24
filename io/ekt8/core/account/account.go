package account

import (
	"bytes"

	"../../crypto"
	"../common"
)

type Account struct {
	Address    common.Address
	PublickKey crypto.PublicKey
}

func (account *Account) Validate(data []byte, sign []byte) bool {
	decryptData := crypto.Decrypt(sign, account.PublickKey)
	return bytes.Equal(data, decryptData)
}
