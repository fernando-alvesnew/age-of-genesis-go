package payment

type ChargeResponse struct {
	ID            string               `json:"id"`
	ErrorMessages []map[string]string `json:"error_messages"`
	Charges       []ChargeResult       `json:"charges"`
}

type ChargeResult struct {
	ReferenceID string             `json:"reference_id"`
	Status      string             `json:"status"`
	Amount      ChargeResultAmount `json:"amount"`
}

type ChargeResultAmount struct {
	Value int64 `json:"value"`
}
