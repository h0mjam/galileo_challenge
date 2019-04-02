package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func jsonRequest(r *http.Request, t interface{}) error {
	err := json.NewDecoder(r.Body).Decode(t)

	if err != nil && err != io.EOF {
		return err
	}

	err = r.Body.Close()

	if err != nil {
		return err
	}

	return nil
}

func writeResponse(w http.ResponseWriter, data []byte, code int) {
	if code != 0 {
		w.WriteHeader(code)
	}

	_, err := w.Write(data)

	if err != nil {
		logrus.Errorf("error write response: %v", err)
	}
}
