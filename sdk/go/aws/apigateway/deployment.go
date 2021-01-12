// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
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
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
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
// 			Variables: pulumi.StringMap{
// 				"answer": pulumi.String("42"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			myDemoIntegration,
// 		}))
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
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

type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput
}

func (*Deployment) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (i *Deployment) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

func (i *Deployment) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *Deployment) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPtrOutput)
}

type DeploymentPtrInput interface {
	pulumi.Input

	ToDeploymentPtrOutput() DeploymentPtrOutput
	ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput
}

type deploymentPtrType DeploymentArgs

func (*deploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil))
}

func (i *deploymentPtrType) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *deploymentPtrType) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput).ToDeploymentPtrOutput()
}

// DeploymentArrayInput is an input type that accepts DeploymentArray and DeploymentArrayOutput values.
// You can construct a concrete instance of `DeploymentArrayInput` via:
//
//          DeploymentArray{ DeploymentArgs{...} }
type DeploymentArrayInput interface {
	pulumi.Input

	ToDeploymentArrayOutput() DeploymentArrayOutput
	ToDeploymentArrayOutputWithContext(context.Context) DeploymentArrayOutput
}

type DeploymentArray []DeploymentInput

func (DeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Deployment)(nil))
}

func (i DeploymentArray) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return i.ToDeploymentArrayOutputWithContext(context.Background())
}

func (i DeploymentArray) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentArrayOutput)
}

// DeploymentMapInput is an input type that accepts DeploymentMap and DeploymentMapOutput values.
// You can construct a concrete instance of `DeploymentMapInput` via:
//
//          DeploymentMap{ "key": DeploymentArgs{...} }
type DeploymentMapInput interface {
	pulumi.Input

	ToDeploymentMapOutput() DeploymentMapOutput
	ToDeploymentMapOutputWithContext(context.Context) DeploymentMapOutput
}

type DeploymentMap map[string]DeploymentInput

func (DeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Deployment)(nil))
}

func (i DeploymentMap) ToDeploymentMapOutput() DeploymentMapOutput {
	return i.ToDeploymentMapOutputWithContext(context.Background())
}

func (i DeploymentMap) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentMapOutput)
}

type DeploymentOutput struct {
	*pulumi.OutputState
}

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil))
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o.ToDeploymentPtrOutputWithContext(context.Background())
}

func (o DeploymentOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o.ApplyT(func(v Deployment) *Deployment {
		return &v
	}).(DeploymentPtrOutput)
}

type DeploymentPtrOutput struct {
	*pulumi.OutputState
}

func (DeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil))
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o
}

type DeploymentArrayOutput struct{ *pulumi.OutputState }

func (DeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Deployment)(nil))
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutput() DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) ToDeploymentArrayOutputWithContext(ctx context.Context) DeploymentArrayOutput {
	return o
}

func (o DeploymentArrayOutput) Index(i pulumi.IntInput) DeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Deployment {
		return vs[0].([]Deployment)[vs[1].(int)]
	}).(DeploymentOutput)
}

type DeploymentMapOutput struct{ *pulumi.OutputState }

func (DeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Deployment)(nil))
}

func (o DeploymentMapOutput) ToDeploymentMapOutput() DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) ToDeploymentMapOutputWithContext(ctx context.Context) DeploymentMapOutput {
	return o
}

func (o DeploymentMapOutput) MapIndex(k pulumi.StringInput) DeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Deployment {
		return vs[0].(map[string]Deployment)[vs[1].(string)]
	}).(DeploymentOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentOutput{})
	pulumi.RegisterOutputType(DeploymentPtrOutput{})
	pulumi.RegisterOutputType(DeploymentArrayOutput{})
	pulumi.RegisterOutputType(DeploymentMapOutput{})
}
