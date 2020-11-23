// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a SES template.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewTemplate(ctx, "myTemplate", &ses.TemplateArgs{
// 			Html:    pulumi.String("<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>"),
// 			Subject: pulumi.String("Greetings, {{name}}!"),
// 			Text: pulumi.String(fmt.Sprintf("%v%v", "Hello {{name}},\n", "Your favorite animal is {{favoriteanimal}}.\n")),
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
// SES templates can be imported using the template name, e.g.
//
// ```sh
//  $ pulumi import aws:ses/template:Template MyTemplate MyTemplate
// ```
type Template struct {
	pulumi.CustomResourceState

	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringPtrOutput `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringOutput `pulumi:"name"`
	// The subject line of the email.
	Subject pulumi.StringPtrOutput `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringPtrOutput `pulumi:"text"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		args = &TemplateArgs{}
	}

	var resource Template
	err := ctx.RegisterResource("aws:ses/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("aws:ses/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html *string `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name *string `pulumi:"name"`
	// The subject line of the email.
	Subject *string `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text *string `pulumi:"text"`
}

type TemplateState struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringPtrInput
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringPtrInput
	// The subject line of the email.
	Subject pulumi.StringPtrInput
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html *string `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name *string `pulumi:"name"`
	// The subject line of the email.
	Subject *string `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text *string `pulumi:"text"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringPtrInput
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringPtrInput
	// The subject line of the email.
	Subject pulumi.StringPtrInput
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (Template) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil)).Elem()
}

func (i Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

type TemplateOutput struct {
	*pulumi.OutputState
}

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateOutput)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TemplateOutput{})
}
