package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func getOffsetFromRequest(r *http.Request) (int, error) {
	offset := 0

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return 0, fmt.Errorf("invalid page: %w", err)
		}

		if offset < 0 {
			return 0, fmt.Errorf("offset less than 0: %d", offset)
		}
	}

	return offset, nil
}

func getLimitFromRequest(r *http.Request) (int, error) {
	limit := 12

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return 0, fmt.Errorf("invalid limit: %w", err)
		}

		if limit < 0 {
			return 0, fmt.Errorf("limit less than 0: %d", limit)
		}

		if limit > 100 {
			return 0, fmt.Errorf("limit greater than 100: %d", limit)
		}
	}

	return limit, nil
}
