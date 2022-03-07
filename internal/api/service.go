package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ryanrolds/go_web_service/internal/domain"
	"github.com/ryanrolds/go_web_service/internal/services/service"
)

func (a *API) listServices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	offset, err := getOffsetFromRequest(r)
	if err != nil {
		WriteError(ctx, w, fmt.Errorf("error getting page from request: %w", err))
		return
	}

	limit, err := getLimitFromRequest(r)
	if err != nil {
		WriteError(ctx, w, fmt.Errorf("error getting limit from request: %w", err))
		return
	}

	var results []domain.Service

	// Check if search query is present
	search := r.URL.Query().Get("search")
	if search != "" {
		// Fetch page of services from DB
		results, err = service.PageSearched(r.Context(), a.db, search, offset, limit)
		if err != nil {
			WriteError(ctx, w, fmt.Errorf("error fetching page of searched services: %w", err))
			return
		}
	} else {
		// Fetch page of services from DB
		results, err = service.Page(r.Context(), a.db, offset, limit)
		if err != nil {
			WriteError(ctx, w, fmt.Errorf("error fetching page of services: %w", err))
			return
		}
	}

	// Get total number of services
	total, err := service.Count(ctx, a.db)
	if err != nil {
		WriteError(ctx, w, fmt.Errorf("error counting services: %w", err))
		return
	}

	WriteOk(ctx, w, http.StatusOK, results, total)
}

func (a *API) getService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the service id from the request
	id, ok := mux.Vars(r)["id"]
	if !ok || id == "" {
		WriteError(ctx, w, fmt.Errorf("missing service id"))
		return
	}

	service, err := service.Get(r.Context(), a.db, id)
	if err != nil {
		WriteError(ctx, w, fmt.Errorf("error listing services: %w", err))
		return
	}

	WriteOk(ctx, w, http.StatusOK, service, 1)
}
