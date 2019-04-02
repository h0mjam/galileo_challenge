package handler

import (
	"galileo/store"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserRegistry(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost/user", strings.NewReader(`{"email":"tester@tester.com"}`))
	s := store.NewUserStore()
	h := UserRegistry(s)

	rr := httptest.NewRecorder()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	h.ServeHTTP(rr, req)

	u, err := s.FindByEmail("tester@tester.com")
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"token":"` + u.Token + `"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDeviceRegistry(t *testing.T) {
	dStore := store.NewDeviceStore()
	uStore := store.NewUserStore()
	u := store.NewUser()
	uStore.Add(u)
	h := DeviceRegistry(dStore, uStore)

	req := httptest.NewRequest(http.MethodPost, "http://localhost/device", strings.NewReader(`{"name": "CR-111-4","email": "tester@tester.ru"}`))
	req.Header.Add("Authorization", u.Token)

	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	d, err := u.FindDevice(1)

	assert.NoError(t, err)

	expected := `{"token":"` + d.Token + `"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
