#!/bin/bash

###############################################################################
# Test Settings

# What are your registry details?
export OCI_ROOT_URL="https://r.myreg.io"
export OCI_NAMESPACE="myorg/myrepo"
export OCI_CROSSMOUNT_NAMESPACE="myorg/other"
export OCI_USERNAME="myuser"
export OCI_PASSWORD="mypass"

# Which workflows to test?
export OCI_TEST_PULL=1
export OCI_TEST_PUSH=1
export OCI_TEST_CONTENT_DISCOVERY=1
export OCI_TEST_CONTENT_MANAGEMENT=1

# Miscellaneous settings
export OCI_HIDE_SKIPPED_WORKFLOWS=0
export OCI_DEBUG=0
export OCI_DELETE_MANIFEST_BEFORE_BLOBS=0

###############################################################################
# Test Execution

set -euo pipefail

go test -c
./conformance.test
