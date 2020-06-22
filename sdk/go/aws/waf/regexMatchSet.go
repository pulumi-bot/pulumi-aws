// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regex Match Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRegexPatternSet, err := waf.NewRegexPatternSet(ctx, "exampleRegexPatternSet", &waf.RegexPatternSetArgs{
// 			RegexPatternStrings: pulumi.StringArray{
// 				pulumi.String("one"),
// 				pulumi.String("two"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = waf.NewRegexMatchSet(ctx, "exampleRegexMatchSet", &waf.RegexMatchSetArgs{
// 			RegexMatchTuples: waf.RegexMatchSetRegexMatchTupleArray{
// 				&waf.RegexMatchSetRegexMatchTupleArgs{
// 					FieldToMatch: &waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs{
// 						Data: pulumi.String("User-Agent"),
// 						Type: pulumi.String("HEADER"),
// 					},
// 					RegexPatternSetId:  exampleRegexPatternSet.ID(),
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
type RegexMatchSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or description of the Regex Match Set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayOutput `pulumi:"regexMatchTuples"`
}

// NewRegexMatchSet registers a new resource with the given unique name, arguments, and options.
func NewRegexMatchSet(ctx *pulumi.Context,
	name string, args *RegexMatchSetArgs, opts ...pulumi.ResourceOption) (*RegexMatchSet, error) {
	if args == nil {
		args = &RegexMatchSetArgs{}
	}
	var resource RegexMatchSet
	err := ctx.RegisterResource("aws:waf/regexMatchSet:RegexMatchSet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:waf/regexMatchSet:RegexMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexMatchSet resources.
type regexMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The name or description of the Regex Match Set.
	Name *string `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

type RegexMatchSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The name or description of the Regex Match Set.
	Name pulumi.StringPtrInput
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetState)(nil)).Elem()
}

type regexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name *string `pulumi:"name"`
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []RegexMatchSetRegexMatchTuple `pulumi:"regexMatchTuples"`
}

// The set of arguments for constructing a RegexMatchSet resource.
type RegexMatchSetArgs struct {
	// The name or description of the Regex Match Set.
	Name pulumi.StringPtrInput
	// The regular expression pattern that you want AWS WAF to search for in web requests,
	// the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples RegexMatchSetRegexMatchTupleArrayInput
}

func (RegexMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexMatchSetArgs)(nil)).Elem()
}
