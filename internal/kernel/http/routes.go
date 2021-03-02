package http

import (
	"cat-unifier/internal/kernel/common/library/decorators"
	"cat-unifier/internal/kernel/common/library/repositiries"
	"fmt"
	"github.com/gorilla/mux"
)

type httpRoutes struct {
	config    repositiries.IConfigRepository
	decorator decorators.IHttpDecorator
	router    *mux.Router
}

type IHttpRoutes interface {
	Init() *mux.Router
}

func (hr *httpRoutes) Init() *mux.Router {
	prefix := fmt.Sprintf("%s", hr.config.Get("app.version", "v1"))

	hr.router.
		Methods("GET").
		PathPrefix(prefix).
		Path("balances/{account}").
		HandlerFunc(hr.decorator.HandleGetBalance).
		Name("GetBalance")

	hr.router.
		Methods("GET").
		PathPrefix(prefix).
		Path("blocks/{block}").
		HandlerFunc(hr.decorator.HandleGetBlock).
		Name("GetBlock")

	hr.router.
		Methods("GET").
		PathPrefix(prefix).
		Path("transactions/{transaction}").
		HandlerFunc(hr.decorator.HandleGetTransaction).
		Name("GetTransaction")

	return hr.router
}

func NewHttpRoutes(r *mux.Router, cfg repositiries.IConfigRepository, http decorators.IHttpDecorator) IHttpRoutes {
	return &httpRoutes{router: r, config: cfg, decorator: http}
}
