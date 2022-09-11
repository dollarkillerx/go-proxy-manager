package server

import (
	"crypto/tls"
	"net/http"
)

func (s *Server) listen() error {
	ser := http.Server{
		Addr:    ":80",
		Handler: &httpHandler{s: s},
	}
	return ser.ListenAndServe()
}

func (s *Server) listenAndTLS() error {
	tlsConfig := &tls.Config{
		GetCertificate: s.getCertificate,
	}

	ser := http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
		Handler:   &httpsHandler{s: s},
	}

	return ser.ListenAndServe()
}

func (s *Server) getCertificate(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return nil, nil
}

type httpHandler struct {
	s *Server
}

func (h *httpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

type httpsHandler struct {
	s *Server
}

func (h *httpsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
