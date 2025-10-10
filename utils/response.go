package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeResponseIntoStruct(r *http.Request, decodingStruct *any) error {
	if err := json.NewDecoder(r.Body).Decode(&decodingStruct); err != nil {
		return err
	}
	return nil
}
