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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ec2Config == nil {
		return nil, errors.New("invalid value for required argument 'Ec2Config'")
	}
	if args.EfsFileSystemArn == nil {
		return nil, errors.New("invalid value for required argument 'EfsFileSystemArn'")
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

func (*EfsLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*EfsLocation)(nil))
}

func (i *EfsLocation) ToEfsLocationOutput() EfsLocationOutput {
	return i.ToEfsLocationOutputWithContext(context.Background())
}

func (i *EfsLocation) ToEfsLocationOutputWithContext(ctx context.Context) EfsLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationOutput)
}

func (i *EfsLocation) ToEfsLocationPtrOutput() EfsLocationPtrOutput {
	return i.ToEfsLocationPtrOutputWithContext(context.Background())
}

func (i *EfsLocation) ToEfsLocationPtrOutputWithContext(ctx context.Context) EfsLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationPtrOutput)
}

type EfsLocationPtrInput interface {
	pulumi.Input

	ToEfsLocationPtrOutput() EfsLocationPtrOutput
	ToEfsLocationPtrOutputWithContext(ctx context.Context) EfsLocationPtrOutput
}

type efsLocationPtrType EfsLocationArgs

func (*efsLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EfsLocation)(nil))
}

func (i *efsLocationPtrType) ToEfsLocationPtrOutput() EfsLocationPtrOutput {
	return i.ToEfsLocationPtrOutputWithContext(context.Background())
}

func (i *efsLocationPtrType) ToEfsLocationPtrOutputWithContext(ctx context.Context) EfsLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationPtrOutput)
}

// EfsLocationArrayInput is an input type that accepts EfsLocationArray and EfsLocationArrayOutput values.
// You can construct a concrete instance of `EfsLocationArrayInput` via:
//
//          EfsLocationArray{ EfsLocationArgs{...} }
type EfsLocationArrayInput interface {
	pulumi.Input

	ToEfsLocationArrayOutput() EfsLocationArrayOutput
	ToEfsLocationArrayOutputWithContext(context.Context) EfsLocationArrayOutput
}

type EfsLocationArray []EfsLocationInput

func (EfsLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EfsLocation)(nil))
}

func (i EfsLocationArray) ToEfsLocationArrayOutput() EfsLocationArrayOutput {
	return i.ToEfsLocationArrayOutputWithContext(context.Background())
}

func (i EfsLocationArray) ToEfsLocationArrayOutputWithContext(ctx context.Context) EfsLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationArrayOutput)
}

// EfsLocationMapInput is an input type that accepts EfsLocationMap and EfsLocationMapOutput values.
// You can construct a concrete instance of `EfsLocationMapInput` via:
//
//          EfsLocationMap{ "key": EfsLocationArgs{...} }
type EfsLocationMapInput interface {
	pulumi.Input

	ToEfsLocationMapOutput() EfsLocationMapOutput
	ToEfsLocationMapOutputWithContext(context.Context) EfsLocationMapOutput
}

type EfsLocationMap map[string]EfsLocationInput

func (EfsLocationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EfsLocation)(nil))
}

func (i EfsLocationMap) ToEfsLocationMapOutput() EfsLocationMapOutput {
	return i.ToEfsLocationMapOutputWithContext(context.Background())
}

func (i EfsLocationMap) ToEfsLocationMapOutputWithContext(ctx context.Context) EfsLocationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationMapOutput)
}

type EfsLocationOutput struct {
	*pulumi.OutputState
}

func (EfsLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EfsLocation)(nil))
}

func (o EfsLocationOutput) ToEfsLocationOutput() EfsLocationOutput {
	return o
}

func (o EfsLocationOutput) ToEfsLocationOutputWithContext(ctx context.Context) EfsLocationOutput {
	return o
}

func (o EfsLocationOutput) ToEfsLocationPtrOutput() EfsLocationPtrOutput {
	return o.ToEfsLocationPtrOutputWithContext(context.Background())
}

func (o EfsLocationOutput) ToEfsLocationPtrOutputWithContext(ctx context.Context) EfsLocationPtrOutput {
	return o.ApplyT(func(v EfsLocation) *EfsLocation {
		return &v
	}).(EfsLocationPtrOutput)
}

type EfsLocationPtrOutput struct {
	*pulumi.OutputState
}

func (EfsLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EfsLocation)(nil))
}

func (o EfsLocationPtrOutput) ToEfsLocationPtrOutput() EfsLocationPtrOutput {
	return o
}

func (o EfsLocationPtrOutput) ToEfsLocationPtrOutputWithContext(ctx context.Context) EfsLocationPtrOutput {
	return o
}

type EfsLocationArrayOutput struct{ *pulumi.OutputState }

func (EfsLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EfsLocation)(nil))
}

func (o EfsLocationArrayOutput) ToEfsLocationArrayOutput() EfsLocationArrayOutput {
	return o
}

func (o EfsLocationArrayOutput) ToEfsLocationArrayOutputWithContext(ctx context.Context) EfsLocationArrayOutput {
	return o
}

func (o EfsLocationArrayOutput) Index(i pulumi.IntInput) EfsLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EfsLocation {
		return vs[0].([]EfsLocation)[vs[1].(int)]
	}).(EfsLocationOutput)
}

type EfsLocationMapOutput struct{ *pulumi.OutputState }

func (EfsLocationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EfsLocation)(nil))
}

func (o EfsLocationMapOutput) ToEfsLocationMapOutput() EfsLocationMapOutput {
	return o
}

func (o EfsLocationMapOutput) ToEfsLocationMapOutputWithContext(ctx context.Context) EfsLocationMapOutput {
	return o
}

func (o EfsLocationMapOutput) MapIndex(k pulumi.StringInput) EfsLocationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EfsLocation {
		return vs[0].(map[string]EfsLocation)[vs[1].(string)]
	}).(EfsLocationOutput)
}

func init() {
	pulumi.RegisterOutputType(EfsLocationOutput{})
	pulumi.RegisterOutputType(EfsLocationPtrOutput{})
	pulumi.RegisterOutputType(EfsLocationArrayOutput{})
	pulumi.RegisterOutputType(EfsLocationMapOutput{})
}
