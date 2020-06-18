// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks instance resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = opsworks.NewInstance(ctx, "my-instance", &opsworks.InstanceArgs{
// 			InstanceType: pulumi.String("t2.micro"),
// 			LayerIds: pulumi.StringArray{
// 				pulumi.String(aws_opsworks_custom_layer.My - layer.Id),
// 			},
// 			Os:      pulumi.String("Amazon Linux 2015.09"),
// 			StackId: pulumi.String(aws_opsworks_stack.Main.Id),
// 			State:   pulumi.String("stopped"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Block devices
//
// Each of the `*_block_device` attributes controls a portion of the AWS
// Instance's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
// Mapping docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
// to understand the implications of using these attributes.
//
// The `rootBlockDevice` mapping supports the following:
//
// * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
//   or `"io1"`. (Default: `"standard"`).
// * `volumeSize` - (Optional) The size of the volume in gigabytes.
// * `iops` - (Optional) The amount of provisioned
//   [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
//   This must be set with a `volumeType` of `"io1"`.
// * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
//   on instance termination (Default: `true`).
//
// Modifying any of the `rootBlockDevice` settings requires resource
// replacement.
//
// Each `ebsBlockDevice` supports the following:
//
// * `deviceName` - The name of the device to mount.
// * `snapshotId` - (Optional) The Snapshot ID to mount.
// * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
//   or `"io1"`. (Default: `"standard"`).
// * `volumeSize` - (Optional) The size of the volume in gigabytes.
// * `iops` - (Optional) The amount of provisioned
//   [IOPS](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
//   This must be set with a `volumeType` of `"io1"`.
// * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
//   on instance termination (Default: `true`).
//
// Modifying any `ebsBlockDevice` currently requires resource replacement.
//
// Each `ephemeralBlockDevice` supports the following:
//
// * `deviceName` - The name of the block device to mount on the instance.
// * `virtualName` - The [Instance Store Device
//   Name](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
//   (e.g. `"ephemeral0"`)
//
// Each AWS Instance type has a different set of Instance Store block devices
// available for attachment. AWS [publishes a
// list](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
// of which ephemeral devices are available on each type. The devices are always
// identified by the `virtualName` in the format `"ephemeral{0..N}"`.
//
// > **NOTE:** Currently, changes to `*_block_device` configuration of _existing_
// resources cannot be automatically detected by this provider. After making updates
// to block device configuration, resource recreation can be manually triggered by
// using the [`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
type Instance struct {
	pulumi.CustomResourceState

	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion pulumi.StringPtrOutput `pulumi:"agentVersion"`
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId pulumi.StringOutput `pulumi:"amiId"`
	// Machine architecture for created instances.  Can be either `"x8664"` (the default) or `"i386"`
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType pulumi.StringPtrOutput `pulumi:"autoScalingType"`
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone pulumi.StringOutput  `pulumi:"availabilityZone"`
	CreatedAt        pulumi.StringOutput  `pulumi:"createdAt"`
	DeleteEbs        pulumi.BoolPtrOutput `pulumi:"deleteEbs"`
	DeleteEip        pulumi.BoolPtrOutput `pulumi:"deleteEip"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices InstanceEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolPtrOutput `pulumi:"ebsOptimized"`
	// EC2 instance ID
	Ec2InstanceId pulumi.StringOutput `pulumi:"ec2InstanceId"`
	EcsClusterArn pulumi.StringOutput `pulumi:"ecsClusterArn"`
	ElasticIp     pulumi.StringOutput `pulumi:"elasticIp"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices InstanceEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	// The instance's host name.
	Hostname            pulumi.StringOutput `pulumi:"hostname"`
	InfrastructureClass pulumi.StringOutput `pulumi:"infrastructureClass"`
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	InstanceProfileArn   pulumi.StringOutput  `pulumi:"instanceProfileArn"`
	// The type of instance to start
	InstanceType       pulumi.StringPtrOutput `pulumi:"instanceType"`
	LastServiceErrorId pulumi.StringOutput    `pulumi:"lastServiceErrorId"`
	// The ids of the layers the instance will belong to.
	LayerIds pulumi.StringArrayOutput `pulumi:"layerIds"`
	// Name of operating system that will be installed.
	Os       pulumi.StringOutput `pulumi:"os"`
	Platform pulumi.StringOutput `pulumi:"platform"`
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns pulumi.StringOutput `pulumi:"privateDns"`
	// The private IP address assigned to the instance
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns pulumi.StringOutput `pulumi:"publicDns"`
	// The public IP address assigned to the instance, if applicable.
	PublicIp             pulumi.StringOutput `pulumi:"publicIp"`
	RegisteredBy         pulumi.StringOutput `pulumi:"registeredBy"`
	ReportedAgentVersion pulumi.StringOutput `pulumi:"reportedAgentVersion"`
	ReportedOsFamily     pulumi.StringOutput `pulumi:"reportedOsFamily"`
	ReportedOsName       pulumi.StringOutput `pulumi:"reportedOsName"`
	ReportedOsVersion    pulumi.StringOutput `pulumi:"reportedOsVersion"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices InstanceRootBlockDeviceArrayOutput `pulumi:"rootBlockDevices"`
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType     pulumi.StringOutput `pulumi:"rootDeviceType"`
	RootDeviceVolumeId pulumi.StringOutput `pulumi:"rootDeviceVolumeId"`
	// The associated security groups.
	SecurityGroupIds         pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	SshHostDsaKeyFingerprint pulumi.StringOutput      `pulumi:"sshHostDsaKeyFingerprint"`
	SshHostRsaKeyFingerprint pulumi.StringOutput      `pulumi:"sshHostRsaKeyFingerprint"`
	// Name of the SSH keypair that instances will have by default.
	SshKeyName pulumi.StringOutput `pulumi:"sshKeyName"`
	// The id of the stack the instance will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State  pulumi.StringPtrOutput `pulumi:"state"`
	Status pulumi.StringOutput    `pulumi:"status"`
	// Subnet ID to attach to
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy pulumi.StringOutput `pulumi:"tenancy"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType pulumi.StringOutput `pulumi:"virtualizationType"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.LayerIds == nil {
		return nil, errors.New("missing required argument 'LayerIds'")
	}
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("aws:opsworks/instance:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:opsworks/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion *string `pulumi:"agentVersion"`
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId *string `pulumi:"amiId"`
	// Machine architecture for created instances.  Can be either `"x8664"` (the default) or `"i386"`
	Architecture *string `pulumi:"architecture"`
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType *string `pulumi:"autoScalingType"`
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	CreatedAt        *string `pulumi:"createdAt"`
	DeleteEbs        *bool   `pulumi:"deleteEbs"`
	DeleteEip        *bool   `pulumi:"deleteEip"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices []InstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// EC2 instance ID
	Ec2InstanceId *string `pulumi:"ec2InstanceId"`
	EcsClusterArn *string `pulumi:"ecsClusterArn"`
	ElasticIp     *string `pulumi:"elasticIp"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices []InstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// The instance's host name.
	Hostname            *string `pulumi:"hostname"`
	InfrastructureClass *string `pulumi:"infrastructureClass"`
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot *bool   `pulumi:"installUpdatesOnBoot"`
	InstanceProfileArn   *string `pulumi:"instanceProfileArn"`
	// The type of instance to start
	InstanceType       *string `pulumi:"instanceType"`
	LastServiceErrorId *string `pulumi:"lastServiceErrorId"`
	// The ids of the layers the instance will belong to.
	LayerIds []string `pulumi:"layerIds"`
	// Name of operating system that will be installed.
	Os       *string `pulumi:"os"`
	Platform *string `pulumi:"platform"`
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns *string `pulumi:"privateDns"`
	// The private IP address assigned to the instance
	PrivateIp *string `pulumi:"privateIp"`
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns *string `pulumi:"publicDns"`
	// The public IP address assigned to the instance, if applicable.
	PublicIp             *string `pulumi:"publicIp"`
	RegisteredBy         *string `pulumi:"registeredBy"`
	ReportedAgentVersion *string `pulumi:"reportedAgentVersion"`
	ReportedOsFamily     *string `pulumi:"reportedOsFamily"`
	ReportedOsName       *string `pulumi:"reportedOsName"`
	ReportedOsVersion    *string `pulumi:"reportedOsVersion"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices []InstanceRootBlockDevice `pulumi:"rootBlockDevices"`
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType     *string `pulumi:"rootDeviceType"`
	RootDeviceVolumeId *string `pulumi:"rootDeviceVolumeId"`
	// The associated security groups.
	SecurityGroupIds         []string `pulumi:"securityGroupIds"`
	SshHostDsaKeyFingerprint *string  `pulumi:"sshHostDsaKeyFingerprint"`
	SshHostRsaKeyFingerprint *string  `pulumi:"sshHostRsaKeyFingerprint"`
	// Name of the SSH keypair that instances will have by default.
	SshKeyName *string `pulumi:"sshKeyName"`
	// The id of the stack the instance will belong to.
	StackId *string `pulumi:"stackId"`
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State  *string `pulumi:"state"`
	Status *string `pulumi:"status"`
	// Subnet ID to attach to
	SubnetId *string `pulumi:"subnetId"`
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy *string `pulumi:"tenancy"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

type InstanceState struct {
	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion pulumi.StringPtrInput
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId pulumi.StringPtrInput
	// Machine architecture for created instances.  Can be either `"x8664"` (the default) or `"i386"`
	Architecture pulumi.StringPtrInput
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType pulumi.StringPtrInput
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone pulumi.StringPtrInput
	CreatedAt        pulumi.StringPtrInput
	DeleteEbs        pulumi.BoolPtrInput
	DeleteEip        pulumi.BoolPtrInput
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices InstanceEbsBlockDeviceArrayInput
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolPtrInput
	// EC2 instance ID
	Ec2InstanceId pulumi.StringPtrInput
	EcsClusterArn pulumi.StringPtrInput
	ElasticIp     pulumi.StringPtrInput
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices InstanceEphemeralBlockDeviceArrayInput
	// The instance's host name.
	Hostname            pulumi.StringPtrInput
	InfrastructureClass pulumi.StringPtrInput
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	InstanceProfileArn   pulumi.StringPtrInput
	// The type of instance to start
	InstanceType       pulumi.StringPtrInput
	LastServiceErrorId pulumi.StringPtrInput
	// The ids of the layers the instance will belong to.
	LayerIds pulumi.StringArrayInput
	// Name of operating system that will be installed.
	Os       pulumi.StringPtrInput
	Platform pulumi.StringPtrInput
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns pulumi.StringPtrInput
	// The private IP address assigned to the instance
	PrivateIp pulumi.StringPtrInput
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns pulumi.StringPtrInput
	// The public IP address assigned to the instance, if applicable.
	PublicIp             pulumi.StringPtrInput
	RegisteredBy         pulumi.StringPtrInput
	ReportedAgentVersion pulumi.StringPtrInput
	ReportedOsFamily     pulumi.StringPtrInput
	ReportedOsName       pulumi.StringPtrInput
	ReportedOsVersion    pulumi.StringPtrInput
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices InstanceRootBlockDeviceArrayInput
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType     pulumi.StringPtrInput
	RootDeviceVolumeId pulumi.StringPtrInput
	// The associated security groups.
	SecurityGroupIds         pulumi.StringArrayInput
	SshHostDsaKeyFingerprint pulumi.StringPtrInput
	SshHostRsaKeyFingerprint pulumi.StringPtrInput
	// Name of the SSH keypair that instances will have by default.
	SshKeyName pulumi.StringPtrInput
	// The id of the stack the instance will belong to.
	StackId pulumi.StringPtrInput
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State  pulumi.StringPtrInput
	Status pulumi.StringPtrInput
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrInput
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy pulumi.StringPtrInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion *string `pulumi:"agentVersion"`
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId *string `pulumi:"amiId"`
	// Machine architecture for created instances.  Can be either `"x8664"` (the default) or `"i386"`
	Architecture *string `pulumi:"architecture"`
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType *string `pulumi:"autoScalingType"`
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	CreatedAt        *string `pulumi:"createdAt"`
	DeleteEbs        *bool   `pulumi:"deleteEbs"`
	DeleteEip        *bool   `pulumi:"deleteEip"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices []InstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized  *bool   `pulumi:"ebsOptimized"`
	EcsClusterArn *string `pulumi:"ecsClusterArn"`
	ElasticIp     *string `pulumi:"elasticIp"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices []InstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// The instance's host name.
	Hostname            *string `pulumi:"hostname"`
	InfrastructureClass *string `pulumi:"infrastructureClass"`
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot *bool   `pulumi:"installUpdatesOnBoot"`
	InstanceProfileArn   *string `pulumi:"instanceProfileArn"`
	// The type of instance to start
	InstanceType       *string `pulumi:"instanceType"`
	LastServiceErrorId *string `pulumi:"lastServiceErrorId"`
	// The ids of the layers the instance will belong to.
	LayerIds []string `pulumi:"layerIds"`
	// Name of operating system that will be installed.
	Os       *string `pulumi:"os"`
	Platform *string `pulumi:"platform"`
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns *string `pulumi:"privateDns"`
	// The private IP address assigned to the instance
	PrivateIp *string `pulumi:"privateIp"`
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns *string `pulumi:"publicDns"`
	// The public IP address assigned to the instance, if applicable.
	PublicIp             *string `pulumi:"publicIp"`
	RegisteredBy         *string `pulumi:"registeredBy"`
	ReportedAgentVersion *string `pulumi:"reportedAgentVersion"`
	ReportedOsFamily     *string `pulumi:"reportedOsFamily"`
	ReportedOsName       *string `pulumi:"reportedOsName"`
	ReportedOsVersion    *string `pulumi:"reportedOsVersion"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices []InstanceRootBlockDevice `pulumi:"rootBlockDevices"`
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType     *string `pulumi:"rootDeviceType"`
	RootDeviceVolumeId *string `pulumi:"rootDeviceVolumeId"`
	// The associated security groups.
	SecurityGroupIds         []string `pulumi:"securityGroupIds"`
	SshHostDsaKeyFingerprint *string  `pulumi:"sshHostDsaKeyFingerprint"`
	SshHostRsaKeyFingerprint *string  `pulumi:"sshHostRsaKeyFingerprint"`
	// Name of the SSH keypair that instances will have by default.
	SshKeyName *string `pulumi:"sshKeyName"`
	// The id of the stack the instance will belong to.
	StackId string `pulumi:"stackId"`
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State  *string `pulumi:"state"`
	Status *string `pulumi:"status"`
	// Subnet ID to attach to
	SubnetId *string `pulumi:"subnetId"`
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy *string `pulumi:"tenancy"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion pulumi.StringPtrInput
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId pulumi.StringPtrInput
	// Machine architecture for created instances.  Can be either `"x8664"` (the default) or `"i386"`
	Architecture pulumi.StringPtrInput
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType pulumi.StringPtrInput
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone pulumi.StringPtrInput
	CreatedAt        pulumi.StringPtrInput
	DeleteEbs        pulumi.BoolPtrInput
	DeleteEip        pulumi.BoolPtrInput
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices InstanceEbsBlockDeviceArrayInput
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized  pulumi.BoolPtrInput
	EcsClusterArn pulumi.StringPtrInput
	ElasticIp     pulumi.StringPtrInput
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices InstanceEphemeralBlockDeviceArrayInput
	// The instance's host name.
	Hostname            pulumi.StringPtrInput
	InfrastructureClass pulumi.StringPtrInput
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	InstanceProfileArn   pulumi.StringPtrInput
	// The type of instance to start
	InstanceType       pulumi.StringPtrInput
	LastServiceErrorId pulumi.StringPtrInput
	// The ids of the layers the instance will belong to.
	LayerIds pulumi.StringArrayInput
	// Name of operating system that will be installed.
	Os       pulumi.StringPtrInput
	Platform pulumi.StringPtrInput
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns pulumi.StringPtrInput
	// The private IP address assigned to the instance
	PrivateIp pulumi.StringPtrInput
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns pulumi.StringPtrInput
	// The public IP address assigned to the instance, if applicable.
	PublicIp             pulumi.StringPtrInput
	RegisteredBy         pulumi.StringPtrInput
	ReportedAgentVersion pulumi.StringPtrInput
	ReportedOsFamily     pulumi.StringPtrInput
	ReportedOsName       pulumi.StringPtrInput
	ReportedOsVersion    pulumi.StringPtrInput
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices InstanceRootBlockDeviceArrayInput
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType     pulumi.StringPtrInput
	RootDeviceVolumeId pulumi.StringPtrInput
	// The associated security groups.
	SecurityGroupIds         pulumi.StringArrayInput
	SshHostDsaKeyFingerprint pulumi.StringPtrInput
	SshHostRsaKeyFingerprint pulumi.StringPtrInput
	// Name of the SSH keypair that instances will have by default.
	SshKeyName pulumi.StringPtrInput
	// The id of the stack the instance will belong to.
	StackId pulumi.StringInput
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State  pulumi.StringPtrInput
	Status pulumi.StringPtrInput
	// Subnet ID to attach to
	SubnetId pulumi.StringPtrInput
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy pulumi.StringPtrInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}
