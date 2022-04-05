package polygonreceiver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HeimdallClient struct {
	apiUrl string
}

func NewHeimdallClient(apiUrl string) *HeimdallClient {
	return &HeimdallClient{
		apiUrl: apiUrl,
	}
}

func (h *HeimdallClient) Block(height *string) (*HeimdallBlock, error) {
	var url string

	if height != nil {
		url = fmt.Sprintf("%s/block?height=%s", h.apiUrl, *height)
	} else {
		url = h.apiUrl + "/block"
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
