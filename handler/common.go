package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func jsonRequest(r *http.Request, t interface{}) error {
	body := make([]byte, r.ContentLength)
	n, err := r.Body.Read(body)

	if n == 0 {
		return fmt.Errorf("request body is empty")
	}

	if err != nil && err != io.EOF {
		return err
	}

	err = json.Unmarshal(body, t)

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
