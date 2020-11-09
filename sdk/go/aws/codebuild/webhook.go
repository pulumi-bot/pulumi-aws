// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a CodeBuild webhook, which is an endpoint accepted by the CodeBuild service to trigger builds from source code repositories. Depending on the source type of the CodeBuild project, the CodeBuild service may also automatically create and delete the actual repository webhook as well.
//
// ## Example Usage
// ### Bitbucket and GitHub
//
// When working with [Bitbucket](https://bitbucket.org) and [GitHub](https://github.com) source CodeBuild webhooks, the CodeBuild service will automatically create (on `codebuild.Webhook` resource creation) and delete (on `codebuild.Webhook` resource deletion) the Bitbucket/GitHub repository webhook using its granted OAuth permissions. This behavior cannot be controlled by this provider.
//
// > **Note:** The AWS account that this provider uses to create this resource *must* have authorized CodeBuild to access Bitbucket/GitHub's OAuth API in each applicable region. This is a manual step that must be done *before* creating webhooks with this resource. If OAuth is not configured, AWS will return an error similar to `ResourceNotFoundException: Could not find access token for server type github`. More information can be found in the CodeBuild User Guide for [Bitbucket](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-bitbucket-pull-request.html) and [GitHub](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-pull-request.html).
//
// > **Note:** Further managing the automatically created Bitbucket/GitHub webhook with the `bitbucketHook`/`githubRepositoryWebhook` resource is only possible with importing that resource after creation of the `codebuild.Webhook` resource. The CodeBuild API does not ever provide the `secret` attribute for the `codebuild.Webhook` resource in this scenario.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codebuild"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codebuild.NewWebhook(ctx, "example", &codebuild.WebhookArgs{
// 			ProjectName: pulumi.Any(aws_codebuild_project.Example.Name),
// 			FilterGroups: codebuild.WebhookFilterGroupArray{
// 				&codebuild.WebhookFilterGroupArgs{
// 					Filters: codebuild.WebhookFilterGroupFilterArray{
// 						&codebuild.WebhookFilterGroupFilterArgs{
// 							Type:    pulumi.String("EVENT"),
// 							Pattern: pulumi.String("PUSH"),
// 						},
// 						&codebuild.WebhookFilterGroupFilterArgs{
// 							Type:    pulumi.String("HEAD_REF"),
// 							Pattern: pulumi.String("master"),
// 						},
// 					},
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
// ### GitHub Enterprise
//
// When working with [GitHub Enterprise](https://enterprise.github.com/) source CodeBuild webhooks, the GHE repository webhook must be separately managed (e.g. manually or with the `githubRepositoryWebhook` resource).
//
// More information creating webhooks with GitHub Enterprise can be found in the [CodeBuild User Guide](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-enterprise.html).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codebuild"
// 	"github.com/pulumi/pulumi-github/sdk/go/github"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleWebhook, err := codebuild.NewWebhook(ctx, "exampleWebhook", &codebuild.WebhookArgs{
// 			ProjectName: pulumi.Any(aws_codebuild_project.Example.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = github.NewRepositoryWebhook(ctx, "exampleRepositoryWebhook", &github.RepositoryWebhookArgs{
// 			Active: pulumi.Bool(true),
// 			Events: pulumi.StringArray{
// 				pulumi.String("push"),
// 			},
// 			Repository: pulumi.Any(github_repository.Example.Name),
// 			Configuration: &github.RepositoryWebhookConfigurationArgs{
// 				Url:         exampleWebhook.PayloadUrl,
// 				Secret:      exampleWebhook.Secret,
// 				ContentType: pulumi.String("json"),
// 				InsecureSsl: pulumi.Bool(false),
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
// CodeBuild Webhooks can be imported using the CodeBuild Project name, e.g.
//
// ```sh
//  $ pulumi import aws:codebuild/webhook:Webhook example MyProjectName
// ```
type Webhook struct {
	pulumi.CustomResourceState

	// A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filterGroup` over `branchFilter`.
	BranchFilter pulumi.StringPtrOutput `pulumi:"branchFilter"`
	// Information about the webhook's trigger. Filter group blocks are documented below.
	FilterGroups WebhookFilterGroupArrayOutput `pulumi:"filterGroups"`
	// The CodeBuild endpoint where webhook events are sent.
	PayloadUrl pulumi.StringOutput `pulumi:"payloadUrl"`
	// The name of the build project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The URL to the webhook.
	Url pulumi.StringOutput `pulumi:"url"`
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
	// A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filterGroup` over `branchFilter`.
	BranchFilter *string `pulumi:"branchFilter"`
	// Information about the webhook's trigger. Filter group blocks are documented below.
	FilterGroups []WebhookFilterGroup `pulumi:"filterGroups"`
	// The CodeBuild endpoint where webhook events are sent.
	PayloadUrl *string `pulumi:"payloadUrl"`
	// The name of the build project.
	ProjectName *string `pulumi:"projectName"`
	// The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
	Secret *string `pulumi:"secret"`
	// The URL to the webhook.
	Url *string `pulumi:"url"`
}

type WebhookState struct {
	// A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filterGroup` over `branchFilter`.
	BranchFilter pulumi.StringPtrInput
	// Information about the webhook's trigger. Filter group blocks are documented below.
	FilterGroups WebhookFilterGroupArrayInput
	// The CodeBuild endpoint where webhook events are sent.
	PayloadUrl pulumi.StringPtrInput
	// The name of the build project.
	ProjectName pulumi.StringPtrInput
	// The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
	Secret pulumi.StringPtrInput
	// The URL to the webhook.
	Url pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filterGroup` over `branchFilter`.
	BranchFilter *string `pulumi:"branchFilter"`
	// Information about the webhook's trigger. Filter group blocks are documented below.
	FilterGroups []WebhookFilterGroup `pulumi:"filterGroups"`
	// The name of the build project.
	ProjectName string `pulumi:"projectName"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filterGroup` over `branchFilter`.
	BranchFilter pulumi.StringPtrInput
	// Information about the webhook's trigger. Filter group blocks are documented below.
	FilterGroups WebhookFilterGroupArrayInput
	// The name of the build project.
	ProjectName pulumi.StringInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}
