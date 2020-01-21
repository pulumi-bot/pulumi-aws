// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ListenerRuleAction

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/alb/ListenerRuleActionAuthenticateCognito"
	"https:/github.com/pulumi/pulumi-aws/alb/ListenerRuleActionAuthenticateOidc"
	"https:/github.com/pulumi/pulumi-aws/alb/ListenerRuleActionFixedResponse"
	"https:/github.com/pulumi/pulumi-aws/alb/ListenerRuleActionRedirect"
)

type ListenerRuleAction struct {
	// Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
	AuthenticateCognito *albListenerRuleActionAuthenticateCognito.ListenerRuleActionAuthenticateCognito `pulumi:"authenticateCognito"`
	// Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
	AuthenticateOidc *albListenerRuleActionAuthenticateOidc.ListenerRuleActionAuthenticateOidc `pulumi:"authenticateOidc"`
	// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
	FixedResponse *albListenerRuleActionFixedResponse.ListenerRuleActionFixedResponse `pulumi:"fixedResponse"`
	Order *int `pulumi:"order"`
	// Information for creating a redirect action. Required if `type` is `redirect`.
	Redirect *albListenerRuleActionRedirect.ListenerRuleActionRedirect `pulumi:"redirect"`
	// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
	TargetGroupArn *string `pulumi:"targetGroupArn"`
	// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
	Type string `pulumi:"type"`
}

type ListenerRuleActionInput interface {
	pulumi.Input

	ToListenerRuleActionOutput() ListenerRuleActionOutput
	ToListenerRuleActionOutputWithContext(context.Context) ListenerRuleActionOutput
}

type ListenerRuleActionArgs struct {
	// Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
	AuthenticateCognito albListenerRuleActionAuthenticateCognito.ListenerRuleActionAuthenticateCognitoPtrInput `pulumi:"authenticateCognito"`
	// Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
	AuthenticateOidc albListenerRuleActionAuthenticateOidc.ListenerRuleActionAuthenticateOidcPtrInput `pulumi:"authenticateOidc"`
	// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
	FixedResponse albListenerRuleActionFixedResponse.ListenerRuleActionFixedResponsePtrInput `pulumi:"fixedResponse"`
	Order pulumi.IntPtrInput `pulumi:"order"`
	// Information for creating a redirect action. Required if `type` is `redirect`.
	Redirect albListenerRuleActionRedirect.ListenerRuleActionRedirectPtrInput `pulumi:"redirect"`
	// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
	TargetGroupArn pulumi.StringPtrInput `pulumi:"targetGroupArn"`
	// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ListenerRuleActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleAction)(nil)).Elem()
}

func (i ListenerRuleActionArgs) ToListenerRuleActionOutput() ListenerRuleActionOutput {
	return i.ToListenerRuleActionOutputWithContext(context.Background())
}

func (i ListenerRuleActionArgs) ToListenerRuleActionOutputWithContext(ctx context.Context) ListenerRuleActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleActionOutput)
}

type ListenerRuleActionArrayInput interface {
	pulumi.Input

	ToListenerRuleActionArrayOutput() ListenerRuleActionArrayOutput
	ToListenerRuleActionArrayOutputWithContext(context.Context) ListenerRuleActionArrayOutput
}

type ListenerRuleActionArray []ListenerRuleActionInput

func (ListenerRuleActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerRuleAction)(nil)).Elem()
}

func (i ListenerRuleActionArray) ToListenerRuleActionArrayOutput() ListenerRuleActionArrayOutput {
	return i.ToListenerRuleActionArrayOutputWithContext(context.Background())
}

func (i ListenerRuleActionArray) ToListenerRuleActionArrayOutputWithContext(ctx context.Context) ListenerRuleActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerRuleActionArrayOutput)
}

type ListenerRuleActionOutput struct { *pulumi.OutputState }

func (ListenerRuleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerRuleAction)(nil)).Elem()
}

func (o ListenerRuleActionOutput) ToListenerRuleActionOutput() ListenerRuleActionOutput {
	return o
}

func (o ListenerRuleActionOutput) ToListenerRuleActionOutputWithContext(ctx context.Context) ListenerRuleActionOutput {
	return o
}

// Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
func (o ListenerRuleActionOutput) AuthenticateCognito() albListenerRuleActionAuthenticateCognito.ListenerRuleActionAuthenticateCognitoPtrOutput {
	return o.ApplyT(func (v ListenerRuleAction) *albListenerRuleActionAuthenticateCognito.ListenerRuleActionAuthenticateCognito { return v.AuthenticateCognito }).(albListenerRuleActionAuthenticateCognito.ListenerRuleActionAuthenticateCognitoPtrOutput)
}

// Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
func (o ListenerRuleActionOutput) AuthenticateOidc() albListenerRuleActionAuthenticateOidc.ListenerRuleActionAuthenticateOidcPtrOutput {
	return o.ApplyT(func (v ListenerRuleAction) *albListenerRuleActionAuthenticateOidc.ListenerRuleActionAuthenticateOidc { return v.AuthenticateOidc }).(albListenerRuleActionAuthenticateOidc.ListenerRuleActionAuthenticateOidcPtrOutput)
}

// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
func (o ListenerRuleActionOutput) FixedResponse() albListenerRuleActionFixedResponse.ListenerRuleActionFixedResponsePtrOutput {
	return o.ApplyT(func (v ListenerRuleAction) *albListenerRuleActionFixedResponse.ListenerRuleActionFixedResponse { return v.FixedResponse }).(albListenerRuleActionFixedResponse.ListenerRuleActionFixedResponsePtrOutput)
}

func (o ListenerRuleActionOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ListenerRuleAction) *int { return v.Order }).(pulumi.IntPtrOutput)
}

// Information for creating a redirect action. Required if `type` is `redirect`.
func (o ListenerRuleActionOutput) Redirect() albListenerRuleActionRedirect.ListenerRuleActionRedirectPtrOutput {
	return o.ApplyT(func (v ListenerRuleAction) *albListenerRuleActionRedirect.ListenerRuleActionRedirect { return v.Redirect }).(albListenerRuleActionRedirect.ListenerRuleActionRedirectPtrOutput)
}

// The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
func (o ListenerRuleActionOutput) TargetGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ListenerRuleAction) *string { return v.TargetGroupArn }).(pulumi.StringPtrOutput)
}

// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
func (o ListenerRuleActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v ListenerRuleAction) string { return v.Type }).(pulumi.StringOutput)
}

type ListenerRuleActionArrayOutput struct { *pulumi.OutputState}

func (ListenerRuleActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerRuleAction)(nil)).Elem()
}

func (o ListenerRuleActionArrayOutput) ToListenerRuleActionArrayOutput() ListenerRuleActionArrayOutput {
	return o
}

func (o ListenerRuleActionArrayOutput) ToListenerRuleActionArrayOutputWithContext(ctx context.Context) ListenerRuleActionArrayOutput {
	return o
}

func (o ListenerRuleActionArrayOutput) Index(i pulumi.IntInput) ListenerRuleActionOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ListenerRuleAction {
		return vs[0].([]ListenerRuleAction)[vs[1].(int)]
	}).(ListenerRuleActionOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenerRuleActionOutput{})
	pulumi.RegisterOutputType(ListenerRuleActionArrayOutput{})
}
