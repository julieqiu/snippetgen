# snippetgen

```
protoc -I googleapis --go_gapic_out generated --go_gapic_opt 'go-gapic-package=cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb,grpc-service-config=googleapis/google/cloud/secretmanager/v1/secretmanager_grpc_service_config.json,api-service-config=googleapis/google/cloud/secretmanager/v1/secretmanager_v1.yaml' googleapis/google/cloud/secretmanager/v1/*.proto
```
