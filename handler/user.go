package handler

import (
	"encoding/json"
	"galileo/store"
	"galileo/types"
	"net/http"
)

func UserRegistry(s types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &types.UserCreateRequest{}
		err := jsonRequest(r, req)

		if err != nil {
			http.Error(w, "Corrupted request: "+err.Error(), 400)
			return
		}

		user, err := s.FindByEmail(req.Email)

		// Если пользователя с таким e-mail нет, создаем нового
		if err != nil {
			user = types.NewUser()
			user.Devices = store.NewDeviceStore()
			user.Email = req.Email

			s.Add(user)
		}

		body, err := json.Marshal(&struct {
			Token string `json:"token"`
		}{
			Token: user.Token,
		})

		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		writeResponse(w, body, 201)
	}
}

func UserDevices(store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("Authorization")

		if userToken == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		user, err := store.FindByToken(userToken)

		if err != nil {
			http.Error(w, "Unknown user", 401)
			return
		}

		body, err := json.Marshal([]*types.Device(user.Devices.All()))

		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}

		writeResponse(w, body, 200)
	}
}
