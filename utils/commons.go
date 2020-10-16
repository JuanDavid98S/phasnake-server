package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// RespondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError return error message using RespondwithJSON method
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}

// GenerateUUID retrieve a UUID string (not standart RFC-4122)
func GenerateUUID() (uuid string) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
	return uuid
}

// ParseInt64 retrieve a int64 value
func ParseInt64(s string) int64 {
	dec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return dec
}
