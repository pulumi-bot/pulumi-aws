// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DynamoDB table item resource
//
// > **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
//   You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/dynamodb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
// 			ReadCapacity:  pulumi.Int(10),
// 			WriteCapacity: pulumi.Int(10),
// 			HashKey:       pulumi.String("exampleHashKey"),
// 			Attributes: dynamodb.TableAttributeArray{
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("exampleHashKey"),
// 					Type: pulumi.String("S"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dynamodb.NewTableItem(ctx, "exampleTableItem", &dynamodb.TableItemArgs{
// 			TableName: exampleTable.Name,
// 			HashKey:   exampleTable.HashKey,
// 			Item:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "{\n", "  \"exampleHashKey\": {\"S\": \"something\"},\n", "  \"one\": {\"N\": \"11111\"},\n", "  \"two\": {\"N\": \"22222\"},\n", "  \"three\": {\"N\": \"33333\"},\n", "  \"four\": {\"N\": \"44444\"}\n", "}\n")),
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
// DynamoDB table items cannot be imported.
type TableItem struct {
	pulumi.CustomResourceState

	// Hash key to use for lookups and identification of the item
	HashKey pulumi.StringOutput `pulumi:"hashKey"`
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item pulumi.StringOutput `pulumi:"item"`
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey pulumi.StringPtrOutput `pulumi:"rangeKey"`
	// The name of the table to contain the item.
	TableName pulumi.StringOutput `pulumi:"tableName"`
}

// NewTableItem registers a new resource with the given unique name, arguments, and options.
func NewTableItem(ctx *pulumi.Context,
	name string, args *TableItemArgs, opts ...pulumi.ResourceOption) (*TableItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HashKey == nil {
		return nil, errors.New("invalid value for required argument 'HashKey'")
	}
	if args.Item == nil {
		return nil, errors.New("invalid value for required argument 'Item'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	var resource TableItem
	err := ctx.RegisterResource("aws:dynamodb/tableItem:TableItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableItem gets an existing TableItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableItemState, opts ...pulumi.ResourceOption) (*TableItem, error) {
	var resource TableItem
	err := ctx.ReadResource("aws:dynamodb/tableItem:TableItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableItem resources.
type tableItemState struct {
	// Hash key to use for lookups and identification of the item
	HashKey *string `pulumi:"hashKey"`
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item *string `pulumi:"item"`
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey *string `pulumi:"rangeKey"`
	// The name of the table to contain the item.
	TableName *string `pulumi:"tableName"`
}

type TableItemState struct {
	// Hash key to use for lookups and identification of the item
	HashKey pulumi.StringPtrInput
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item pulumi.StringPtrInput
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey pulumi.StringPtrInput
	// The name of the table to contain the item.
	TableName pulumi.StringPtrInput
}

func (TableItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableItemState)(nil)).Elem()
}

type tableItemArgs struct {
	// Hash key to use for lookups and identification of the item
	HashKey string `pulumi:"hashKey"`
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item string `pulumi:"item"`
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey *string `pulumi:"rangeKey"`
	// The name of the table to contain the item.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a TableItem resource.
type TableItemArgs struct {
	// Hash key to use for lookups and identification of the item
	HashKey pulumi.StringInput
	// JSON representation of a map of attribute name/value pairs, one for each attribute.
	// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item pulumi.StringInput
	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey pulumi.StringPtrInput
	// The name of the table to contain the item.
	TableName pulumi.StringInput
}

func (TableItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableItemArgs)(nil)).Elem()
}

type TableItemInput interface {
	pulumi.Input

	ToTableItemOutput() TableItemOutput
	ToTableItemOutputWithContext(ctx context.Context) TableItemOutput
}

type TableItemPtrInput interface {
	pulumi.Input

	ToTableItemPtrOutput() TableItemPtrOutput
	ToTableItemPtrOutputWithContext(ctx context.Context) TableItemPtrOutput
}

func (TableItem) ElementType() reflect.Type {
	return reflect.TypeOf((*TableItem)(nil)).Elem()
}

func (i TableItem) ToTableItemOutput() TableItemOutput {
	return i.ToTableItemOutputWithContext(context.Background())
}

func (i TableItem) ToTableItemOutputWithContext(ctx context.Context) TableItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableItemOutput)
}

func (i TableItem) ToTableItemPtrOutput() TableItemPtrOutput {
	return i.ToTableItemPtrOutputWithContext(context.Background())
}

func (i TableItem) ToTableItemPtrOutputWithContext(ctx context.Context) TableItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableItemPtrOutput)
}

type TableItemOutput struct {
	*pulumi.OutputState
}

func (TableItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableItemOutput)(nil)).Elem()
}

func (o TableItemOutput) ToTableItemOutput() TableItemOutput {
	return o
}

func (o TableItemOutput) ToTableItemOutputWithContext(ctx context.Context) TableItemOutput {
	return o
}

type TableItemPtrOutput struct {
	*pulumi.OutputState
}

func (TableItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableItem)(nil)).Elem()
}

func (o TableItemPtrOutput) ToTableItemPtrOutput() TableItemPtrOutput {
	return o
}

func (o TableItemPtrOutput) ToTableItemPtrOutputWithContext(ctx context.Context) TableItemPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TableItemOutput{})
	pulumi.RegisterOutputType(TableItemPtrOutput{})
}
