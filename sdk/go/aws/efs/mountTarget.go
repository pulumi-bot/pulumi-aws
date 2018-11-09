// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic File System (EFS) mount target.
type MountTarget struct {
	s *pulumi.ResourceState
}

// NewMountTarget registers a new resource with the given unique name, arguments, and options.
func NewMountTarget(ctx *pulumi.Context,
	name string, args *MountTargetArgs, opts ...pulumi.ResourceOpt) (*MountTarget, error) {
	if args == nil || args.FileSystemId == nil {
		return nil, errors.New("missing required argument 'FileSystemId'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["fileSystemId"] = nil
		inputs["ipAddress"] = nil
		inputs["securityGroups"] = nil
		inputs["subnetId"] = nil
	} else {
		inputs["fileSystemId"] = args.FileSystemId
		inputs["ipAddress"] = args.IpAddress
		inputs["securityGroups"] = args.SecurityGroups
		inputs["subnetId"] = args.SubnetId
	}
	inputs["dnsName"] = nil
	inputs["fileSystemArn"] = nil
	inputs["networkInterfaceId"] = nil
	s, err := ctx.RegisterResource("aws:efs/mountTarget:MountTarget", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MountTarget{s: s}, nil
}

// GetMountTarget gets an existing MountTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMountTarget(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MountTargetState, opts ...pulumi.ResourceOpt) (*MountTarget, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dnsName"] = state.DnsName
		inputs["fileSystemArn"] = state.FileSystemArn
		inputs["fileSystemId"] = state.FileSystemId
		inputs["ipAddress"] = state.IpAddress
		inputs["networkInterfaceId"] = state.NetworkInterfaceId
		inputs["securityGroups"] = state.SecurityGroups
		inputs["subnetId"] = state.SubnetId
	}
	s, err := ctx.ReadResource("aws:efs/mountTarget:MountTarget", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MountTarget{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MountTarget) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MountTarget) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
func (r *MountTarget) DnsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsName"])
}

// Amazon Resource Name of the file system.
func (r *MountTarget) FileSystemArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fileSystemArn"])
}

// The ID of the file system for which the mount target is intended.
func (r *MountTarget) FileSystemId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fileSystemId"])
}

// The address (within the address range of the specified subnet) at
// which the file system may be mounted via the mount target.
func (r *MountTarget) IpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipAddress"])
}

// The ID of the network interface that Amazon EFS created when it created the mount target.
func (r *MountTarget) NetworkInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkInterfaceId"])
}

// A list of up to 5 VPC security group IDs (that must
// be for the same VPC as subnet specified) in effect for the mount target.
func (r *MountTarget) SecurityGroups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroups"])
}

// The ID of the subnet to add the mount target in.
func (r *MountTarget) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// Input properties used for looking up and filtering MountTarget resources.
type MountTargetState struct {
	// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	DnsName interface{}
	// Amazon Resource Name of the file system.
	FileSystemArn interface{}
	// The ID of the file system for which the mount target is intended.
	FileSystemId interface{}
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress interface{}
	// The ID of the network interface that Amazon EFS created when it created the mount target.
	NetworkInterfaceId interface{}
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups interface{}
	// The ID of the subnet to add the mount target in.
	SubnetId interface{}
}

// The set of arguments for constructing a MountTarget resource.
type MountTargetArgs struct {
	// The ID of the file system for which the mount target is intended.
	FileSystemId interface{}
	// The address (within the address range of the specified subnet) at
	// which the file system may be mounted via the mount target.
	IpAddress interface{}
	// A list of up to 5 VPC security group IDs (that must
	// be for the same VPC as subnet specified) in effect for the mount target.
	SecurityGroups interface{}
	// The ID of the subnet to add the mount target in.
	SubnetId interface{}
}
