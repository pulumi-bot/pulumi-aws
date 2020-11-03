// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage a Default VPC Routing Table.
//
// Each VPC created in AWS comes with a Default Route Table that can be managed, but not
// destroyed. **This is an advanced resource**, and has special caveats to be aware
// of when using it. Please read this document in its entirety before using this
// resource. It is recommended you **do not** use both `ec2.DefaultRouteTable` to
// manage the default route table **and** use the `ec2.MainRouteTableAssociation`,
// due to possible conflict in routes.
//
// The `ec2.DefaultRouteTable` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead attempts to "adopt" it
// into management. We can do this because each VPC created has a Default Route
// Table that cannot be destroyed, and is created with a single route.
//
// When this provider first adopts the Default Route Table, it **immediately removes all
// defined routes**. It then proceeds to create any routes specified in the
// configuration. This step is required so that only the routes specified in the
// configuration present in the Default Route Table.
//
// For more information about Route Tables, see the AWS Documentation on
// [Route Tables][aws-route-tables].
//
// For more information about managing normal Route Tables in this provider, see our
// documentation on [ec2.RouteTable][tf-route-tables].
//
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource and a Route Table resource with routes
// defined in-line. At this time you cannot use a Route Table with in-line routes
// in conjunction with any Route resources. Doing so will cause
// a conflict of rule settings and will overwrite routes.
//
// ## Example Usage
// ### With Tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewDefaultRouteTable(ctx, "defaultRouteTable", &ec2.DefaultRouteTableArgs{
// 			DefaultRouteTableId: pulumi.Any(aws_vpc.Foo.Default_route_table_id),
// 			Routes: ec2.DefaultRouteTableRouteArray{
// 				nil,
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("default table"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DefaultRouteTable struct {
	pulumi.CustomResourceState

	// The ID of the Default Routing Table.
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`
	// The ID of the AWS account that owns the route table
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayOutput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	Routes DefaultRouteTableRouteArrayOutput `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags  pulumi.StringMapOutput `pulumi:"tags"`
	VpcId pulumi.StringOutput    `pulumi:"vpcId"`
}

// NewDefaultRouteTable registers a new resource with the given unique name, arguments, and options.
func NewDefaultRouteTable(ctx *pulumi.Context,
	name string, args *DefaultRouteTableArgs, opts ...pulumi.ResourceOption) (*DefaultRouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DefaultRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'DefaultRouteTableId'")
	}
	var resource DefaultRouteTable
	err := ctx.RegisterResource("aws:ec2/defaultRouteTable:DefaultRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultRouteTable gets an existing DefaultRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultRouteTableState, opts ...pulumi.ResourceOption) (*DefaultRouteTable, error) {
	var resource DefaultRouteTable
	err := ctx.ReadResource("aws:ec2/defaultRouteTable:DefaultRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultRouteTable resources.
type defaultRouteTableState struct {
	// The ID of the Default Routing Table.
	DefaultRouteTableId *string `pulumi:"defaultRouteTableId"`
	// The ID of the AWS account that owns the route table
	OwnerId *string `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws []string `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	Routes []DefaultRouteTableRoute `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags  map[string]string `pulumi:"tags"`
	VpcId *string           `pulumi:"vpcId"`
}

type DefaultRouteTableState struct {
	// The ID of the Default Routing Table.
	DefaultRouteTableId pulumi.StringPtrInput
	// The ID of the AWS account that owns the route table
	OwnerId pulumi.StringPtrInput
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayInput
	// A list of route objects. Their keys are documented below.
	Routes DefaultRouteTableRouteArrayInput
	// A mapping of tags to assign to the resource.
	Tags  pulumi.StringMapInput
	VpcId pulumi.StringPtrInput
}

func (DefaultRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRouteTableState)(nil)).Elem()
}

type defaultRouteTableArgs struct {
	// The ID of the Default Routing Table.
	DefaultRouteTableId string `pulumi:"defaultRouteTableId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws []string `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	Routes []DefaultRouteTableRoute `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultRouteTable resource.
type DefaultRouteTableArgs struct {
	// The ID of the Default Routing Table.
	DefaultRouteTableId pulumi.StringInput
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayInput
	// A list of route objects. Their keys are documented below.
	Routes DefaultRouteTableRouteArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefaultRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRouteTableArgs)(nil)).Elem()
}
