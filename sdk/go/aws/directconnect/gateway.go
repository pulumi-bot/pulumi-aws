// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AmazonSideAsn == nil {
		return nil, errors.New("invalid value for required argument 'AmazonSideAsn'")
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

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

func (i *Gateway) ToGatewayPtrOutput() GatewayPtrOutput {
	return i.ToGatewayPtrOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPtrOutput)
}

type GatewayPtrInput interface {
	pulumi.Input

	ToGatewayPtrOutput() GatewayPtrOutput
	ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput
}

type gatewayPtrType GatewayArgs

func (*gatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil))
}

func (i *gatewayPtrType) ToGatewayPtrOutput() GatewayPtrOutput {
	return i.ToGatewayPtrOutputWithContext(context.Background())
}

func (i *gatewayPtrType) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPtrOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//          GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Gateway)(nil))
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//          GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Gateway)(nil))
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct {
	*pulumi.OutputState
}

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayPtrOutput() GatewayPtrOutput {
	return o.ToGatewayPtrOutputWithContext(context.Background())
}

func (o GatewayOutput) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return o.ApplyT(func(v Gateway) *Gateway {
		return &v
	}).(GatewayPtrOutput)
}

type GatewayPtrOutput struct {
	*pulumi.OutputState
}

func (GatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil))
}

func (o GatewayPtrOutput) ToGatewayPtrOutput() GatewayPtrOutput {
	return o
}

func (o GatewayPtrOutput) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return o
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Gateway)(nil))
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Gateway {
		return vs[0].([]Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Gateway)(nil))
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Gateway {
		return vs[0].(map[string]Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayPtrOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
