package delete

import (
	"errors"
	"go-urlshortener/internal/lib/api/response"
	"go-urlshortener/internal/lib/logger/sl"
	"go-urlshortener/internal/storage"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(logger *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		logger := logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			logger.Warn("alias is empty")
			render.JSON(w, r, response.Error("not found"))
			return
		}

		err := urlDeleter.DeleteURL(alias)

		if errors.Is(err, storage.ErrURLNotFound) {
			logger.Info("url not found", "alias", alias)
			render.JSON(w, r, response.Error("not found"))
			return
		}

		if err != nil {
			logger.Error("failed to delete url", sl.Err(err))
			render.JSON(w, r, response.Error("internal error"))
			return
		}

		logger.Info("url deleted", "alias", alias)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
