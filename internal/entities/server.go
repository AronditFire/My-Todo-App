package entities

import (
	"context"
	"net/http"
)

type Service struct {
	httpsServer *http.Server
}

func (sv *Service) RunServer(port string, h http.Handler) error {
	sv.httpsServer = &http.Server{
		Addr:    ":" + port,
		Handler: h,
	}

	return sv.httpsServer.ListenAndServe()
}

func (sv *Service) Shutdown() error {
	return sv.httpsServer.Shutdown(context.Background())
}
