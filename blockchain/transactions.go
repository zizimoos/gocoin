package blockchain

import (
	"errors"
	"time"

	"github.com/zizimoos/gocoin/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

type Tx struct {
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.ID = utils.Hash(t)
}

type TxIn struct {
	TxID  string `json:"txId"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type UTxOut struct {
	TxID   string `json:"txId"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			if input.TxID == uTxOut.TxID && input.Index == uTxOut.Index {
				exists = true
			}
		}
	}
	return exists
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "Coinbase"},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough coin")
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	UTxOuts := Blockchain().UTxOutsByAddress(from)
	for _, uTxOut := range UTxOuts {
		if total > amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("zizi", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("zizi")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
