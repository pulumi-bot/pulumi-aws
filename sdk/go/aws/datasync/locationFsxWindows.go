// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS DataSync FSx Windows Location.
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
// 		_, err := datasync.NewLocationFsxWindows(ctx, "example", &datasync.LocationFsxWindowsArgs{
// 			FsxFilesystemArn: pulumi.Any(aws_fsx_windows_file_system.Example.Arn),
// 			User:             pulumi.String("SomeUser"),
// 			Password:         pulumi.String("SuperSecretPassw0rd"),
// 			SecurityGroupArns: pulumi.StringArray{
// 				pulumi.Any(aws_security_group.Example.Arn),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocationFsxWindows struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The time that the FSx for Windows location was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn pulumi.StringOutput `pulumi:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password pulumi.StringOutput `pulumi:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns pulumi.StringArrayOutput `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The URL of the FSx for Windows location that was described.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewLocationFsxWindows registers a new resource with the given unique name, arguments, and options.
func NewLocationFsxWindows(ctx *pulumi.Context,
	name string, args *LocationFsxWindowsArgs, opts ...pulumi.ResourceOption) (*LocationFsxWindows, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.FsxFilesystemArn == nil {
		return nil, errors.New("invalid value for required argument 'FsxFilesystemArn'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.SecurityGroupArns == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupArns'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource LocationFsxWindows
	err := ctx.RegisterResource("aws:datasync/locationFsxWindows:LocationFsxWindows", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationFsxWindows gets an existing LocationFsxWindows resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationFsxWindows(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationFsxWindowsState, opts ...pulumi.ResourceOption) (*LocationFsxWindows, error) {
	var resource LocationFsxWindows
	err := ctx.ReadResource("aws:datasync/locationFsxWindows:LocationFsxWindows", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationFsxWindows resources.
type locationFsxWindowsState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// The time that the FSx for Windows location was created.
	CreationTime *string `pulumi:"creationTime"`
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain *string `pulumi:"domain"`
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn *string `pulumi:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password *string `pulumi:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]string `pulumi:"tags"`
	// The URL of the FSx for Windows location that was described.
	Uri *string `pulumi:"uri"`
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User *string `pulumi:"user"`
}

type LocationFsxWindowsState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// The time that the FSx for Windows location was created.
	CreationTime pulumi.StringPtrInput
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn pulumi.StringPtrInput
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password pulumi.StringPtrInput
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns pulumi.StringArrayInput
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapInput
	// The URL of the FSx for Windows location that was described.
	Uri pulumi.StringPtrInput
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User pulumi.StringPtrInput
}

func (LocationFsxWindowsState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFsxWindowsState)(nil)).Elem()
}

type locationFsxWindowsArgs struct {
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain *string `pulumi:"domain"`
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn string `pulumi:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password string `pulumi:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]string `pulumi:"tags"`
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a LocationFsxWindows resource.
type LocationFsxWindowsArgs struct {
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn pulumi.StringInput
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password pulumi.StringInput
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns pulumi.StringArrayInput
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapInput
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User pulumi.StringInput
}

func (LocationFsxWindowsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFsxWindowsArgs)(nil)).Elem()
}
