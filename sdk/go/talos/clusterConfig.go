// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package talos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterConfig struct {
	pulumi.ResourceState

	// Provides control plane specific configuration options.
	ConfigYAML pulumi.StringOutput `pulumi:"configYAML"`
}

// NewClusterConfig registers a new resource with the given unique name, arguments, and options.
func NewClusterConfig(ctx *pulumi.Context,
	name string, args *ClusterConfigArgs, opts ...pulumi.ResourceOption) (*ClusterConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ControlPlaneConfig == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlaneConfig'")
	}
	var resource ClusterConfig
	err := ctx.RegisterRemoteComponentResource("talos:index:ClusterConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type clusterConfigArgs struct {
	// Configures the cluster's name.
	ClusterName string `pulumi:"clusterName"`
	// Provides control plane specific configuration options.
	ControlPlaneConfig ControlPlaneConfig `pulumi:"controlPlaneConfig"`
}

// The set of arguments for constructing a ClusterConfig resource.
type ClusterConfigArgs struct {
	// Configures the cluster's name.
	ClusterName pulumi.StringInput
	// Provides control plane specific configuration options.
	ControlPlaneConfig ControlPlaneConfigInput
}

func (ClusterConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterConfigArgs)(nil)).Elem()
}

type ClusterConfigInput interface {
	pulumi.Input

	ToClusterConfigOutput() ClusterConfigOutput
	ToClusterConfigOutputWithContext(ctx context.Context) ClusterConfigOutput
}

func (*ClusterConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterConfig)(nil)).Elem()
}

func (i *ClusterConfig) ToClusterConfigOutput() ClusterConfigOutput {
	return i.ToClusterConfigOutputWithContext(context.Background())
}

func (i *ClusterConfig) ToClusterConfigOutputWithContext(ctx context.Context) ClusterConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterConfigOutput)
}

// ClusterConfigArrayInput is an input type that accepts ClusterConfigArray and ClusterConfigArrayOutput values.
// You can construct a concrete instance of `ClusterConfigArrayInput` via:
//
//          ClusterConfigArray{ ClusterConfigArgs{...} }
type ClusterConfigArrayInput interface {
	pulumi.Input

	ToClusterConfigArrayOutput() ClusterConfigArrayOutput
	ToClusterConfigArrayOutputWithContext(context.Context) ClusterConfigArrayOutput
}

type ClusterConfigArray []ClusterConfigInput

func (ClusterConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterConfig)(nil)).Elem()
}

func (i ClusterConfigArray) ToClusterConfigArrayOutput() ClusterConfigArrayOutput {
	return i.ToClusterConfigArrayOutputWithContext(context.Background())
}

func (i ClusterConfigArray) ToClusterConfigArrayOutputWithContext(ctx context.Context) ClusterConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterConfigArrayOutput)
}

// ClusterConfigMapInput is an input type that accepts ClusterConfigMap and ClusterConfigMapOutput values.
// You can construct a concrete instance of `ClusterConfigMapInput` via:
//
//          ClusterConfigMap{ "key": ClusterConfigArgs{...} }
type ClusterConfigMapInput interface {
	pulumi.Input

	ToClusterConfigMapOutput() ClusterConfigMapOutput
	ToClusterConfigMapOutputWithContext(context.Context) ClusterConfigMapOutput
}

type ClusterConfigMap map[string]ClusterConfigInput

func (ClusterConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterConfig)(nil)).Elem()
}

func (i ClusterConfigMap) ToClusterConfigMapOutput() ClusterConfigMapOutput {
	return i.ToClusterConfigMapOutputWithContext(context.Background())
}

func (i ClusterConfigMap) ToClusterConfigMapOutputWithContext(ctx context.Context) ClusterConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterConfigMapOutput)
}

type ClusterConfigOutput struct{ *pulumi.OutputState }

func (ClusterConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterConfig)(nil)).Elem()
}

func (o ClusterConfigOutput) ToClusterConfigOutput() ClusterConfigOutput {
	return o
}

func (o ClusterConfigOutput) ToClusterConfigOutputWithContext(ctx context.Context) ClusterConfigOutput {
	return o
}

type ClusterConfigArrayOutput struct{ *pulumi.OutputState }

func (ClusterConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterConfig)(nil)).Elem()
}

func (o ClusterConfigArrayOutput) ToClusterConfigArrayOutput() ClusterConfigArrayOutput {
	return o
}

func (o ClusterConfigArrayOutput) ToClusterConfigArrayOutputWithContext(ctx context.Context) ClusterConfigArrayOutput {
	return o
}

func (o ClusterConfigArrayOutput) Index(i pulumi.IntInput) ClusterConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterConfig {
		return vs[0].([]*ClusterConfig)[vs[1].(int)]
	}).(ClusterConfigOutput)
}

type ClusterConfigMapOutput struct{ *pulumi.OutputState }

func (ClusterConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterConfig)(nil)).Elem()
}

func (o ClusterConfigMapOutput) ToClusterConfigMapOutput() ClusterConfigMapOutput {
	return o
}

func (o ClusterConfigMapOutput) ToClusterConfigMapOutputWithContext(ctx context.Context) ClusterConfigMapOutput {
	return o
}

func (o ClusterConfigMapOutput) MapIndex(k pulumi.StringInput) ClusterConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterConfig {
		return vs[0].(map[string]*ClusterConfig)[vs[1].(string)]
	}).(ClusterConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterConfigInput)(nil)).Elem(), &ClusterConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterConfigArrayInput)(nil)).Elem(), ClusterConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterConfigMapInput)(nil)).Elem(), ClusterConfigMap{})
	pulumi.RegisterOutputType(ClusterConfigOutput{})
	pulumi.RegisterOutputType(ClusterConfigArrayOutput{})
	pulumi.RegisterOutputType(ClusterConfigMapOutput{})
}
