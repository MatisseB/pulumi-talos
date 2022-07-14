# Pulumi Component Provider Boilerplate (Go)

This repo is a boilerplate showing how to create a Pulumi component provider written in Go. You can search-replace `talos` with the name of your desired provider as a starting point for creating a component provider for your component resources.

## Background
This repository is part of the [guide for authoring and publishing a Pulumi Package](https://www.pulumi.com/docs/guides/pulumi-packages/how-to-author).

Learn about the concepts behind [Pulumi Packages](https://www.pulumi.com/docs/guides/pulumi-packages/#pulumi-packages) and, more specifically, [Pulumi Components](https://www.pulumi.com/docs/intro/concepts/resources/components/)

## Sample talos Component Provider

An example `ClusterConfig` [component resource](https://www.pulumi.com/docs/intro/concepts/resources/#components) is available in `provider/pkg/provider/clusterConfig.go`. This component creates a static web page hosted in an AWS S3 Bucket. There is nothing special about `ClusterConfig` -- it is a typical component resource written in Go.

The component provider makes component resources available to other languages. The implementation is in `provider/pkg/provider/provider.go`. Each component resource in the provider must have an implementation in the `Construct` function to create an instance of the requested component resource and return its `URN` and state (outputs). There is an initial implementation that demonstrates an implementation of `Construct` for the example `ClusterConfig` component.

A code generator is available which generates SDKs in TypeScript, Python, Go and .NET which are also checked in to the `sdk` folder. The SDKs are generated from a schema in `schema.json`. This file should be kept aligned with the component resources supported by the component provider implementation.

An example of using the `ClusterConfig` component in TypeScript is in `examples/simple`.

Note that the generated provider plugin (`pulumi-resource-talos`) must be on your `PATH` to be used by Pulumi deployments. If creating a provider for distribution to other users, you should ensure they install this plugin to their `PATH`.

## Prerequisites

- Go 1.17
- Pulumi CLI
- Node.js (to build the Node.js SDK)
- Yarn (to build the Node.js SDK)
- Python 3.6+ (to build the Python SDK)
- .NET Core SDK (to build the .NET SDK)

## Build and Test

```bash
# Build and install the provider (plugin copied to $GOPATH/bin)
make install_provider

# Regenerate SDKs
make generate

# Test Node.js SDK
$ make install_nodejs_sdk
$ cd examples/simple
$ yarn install
$ yarn link @pulumi/talos
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up
```

## Naming

The `talos` provider's plugin binary must be named `pulumi-resource-talos` (in the format `pulumi-resource-<provider>`).

While the provider plugin must follow this naming convention, the SDK package naming can be customized. TODO explain.

## Configuring CI and releases

1. Follow the instructions laid out in the [deployment templates](./deployment-templates/README-DEPLOYMENT.md).

## Example component

Let's look at the example `ClusterConfig` component resource in more detail.

### Schema

The example `ClusterConfig` component resource is defined in `schema.json`:

```yaml
resources:
  talos:index:ClusterConfig:
    isComponent: true
    inputProperties:
      indexContent:
        type: string
        description: The HTML content for index.html.
    requiredInputs:
      - indexContent
    properties:
      bucket:
        "$ref": "/aws/v4.0.0/schema.json#/resources/aws:s3%2Fbucket:Bucket"
        description: The bucket resource.
      websiteUrl:
        type: string
        description: The website URL.
    required:
      - bucket
      - websiteUrl
```

The component resource's type token is `talos:index:ClusterConfig` in the format of `<package>:<module>:<type>`. In this case, it's in the `talos` package and `index` module. This is the same type token passed to `RegisterComponentResource` inside the implementation of `NewClusterConfig` in `provider/pkg/provider/clusterConfig.go`, and also the same token referenced in `Construct` in `provider/pkg/provider/provider.go`.

This component has a required `indexContent` input property typed as `string`, and two required output properties: `bucket` and `websiteUrl`. Note that `bucket` is typed as the `aws:s3/bucket:Bucket` resource from the `aws` provider (in the schema the `/` is escaped as `%2F`).

Since this component returns a type from the `aws` provider, each SDK must reference the associated Pulumi `aws` SDK for the language. For the .NET, Node.js, and Python SDKs, dependencies are specified in the `language` section of the schema:

```yaml
language:
  csharp:
    packageReferences:
      Pulumi: 3.*
      Pulumi.Aws: 4.*
  go:
    generateResourceContainerTypes: true
    importBasePath: github.com/MatisseB/pulumi-talos/sdk/go/talos
  nodejs:
    dependencies:
      "@pulumi/aws": "^4.0.0"
    devDependencies:
      typescript: "^3.7.0"
  python:
    requires:
      pulumi: ">=3.0.0,<4.0.0"
      pulumi-aws: ">=4.0.0,<5.0.0"
```

For the Go SDK, dependencies are specified in the `sdk/go.mod` file.

### Implementation

The implementation of this component is in `provider/pkg/provider/clusterConfig.go` and the structure of the component's inputs and outputs aligns with what is defined in `schema.json`:

```go
// The set of arguments for creating a ClusterConfig component resource.
type ClusterConfigArgs struct {
    IndexContent pulumi.StringInput `pulumi:"indexContent"`
}

// The ClusterConfig component resource.
type ClusterConfig struct {
    pulumi.ResourceState

    Bucket     *s3.Bucket          `pulumi:"bucket"`
    WebsiteUrl pulumi.StringOutput `pulumi:"websiteUrl"`
}

// NewClusterConfig creates a new ClusterConfig component resource.
func NewClusterConfig(ctx *pulumi.Context, name string, args *ClusterConfigArgs, opts ...pulumi.ResourceOption) (*ClusterConfig, error) {
    ...
}
```

The provider makes this component resource available in the `construct` function in `provider/pkg/provider/provider.go`. When `construct` is called and the `typ` argument is `talos:index:ClusterConfig`, we create an instance of the `ClusterConfig` component resource and return its `URN` and state.

```go
func constructClusterConfig(ctx *pulumi.Context, name string, inputs provider.ConstructInputs,
    options pulumi.ResourceOption) (*provider.ConstructResult, error) {

    // Copy the raw inputs to ClusterConfigArgs. `inputs.CopyTo` uses the types and `pulumi:` tags
    // on the struct's fields to convert the raw values to the appropriate Input types.
    args := &ClusterConfigArgs{}
    if err := inputs.CopyTo(args); err != nil {
        return nil, errors.Wrap(err, "setting args")
    }

    // Create the component resource.
    clusterConfig, err := NewClusterConfig(ctx, name, args, options)
    if err != nil {
        return nil, errors.Wrap(err, "creating component")
    }

    // Return the component resource's URN and state. `NewConstructResult` automatically sets the
    // ConstructResult's state based on resource struct fields tagged with `pulumi:` tags with a value
    // that is convertible to `pulumi.Input`.
    return provider.NewConstructResult(clusterConfig)
}
```