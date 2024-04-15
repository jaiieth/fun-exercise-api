package wallet

import "time"

type Wallet struct {
	ID         int       `json:"id" example:"1"`
	UserID     int       `json:"user_id" example:"1" validate:"required,numeric"`
	UserName   string    `json:"user_name" example:"John Doe" validate:"required"`
	WalletName string    `json:"wallet_name" example:"John's Wallet" validate:"required"`
	WalletType string    `json:"wallet_type" example:"Create Card" validate:"required,oneof=Savings 'Credit Card' 'Crypto Wallet'"`
	Balance    float64   `json:"balance" example:"100.00" validate:"required"`
	CreatedAt  time.Time `json:"created_at" example:"2024-03-25T14:19:00.729237Z"`
}

const (
	Savings      = "Savings"
	CreditCard   = "Credit Card"
	CryptoWallet = "Crypto Wallet"
)

var AllowedWalletTypes = []string{
	Savings, CreditCard, CryptoWallet,
}
