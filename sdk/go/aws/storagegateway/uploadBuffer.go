// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Storage Gateway upload buffer.
//
// > **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this resource does not perform any Storage Gateway actions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewUploadBuffer(ctx, "example", &storagegateway.UploadBufferArgs{
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
type UploadBuffer struct {
	pulumi.CustomResourceState

	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
}

// NewUploadBuffer registers a new resource with the given unique name, arguments, and options.
func NewUploadBuffer(ctx *pulumi.Context,
	name string, args *UploadBufferArgs, opts ...pulumi.ResourceOption) (*UploadBuffer, error) {
	if args == nil || args.DiskId == nil {
		return nil, errors.New("missing required argument 'DiskId'")
	}
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	if args == nil {
		args = &UploadBufferArgs{}
	}
	var resource UploadBuffer
	err := ctx.RegisterResource("aws:storagegateway/uploadBuffer:UploadBuffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUploadBuffer gets an existing UploadBuffer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUploadBuffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UploadBufferState, opts ...pulumi.ResourceOption) (*UploadBuffer, error) {
	var resource UploadBuffer
	err := ctx.ReadResource("aws:storagegateway/uploadBuffer:UploadBuffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UploadBuffer resources.
type uploadBufferState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId *string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
}

type UploadBufferState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
}

func (UploadBufferState) ElementType() reflect.Type {
	return reflect.TypeOf((*uploadBufferState)(nil)).Elem()
}

type uploadBufferArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
}

// The set of arguments for constructing a UploadBuffer resource.
type UploadBufferArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
}

func (UploadBufferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uploadBufferArgs)(nil)).Elem()
}
