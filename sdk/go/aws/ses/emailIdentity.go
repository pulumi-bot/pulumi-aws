// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SES email identity resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewEmailIdentity(ctx, "example", &ses.EmailIdentityArgs{
// 			Email: pulumi.String("email@example.com"),
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
// SES email identities can be imported using the email address.
//
// ```sh
//  $ pulumi import aws:ses/emailIdentity:EmailIdentity example email@example.com
// ```
type EmailIdentity struct {
	pulumi.CustomResourceState

	// The ARN of the email identity.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The email address to assign to SES
	Email pulumi.StringOutput `pulumi:"email"`
}

// NewEmailIdentity registers a new resource with the given unique name, arguments, and options.
func NewEmailIdentity(ctx *pulumi.Context,
	name string, args *EmailIdentityArgs, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	var resource EmailIdentity
	err := ctx.RegisterResource("aws:ses/emailIdentity:EmailIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailIdentity gets an existing EmailIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailIdentityState, opts ...pulumi.ResourceOption) (*EmailIdentity, error) {
	var resource EmailIdentity
	err := ctx.ReadResource("aws:ses/emailIdentity:EmailIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailIdentity resources.
type emailIdentityState struct {
	// The ARN of the email identity.
	Arn *string `pulumi:"arn"`
	// The email address to assign to SES
	Email *string `pulumi:"email"`
}

type EmailIdentityState struct {
	// The ARN of the email identity.
	Arn pulumi.StringPtrInput
	// The email address to assign to SES
	Email pulumi.StringPtrInput
}

func (EmailIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityState)(nil)).Elem()
}

type emailIdentityArgs struct {
	// The email address to assign to SES
	Email string `pulumi:"email"`
}

// The set of arguments for constructing a EmailIdentity resource.
type EmailIdentityArgs struct {
	// The email address to assign to SES
	Email pulumi.StringInput
}

func (EmailIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailIdentityArgs)(nil)).Elem()
}

type EmailIdentityInput interface {
	pulumi.Input

	ToEmailIdentityOutput() EmailIdentityOutput
	ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput
}

func (*EmailIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailIdentity)(nil))
}

func (i *EmailIdentity) ToEmailIdentityOutput() EmailIdentityOutput {
	return i.ToEmailIdentityOutputWithContext(context.Background())
}

func (i *EmailIdentity) ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityOutput)
}

func (i *EmailIdentity) ToEmailIdentityPtrOutput() EmailIdentityPtrOutput {
	return i.ToEmailIdentityPtrOutputWithContext(context.Background())
}

func (i *EmailIdentity) ToEmailIdentityPtrOutputWithContext(ctx context.Context) EmailIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityPtrOutput)
}

type EmailIdentityPtrInput interface {
	pulumi.Input

	ToEmailIdentityPtrOutput() EmailIdentityPtrOutput
	ToEmailIdentityPtrOutputWithContext(ctx context.Context) EmailIdentityPtrOutput
}

type emailIdentityPtrType EmailIdentityArgs

func (*emailIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailIdentity)(nil))
}

func (i *emailIdentityPtrType) ToEmailIdentityPtrOutput() EmailIdentityPtrOutput {
	return i.ToEmailIdentityPtrOutputWithContext(context.Background())
}

func (i *emailIdentityPtrType) ToEmailIdentityPtrOutputWithContext(ctx context.Context) EmailIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityOutput).ToEmailIdentityPtrOutput()
}

// EmailIdentityArrayInput is an input type that accepts EmailIdentityArray and EmailIdentityArrayOutput values.
// You can construct a concrete instance of `EmailIdentityArrayInput` via:
//
//          EmailIdentityArray{ EmailIdentityArgs{...} }
type EmailIdentityArrayInput interface {
	pulumi.Input

	ToEmailIdentityArrayOutput() EmailIdentityArrayOutput
	ToEmailIdentityArrayOutputWithContext(context.Context) EmailIdentityArrayOutput
}

type EmailIdentityArray []EmailIdentityInput

func (EmailIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailIdentity)(nil))
}

func (i EmailIdentityArray) ToEmailIdentityArrayOutput() EmailIdentityArrayOutput {
	return i.ToEmailIdentityArrayOutputWithContext(context.Background())
}

func (i EmailIdentityArray) ToEmailIdentityArrayOutputWithContext(ctx context.Context) EmailIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityArrayOutput)
}

// EmailIdentityMapInput is an input type that accepts EmailIdentityMap and EmailIdentityMapOutput values.
// You can construct a concrete instance of `EmailIdentityMapInput` via:
//
//          EmailIdentityMap{ "key": EmailIdentityArgs{...} }
type EmailIdentityMapInput interface {
	pulumi.Input

	ToEmailIdentityMapOutput() EmailIdentityMapOutput
	ToEmailIdentityMapOutputWithContext(context.Context) EmailIdentityMapOutput
}

type EmailIdentityMap map[string]EmailIdentityInput

func (EmailIdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EmailIdentity)(nil))
}

func (i EmailIdentityMap) ToEmailIdentityMapOutput() EmailIdentityMapOutput {
	return i.ToEmailIdentityMapOutputWithContext(context.Background())
}

func (i EmailIdentityMap) ToEmailIdentityMapOutputWithContext(ctx context.Context) EmailIdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailIdentityMapOutput)
}

type EmailIdentityOutput struct {
	*pulumi.OutputState
}

func (EmailIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailIdentity)(nil))
}

func (o EmailIdentityOutput) ToEmailIdentityOutput() EmailIdentityOutput {
	return o
}

func (o EmailIdentityOutput) ToEmailIdentityOutputWithContext(ctx context.Context) EmailIdentityOutput {
	return o
}

func (o EmailIdentityOutput) ToEmailIdentityPtrOutput() EmailIdentityPtrOutput {
	return o.ToEmailIdentityPtrOutputWithContext(context.Background())
}

func (o EmailIdentityOutput) ToEmailIdentityPtrOutputWithContext(ctx context.Context) EmailIdentityPtrOutput {
	return o.ApplyT(func(v EmailIdentity) *EmailIdentity {
		return &v
	}).(EmailIdentityPtrOutput)
}

type EmailIdentityPtrOutput struct {
	*pulumi.OutputState
}

func (EmailIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailIdentity)(nil))
}

func (o EmailIdentityPtrOutput) ToEmailIdentityPtrOutput() EmailIdentityPtrOutput {
	return o
}

func (o EmailIdentityPtrOutput) ToEmailIdentityPtrOutputWithContext(ctx context.Context) EmailIdentityPtrOutput {
	return o
}

type EmailIdentityArrayOutput struct{ *pulumi.OutputState }

func (EmailIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailIdentity)(nil))
}

func (o EmailIdentityArrayOutput) ToEmailIdentityArrayOutput() EmailIdentityArrayOutput {
	return o
}

func (o EmailIdentityArrayOutput) ToEmailIdentityArrayOutputWithContext(ctx context.Context) EmailIdentityArrayOutput {
	return o
}

func (o EmailIdentityArrayOutput) Index(i pulumi.IntInput) EmailIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailIdentity {
		return vs[0].([]EmailIdentity)[vs[1].(int)]
	}).(EmailIdentityOutput)
}

type EmailIdentityMapOutput struct{ *pulumi.OutputState }

func (EmailIdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EmailIdentity)(nil))
}

func (o EmailIdentityMapOutput) ToEmailIdentityMapOutput() EmailIdentityMapOutput {
	return o
}

func (o EmailIdentityMapOutput) ToEmailIdentityMapOutputWithContext(ctx context.Context) EmailIdentityMapOutput {
	return o
}

func (o EmailIdentityMapOutput) MapIndex(k pulumi.StringInput) EmailIdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EmailIdentity {
		return vs[0].(map[string]EmailIdentity)[vs[1].(string)]
	}).(EmailIdentityOutput)
}

func init() {
	pulumi.RegisterOutputType(EmailIdentityOutput{})
	pulumi.RegisterOutputType(EmailIdentityPtrOutput{})
	pulumi.RegisterOutputType(EmailIdentityArrayOutput{})
	pulumi.RegisterOutputType(EmailIdentityMapOutput{})
}
