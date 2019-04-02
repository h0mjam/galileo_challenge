package server

import "net/http"

type Server struct {
	Addr    string
	Handler *http.ServeMux
}

func NewServer(addres string) *Server {
	return &Server{
		Addr:    addres,
		Handler: http.NewServeMux(),
	}
}

func (s *Server) Listen() error {
	return http.ListenAndServe(s.Addr, s.Handler)
}

func (s *Server) AddHandler(
	route string,
	method string,
	handler http.HandlerFunc,
) {
	s.Handler.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if method != "Any" && method != r.Method {
			http.Error(w, "Method not allowed", 405)
			return
		}

		handler(w, r)
	})
}

func (s *Server) Post(route string, handler http.HandlerFunc) {
	s.AddHandler(route, http.MethodPost, handler)
}

func (s *Server) Get(route string, handler http.HandlerFunc) {
	s.AddHandler(route, http.MethodGet, handler)
}
