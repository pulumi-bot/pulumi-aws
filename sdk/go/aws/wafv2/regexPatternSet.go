// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS WAFv2 Regex Pattern Set Resource
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := wafv2.NewRegexPatternSet(ctx, "example", &wafv2.RegexPatternSetArgs{
// 			Description: pulumi.String("Example regex pattern set"),
// 			RegularExpressions: wafv2.RegexPatternSetRegularExpressionArray{
// 				&wafv2.RegexPatternSetRegularExpressionArgs{
// 					RegexString: pulumi.String("one"),
// 				},
// 				&wafv2.RegexPatternSetRegularExpressionArgs{
// 					RegexString: pulumi.String("two"),
// 				},
// 			},
// 			Scope: pulumi.String("REGIONAL"),
// 			Tags: map[string]interface{}{
// 				"Tag1": "Value1",
// 				"Tag2": "Value2",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RegexPatternSet struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A friendly description of the regular expression pattern set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	LockToken   pulumi.StringOutput    `pulumi:"lockToken"`
	// A friendly name of the regular expression pattern set.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
	RegularExpressions RegexPatternSetRegularExpressionArrayOutput `pulumi:"regularExpressions"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
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
	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn *string `pulumi:"arn"`
	// A friendly description of the regular expression pattern set.
	Description *string `pulumi:"description"`
	LockToken   *string `pulumi:"lockToken"`
	// A friendly name of the regular expression pattern set.
	Name *string `pulumi:"name"`
	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
	RegularExpressions []RegexPatternSetRegularExpression `pulumi:"regularExpressions"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope *string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type RegexPatternSetState struct {
	// The Amazon Resource Name (ARN) that identifies the cluster.
	Arn pulumi.StringPtrInput
	// A friendly description of the regular expression pattern set.
	Description pulumi.StringPtrInput
	LockToken   pulumi.StringPtrInput
	// A friendly name of the regular expression pattern set.
	Name pulumi.StringPtrInput
	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
	RegularExpressions RegexPatternSetRegularExpressionArrayInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringPtrInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.MapInput
}

func (RegexPatternSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetState)(nil)).Elem()
}

type regexPatternSetArgs struct {
	// A friendly description of the regular expression pattern set.
	Description *string `pulumi:"description"`
	// A friendly name of the regular expression pattern set.
	Name *string `pulumi:"name"`
	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
	RegularExpressions []RegexPatternSetRegularExpression `pulumi:"regularExpressions"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	// A friendly description of the regular expression pattern set.
	Description pulumi.StringPtrInput
	// A friendly name of the regular expression pattern set.
	Name pulumi.StringPtrInput
	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
	RegularExpressions RegexPatternSetRegularExpressionArrayInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringInput
	// An array of key:value pairs to associate with the resource.
	Tags pulumi.MapInput
}

func (RegexPatternSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetArgs)(nil)).Elem()
}
