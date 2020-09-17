package main

import (
	"coding-grpc-gateway/endpoint"
	_ "coding-grpc-gateway/endpoint"
	"coding-grpc-gateway/swagger"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = endpoint.RegisterEndpoint(ctx, grpcMux, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	mux.HandleFunc("/swagger/", swagger.ServeSwaggerFile)
	swagger.ServeSwaggerUI(mux)

	log.Print("coding-grpc-gateway gRPC Server gateway start at port 8080...")
	http.ListenAndServe(":8080", mux)
	return nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
