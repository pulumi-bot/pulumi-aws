// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF Regional SQL Injection Match Set Resource for use with Application Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = wafregional.NewSqlInjectionMatchSet(ctx, "sqlInjectionMatchSet", &wafregional.SqlInjectionMatchSetArgs{
// 			SqlInjectionMatchTuples: wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleArray{
// 				&wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleArgs{
// 					FieldToMatch: &wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs{
// 						Type: pulumi.String("QUERY_STRING"),
// 					},
// 					TextTransformation: pulumi.String("URL_DECODE"),
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
type SqlInjectionMatchSet struct {
	pulumi.CustomResourceState

	// The name or description of the SizeConstraintSet.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayOutput `pulumi:"sqlInjectionMatchTuples"`
}

// NewSqlInjectionMatchSet registers a new resource with the given unique name, arguments, and options.
func NewSqlInjectionMatchSet(ctx *pulumi.Context,
	name string, args *SqlInjectionMatchSetArgs, opts ...pulumi.ResourceOption) (*SqlInjectionMatchSet, error) {
	if args == nil {
		args = &SqlInjectionMatchSetArgs{}
	}
	var resource SqlInjectionMatchSet
	err := ctx.RegisterResource("aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlInjectionMatchSet gets an existing SqlInjectionMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlInjectionMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlInjectionMatchSetState, opts ...pulumi.ResourceOption) (*SqlInjectionMatchSet, error) {
	var resource SqlInjectionMatchSet
	err := ctx.ReadResource("aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlInjectionMatchSet resources.
type sqlInjectionMatchSetState struct {
	// The name or description of the SizeConstraintSet.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples []SqlInjectionMatchSetSqlInjectionMatchTuple `pulumi:"sqlInjectionMatchTuples"`
}

type SqlInjectionMatchSetState struct {
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayInput
}

func (SqlInjectionMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlInjectionMatchSetState)(nil)).Elem()
}

type sqlInjectionMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples []SqlInjectionMatchSetSqlInjectionMatchTuple `pulumi:"sqlInjectionMatchTuples"`
}

// The set of arguments for constructing a SqlInjectionMatchSet resource.
type SqlInjectionMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayInput
}

func (SqlInjectionMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlInjectionMatchSetArgs)(nil)).Elem()
}
