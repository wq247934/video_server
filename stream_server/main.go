package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests!")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/:vid-id", StreamHandler)
	router.POST("/video/:vid-id", UploadHandler)
	return router
}
func main() {
	r := RegisterHandlers()
	middleware := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", middleware)
}
