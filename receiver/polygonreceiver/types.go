package polygonreceiver

type CheckpointSignatures struct {
	Success bool `json:"success"`
	Result  []struct {
		SignerAddress string `json:"signerAddress"`
		HasSigned     bool   `json:"hasSigned"`
	} `json:"result"`
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

type HeimdallSpan struct {
	Height string `json:"height"`
	Result struct {
		SpanID     int64 `json:"span_id"`
		StartBlock int64 `json:"start_block"`
		EndBlock   int64 `json:"end_block"`
	} `json:"result"`
}
