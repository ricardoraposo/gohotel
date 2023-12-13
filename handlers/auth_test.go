package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAuthenticate(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/", authHandler.Authenticate)

	params := AuthParams{
		Email:    "rick@renner.com",
		Password: "superSecretFodase",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected http status code of 200 but got %d", resp.StatusCode)
	}
	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
