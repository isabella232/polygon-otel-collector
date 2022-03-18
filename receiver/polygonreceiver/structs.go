package polygonreceiver

type Result struct {
	SignerAddress string `json:"signerAddress"`
	ValidatorID   string `json:"validatorId"`
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
