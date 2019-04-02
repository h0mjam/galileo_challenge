package handler

import (
	"fmt"
	"galileo/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserRegistry(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost/user", strings.NewReader(`{"email":"tester@tester.com"}`))
	rr := httptest.NewRecorder()

	s := store.NewUserStore()
	h := UserRegistry(s)
	h.ServeHTTP(rr, req)

	u, err := s.FindByEmail("tester@tester.com")

	if err != nil {
		t.Fatal(err)
	}

	expected := `{"token":"` + u.Token + `"}`

	assert.Equal(t, rr.Code, http.StatusCreated, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))
	assert.Equal(t, rr.Body.String(), expected, fmt.Sprintf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected))
}

func TestDeviceRegistry(t *testing.T) {
	u := store.NewUser()
	dStore := store.NewDeviceStore()
	uStore := store.NewUserStore()
	uStore.Add(u)

	req := httptest.NewRequest(http.MethodPost, "http://localhost/device", strings.NewReader(`{"name": "CR-111-4","email": "tester@tester.ru"}`))
	req.Header.Add("Authorization", u.Token)

	rr := httptest.NewRecorder()

	h := DeviceRegistry(dStore, uStore)
	h.ServeHTTP(rr, req)

	d, err := u.FindDevice(1)

	if err != nil {
		t.Fatal(err)
	}

	expected := `{"token":"` + d.Token + `"}`

	assert.Equal(t, rr.Code, http.StatusCreated, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))
	assert.Equal(t, rr.Body.String(), expected, fmt.Sprintf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected))
}
