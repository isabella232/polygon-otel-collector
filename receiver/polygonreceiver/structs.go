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
