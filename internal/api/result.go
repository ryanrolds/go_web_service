package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	StatusOk    = "ok"
	StatusError = "error"
)

type ApiResult struct {
	Status string      `json:"status"`
	Reason string      `json:"reason,omitempty"` // For returning error messages
	Total  int         `json:"total,omitempty"`  // For returning total count
	Result interface{} `json:"result,omitempty"` // For returning data
}

func WriteOk(ctx context.Context, w http.ResponseWriter, statusCode int, result interface{}, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload, err := json.Marshal(ApiResult{
		Status: StatusOk,
		Result: result,
		Total:  total,
	})
	if err != nil {
		log.Error(fmt.Errorf("error marshalling json: %w", err))
		return
	}

	_, err = w.Write(payload)
	if err != nil {
		log.Error(fmt.Errorf("error writing response: %w", err))
		return
	}
}

func WriteError(ctx context.Context, w http.ResponseWriter, err error) {
	log.Error(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	payload, err := json.Marshal(ApiResult{
		Status: StatusError,
		Reason: err.Error(),
	})
	if err != nil {
		log.Error(fmt.Errorf("error marshalling json: %w", err))
		return
	}

	_, err = w.Write(payload)
	if err != nil {
		log.Error(fmt.Errorf("error writing error response: %w", err))
		return
	}
}
