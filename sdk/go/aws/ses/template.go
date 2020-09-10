// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Template struct {
	pulumi.CustomResourceState

	Html    pulumi.StringPtrOutput `pulumi:"html"`
	Name    pulumi.StringOutput    `pulumi:"name"`
	Subject pulumi.StringPtrOutput `pulumi:"subject"`
	Text    pulumi.StringPtrOutput `pulumi:"text"`
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
	Html    *string `pulumi:"html"`
	Name    *string `pulumi:"name"`
	Subject *string `pulumi:"subject"`
	Text    *string `pulumi:"text"`
}

type TemplateState struct {
	Html    pulumi.StringPtrInput
	Name    pulumi.StringPtrInput
	Subject pulumi.StringPtrInput
	Text    pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	Html    *string `pulumi:"html"`
	Name    *string `pulumi:"name"`
	Subject *string `pulumi:"subject"`
	Text    *string `pulumi:"text"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	Html    pulumi.StringPtrInput
	Name    pulumi.StringPtrInput
	Subject pulumi.StringPtrInput
	Text    pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}
