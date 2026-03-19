package handler

import (
	"net/http"

	appPayment "github.com/alves/age-of-genesis/internal/application/payment"
	"github.com/alves/age-of-genesis/internal/domain/payment"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	chargeService *appPayment.ChargeService
}

type CreditCardChargeRequest struct {
	UserID           int64          `json:"user_id" binding:"required"`
	StoreCartID      int64          `json:"store_carts_id" binding:"required"`
	CreditCardHolder string         `json:"credit_card_holder" binding:"required"`
	CPFForCard       string         `json:"cpf_for_card" binding:"required"`
	EncryptedCard    string         `json:"encrypted_card" binding:"required"`
	AmountInCents    int64          `json:"amount" binding:"required"`
	Description      string         `json:"description"`
	CustomerEmail    string         `json:"customer_email" binding:"required,email"`
	NotificationURL  string         `json:"notification_url" binding:"required,url"`
	Items            []payment.Item `json:"items" binding:"required,dive"`
}

func NewPaymentHandler(chargeService *appPayment.ChargeService) *PaymentHandler {
	return &PaymentHandler{chargeService: chargeService}
}

func (h *PaymentHandler) ChargeCreditCard(c *gin.Context) {
	var req CreditCardChargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := h.chargeService.Execute(c.Request.Context(), payment.CreditChargeInput{
		UserID:           req.UserID,
		StoreCartID:      req.StoreCartID,
		CreditCardHolder: req.CreditCardHolder,
		CPFForCard:       req.CPFForCard,
		EncryptedCard:    req.EncryptedCard,
		AmountInCents:    req.AmountInCents,
		Description:      req.Description,
		CustomerEmail:    req.CustomerEmail,
		NotificationURL:  req.NotificationURL,
		Items:            req.Items,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, out)
}
