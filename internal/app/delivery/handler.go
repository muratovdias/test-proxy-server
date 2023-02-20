package delivery

import (
	"net/http"

	"github.com/muratovdias/test-proxy-server/internal/app/usecase"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(u *usecase.Usecase) *Handler {
	return &Handler{
		usecase: u,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", h.reverseProxyRequest)
	return router
}
