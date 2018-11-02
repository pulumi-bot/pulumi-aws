// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a CodeBuild webhook, which is an endpoint accepted by the CodeBuild service to trigger builds from source code repositories. Depending on the source type of the CodeBuild project, the CodeBuild service may also automatically create and delete the actual repository webhook as well.
type Webhook struct {
	s *pulumi.ResourceState
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOpt) (*Webhook, error) {
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["branchFilter"] = nil
		inputs["projectName"] = nil
	} else {
		inputs["branchFilter"] = args.BranchFilter
		inputs["projectName"] = args.ProjectName
	}
	inputs["payloadUrl"] = nil
	inputs["secret"] = nil
	inputs["url"] = nil
	s, err := ctx.RegisterResource("aws:codebuild/webhook:Webhook", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Webhook{s: s}, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WebhookState, opts ...pulumi.ResourceOpt) (*Webhook, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["branchFilter"] = state.BranchFilter
		inputs["payloadUrl"] = state.PayloadUrl
		inputs["projectName"] = state.ProjectName
		inputs["secret"] = state.Secret
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("aws:codebuild/webhook:Webhook", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Webhook{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Webhook) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Webhook) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A regular expression used to determine which branches get built. Default is all branches are built.
func (r *Webhook) BranchFilter() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["branchFilter"])
}

// The CodeBuild endpoint where webhook events are sent.
func (r *Webhook) PayloadUrl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["payloadUrl"])
}

// The name of the build project.
func (r *Webhook) ProjectName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["projectName"])
}

// The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
func (r *Webhook) Secret() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["secret"])
}

// The URL to the webhook.
func (r *Webhook) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering Webhook resources.
type WebhookState struct {
	// A regular expression used to determine which branches get built. Default is all branches are built.
	BranchFilter interface{}
	// The CodeBuild endpoint where webhook events are sent.
	PayloadUrl interface{}
	// The name of the build project.
	ProjectName interface{}
	// The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
	Secret interface{}
	// The URL to the webhook.
	Url interface{}
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// A regular expression used to determine which branches get built. Default is all branches are built.
	BranchFilter interface{}
	// The name of the build project.
	ProjectName interface{}
}
