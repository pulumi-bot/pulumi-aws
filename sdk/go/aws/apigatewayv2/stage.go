// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 stage.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewStage(ctx, "example", &apigatewayv2.StageArgs{
// 			ApiId: pulumi.Any(aws_apigatewayv2_api.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Stage struct {
	pulumi.CustomResourceState

	// Settings for logging access in this stage.
	// Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
	AccessLogSettings StageAccessLogSettingsPtrOutput `pulumi:"accessLogSettings"`
	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The ARN of the stage.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
	AutoDeploy pulumi.BoolPtrOutput `pulumi:"autoDeploy"`
	// The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateId pulumi.StringPtrOutput `pulumi:"clientCertificateId"`
	// The default route settings for the stage.
	DefaultRouteSettings StageDefaultRouteSettingsPtrOutput `pulumi:"defaultRouteSettings"`
	// The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The description for the stage.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ARN prefix to be used in an `lambda.Permission` `sourceArn` attribute.
	// For WebSocket APIs this attribute can additionally be used in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
	InvokeUrl pulumi.StringOutput `pulumi:"invokeUrl"`
	// The name of the stage.
	Name pulumi.StringOutput `pulumi:"name"`
	// Route settings for the stage.
	RouteSettings StageRouteSettingArrayOutput `pulumi:"routeSettings"`
	// A map that defines the stage variables for the stage.
	StageVariables pulumi.StringMapOutput `pulumi:"stageVariables"`
	// A map of tags to assign to the stage.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewStage registers a new resource with the given unique name, arguments, and options.
func NewStage(ctx *pulumi.Context,
	name string, args *StageArgs, opts ...pulumi.ResourceOption) (*Stage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource Stage
	err := ctx.RegisterResource("aws:apigatewayv2/stage:Stage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStage gets an existing Stage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageState, opts ...pulumi.ResourceOption) (*Stage, error) {
	var resource Stage
	err := ctx.ReadResource("aws:apigatewayv2/stage:Stage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stage resources.
type stageState struct {
	// Settings for logging access in this stage.
	// Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
	AccessLogSettings *StageAccessLogSettings `pulumi:"accessLogSettings"`
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// The ARN of the stage.
	Arn *string `pulumi:"arn"`
	// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
	AutoDeploy *bool `pulumi:"autoDeploy"`
	// The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The default route settings for the stage.
	DefaultRouteSettings *StageDefaultRouteSettings `pulumi:"defaultRouteSettings"`
	// The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
	DeploymentId *string `pulumi:"deploymentId"`
	// The description for the stage.
	Description *string `pulumi:"description"`
	// The ARN prefix to be used in an `lambda.Permission` `sourceArn` attribute.
	// For WebSocket APIs this attribute can additionally be used in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
	ExecutionArn *string `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
	InvokeUrl *string `pulumi:"invokeUrl"`
	// The name of the stage.
	Name *string `pulumi:"name"`
	// Route settings for the stage.
	RouteSettings []StageRouteSetting `pulumi:"routeSettings"`
	// A map that defines the stage variables for the stage.
	StageVariables map[string]string `pulumi:"stageVariables"`
	// A map of tags to assign to the stage.
	Tags map[string]string `pulumi:"tags"`
}

type StageState struct {
	// Settings for logging access in this stage.
	// Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
	AccessLogSettings StageAccessLogSettingsPtrInput
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// The ARN of the stage.
	Arn pulumi.StringPtrInput
	// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
	AutoDeploy pulumi.BoolPtrInput
	// The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateId pulumi.StringPtrInput
	// The default route settings for the stage.
	DefaultRouteSettings StageDefaultRouteSettingsPtrInput
	// The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
	DeploymentId pulumi.StringPtrInput
	// The description for the stage.
	Description pulumi.StringPtrInput
	// The ARN prefix to be used in an `lambda.Permission` `sourceArn` attribute.
	// For WebSocket APIs this attribute can additionally be used in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
	// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
	ExecutionArn pulumi.StringPtrInput
	// The URL to invoke the API pointing to the stage,
	// e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
	InvokeUrl pulumi.StringPtrInput
	// The name of the stage.
	Name pulumi.StringPtrInput
	// Route settings for the stage.
	RouteSettings StageRouteSettingArrayInput
	// A map that defines the stage variables for the stage.
	StageVariables pulumi.StringMapInput
	// A map of tags to assign to the stage.
	Tags pulumi.StringMapInput
}

func (StageState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageState)(nil)).Elem()
}

type stageArgs struct {
	// Settings for logging access in this stage.
	// Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
	AccessLogSettings *StageAccessLogSettings `pulumi:"accessLogSettings"`
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
	AutoDeploy *bool `pulumi:"autoDeploy"`
	// The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The default route settings for the stage.
	DefaultRouteSettings *StageDefaultRouteSettings `pulumi:"defaultRouteSettings"`
	// The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
	DeploymentId *string `pulumi:"deploymentId"`
	// The description for the stage.
	Description *string `pulumi:"description"`
	// The name of the stage.
	Name *string `pulumi:"name"`
	// Route settings for the stage.
	RouteSettings []StageRouteSetting `pulumi:"routeSettings"`
	// A map that defines the stage variables for the stage.
	StageVariables map[string]string `pulumi:"stageVariables"`
	// A map of tags to assign to the stage.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Stage resource.
type StageArgs struct {
	// Settings for logging access in this stage.
	// Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
	AccessLogSettings StageAccessLogSettingsPtrInput
	// The API identifier.
	ApiId pulumi.StringInput
	// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
	AutoDeploy pulumi.BoolPtrInput
	// The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
	// Supported only for WebSocket APIs.
	ClientCertificateId pulumi.StringPtrInput
	// The default route settings for the stage.
	DefaultRouteSettings StageDefaultRouteSettingsPtrInput
	// The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
	DeploymentId pulumi.StringPtrInput
	// The description for the stage.
	Description pulumi.StringPtrInput
	// The name of the stage.
	Name pulumi.StringPtrInput
	// Route settings for the stage.
	RouteSettings StageRouteSettingArrayInput
	// A map that defines the stage variables for the stage.
	StageVariables pulumi.StringMapInput
	// A map of tags to assign to the stage.
	Tags pulumi.StringMapInput
}

func (StageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageArgs)(nil)).Elem()
}
