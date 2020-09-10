// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RegexPatternSet struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput                         `pulumi:"arn"`
	Description        pulumi.StringPtrOutput                      `pulumi:"description"`
	LockToken          pulumi.StringOutput                         `pulumi:"lockToken"`
	Name               pulumi.StringOutput                         `pulumi:"name"`
	RegularExpressions RegexPatternSetRegularExpressionArrayOutput `pulumi:"regularExpressions"`
	Scope              pulumi.StringOutput                         `pulumi:"scope"`
	Tags               pulumi.StringMapOutput                      `pulumi:"tags"`
}

// NewRegexPatternSet registers a new resource with the given unique name, arguments, and options.
func NewRegexPatternSet(ctx *pulumi.Context,
	name string, args *RegexPatternSetArgs, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RegexPatternSetArgs{}
	}
	var resource RegexPatternSet
	err := ctx.RegisterResource("aws:wafv2/regexPatternSet:RegexPatternSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexPatternSet gets an existing RegexPatternSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexPatternSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexPatternSetState, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	var resource RegexPatternSet
	err := ctx.ReadResource("aws:wafv2/regexPatternSet:RegexPatternSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexPatternSet resources.
type regexPatternSetState struct {
	Arn                *string                            `pulumi:"arn"`
	Description        *string                            `pulumi:"description"`
	LockToken          *string                            `pulumi:"lockToken"`
	Name               *string                            `pulumi:"name"`
	RegularExpressions []RegexPatternSetRegularExpression `pulumi:"regularExpressions"`
	Scope              *string                            `pulumi:"scope"`
	Tags               map[string]string                  `pulumi:"tags"`
}

type RegexPatternSetState struct {
	Arn                pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	LockToken          pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	RegularExpressions RegexPatternSetRegularExpressionArrayInput
	Scope              pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
}

func (RegexPatternSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetState)(nil)).Elem()
}

type regexPatternSetArgs struct {
	Description        *string                            `pulumi:"description"`
	Name               *string                            `pulumi:"name"`
	RegularExpressions []RegexPatternSetRegularExpression `pulumi:"regularExpressions"`
	Scope              string                             `pulumi:"scope"`
	Tags               map[string]string                  `pulumi:"tags"`
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	Description        pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	RegularExpressions RegexPatternSetRegularExpressionArrayInput
	Scope              pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (RegexPatternSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetArgs)(nil)).Elem()
}
