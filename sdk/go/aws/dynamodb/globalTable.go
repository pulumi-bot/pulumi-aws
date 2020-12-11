// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html). These are layered on top of existing DynamoDB Tables.
//
// > **NOTE:** To instead manage [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html), use the `dynamodb.Table` resource `replica` configuration block.
//
// > Note: There are many restrictions before you can properly create DynamoDB Global Tables in multiple regions. See the [AWS DynamoDB Global Table Requirements](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables_reqs_bestpractices.html) for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/dynamodb"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/providers"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "us_east_1", &providers.awsArgs{
// 			Region: pulumi.String("us-east-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = providers.Newaws(ctx, "us_west_2", &providers.awsArgs{
// 			Region: pulumi.String("us-west-2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dynamodb.NewTable(ctx, "us_east_1Table", &dynamodb.TableArgs{
// 			HashKey:        pulumi.String("myAttribute"),
// 			StreamEnabled:  pulumi.Bool(true),
// 			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
// 			ReadCapacity:   pulumi.Int(1),
// 			WriteCapacity:  pulumi.Int(1),
// 			Attributes: dynamodb.TableAttributeArray{
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("myAttribute"),
// 					Type: pulumi.String("S"),
// 				},
// 			},
// 		}, pulumi.Provider(aws.Us-east-1))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dynamodb.NewTable(ctx, "us_west_2Table", &dynamodb.TableArgs{
// 			HashKey:        pulumi.String("myAttribute"),
// 			StreamEnabled:  pulumi.Bool(true),
// 			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
// 			ReadCapacity:   pulumi.Int(1),
// 			WriteCapacity:  pulumi.Int(1),
// 			Attributes: dynamodb.TableAttributeArray{
// 				&dynamodb.TableAttributeArgs{
// 					Name: pulumi.String("myAttribute"),
// 					Type: pulumi.String("S"),
// 				},
// 			},
// 		}, pulumi.Provider(aws.Us-west-2))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dynamodb.NewGlobalTable(ctx, "myTable", &dynamodb.GlobalTableArgs{
// 			Replicas: dynamodb.GlobalTableReplicaArray{
// 				&dynamodb.GlobalTableReplicaArgs{
// 					RegionName: pulumi.String("us-east-1"),
// 				},
// 				&dynamodb.GlobalTableReplicaArgs{
// 					RegionName: pulumi.String("us-west-2"),
// 				},
// 			},
// 		}, pulumi.Provider(aws.Us-east-1), pulumi.DependsOn([]pulumi.Resource{
// 			us_east_1Table,
// 			us_west_2Table,
// 		}))
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
// DynamoDB Global Tables can be imported using the global table name, e.g.
//
// ```sh
//  $ pulumi import aws:dynamodb/globalTable:GlobalTable MyTable MyTable
// ```
type GlobalTable struct {
	pulumi.CustomResourceState

	// The ARN of the DynamoDB Global Table
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the global table. Must match underlying DynamoDB Table names in all regions.
	Name pulumi.StringOutput `pulumi:"name"`
	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replicas GlobalTableReplicaArrayOutput `pulumi:"replicas"`
}

// NewGlobalTable registers a new resource with the given unique name, arguments, and options.
func NewGlobalTable(ctx *pulumi.Context,
	name string, args *GlobalTableArgs, opts ...pulumi.ResourceOption) (*GlobalTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Replicas == nil {
		return nil, errors.New("invalid value for required argument 'Replicas'")
	}
	var resource GlobalTable
	err := ctx.RegisterResource("aws:dynamodb/globalTable:GlobalTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalTable gets an existing GlobalTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalTableState, opts ...pulumi.ResourceOption) (*GlobalTable, error) {
	var resource GlobalTable
	err := ctx.ReadResource("aws:dynamodb/globalTable:GlobalTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalTable resources.
type globalTableState struct {
	// The ARN of the DynamoDB Global Table
	Arn *string `pulumi:"arn"`
	// The name of the global table. Must match underlying DynamoDB Table names in all regions.
	Name *string `pulumi:"name"`
	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replicas []GlobalTableReplica `pulumi:"replicas"`
}

type GlobalTableState struct {
	// The ARN of the DynamoDB Global Table
	Arn pulumi.StringPtrInput
	// The name of the global table. Must match underlying DynamoDB Table names in all regions.
	Name pulumi.StringPtrInput
	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replicas GlobalTableReplicaArrayInput
}

func (GlobalTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalTableState)(nil)).Elem()
}

type globalTableArgs struct {
	// The name of the global table. Must match underlying DynamoDB Table names in all regions.
	Name *string `pulumi:"name"`
	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replicas []GlobalTableReplica `pulumi:"replicas"`
}

// The set of arguments for constructing a GlobalTable resource.
type GlobalTableArgs struct {
	// The name of the global table. Must match underlying DynamoDB Table names in all regions.
	Name pulumi.StringPtrInput
	// Underlying DynamoDB Table. At least 1 replica must be defined. See below.
	Replicas GlobalTableReplicaArrayInput
}

func (GlobalTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalTableArgs)(nil)).Elem()
}

type GlobalTableInput interface {
	pulumi.Input

	ToGlobalTableOutput() GlobalTableOutput
	ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput
}

type GlobalTablePtrInput interface {
	pulumi.Input

	ToGlobalTablePtrOutput() GlobalTablePtrOutput
	ToGlobalTablePtrOutputWithContext(ctx context.Context) GlobalTablePtrOutput
}

func (GlobalTable) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTable)(nil)).Elem()
}

func (i GlobalTable) ToGlobalTableOutput() GlobalTableOutput {
	return i.ToGlobalTableOutputWithContext(context.Background())
}

func (i GlobalTable) ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalTableOutput)
}

func (i GlobalTable) ToGlobalTablePtrOutput() GlobalTablePtrOutput {
	return i.ToGlobalTablePtrOutputWithContext(context.Background())
}

func (i GlobalTable) ToGlobalTablePtrOutputWithContext(ctx context.Context) GlobalTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalTablePtrOutput)
}

type GlobalTableOutput struct {
	*pulumi.OutputState
}

func (GlobalTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTableOutput)(nil)).Elem()
}

func (o GlobalTableOutput) ToGlobalTableOutput() GlobalTableOutput {
	return o
}

func (o GlobalTableOutput) ToGlobalTableOutputWithContext(ctx context.Context) GlobalTableOutput {
	return o
}

type GlobalTablePtrOutput struct {
	*pulumi.OutputState
}

func (GlobalTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalTable)(nil)).Elem()
}

func (o GlobalTablePtrOutput) ToGlobalTablePtrOutput() GlobalTablePtrOutput {
	return o
}

func (o GlobalTablePtrOutput) ToGlobalTablePtrOutputWithContext(ctx context.Context) GlobalTablePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GlobalTableOutput{})
	pulumi.RegisterOutputType(GlobalTablePtrOutput{})
}
