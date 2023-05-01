package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type RouterImpl struct {
	logger *zerolog.Logger
}

func NewRouter(logger *zerolog.Logger) *chi.Mux {
	impl := &RouterImpl{logger}

	router := chi.NewRouter()
	router.Get("/api/_hc", impl.healthCheck)
	return router
}

func (impl *RouterImpl) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
