// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic File System (EFS) access point.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.NewAccessPoint(ctx, "test", &efs.AccessPointArgs{
// 			FileSystemId: pulumi.Any(aws_efs_file_system.Foo.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AccessPoint struct {
	pulumi.CustomResourceState

	// Amazon Resource Name of the access point.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Amazon Resource Name of the file system.
	FileSystemArn pulumi.StringOutput `pulumi:"fileSystemArn"`
	// The ID of the file system for which the access point is intended.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	OwnerId      pulumi.StringOutput `pulumi:"ownerId"`
	// The operating system user and group applied to all file system requests made using the access point. See Posix User below.
	PosixUser AccessPointPosixUserPtrOutput `pulumi:"posixUser"`
	// Specifies the directory on the Amazon EFS file system that the access point provides access to. See Root Directory below.
	RootDirectory AccessPointRootDirectoryOutput `pulumi:"rootDirectory"`
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewAccessPoint(ctx *pulumi.Context,
	name string, args *AccessPointArgs, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	if args == nil || args.FileSystemId == nil {
		return nil, errors.New("missing required argument 'FileSystemId'")
	}
	if args == nil {
		args = &AccessPointArgs{}
	}
	var resource AccessPoint
	err := ctx.RegisterResource("aws:efs/accessPoint:AccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPoint gets an existing AccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPointState, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	var resource AccessPoint
	err := ctx.ReadResource("aws:efs/accessPoint:AccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPoint resources.
type accessPointState struct {
	// Amazon Resource Name of the access point.
	Arn *string `pulumi:"arn"`
	// Amazon Resource Name of the file system.
	FileSystemArn *string `pulumi:"fileSystemArn"`
	// The ID of the file system for which the access point is intended.
	FileSystemId *string `pulumi:"fileSystemId"`
	OwnerId      *string `pulumi:"ownerId"`
	// The operating system user and group applied to all file system requests made using the access point. See Posix User below.
	PosixUser *AccessPointPosixUser `pulumi:"posixUser"`
	// Specifies the directory on the Amazon EFS file system that the access point provides access to. See Root Directory below.
	RootDirectory *AccessPointRootDirectory `pulumi:"rootDirectory"`
	// Key-value mapping of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type AccessPointState struct {
	// Amazon Resource Name of the access point.
	Arn pulumi.StringPtrInput
	// Amazon Resource Name of the file system.
	FileSystemArn pulumi.StringPtrInput
	// The ID of the file system for which the access point is intended.
	FileSystemId pulumi.StringPtrInput
	OwnerId      pulumi.StringPtrInput
	// The operating system user and group applied to all file system requests made using the access point. See Posix User below.
	PosixUser AccessPointPosixUserPtrInput
	// Specifies the directory on the Amazon EFS file system that the access point provides access to. See Root Directory below.
	RootDirectory AccessPointRootDirectoryPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapInput
}

func (AccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointState)(nil)).Elem()
}

type accessPointArgs struct {
	// The ID of the file system for which the access point is intended.
	FileSystemId string `pulumi:"fileSystemId"`
	// The operating system user and group applied to all file system requests made using the access point. See Posix User below.
	PosixUser *AccessPointPosixUser `pulumi:"posixUser"`
	// Specifies the directory on the Amazon EFS file system that the access point provides access to. See Root Directory below.
	RootDirectory *AccessPointRootDirectory `pulumi:"rootDirectory"`
	// Key-value mapping of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AccessPoint resource.
type AccessPointArgs struct {
	// The ID of the file system for which the access point is intended.
	FileSystemId pulumi.StringInput
	// The operating system user and group applied to all file system requests made using the access point. See Posix User below.
	PosixUser AccessPointPosixUserPtrInput
	// Specifies the directory on the Amazon EFS file system that the access point provides access to. See Root Directory below.
	RootDirectory AccessPointRootDirectoryPtrInput
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapInput
}

func (AccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointArgs)(nil)).Elem()
}

type AccessPointInput interface {
	pulumi.Input

	ToAccessPointOutput() AccessPointOutput
	ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput
}

func (AccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPoint)(nil)).Elem()
}

func (i AccessPoint) ToAccessPointOutput() AccessPointOutput {
	return i.ToAccessPointOutputWithContext(context.Background())
}

func (i AccessPoint) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointOutput)
}

type AccessPointOutput struct {
	*pulumi.OutputState
}

func (AccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPointOutput)(nil)).Elem()
}

func (o AccessPointOutput) ToAccessPointOutput() AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessPointOutput{})
}
