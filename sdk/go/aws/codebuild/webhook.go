// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Webhook struct {
	pulumi.CustomResourceState

	BranchFilter pulumi.StringPtrOutput        `pulumi:"branchFilter"`
	FilterGroups WebhookFilterGroupArrayOutput `pulumi:"filterGroups"`
	PayloadUrl   pulumi.StringOutput           `pulumi:"payloadUrl"`
	ProjectName  pulumi.StringOutput           `pulumi:"projectName"`
	Secret       pulumi.StringOutput           `pulumi:"secret"`
	Url          pulumi.StringOutput           `pulumi:"url"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil {
		args = &WebhookArgs{}
	}
	var resource Webhook
	err := ctx.RegisterResource("aws:codebuild/webhook:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("aws:codebuild/webhook:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	BranchFilter *string              `pulumi:"branchFilter"`
	FilterGroups []WebhookFilterGroup `pulumi:"filterGroups"`
	PayloadUrl   *string              `pulumi:"payloadUrl"`
	ProjectName  *string              `pulumi:"projectName"`
	Secret       *string              `pulumi:"secret"`
	Url          *string              `pulumi:"url"`
}

type WebhookState struct {
	BranchFilter pulumi.StringPtrInput
	FilterGroups WebhookFilterGroupArrayInput
	PayloadUrl   pulumi.StringPtrInput
	ProjectName  pulumi.StringPtrInput
	Secret       pulumi.StringPtrInput
	Url          pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	BranchFilter *string              `pulumi:"branchFilter"`
	FilterGroups []WebhookFilterGroup `pulumi:"filterGroups"`
	ProjectName  string               `pulumi:"projectName"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	BranchFilter pulumi.StringPtrInput
	FilterGroups WebhookFilterGroupArrayInput
	ProjectName  pulumi.StringInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}
