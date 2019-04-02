package handler

import (
	"fmt"
	"galileo/types"
	"github.com/sirupsen/logrus"
	"net/http"
)

func MetricsAppend(store types.DeviceStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		measures := make([]types.Measure, 0)
		err := jsonRequest(r, &measures)

		if err != nil {
			http.Error(w, fmt.Sprintf("Corrupted request: %v", err), 422)
			return
		}

		device, err := store.FindByToken(token)

		if err != nil {
			logrus.Errorf("device error: %v", err)
			http.Error(w, "Unknown device", 422)
			return
		}

		for _, m := range measures {
			device.AppendMeasure(&m)
		}

		writeResponse(w, []byte("OK"), 201)
	}
}
