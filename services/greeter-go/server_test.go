package main

import (
	"os"
	"testing"
)

func TestEnvVarEmptyOk(t *testing.T) {
	os.Setenv("MyTestVar", "1")
	defer func() {
		os.Unsetenv("MyTestVar")
	}()

	isEmpty := envVarEmpty("MyTestVar")
	if isEmpty {
		t.Errorf("Variable should have value of 1 , but instead was empty")
	}
}

func TestEnvVarEmptyNotOk(t *testing.T) {
	os.Setenv("MyTestVar", "")
	defer func() {
		os.Unsetenv("MyTestVar")
	}()

	isEmpty := envVarEmpty("MyTestVar")
	if !isEmpty {
		t.Errorf("Variable should have been empty , but instead had value: %s", os.Getenv("MyTestVar"))
	}
}

func TestCheckRequiredEnvNoError(t *testing.T) {
	os.Setenv(GREETER_SERVER_HOST, "1")
	os.Setenv(GREETER_SERVER_PORT, "1")
	os.Setenv(GREETER_SERVER_CERT, "1")
	os.Setenv(GREETER_SERVER_PRIVATE_KEY, "1")
	defer func() {
		os.Unsetenv(GREETER_SERVER_HOST)
		os.Unsetenv(GREETER_SERVER_PORT)
		os.Unsetenv(GREETER_SERVER_CERT)
		os.Unsetenv(GREETER_SERVER_PRIVATE_KEY)
	}()

	err := checkRequiredEnv()
	if err != nil {
		t.Errorf("checkRequireEnv should not return error: %s", err)
	}
}

func TestCheckRequiredEnvError(t *testing.T) {
	os.Setenv(GREETER_SERVER_HOST, "")
	os.Setenv(GREETER_SERVER_PORT, "1")
	os.Setenv(GREETER_SERVER_CERT, "1")
	os.Setenv(GREETER_SERVER_PRIVATE_KEY, "1")
	defer func() {
		os.Unsetenv(GREETER_SERVER_HOST)
		os.Unsetenv(GREETER_SERVER_PORT)
		os.Unsetenv(GREETER_SERVER_CERT)
		os.Unsetenv(GREETER_SERVER_PRIVATE_KEY)
	}()

	err := checkRequiredEnv()
	if err == nil {
		t.Errorf("checkRequireEnv should return error as GREETER_SERVER_HOST is empty")
	}
}

func TestGreetingLangToStr(t *testing.T) {
	testLang := GreeterRequest_English
	greetingStr := greetingLangToStr(testLang)
	if greetingStr != mapping.greetingLanguages[GreeterRequest_English] {
		t.Errorf("greeting string must be: %s", mapping.greetingLanguages[GreeterRequest_English])
	}
}
