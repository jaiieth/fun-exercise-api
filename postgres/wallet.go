package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) GetWallets(walletType string) ([]wallet.Wallet, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if walletType == "" {
		rows, err = p.Db.Query("SELECT * FROM user_wallet")
	} else {
		rows, err = p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type = $1", walletType)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) GetWalletsById(userID int) ([]wallet.Wallet, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = p.Db.Query("SELECT * FROM user_wallet WHERE user_id = $1", userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) CreateWallet(wallet *wallet.Wallet) error {
	args := []interface{}{wallet.UserID, wallet.UserName, wallet.WalletName, wallet.WalletType, wallet.Balance}

	stmt := `INSERT INTO 
	user_wallet (user_id, user_name, wallet_name, wallet_type, balance) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id , created_at`

	err := p.Db.QueryRow(stmt, args...).Scan(&wallet.ID, &wallet.CreatedAt)

	if err != nil {
		return err
	}
	return nil

}

func (p *Postgres) UpdateWallet(wallet *wallet.Wallet) error {
	args := []interface{}{wallet.UserID, wallet.UserName, wallet.WalletName, wallet.WalletType, wallet.Balance, wallet.ID}

	stmt := `UPDATE user_wallet
	SET user_id = $1, user_name = $2, wallet_name = $3, wallet_type = $4, balance = $5
	WHERE id = $6
	RETURNING created_at`

	// _, err := p.Db.Exec(stmt, args...)
	err := p.Db.QueryRow(stmt, args...).Scan(&wallet.CreatedAt)

	if err != nil {
		fmt.Println("🚀 | file: wallet.go | line 127 | func | err : ", err)
		return err
	}

	return nil
}

func (p *Postgres) DeleteWallet(userID int) error {

	stmt := `DELETE FROM user_wallet WHERE user_id = $1`
	_, err := p.Db.Exec(stmt, userID)

	if err != nil {
		return err
	}
	return nil
}
