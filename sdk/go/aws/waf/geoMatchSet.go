// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Geo Match Set Resource
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
// 		_, err := waf.NewGeoMatchSet(ctx, "geoMatchSet", &waf.GeoMatchSetArgs{
// 			GeoMatchConstraints: waf.GeoMatchSetGeoMatchConstraintArray{
// 				&waf.GeoMatchSetGeoMatchConstraintArgs{
// 					Type:  pulumi.String("Country"),
// 					Value: pulumi.String("US"),
// 				},
// 				&waf.GeoMatchSetGeoMatchConstraintArgs{
// 					Type:  pulumi.String("Country"),
// 					Value: pulumi.String("CA"),
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
// WAF Geo Match Set can be imported using their ID, e.g.
//
// ```sh
//  $ pulumi import aws:waf/geoMatchSet:GeoMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type GeoMatchSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayOutput `pulumi:"geoMatchConstraints"`
	// The name or description of the GeoMatchSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGeoMatchSet registers a new resource with the given unique name, arguments, and options.
func NewGeoMatchSet(ctx *pulumi.Context,
	name string, args *GeoMatchSetArgs, opts ...pulumi.ResourceOption) (*GeoMatchSet, error) {
	if args == nil {
		args = &GeoMatchSetArgs{}
	}
	var resource GeoMatchSet
	err := ctx.RegisterResource("aws:waf/geoMatchSet:GeoMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeoMatchSet gets an existing GeoMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeoMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeoMatchSetState, opts ...pulumi.ResourceOption) (*GeoMatchSet, error) {
	var resource GeoMatchSet
	err := ctx.ReadResource("aws:waf/geoMatchSet:GeoMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeoMatchSet resources.
type geoMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	// The name or description of the GeoMatchSet.
	Name *string `pulumi:"name"`
}

type GeoMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayInput
	// The name or description of the GeoMatchSet.
	Name pulumi.StringPtrInput
}

func (GeoMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*geoMatchSetState)(nil)).Elem()
}

type geoMatchSetArgs struct {
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints []GeoMatchSetGeoMatchConstraint `pulumi:"geoMatchConstraints"`
	// The name or description of the GeoMatchSet.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a GeoMatchSet resource.
type GeoMatchSetArgs struct {
	// The GeoMatchConstraint objects which contain the country that you want AWS WAF to search for.
	GeoMatchConstraints GeoMatchSetGeoMatchConstraintArrayInput
	// The name or description of the GeoMatchSet.
	Name pulumi.StringPtrInput
}

func (GeoMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geoMatchSetArgs)(nil)).Elem()
}

type GeoMatchSetInput interface {
	pulumi.Input

	ToGeoMatchSetOutput() GeoMatchSetOutput
	ToGeoMatchSetOutputWithContext(ctx context.Context) GeoMatchSetOutput
}

func (GeoMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoMatchSet)(nil)).Elem()
}

func (i GeoMatchSet) ToGeoMatchSetOutput() GeoMatchSetOutput {
	return i.ToGeoMatchSetOutputWithContext(context.Background())
}

func (i GeoMatchSet) ToGeoMatchSetOutputWithContext(ctx context.Context) GeoMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetOutput)
}

type GeoMatchSetOutput struct {
	*pulumi.OutputState
}

func (GeoMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoMatchSetOutput)(nil)).Elem()
}

func (o GeoMatchSetOutput) ToGeoMatchSetOutput() GeoMatchSetOutput {
	return o
}

func (o GeoMatchSetOutput) ToGeoMatchSetOutputWithContext(ctx context.Context) GeoMatchSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GeoMatchSetOutput{})
}
