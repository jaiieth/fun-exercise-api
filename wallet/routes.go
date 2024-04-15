package wallet

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRoutes(echo *echo.Echo) {
	v1 := echo.Group("/api/v1")
	v1.Add("GET", "/wallets", h.GetWallets)
	v1.Add("GET", "/users/:id/wallets", h.GetWalletsByUserID)
	v1.Add("POST", "/wallets", h.CreateWallet)
	v1.Add("PUT", "/wallets", h.UpdateWallet)
	v1.Add("DELETE", "/users/:id/wallets", h.DeleteWallet)
}
