// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway REST Deployment.
//
// > **Note:** This resource depends on having at least one `apigateway.Integration` created in the REST API, which
// itself has other dependencies. To avoid race conditions when all resources are being created together, you need to add
// implicit resource references via the `triggers` argument or explicit resource references using the
// [resource `dependsOn` meta-argument](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myDemoResource, err := apigateway.NewResource(ctx, "myDemoResource", &apigateway.ResourceArgs{
// 			RestApi:  myDemoAPI.ID(),
// 			ParentId: myDemoAPI.RootResourceId,
// 			PathPart: pulumi.String("test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myDemoMethod, err := apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
// 			RestApi:       myDemoAPI.ID(),
// 			ResourceId:    myDemoResource.ID(),
// 			HttpMethod:    pulumi.String("GET"),
// 			Authorization: pulumi.String("NONE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myDemoIntegration, err := apigateway.NewIntegration(ctx, "myDemoIntegration", &apigateway.IntegrationArgs{
// 			RestApi:    myDemoAPI.ID(),
// 			ResourceId: myDemoResource.ID(),
// 			HttpMethod: myDemoMethod.HttpMethod,
// 			Type:       pulumi.String("MOCK"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewDeployment(ctx, "myDemoDeployment", &apigateway.DeploymentArgs{
// 			RestApi:   myDemoAPI.ID(),
// 			StageName: pulumi.String("test"),
// 			Variables: pulumi.Map{
// 				"answer": pulumi.String("42"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Deployment struct {
	pulumi.CustomResourceState

	// The creation date of the deployment
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the deployment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The execution ARN to be used in `lambdaPermission` resource's `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringOutput `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The description of the stage
	StageDescription pulumi.StringPtrOutput `pulumi:"stageDescription"`
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName pulumi.StringPtrOutput `pulumi:"stageName"`
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers pulumi.StringMapOutput `pulumi:"triggers"`
	// A map that defines variables for the stage
	Variables pulumi.StringMapOutput `pulumi:"variables"`
}

// NewDeployment registers a new resource with the given unique name, arguments, and options.
func NewDeployment(ctx *pulumi.Context,
	name string, args *DeploymentArgs, opts ...pulumi.ResourceOption) (*Deployment, error) {
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil {
		args = &DeploymentArgs{}
	}
	var resource Deployment
	err := ctx.RegisterResource("aws:apigateway/deployment:Deployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployment gets an existing Deployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentState, opts ...pulumi.ResourceOption) (*Deployment, error) {
	var resource Deployment
	err := ctx.ReadResource("aws:apigateway/deployment:Deployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Deployment resources.
type deploymentState struct {
	// The creation date of the deployment
	CreatedDate *string `pulumi:"createdDate"`
	// The description of the deployment
	Description *string `pulumi:"description"`
	// The execution ARN to be used in `lambdaPermission` resource's `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn *string `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl *string `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The description of the stage
	StageDescription *string `pulumi:"stageDescription"`
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName *string `pulumi:"stageName"`
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]string `pulumi:"triggers"`
	// A map that defines variables for the stage
	Variables map[string]string `pulumi:"variables"`
}

type DeploymentState struct {
	// The creation date of the deployment
	CreatedDate pulumi.StringPtrInput
	// The description of the deployment
	Description pulumi.StringPtrInput
	// The execution ARN to be used in `lambdaPermission` resource's `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringPtrInput
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The description of the stage
	StageDescription pulumi.StringPtrInput
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName pulumi.StringPtrInput
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers pulumi.StringMapInput
	// A map that defines variables for the stage
	Variables pulumi.StringMapInput
}

func (DeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentState)(nil)).Elem()
}

type deploymentArgs struct {
	// The description of the deployment
	Description *string `pulumi:"description"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The description of the stage
	StageDescription *string `pulumi:"stageDescription"`
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName *string `pulumi:"stageName"`
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]string `pulumi:"triggers"`
	// A map that defines variables for the stage
	Variables map[string]string `pulumi:"variables"`
}

// The set of arguments for constructing a Deployment resource.
type DeploymentArgs struct {
	// The description of the deployment
	Description pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The description of the stage
	StageDescription pulumi.StringPtrInput
	// The name of the stage. If the specified stage already exists, it will be updated to point to the new deployment. If the stage does not exist, a new one will be created and point to this deployment.
	StageName pulumi.StringPtrInput
	// A map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers pulumi.StringMapInput
	// A map that defines variables for the stage
	Variables pulumi.StringMapInput
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentArgs)(nil)).Elem()
}
