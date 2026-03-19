package payment

type CreditChargeInput struct {
	UserID           int64
	StoreCartID      int64
	CreditCardHolder string
	CPFForCard       string
	EncryptedCard    string
	AmountInCents    int64
	Description      string
	CustomerEmail    string
	Items            []Item
	NotificationURL  string
}

type Item struct {
	ReferenceID string `json:"reference_id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	UnitAmount  int64  `json:"unit_amount"`
}
