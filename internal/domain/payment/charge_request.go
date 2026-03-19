package payment

type ChargeRequest struct {
	ReferenceID      string           `json:"reference_id"`
	Customer         ChargeCustomer   `json:"customer"`
	Items            []Item           `json:"items"`
	NotificationURLs []string         `json:"notification_urls"`
	Charges          []ChargeEnvelope `json:"charges"`
}

type ChargeCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	TaxID string `json:"tax_id"`
}

type ChargeEnvelope struct {
	ReferenceID  string        `json:"reference_id"`
	Description  string        `json:"description"`
	Amount       ChargeAmount  `json:"amount"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}

type ChargeAmount struct {
	Value    int64  `json:"value"`
	Currency string `json:"currency"`
}

type PaymentMethod struct {
	Type         string     `json:"type"`
	Installments int        `json:"installments"`
	Capture      bool       `json:"capture"`
	Card         CardData   `json:"card"`
	Holder       CardHolder `json:"holder"`
}

type CardData struct {
	Encrypted string `json:"encrypted"`
	Store     bool   `json:"store"`
}

type CardHolder struct {
	Name  string `json:"name"`
	TaxID string `json:"tax_id"`
}
