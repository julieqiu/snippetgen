package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/julieqiu/snippetgen/gapic-generator-go/internal/gengapic"
	"github.com/julieqiu/snippetgen/gapic-generator-go/internal/testing/sample"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const options = "go-gapic-package=cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb,api-service-config=secretmanager_v1.yaml,grpc-service-config=secretmanager_grpc_service_config.json"

func TestGenSecretManager(t *testing.T) {
	input := &pluginpb.CodeGeneratorRequest{
		FileToGenerate: []string{"google/cloud/secretmanager/v1/service.proto"},
		Parameter:      proto.String(options),
		ProtoFile:      []*descriptorpb.FileDescriptorProto{sample.ProtoFile()},
	}
	output, err := gengapic.Gen(input)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range output.File {
		if file.Name != nil {
			if strings.Contains(*file.Name, "Create") {
				fmt.Println(*file.Name)
				// fmt.Println(*file.Content)
			}
			// fmt.Println(strings.Repeat("-", 80))
		}
	}
}
