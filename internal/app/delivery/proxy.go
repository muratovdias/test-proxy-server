package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/muratovdias/test-proxy-server/internal/app/entities"
)

func (h *Handler) reverseProxyRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var proxyRequest entities.ProxyRequest
		if err := json.NewDecoder(r.Body).Decode(&proxyRequest); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response, err := h.usecase.ProxyRequest(proxyRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
