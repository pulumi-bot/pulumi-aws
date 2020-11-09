// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a SageMaker model resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "Service",
// 							Identifiers: []string{
// 								"sagemaker.amazonaws.com",
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(assumeRole.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sagemaker.NewModel(ctx, "exampleModel", &sagemaker.ModelArgs{
// 			ExecutionRoleArn: exampleRole.Arn,
// 			PrimaryContainer: &sagemaker.ModelPrimaryContainerArgs{
// 				Image: pulumi.String("174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1"),
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
// Models can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:sagemaker/model:Model test_model model-foo
// ```
type Model struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this model.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers ModelContainerArrayOutput `pulumi:"containers"`
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation pulumi.BoolPtrOutput `pulumi:"enableNetworkIsolation"`
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer ModelPrimaryContainerPtrOutput `pulumi:"primaryContainer"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig ModelVpcConfigPtrOutput `pulumi:"vpcConfig"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil || args.ExecutionRoleArn == nil {
		return nil, errors.New("missing required argument 'ExecutionRoleArn'")
	}
	if args == nil {
		args = &ModelArgs{}
	}
	var resource Model
	err := ctx.RegisterResource("aws:sagemaker/model:Model", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelState, opts ...pulumi.ResourceOption) (*Model, error) {
	var resource Model
	err := ctx.ReadResource("aws:sagemaker/model:Model", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Model resources.
type modelState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this model.
	Arn *string `pulumi:"arn"`
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers []ModelContainer `pulumi:"containers"`
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation *bool `pulumi:"enableNetworkIsolation"`
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer *ModelPrimaryContainer `pulumi:"primaryContainer"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig *ModelVpcConfig `pulumi:"vpcConfig"`
}

type ModelState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this model.
	Arn pulumi.StringPtrInput
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers ModelContainerArrayInput
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation pulumi.BoolPtrInput
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn pulumi.StringPtrInput
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer ModelPrimaryContainerPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig ModelVpcConfigPtrInput
}

func (ModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelState)(nil)).Elem()
}

type modelArgs struct {
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers []ModelContainer `pulumi:"containers"`
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation *bool `pulumi:"enableNetworkIsolation"`
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer *ModelPrimaryContainer `pulumi:"primaryContainer"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig *ModelVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
	Containers ModelContainerArrayInput
	// Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
	EnableNetworkIsolation pulumi.BoolPtrInput
	// A role that SageMaker can assume to access model artifacts and docker images for deployment.
	ExecutionRoleArn pulumi.StringInput
	// The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
	PrimaryContainer ModelPrimaryContainerPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
	VpcConfig ModelVpcConfigPtrInput
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelArgs)(nil)).Elem()
}
