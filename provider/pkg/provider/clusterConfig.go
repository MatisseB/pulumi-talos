// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"net/url"

	yaml "gopkg.in/yaml.v3"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1"
)

// The set of arguments for creating a ClusterConfig component resource.
type ClusterConfig struct {
	pulumi.ResourceState
	// Cluster configuration generated yaml.
	ConfigYAML string `pulumi:"configYAML"`
}

// The ClusterConfig component resource.
type ClusterConfigArgs struct {
	// Configures the cluster's name.
	ClusterName string `pulumi:"clusterName"`
	// Provides control plane specific configuration options.
	ControlPlaneConfig *ControlPlaneConfig `pulumi:"controlPlaneConfig"`
}

// The ClusterConfig component resource.
type ControlPlaneConfig struct {
	Endpoint           string `pulumi:"endpoint"`
	LocalAPIServerPort int    `pulumi:"localAPIServerPort"`
}

func (c *ControlPlaneConfig) unmarshallPulumi() (*v1alpha1.ControlPlaneConfig, error) {
	url, err := url.Parse(c.Endpoint)
	return &v1alpha1.ControlPlaneConfig{
		LocalAPIServerPort: c.LocalAPIServerPort,
		Endpoint:           &v1alpha1.Endpoint{url},
	}, err
}

// NewClusterConfig creates a new ClusterConfig component resource.
func NewClusterConfig(ctx *pulumi.Context,
	name string, args *ClusterConfigArgs, opts ...pulumi.ResourceOption) (*ClusterConfig, error) {
	if args == nil {
		args = &ClusterConfigArgs{}
	}

	component := &ClusterConfig{}
	err := ctx.RegisterComponentResource("talos:index:ClusterConfig", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// (v1alpha1.ControlPlaneConfig)
	test := v1alpha1.ClusterConfig{ClusterName: args.ClusterName}

	// Specific types
	if test.ControlPlane, err = args.ControlPlaneConfig.unmarshallPulumi(); err != nil {
		return nil, err
	}
	// Create a bucket and expose a website index document.
	// bucket, err := s3.NewBucket(ctx, name, &s3.BucketArgs{
	// 	Website: s3.BucketWebsiteArgs{
	// 		IndexDocument: pulumi.String("index.html"),
	// 	},
	// }, pulumi.Parent(component))
	// if err != nil {
	// 	return nil, err
	// }

	// Create a bucket object for the index document.
	// if _, err := s3.NewBucketObject(ctx, name, &s3.BucketObjectArgs{
	// 	Bucket:      bucket.ID(),
	// 	Key:         pulumi.String("index.html"),
	// 	Content:     args.IndexContent,
	// 	ContentType: pulumi.String("text/html"),
	// }, pulumi.Parent(bucket)); err != nil {
	// 	return nil, err
	// }

	// Set the access policy for the bucket so all objects are readable.
	// if _, err := s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
	// 	Bucket: bucket.ID(),
	// 	Policy: pulumi.Any(map[string]interface{}{
	// 		"Version": "2012-10-17",
	// 		"Statement": []map[string]interface{}{
	// 			{
	// 				"Effect":    "Allow",
	// 				"Principal": "*",
	// 				"Action": []interface{}{
	// 					"s3:GetObject",
	// 				},
	// 				"Resource": []interface{}{
	// 					pulumi.Sprintf("arn:aws:s3:::%s/*", bucket.ID()), // policy refers to bucket name explicitly
	// 				},
	// 			},
	// 		},
	// 	}),
	// }, pulumi.Parent(bucket)); err != nil {
	// 	return nil, err
	// }

	// component.Bucket = bucket
	// component.WebsiteUrl = bucket.WebsiteEndpoint
	out, err := yaml.Marshal(test)
	if err != nil {
		return nil, err
	}

	component.ConfigYAML = string(out)

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"configYAML": pulumi.StringPtr(component.ConfigYAML),
	}); err != nil {
		return nil, err
	}

	return component, nil
}
