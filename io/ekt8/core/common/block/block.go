package block

import (
	"fmt"
	"sort"

	"../"
	"../transaction"
	"golang.org/x/crypto/sha3"
)

type Block struct {
	Height       int64
	PreviousHash common.Hash
	CurrentHash  common.Hash
	Transactions transaction.Transactions
}

func (block *Block) sort() {
	sort.Sort(block.Transactions)
}

func (block *Block) Hash() {
	hash := sha3.New256()
	result := hash.Sum(nil)
	fmt.Println(string(result))
}
