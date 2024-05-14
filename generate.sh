#!/usr/bin/env bash
#
# GOOGLEAPIS DIRECTORY
export GOOGLEAPIS="/Users/julieqiu/code/googleapis/googleapis"

# OUTPUT DIRECTORY
export OUTPUT=generated

# OPTIONS
export GO_GAPIC_PACKAGE="go-gapic-package=cloud.google.com/go/${1}/api${2}/${1}pb;${1}pb"
export API_SERVICE_CONFIG="api-service-config=${GOOGLEAPIS}/google/cloud/${1}/${2}/${1}_${2}.yaml"
export OPTIONS="${GO_GAPIC_PACKAGE},${API_SERVICE_CONFIG}"

# API
export API="${GOOGLEAPIS}/google/cloud/${1}/${2}/*.proto"

export EXIT_CODE=0

info() { echo -e "$@}" 1>&2; }
err() { echo -e "$@" 1>&2; EXIT_CODE=1; }

# runcmd prints an info log describing the command that is about to be run, and
# then runs it. It sets EXIT_CODE to non-zero if the command fails, but does not exit
# the script.
runcmd() {
  msg="$@"
  # Truncate command logging for narrow terminals.
  # Account for the 2 characters of '$ '.
  if [[ $MAXWIDTH -gt 0 && ${#msg} -gt $MAXWIDTH ]]; then
    msg="${msg::$(( MAXWIDTH - 3 ))}..."
  fi

  info "\$ $msg \n"
  $@ || err "command failed"
}

main() {
    if [[ $1 == "" ]]; then
        err "FAILED; the first argument should be the name of the API, for example, secretmanager"
        exit $EXIT_CODE
    fi
    if [[ $2 == "" ]]; then
        err "FAILED; `$2` should be the version of the API, for example, v1"
        exit $EXIT_CODE
    fi

    go install ./gapic-generator-go/cmd/protoc-gen-go_gapic

    # protoc -I googleapis --go_gapic_out generated --go_gapic_opt 'go-gapic-package=cloud.google.com/go/secretmanager/apiv1/secretmanagerpb;secretmanagerpb,grpc-service-config=googleapis/google/cloud/secretmanager/v1/secretmanager_grpc_service_config.json,api-service-config=googleapis/google/cloud/secretmanager/v1/secretmanager_v1.yaml' googleapis/google/cloud/secretmanager/v1/*.proto
    runcmd protoc -I $GOOGLEAPIS \
        --go_gapic_out "$OUTPUT" \
        --go_gapic_opt "$OPTIONS" \
        "$API"
}
main $@
