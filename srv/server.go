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

func (*Server) NewRouter() *mux.Router {
	R := mux.NewRouter()

	R.HandleFunc("/hashicorp-vault/new/{secretName}", environmentalistrestful.NewHashicorpVaultSecret)
	R.HandleFunc("/hashicorp-vault/update/{secretName}", environmentalistrestful.UpdateHashicorpVaultSecret)
	R.HandleFunc("/hashicorp-vault/delte/{secretName}", environmentalistrestful.DeleteHashicorpVaultSecret)
	R.HandleFunc("/hashicorp-vault/view/{secretName}", environmentalistrestful.ViewHashicorpVaultSecret)
	R.HandleFunc("/aws-ssm/new/{secretName}", environmentalistrestful.NewAWSSSMSecret)
	R.HandleFunc("/aws-ssm/update/{secretName}", environmentalistrestful.UpdateAWSSSMSecret)
	R.HandleFunc("/aws-ssm/delte/{secretName}", environmentalistrestful.DeleteAWSSSMSecret)
	R.HandleFunc("/aws-ssm/view/{secretName}", environmentalistrestful.ViewAWSSSMSecret)

	return R
}
