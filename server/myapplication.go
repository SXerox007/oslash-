package main

import (
	"OSlash/api/onboarding/register"
	"OSlash/api/user"
	"OSlash/base/server"
	db "OSlash/db/postgres"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	PgSetup()
	ServerSetup()
}

func PgSetup() {
	db.DBConnecting()
}

func main() {
	Init()
}

//brain setup
func ServerSetup() {
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("Error in server start:", err)
		return
	}
	//Create the New gRPC Server
	srv := server.CreateNewgRPCServer(false, nil)
	//Register reflection on gRPC server
	reflection.Register(srv)
	// all the rpc services
	rpcServices(srv)

	go func() {
		fmt.Println("Server start on Port:50051")
		if err := srv.Serve(listner); err != nil {
			log.Println("Error in Serve:", err)
			return
		}
	}()
	//make a channnel that will wait for server to close or interrupt
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// this will block the code while server
	<-ch
	log.Println("Exit the Server with:", os.Interrupt)
}

// All the services
func rpcServices(srv *grpc.Server) {
	//Register the User (on the base of role ADMIN/USER/SUPER-ADMIN)
	register.RegisterUserService(srv)
	// users
	user.RegisterTweetService(srv)
}
