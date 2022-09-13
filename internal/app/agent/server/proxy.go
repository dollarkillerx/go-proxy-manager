package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/dollarkillerx/go-proxy-manager/internal/utils"
	"github.com/dollarkillerx/go-proxy-manager/proto/common"
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

	return ser.ListenAndServeTLS("", "")
}

func (s *Server) getCertificate(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
	log.Println(info.ServerName) // domain
	domain := s.cache.GetTaskByDomain(info.ServerName)
	if domain != nil {
		if domain.Certificate != nil && domain.EnableSsl {
			pair, err := tls.X509KeyPair([]byte(domain.Certificate.PublicKey), []byte(domain.Certificate.PrivateKey))
			if err != nil {
				return nil, err
			}

			return &pair, nil
		}
	}
	return nil, nil
}

type httpHandler struct {
	s *Server
}

func (h *httpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	host := request.Host

	domain := h.s.cache.GetTaskByDomain(host)
	if domain == nil {
		writer.WriteHeader(404)
		writer.Write([]byte("404"))
		return
	}

	switch domain.ProxyType {
	case common.ProxyType_ReverseProxyEnum:
		if domain.ReverseProxy != nil {
			target, err := url.Parse(domain.ReverseProxy.Target)
			if err != nil {
				log.Println(err)
				writer.WriteHeader(500)
				return
			}
			r := &httputil.ReverseProxy{
				Director: func(req *http.Request) {
					req.URL.Scheme = target.Scheme
					req.URL.Host = target.Host
					req.Host = target.Host //这才是关键
					req.URL.Path = utils.SingleJoiningSlash(target.Path, req.URL.Path)
					if target.RawQuery == "" || req.URL.RawQuery == "" {
						req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
					} else {
						req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
					}
				},
			}
			r.ServeHTTP(writer, request)
		}
	case common.ProxyType_RedirectProxyEnum:
		if domain.RedirectProxy != nil {
			switch domain.RedirectProxy.RedirectStateType {
			case common.RedirectStateType_MultipleChoices300:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 300)
			case common.RedirectStateType_MovedPermanently301:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 301)
			case common.RedirectStateType_Found302:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 302)
			case common.RedirectStateType_SeeOther303:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 303)
			case common.RedirectStateType_TemporaryRedirect307:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 307)
			case common.RedirectStateType_PermanentRedirect308:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 308)
			}
		}
		return
	case common.ProxyType_StreamProxyEnum:

	case common.ProxyType_R404ProxyEnum:
		writer.WriteHeader(404)
		writer.Write([]byte("404"))
		return
	}
}

type httpsHandler struct {
	s *Server
}

func (h *httpsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	host := request.Host

	domain := h.s.cache.GetTaskByDomain(host)
	if domain == nil {
		writer.WriteHeader(404)
		writer.Write([]byte("404"))
		return
	}

	switch domain.ProxyType {
	case common.ProxyType_ReverseProxyEnum:
		if domain.ReverseProxy != nil {
			target, err := url.Parse(domain.ReverseProxy.Target)
			if err != nil {
				log.Println(err)
				writer.WriteHeader(500)
				return
			}
			r := &httputil.ReverseProxy{
				Director: func(req *http.Request) {
					req.URL.Scheme = target.Scheme
					req.URL.Host = target.Host
					req.Host = target.Host //这才是关键
					req.URL.Path = utils.SingleJoiningSlash(target.Path, req.URL.Path)
					if target.RawQuery == "" || req.URL.RawQuery == "" {
						req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
					} else {
						req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
					}
				},
			}
			r.ServeHTTP(writer, request)
		}
	case common.ProxyType_RedirectProxyEnum:
		if domain.RedirectProxy != nil {
			switch domain.RedirectProxy.RedirectStateType {
			case common.RedirectStateType_MultipleChoices300:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 300)
			case common.RedirectStateType_MovedPermanently301:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 301)
			case common.RedirectStateType_Found302:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 302)
			case common.RedirectStateType_SeeOther303:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 303)
			case common.RedirectStateType_TemporaryRedirect307:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 307)
			case common.RedirectStateType_PermanentRedirect308:
				http.Redirect(writer, request, domain.RedirectProxy.Target, 308)
			}
		}
		return
	case common.ProxyType_StreamProxyEnum:

	case common.ProxyType_R404ProxyEnum:
		writer.WriteHeader(404)
		writer.Write([]byte("404"))
		return
	}
}
