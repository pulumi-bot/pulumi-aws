// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF XSS Match Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/waf"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waf.NewXssMatchSet(ctx, "xssMatchSet", &waf.XssMatchSetArgs{
// 			XssMatchTuples: waf.XssMatchSetXssMatchTupleArray{
// 				&waf.XssMatchSetXssMatchTupleArgs{
// 					FieldToMatch: &waf.XssMatchSetXssMatchTupleFieldToMatchArgs{
// 						Type: pulumi.String("URI"),
// 					},
// 					TextTransformation: pulumi.String("NONE"),
// 				},
// 				&waf.XssMatchSetXssMatchTupleArgs{
// 					FieldToMatch: &waf.XssMatchSetXssMatchTupleFieldToMatchArgs{
// 						Type: pulumi.String("QUERY_STRING"),
// 					},
// 					TextTransformation: pulumi.String("NONE"),
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
type XssMatchSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayOutput `pulumi:"xssMatchTuples"`
}

// NewXssMatchSet registers a new resource with the given unique name, arguments, and options.
func NewXssMatchSet(ctx *pulumi.Context,
	name string, args *XssMatchSetArgs, opts ...pulumi.ResourceOption) (*XssMatchSet, error) {
	if args == nil {
		args = &XssMatchSetArgs{}
	}

	var resource XssMatchSet
	err := ctx.RegisterResource("aws:waf/xssMatchSet:XssMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetXssMatchSet gets an existing XssMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetXssMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *XssMatchSetState, opts ...pulumi.ResourceOption) (*XssMatchSet, error) {
	var resource XssMatchSet
	err := ctx.ReadResource("aws:waf/xssMatchSet:XssMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering XssMatchSet resources.
type xssMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The name or description of the SizeConstraintSet.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []XssMatchSetXssMatchTuple `pulumi:"xssMatchTuples"`
}

type XssMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayInput
}

func (XssMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*xssMatchSetState)(nil)).Elem()
}

type xssMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []XssMatchSetXssMatchTuple `pulumi:"xssMatchTuples"`
}

// The set of arguments for constructing a XssMatchSet resource.
type XssMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples XssMatchSetXssMatchTupleArrayInput
}

func (XssMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*xssMatchSetArgs)(nil)).Elem()
}

type XssMatchSetInput interface {
	pulumi.Input

	ToXssMatchSetOutput() XssMatchSetOutput
	ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput
}

func (XssMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*XssMatchSet)(nil)).Elem()
}

func (i XssMatchSet) ToXssMatchSetOutput() XssMatchSetOutput {
	return i.ToXssMatchSetOutputWithContext(context.Background())
}

func (i XssMatchSet) ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(XssMatchSetOutput)
}

type XssMatchSetOutput struct {
	*pulumi.OutputState
}

func (XssMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*XssMatchSetOutput)(nil)).Elem()
}

func (o XssMatchSetOutput) ToXssMatchSetOutput() XssMatchSetOutput {
	return o
}

func (o XssMatchSetOutput) ToXssMatchSetOutputWithContext(ctx context.Context) XssMatchSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(XssMatchSetOutput{})
}
