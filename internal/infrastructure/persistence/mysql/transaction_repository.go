package mysql

import (
	"context"
	"database/sql"

	"github.com/alves/age-of-genesis/internal/domain/payment"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) UpsertByReferenceID(ctx context.Context, tx payment.Transaction) error {
	query := `
		INSERT INTO pagseguro_credit_card (
			store_carts_id, users_id, payment_id, reference_id, amount, status, description, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE
			payment_id = VALUES(payment_id),
			amount = VALUES(amount),
			status = VALUES(status),
			description = VALUES(description),
			updated_at = NOW()
	`
	_, err := r.db.ExecContext(
		ctx, query,
		tx.StoreCartID, tx.UserID, tx.PaymentID, tx.ReferenceID, tx.Amount, tx.Status, tx.Description,
	)
	return err
}
