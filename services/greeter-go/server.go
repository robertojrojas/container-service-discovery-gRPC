package main

import (
	"fmt"
	"log"
	"net"
	"os"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type greeterService struct{}
type config struct {
	serverHost  string
	serverPort  string
	greeterCert string
	greeterKey  string
}

func startServer() error {

	config, err := getConfig()
	if err != nil {
		return err
	}

	serverHostPort := fmt.Sprintf("%s:%s", config.serverHost, config.serverPort)
	lis, err := net.Listen("tcp", serverHostPort)
	if err != nil {
		return err
	}
	log.Printf("Listening on [%s]....\n", serverHostPort)

	creds, err := credentials.NewServerTLSFromFile(config.greeterCert, config.greeterKey)
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)
	cs := &greeterService{}

	RegisterGreeterServiceServer(s, cs)

	return s.Serve(lis)

}

func (gs *greeterService) Greet(ctx context.Context, gr *GreeterRequest) (*GreeterResponse, error) {
	var response string

	switch gr.GetLang() {
	case GreeterRequest_English:
		response = "Hello"
	case GreeterRequest_Spanish:
		response = "Hola"
	case GreeterRequest_French:
		response = "Bonjour"
	}

	return &GreeterResponse{
		Greet: response,
	}, nil
}

func getConfig() (*config, error) {

	//TODO: check environment variables and return error if one is not foud

	envConfig := config{}

	envConfig.serverHost = os.Getenv("GREETER_SERVER_HOST")

	envConfig.serverPort = os.Getenv("GREETER_SERVER_PORT")

	envConfig.greeterCert = os.Getenv("GREETER_SERVER_CERT")

	envConfig.greeterKey = os.Getenv("GREETER_SERVER_PRIVATE_KEY")

	return &envConfig, nil
}
