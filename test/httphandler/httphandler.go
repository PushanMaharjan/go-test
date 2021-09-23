// package httphandler

// import (
// 	"net/http"

// 	"go.uber.org/zap"
// )

// type Handler struct {
// 	mux    *http.ServeMux
// 	logger *zap.SugaredLogger
// }

// func HandlerFunc(s *http.ServeMux, logger *zap.SugaredLogger) *Handler {
// 	h := Handler{s, logger}
// 	h.registerRoutes()

// 	return &h
// }

// func (h *Handler) registerRoutes() {
// 	h.mux.HandleFunc("/", h.hello)
// }

// func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(200)
// 	w.Write([]byte("Hello World"))
// }
