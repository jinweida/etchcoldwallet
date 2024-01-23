package entity

type SendTransaction struct {
	from     string `json:"from"`
	nonce    string `json:"nonce"`
	gasPrice string `json:"gas_price"`
	gasLimit string `json:"gas_limit"`
	to       string `json:"to"`
	value    string `json:"value"`
	data     string `json:"data"`
	chainId  string `json:"chain_id"`
	tx       string `json:"tx"`
}
