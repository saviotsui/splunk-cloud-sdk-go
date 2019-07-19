#!/bin/bash

echo "==============================================="
echo "Beginning integration tests"
echo "==============================================="
env | grep SPLUNK_CLOUD_HOST
env | grep TENANT_ID

echo "==============================================="

# Get the BEARER_TOKEN setup
CONFIG_FILE="./.token"
if [ -f $CONFIG_FILE ]; then
    echo "Token found in $CONFIG_FILE"
    export BEARER_TOKEN=$(cat $CONFIG_FILE)
else
    echo "Token was not set to $CONFIG_FILE"
    exit 1
fi

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./cicd/integration/
FULL_INTEGRATION_TEST_CODECOV_FILE_NAME=integration_test_codecov.out
FULL_INTEGRATION_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_INTEGRATION_TEST_CODECOV_FILE_NAME

# Required to run just the service tests
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
               -covermode=count \
               -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
               ./test/integration/...
else
    echo "Running integration tests and gating on failures..."
    go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
               -covermode=count \
               -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
               ./test/integration/... \
        || exit 1
fi

# Upload code cov report
echo "TODO: CODE COVERAGE IS NOT CURRENTLY SUPPORTED! CODECOV REPORT WILL NOT BE UPLOADED."