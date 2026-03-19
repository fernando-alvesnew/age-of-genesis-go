package payment

import (
	"context"
	"errors"

	"github.com/alves/age-of-genesis/internal/domain/payment"
	"github.com/google/uuid"
)

var ErrInvalidEncryptedCard = errors.New("encrypted_card is required")
var ErrEmptyGatewayCharges = errors.New("gateway returned no charges")

type TransactionRepository interface {
	UpsertByReferenceID(ctx context.Context, tx payment.Transaction) error
}

type ChargeService struct {
	gateway      payment.Gateway
	txRepo       TransactionRepository
	statusMapper StatusMapper
}

func NewChargeService(gateway payment.Gateway, txRepo TransactionRepository) *ChargeService {
	return &ChargeService{
		gateway:      gateway,
		txRepo:       txRepo,
		statusMapper: NewDefaultStatusMapper(),
	}
}

func (s *ChargeService) Execute(ctx context.Context, in payment.CreditChargeInput) (*payment.Transaction, error) {
	if in.EncryptedCard == "" {
		return nil, ErrInvalidEncryptedCard
	}

	referenceID := uuid.NewString()
	req := buildChargeRequest(in, referenceID)

	resp, err := s.gateway.Charge(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resp.ErrorMessages) > 0 {
		return nil, errors.New("pagseguro returned charge errors")
	}
	if len(resp.Charges) == 0 {
		return nil, ErrEmptyGatewayCharges
	}

	status := s.statusMapper.Map(resp.Charges[0].Status)

	tx := payment.Transaction{
		StoreCartID: in.StoreCartID,
		UserID:      in.UserID,
		PaymentID:   resp.ID,
		ReferenceID: resp.Charges[0].ReferenceID,
		Amount:      resp.Charges[0].Amount.Value,
		Status:      status,
		Description: in.Description,
	}

	if err := s.txRepo.UpsertByReferenceID(ctx, tx); err != nil {
		return nil, err
	}

	return &tx, nil
}

func buildChargeRequest(in payment.CreditChargeInput, referenceID string) payment.ChargeRequest {
	return payment.ChargeRequest{
		ReferenceID: referenceID,
		Customer: payment.ChargeCustomer{
			Name:  in.CreditCardHolder,
			Email: in.CustomerEmail,
			TaxID: in.CPFForCard,
		},
		Items:            in.Items,
		NotificationURLs: []string{in.NotificationURL},
		Charges: []payment.ChargeEnvelope{
			{
				ReferenceID: referenceID,
				Description: in.Description,
				Amount: payment.ChargeAmount{
					Value:    in.AmountInCents,
					Currency: "BRL",
				},
				PaymentMethod: payment.PaymentMethod{
					Type:         payment.PaymentMethodCreditCard,
					Installments: 1,
					Capture:      true,
					Card: payment.CardData{
						Encrypted: in.EncryptedCard,
						Store:     false,
					},
					Holder: payment.CardHolder{
						Name:  in.CreditCardHolder,
						TaxID: in.CPFForCard,
					},
				},
			},
		},
	}
}
