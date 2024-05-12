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
	root = "/Users/julieqiu/code/julieqiu/snippetgen"
)

var (
	googleapisDir = fmt.Sprintf("%s/testdata/googleapis", root)
	outputDir     = fmt.Sprintf("%s/generated", root)
	files         = []string{
		fmt.Sprintf("%s/google/cloud/secretmanager/v1/resources.proto", googleapisDir),
		fmt.Sprintf("%s/google/cloud/secretmanager/v1/service.proto", googleapisDir),
	}
	options = []string{
		"go-gapic-package=cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb",
		fmt.Sprintf("api-service-config=%s/google/datastore/v1/datastore_v1.yaml", googleapisDir),
		fmt.Sprintf("api-service-config=%s/google/cloud/secretmanager/v1/secretmanager_v1.yaml", googleapisDir),
		fmt.Sprintf("grpc-service-config=%s/google/cloud/secretmanager/v1/secretmanager_grpc_service_config.json", googleapisDir),
	}
)

func run() error {
	cmd := exec.Command(
		"protoc",
		"--go_gapic_out", outputDir,
		"--go_gapic_opt", fmt.Sprintf("'%s'", strings.Join(options, ",")),
		"-I", googleapisDir,
		fmt.Sprintf("%s/google/cloud/secretmanager/v1/*.proto", googleapisDir),
	)

	fmt.Println(cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
