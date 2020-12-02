// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Glue Connection resource.
//
// ## Example Usage
// ### Non-VPC Connection
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewConnection(ctx, "example", &glue.ConnectionArgs{
// 			ConnectionProperties: pulumi.StringMap{
// 				"JDBC_CONNECTION_URL": pulumi.String("jdbc:mysql://example.com/exampledatabase"),
// 				"PASSWORD":            pulumi.String("examplepassword"),
// 				"USERNAME":            pulumi.String("exampleusername"),
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
// Glue Connections can be imported using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`, e.g.
//
// ```sh
//  $ pulumi import aws:glue/connection:Connection MyConnection 123456789012:MyConnection
// ```
type Connection struct {
	pulumi.CustomResourceState

	// The ARN of the Glue Connection.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// A map of key-value pairs used as parameters for this connection.
	ConnectionProperties pulumi.StringMapOutput `pulumi:"connectionProperties"`
	// The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
	ConnectionType pulumi.StringPtrOutput `pulumi:"connectionType"`
	// Description of the connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of criteria that can be used in selecting this connection.
	MatchCriterias pulumi.StringArrayOutput `pulumi:"matchCriterias"`
	// The name of the connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements ConnectionPhysicalConnectionRequirementsPtrOutput `pulumi:"physicalConnectionRequirements"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProperties'")
	}
	var resource Connection
	err := ctx.RegisterResource("aws:glue/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("aws:glue/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// The ARN of the Glue Connection.
	Arn *string `pulumi:"arn"`
	// The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId *string `pulumi:"catalogId"`
	// A map of key-value pairs used as parameters for this connection.
	ConnectionProperties map[string]string `pulumi:"connectionProperties"`
	// The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
	ConnectionType *string `pulumi:"connectionType"`
	// Description of the connection.
	Description *string `pulumi:"description"`
	// A list of criteria that can be used in selecting this connection.
	MatchCriterias []string `pulumi:"matchCriterias"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements *ConnectionPhysicalConnectionRequirements `pulumi:"physicalConnectionRequirements"`
}

type ConnectionState struct {
	// The ARN of the Glue Connection.
	Arn pulumi.StringPtrInput
	// The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId pulumi.StringPtrInput
	// A map of key-value pairs used as parameters for this connection.
	ConnectionProperties pulumi.StringMapInput
	// The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
	ConnectionType pulumi.StringPtrInput
	// Description of the connection.
	Description pulumi.StringPtrInput
	// A list of criteria that can be used in selecting this connection.
	MatchCriterias pulumi.StringArrayInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements ConnectionPhysicalConnectionRequirementsPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId *string `pulumi:"catalogId"`
	// A map of key-value pairs used as parameters for this connection.
	ConnectionProperties map[string]string `pulumi:"connectionProperties"`
	// The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
	ConnectionType *string `pulumi:"connectionType"`
	// Description of the connection.
	Description *string `pulumi:"description"`
	// A list of criteria that can be used in selecting this connection.
	MatchCriterias []string `pulumi:"matchCriterias"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements *ConnectionPhysicalConnectionRequirements `pulumi:"physicalConnectionRequirements"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId pulumi.StringPtrInput
	// A map of key-value pairs used as parameters for this connection.
	ConnectionProperties pulumi.StringMapInput
	// The type of the connection. Supported are: `JDBC`, `MONGODB`, `KAFKA`, and `NETWORK`. Defaults to `JBDC`.
	ConnectionType pulumi.StringPtrInput
	// Description of the connection.
	Description pulumi.StringPtrInput
	// A list of criteria that can be used in selecting this connection.
	MatchCriterias pulumi.StringArrayInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements ConnectionPhysicalConnectionRequirementsPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil)).Elem()
}

func (i Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct {
	*pulumi.OutputState
}

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionOutput)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
}
