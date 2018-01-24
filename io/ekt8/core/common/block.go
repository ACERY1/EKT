package common

import (
	"fmt"
	"sort"

	"golang.org/x/crypto/sha3"
)

type Block struct {
	Height       int64
	PreviousHash Hash
	CurrentHash  Hash
	Transactions Transactions
}

func (block *Block) sort() {
	sort.Sort(block.Transactions)
}

func (block *Block) Hash() {
	hash := sha3.New256()
	result := hash.Sum(nil)
	fmt.Println(string(result))
}
