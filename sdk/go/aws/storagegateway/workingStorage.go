// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type WorkingStorage struct {
	s *pulumi.ResourceState
}

// NewWorkingStorage registers a new resource with the given unique name, arguments, and options.
func NewWorkingStorage(ctx *pulumi.Context,
	name string, args *WorkingStorageArgs, opts ...pulumi.ResourceOpt) (*WorkingStorage, error) {
	if args == nil || args.DiskId == nil {
		return nil, errors.New("missing required argument 'DiskId'")
	}
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["diskId"] = nil
		inputs["gatewayArn"] = nil
	} else {
		inputs["diskId"] = args.DiskId
		inputs["gatewayArn"] = args.GatewayArn
	}
	s, err := ctx.RegisterResource("aws:storagegateway/workingStorage:WorkingStorage", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WorkingStorage{s: s}, nil
}

// GetWorkingStorage gets an existing WorkingStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkingStorage(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WorkingStorageState, opts ...pulumi.ResourceOpt) (*WorkingStorage, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["diskId"] = state.DiskId
		inputs["gatewayArn"] = state.GatewayArn
	}
	s, err := ctx.ReadResource("aws:storagegateway/workingStorage:WorkingStorage", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WorkingStorage{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *WorkingStorage) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *WorkingStorage) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
func (r *WorkingStorage) DiskId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["diskId"])
}

// The Amazon Resource Name (ARN) of the gateway.
func (r *WorkingStorage) GatewayArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gatewayArn"])
}

// Input properties used for looking up and filtering WorkingStorage resources.
type WorkingStorageState struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId interface{}
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn interface{}
}

// The set of arguments for constructing a WorkingStorage resource.
type WorkingStorageArgs struct {
	// Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
	DiskId interface{}
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn interface{}
}
