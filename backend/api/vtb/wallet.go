package vtb

import (
	"encoding/json"
	"fmt"
	"io"
)

type NewWalletResponse struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type WalletBalanceResponse struct {
	MaticAmount float32 `json:"maticAmount"`
	CoinsAmount float32 `json:"coinsAmount"`
}

func (v *vtb) NewWallet() (string, string, error) {
	path := "/v1/wallets/new"
	resp, err := v.client.Post(baseUrl+path, "application/json", nil)
	if err != nil {
		return "", "", err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	j := new(NewWalletResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return "", "", err
	}

	return j.PublicKey, j.PrivateKey, nil
}

func (v *vtb) WalletBalance(publicKey string) (float32, float32, error) {
	path := fmt.Sprintf("/v1/wallets/%s/balance", publicKey)
	resp, err := v.client.Post(baseUrl+path, "application/json", nil)
	if err != nil {
		return 0, 0, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	j := new(WalletBalanceResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return 0, 0, err
	}

	return j.MaticAmount, j.CoinsAmount, nil
}
