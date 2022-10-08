package vtb

import (
	"bytes"
	"encoding/json"
	"io"
)

type transferMaticRequest struct {
	FromPrivateKey string  `json:"fromPrivateKey"`
	ToPublicKey    string  `json:"toPublicKey"`
	Amount         float32 `json:"amount"`
}

type transferRubleRequest struct {
	FromPrivateKey string  `json:"fromPrivateKey"`
	ToPublicKey    string  `json:"toPublicKey"`
	Amount         float32 `json:"amount"`
}

type transferNftRequest struct {
	FromPrivateKey string `json:"fromPrivateKey"`
	ToPublicKey    string `json:"toPublicKey"`
	TokenId        int    `json:"tokenId"`
}

type transferResponse struct {
	TransactionHash string `json:"transactionHash"`
}

func (v *vtb) TransferMatic(fromPrivateKey string, toPublicKey string, amount float32) (string, error) {
	path := "/v1/transfers/matic"

	req := &transferMaticRequest{
		FromPrivateKey: fromPrivateKey,
		ToPublicKey:    toPublicKey,
		Amount:         amount,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := v.client.Post(path, "application/json", io.NopCloser(bytes.NewReader(b)))
	if err != nil {
		return "", err
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	j := new(transferResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return "", err
	}

	return j.TransactionHash, nil
}

func (v *vtb) TransferRuble(fromPrivateKey string, toPublicKey string, amount float32) (string, error) {
	path := "/v1/transfers/rubble"

	req := &transferRubleRequest{
		FromPrivateKey: fromPrivateKey,
		ToPublicKey:    toPublicKey,
		Amount:         amount,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := v.client.Post(path, "application/json", io.NopCloser(bytes.NewReader(b)))
	if err != nil {
		return "", err
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	j := new(transferResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return "", err
	}

	return j.TransactionHash, nil
}

func (v *vtb) TransferNft(fromPrivateKey string, toPublicKey string, tokenId int) (string, error) {
	path := "/v1/transfers/nft"

	req := &transferNftRequest{
		FromPrivateKey: fromPrivateKey,
		ToPublicKey:    toPublicKey,
		TokenId:        tokenId,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := v.client.Post(path, "application/json", io.NopCloser(bytes.NewReader(b)))
	if err != nil {
		return "", err
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	j := new(transferResponse)

	if err = json.Unmarshal(b, j); err != nil {
		return "", err
	}

	return j.TransactionHash, nil
}
