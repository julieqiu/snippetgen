package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const (
	inputDir = "/Users/julieqiu/code/julieqiu/snippetgen/testdata/googleapis"
	root     = "/Users/julieqiu/code/julieqiu/snippetgen"
)

var (
	files = []string{
		inputDir + "/google/api/annotations.proto",
		inputDir + "/google/api/client.proto",
		inputDir + "/google/api/field_behavior.proto",
		inputDir + "/google/api/http.proto",
		inputDir + "/google/api/launch_stage.proto",
		inputDir + "/google/api/resource.proto",
		inputDir + "/google/cloud/secretmanager/v1/resources.proto",
		inputDir + "/google/cloud/secretmanager/v1/service.proto",
		inputDir + "/google/iam/v1/iam_policy.proto",
		inputDir + "/google/iam/v1/logging/audit_data.proto",
		inputDir + "/google/iam/v1/options.proto",
		inputDir + "/google/iam/v1/policy.proto",
		inputDir + "/google/type/expr.proto",
	}
	opts = map[string]string{
		"go-gapic-package":    "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb",
		"grpc-service-config": inputDir + "/google/cloud/secretmanager/v1/secretmanager_grpc_service_config.json",
		"api-service-config":  inputDir + "/google/cloud/secretmanager/v1/secretmanager_v1.yaml",
	}
)

func run() error {
	var parameters []string
	for k, v := range opts {
		parameters = append(parameters, fmt.Sprintf("%s=%s", k, v))
	}

	cmd := exec.Command(
		"protoc",
		"-I",
		inputDir,
		"--go_gapic_out",
		root+"/generated",
		"--go_gapic_opt",
		"'go-gapic-package=go-gapic-package=cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb,grpc-service-config=/Users/julieqiu/code/julieqiu/snippetgen/testdata/googleapis/google/cloud/secretmanager/v1/secretmanager_grpc_service_config.json,api-service-config=/Users/julieqiu/code/julieqiu/snippetgen/testdata/googleapis/google/cloud/secretmanager/v1/secretmanager_v1.yaml'",
		strings.Join(files, " "),
	)
	fmt.Println(cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
