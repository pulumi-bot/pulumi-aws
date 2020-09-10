// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RegexMatchSet struct {
	pulumi.CustomResourceState

	Name             pulumi.StringOutput                     `pulumi:"name"`
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayOutput `pulumi:"regexMatchTuples"`
}

// NewRegexMatchSet registers a new resource with the given unique name, arguments, and options.
func NewRegexMatchSet(ctx *pulumi.Context,
	name string, args *RegexMatchSetArgs, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	if args == nil {
		args = &RegexMatchSetArgs{}
	}
	var resource RegexMatchSet
	err := ctx.RegisterResource("aws:wafregional/regexMatchSet:RegexMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexMatchSet gets an existing RegexMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexMatchSetState, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	var resource RegexMatchSet
	err := ctx.ReadResource("aws:wafregional/regexMatchSet:RegexMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexMatchSet resources.
type regexMatchSetState struct {
	Name             *string                        `pulumi:"name"`
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

type RegexMatchSetState struct {
	Name             pulumi.StringPtrInput
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetState)(nil)).Elem()
}

type regexMatchSetArgs struct {
	Name             *string                        `pulumi:"name"`
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

// The set of arguments for constructing a RegexMatchSet resource.
type RegexMatchSetArgs struct {
	Name             pulumi.StringPtrInput
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetArgs)(nil)).Elem()
}
