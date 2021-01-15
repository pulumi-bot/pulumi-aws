// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a WAF SQL Injection Match Set Resource
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
// 		_, err := waf.NewSqlInjectionMatchSet(ctx, "sqlInjectionMatchSet", &waf.SqlInjectionMatchSetArgs{
// 			SqlInjectionMatchTuples: waf.SqlInjectionMatchSetSqlInjectionMatchTupleArray{
// 				&waf.SqlInjectionMatchSetSqlInjectionMatchTupleArgs{
// 					FieldToMatch: &waf.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs{
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
//
// ## Import
//
// AWS WAF SQL Injection Match Set can be imported using their ID, e.g.
//
// ```sh
//  $ pulumi import aws:waf/sqlInjectionMatchSet:SqlInjectionMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type SqlInjectionMatchSet struct {
	pulumi.CustomResourceState

	// The name or description of the SQL Injection Match Set.
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
	err := ctx.RegisterResource("aws:waf/sqlInjectionMatchSet:SqlInjectionMatchSet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:waf/sqlInjectionMatchSet:SqlInjectionMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlInjectionMatchSet resources.
type sqlInjectionMatchSetState struct {
	// The name or description of the SQL Injection Match Set.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples []SqlInjectionMatchSetSqlInjectionMatchTuple `pulumi:"sqlInjectionMatchTuples"`
}

type SqlInjectionMatchSetState struct {
	// The name or description of the SQL Injection Match Set.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayInput
}

func (SqlInjectionMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlInjectionMatchSetState)(nil)).Elem()
}

type sqlInjectionMatchSetArgs struct {
	// The name or description of the SQL Injection Match Set.
	Name *string `pulumi:"name"`
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples []SqlInjectionMatchSetSqlInjectionMatchTuple `pulumi:"sqlInjectionMatchTuples"`
}

// The set of arguments for constructing a SqlInjectionMatchSet resource.
type SqlInjectionMatchSetArgs struct {
	// The name or description of the SQL Injection Match Set.
	Name pulumi.StringPtrInput
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples SqlInjectionMatchSetSqlInjectionMatchTupleArrayInput
}

func (SqlInjectionMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlInjectionMatchSetArgs)(nil)).Elem()
}

type SqlInjectionMatchSetInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetOutput() SqlInjectionMatchSetOutput
	ToSqlInjectionMatchSetOutputWithContext(ctx context.Context) SqlInjectionMatchSetOutput
}

func (*SqlInjectionMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlInjectionMatchSet)(nil))
}

func (i *SqlInjectionMatchSet) ToSqlInjectionMatchSetOutput() SqlInjectionMatchSetOutput {
	return i.ToSqlInjectionMatchSetOutputWithContext(context.Background())
}

func (i *SqlInjectionMatchSet) ToSqlInjectionMatchSetOutputWithContext(ctx context.Context) SqlInjectionMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetOutput)
}

func (i *SqlInjectionMatchSet) ToSqlInjectionMatchSetPtrOutput() SqlInjectionMatchSetPtrOutput {
	return i.ToSqlInjectionMatchSetPtrOutputWithContext(context.Background())
}

func (i *SqlInjectionMatchSet) ToSqlInjectionMatchSetPtrOutputWithContext(ctx context.Context) SqlInjectionMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetPtrOutput)
}

type SqlInjectionMatchSetPtrInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetPtrOutput() SqlInjectionMatchSetPtrOutput
	ToSqlInjectionMatchSetPtrOutputWithContext(ctx context.Context) SqlInjectionMatchSetPtrOutput
}

type sqlInjectionMatchSetPtrType SqlInjectionMatchSetArgs

func (*sqlInjectionMatchSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlInjectionMatchSet)(nil))
}

func (i *sqlInjectionMatchSetPtrType) ToSqlInjectionMatchSetPtrOutput() SqlInjectionMatchSetPtrOutput {
	return i.ToSqlInjectionMatchSetPtrOutputWithContext(context.Background())
}

func (i *sqlInjectionMatchSetPtrType) ToSqlInjectionMatchSetPtrOutputWithContext(ctx context.Context) SqlInjectionMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetPtrOutput)
}

// SqlInjectionMatchSetArrayInput is an input type that accepts SqlInjectionMatchSetArray and SqlInjectionMatchSetArrayOutput values.
// You can construct a concrete instance of `SqlInjectionMatchSetArrayInput` via:
//
//          SqlInjectionMatchSetArray{ SqlInjectionMatchSetArgs{...} }
type SqlInjectionMatchSetArrayInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetArrayOutput() SqlInjectionMatchSetArrayOutput
	ToSqlInjectionMatchSetArrayOutputWithContext(context.Context) SqlInjectionMatchSetArrayOutput
}

type SqlInjectionMatchSetArray []SqlInjectionMatchSetInput

func (SqlInjectionMatchSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SqlInjectionMatchSet)(nil))
}

func (i SqlInjectionMatchSetArray) ToSqlInjectionMatchSetArrayOutput() SqlInjectionMatchSetArrayOutput {
	return i.ToSqlInjectionMatchSetArrayOutputWithContext(context.Background())
}

func (i SqlInjectionMatchSetArray) ToSqlInjectionMatchSetArrayOutputWithContext(ctx context.Context) SqlInjectionMatchSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetArrayOutput)
}

// SqlInjectionMatchSetMapInput is an input type that accepts SqlInjectionMatchSetMap and SqlInjectionMatchSetMapOutput values.
// You can construct a concrete instance of `SqlInjectionMatchSetMapInput` via:
//
//          SqlInjectionMatchSetMap{ "key": SqlInjectionMatchSetArgs{...} }
type SqlInjectionMatchSetMapInput interface {
	pulumi.Input

	ToSqlInjectionMatchSetMapOutput() SqlInjectionMatchSetMapOutput
	ToSqlInjectionMatchSetMapOutputWithContext(context.Context) SqlInjectionMatchSetMapOutput
}

type SqlInjectionMatchSetMap map[string]SqlInjectionMatchSetInput

func (SqlInjectionMatchSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SqlInjectionMatchSet)(nil))
}

func (i SqlInjectionMatchSetMap) ToSqlInjectionMatchSetMapOutput() SqlInjectionMatchSetMapOutput {
	return i.ToSqlInjectionMatchSetMapOutputWithContext(context.Background())
}

func (i SqlInjectionMatchSetMap) ToSqlInjectionMatchSetMapOutputWithContext(ctx context.Context) SqlInjectionMatchSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlInjectionMatchSetMapOutput)
}

type SqlInjectionMatchSetOutput struct {
	*pulumi.OutputState
}

func (SqlInjectionMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlInjectionMatchSet)(nil))
}

func (o SqlInjectionMatchSetOutput) ToSqlInjectionMatchSetOutput() SqlInjectionMatchSetOutput {
	return o
}

func (o SqlInjectionMatchSetOutput) ToSqlInjectionMatchSetOutputWithContext(ctx context.Context) SqlInjectionMatchSetOutput {
	return o
}

func (o SqlInjectionMatchSetOutput) ToSqlInjectionMatchSetPtrOutput() SqlInjectionMatchSetPtrOutput {
	return o.ToSqlInjectionMatchSetPtrOutputWithContext(context.Background())
}

func (o SqlInjectionMatchSetOutput) ToSqlInjectionMatchSetPtrOutputWithContext(ctx context.Context) SqlInjectionMatchSetPtrOutput {
	return o.ApplyT(func(v SqlInjectionMatchSet) *SqlInjectionMatchSet {
		return &v
	}).(SqlInjectionMatchSetPtrOutput)
}

type SqlInjectionMatchSetPtrOutput struct {
	*pulumi.OutputState
}

func (SqlInjectionMatchSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlInjectionMatchSet)(nil))
}

func (o SqlInjectionMatchSetPtrOutput) ToSqlInjectionMatchSetPtrOutput() SqlInjectionMatchSetPtrOutput {
	return o
}

func (o SqlInjectionMatchSetPtrOutput) ToSqlInjectionMatchSetPtrOutputWithContext(ctx context.Context) SqlInjectionMatchSetPtrOutput {
	return o
}

type SqlInjectionMatchSetArrayOutput struct{ *pulumi.OutputState }

func (SqlInjectionMatchSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SqlInjectionMatchSet)(nil))
}

func (o SqlInjectionMatchSetArrayOutput) ToSqlInjectionMatchSetArrayOutput() SqlInjectionMatchSetArrayOutput {
	return o
}

func (o SqlInjectionMatchSetArrayOutput) ToSqlInjectionMatchSetArrayOutputWithContext(ctx context.Context) SqlInjectionMatchSetArrayOutput {
	return o
}

func (o SqlInjectionMatchSetArrayOutput) Index(i pulumi.IntInput) SqlInjectionMatchSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SqlInjectionMatchSet {
		return vs[0].([]SqlInjectionMatchSet)[vs[1].(int)]
	}).(SqlInjectionMatchSetOutput)
}

type SqlInjectionMatchSetMapOutput struct{ *pulumi.OutputState }

func (SqlInjectionMatchSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SqlInjectionMatchSet)(nil))
}

func (o SqlInjectionMatchSetMapOutput) ToSqlInjectionMatchSetMapOutput() SqlInjectionMatchSetMapOutput {
	return o
}

func (o SqlInjectionMatchSetMapOutput) ToSqlInjectionMatchSetMapOutputWithContext(ctx context.Context) SqlInjectionMatchSetMapOutput {
	return o
}

func (o SqlInjectionMatchSetMapOutput) MapIndex(k pulumi.StringInput) SqlInjectionMatchSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SqlInjectionMatchSet {
		return vs[0].(map[string]SqlInjectionMatchSet)[vs[1].(string)]
	}).(SqlInjectionMatchSetOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlInjectionMatchSetOutput{})
	pulumi.RegisterOutputType(SqlInjectionMatchSetPtrOutput{})
	pulumi.RegisterOutputType(SqlInjectionMatchSetArrayOutput{})
	pulumi.RegisterOutputType(SqlInjectionMatchSetMapOutput{})
}
