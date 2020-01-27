// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetGroupUser struct {
	// The Amazon Resource Name (ARN) specifying the iam user.
	Arn string `pulumi:"arn"`
	// The path to the iam user.
	Path string `pulumi:"path"`
	// The stable and unique string identifying the iam user.
	UserId string `pulumi:"userId"`
	// The name of the iam user.
	UserName string `pulumi:"userName"`
}

type GetGroupUserInput interface {
	pulumi.Input

	ToGetGroupUserOutput() GetGroupUserOutput
	ToGetGroupUserOutputWithContext(context.Context) GetGroupUserOutput
}

type GetGroupUserArgs struct {
	// The Amazon Resource Name (ARN) specifying the iam user.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The path to the iam user.
	Path pulumi.StringInput `pulumi:"path"`
	// The stable and unique string identifying the iam user.
	UserId pulumi.StringInput `pulumi:"userId"`
	// The name of the iam user.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (GetGroupUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUser)(nil)).Elem()
}

func (i GetGroupUserArgs) ToGetGroupUserOutput() GetGroupUserOutput {
	return i.ToGetGroupUserOutputWithContext(context.Background())
}

func (i GetGroupUserArgs) ToGetGroupUserOutputWithContext(ctx context.Context) GetGroupUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupUserOutput)
}

type GetGroupUserArrayInput interface {
	pulumi.Input

	ToGetGroupUserArrayOutput() GetGroupUserArrayOutput
	ToGetGroupUserArrayOutputWithContext(context.Context) GetGroupUserArrayOutput
}

type GetGroupUserArray []GetGroupUserInput

func (GetGroupUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupUser)(nil)).Elem()
}

func (i GetGroupUserArray) ToGetGroupUserArrayOutput() GetGroupUserArrayOutput {
	return i.ToGetGroupUserArrayOutputWithContext(context.Background())
}

func (i GetGroupUserArray) ToGetGroupUserArrayOutputWithContext(ctx context.Context) GetGroupUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupUserArrayOutput)
}

type GetGroupUserOutput struct { *pulumi.OutputState }

func (GetGroupUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUser)(nil)).Elem()
}

func (o GetGroupUserOutput) ToGetGroupUserOutput() GetGroupUserOutput {
	return o
}

func (o GetGroupUserOutput) ToGetGroupUserOutputWithContext(ctx context.Context) GetGroupUserOutput {
	return o
}

// The Amazon Resource Name (ARN) specifying the iam user.
func (o GetGroupUserOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func (v GetGroupUser) string { return v.Arn }).(pulumi.StringOutput)
}

// The path to the iam user.
func (o GetGroupUserOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func (v GetGroupUser) string { return v.Path }).(pulumi.StringOutput)
}

// The stable and unique string identifying the iam user.
func (o GetGroupUserOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func (v GetGroupUser) string { return v.UserId }).(pulumi.StringOutput)
}

// The name of the iam user.
func (o GetGroupUserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func (v GetGroupUser) string { return v.UserName }).(pulumi.StringOutput)
}

type GetGroupUserArrayOutput struct { *pulumi.OutputState}

func (GetGroupUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupUser)(nil)).Elem()
}

func (o GetGroupUserArrayOutput) ToGetGroupUserArrayOutput() GetGroupUserArrayOutput {
	return o
}

func (o GetGroupUserArrayOutput) ToGetGroupUserArrayOutputWithContext(ctx context.Context) GetGroupUserArrayOutput {
	return o
}

func (o GetGroupUserArrayOutput) Index(i pulumi.IntInput) GetGroupUserOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetGroupUser {
		return vs[0].([]GetGroupUser)[vs[1].(int)]
	}).(GetGroupUserOutput)
}

type GetPolicyDocumentStatement struct {
	// A list of actions that this statement either allows
	// or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
	Actions []string `pulumi:"actions"`
	// A nested configuration block (described below)
	// that defines a further, possibly-service-specific condition that constrains
	// whether this statement applies.
	Conditions []GetPolicyDocumentStatementCondition `pulumi:"conditions"`
	// Either "Allow" or "Deny", to specify whether this
	// statement allows or denies the given actions. The default is "Allow".
	Effect *string `pulumi:"effect"`
	// A list of actions that this statement does *not*
	// apply to. Used to apply a policy statement to all actions *except* those
	// listed.
	NotActions []string `pulumi:"notActions"`
	// Like `principals` except gives resources that
	// the statement does *not* apply to.
	NotPrincipals []GetPolicyDocumentStatementNotPrincipal `pulumi:"notPrincipals"`
	// A list of resource ARNs that this statement
	// does *not* apply to. Used to apply a policy statement to all resources
	// *except* those listed.
	NotResources []string `pulumi:"notResources"`
	// A nested configuration block (described below)
	// specifying a resource (or resource pattern) to which this statement applies.
	Principals []GetPolicyDocumentStatementPrincipal `pulumi:"principals"`
	// A list of resource ARNs that this statement applies
	// to. This is required by AWS if used for an IAM policy.
	Resources []string `pulumi:"resources"`
	// An ID for the policy statement.
	Sid *string `pulumi:"sid"`
}

type GetPolicyDocumentStatementInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementOutput() GetPolicyDocumentStatementOutput
	ToGetPolicyDocumentStatementOutputWithContext(context.Context) GetPolicyDocumentStatementOutput
}

type GetPolicyDocumentStatementArgs struct {
	// A list of actions that this statement either allows
	// or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
	Actions pulumi.StringArrayInput `pulumi:"actions"`
	// A nested configuration block (described below)
	// that defines a further, possibly-service-specific condition that constrains
	// whether this statement applies.
	Conditions GetPolicyDocumentStatementConditionArrayInput `pulumi:"conditions"`
	// Either "Allow" or "Deny", to specify whether this
	// statement allows or denies the given actions. The default is "Allow".
	Effect pulumi.StringPtrInput `pulumi:"effect"`
	// A list of actions that this statement does *not*
	// apply to. Used to apply a policy statement to all actions *except* those
	// listed.
	NotActions pulumi.StringArrayInput `pulumi:"notActions"`
	// Like `principals` except gives resources that
	// the statement does *not* apply to.
	NotPrincipals GetPolicyDocumentStatementNotPrincipalArrayInput `pulumi:"notPrincipals"`
	// A list of resource ARNs that this statement
	// does *not* apply to. Used to apply a policy statement to all resources
	// *except* those listed.
	NotResources pulumi.StringArrayInput `pulumi:"notResources"`
	// A nested configuration block (described below)
	// specifying a resource (or resource pattern) to which this statement applies.
	Principals GetPolicyDocumentStatementPrincipalArrayInput `pulumi:"principals"`
	// A list of resource ARNs that this statement applies
	// to. This is required by AWS if used for an IAM policy.
	Resources pulumi.StringArrayInput `pulumi:"resources"`
	// An ID for the policy statement.
	Sid pulumi.StringPtrInput `pulumi:"sid"`
}

func (GetPolicyDocumentStatementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatement)(nil)).Elem()
}

func (i GetPolicyDocumentStatementArgs) ToGetPolicyDocumentStatementOutput() GetPolicyDocumentStatementOutput {
	return i.ToGetPolicyDocumentStatementOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementArgs) ToGetPolicyDocumentStatementOutputWithContext(ctx context.Context) GetPolicyDocumentStatementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementOutput)
}

type GetPolicyDocumentStatementArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementArrayOutput() GetPolicyDocumentStatementArrayOutput
	ToGetPolicyDocumentStatementArrayOutputWithContext(context.Context) GetPolicyDocumentStatementArrayOutput
}

type GetPolicyDocumentStatementArray []GetPolicyDocumentStatementInput

func (GetPolicyDocumentStatementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatement)(nil)).Elem()
}

func (i GetPolicyDocumentStatementArray) ToGetPolicyDocumentStatementArrayOutput() GetPolicyDocumentStatementArrayOutput {
	return i.ToGetPolicyDocumentStatementArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementArray) ToGetPolicyDocumentStatementArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementArrayOutput)
}

type GetPolicyDocumentStatementOutput struct { *pulumi.OutputState }

func (GetPolicyDocumentStatementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatement)(nil)).Elem()
}

func (o GetPolicyDocumentStatementOutput) ToGetPolicyDocumentStatementOutput() GetPolicyDocumentStatementOutput {
	return o
}

func (o GetPolicyDocumentStatementOutput) ToGetPolicyDocumentStatementOutputWithContext(ctx context.Context) GetPolicyDocumentStatementOutput {
	return o
}

// A list of actions that this statement either allows
// or denies. For example, ``["ec2:RunInstances", "s3:*"]``.
func (o GetPolicyDocumentStatementOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

// A nested configuration block (described below)
// that defines a further, possibly-service-specific condition that constrains
// whether this statement applies.
func (o GetPolicyDocumentStatementOutput) Conditions() GetPolicyDocumentStatementConditionArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []GetPolicyDocumentStatementCondition { return v.Conditions }).(GetPolicyDocumentStatementConditionArrayOutput)
}

// Either "Allow" or "Deny", to specify whether this
// statement allows or denies the given actions. The default is "Allow".
func (o GetPolicyDocumentStatementOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) *string { return v.Effect }).(pulumi.StringPtrOutput)
}

// A list of actions that this statement does *not*
// apply to. Used to apply a policy statement to all actions *except* those
// listed.
func (o GetPolicyDocumentStatementOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

// Like `principals` except gives resources that
// the statement does *not* apply to.
func (o GetPolicyDocumentStatementOutput) NotPrincipals() GetPolicyDocumentStatementNotPrincipalArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []GetPolicyDocumentStatementNotPrincipal { return v.NotPrincipals }).(GetPolicyDocumentStatementNotPrincipalArrayOutput)
}

// A list of resource ARNs that this statement
// does *not* apply to. Used to apply a policy statement to all resources
// *except* those listed.
func (o GetPolicyDocumentStatementOutput) NotResources() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []string { return v.NotResources }).(pulumi.StringArrayOutput)
}

// A nested configuration block (described below)
// specifying a resource (or resource pattern) to which this statement applies.
func (o GetPolicyDocumentStatementOutput) Principals() GetPolicyDocumentStatementPrincipalArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []GetPolicyDocumentStatementPrincipal { return v.Principals }).(GetPolicyDocumentStatementPrincipalArrayOutput)
}

// A list of resource ARNs that this statement applies
// to. This is required by AWS if used for an IAM policy.
func (o GetPolicyDocumentStatementOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

// An ID for the policy statement.
func (o GetPolicyDocumentStatementOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatement) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

type GetPolicyDocumentStatementArrayOutput struct { *pulumi.OutputState}

func (GetPolicyDocumentStatementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatement)(nil)).Elem()
}

func (o GetPolicyDocumentStatementArrayOutput) ToGetPolicyDocumentStatementArrayOutput() GetPolicyDocumentStatementArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementArrayOutput) ToGetPolicyDocumentStatementArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentStatementOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetPolicyDocumentStatement {
		return vs[0].([]GetPolicyDocumentStatement)[vs[1].(int)]
	}).(GetPolicyDocumentStatementOutput)
}

type GetPolicyDocumentStatementCondition struct {
	// The name of the
	// [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
	// to evaluate.
	Test string `pulumi:"test"`
	// The values to evaluate the condition against. If multiple
	// values are provided, the condition matches if at least one of them applies.
	// (That is, the tests are combined with the "OR" boolean operation.)
	Values []string `pulumi:"values"`
	// The name of a
	// [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
	// to apply the condition to. Context variables may either be standard AWS
	// variables starting with `aws:`, or service-specific variables prefixed with
	// the service name.
	Variable string `pulumi:"variable"`
}

type GetPolicyDocumentStatementConditionInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementConditionOutput() GetPolicyDocumentStatementConditionOutput
	ToGetPolicyDocumentStatementConditionOutputWithContext(context.Context) GetPolicyDocumentStatementConditionOutput
}

type GetPolicyDocumentStatementConditionArgs struct {
	// The name of the
	// [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
	// to evaluate.
	Test pulumi.StringInput `pulumi:"test"`
	// The values to evaluate the condition against. If multiple
	// values are provided, the condition matches if at least one of them applies.
	// (That is, the tests are combined with the "OR" boolean operation.)
	Values pulumi.StringArrayInput `pulumi:"values"`
	// The name of a
	// [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
	// to apply the condition to. Context variables may either be standard AWS
	// variables starting with `aws:`, or service-specific variables prefixed with
	// the service name.
	Variable pulumi.StringInput `pulumi:"variable"`
}

func (GetPolicyDocumentStatementConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatementCondition)(nil)).Elem()
}

func (i GetPolicyDocumentStatementConditionArgs) ToGetPolicyDocumentStatementConditionOutput() GetPolicyDocumentStatementConditionOutput {
	return i.ToGetPolicyDocumentStatementConditionOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementConditionArgs) ToGetPolicyDocumentStatementConditionOutputWithContext(ctx context.Context) GetPolicyDocumentStatementConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementConditionOutput)
}

type GetPolicyDocumentStatementConditionArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementConditionArrayOutput() GetPolicyDocumentStatementConditionArrayOutput
	ToGetPolicyDocumentStatementConditionArrayOutputWithContext(context.Context) GetPolicyDocumentStatementConditionArrayOutput
}

type GetPolicyDocumentStatementConditionArray []GetPolicyDocumentStatementConditionInput

func (GetPolicyDocumentStatementConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatementCondition)(nil)).Elem()
}

func (i GetPolicyDocumentStatementConditionArray) ToGetPolicyDocumentStatementConditionArrayOutput() GetPolicyDocumentStatementConditionArrayOutput {
	return i.ToGetPolicyDocumentStatementConditionArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementConditionArray) ToGetPolicyDocumentStatementConditionArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementConditionArrayOutput)
}

type GetPolicyDocumentStatementConditionOutput struct { *pulumi.OutputState }

func (GetPolicyDocumentStatementConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatementCondition)(nil)).Elem()
}

func (o GetPolicyDocumentStatementConditionOutput) ToGetPolicyDocumentStatementConditionOutput() GetPolicyDocumentStatementConditionOutput {
	return o
}

func (o GetPolicyDocumentStatementConditionOutput) ToGetPolicyDocumentStatementConditionOutputWithContext(ctx context.Context) GetPolicyDocumentStatementConditionOutput {
	return o
}

// The name of the
// [IAM condition operator](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition_operators.html)
// to evaluate.
func (o GetPolicyDocumentStatementConditionOutput) Test() pulumi.StringOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementCondition) string { return v.Test }).(pulumi.StringOutput)
}

// The values to evaluate the condition against. If multiple
// values are provided, the condition matches if at least one of them applies.
// (That is, the tests are combined with the "OR" boolean operation.)
func (o GetPolicyDocumentStatementConditionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementCondition) []string { return v.Values }).(pulumi.StringArrayOutput)
}

// The name of a
// [Context Variable](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html#AvailableKeys)
// to apply the condition to. Context variables may either be standard AWS
// variables starting with `aws:`, or service-specific variables prefixed with
// the service name.
func (o GetPolicyDocumentStatementConditionOutput) Variable() pulumi.StringOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementCondition) string { return v.Variable }).(pulumi.StringOutput)
}

type GetPolicyDocumentStatementConditionArrayOutput struct { *pulumi.OutputState}

func (GetPolicyDocumentStatementConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatementCondition)(nil)).Elem()
}

func (o GetPolicyDocumentStatementConditionArrayOutput) ToGetPolicyDocumentStatementConditionArrayOutput() GetPolicyDocumentStatementConditionArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementConditionArrayOutput) ToGetPolicyDocumentStatementConditionArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementConditionArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementConditionArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentStatementConditionOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetPolicyDocumentStatementCondition {
		return vs[0].([]GetPolicyDocumentStatementCondition)[vs[1].(int)]
	}).(GetPolicyDocumentStatementConditionOutput)
}

type GetPolicyDocumentStatementNotPrincipal struct {
	// List of identifiers for principals. When `type`
	// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
	Identifiers []string `pulumi:"identifiers"`
	// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
	Type string `pulumi:"type"`
}

type GetPolicyDocumentStatementNotPrincipalInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementNotPrincipalOutput() GetPolicyDocumentStatementNotPrincipalOutput
	ToGetPolicyDocumentStatementNotPrincipalOutputWithContext(context.Context) GetPolicyDocumentStatementNotPrincipalOutput
}

type GetPolicyDocumentStatementNotPrincipalArgs struct {
	// List of identifiers for principals. When `type`
	// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
	Identifiers pulumi.StringArrayInput `pulumi:"identifiers"`
	// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetPolicyDocumentStatementNotPrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatementNotPrincipal)(nil)).Elem()
}

func (i GetPolicyDocumentStatementNotPrincipalArgs) ToGetPolicyDocumentStatementNotPrincipalOutput() GetPolicyDocumentStatementNotPrincipalOutput {
	return i.ToGetPolicyDocumentStatementNotPrincipalOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementNotPrincipalArgs) ToGetPolicyDocumentStatementNotPrincipalOutputWithContext(ctx context.Context) GetPolicyDocumentStatementNotPrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementNotPrincipalOutput)
}

type GetPolicyDocumentStatementNotPrincipalArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementNotPrincipalArrayOutput() GetPolicyDocumentStatementNotPrincipalArrayOutput
	ToGetPolicyDocumentStatementNotPrincipalArrayOutputWithContext(context.Context) GetPolicyDocumentStatementNotPrincipalArrayOutput
}

type GetPolicyDocumentStatementNotPrincipalArray []GetPolicyDocumentStatementNotPrincipalInput

func (GetPolicyDocumentStatementNotPrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatementNotPrincipal)(nil)).Elem()
}

func (i GetPolicyDocumentStatementNotPrincipalArray) ToGetPolicyDocumentStatementNotPrincipalArrayOutput() GetPolicyDocumentStatementNotPrincipalArrayOutput {
	return i.ToGetPolicyDocumentStatementNotPrincipalArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementNotPrincipalArray) ToGetPolicyDocumentStatementNotPrincipalArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementNotPrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementNotPrincipalArrayOutput)
}

type GetPolicyDocumentStatementNotPrincipalOutput struct { *pulumi.OutputState }

func (GetPolicyDocumentStatementNotPrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatementNotPrincipal)(nil)).Elem()
}

func (o GetPolicyDocumentStatementNotPrincipalOutput) ToGetPolicyDocumentStatementNotPrincipalOutput() GetPolicyDocumentStatementNotPrincipalOutput {
	return o
}

func (o GetPolicyDocumentStatementNotPrincipalOutput) ToGetPolicyDocumentStatementNotPrincipalOutputWithContext(ctx context.Context) GetPolicyDocumentStatementNotPrincipalOutput {
	return o
}

// List of identifiers for principals. When `type`
// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
func (o GetPolicyDocumentStatementNotPrincipalOutput) Identifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementNotPrincipal) []string { return v.Identifiers }).(pulumi.StringArrayOutput)
}

// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
func (o GetPolicyDocumentStatementNotPrincipalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementNotPrincipal) string { return v.Type }).(pulumi.StringOutput)
}

type GetPolicyDocumentStatementNotPrincipalArrayOutput struct { *pulumi.OutputState}

func (GetPolicyDocumentStatementNotPrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatementNotPrincipal)(nil)).Elem()
}

func (o GetPolicyDocumentStatementNotPrincipalArrayOutput) ToGetPolicyDocumentStatementNotPrincipalArrayOutput() GetPolicyDocumentStatementNotPrincipalArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementNotPrincipalArrayOutput) ToGetPolicyDocumentStatementNotPrincipalArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementNotPrincipalArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementNotPrincipalArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentStatementNotPrincipalOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetPolicyDocumentStatementNotPrincipal {
		return vs[0].([]GetPolicyDocumentStatementNotPrincipal)[vs[1].(int)]
	}).(GetPolicyDocumentStatementNotPrincipalOutput)
}

type GetPolicyDocumentStatementPrincipal struct {
	// List of identifiers for principals. When `type`
	// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
	Identifiers []string `pulumi:"identifiers"`
	// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
	Type string `pulumi:"type"`
}

type GetPolicyDocumentStatementPrincipalInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementPrincipalOutput() GetPolicyDocumentStatementPrincipalOutput
	ToGetPolicyDocumentStatementPrincipalOutputWithContext(context.Context) GetPolicyDocumentStatementPrincipalOutput
}

type GetPolicyDocumentStatementPrincipalArgs struct {
	// List of identifiers for principals. When `type`
	// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
	Identifiers pulumi.StringArrayInput `pulumi:"identifiers"`
	// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetPolicyDocumentStatementPrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatementPrincipal)(nil)).Elem()
}

func (i GetPolicyDocumentStatementPrincipalArgs) ToGetPolicyDocumentStatementPrincipalOutput() GetPolicyDocumentStatementPrincipalOutput {
	return i.ToGetPolicyDocumentStatementPrincipalOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementPrincipalArgs) ToGetPolicyDocumentStatementPrincipalOutputWithContext(ctx context.Context) GetPolicyDocumentStatementPrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementPrincipalOutput)
}

type GetPolicyDocumentStatementPrincipalArrayInput interface {
	pulumi.Input

	ToGetPolicyDocumentStatementPrincipalArrayOutput() GetPolicyDocumentStatementPrincipalArrayOutput
	ToGetPolicyDocumentStatementPrincipalArrayOutputWithContext(context.Context) GetPolicyDocumentStatementPrincipalArrayOutput
}

type GetPolicyDocumentStatementPrincipalArray []GetPolicyDocumentStatementPrincipalInput

func (GetPolicyDocumentStatementPrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatementPrincipal)(nil)).Elem()
}

func (i GetPolicyDocumentStatementPrincipalArray) ToGetPolicyDocumentStatementPrincipalArrayOutput() GetPolicyDocumentStatementPrincipalArrayOutput {
	return i.ToGetPolicyDocumentStatementPrincipalArrayOutputWithContext(context.Background())
}

func (i GetPolicyDocumentStatementPrincipalArray) ToGetPolicyDocumentStatementPrincipalArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementPrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPolicyDocumentStatementPrincipalArrayOutput)
}

type GetPolicyDocumentStatementPrincipalOutput struct { *pulumi.OutputState }

func (GetPolicyDocumentStatementPrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyDocumentStatementPrincipal)(nil)).Elem()
}

func (o GetPolicyDocumentStatementPrincipalOutput) ToGetPolicyDocumentStatementPrincipalOutput() GetPolicyDocumentStatementPrincipalOutput {
	return o
}

func (o GetPolicyDocumentStatementPrincipalOutput) ToGetPolicyDocumentStatementPrincipalOutputWithContext(ctx context.Context) GetPolicyDocumentStatementPrincipalOutput {
	return o
}

// List of identifiers for principals. When `type`
// is "AWS", these are IAM user or role ARNs.  When `type` is "Service", these are AWS Service roles e.g. `lambda.amazonaws.com`.
func (o GetPolicyDocumentStatementPrincipalOutput) Identifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementPrincipal) []string { return v.Identifiers }).(pulumi.StringArrayOutput)
}

// The type of principal. For AWS ARNs this is "AWS".  For AWS services (e.g. Lambda), this is "Service".
func (o GetPolicyDocumentStatementPrincipalOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v GetPolicyDocumentStatementPrincipal) string { return v.Type }).(pulumi.StringOutput)
}

type GetPolicyDocumentStatementPrincipalArrayOutput struct { *pulumi.OutputState}

func (GetPolicyDocumentStatementPrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetPolicyDocumentStatementPrincipal)(nil)).Elem()
}

func (o GetPolicyDocumentStatementPrincipalArrayOutput) ToGetPolicyDocumentStatementPrincipalArrayOutput() GetPolicyDocumentStatementPrincipalArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementPrincipalArrayOutput) ToGetPolicyDocumentStatementPrincipalArrayOutputWithContext(ctx context.Context) GetPolicyDocumentStatementPrincipalArrayOutput {
	return o
}

func (o GetPolicyDocumentStatementPrincipalArrayOutput) Index(i pulumi.IntInput) GetPolicyDocumentStatementPrincipalOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetPolicyDocumentStatementPrincipal {
		return vs[0].([]GetPolicyDocumentStatementPrincipal)[vs[1].(int)]
	}).(GetPolicyDocumentStatementPrincipalOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupUserOutput{})
	pulumi.RegisterOutputType(GetGroupUserArrayOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementArrayOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementConditionOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementConditionArrayOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementNotPrincipalOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementNotPrincipalArrayOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementPrincipalOutput{})
	pulumi.RegisterOutputType(GetPolicyDocumentStatementPrincipalArrayOutput{})
}
