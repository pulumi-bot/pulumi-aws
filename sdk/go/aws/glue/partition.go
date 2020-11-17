// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Glue Partition Resource.
//
// ## Import
//
// Glue Partitions can be imported with their catalog ID (usually AWS account ID), database name, table name and partition values e.g.
//
// ```sh
//  $ pulumi import aws:glue/partition:Partition part 123456789012:MyDatabase:MyTable:val1#val2
// ```
type Partition struct {
	pulumi.CustomResourceState

	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// The time at which the partition was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The last time at which the partition was accessed.
	LastAccessedTime pulumi.StringOutput `pulumi:"lastAccessedTime"`
	// The last time at which column statistics were computed for this partition.
	LastAnalyzedTime pulumi.StringOutput `pulumi:"lastAnalyzedTime"`
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The values that define the partition.
	PartitionValues pulumi.StringArrayOutput `pulumi:"partitionValues"`
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor PartitionStorageDescriptorPtrOutput `pulumi:"storageDescriptor"`
	TableName         pulumi.StringOutput                 `pulumi:"tableName"`
}

// NewPartition registers a new resource with the given unique name, arguments, and options.
func NewPartition(ctx *pulumi.Context,
	name string, args *PartitionArgs, opts ...pulumi.ResourceOption) (*Partition, error) {
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.PartitionValues == nil {
		return nil, errors.New("missing required argument 'PartitionValues'")
	}
	if args == nil || args.TableName == nil {
		return nil, errors.New("missing required argument 'TableName'")
	}
	if args == nil {
		args = &PartitionArgs{}
	}
	var resource Partition
	err := ctx.RegisterResource("aws:glue/partition:Partition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartition gets an existing Partition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartitionState, opts ...pulumi.ResourceOption) (*Partition, error) {
	var resource Partition
	err := ctx.ReadResource("aws:glue/partition:Partition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Partition resources.
type partitionState struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId *string `pulumi:"catalogId"`
	// The time at which the partition was created.
	CreationTime *string `pulumi:"creationTime"`
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName *string `pulumi:"databaseName"`
	// The last time at which the partition was accessed.
	LastAccessedTime *string `pulumi:"lastAccessedTime"`
	// The last time at which column statistics were computed for this partition.
	LastAnalyzedTime *string `pulumi:"lastAnalyzedTime"`
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `pulumi:"parameters"`
	// The values that define the partition.
	PartitionValues []string `pulumi:"partitionValues"`
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor *PartitionStorageDescriptor `pulumi:"storageDescriptor"`
	TableName         *string                     `pulumi:"tableName"`
}

type PartitionState struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId pulumi.StringPtrInput
	// The time at which the partition was created.
	CreationTime pulumi.StringPtrInput
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName pulumi.StringPtrInput
	// The last time at which the partition was accessed.
	LastAccessedTime pulumi.StringPtrInput
	// The last time at which column statistics were computed for this partition.
	LastAnalyzedTime pulumi.StringPtrInput
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapInput
	// The values that define the partition.
	PartitionValues pulumi.StringArrayInput
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor PartitionStorageDescriptorPtrInput
	TableName         pulumi.StringPtrInput
}

func (PartitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionState)(nil)).Elem()
}

type partitionArgs struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId *string `pulumi:"catalogId"`
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName string `pulumi:"databaseName"`
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]string `pulumi:"parameters"`
	// The values that define the partition.
	PartitionValues []string `pulumi:"partitionValues"`
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor *PartitionStorageDescriptor `pulumi:"storageDescriptor"`
	TableName         string                      `pulumi:"tableName"`
}

// The set of arguments for constructing a Partition resource.
type PartitionArgs struct {
	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogId pulumi.StringPtrInput
	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName pulumi.StringInput
	// A map of initialization parameters for the SerDe, in key-value form.
	Parameters pulumi.StringMapInput
	// The values that define the partition.
	PartitionValues pulumi.StringArrayInput
	// A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
	StorageDescriptor PartitionStorageDescriptorPtrInput
	TableName         pulumi.StringInput
}

func (PartitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionArgs)(nil)).Elem()
}

type PartitionInput interface {
	pulumi.Input

	ToPartitionOutput() PartitionOutput
	ToPartitionOutputWithContext(ctx context.Context) PartitionOutput
}

func (Partition) ElementType() reflect.Type {
	return reflect.TypeOf((*Partition)(nil)).Elem()
}

func (i Partition) ToPartitionOutput() PartitionOutput {
	return i.ToPartitionOutputWithContext(context.Background())
}

func (i Partition) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionOutput)
}

type PartitionOutput struct {
	*pulumi.OutputState
}

func (PartitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionOutput)(nil)).Elem()
}

func (o PartitionOutput) ToPartitionOutput() PartitionOutput {
	return o
}

func (o PartitionOutput) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PartitionOutput{})
}
