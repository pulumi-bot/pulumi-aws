// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SES receipt filter resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewReceiptFilter(ctx, "filter", &ses.ReceiptFilterArgs{
// 			Cidr:   pulumi.String("10.10.10.10"),
// 			Policy: pulumi.String("Block"),
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
// SES Receipt Filter can be imported using their `name`, e.g.
//
// ```sh
//  $ pulumi import aws:ses/receiptFilter:ReceiptFilter test some-filter
// ```
type ReceiptFilter struct {
	pulumi.CustomResourceState

	// The SES receipt filter ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The IP address or address range to filter, in CIDR notation
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// The name of the filter
	Name pulumi.StringOutput `pulumi:"name"`
	// Block or Allow
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewReceiptFilter registers a new resource with the given unique name, arguments, and options.
func NewReceiptFilter(ctx *pulumi.Context,
	name string, args *ReceiptFilterArgs, opts ...pulumi.ResourceOption) (*ReceiptFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cidr == nil {
		return nil, errors.New("invalid value for required argument 'Cidr'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource ReceiptFilter
	err := ctx.RegisterResource("aws:ses/receiptFilter:ReceiptFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiptFilter gets an existing ReceiptFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiptFilterState, opts ...pulumi.ResourceOption) (*ReceiptFilter, error) {
	var resource ReceiptFilter
	err := ctx.ReadResource("aws:ses/receiptFilter:ReceiptFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReceiptFilter resources.
type receiptFilterState struct {
	// The SES receipt filter ARN.
	Arn *string `pulumi:"arn"`
	// The IP address or address range to filter, in CIDR notation
	Cidr *string `pulumi:"cidr"`
	// The name of the filter
	Name *string `pulumi:"name"`
	// Block or Allow
	Policy *string `pulumi:"policy"`
}

type ReceiptFilterState struct {
	// The SES receipt filter ARN.
	Arn pulumi.StringPtrInput
	// The IP address or address range to filter, in CIDR notation
	Cidr pulumi.StringPtrInput
	// The name of the filter
	Name pulumi.StringPtrInput
	// Block or Allow
	Policy pulumi.StringPtrInput
}

func (ReceiptFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptFilterState)(nil)).Elem()
}

type receiptFilterArgs struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr string `pulumi:"cidr"`
	// The name of the filter
	Name *string `pulumi:"name"`
	// Block or Allow
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a ReceiptFilter resource.
type ReceiptFilterArgs struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr pulumi.StringInput
	// The name of the filter
	Name pulumi.StringPtrInput
	// Block or Allow
	Policy pulumi.StringInput
}

func (ReceiptFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptFilterArgs)(nil)).Elem()
}

type ReceiptFilterInput interface {
	pulumi.Input

	ToReceiptFilterOutput() ReceiptFilterOutput
	ToReceiptFilterOutputWithContext(ctx context.Context) ReceiptFilterOutput
}

func (*ReceiptFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptFilter)(nil))
}

func (i *ReceiptFilter) ToReceiptFilterOutput() ReceiptFilterOutput {
	return i.ToReceiptFilterOutputWithContext(context.Background())
}

func (i *ReceiptFilter) ToReceiptFilterOutputWithContext(ctx context.Context) ReceiptFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptFilterOutput)
}

func (i *ReceiptFilter) ToReceiptFilterPtrOutput() ReceiptFilterPtrOutput {
	return i.ToReceiptFilterPtrOutputWithContext(context.Background())
}

func (i *ReceiptFilter) ToReceiptFilterPtrOutputWithContext(ctx context.Context) ReceiptFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptFilterPtrOutput)
}

type ReceiptFilterPtrInput interface {
	pulumi.Input

	ToReceiptFilterPtrOutput() ReceiptFilterPtrOutput
	ToReceiptFilterPtrOutputWithContext(ctx context.Context) ReceiptFilterPtrOutput
}

type receiptFilterPtrType ReceiptFilterArgs

func (*receiptFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReceiptFilter)(nil))
}

func (i *receiptFilterPtrType) ToReceiptFilterPtrOutput() ReceiptFilterPtrOutput {
	return i.ToReceiptFilterPtrOutputWithContext(context.Background())
}

func (i *receiptFilterPtrType) ToReceiptFilterPtrOutputWithContext(ctx context.Context) ReceiptFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptFilterPtrOutput)
}

// ReceiptFilterArrayInput is an input type that accepts ReceiptFilterArray and ReceiptFilterArrayOutput values.
// You can construct a concrete instance of `ReceiptFilterArrayInput` via:
//
//          ReceiptFilterArray{ ReceiptFilterArgs{...} }
type ReceiptFilterArrayInput interface {
	pulumi.Input

	ToReceiptFilterArrayOutput() ReceiptFilterArrayOutput
	ToReceiptFilterArrayOutputWithContext(context.Context) ReceiptFilterArrayOutput
}

type ReceiptFilterArray []ReceiptFilterInput

func (ReceiptFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ReceiptFilter)(nil))
}

func (i ReceiptFilterArray) ToReceiptFilterArrayOutput() ReceiptFilterArrayOutput {
	return i.ToReceiptFilterArrayOutputWithContext(context.Background())
}

func (i ReceiptFilterArray) ToReceiptFilterArrayOutputWithContext(ctx context.Context) ReceiptFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptFilterArrayOutput)
}

// ReceiptFilterMapInput is an input type that accepts ReceiptFilterMap and ReceiptFilterMapOutput values.
// You can construct a concrete instance of `ReceiptFilterMapInput` via:
//
//          ReceiptFilterMap{ "key": ReceiptFilterArgs{...} }
type ReceiptFilterMapInput interface {
	pulumi.Input

	ToReceiptFilterMapOutput() ReceiptFilterMapOutput
	ToReceiptFilterMapOutputWithContext(context.Context) ReceiptFilterMapOutput
}

type ReceiptFilterMap map[string]ReceiptFilterInput

func (ReceiptFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ReceiptFilter)(nil))
}

func (i ReceiptFilterMap) ToReceiptFilterMapOutput() ReceiptFilterMapOutput {
	return i.ToReceiptFilterMapOutputWithContext(context.Background())
}

func (i ReceiptFilterMap) ToReceiptFilterMapOutputWithContext(ctx context.Context) ReceiptFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptFilterMapOutput)
}

type ReceiptFilterOutput struct {
	*pulumi.OutputState
}

func (ReceiptFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptFilter)(nil))
}

func (o ReceiptFilterOutput) ToReceiptFilterOutput() ReceiptFilterOutput {
	return o
}

func (o ReceiptFilterOutput) ToReceiptFilterOutputWithContext(ctx context.Context) ReceiptFilterOutput {
	return o
}

func (o ReceiptFilterOutput) ToReceiptFilterPtrOutput() ReceiptFilterPtrOutput {
	return o.ToReceiptFilterPtrOutputWithContext(context.Background())
}

func (o ReceiptFilterOutput) ToReceiptFilterPtrOutputWithContext(ctx context.Context) ReceiptFilterPtrOutput {
	return o.ApplyT(func(v ReceiptFilter) *ReceiptFilter {
		return &v
	}).(ReceiptFilterPtrOutput)
}

type ReceiptFilterPtrOutput struct {
	*pulumi.OutputState
}

func (ReceiptFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReceiptFilter)(nil))
}

func (o ReceiptFilterPtrOutput) ToReceiptFilterPtrOutput() ReceiptFilterPtrOutput {
	return o
}

func (o ReceiptFilterPtrOutput) ToReceiptFilterPtrOutputWithContext(ctx context.Context) ReceiptFilterPtrOutput {
	return o
}

type ReceiptFilterArrayOutput struct{ *pulumi.OutputState }

func (ReceiptFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReceiptFilter)(nil))
}

func (o ReceiptFilterArrayOutput) ToReceiptFilterArrayOutput() ReceiptFilterArrayOutput {
	return o
}

func (o ReceiptFilterArrayOutput) ToReceiptFilterArrayOutputWithContext(ctx context.Context) ReceiptFilterArrayOutput {
	return o
}

func (o ReceiptFilterArrayOutput) Index(i pulumi.IntInput) ReceiptFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReceiptFilter {
		return vs[0].([]ReceiptFilter)[vs[1].(int)]
	}).(ReceiptFilterOutput)
}

type ReceiptFilterMapOutput struct{ *pulumi.OutputState }

func (ReceiptFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReceiptFilter)(nil))
}

func (o ReceiptFilterMapOutput) ToReceiptFilterMapOutput() ReceiptFilterMapOutput {
	return o
}

func (o ReceiptFilterMapOutput) ToReceiptFilterMapOutputWithContext(ctx context.Context) ReceiptFilterMapOutput {
	return o
}

func (o ReceiptFilterMapOutput) MapIndex(k pulumi.StringInput) ReceiptFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReceiptFilter {
		return vs[0].(map[string]ReceiptFilter)[vs[1].(string)]
	}).(ReceiptFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(ReceiptFilterOutput{})
	pulumi.RegisterOutputType(ReceiptFilterPtrOutput{})
	pulumi.RegisterOutputType(ReceiptFilterArrayOutput{})
	pulumi.RegisterOutputType(ReceiptFilterMapOutput{})
}
