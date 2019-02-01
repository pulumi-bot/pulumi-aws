// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
// with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
// for more information.
// 
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
// 
// ## Availability Zones
// 
// Lightsail currently supports the following Availability Zones (e.g. `us-east-1a`):
// 
// - `ap-northeast-1{a,c,d}`
// - `ap-northeast-2{a,c}`
// - `ap-south-1{a,b}`
// - `ap-southeast-1{a,b,c}`
// - `ap-southeast-2{a,b,c}`
// - `ca-central-1{a,b}`
// - `eu-central-1{a,b,c}`
// - `eu-west-1{a,b,c}`
// - `eu-west-2{a,b,c}`
// - `eu-west-3{a,b,c}`
// - `us-east-1{a,b,c,d,e,f}`
// - `us-east-2{a,b,c}`
// - `us-west-2{a,b,c}`
// 
// ## Blueprints
// 
// Lightsail currently supports the following Blueprint IDs:
// 
// ### OS Only
// 
// - `amazon_linux_2018_03_0_2`
// - `centos_7_1805_01`
// - `debian_8_7`
// - `debian_9_5`
// - `freebsd_11_1`
// - `opensuse_42_2`
// - `ubuntu_16_04_2`
// - `ubuntu_18_04`
// 
// ### Apps and OS
// 
// - `drupal_8_5_6`
// - `gitlab_11_1_4_1`
// - `joomla_3_8_11`
// - `lamp_5_6_37_2`
// - `lamp_7_1_20_1`
// - `magento_2_2_5`
// - `mean_4_0_1`
// - `nginx_1_14_0_1`
// - `nodejs_10_8_0`
// - `plesk_ubuntu_17_8_11_1`
// - `redmine_3_4_6`
// - `wordpress_4_9_8`
// - `wordpress_multisite_4_9_8`
// 
// ## Bundles
// 
// Lightsail currently supports the following Bundle IDs (e.g. an instance in `ap-northeast-1` would use `small_2_0`):
// 
// ### Prefix
// 
// A Bundle ID starts with one of the below size prefixes:
// 
// - `nano_`
// - `micro_`
// - `small_`
// - `medium_`
// - `large_`
// - `xlarge_`
// - `2xlarge_`
// 
// ### Suffix
// 
// A Bundle ID ends with one of the following suffixes depending on Availability Zone:
// 
// - ap-northeast-1: `2_0`
// - ap-northeast-2: `2_0`
// - ap-south-1: `2_1`
// - ap-southeast-1: `2_0`
// - ap-southeast-2: `2_2`
// - ca-central-1: `2_0`
// - eu-central-1: `2_0`
// - eu-west-1: `2_0`
// - eu-west-2: `2_0`
// - eu-west-3: `2_0`
// - us-east-1: `2_0`
// - us-east-2: `2_0`
// - us-west-2: `2_0`
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil || args.BlueprintId == nil {
		return nil, errors.New("missing required argument 'BlueprintId'")
	}
	if args == nil || args.BundleId == nil {
		return nil, errors.New("missing required argument 'BundleId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["availabilityZone"] = nil
		inputs["blueprintId"] = nil
		inputs["bundleId"] = nil
		inputs["keyPairName"] = nil
		inputs["name"] = nil
		inputs["userData"] = nil
	} else {
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["blueprintId"] = args.BlueprintId
		inputs["bundleId"] = args.BundleId
		inputs["keyPairName"] = args.KeyPairName
		inputs["name"] = args.Name
		inputs["userData"] = args.UserData
	}
	inputs["arn"] = nil
	inputs["cpuCount"] = nil
	inputs["createdAt"] = nil
	inputs["ipv6Address"] = nil
	inputs["isStaticIp"] = nil
	inputs["privateIpAddress"] = nil
	inputs["publicIpAddress"] = nil
	inputs["ramSize"] = nil
	inputs["username"] = nil
	s, err := ctx.RegisterResource("aws:lightsail/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["blueprintId"] = state.BlueprintId
		inputs["bundleId"] = state.BundleId
		inputs["cpuCount"] = state.CpuCount
		inputs["createdAt"] = state.CreatedAt
		inputs["ipv6Address"] = state.Ipv6Address
		inputs["isStaticIp"] = state.IsStaticIp
		inputs["keyPairName"] = state.KeyPairName
		inputs["name"] = state.Name
		inputs["privateIpAddress"] = state.PrivateIpAddress
		inputs["publicIpAddress"] = state.PublicIpAddress
		inputs["ramSize"] = state.RamSize
		inputs["userData"] = state.UserData
		inputs["username"] = state.Username
	}
	s, err := ctx.ReadResource("aws:lightsail/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the Lightsail instance (matches `id`).
// * `availability_zone`
// * `blueprint_id`
// * `bundle_id`
// * `key_pair_name`
// * `user_data`
func (r *Instance) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The Availability Zone in which to create your
// instance (see list below)
func (r *Instance) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The ID for a virtual private server image
// (see list below)
func (r *Instance) BlueprintId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["blueprintId"])
}

// The bundle of specification information (see list below)
func (r *Instance) BundleId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bundleId"])
}

func (r *Instance) CpuCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuCount"])
}

func (r *Instance) CreatedAt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createdAt"])
}

func (r *Instance) Ipv6Address() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6Address"])
}

func (r *Instance) IsStaticIp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["isStaticIp"])
}

// The name of your key pair. Created in the
// Lightsail console (cannot use `aws_key_pair` at this time)
func (r *Instance) KeyPairName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyPairName"])
}

// The name of the Lightsail Instance
func (r *Instance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Instance) PrivateIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIpAddress"])
}

func (r *Instance) PublicIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicIpAddress"])
}

func (r *Instance) RamSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ramSize"])
}

// launch script to configure server with additional user data
func (r *Instance) UserData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userData"])
}

func (r *Instance) Username() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["username"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// The ARN of the Lightsail instance (matches `id`).
	// * `availability_zone`
	// * `blueprint_id`
	// * `bundle_id`
	// * `key_pair_name`
	// * `user_data`
	Arn interface{}
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone interface{}
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId interface{}
	// The bundle of specification information (see list below)
	BundleId interface{}
	CpuCount interface{}
	CreatedAt interface{}
	Ipv6Address interface{}
	IsStaticIp interface{}
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `aws_key_pair` at this time)
	KeyPairName interface{}
	// The name of the Lightsail Instance
	Name interface{}
	PrivateIpAddress interface{}
	PublicIpAddress interface{}
	RamSize interface{}
	// launch script to configure server with additional user data
	UserData interface{}
	Username interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone interface{}
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId interface{}
	// The bundle of specification information (see list below)
	BundleId interface{}
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `aws_key_pair` at this time)
	KeyPairName interface{}
	// The name of the Lightsail Instance
	Name interface{}
	// launch script to configure server with additional user data
	UserData interface{}
}
