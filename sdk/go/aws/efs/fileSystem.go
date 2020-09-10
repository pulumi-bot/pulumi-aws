// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type FileSystem struct {
	pulumi.CustomResourceState

	Arn                          pulumi.StringOutput                `pulumi:"arn"`
	CreationToken                pulumi.StringOutput                `pulumi:"creationToken"`
	DnsName                      pulumi.StringOutput                `pulumi:"dnsName"`
	Encrypted                    pulumi.BoolOutput                  `pulumi:"encrypted"`
	KmsKeyId                     pulumi.StringOutput                `pulumi:"kmsKeyId"`
	LifecyclePolicy              FileSystemLifecyclePolicyPtrOutput `pulumi:"lifecyclePolicy"`
	PerformanceMode              pulumi.StringOutput                `pulumi:"performanceMode"`
	ProvisionedThroughputInMibps pulumi.Float64PtrOutput            `pulumi:"provisionedThroughputInMibps"`
	Tags                         pulumi.StringMapOutput             `pulumi:"tags"`
	ThroughputMode               pulumi.StringPtrOutput             `pulumi:"throughputMode"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		args = &FileSystemArgs{}
	}
	var resource FileSystem
	err := ctx.RegisterResource("aws:efs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("aws:efs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	Arn                          *string                    `pulumi:"arn"`
	CreationToken                *string                    `pulumi:"creationToken"`
	DnsName                      *string                    `pulumi:"dnsName"`
	Encrypted                    *bool                      `pulumi:"encrypted"`
	KmsKeyId                     *string                    `pulumi:"kmsKeyId"`
	LifecyclePolicy              *FileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	PerformanceMode              *string                    `pulumi:"performanceMode"`
	ProvisionedThroughputInMibps *float64                   `pulumi:"provisionedThroughputInMibps"`
	Tags                         map[string]string          `pulumi:"tags"`
	ThroughputMode               *string                    `pulumi:"throughputMode"`
}

type FileSystemState struct {
	Arn                          pulumi.StringPtrInput
	CreationToken                pulumi.StringPtrInput
	DnsName                      pulumi.StringPtrInput
	Encrypted                    pulumi.BoolPtrInput
	KmsKeyId                     pulumi.StringPtrInput
	LifecyclePolicy              FileSystemLifecyclePolicyPtrInput
	PerformanceMode              pulumi.StringPtrInput
	ProvisionedThroughputInMibps pulumi.Float64PtrInput
	Tags                         pulumi.StringMapInput
	ThroughputMode               pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	CreationToken                *string                    `pulumi:"creationToken"`
	Encrypted                    *bool                      `pulumi:"encrypted"`
	KmsKeyId                     *string                    `pulumi:"kmsKeyId"`
	LifecyclePolicy              *FileSystemLifecyclePolicy `pulumi:"lifecyclePolicy"`
	PerformanceMode              *string                    `pulumi:"performanceMode"`
	ProvisionedThroughputInMibps *float64                   `pulumi:"provisionedThroughputInMibps"`
	Tags                         map[string]string          `pulumi:"tags"`
	ThroughputMode               *string                    `pulumi:"throughputMode"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	CreationToken                pulumi.StringPtrInput
	Encrypted                    pulumi.BoolPtrInput
	KmsKeyId                     pulumi.StringPtrInput
	LifecyclePolicy              FileSystemLifecyclePolicyPtrInput
	PerformanceMode              pulumi.StringPtrInput
	ProvisionedThroughputInMibps pulumi.Float64PtrInput
	Tags                         pulumi.StringMapInput
	ThroughputMode               pulumi.StringPtrInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}
