package docs

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed openapi.yaml swagger-ui.html
var files embed.FS

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	sub, err := fs.Sub(files, ".")
	if err != nil {
		panic(err)
	}
	mux.Handle("GET /docs/", http.StripPrefix("/docs/", http.FileServer(http.FS(sub))))
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/swagger-ui.html", http.StatusMovedPermanently)
	})
}