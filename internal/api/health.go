package api

import (
	"net/http"
)

func (a *API) health(w http.ResponseWriter, r *http.Request) {
	WriteOk(r.Context(), w, http.StatusOK, ApiResult{Status: "ok"}, 0)
}
