package payment

import "context"

type Gateway interface {
	Charge(ctx context.Context, req ChargeRequest) (*ChargeResponse, error)
}
