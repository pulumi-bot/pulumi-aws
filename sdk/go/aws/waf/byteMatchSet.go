// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Byte Match Set Resource
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
// 		_, err := waf.NewByteMatchSet(ctx, "byteSet", &waf.ByteMatchSetArgs{
// 			ByteMatchTuples: waf.ByteMatchSetByteMatchTupleArray{
// 				&waf.ByteMatchSetByteMatchTupleArgs{
// 					FieldToMatch: &waf.ByteMatchSetByteMatchTupleFieldToMatchArgs{
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
// WAF Byte Match Set can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import aws:waf/byteMatchSet:ByteMatchSet byte_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type ByteMatchSet struct {
	pulumi.CustomResourceState

	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayOutput `pulumi:"byteMatchTuples"`
	// The name or description of the Byte Match Set.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewByteMatchSet registers a new resource with the given unique name, arguments, and options.
func NewByteMatchSet(ctx *pulumi.Context,
	name string, args *ByteMatchSetArgs, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	if args == nil {
		args = &ByteMatchSetArgs{}
	}
	var resource ByteMatchSet
	err := ctx.RegisterResource("aws:waf/byteMatchSet:ByteMatchSet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:waf/byteMatchSet:ByteMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ByteMatchSet resources.
type byteMatchSetState struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	// The name or description of the Byte Match Set.
	Name *string `pulumi:"name"`
}

type ByteMatchSetState struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	// The name or description of the Byte Match Set.
	Name pulumi.StringPtrInput
}

func (ByteMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetState)(nil)).Elem()
}

type byteMatchSetArgs struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	// The name or description of the Byte Match Set.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ByteMatchSet resource.
type ByteMatchSetArgs struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	// The name or description of the Byte Match Set.
	Name pulumi.StringPtrInput
}

func (ByteMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetArgs)(nil)).Elem()
}
