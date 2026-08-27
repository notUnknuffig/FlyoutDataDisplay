package state

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func ReadString(r string, data any) error {
	var byteData []byte = []byte(r)
	return json.Unmarshal(byteData, data)
}

func WriteString(data any) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), err
}
