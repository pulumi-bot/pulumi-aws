// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
// with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
// for more information.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = lightsail.NewInstance(ctx, "gitlabTest", &lightsail.InstanceArgs{
// 			AvailabilityZone: pulumi.String("us-east-1b"),
// 			BlueprintId:      pulumi.String("string"),
// 			BundleId:         pulumi.String("string"),
// 			KeyPairName:      pulumi.String("some_key_name"),
// 			Tags: pulumi.Map{
// 				"foo": pulumi.String("bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
// - `amazonLinux20180302`
// - `centos7190101`
// - `debian87`
// - `debian95`
// - `freebsd111`
// - `opensuse422`
// - `ubuntu16042`
// - `ubuntu1804`
//
// ### Apps and OS
//
// - `drupal856`
// - `gitlab11141`
// - `joomla3811`
// - `lamp56372`
// - `lamp71201`
// - `magento225`
// - `mean401`
// - `nginx11401`
// - `nodejs1080`
// - `pleskUbuntu178111`
// - `redmine346`
// - `wordpress498`
// - `wordpressMultisite498`
//
// ## Bundles
//
// Lightsail currently supports the following Bundle IDs (e.g. an instance in `ap-northeast-1` would use `small20`):
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
// - ap-northeast-1: `20`
// - ap-northeast-2: `20`
// - ap-south-1: `21`
// - ap-southeast-1: `20`
// - ap-southeast-2: `22`
// - ca-central-1: `20`
// - eu-central-1: `20`
// - eu-west-1: `20`
// - eu-west-2: `20`
// - eu-west-3: `20`
// - us-east-1: `20`
// - us-east-2: `20`
// - us-west-2: `20`
type Instance struct {
	pulumi.CustomResourceState

	// The ARN of the Lightsail instance (matches `id`).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId pulumi.StringOutput `pulumi:"blueprintId"`
	// The bundle of specification information (see list below)
	BundleId pulumi.StringOutput `pulumi:"bundleId"`
	CpuCount pulumi.IntOutput    `pulumi:"cpuCount"`
	// The timestamp when the instance was created.
	// * `availabilityZone`
	// * `blueprintId`
	// * `bundleId`
	// * `keyPairName`
	// * `userData`
	CreatedAt   pulumi.StringOutput `pulumi:"createdAt"`
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	IsStaticIp  pulumi.BoolOutput   `pulumi:"isStaticIp"`
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `ec2.KeyPair` at this time)
	KeyPairName pulumi.StringPtrOutput `pulumi:"keyPairName"`
	// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
	Name             pulumi.StringOutput  `pulumi:"name"`
	PrivateIpAddress pulumi.StringOutput  `pulumi:"privateIpAddress"`
	PublicIpAddress  pulumi.StringOutput  `pulumi:"publicIpAddress"`
	RamSize          pulumi.Float64Output `pulumi:"ramSize"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// launch script to configure server with additional user data
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	Username pulumi.StringOutput    `pulumi:"username"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.AvailabilityZone == nil {
		return nil, errors.New("missing required argument 'AvailabilityZone'")
	}
	if args == nil || args.BlueprintId == nil {
		return nil, errors.New("missing required argument 'BlueprintId'")
	}
	if args == nil || args.BundleId == nil {
		return nil, errors.New("missing required argument 'BundleId'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("aws:lightsail/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("aws:lightsail/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The ARN of the Lightsail instance (matches `id`).
	Arn *string `pulumi:"arn"`
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId *string `pulumi:"blueprintId"`
	// The bundle of specification information (see list below)
	BundleId *string `pulumi:"bundleId"`
	CpuCount *int    `pulumi:"cpuCount"`
	// The timestamp when the instance was created.
	// * `availabilityZone`
	// * `blueprintId`
	// * `bundleId`
	// * `keyPairName`
	// * `userData`
	CreatedAt   *string `pulumi:"createdAt"`
	Ipv6Address *string `pulumi:"ipv6Address"`
	IsStaticIp  *bool   `pulumi:"isStaticIp"`
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `ec2.KeyPair` at this time)
	KeyPairName *string `pulumi:"keyPairName"`
	// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
	Name             *string  `pulumi:"name"`
	PrivateIpAddress *string  `pulumi:"privateIpAddress"`
	PublicIpAddress  *string  `pulumi:"publicIpAddress"`
	RamSize          *float64 `pulumi:"ramSize"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// launch script to configure server with additional user data
	UserData *string `pulumi:"userData"`
	Username *string `pulumi:"username"`
}

type InstanceState struct {
	// The ARN of the Lightsail instance (matches `id`).
	Arn pulumi.StringPtrInput
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone pulumi.StringPtrInput
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId pulumi.StringPtrInput
	// The bundle of specification information (see list below)
	BundleId pulumi.StringPtrInput
	CpuCount pulumi.IntPtrInput
	// The timestamp when the instance was created.
	// * `availabilityZone`
	// * `blueprintId`
	// * `bundleId`
	// * `keyPairName`
	// * `userData`
	CreatedAt   pulumi.StringPtrInput
	Ipv6Address pulumi.StringPtrInput
	IsStaticIp  pulumi.BoolPtrInput
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `ec2.KeyPair` at this time)
	KeyPairName pulumi.StringPtrInput
	// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
	Name             pulumi.StringPtrInput
	PrivateIpAddress pulumi.StringPtrInput
	PublicIpAddress  pulumi.StringPtrInput
	RamSize          pulumi.Float64PtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// launch script to configure server with additional user data
	UserData pulumi.StringPtrInput
	Username pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId string `pulumi:"blueprintId"`
	// The bundle of specification information (see list below)
	BundleId string `pulumi:"bundleId"`
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `ec2.KeyPair` at this time)
	KeyPairName *string `pulumi:"keyPairName"`
	// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// launch script to configure server with additional user data
	UserData *string `pulumi:"userData"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The Availability Zone in which to create your
	// instance (see list below)
	AvailabilityZone pulumi.StringInput
	// The ID for a virtual private server image
	// (see list below)
	BlueprintId pulumi.StringInput
	// The bundle of specification information (see list below)
	BundleId pulumi.StringInput
	// The name of your key pair. Created in the
	// Lightsail console (cannot use `ec2.KeyPair` at this time)
	KeyPairName pulumi.StringPtrInput
	// The name of the Lightsail Instance. Names be unique within each AWS Region in your Lightsail account.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// launch script to configure server with additional user data
	UserData pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
