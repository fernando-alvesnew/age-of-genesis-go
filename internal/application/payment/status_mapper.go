package payment

import domainPayment "github.com/alves/age-of-genesis/internal/domain/payment"

type StatusMapper interface {
	Map(gatewayStatus string) string
}

type DefaultStatusMapper struct{}

func NewDefaultStatusMapper() *DefaultStatusMapper {
	return &DefaultStatusMapper{}
}

func (m *DefaultStatusMapper) Map(gatewayStatus string) string {
	switch gatewayStatus {
	case domainPayment.StatusPaid:
		return domainPayment.StatusWaiting
	default:
		return gatewayStatus
	}
}
