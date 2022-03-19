# Conformance Tests for the OCI Distribution Spec

***Table of Contents***

1. [What is the OCI Distribution Spec?](#oci-distribution-spec)
1. [What are Conformance Tests?](#what-are-conformance-tests)
1. [How to run the Conformance Tests](#how-to-run-the-conformance-tests)
1. [Workflows](#workflows)
1. [Test Reports in HTML](#test-reports-in-html)
1. [Teardown Order](#teardown-order)

## What is the OCI Distribution Spec?

The [OCI distribution specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md) is an API standard for ***registries***, which are essentially *HTTP servers that implement the following*:

* pulling container images
* pushing container images
* discovering container images
* managing container images

[Link to the official OCI distribution specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md)

## What are Conformance Tests?

In order for a registry to conform to the [OCI distribution specification](https://github.com/opencontainers/distribution-spec/blob/main/spec.md), it must correctly implement the following ***workflows***:

**Workflow #1:** Pulling container images.

**Workflow #2:** Pushing container images.

**Workflow #3:** Discovering container images.

**Workflow #4:** Managing container images.

The details of each workflow are described below in the section [Workflows](#workflows)

## How to Run the Conformance Tests

### 1. Build the testing binary.

Requires Go 1.17+.

In this directory, build the testing binary:

```
$ go test -c
```

This will produce an executable at `conformance.test`.

### 2. Set the environment variables in the testing script.

The testing script `conformance.sh` will run the conformance tests on a registry. Before you run the testing script, you should set the environment variables in the testing script for your registry.

```
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
```

### 3.Run the testing script.

Run the following in your shell:

```bash
$ ./conformance.sh
```

This will produce a `junit.xml` and `report.html` with the test results.

Note: for some registries, you may need to create `OCI_NAMESPACE` ahead of time.

## Workflows

The tests are broken down into 4 major categories:

1. **Pull** - Highest priority - All OCI registries MUST support pulling OCI container
images.
2. **Push** - Registries need a way to get content to be pulled, but clients can/should
be more forgiving here. For example, if needing to fallback after an unsupported endpoint.
3. **Content Discovery** - Includes tag listing (and possibly search in the future).
4. **Content Management** - Lowest Priority - Includes tag, blob, and repo deletion.
(Note: Many registries may have other ways to accomplish this than the OCI API.)

In addition, each category has its own setup and teardown processes where appropriate.

### Pull

The Pull tests validate that content can be retrieved from a registry.

These tests are run when the following is set in the environment:
```
OCI_TEST_PULL=1
```

Regardless of whether the Push tests are enabled, as part of setup for the Pull tests,
content will be uploaded to the registry.
If you wish to prevent this, you can set the following environment variables pointing
to content already present in the registry:

```
# Optional: set to prevent automatic setup
OCI_MANIFEST_DIGEST=<digest>
OCI_TAG_NAME=<tag>
OCI_BLOB_DIGEST=<digest>
```

### Push

The Push tests validate that content can be uploaded to a registry.

To enable the Push tests, you must explicitly set the following in the environment:

```
# Required to enable
OCI_TEST_PUSH=1
```

Some registries may require a workaround for Authorization during the push flow. To set your own scope, set the following in the environment:

```
# Set the auth scope
OCI_AUTH_SCOPE="repository:mystuff/myrepo:pull,push"
```

Most registries currently require at least one layer to be uploaded (and referenced in the appropriate section of the manifest)
before a manifest upload will succeed. By default, the push tests will attempt to push two manifests: one with a single layer,
and another with no layers. If the empty-layer test is causing a failure, it can be skipped by setting the following in the
environment:

```
# Enable layer upload
OCI_SKIP_EMPTY_LAYER_PUSH_TEST=1
```

The test suite will need access to a second namespace. This namespace is used to check support for cross-repository mounting
of blobs, and may need to be configured on the server-side in advance. It is specified by setting the following in
the environment:

```
# The destination repository for cross-repository mounting:
OCI_CROSSMOUNT_NAMESPACE="myorg/other"
```

If you want to test the behaviour of automatic content discovery, you should set the `OCI_AUTOMATIC_CROSSMOUNT` variable.

```
# Do not test automatic cross mounting
unset OCI_AUTOMATIC_CROSSMOUNT

# Test that automatic cross mounting is working as expected
OCI_AUTOMATIC_CROSSMOUNT=1

# Test that automatic cross mounting is disabled
OCI_AUTOMATIC_CROSSMOUNT=0
```

### Content Discovery

The Content Discovery tests validate that the contents of a registry can be discovered.

To enable the Content Discovery tests, you must explicitly set the following in the environment:

```
# Required to enable
OCI_TEST_CONTENT_DISCOVERY=1
```

As part of setup of these tests, a manifest and associated tags will be pushed to the registry.
If you wish to prevent this, you can set the following environment variable pointing
to list of tags to be returned from `GET /v2/<name>/tags/list`:

```
# Optional: set to prevent automatic setup
OCI_TAG_LIST=<tag1>,<tag2>,<tag3>,<tag4>
```

### Content Management

The Content Management tests validate that the contents of a registry can be deleted or otherwise modified.

To enable the Content Management tests, you must explicitly set the following in the environment:

```
# Required to enable
OCI_TEST_CONTENT_MANAGEMENT=1
```

Note: The Content Management tests explicitly depend upon the Push and Content Discovery tests, as there is no
way to test content management without also supporting push and content discovery.

## Test Reports in HTML
By default, the HTML report will show tests from all workflows. To hide workflows that have been disabled from
the report, you must set the following in the environment:

```
# Required to hide disabled workflows
OCI_HIDE_SKIPPED_WORKFLOWS=1
```

## Teardown Order

By default, the teardown phase of each test deletes blobs before manifests. Some registries require the opposite order, deleting manifests before blobs. In this case, you must set the following in the environment:

```
# Required to delete manifests before blobs
OCI_DELETE_MANIFEST_BEFORE_BLOBS=1
```
