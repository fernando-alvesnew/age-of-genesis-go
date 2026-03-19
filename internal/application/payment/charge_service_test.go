package payment

import (
	"context"
	"testing"

	"github.com/alves/age-of-genesis/internal/domain/payment"
)

type gatewayStub struct{}

func (g *gatewayStub) Charge(context.Context, payment.ChargeRequest) (*payment.ChargeResponse, error) {
	return &payment.ChargeResponse{
		ID: "order_123",
		Charges: []payment.ChargeResult{
			{
				ReferenceID: "ref_1",
				Status:      payment.StatusPaid,
				Amount:      payment.ChargeResultAmount{Value: 1000},
			},
		},
	}, nil
}

type txRepoStub struct {
	last payment.Transaction
}

func (s *txRepoStub) UpsertByReferenceID(_ context.Context, tx payment.Transaction) error {
	s.last = tx
	return nil
}

func TestChargeService_ExecuteSuccess(t *testing.T) {
	repo := &txRepoStub{}
	svc := NewChargeService(&gatewayStub{}, repo)

	out, err := svc.Execute(context.Background(), payment.CreditChargeInput{
		UserID:           1,
		StoreCartID:      10,
		CreditCardHolder: "John Doe",
		CPFForCard:       "12345678910",
		EncryptedCard:    "encrypted_data",
		AmountInCents:    1000,
		Description:      "Compra teste",
		CustomerEmail:    "john@example.com",
		NotificationURL:  "https://example.com/api/payment-notification",
		Items: []payment.Item{
			{ReferenceID: "1", Name: "Produto", Quantity: 1, UnitAmount: 1000},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Status != payment.StatusWaiting {
		t.Fatalf("expected waiting status, got %s", out.Status)
	}
}

type gatewayEmptyChargesStub struct{}

func (g *gatewayEmptyChargesStub) Charge(context.Context, payment.ChargeRequest) (*payment.ChargeResponse, error) {
	return &payment.ChargeResponse{
		ID:      "order_empty",
		Charges: []payment.ChargeResult{},
	}, nil
}

func TestChargeService_ExecuteReturnsErrorWhenGatewayHasNoCharges(t *testing.T) {
	repo := &txRepoStub{}
	svc := NewChargeService(&gatewayEmptyChargesStub{}, repo)

	_, err := svc.Execute(context.Background(), payment.CreditChargeInput{
		UserID:           1,
		StoreCartID:      10,
		CreditCardHolder: "John Doe",
		CPFForCard:       "12345678910",
		EncryptedCard:    "encrypted_data",
		AmountInCents:    1000,
		Description:      "Compra teste",
		CustomerEmail:    "john@example.com",
		NotificationURL:  "https://example.com/api/payment-notification",
		Items: []payment.Item{
			{ReferenceID: "1", Name: "Produto", Quantity: 1, UnitAmount: 1000},
		},
	})
	if err == nil {
		t.Fatalf("expected error")
	}
	if err != ErrEmptyGatewayCharges {
		t.Fatalf("expected ErrEmptyGatewayCharges, got %v", err)
	}
}
