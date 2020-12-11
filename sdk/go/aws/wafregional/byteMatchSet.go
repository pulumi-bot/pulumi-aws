// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regional Byte Match Set Resource for use with Application Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafregional.NewByteMatchSet(ctx, "byteSet", &wafregional.ByteMatchSetArgs{
// 			ByteMatchTuples: wafregional.ByteMatchSetByteMatchTupleArray{
// 				&wafregional.ByteMatchSetByteMatchTupleArgs{
// 					FieldToMatch: &wafregional.ByteMatchSetByteMatchTupleFieldToMatchArgs{
// 						Data: pulumi.String("referer"),
// 						Type: pulumi.String("HEADER"),
// 					},
// 					PositionalConstraint: pulumi.String("CONTAINS"),
// 					TargetString:         pulumi.String("badrefer1"),
// 					TextTransformation:   pulumi.String("NONE"),
// 				},
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
// WAF Regional Byte Match Set can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:wafregional/byteMatchSet:ByteMatchSet byte_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type ByteMatchSet struct {
	pulumi.CustomResourceState

	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayOutput `pulumi:"byteMatchTuples"`
	// The name or description of the ByteMatchSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewByteMatchSet registers a new resource with the given unique name, arguments, and options.
func NewByteMatchSet(ctx *pulumi.Context,
	name string, args *ByteMatchSetArgs, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	if args == nil {
		args = &ByteMatchSetArgs{}
	}

	var resource ByteMatchSet
	err := ctx.RegisterResource("aws:wafregional/byteMatchSet:ByteMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetByteMatchSet gets an existing ByteMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetByteMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ByteMatchSetState, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	var resource ByteMatchSet
	err := ctx.ReadResource("aws:wafregional/byteMatchSet:ByteMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ByteMatchSet resources.
type byteMatchSetState struct {
	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	// The name or description of the ByteMatchSet.
	Name *string `pulumi:"name"`
}

type ByteMatchSetState struct {
	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	// The name or description of the ByteMatchSet.
	Name pulumi.StringPtrInput
}

func (ByteMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetState)(nil)).Elem()
}

type byteMatchSetArgs struct {
	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	// The name or description of the ByteMatchSet.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ByteMatchSet resource.
type ByteMatchSetArgs struct {
	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	// The name or description of the ByteMatchSet.
	Name pulumi.StringPtrInput
}

func (ByteMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetArgs)(nil)).Elem()
}

type ByteMatchSetInput interface {
	pulumi.Input

	ToByteMatchSetOutput() ByteMatchSetOutput
	ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput
}

type ByteMatchSetPtrInput interface {
	pulumi.Input

	ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput
	ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput
}

func (ByteMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ByteMatchSet)(nil)).Elem()
}

func (i ByteMatchSet) ToByteMatchSetOutput() ByteMatchSetOutput {
	return i.ToByteMatchSetOutputWithContext(context.Background())
}

func (i ByteMatchSet) ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetOutput)
}

func (i ByteMatchSet) ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput {
	return i.ToByteMatchSetPtrOutputWithContext(context.Background())
}

func (i ByteMatchSet) ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetPtrOutput)
}

type ByteMatchSetOutput struct {
	*pulumi.OutputState
}

func (ByteMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ByteMatchSetOutput)(nil)).Elem()
}

func (o ByteMatchSetOutput) ToByteMatchSetOutput() ByteMatchSetOutput {
	return o
}

func (o ByteMatchSetOutput) ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput {
	return o
}

type ByteMatchSetPtrOutput struct {
	*pulumi.OutputState
}

func (ByteMatchSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ByteMatchSet)(nil)).Elem()
}

func (o ByteMatchSetPtrOutput) ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput {
	return o
}

func (o ByteMatchSetPtrOutput) ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ByteMatchSetOutput{})
	pulumi.RegisterOutputType(ByteMatchSetPtrOutput{})
}
