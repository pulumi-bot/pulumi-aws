// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS DataSync EFS Location.
//
// > **NOTE:** The EFS File System must have a mounted EFS Mount Target before creating this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/datasync"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datasync.NewEfsLocation(ctx, "example", &datasync.EfsLocationArgs{
// 			EfsFileSystemArn: pulumi.Any(aws_efs_mount_target.Example.File_system_arn),
// 			Ec2Config: &datasync.EfsLocationEc2ConfigArgs{
// 				SecurityGroupArns: pulumi.StringArray{
// 					pulumi.Any(aws_security_group.Example.Arn),
// 				},
// 				SubnetArn: pulumi.Any(aws_subnet.Example.Arn),
// 			},
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
// `aws_datasync_location_efs` can be imported by using the DataSync Task Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:datasync/efsLocation:EfsLocation example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
// ```
type EfsLocation struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2ConfigOutput `pulumi:"ec2Config"`
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn pulumi.StringOutput `pulumi:"efsFileSystemArn"`
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory pulumi.StringPtrOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	Uri  pulumi.StringOutput    `pulumi:"uri"`
}

// NewEfsLocation registers a new resource with the given unique name, arguments, and options.
func NewEfsLocation(ctx *pulumi.Context,
	name string, args *EfsLocationArgs, opts ...pulumi.ResourceOption) (*EfsLocation, error) {
	if args == nil || args.Ec2Config == nil {
		return nil, errors.New("missing required argument 'Ec2Config'")
	}
	if args == nil || args.EfsFileSystemArn == nil {
		return nil, errors.New("missing required argument 'EfsFileSystemArn'")
	}
	if args == nil {
		args = &EfsLocationArgs{}
	}
	var resource EfsLocation
	err := ctx.RegisterResource("aws:datasync/efsLocation:EfsLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEfsLocation gets an existing EfsLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEfsLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EfsLocationState, opts ...pulumi.ResourceOption) (*EfsLocation, error) {
	var resource EfsLocation
	err := ctx.ReadResource("aws:datasync/efsLocation:EfsLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EfsLocation resources.
type efsLocationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config *EfsLocationEc2Config `pulumi:"ec2Config"`
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn *string `pulumi:"efsFileSystemArn"`
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]string `pulumi:"tags"`
	Uri  *string           `pulumi:"uri"`
}

type EfsLocationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2ConfigPtrInput
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn pulumi.StringPtrInput
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapInput
	Uri  pulumi.StringPtrInput
}

func (EfsLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*efsLocationState)(nil)).Elem()
}

type efsLocationArgs struct {
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2Config `pulumi:"ec2Config"`
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn string `pulumi:"efsFileSystemArn"`
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EfsLocation resource.
type EfsLocationArgs struct {
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2ConfigInput
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn pulumi.StringInput
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapInput
}

func (EfsLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*efsLocationArgs)(nil)).Elem()
}

type EfsLocationInput interface {
	pulumi.Input

	ToEfsLocationOutput() EfsLocationOutput
	ToEfsLocationOutputWithContext(ctx context.Context) EfsLocationOutput
}

func (EfsLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*EfsLocation)(nil)).Elem()
}

func (i EfsLocation) ToEfsLocationOutput() EfsLocationOutput {
	return i.ToEfsLocationOutputWithContext(context.Background())
}

func (i EfsLocation) ToEfsLocationOutputWithContext(ctx context.Context) EfsLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationOutput)
}

type EfsLocationOutput struct {
	*pulumi.OutputState
}

func (EfsLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EfsLocationOutput)(nil)).Elem()
}

func (o EfsLocationOutput) ToEfsLocationOutput() EfsLocationOutput {
	return o
}

func (o EfsLocationOutput) ToEfsLocationOutputWithContext(ctx context.Context) EfsLocationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EfsLocationOutput{})
}
