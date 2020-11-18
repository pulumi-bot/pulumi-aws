// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Storage Gateway working storage.
//
// > **NOTE:** The Storage Gateway API provides no method to remove a working storage disk. Destroying this resource does not perform any Storage Gateway actions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewWorkingStorage(ctx, "example", &storagegateway.WorkingStorageArgs{
// 			DiskId:     pulumi.Any(data.Aws_storagegateway_local_disk.Example.Id),
// 			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type WorkingStorage struct {
	pulumi.CustomResourceState

	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
}

// NewWorkingStorage registers a new resource with the given unique name, arguments, and options.
func NewWorkingStorage(ctx *pulumi.Context,
	name string, args *WorkingStorageArgs, opts ...pulumi.ResourceOption) (*WorkingStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	if args.GatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'GatewayArn'")
	}
	var resource WorkingStorage
	err := ctx.RegisterResource("aws:storagegateway/workingStorage:WorkingStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkingStorage gets an existing WorkingStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkingStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkingStorageState, opts ...pulumi.ResourceOption) (*WorkingStorage, error) {
	var resource WorkingStorage
	err := ctx.ReadResource("aws:storagegateway/workingStorage:WorkingStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkingStorage resources.
type workingStorageState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId *string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
}

type WorkingStorageState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
}

func (WorkingStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*workingStorageState)(nil)).Elem()
}

type workingStorageArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
}

// The set of arguments for constructing a WorkingStorage resource.
type WorkingStorageArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
}

func (WorkingStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workingStorageArgs)(nil)).Elem()
}

type WorkingStorageInput interface {
	pulumi.Input

	ToWorkingStorageOutput() WorkingStorageOutput
	ToWorkingStorageOutputWithContext(ctx context.Context) WorkingStorageOutput
}

func (WorkingStorage) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkingStorage)(nil)).Elem()
}

func (i WorkingStorage) ToWorkingStorageOutput() WorkingStorageOutput {
	return i.ToWorkingStorageOutputWithContext(context.Background())
}

func (i WorkingStorage) ToWorkingStorageOutputWithContext(ctx context.Context) WorkingStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkingStorageOutput)
}

type WorkingStorageOutput struct {
	*pulumi.OutputState
}

func (WorkingStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkingStorageOutput)(nil)).Elem()
}

func (o WorkingStorageOutput) ToWorkingStorageOutput() WorkingStorageOutput {
	return o
}

func (o WorkingStorageOutput) ToWorkingStorageOutputWithContext(ctx context.Context) WorkingStorageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkingStorageOutput{})
}
