package vtb

import (
	"net/http"
)

const baseUrl = "https://hackathon.lsp.team/hk"

type Vtb interface {
	Wallet
	Transfer
}

type Wallet interface {
	NewWallet() (string, string, error)
	WalletBalance(publicKey string) (float32, float32, error)
}

type Transfer interface {
	TransferMatic(fromPrivateKey string, toPublicKey string, amount float32) (string, error)
	TransferRuble(fromPrivateKey string, toPublicKey string, amount float32) (string, error)
	TransferNft(fromPrivateKey string, toPublicKey string, tokenId int) (string, error)
}

type vtb struct {
	client *http.Client
}

func NewVtb(client *http.Client) Vtb {
	return &vtb{
		client: client,
	}
}
