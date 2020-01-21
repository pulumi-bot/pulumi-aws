// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package table

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/dynamodb/TableAttribute"
	"https:/github.com/pulumi/pulumi-aws/dynamodb/TableGlobalSecondaryIndex"
	"https:/github.com/pulumi/pulumi-aws/dynamodb/TableLocalSecondaryIndex"
	"https:/github.com/pulumi/pulumi-aws/dynamodb/TablePointInTimeRecovery"
	"https:/github.com/pulumi/pulumi-aws/dynamodb/TableServerSideEncryption"
	"https:/github.com/pulumi/pulumi-aws/dynamodb/TableTtl"
)

// Provides a DynamoDB table resource
// 
// > **Note:** It is recommended to use `lifecycle` [`ignoreChanges`](https://www.terraform.io/docs/configuration/resources.html#ignore_changes) for `readCapacity` and/or `writeCapacity` if there's [autoscaling policy](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html) attached to the table.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dynamodb_table.html.markdown.
type Table struct {
	pulumi.CustomResourceState

	// The arn of the table
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes dynamodbTableAttribute.TableAttributeArrayOutput `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringPtrOutput `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes dynamodbTableGlobalSecondaryIndex.TableGlobalSecondaryIndexArrayOutput `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringOutput `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes dynamodbTableLocalSecondaryIndex.TableLocalSecondaryIndexArrayOutput `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name pulumi.StringOutput `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery dynamodbTablePointInTimeRecovery.TablePointInTimeRecoveryOutput `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey pulumi.StringPtrOutput `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntPtrOutput `pulumi:"readCapacity"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption dynamodbTableServerSideEncryption.TableServerSideEncryptionOutput `pulumi:"serverSideEncryption"`
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn pulumi.StringOutput `pulumi:"streamArn"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolPtrOutput `pulumi:"streamEnabled"`
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel pulumi.StringOutput `pulumi:"streamLabel"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringOutput `pulumi:"streamViewType"`
	// A map of tags to populate on the created table.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl dynamodbTableTtl.TableTtlPtrOutput `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntPtrOutput `pulumi:"writeCapacity"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil || args.Attributes == nil {
		return nil, errors.New("missing required argument 'Attributes'")
	}
	if args == nil || args.HashKey == nil {
		return nil, errors.New("missing required argument 'HashKey'")
	}
	if args == nil {
		args = &TableArgs{}
	}
	var resource Table
	err := ctx.RegisterResource("aws:dynamodb/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("aws:dynamodb/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// The arn of the table
	Arn *string `pulumi:"arn"`
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes []dynamodbTableAttribute.TableAttribute `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode *string `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes []dynamodbTableGlobalSecondaryIndex.TableGlobalSecondaryIndex `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey *string `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes []dynamodbTableLocalSecondaryIndex.TableLocalSecondaryIndex `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name *string `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery *dynamodbTablePointInTimeRecovery.TablePointInTimeRecovery `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey *string `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity *int `pulumi:"readCapacity"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption *dynamodbTableServerSideEncryption.TableServerSideEncryption `pulumi:"serverSideEncryption"`
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn *string `pulumi:"streamArn"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled *bool `pulumi:"streamEnabled"`
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel *string `pulumi:"streamLabel"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType *string `pulumi:"streamViewType"`
	// A map of tags to populate on the created table.
	Tags map[string]interface{} `pulumi:"tags"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl *dynamodbTableTtl.TableTtl `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity *int `pulumi:"writeCapacity"`
}

type TableState struct {
	// The arn of the table
	Arn pulumi.StringPtrInput
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes dynamodbTableAttribute.TableAttributeArrayInput
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringPtrInput
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes dynamodbTableGlobalSecondaryIndex.TableGlobalSecondaryIndexArrayInput
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringPtrInput
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes dynamodbTableLocalSecondaryIndex.TableLocalSecondaryIndexArrayInput
	// The name of the index
	Name pulumi.StringPtrInput
	// Point-in-time recovery options.
	PointInTimeRecovery dynamodbTablePointInTimeRecovery.TablePointInTimeRecoveryPtrInput
	// The name of the range key; must be defined
	RangeKey pulumi.StringPtrInput
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntPtrInput
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption dynamodbTableServerSideEncryption.TableServerSideEncryptionPtrInput
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn pulumi.StringPtrInput
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolPtrInput
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel pulumi.StringPtrInput
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringPtrInput
	// A map of tags to populate on the created table.
	Tags pulumi.MapInput
	// Defines ttl, has two properties, and can only be specified once:
	Ttl dynamodbTableTtl.TableTtlPtrInput
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes []dynamodbTableAttribute.TableAttribute `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode *string `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes []dynamodbTableGlobalSecondaryIndex.TableGlobalSecondaryIndex `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey string `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes []dynamodbTableLocalSecondaryIndex.TableLocalSecondaryIndex `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name *string `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery *dynamodbTablePointInTimeRecovery.TablePointInTimeRecovery `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey *string `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity *int `pulumi:"readCapacity"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption *dynamodbTableServerSideEncryption.TableServerSideEncryption `pulumi:"serverSideEncryption"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled *bool `pulumi:"streamEnabled"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType *string `pulumi:"streamViewType"`
	// A map of tags to populate on the created table.
	Tags map[string]interface{} `pulumi:"tags"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl *dynamodbTableTtl.TableTtl `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity *int `pulumi:"writeCapacity"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes dynamodbTableAttribute.TableAttributeArrayInput
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringPtrInput
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes dynamodbTableGlobalSecondaryIndex.TableGlobalSecondaryIndexArrayInput
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringInput
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes dynamodbTableLocalSecondaryIndex.TableLocalSecondaryIndexArrayInput
	// The name of the index
	Name pulumi.StringPtrInput
	// Point-in-time recovery options.
	PointInTimeRecovery dynamodbTablePointInTimeRecovery.TablePointInTimeRecoveryPtrInput
	// The name of the range key; must be defined
	RangeKey pulumi.StringPtrInput
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntPtrInput
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption dynamodbTableServerSideEncryption.TableServerSideEncryptionPtrInput
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolPtrInput
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringPtrInput
	// A map of tags to populate on the created table.
	Tags pulumi.MapInput
	// Defines ttl, has two properties, and can only be specified once:
	Ttl dynamodbTableTtl.TableTtlPtrInput
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

