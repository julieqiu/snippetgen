# snippetgen

```
protoc -I testdata/googleapis --go_gapic_out generated --go_gapic_opt 'go-gapic-package=cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb,grpc-service-config=testdata/googleapis/google/cloud/secretmanager/v1/secretmanager_grpc_service_config.json'  testdata/googleapis/google/cloud/secretmanager/v1/resources.proto testdata/googleapis/google/cloud/secretmanager/v1/service.proto
```
