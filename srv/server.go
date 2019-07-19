package srv

import (
	"net/http"
	"time"

	"github.com/j4ng5y/environmentalist/srv/environmentalistrestful"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"
)

type Server struct {
	GRPCServer *grpc.Server
	HTTPServer *http.Server
	Router     *mux.Router
}

func NewServer() *Server {
	var S Server
	S.GRPCServer = grpc.NewServer()
	S.HTTPServer = &http.Server{
		Addr:         "0.0.0.0:5005",
		Handler:      S.NewRouter(),
		IdleTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(15) * time.Second,
		ReadTimeout:  time.Duration(15) * time.Second,
	}
	return &S
}

func (S *Server) SetHTTPAddress(addr string) *Server {
	S.HTTPServer = &http.Server{
		Addr:         addr,
		Handler:      S.NewRouter(),
		IdleTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(15) * time.Second,
		ReadTimeout:  time.Duration(15) * time.Second,
	}
	return S
}

func (*Server) NewRouter() *mux.Router {
	R := mux.NewRouter()

	R.HandleFunc("/new/{secretName}", environmentalistrestful.NewSecret)
	R.HandleFunc("/update/{secretName}", environmentalistrestful.UpdateSecret)
	R.HandleFunc("/delte/{secretName}", environmentalistrestful.DeleteSecret)
	R.HandleFunc("/view/{secretName}", environmentalistrestful.ViewSecret)

	return R
}
