package main

import (
	"OSlash/protos/onboarding/register"
	"OSlash/protos/user/tweet"
	"context"
	"crypto/x509"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type myService struct{}

var (
	authpoint    = flag.String("auth_end_points", "localhost:50051", "expose gamezop end point ")
	demoCertPool *x509.CertPool
)

func newServer() *myService {
	return new(myService)
}

func ExposePoint(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := regsiterpb.RegisterRegisterServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	err = tweetpb.RegisterTweetServiceHandlerFromEndpoint(ctx, mux, *authpoint, dialOpts)
	if err != nil {
		return err
	}
	grpcMux := http.NewServeMux()
	grpcMux.Handle("/", mux)
	//grpcMux.HandleFunc("/swagger/", serveSwagger)
	log.Println("Starting Endpoint Exposed Server: localhost:5051")
	http.ListenAndServe(address, grpcMux)
	return nil
}

func main() {
	Init()
}

// initilization
func Init() {
	if err := ExposePoint(":5051"); err != nil {
		log.Fatal("Error in serve", err)
	}
}

// // serve swagger
// func serveSwagger(w http.ResponseWriter, r *http.Request) {
// 	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
// 	p = path.Join("swagger/swagger-ui/", p)
// 	http.ServeFile(w, r, p)
// }
