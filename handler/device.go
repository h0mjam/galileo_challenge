package handler

import (
	"encoding/json"
	"galileo/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// DeviceRegistry Регитрация устройства
func DeviceRegistry(devices types.DeviceStore, users types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("Authorization")
		if userToken == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		user, err := users.FindByToken(userToken)
		if err != nil {
			http.Error(w, "Unknown user", 400)
			return
		}

		d := &types.DeviceCreateRequest{}
		err = jsonRequest(r, d)

		if err != nil {
			logrus.Errorf("device registry invalid json: %v", err)
			http.Error(w, "Invalid json", 400)
			return
		}

		device := types.NewDevice()
		device.Name = d.Name
		devices.Add(device)

		user.AppendDevice(device)

		res, err := json.Marshal(&struct {
			Token string `json:"token"`
		}{
			Token: device.Token,
		})

		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		writeResponse(w, res, 201)
	}
}

func DeviceStat(store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		user, err := store.FindByToken(token)

		if err != nil {
			http.Error(w, "Unknown user", 422)
			return
		}

		id, err := strconv.Atoi(r.URL.Path[len("/device/stat/"):])

		if err != nil {
			http.Error(w, "Corrupted request", 422)
			return
		}

		device, err := user.FindDevice(uint64(id))

		if err != nil {
			http.Error(w, "Unknown device", 422)
			return
		}

		resp, err := json.Marshal(device.Measures)

		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}

		writeResponse(w, resp, 200)
	}
}
