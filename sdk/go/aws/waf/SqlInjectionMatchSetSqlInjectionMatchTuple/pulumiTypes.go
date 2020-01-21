// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package SqlInjectionMatchSetSqlInjectionMatchTuple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/waf/SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch"
)

type SqlInjectionMatchSetSqlInjectionMatchTuple struct {
	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch wafSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch `pulumi:"fieldToMatch"`
	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
	// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SqlInjectionMatchTuple.html#WAF-Type-SqlInjectionMatchTuple-TextTransformation)
	// for all supported values.
	TextTransformation string `pulumi:"textTransformation"`
}

type SqlInjectionMatchSetSqlInjectionMatchTupleInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetSqlInjectionMatchTupleOutput() SqlInjectionMatchSetSqlInjectionMatchTupleOutput
	ToSqlInjectionMatchSetSqlInjectionMatchTupleOutputWithContext(context.Context) SqlInjectionMatchSetSqlInjectionMatchTupleOutput
}

type SqlInjectionMatchSetSqlInjectionMatchTupleArgs struct {
	// Specifies where in a web request to look for snippets of malicious SQL code.
	FieldToMatch wafSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchInput `pulumi:"fieldToMatch"`
	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
	// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SqlInjectionMatchTuple.html#WAF-Type-SqlInjectionMatchTuple-TextTransformation)
	// for all supported values.
	TextTransformation pulumi.StringInput `pulumi:"textTransformation"`
}

func (SqlInjectionMatchSetSqlInjectionMatchTupleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlInjectionMatchSetSqlInjectionMatchTuple)(nil)).Elem()
}

func (i SqlInjectionMatchSetSqlInjectionMatchTupleArgs) ToSqlInjectionMatchSetSqlInjectionMatchTupleOutput() SqlInjectionMatchSetSqlInjectionMatchTupleOutput {
	return i.ToSqlInjectionMatchSetSqlInjectionMatchTupleOutputWithContext(context.Background())
}

func (i SqlInjectionMatchSetSqlInjectionMatchTupleArgs) ToSqlInjectionMatchSetSqlInjectionMatchTupleOutputWithContext(ctx context.Context) SqlInjectionMatchSetSqlInjectionMatchTupleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetSqlInjectionMatchTupleOutput)
}

type SqlInjectionMatchSetSqlInjectionMatchTupleArrayInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput() SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput
	ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutputWithContext(context.Context) SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput
}

type SqlInjectionMatchSetSqlInjectionMatchTupleArray []SqlInjectionMatchSetSqlInjectionMatchTupleInput

func (SqlInjectionMatchSetSqlInjectionMatchTupleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlInjectionMatchSetSqlInjectionMatchTuple)(nil)).Elem()
}

func (i SqlInjectionMatchSetSqlInjectionMatchTupleArray) ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput() SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput {
	return i.ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutputWithContext(context.Background())
}

func (i SqlInjectionMatchSetSqlInjectionMatchTupleArray) ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutputWithContext(ctx context.Context) SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput)
}

type SqlInjectionMatchSetSqlInjectionMatchTupleOutput struct { *pulumi.OutputState }

func (SqlInjectionMatchSetSqlInjectionMatchTupleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlInjectionMatchSetSqlInjectionMatchTuple)(nil)).Elem()
}

func (o SqlInjectionMatchSetSqlInjectionMatchTupleOutput) ToSqlInjectionMatchSetSqlInjectionMatchTupleOutput() SqlInjectionMatchSetSqlInjectionMatchTupleOutput {
	return o
}

func (o SqlInjectionMatchSetSqlInjectionMatchTupleOutput) ToSqlInjectionMatchSetSqlInjectionMatchTupleOutputWithContext(ctx context.Context) SqlInjectionMatchSetSqlInjectionMatchTupleOutput {
	return o
}

// Specifies where in a web request to look for snippets of malicious SQL code.
func (o SqlInjectionMatchSetSqlInjectionMatchTupleOutput) FieldToMatch() wafSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchOutput {
	return o.ApplyT(func (v SqlInjectionMatchSetSqlInjectionMatchTuple) wafSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch { return v.FieldToMatch }).(wafSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchOutput)
}

// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
// If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SqlInjectionMatchTuple.html#WAF-Type-SqlInjectionMatchTuple-TextTransformation)
// for all supported values.
func (o SqlInjectionMatchSetSqlInjectionMatchTupleOutput) TextTransformation() pulumi.StringOutput {
	return o.ApplyT(func (v SqlInjectionMatchSetSqlInjectionMatchTuple) string { return v.TextTransformation }).(pulumi.StringOutput)
}

type SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput struct { *pulumi.OutputState}

func (SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlInjectionMatchSetSqlInjectionMatchTuple)(nil)).Elem()
}

func (o SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput) ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput() SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput {
	return o
}

func (o SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput) ToSqlInjectionMatchSetSqlInjectionMatchTupleArrayOutputWithContext(ctx context.Context) SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput {
	return o
}

func (o SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput) Index(i pulumi.IntInput) SqlInjectionMatchSetSqlInjectionMatchTupleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) SqlInjectionMatchSetSqlInjectionMatchTuple {
		return vs[0].([]SqlInjectionMatchSetSqlInjectionMatchTuple)[vs[1].(int)]
	}).(SqlInjectionMatchSetSqlInjectionMatchTupleOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlInjectionMatchSetSqlInjectionMatchTupleOutput{})
	pulumi.RegisterOutputType(SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput{})
}
