package transaction

import (
	"strings"

	"../"
)

type Transactions []*Transaction

type Transaction struct {
	From      common.Address
	To        common.Address
	TimeStamp common.Time // UnixTimeStamp
	Amount    int64
	Type      common.CoinType
	Nonce     int64
	Sign      string
}

func (transactions Transactions) Len() int {
	return len(transactions)
}

func (transactions Transactions) Less(i, j int) bool {
	return strings.Compare(transactions[i].Sign, transactions[j].Sign) > 0
}

func (transactions Transactions) Swap(i, j int) {
	transactions[i], transactions[j] = transactions[j], transactions[i]
}

func (transactions Transactions) Hash() string {
	return ""
}
