package wallet

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	GetWallets(walletType string) ([]Wallet, error)
	GetWalletsById(id int) ([]Wallet, error)
	CreateWallet(wallet *Wallet) error
	UpdateWallet(wallet *Wallet) error
	DeleteWallet(id int) error
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}
type Success struct {
	Message string `json:"message"`
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

// @Summary		Get all wallets
// @Description	Get all wallets or wallets with given wallet type
// @Tags			wallet
// @Param			wallet_type	query	string	false	"wallet type"	Enums(Savings, Crypto Wallet, Credit Card)
// @Accept			json
// @Produce		json
// @Success		200	{array}	Wallet
// @Router			/api/v1/wallets [get]
// @Failure		500	{object}	Err
func (h *Handler) GetWallets(c echo.Context) error {
	var filter string
	if walletType := c.QueryParam("wallet_type"); walletType != "" {
		ruleTag := "required,oneof=Savings 'Credit Card' 'Crypto Wallet'"

		if err := validate.Var(walletType, ruleTag); err != nil {
			return c.JSON(http.StatusBadRequest, new(struct{}))
		}

		filter = walletType
	}

	wallets, err := h.store.GetWallets(filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "Unable to get wallets"})
	}
	return c.JSON(http.StatusOK, wallets)

}

// @Summary		Get all user wallets based on given id
// @Description	Get all wallets based on given id
// @Tags			wallet
// @Param			id	path	number	false	"User ID"
// @Accept			json
// @Produce		json
// @Success		200	{array}	Wallet
// @Router			/api/v1/users/{id}/wallets [get]
// @Failure		500	{object}	Err
// @Failure 		400 {object} 	Err
func (h *Handler) GetWalletsByUserID(c echo.Context) error {
	userId := c.Param("id")

	id, _ := strconv.Atoi(userId)

	ruleTag := "required,numeric"

	if err := validate.Var(id, ruleTag); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid user id"})
	}

	wallets, err := h.store.GetWalletsById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "Unable to get wallets"})
	}

	if len(wallets) == 0 {
		return c.JSON(http.StatusOK, []Wallet{})
	}

	return c.JSON(http.StatusOK, wallets)

}

// @Summary		Create a wallet
// @Description Create a wallet
// @Tags			wallet
// @Param 		wallet body CreateWalletBody true "Wallet"
// @Accept			json
// @Produce		json
// @Success		200	{array}	Wallet
// @Router			/api/v1/wallets [post]
// @Failure		500	{object}	Err
// @Failure 		400 {object} 	Err
func (h *Handler) CreateWallet(c echo.Context) (err error) {
	wallet := new(Wallet)
	if err = c.Bind(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid request body"})
	}

	if err = c.Validate(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid request body"})
	}

	err = h.store.CreateWallet(wallet)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "Unable to create wallet"})
	}

	return c.JSON(http.StatusCreated, wallet)

}

// @Summary		Update a wallet
// @Description Update a wallet
// @Tags			wallet
// @Param 		wallet body CreateWalletBody true "Wallet"
// @Accept			json
// @Produce		json
// @Success		200	{array}	Wallet
// @Router			/api/v1/wallets [post]
// @Failure		500	{object}	Err
// @Failure 		400 {object} 	Err
func (h *Handler) UpdateWallet(c echo.Context) (err error) {
	wallet := new(Wallet)

	if err = c.Bind(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid request body"})
	}

	if err = c.Validate(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid request body"})
	}

	err = h.store.UpdateWallet(wallet)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "Unable to update wallet"})
	}
	return c.JSON(http.StatusOK, wallet)

}

func (h *Handler) DeleteWallet(c echo.Context) (err error) {

	param := c.Param("id")

	id, _ := strconv.Atoi(param)

	ruleTag := "required,numeric"

	if err := validate.Var(id, ruleTag); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid user id"})
	}

	err = h.store.DeleteWallet(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "Unable to delete wallet"})
	}
	return c.JSON(http.StatusOK, Success{Message: "Delete wallet successfully"})

}
