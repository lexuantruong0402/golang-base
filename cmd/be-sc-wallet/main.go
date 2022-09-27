package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	db "smc-wallet-be/internal/database"
	grpc_handler "smc-wallet-be/internal/module/user/grpc"
	"smc-wallet-be/proto/pb"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGRPCServer(
	userServer pb.UserServiceServer,
	listener net.Listener,
	enableTLS bool,
) error {
	//accessibleRoles := map[string][]string{}
	//interceptor := services.NewAuthInterceptor(jwtManager, accessibleRoles)
	serverOptions := []grpc.ServerOption{
		//grpc.UnaryInterceptor(interceptor.Unary()),
		//grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := grpc.NewServer(serverOptions...)

	//------START register service Server------
	pb.RegisterUserServiceServer(grpcServer, userServer)
	//------END register service Server------
	reflection.Register(grpcServer)

	log.Printf("Start GRPC server at %s, TLS = %t", listener.Addr().String(), enableTLS)
	return grpcServer.Serve(listener)
}

func main() {
	fmt.Println("Starting API Blockchain Internal!")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	//Create connection db
	db := db.ConnectToDB()
	host := os.Getenv("HOST_GRPC")
	log.Println(os.Getenv("RUN_GRPC_ENDPOINT"))
	if runGRPC, _ := strconv.ParseBool(os.Getenv("RUN_GRPC_ENDPOINT")); runGRPC {
		//run grpc serve
		port := os.Getenv("LISTEN_GRPC_PORT")
		//var address string
		address := fmt.Sprintf("%s:%s", host, port)
		listener, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("[Internal Block Chain API] Cannot start server at address: %s", address)
		}

		userServer := grpc_handler.NewUserGRPCServer(db)

		err = runGRPCServer(
			userServer,
			listener,
			false,
		)
		if err != nil {
			log.Fatalf("[ProductService][Main] Cannot start event with err: %s", err)
		}
	}
}
