// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Stage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testRestApi, err := apigateway.NewRestApi(ctx, "testRestApi", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testResource, err := apigateway.NewResource(ctx, "testResource", &apigateway.ResourceArgs{
// 			RestApi:  testRestApi.ID(),
// 			ParentId: testRestApi.RootResourceId,
// 			PathPart: pulumi.String("mytestresource"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testMethod, err := apigateway.NewMethod(ctx, "testMethod", &apigateway.MethodArgs{
// 			RestApi:       testRestApi.ID(),
// 			ResourceId:    testResource.ID(),
// 			HttpMethod:    pulumi.String("GET"),
// 			Authorization: pulumi.String("NONE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testIntegration, err := apigateway.NewIntegration(ctx, "testIntegration", &apigateway.IntegrationArgs{
// 			RestApi:    testRestApi.ID(),
// 			ResourceId: testResource.ID(),
// 			HttpMethod: testMethod.HttpMethod,
// 			Type:       pulumi.String("MOCK"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testDeployment, err := apigateway.NewDeployment(ctx, "testDeployment", &apigateway.DeploymentArgs{
// 			RestApi:   testRestApi.ID(),
// 			StageName: pulumi.String("dev"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			testIntegration,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		testStage, err := apigateway.NewStage(ctx, "testStage", &apigateway.StageArgs{
// 			StageName:  pulumi.String("prod"),
// 			RestApi:    testRestApi.ID(),
// 			Deployment: testDeployment.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewMethodSettings(ctx, "methodSettings", &apigateway.MethodSettingsArgs{
// 			RestApi:   testRestApi.ID(),
// 			StageName: testStage.StageName,
// 			MethodPath: pulumi.All(testResource.PathPart, testMethod.HttpMethod).ApplyT(func(_args []interface{}) (string, error) {
// 				pathPart := _args[0].(string)
// 				httpMethod := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v", pathPart, "/", httpMethod), nil
// 			}).(pulumi.StringOutput),
// 			Settings: &apigateway.MethodSettingsSettingsArgs{
// 				MetricsEnabled: pulumi.Bool(true),
// 				LoggingLevel:   pulumi.String("INFO"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Managing the API Logging CloudWatch Log Group
//
// API Gateway provides the ability to [enable CloudWatch API logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html). To manage the CloudWatch Log Group when this feature is enabled, the `cloudwatch.LogGroup` resource can be used where the name matches the API Gateway naming convention. If the CloudWatch Log Group previously exists, the `cloudwatch.LogGroup` resource can be imported as a one time operation and recreation of the environment can occur without import.
//
// > The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with API Gateway automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		stageName := "example"
// 		if param := cfg.Get("stageName"); param != "" {
// 			stageName = param
// 		}
// 		_, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
// 			RetentionInDays: pulumi.Int(7),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
// 			StageName: pulumi.String(stageName),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleLogGroup,
// 		}))
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
// `aws_api_gateway_stage` can be imported using `REST-API-ID/STAGE-NAME`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/stage:Stage example 12345abcde/example
// ```
type Stage struct {
	pulumi.CustomResourceState

	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings StageAccessLogSettingsPtrOutput `pulumi:"accessLogSettings"`
	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled pulumi.BoolPtrOutput `pulumi:"cacheClusterEnabled"`
	// The size of the cache cluster for the stage, if enabled.
	// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize pulumi.StringPtrOutput `pulumi:"cacheClusterSize"`
	// The identifier of a client certificate for the stage.
	ClientCertificateId pulumi.StringPtrOutput `pulumi:"clientCertificateId"`
	// The ID of the deployment that the stage points to
	Deployment pulumi.StringOutput `pulumi:"deployment"`
	// The description of the stage
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The version of the associated API documentation
	DocumentationVersion pulumi.StringPtrOutput `pulumi:"documentationVersion"`
	// The execution ARN to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringOutput `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The name of the stage
	StageName pulumi.StringOutput `pulumi:"stageName"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map that defines the stage variables
	Variables pulumi.StringMapOutput `pulumi:"variables"`
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled pulumi.BoolPtrOutput `pulumi:"xrayTracingEnabled"`
}

// NewStage registers a new resource with the given unique name, arguments, and options.
func NewStage(ctx *pulumi.Context,
	name string, args *StageArgs, opts ...pulumi.ResourceOption) (*Stage, error) {
	if args == nil || args.Deployment == nil {
		return nil, errors.New("missing required argument 'Deployment'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.StageName == nil {
		return nil, errors.New("missing required argument 'StageName'")
	}
	if args == nil {
		args = &StageArgs{}
	}
	var resource Stage
	err := ctx.RegisterResource("aws:apigateway/stage:Stage", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:apigateway/stage:Stage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stage resources.
type stageState struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings *StageAccessLogSettings `pulumi:"accessLogSettings"`
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled *bool `pulumi:"cacheClusterEnabled"`
	// The size of the cache cluster for the stage, if enabled.
	// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize *string `pulumi:"cacheClusterSize"`
	// The identifier of a client certificate for the stage.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The ID of the deployment that the stage points to
	Deployment *string `pulumi:"deployment"`
	// The description of the stage
	Description *string `pulumi:"description"`
	// The version of the associated API documentation
	DocumentationVersion *string `pulumi:"documentationVersion"`
	// The execution ARN to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn *string `pulumi:"executionArn"`
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl *string `pulumi:"invokeUrl"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The name of the stage
	StageName *string `pulumi:"stageName"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A map that defines the stage variables
	Variables map[string]string `pulumi:"variables"`
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled *bool `pulumi:"xrayTracingEnabled"`
}

type StageState struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings StageAccessLogSettingsPtrInput
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled pulumi.BoolPtrInput
	// The size of the cache cluster for the stage, if enabled.
	// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize pulumi.StringPtrInput
	// The identifier of a client certificate for the stage.
	ClientCertificateId pulumi.StringPtrInput
	// The ID of the deployment that the stage points to
	Deployment pulumi.StringPtrInput
	// The description of the stage
	Description pulumi.StringPtrInput
	// The version of the associated API documentation
	DocumentationVersion pulumi.StringPtrInput
	// The execution ARN to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
	ExecutionArn pulumi.StringPtrInput
	// The URL to invoke the API pointing to the stage,
	// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
	InvokeUrl pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The name of the stage
	StageName pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A map that defines the stage variables
	Variables pulumi.StringMapInput
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled pulumi.BoolPtrInput
}

func (StageState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageState)(nil)).Elem()
}

type stageArgs struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings *StageAccessLogSettings `pulumi:"accessLogSettings"`
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled *bool `pulumi:"cacheClusterEnabled"`
	// The size of the cache cluster for the stage, if enabled.
	// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize *string `pulumi:"cacheClusterSize"`
	// The identifier of a client certificate for the stage.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The ID of the deployment that the stage points to
	Deployment interface{} `pulumi:"deployment"`
	// The description of the stage
	Description *string `pulumi:"description"`
	// The version of the associated API documentation
	DocumentationVersion *string `pulumi:"documentationVersion"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The name of the stage
	StageName string `pulumi:"stageName"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A map that defines the stage variables
	Variables map[string]string `pulumi:"variables"`
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled *bool `pulumi:"xrayTracingEnabled"`
}

// The set of arguments for constructing a Stage resource.
type StageArgs struct {
	// Enables access logs for the API stage. Detailed below.
	AccessLogSettings StageAccessLogSettingsPtrInput
	// Specifies whether a cache cluster is enabled for the stage
	CacheClusterEnabled pulumi.BoolPtrInput
	// The size of the cache cluster for the stage, if enabled.
	// Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
	CacheClusterSize pulumi.StringPtrInput
	// The identifier of a client certificate for the stage.
	ClientCertificateId pulumi.StringPtrInput
	// The ID of the deployment that the stage points to
	Deployment pulumi.Input
	// The description of the stage
	Description pulumi.StringPtrInput
	// The version of the associated API documentation
	DocumentationVersion pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The name of the stage
	StageName pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A map that defines the stage variables
	Variables pulumi.StringMapInput
	// Whether active tracing with X-ray is enabled. Defaults to `false`.
	XrayTracingEnabled pulumi.BoolPtrInput
}

func (StageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageArgs)(nil)).Elem()
}

type StageInput interface {
	pulumi.Input

	ToStageOutput() StageOutput
	ToStageOutputWithContext(ctx context.Context) StageOutput
}

func (Stage) ElementType() reflect.Type {
	return reflect.TypeOf((*Stage)(nil)).Elem()
}

func (i Stage) ToStageOutput() StageOutput {
	return i.ToStageOutputWithContext(context.Background())
}

func (i Stage) ToStageOutputWithContext(ctx context.Context) StageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageOutput)
}

type StageOutput struct {
	*pulumi.OutputState
}

func (StageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StageOutput)(nil)).Elem()
}

func (o StageOutput) ToStageOutput() StageOutput {
	return o
}

func (o StageOutput) ToStageOutputWithContext(ctx context.Context) StageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StageOutput{})
}
