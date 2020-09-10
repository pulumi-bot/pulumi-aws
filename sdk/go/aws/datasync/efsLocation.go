// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EfsLocation struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput        `pulumi:"arn"`
	Ec2Config        EfsLocationEc2ConfigOutput `pulumi:"ec2Config"`
	EfsFileSystemArn pulumi.StringOutput        `pulumi:"efsFileSystemArn"`
	Subdirectory     pulumi.StringPtrOutput     `pulumi:"subdirectory"`
	Tags             pulumi.StringMapOutput     `pulumi:"tags"`
	Uri              pulumi.StringOutput        `pulumi:"uri"`
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
	Arn              *string               `pulumi:"arn"`
	Ec2Config        *EfsLocationEc2Config `pulumi:"ec2Config"`
	EfsFileSystemArn *string               `pulumi:"efsFileSystemArn"`
	Subdirectory     *string               `pulumi:"subdirectory"`
	Tags             map[string]string     `pulumi:"tags"`
	Uri              *string               `pulumi:"uri"`
}

type EfsLocationState struct {
	Arn              pulumi.StringPtrInput
	Ec2Config        EfsLocationEc2ConfigPtrInput
	EfsFileSystemArn pulumi.StringPtrInput
	Subdirectory     pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
	Uri              pulumi.StringPtrInput
}

func (EfsLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*efsLocationState)(nil)).Elem()
}

type efsLocationArgs struct {
	Ec2Config        EfsLocationEc2Config `pulumi:"ec2Config"`
	EfsFileSystemArn string               `pulumi:"efsFileSystemArn"`
	Subdirectory     *string              `pulumi:"subdirectory"`
	Tags             map[string]string    `pulumi:"tags"`
}

// The set of arguments for constructing a EfsLocation resource.
type EfsLocationArgs struct {
	Ec2Config        EfsLocationEc2ConfigInput
	EfsFileSystemArn pulumi.StringInput
	Subdirectory     pulumi.StringPtrInput
	Tags             pulumi.StringMapInput
}

func (EfsLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*efsLocationArgs)(nil)).Elem()
}
