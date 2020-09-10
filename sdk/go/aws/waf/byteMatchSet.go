// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ByteMatchSet struct {
	pulumi.CustomResourceState

	ByteMatchTuples ByteMatchSetByteMatchTupleArrayOutput `pulumi:"byteMatchTuples"`
	Name            pulumi.StringOutput                   `pulumi:"name"`
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
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	Name            *string                      `pulumi:"name"`
}

type ByteMatchSetState struct {
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	Name            pulumi.StringPtrInput
}

func (ByteMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetState)(nil)).Elem()
}

type byteMatchSetArgs struct {
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	Name            *string                      `pulumi:"name"`
}

// The set of arguments for constructing a ByteMatchSet resource.
type ByteMatchSetArgs struct {
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	Name            pulumi.StringPtrInput
}

func (ByteMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetArgs)(nil)).Elem()
}
