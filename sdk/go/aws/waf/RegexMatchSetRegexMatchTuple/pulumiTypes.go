// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package RegexMatchSetRegexMatchTuple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/waf/RegexMatchSetRegexMatchTupleFieldToMatch"
)

type RegexMatchSetRegexMatchTuple struct {
	// The part of a web request that you want to search, such as a specified header or a query string.
	FieldToMatch wafRegexMatchSetRegexMatchTupleFieldToMatch.RegexMatchSetRegexMatchTupleFieldToMatch `pulumi:"fieldToMatch"`
	// The ID of a [Regex Pattern Set](https://www.terraform.io/docs/providers/aws/r/waf_regex_pattern_set.html).
	RegexPatternSetId string `pulumi:"regexPatternSetId"`
	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
	// for all supported values.
	TextTransformation string `pulumi:"textTransformation"`
}

type RegexMatchSetRegexMatchTupleInput interface {
	pulumi.Input

	ToRegexMatchSetRegexMatchTupleOutput() RegexMatchSetRegexMatchTupleOutput
	ToRegexMatchSetRegexMatchTupleOutputWithContext(context.Context) RegexMatchSetRegexMatchTupleOutput
}

type RegexMatchSetRegexMatchTupleArgs struct {
	// The part of a web request that you want to search, such as a specified header or a query string.
	FieldToMatch wafRegexMatchSetRegexMatchTupleFieldToMatch.RegexMatchSetRegexMatchTupleFieldToMatchInput `pulumi:"fieldToMatch"`
	// The ID of a [Regex Pattern Set](https://www.terraform.io/docs/providers/aws/r/waf_regex_pattern_set.html).
	RegexPatternSetId pulumi.StringInput `pulumi:"regexPatternSetId"`
	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
	// for all supported values.
	TextTransformation pulumi.StringInput `pulumi:"textTransformation"`
}

func (RegexMatchSetRegexMatchTupleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexMatchSetRegexMatchTuple)(nil)).Elem()
}

func (i RegexMatchSetRegexMatchTupleArgs) ToRegexMatchSetRegexMatchTupleOutput() RegexMatchSetRegexMatchTupleOutput {
	return i.ToRegexMatchSetRegexMatchTupleOutputWithContext(context.Background())
}

func (i RegexMatchSetRegexMatchTupleArgs) ToRegexMatchSetRegexMatchTupleOutputWithContext(ctx context.Context) RegexMatchSetRegexMatchTupleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetRegexMatchTupleOutput)
}

type RegexMatchSetRegexMatchTupleArrayInput interface {
	pulumi.Input

	ToRegexMatchSetRegexMatchTupleArrayOutput() RegexMatchSetRegexMatchTupleArrayOutput
	ToRegexMatchSetRegexMatchTupleArrayOutputWithContext(context.Context) RegexMatchSetRegexMatchTupleArrayOutput
}

type RegexMatchSetRegexMatchTupleArray []RegexMatchSetRegexMatchTupleInput

func (RegexMatchSetRegexMatchTupleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegexMatchSetRegexMatchTuple)(nil)).Elem()
}

func (i RegexMatchSetRegexMatchTupleArray) ToRegexMatchSetRegexMatchTupleArrayOutput() RegexMatchSetRegexMatchTupleArrayOutput {
	return i.ToRegexMatchSetRegexMatchTupleArrayOutputWithContext(context.Background())
}

func (i RegexMatchSetRegexMatchTupleArray) ToRegexMatchSetRegexMatchTupleArrayOutputWithContext(ctx context.Context) RegexMatchSetRegexMatchTupleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexMatchSetRegexMatchTupleArrayOutput)
}

type RegexMatchSetRegexMatchTupleOutput struct { *pulumi.OutputState }

func (RegexMatchSetRegexMatchTupleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexMatchSetRegexMatchTuple)(nil)).Elem()
}

func (o RegexMatchSetRegexMatchTupleOutput) ToRegexMatchSetRegexMatchTupleOutput() RegexMatchSetRegexMatchTupleOutput {
	return o
}

func (o RegexMatchSetRegexMatchTupleOutput) ToRegexMatchSetRegexMatchTupleOutputWithContext(ctx context.Context) RegexMatchSetRegexMatchTupleOutput {
	return o
}

// The part of a web request that you want to search, such as a specified header or a query string.
func (o RegexMatchSetRegexMatchTupleOutput) FieldToMatch() wafRegexMatchSetRegexMatchTupleFieldToMatch.RegexMatchSetRegexMatchTupleFieldToMatchOutput {
	return o.ApplyT(func (v RegexMatchSetRegexMatchTuple) wafRegexMatchSetRegexMatchTupleFieldToMatch.RegexMatchSetRegexMatchTupleFieldToMatch { return v.FieldToMatch }).(wafRegexMatchSetRegexMatchTupleFieldToMatch.RegexMatchSetRegexMatchTupleFieldToMatchOutput)
}

// The ID of a [Regex Pattern Set](https://www.terraform.io/docs/providers/aws/r/waf_regex_pattern_set.html).
func (o RegexMatchSetRegexMatchTupleOutput) RegexPatternSetId() pulumi.StringOutput {
	return o.ApplyT(func (v RegexMatchSetRegexMatchTuple) string { return v.RegexPatternSetId }).(pulumi.StringOutput)
}

// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_ByteMatchTuple.html#WAF-Type-ByteMatchTuple-TextTransformation)
// for all supported values.
func (o RegexMatchSetRegexMatchTupleOutput) TextTransformation() pulumi.StringOutput {
	return o.ApplyT(func (v RegexMatchSetRegexMatchTuple) string { return v.TextTransformation }).(pulumi.StringOutput)
}

type RegexMatchSetRegexMatchTupleArrayOutput struct { *pulumi.OutputState}

func (RegexMatchSetRegexMatchTupleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegexMatchSetRegexMatchTuple)(nil)).Elem()
}

func (o RegexMatchSetRegexMatchTupleArrayOutput) ToRegexMatchSetRegexMatchTupleArrayOutput() RegexMatchSetRegexMatchTupleArrayOutput {
	return o
}

func (o RegexMatchSetRegexMatchTupleArrayOutput) ToRegexMatchSetRegexMatchTupleArrayOutputWithContext(ctx context.Context) RegexMatchSetRegexMatchTupleArrayOutput {
	return o
}

func (o RegexMatchSetRegexMatchTupleArrayOutput) Index(i pulumi.IntInput) RegexMatchSetRegexMatchTupleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) RegexMatchSetRegexMatchTuple {
		return vs[0].([]RegexMatchSetRegexMatchTuple)[vs[1].(int)]
	}).(RegexMatchSetRegexMatchTupleOutput)
}

func init() {
	pulumi.RegisterOutputType(RegexMatchSetRegexMatchTupleOutput{})
	pulumi.RegisterOutputType(RegexMatchSetRegexMatchTupleArrayOutput{})
}
