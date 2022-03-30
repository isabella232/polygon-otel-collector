package polygonreceiver

type Result struct {
	SignerAddress string `json:"signerAddress"`
	HasSigned     bool   `json:"hasSigned"`
}
type CheckpointSignatures struct {
	Success bool     `json:"success"`
	Result  []Result `json:"result"`
}

type HeimdallUnconfirmedTransactions struct {
	Result struct {
		Ntxs  string `json:"n_txs"`
		Total string `json:"total"`
	} `json:"result"`
}

type HeimdallBlock struct {
	Result struct {
		Block struct {
			Header struct {
				Time   string `json:"time"`
				Height string `json:"height"`
			} `json:"header"`
		} `json:"block"`
	} `json:"result"`
}
