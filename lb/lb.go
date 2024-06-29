package lb

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Backend struct {
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
}

func NewBackedServer(url *url.URL) Backend {
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("%s :: %s", url.Host, err.Error())
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}
	return Backend{
		URL:          url,
		ReverseProxy: proxy,
	}
}

type ServerPool struct {
	servers     []*Backend
	connections uint64
}

func (sp *ServerPool) AddServer(be *Backend) {
	sp.servers = append(sp.servers, be)
	log.Println("Added new server to pool")
}

func (sp *ServerPool) getNextServer() *Backend {
	next := int(sp.connections) % len(sp.servers)
	return sp.servers[next]
}

type LoadBalancer struct {
	Port    int
	Servers ServerPool
}

func (lb *LoadBalancer) Start() error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", lb.Port),
		Handler: http.HandlerFunc(lb.processRequests),
	}

	log.Println("Starting Load Balancer on port", lb.Port)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (lb *LoadBalancer) processRequests(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request")
	server := lb.Servers.getNextServer()
	if server != nil {
		server.ReverseProxy.ServeHTTP(w, r)
		atomic.StoreUint64(&lb.Servers.connections, atomic.AddUint64(&lb.Servers.connections, 1))
		return
	}
	log.Println("Failed to get a peer")
	http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
}
