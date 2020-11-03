// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codepipeline

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodePipeline Webhook.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codepipeline"
// 	"github.com/pulumi/pulumi-github/sdk/go/github"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		barPipeline, err := codepipeline.NewPipeline(ctx, "barPipeline", &codepipeline.PipelineArgs{
// 			RoleArn: pulumi.Any(aws_iam_role.Bar.Arn),
// 			ArtifactStore: &codepipeline.PipelineArtifactStoreArgs{
// 				Location: pulumi.Any(aws_s3_bucket.Bar.Bucket),
// 				Type:     pulumi.String("S3"),
// 				EncryptionKey: &codepipeline.PipelineArtifactStoreEncryptionKeyArgs{
// 					Id:   pulumi.Any(data.Aws_kms_alias.S3kmskey.Arn),
// 					Type: pulumi.String("KMS"),
// 				},
// 			},
// 			Stages: codepipeline.PipelineStageArray{
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Source"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						&codepipeline.PipelineStageActionArgs{
// 							Name:     pulumi.String("Source"),
// 							Category: pulumi.String("Source"),
// 							Owner:    pulumi.String("ThirdParty"),
// 							Provider: pulumi.String("GitHub"),
// 							Version:  pulumi.String("1"),
// 							OutputArtifacts: pulumi.StringArray{
// 								pulumi.String("test"),
// 							},
// 							Configuration: pulumi.StringMap{
// 								"Owner":  pulumi.String("my-organization"),
// 								"Repo":   pulumi.String("test"),
// 								"Branch": pulumi.String("master"),
// 							},
// 						},
// 					},
// 				},
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Build"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						&codepipeline.PipelineStageActionArgs{
// 							Name:     pulumi.String("Build"),
// 							Category: pulumi.String("Build"),
// 							Owner:    pulumi.String("AWS"),
// 							Provider: pulumi.String("CodeBuild"),
// 							InputArtifacts: pulumi.StringArray{
// 								pulumi.String("test"),
// 							},
// 							Version: pulumi.String("1"),
// 							Configuration: pulumi.StringMap{
// 								"ProjectName": pulumi.String("test"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		webhookSecret := "super-secret"
// 		barWebhook, err := codepipeline.NewWebhook(ctx, "barWebhook", &codepipeline.WebhookArgs{
// 			Authentication: pulumi.String("GITHUB_HMAC"),
// 			TargetAction:   pulumi.String("Source"),
// 			TargetPipeline: barPipeline.Name,
// 			AuthenticationConfiguration: &codepipeline.WebhookAuthenticationConfigurationArgs{
// 				SecretToken: pulumi.String(webhookSecret),
// 			},
// 			Filters: codepipeline.WebhookFilterArray{
// 				&codepipeline.WebhookFilterArgs{
// 					JsonPath:    pulumi.String(fmt.Sprintf("%v%v", "$", ".ref")),
// 					MatchEquals: pulumi.String("refs/heads/{Branch}"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = github.NewRepositoryWebhook(ctx, "barRepositoryWebhook", &github.RepositoryWebhookArgs{
// 			Repository: pulumi.Any(github_repository.Repo.Name),
// 			Configuration: &github.RepositoryWebhookConfigurationArgs{
// 				Url:         barWebhook.Url,
// 				ContentType: pulumi.String("json"),
// 				InsecureSsl: pulumi.Bool(true),
// 				Secret:      pulumi.String(webhookSecret),
// 			},
// 			Events: pulumi.StringArray{
// 				pulumi.String("push"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Webhook struct {
	pulumi.CustomResourceState

	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration WebhookAuthenticationConfigurationPtrOutput `pulumi:"authenticationConfiguration"`
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters WebhookFilterArrayOutput `pulumi:"filters"`
	// The name of the webhook.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction pulumi.StringOutput `pulumi:"targetAction"`
	// The name of the pipeline.
	TargetPipeline pulumi.StringOutput `pulumi:"targetPipeline"`
	// The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Authentication == nil {
		return nil, errors.New("invalid value for required argument 'Authentication'")
	}
	if args.Filters == nil {
		return nil, errors.New("invalid value for required argument 'Filters'")
	}
	if args.TargetAction == nil {
		return nil, errors.New("invalid value for required argument 'TargetAction'")
	}
	if args.TargetPipeline == nil {
		return nil, errors.New("invalid value for required argument 'TargetPipeline'")
	}
	var resource Webhook
	err := ctx.RegisterResource("aws:codepipeline/webhook:Webhook", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:codepipeline/webhook:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication *string `pulumi:"authentication"`
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration *WebhookAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters []WebhookFilter `pulumi:"filters"`
	// The name of the webhook.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction *string `pulumi:"targetAction"`
	// The name of the pipeline.
	TargetPipeline *string `pulumi:"targetPipeline"`
	// The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
	Url *string `pulumi:"url"`
}

type WebhookState struct {
	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication pulumi.StringPtrInput
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration WebhookAuthenticationConfigurationPtrInput
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters WebhookFilterArrayInput
	// The name of the webhook.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction pulumi.StringPtrInput
	// The name of the pipeline.
	TargetPipeline pulumi.StringPtrInput
	// The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
	Url pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication string `pulumi:"authentication"`
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration *WebhookAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters []WebhookFilter `pulumi:"filters"`
	// The name of the webhook.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction string `pulumi:"targetAction"`
	// The name of the pipeline.
	TargetPipeline string `pulumi:"targetPipeline"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication pulumi.StringInput
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration WebhookAuthenticationConfigurationPtrInput
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters WebhookFilterArrayInput
	// The name of the webhook.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction pulumi.StringInput
	// The name of the pipeline.
	TargetPipeline pulumi.StringInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}
