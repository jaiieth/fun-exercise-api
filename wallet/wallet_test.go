package wallet_test

<<<<<<< Updated upstream
import "testing"

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {

	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {

	})
=======
import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KKGo-Software-engineering/fun-exercise-api/helper"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type stubStorer struct {
	wallets []wallet.Wallet
	err     error
	mock.Mock
}

func (s *stubStorer) GetWallets(walletType string) ([]wallet.Wallet, error) {
	s.Called(walletType)
	return s.wallets, s.err
}

func (s *stubStorer) GetWalletsById(id int) ([]wallet.Wallet, error) {
	s.Called(id)
	return s.wallets, s.err
}

func (s *stubStorer) CreateWallet(wallet *wallet.Wallet) error {
	s.Called(wallet)
	return s.err
}
func (s *stubStorer) UpdateWallet(wallet *wallet.Wallet) error {
	s.Called(wallet)
	return s.err
}

func (s *stubStorer) DeleteWallet(id int) error {
	s.Called(id)
	return s.err
}

func NewContext(method string, target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	rec := httptest.NewRecorder()

	e.Validator = helper.NewValidator()

	context := e.NewContext(req, rec)
	return context, rec
}

// func TestWallet(t *testing.T) {
// 	t.Run("given unable to get wallets should return 500", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
// 		rec := httptest.NewRecorder()

// 		c := e.NewContext(req, rec)

// 		mockStore := &stubStorer{
// 			err: echo.ErrInternalServerError,
// 		}

// 		mockStore.On("GetWallets", "").Return()

// 		stubHandler := wallet.New(mockStore)

// 		stubHandler.GetWallets(c)

// 		if rec.Code != http.StatusInternalServerError {
// 			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
// 		}

// 	})

// 	t.Run("given user able to getting wallet should return 200 and list of wallets", func(t *testing.T) {
// 		c, rec := NewContext(http.MethodGet, "/api/v1/wallets", nil)

// 		want := []wallet.Wallet{
// 			{ID: 1, UserID: 1},
// 			{ID: 2, UserID: 2},
// 			{ID: 3, UserID: 3},
// 		}

// 		mockStore := &stubStorer{
// 			wallets: want,
// 		}

// 		mockStore.On("GetWallets", "").Return()

// 		stubHandler := wallet.New(mockStore)

// 		if assert.NoError(t, stubHandler.GetWallets(c)) {
// 			assert.Equal(t, http.StatusOK, rec.Code)

// 			var got []wallet.Wallet
// 			if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
// 				t.Errorf("Expected response to be JSON format")
// 			}

// 			expectedLegth := len(want)
// 			assert.Len(t, got, expectedLegth)
// 		}
// 		mockStore.AssertCalled(t, "GetWallets", "")

// 	})

// 	t.Run("Given wrong wallet_type should return 404", func(t *testing.T) {
// 		var wantFilter = "InvalidType"
// 		c, rec := NewContext(http.MethodGet, fmt.Sprintf("/api/v1/wallets?wallet_type=%s", wantFilter), nil)

// 		// Create a mock store
// 		mockStore := &stubStorer{}

// 		// Expected Wallets() call
// 		mockStore.On("GetWallets").Return()

// 		// Create a handler
// 		stubHandler := wallet.New(mockStore)

// 		if assert.NoError(t, stubHandler.GetWallets(c)) {
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 		}

// 	})
// 	t.Run("Given wallet_type within AllowedWalletTypes type should return 200 and list of wallets, called Wallets() with wallet_type", func(t *testing.T) {
// 		var wantFilter = wallet.Savings
// 		c, rec := NewContext(http.MethodGet, fmt.Sprintf("/api/v1/wallets?wallet_type=%s", wantFilter), nil)

// 		want := []wallet.Wallet{
// 			{ID: 1, UserID: 1},
// 			{ID: 2, UserID: 2},
// 		}
// 		// Create a mock store
// 		mockStore := &stubStorer{
// 			wallets: want,
// 		}

// 		// Expected Wallets() call
// 		mockStore.On("GetWallets", wantFilter).Return()

// 		// Create a handler
// 		stubHandler := wallet.New(mockStore)

// 		if assert.NoError(t, stubHandler.GetWallets(c)) {

// 			var got []wallet.Wallet
// 			if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
// 				t.Errorf("Expected response to be JSON format")
// 			}

// 			expectedLegth := len(want)
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			assert.Len(t, got, expectedLegth)
// 		}

// 		// Assert Wallets() called with wantFilter
// 		mockStore.AssertCalled(t, "GetWallets", wantFilter)

// 	})

// }

// func TestGetUserWallet(t *testing.T) {
// 	t.Run("Given invalid id should return 400", func(t *testing.T) {
// 		targetId := "0"
// 		c, rec := NewContext(http.MethodGet, fmt.Sprintf("/api/v1/users/%v/wallets", targetId), nil)
// 		c.SetPath("/api/v1/users/:id/wallets")
// 		c.SetParamNames("id")
// 		c.SetParamValues(targetId)

// 		stubStore := &stubStorer{
// 			wallets: []wallet.Wallet{
// 				{ID: 1, UserID: 1},
// 				{ID: 1, UserID: 2},
// 			},
// 		}

// 		stubStore.On("GetWalletsById", "").Return()

// 		stubHandler := wallet.New(stubStore)

// 		err := stubHandler.GetWalletsByUserID(c)

// 		if err != nil {
// 			t.Errorf("Unexpected error %v", err)
// 		}

// 		assert.Equal(t, rec.Code, http.StatusBadRequest)

// 	})

// 	t.Run("Given string id should return 400", func(t *testing.T) {
// 		targetId := "asd"
// 		c, rec := NewContext(http.MethodGet, fmt.Sprintf("/api/v1/users/%v/wallets", targetId), nil)

// 		stubStore := &stubStorer{
// 			wallets: []wallet.Wallet{
// 				{ID: 1, UserID: 1},
// 				{ID: 1, UserID: 2},
// 			},
// 		}

// 		stubStore.On("GetWalletsById", targetId).Return()

// 		stubHandler := wallet.New(stubStore)

// 		err := stubHandler.GetWalletsByUserID(c)

// 		if err != nil {
// 			t.Errorf("Unexpected error %v", err)
// 		}

// 		assert.Equal(t, rec.Code, http.StatusBadRequest)

// 	})

// 	t.Run("Given valid id should calls GetWalletsById with id, return 200 and json", func(t *testing.T) {
// 		targetId := "1"
// 		c, rec := NewContext(http.MethodGet, "/", nil)
// 		c.SetPath("/api/v1/users/:id/wallets")
// 		c.SetParamNames("id")
// 		c.SetParamValues(targetId)

// 		stubStore := &stubStorer{
// 			wallets: []wallet.Wallet{{}, {}},
// 		}

// 		stubStore.On("GetWalletsById", 1).Return()

// 		stubHandler := wallet.New(stubStore)

// 		err := stubHandler.GetWalletsByUserID(c)

// 		if err != nil {
// 			t.Errorf("Unexpected error %v", err)
// 		}

// 		var got []wallet.Wallet
// 		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
// 			t.Errorf("Expected response to be JSON format")
// 		}

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		stubStore.AssertCalled(t, "GetWalletsById", 1)

// 	})

// 	t.Run("Given valid id should calls GetWalletsById with id, return 500 if GetWalletsById return err", func(t *testing.T) {
// 		targetId := "1"
// 		c, rec := NewContext(http.MethodGet, "/", nil)
// 		c.SetPath("/api/v1/users/:id/wallets")
// 		c.SetParamNames("id")
// 		c.SetParamValues(targetId)

// 		stubStore := &stubStorer{
// 			wallets: []wallet.Wallet{{}, {}},
// 			err:     echo.ErrInternalServerError,
// 		}

// 		stubStore.On("GetWalletsById", 1).Return()

// 		stubHandler := wallet.New(stubStore)

// 		err := stubHandler.GetWalletsByUserID(c)

// 		if err != nil {
// 			t.Errorf("Unexpected error %v", err)
// 		}

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 		stubStore.AssertCalled(t, "GetWalletsById", 1)
// 	})

// }

// func TestCreateWallet(t *testing.T) {
// 	t.Run("Given invalid body, should return 400", func(t *testing.T) {

// 		invalidBody := "{not valid}"

// 		c, rec := NewContext(http.MethodPost, "/api/v1/wallets", bytes.NewReader([]byte(invalidBody)))

// 		stubStorer := &stubStorer{}
// 		stubHandler := wallet.New(stubStorer)

// 		stubStorer.On("CreateWallet", &wallet.Wallet{}).Return()

// 		stubHandler.CreateWallet(c)

// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	})

// 	// {
// 	// 		"UserID":     99,
// 	// 		"UserName":   "Tester",
// 	// 		"WalletName": "Test Wallet",
// 	// 		"WalletType": "Savings",
// 	// 		"Balance":    500,
// 	// 	}
// 	t.Run("Given valid body, but missing some field should return 400", func(t *testing.T) {

// 		want := wallet.Wallet{}

// 		body, err := json.Marshal(want)

// 		if err != nil {
// 			t.Errorf("Unexpected error: %v", err)
// 		}

// 		c, rec := NewContext(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(body))

// 		c.Request().Header.Set("Content-Type", "application/json")

// 		stubStorer := &stubStorer{}

// 		stubHandler := wallet.New(stubStorer)

// 		stubStorer.On("CreateWallet", &want).Return()

// 		stubHandler.CreateWallet(c)

// 		// assert
// 		assert.Equal(t, http.StatusBadRequest, rec.Code)
// 	})
// 	t.Run("Given valid body, got error from CreateWallet() should return 500 ", func(t *testing.T) {

// 		want := wallet.Wallet{
// 			UserID:     99,
// 			UserName:   "Tester",
// 			WalletName: "Test Wallet",
// 			WalletType: "Savings",
// 			Balance:    500,
// 		}

// 		body, err := json.Marshal(want)

// 		if err != nil {
// 			t.Errorf("Unexpected error: %v", err)
// 		}

// 		c, rec := NewContext(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(body))

// 		c.Request().Header.Set("Content-Type", "application/json")

// 		stubStorer := &stubStorer{
// 			err: echo.ErrInternalServerError,
// 		}

// 		stubHandler := wallet.New(stubStorer)

// 		stubStorer.On("CreateWallet", &want).Return()

// 		stubHandler.CreateWallet(c)
// 		// assert
// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 		stubStorer.AssertCalled(t, "CreateWallet", &want)
// 	})
// 	t.Run("Given valid body, should return 201 and wallet with id", func(t *testing.T) {

// 		want := wallet.Wallet{
// 			UserID:     99,
// 			UserName:   "Tester",
// 			WalletName: "Test Wallet",
// 			WalletType: "Savings",
// 			Balance:    500,
// 		}

// 		body, err := json.Marshal(want)

// 		if err != nil {
// 			t.Errorf("Unexpected error: %v", err)
// 		}

// 		c, rec := NewContext(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(body))

// 		c.Request().Header.Set("Content-Type", "application/json")

// 		stubStorer := &stubStorer{}

// 		stubHandler := wallet.New(stubStorer)

// 		stubStorer.On("CreateWallet", &want).Return()

// 		stubHandler.CreateWallet(c)

// 		var got wallet.Wallet
// 		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
// 			t.Errorf("Unexpected error: %v", err)
// 		}

// 		// assert
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 		assert.Equal(t, got.UserID, 99)
// 		assert.Equal(t, got.UserName, "Tester")
// 		assert.Equal(t, got.WalletName, "Test Wallet")
// 		assert.Equal(t, got.WalletType, "Savings")
// 		assert.EqualValues(t, got.Balance, 500)
// 		stubStorer.AssertCalled(t, "CreateWallet", &want)
// 	})
// }

func TestUpdateWallet(t *testing.T) {
	t.Run("Given valid body, should return 200, and success message", func(t *testing.T) {
		want := wallet.Wallet{
			ID:         10,
			UserID:     10,
			UserName:   "TestUpdate",
			WalletName: "UpdateWalletName",
			WalletType: wallet.AllowedWalletTypes[1],
			Balance:    999,
		}

		body, err := json.Marshal(want)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		c, rec := NewContext("PUT", "/api/v1/wallets", bytes.NewBuffer(body))
		c.Request().Header.Set("Content-Type", "application/json")

		stubStorer := &stubStorer{}
		stubStorer.On("UpdateWallet", &want).Return()

		stubHandler := wallet.New(stubStorer)
		stubHandler.UpdateWallet(c)

		var got wallet.Success
		json.Unmarshal(rec.Body.Bytes(), &got)
		stubStorer.AssertCalled(t, "UpdateWallet", &want)
		assert.Equal(t, got.Message, "Wallet updated successfully")
		assert.Equal(t, rec.Code, http.StatusOK)
	})

>>>>>>> Stashed changes
}
