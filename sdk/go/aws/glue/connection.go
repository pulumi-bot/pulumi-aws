// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// ### VPC Connection
//
// For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html#connection-JDBC-VPC).
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewConnection(ctx, "example", &glue.ConnectionArgs{
// 			ConnectionProperties: pulumi.StringMap{
// 				"JDBC_CONNECTION_URL": pulumi.String(fmt.Sprintf("%v%v%v", "jdbc:mysql://", aws_rds_cluster.Example.Endpoint, "/exampledatabase")),
// 				"PASSWORD":            pulumi.String("examplepassword"),
// 				"USERNAME":            pulumi.String("exampleusername"),
// 			},
// 			PhysicalConnectionRequirements: &glue.ConnectionPhysicalConnectionRequirementsArgs{
// 				AvailabilityZone: pulumi.Any(aws_subnet.Example.Availability_zone),
// 				SecurityGroupIdLists: pulumi.StringArray{
// 					pulumi.Any(aws_security_group.Example.Id),
// 				},
// 				SubnetId: pulumi.Any(aws_subnet.Example.Id),
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
		args = &ConnectionArgs{}
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

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i *Connection) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

type ConnectionPtrInput interface {
	pulumi.Input

	ToConnectionPtrOutput() ConnectionPtrOutput
	ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput
}

type connectionPtrType ConnectionArgs

func (*connectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil))
}

func (i *connectionPtrType) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *connectionPtrType) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//          ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Connection)(nil))
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//          ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Connection)(nil))
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct {
	*pulumi.OutputState
}

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o.ToConnectionPtrOutputWithContext(context.Background())
}

func (o ConnectionOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o.ApplyT(func(v Connection) *Connection {
		return &v
	}).(ConnectionPtrOutput)
}

type ConnectionPtrOutput struct {
	*pulumi.OutputState
}

func (ConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil))
}

func (o ConnectionPtrOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Connection)(nil))
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Connection {
		return vs[0].([]Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Connection)(nil))
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Connection {
		return vs[0].(map[string]Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionPtrOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
