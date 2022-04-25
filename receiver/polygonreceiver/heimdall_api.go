package polygonreceiver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HeimdallClient struct {
	tendermintApiUrl string
	heimdallApiUrl   string
}

func NewHeimdallClient(tendermintApiUrl string, heimdallApiUrl string) *HeimdallClient {
	return &HeimdallClient{
		tendermintApiUrl: tendermintApiUrl,
		heimdallApiUrl:   heimdallApiUrl,
	}
}

func (h *HeimdallClient) Block(height *string) (*HeimdallBlock, error) {
	var url string

	if height != nil {
		url = fmt.Sprintf("%s/block?height=%s", h.tendermintApiUrl, *height)
	} else {
		url = h.tendermintApiUrl + "/block"
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	block := &HeimdallBlock{}
	err = json.Unmarshal(body, &block)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (h *HeimdallClient) LatestSpan() (*HeimdallSpan, error) {
	res, err := http.Get(h.heimdallApiUrl + "/bor/latest-span")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	span := &HeimdallSpan{}
	err = json.Unmarshal(body, &span)
	if err != nil {
		return nil, err
	}

	return span, nil
}

func (h *HeimdallClient) UnconfirmedTransactions() (*HeimdallUnconfirmedTransactions, error) {
	res, err := http.Get(h.tendermintApiUrl + "/num_unconfirmed_txs")
	if err != nil {
		return nil, fmt.Errorf("failed to get unconfirmed transactions: %w", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read unconfirmed transactions: %w", err)
	}

	transactions := &HeimdallUnconfirmedTransactions{}
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal unconfirmed transactions: %w", err)
	}

	return transactions, nil
}
