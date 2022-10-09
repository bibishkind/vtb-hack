package vtb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type NewWalletResponse struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type WalletBalanceResponse struct {
	MaticAmount float32 `json:"maticAmount"`
	CoinsAmount float32 `json:"coinsAmount"`
}

func (v *vtb) CreateWallet() (string, string, error) {
	path := "/v1/wallets/new"
	resp, err := v.client.Post(baseUrl+path, "application/json", nil)
	if err != nil {
		return "", "", err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", errors.New(string(b))
	}

	j := new(NewWalletResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return "", "", err
	}

	return j.PublicKey, j.PrivateKey, nil
}

func (v *vtb) GetBalance(publicKey string) (float32, float32, error) {
	path := fmt.Sprintf("/v1/wallets/%s/balance", publicKey)
	resp, err := v.client.Get(baseUrl + path)
	if err != nil {
		return 0, 0, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, 0, errors.New(string(b))
	}

	j := new(WalletBalanceResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return 0, 0, err
	}

	return j.MaticAmount, j.CoinsAmount, nil
}
