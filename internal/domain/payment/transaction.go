package payment

type Transaction struct {
	ID          int64
	StoreCartID int64
	UserID      int64
	PaymentID   string
	ReferenceID string
	Amount      int64
	Status      string
	Description string
}

const (
	PaymentMethodCreditCard = "CREDIT_CARD"
	StatusPaid              = "PAID"
	StatusWaiting           = "WAITING"
)
