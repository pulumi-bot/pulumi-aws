// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic File System (EFS) mount target.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := ec2.NewVpc(ctx, "foo", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		alphaSubnet, err := ec2.NewSubnet(ctx, "alphaSubnet", &ec2.SubnetArgs{
// 			VpcId:            foo.ID(),
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			CidrBlock:        pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = efs.NewMountTarget(ctx, "alphaMountTarget", &efs.MountTargetArgs{
// 			FileSystemId: pulumi.Any(aws_efs_file_system.Foo.Id),
// 			SubnetId:     alphaSubnet.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// The EFS mount targets can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:efs/mountTarget:MountTarget alpha fsmt-52a643fb
// ```
type MountTarget struct {
	pulumi.CustomResourceState

	// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneId pulumi.StringOutput `pulumi:"availabilityZoneId"`
	// The name of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneName pulumi.StringOutput `pulumi:"availabilityZoneName"`
	// The DNS name for the EFS file system.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Amazon Resource Name of the file system.
	FileSystemArn pulumi.StringOutput `pulumi:"fileSystemArn"`
	// The ID of the file system for which the mount target is intended.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	MountTargetDnsName pulumi.StringOutput `pulumi:"mountTargetDnsName"`
	// The ID of the network interface that Amazon EFS created when it created the mount target.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// AWS account ID that owns the resource.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// The ID of the subnet to add the mount target in.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewMountTarget registers a new resource with the given unique name, arguments, and options.
func NewMountTarget(ctx *pulumi.Context,
	name string, args *MountTargetArgs, opts ...pulumi.ResourceOption) (*MountTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource MountTarget
	err := ctx.RegisterResource("aws:efs/mountTarget:MountTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMountTarget gets an existing MountTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMountTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MountTargetState, opts ...pulumi.ResourceOption) (*MountTarget, error) {
	var resource MountTarget
	err := ctx.ReadResource("aws:efs/mountTarget:MountTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MountTarget resources.
type mountTargetState struct {
	// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneId *string `pulumi:"availabilityZoneId"`
	// The name of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneName *string `pulumi:"availabilityZoneName"`
	// The DNS name for the EFS file system.
	DnsName *string `pulumi:"dnsName"`
	// Amazon Resource Name of the file system.
	FileSystemArn *string `pulumi:"fileSystemArn"`
	// The ID of the file system for which the mount target is intended.
	FileSystemId *string `pulumi:"fileSystemId"`
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress *string `pulumi:"ipAddress"`
	// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	MountTargetDnsName *string `pulumi:"mountTargetDnsName"`
	// The ID of the network interface that Amazon EFS created when it created the mount target.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// AWS account ID that owns the resource.
	OwnerId *string `pulumi:"ownerId"`
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The ID of the subnet to add the mount target in.
	SubnetId *string `pulumi:"subnetId"`
}

type MountTargetState struct {
	// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneId pulumi.StringPtrInput
	// The name of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneName pulumi.StringPtrInput
	// The DNS name for the EFS file system.
	DnsName pulumi.StringPtrInput
	// Amazon Resource Name of the file system.
	FileSystemArn pulumi.StringPtrInput
	// The ID of the file system for which the mount target is intended.
	FileSystemId pulumi.StringPtrInput
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress pulumi.StringPtrInput
	// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	MountTargetDnsName pulumi.StringPtrInput
	// The ID of the network interface that Amazon EFS created when it created the mount target.
	NetworkInterfaceId pulumi.StringPtrInput
	// AWS account ID that owns the resource.
	OwnerId pulumi.StringPtrInput
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups pulumi.StringArrayInput
	// The ID of the subnet to add the mount target in.
	SubnetId pulumi.StringPtrInput
}

func (MountTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*mountTargetState)(nil)).Elem()
}

type mountTargetArgs struct {
	// The ID of the file system for which the mount target is intended.
	FileSystemId string `pulumi:"fileSystemId"`
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress *string `pulumi:"ipAddress"`
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups []string `pulumi:"securityGroups"`
	// The ID of the subnet to add the mount target in.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a MountTarget resource.
type MountTargetArgs struct {
	// The ID of the file system for which the mount target is intended.
	FileSystemId pulumi.StringInput
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress pulumi.StringPtrInput
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups pulumi.StringArrayInput
	// The ID of the subnet to add the mount target in.
	SubnetId pulumi.StringInput
}

func (MountTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mountTargetArgs)(nil)).Elem()
}

type MountTargetInput interface {
	pulumi.Input

	ToMountTargetOutput() MountTargetOutput
	ToMountTargetOutputWithContext(ctx context.Context) MountTargetOutput
}

func (MountTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTarget)(nil)).Elem()
}

func (i MountTarget) ToMountTargetOutput() MountTargetOutput {
	return i.ToMountTargetOutputWithContext(context.Background())
}

func (i MountTarget) ToMountTargetOutputWithContext(ctx context.Context) MountTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountTargetOutput)
}

type MountTargetOutput struct {
	*pulumi.OutputState
}

func (MountTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MountTargetOutput)(nil)).Elem()
}

func (o MountTargetOutput) ToMountTargetOutput() MountTargetOutput {
	return o
}

func (o MountTargetOutput) ToMountTargetOutputWithContext(ctx context.Context) MountTargetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MountTargetOutput{})
}
