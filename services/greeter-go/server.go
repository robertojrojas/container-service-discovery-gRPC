package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	GREETER_SERVER_HOST        = "GREETER_SERVER_HOST"
	GREETER_SERVER_PORT        = "GREETER_SERVER_PORT"
	GREETER_SERVER_CERT        = "GREETER_SERVER_CERT"
	GREETER_SERVER_PRIVATE_KEY = "GREETER_SERVER_PRIVATE_KEY"
)

type greeterService struct{}
type config struct {
	serverHost  string
	serverPort  string
	greeterCert string
	greeterKey  string
}
type greetingLangs struct {
	greetingLanguages map[GreeterRequest_Language]string
}

var mapping = greetingLangs{
	greetingLanguages: map[GreeterRequest_Language]string{
		GreeterRequest_English: "Hello",
		GreeterRequest_Spanish: "Hola",
		GreeterRequest_French:  "Bonjour",
	},
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
	return &GreeterResponse{
		Greet: greetingLangToStr(gr.GetLang()),
	}, nil
}

func greetingLangToStr(greetingLang GreeterRequest_Language) string {
	response, found := mapping.greetingLanguages[greetingLang]
	if !found {
		response = mapping.greetingLanguages[GreeterRequest_English]
	}

	return response
}

func getConfig() (*config, error) {
	err := checkRequiredEnv()
	if err != nil {
		return nil, err
	}

	envConfig := config{}
	envConfig.serverHost = os.Getenv(GREETER_SERVER_HOST)
	envConfig.serverPort = os.Getenv(GREETER_SERVER_PORT)
	envConfig.greeterCert = os.Getenv(GREETER_SERVER_CERT)
	envConfig.greeterKey = os.Getenv(GREETER_SERVER_PRIVATE_KEY)

	return &envConfig, nil
}

func checkRequiredEnv() error {
	reqEnvs := []string{
		GREETER_SERVER_HOST,
		GREETER_SERVER_PORT,
		GREETER_SERVER_CERT,
		GREETER_SERVER_PRIVATE_KEY,
	}
	envsNotFound := make([]interface{}, 0)
	for _, reqEnv := range reqEnvs {
		if envVarEmpty(reqEnv) {
			envsNotFound = append(envsNotFound, reqEnv)
		}
	}
	if len(envsNotFound) == 0 {
		return nil
	}
	errorStr := fmt.Sprintln(envsNotFound...)
	errorStr = fmt.Sprintf("USAGE: GreeterServer - The following environment variables are required:\n %s \n", errorStr)
	return errors.New(errorStr)
}

func envVarEmpty(varKey string) bool {
	if envVar, found := os.LookupEnv(varKey); !found || (strings.TrimSpace(envVar) == "") {
		return true
	}

	return false
}
