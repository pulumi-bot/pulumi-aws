// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Direct Connect Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := directconnect.NewGateway(ctx, "example", &directconnect.GatewayArgs{
// 			AmazonSideAsn: pulumi.String("64512"),
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
// Direct Connect Gateways can be imported using the `gateway id`, e.g.
//
// ```sh
//  $ pulumi import aws:directconnect/gateway:Gateway test abcd1234-dcba-5678-be23-cdef9876ab45
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
	AmazonSideAsn pulumi.StringOutput `pulumi:"amazonSideAsn"`
	// The name of the connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// AWS Account ID of the gateway.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil || args.AmazonSideAsn == nil {
		return nil, errors.New("missing required argument 'AmazonSideAsn'")
	}
	if args == nil {
		args = &GatewayArgs{}
	}
	var resource Gateway
	err := ctx.RegisterResource("aws:directconnect/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("aws:directconnect/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The name of the connection.
	Name *string `pulumi:"name"`
	// AWS Account ID of the gateway.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
}

type GatewayState struct {
	// The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
	AmazonSideAsn pulumi.StringPtrInput
	// The name of the connection.
	Name pulumi.StringPtrInput
	// AWS Account ID of the gateway.
	OwnerAccountId pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
	AmazonSideAsn string `pulumi:"amazonSideAsn"`
	// The name of the connection.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// The ASN to be configured on the Amazon side of the connection. The ASN must be in the private range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294.
	AmazonSideAsn pulumi.StringInput
	// The name of the connection.
	Name pulumi.StringPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}
