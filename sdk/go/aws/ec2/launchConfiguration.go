// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a new launch configuration, used for autoscaling groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		ubuntu, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
// 			Filters: []aws.GetAmiFilter{
// 				aws.GetAmiFilter{
// 					Name: "name",
// 					Values: []string{
// 						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
// 					},
// 				},
// 				aws.GetAmiFilter{
// 					Name: "virtualization-type",
// 					Values: []string{
// 						"hvm",
// 					},
// 				},
// 			},
// 			MostRecent: &opt0,
// 			Owners: []string{
// 				"099720109477",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
// 			ImageId:      pulumi.String(ubuntu.Id),
// 			InstanceType: pulumi.String("t2.micro"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Using with AutoScaling Groups
//
// Launch Configurations cannot be updated after creation with the Amazon
// Web Service API. In order to update a Launch Configuration, this provider will
// destroy the existing resource and create a replacement. In order to effectively
// use a Launch Configuration resource with an [AutoScaling Group resource](https://www.terraform.io/docs/providers/aws/r/autoscaling_group.html),
// it's recommended to specify `createBeforeDestroy` in a [lifecycle](https://www.terraform.io/docs/configuration/resources.html#lifecycle) block.
// Either omit the Launch Configuration `name` attribute, or specify a partial name
// with `namePrefix`.  Example:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		ubuntu, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
// 			Filters: []aws.GetAmiFilter{
// 				aws.GetAmiFilter{
// 					Name: "name",
// 					Values: []string{
// 						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
// 					},
// 				},
// 				aws.GetAmiFilter{
// 					Name: "virtualization-type",
// 					Values: []string{
// 						"hvm",
// 					},
// 				},
// 			},
// 			MostRecent: &opt0,
// 			Owners: []string{
// 				"099720109477",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		asConf, err := ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
// 			ImageId:      pulumi.String(ubuntu.Id),
// 			InstanceType: pulumi.String("t2.micro"),
// 			NamePrefix:   pulumi.String("lc-example-"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
// 			LaunchConfiguration: asConf.Name,
// 			MaxSize:             pulumi.Int(2),
// 			MinSize:             pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With this setup this provider generates a unique name for your Launch
// Configuration and can then update the AutoScaling Group without conflict before
// destroying the previous Launch Configuration.
//
// ## Using with Spot Instances
//
// Launch configurations can set the spot instance pricing to be used for the
// Auto Scaling Group to reserve instances. Simply specifying the `spotPrice`
// parameter will set the price on the Launch Configuration which will attempt to
// reserve your instances at this price.  See the [AWS Spot Instance
// documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
// for more information or how to launch [Spot Instances](https://www.terraform.io/docs/providers/aws/r/spot_instance_request.html) with this provider.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		ubuntu, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
// 			Filters: []aws.GetAmiFilter{
// 				aws.GetAmiFilter{
// 					Name: "name",
// 					Values: []string{
// 						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
// 					},
// 				},
// 				aws.GetAmiFilter{
// 					Name: "virtualization-type",
// 					Values: []string{
// 						"hvm",
// 					},
// 				},
// 			},
// 			MostRecent: &opt0,
// 			Owners: []string{
// 				"099720109477",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		asConf, err := ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
// 			ImageId:      pulumi.String(ubuntu.Id),
// 			InstanceType: pulumi.String("m4.large"),
// 			SpotPrice:    pulumi.String("0.001"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
// 			LaunchConfiguration: asConf.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Block devices
//
// Each of the `*_block_device` attributes controls a portion of the AWS
// Launch Configuration's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
// Mapping docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
// to understand the implications of using these attributes.
//
// The `rootBlockDevice` mapping supports the following:
//
// * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
//   or `"io1"`. (Default: `"standard"`).
// * `volumeSize` - (Optional) The size of the volume in gigabytes.
// * `iops` - (Optional) The amount of provisioned
//   [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
//   This must be set with a `volumeType` of `"io1"`.
// * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
//   on instance termination (Default: `true`).
// * `encrypted` - (Optional) Whether the volume should be encrypted or not. (Default: `false`).
//
// Modifying any of the `rootBlockDevice` settings requires resource
// replacement.
//
// Each `ebsBlockDevice` supports the following:
//
// * `deviceName` - (Required) The name of the device to mount.
// * `snapshotId` - (Optional) The Snapshot ID to mount.
// * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
//   or `"io1"`. (Default: `"standard"`).
// * `volumeSize` - (Optional) The size of the volume in gigabytes.
// * `iops` - (Optional) The amount of provisioned
//   [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
//   This must be set with a `volumeType` of `"io1"`.
// * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
//   on instance termination (Default: `true`).
// * `encrypted` - (Optional) Whether the volume should be encrypted or not. Do not use this option if you are using `snapshotId` as the encrypted flag will be determined by the snapshot. (Default: `false`).
//
// Modifying any `ebsBlockDevice` currently requires resource replacement.
//
// Each `ephemeralBlockDevice` supports the following:
//
// * `deviceName` - The name of the block device to mount on the instance.
// * `virtualName` - The [Instance Store Device
//   Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
//   (e.g. `"ephemeral0"`)
//
// Each AWS Instance type has a different set of Instance Store block devices
// available for attachment. AWS [publishes a
// list](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
// of which ephemeral devices are available on each type. The devices are always
// identified by the `virtualName` in the format `"ephemeral{0..N}"`.
//
// > **NOTE:** Changes to `*_block_device` configuration of _existing_ resources
// cannot currently be detected by this provider. After updating to block device
// configuration, resource recreation can be manually triggered by using the
// [`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
type LaunchConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the launch configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress pulumi.BoolPtrOutput `pulumi:"associatePublicIpAddress"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices LaunchConfigurationEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolOutput `pulumi:"ebsOptimized"`
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring pulumi.BoolPtrOutput `pulumi:"enableMonitoring"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices LaunchConfigurationEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile pulumi.StringPtrOutput `pulumi:"iamInstanceProfile"`
	// The EC2 image ID to launch.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The size of instance to launch.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The key name that should be used for the instance.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy pulumi.StringPtrOutput `pulumi:"placementTenancy"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice LaunchConfigurationRootBlockDeviceOutput `pulumi:"rootBlockDevice"`
	// A list of associated security group IDS.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The maximum price to use for reserving spot instances.
	SpotPrice pulumi.StringPtrOutput `pulumi:"spotPrice"`
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 pulumi.StringPtrOutput `pulumi:"userDataBase64"`
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId pulumi.StringPtrOutput `pulumi:"vpcClassicLinkId"`
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups pulumi.StringArrayOutput `pulumi:"vpcClassicLinkSecurityGroups"`
}

// NewLaunchConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLaunchConfiguration(ctx *pulumi.Context,
	name string, args *LaunchConfigurationArgs, opts ...pulumi.ResourceOption) (*LaunchConfiguration, error) {
	if args == nil || args.ImageId == nil {
		return nil, errors.New("missing required argument 'ImageId'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	if args == nil {
		args = &LaunchConfigurationArgs{}
	}
	var resource LaunchConfiguration
	err := ctx.RegisterResource("aws:ec2/launchConfiguration:LaunchConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchConfiguration gets an existing LaunchConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchConfigurationState, opts ...pulumi.ResourceOption) (*LaunchConfiguration, error) {
	var resource LaunchConfiguration
	err := ctx.ReadResource("aws:ec2/launchConfiguration:LaunchConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchConfiguration resources.
type launchConfigurationState struct {
	// The Amazon Resource Name of the launch configuration.
	Arn *string `pulumi:"arn"`
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress *bool `pulumi:"associatePublicIpAddress"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices []LaunchConfigurationEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring *bool `pulumi:"enableMonitoring"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices []LaunchConfigurationEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile *string `pulumi:"iamInstanceProfile"`
	// The EC2 image ID to launch.
	ImageId *string `pulumi:"imageId"`
	// The size of instance to launch.
	InstanceType *string `pulumi:"instanceType"`
	// The key name that should be used for the instance.
	KeyName *string `pulumi:"keyName"`
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy *string `pulumi:"placementTenancy"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice *LaunchConfigurationRootBlockDevice `pulumi:"rootBlockDevice"`
	// A list of associated security group IDS.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The maximum price to use for reserving spot instances.
	SpotPrice *string `pulumi:"spotPrice"`
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData *string `pulumi:"userData"`
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 *string `pulumi:"userDataBase64"`
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId *string `pulumi:"vpcClassicLinkId"`
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups []string `pulumi:"vpcClassicLinkSecurityGroups"`
}

type LaunchConfigurationState struct {
	// The Amazon Resource Name of the launch configuration.
	Arn pulumi.StringPtrInput
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress pulumi.BoolPtrInput
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices LaunchConfigurationEbsBlockDeviceArrayInput
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolPtrInput
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring pulumi.BoolPtrInput
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices LaunchConfigurationEphemeralBlockDeviceArrayInput
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile pulumi.StringPtrInput
	// The EC2 image ID to launch.
	ImageId pulumi.StringPtrInput
	// The size of instance to launch.
	InstanceType pulumi.StringPtrInput
	// The key name that should be used for the instance.
	KeyName pulumi.StringPtrInput
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy pulumi.StringPtrInput
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice LaunchConfigurationRootBlockDevicePtrInput
	// A list of associated security group IDS.
	SecurityGroups pulumi.StringArrayInput
	// The maximum price to use for reserving spot instances.
	SpotPrice pulumi.StringPtrInput
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData pulumi.StringPtrInput
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 pulumi.StringPtrInput
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId pulumi.StringPtrInput
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups pulumi.StringArrayInput
}

func (LaunchConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchConfigurationState)(nil)).Elem()
}

type launchConfigurationArgs struct {
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress *bool `pulumi:"associatePublicIpAddress"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices []LaunchConfigurationEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring *bool `pulumi:"enableMonitoring"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices []LaunchConfigurationEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile interface{} `pulumi:"iamInstanceProfile"`
	// The EC2 image ID to launch.
	ImageId string `pulumi:"imageId"`
	// The size of instance to launch.
	InstanceType string `pulumi:"instanceType"`
	// The key name that should be used for the instance.
	KeyName *string `pulumi:"keyName"`
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy *string `pulumi:"placementTenancy"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice *LaunchConfigurationRootBlockDevice `pulumi:"rootBlockDevice"`
	// A list of associated security group IDS.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The maximum price to use for reserving spot instances.
	SpotPrice *string `pulumi:"spotPrice"`
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData *string `pulumi:"userData"`
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 *string `pulumi:"userDataBase64"`
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId *string `pulumi:"vpcClassicLinkId"`
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups []string `pulumi:"vpcClassicLinkSecurityGroups"`
}

// The set of arguments for constructing a LaunchConfiguration resource.
type LaunchConfigurationArgs struct {
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress pulumi.BoolPtrInput
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices LaunchConfigurationEbsBlockDeviceArrayInput
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolPtrInput
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring pulumi.BoolPtrInput
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices LaunchConfigurationEphemeralBlockDeviceArrayInput
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile pulumi.Input
	// The EC2 image ID to launch.
	ImageId pulumi.StringInput
	// The size of instance to launch.
	InstanceType pulumi.StringInput
	// The key name that should be used for the instance.
	KeyName pulumi.StringPtrInput
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy pulumi.StringPtrInput
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice LaunchConfigurationRootBlockDevicePtrInput
	// A list of associated security group IDS.
	SecurityGroups pulumi.StringArrayInput
	// The maximum price to use for reserving spot instances.
	SpotPrice pulumi.StringPtrInput
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData pulumi.StringPtrInput
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 pulumi.StringPtrInput
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId pulumi.StringPtrInput
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups pulumi.StringArrayInput
}

func (LaunchConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchConfigurationArgs)(nil)).Elem()
}
