// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Method Settings, e.g. logging or monitoring.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
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
// 		testDeployment, err := apigateway.NewDeployment(ctx, "testDeployment", &apigateway.DeploymentArgs{
// 			RestApi:   testRestApi.ID(),
// 			StageName: pulumi.String("dev"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testStage, err := apigateway.NewStage(ctx, "testStage", &apigateway.StageArgs{
// 			Deployment: testDeployment.ID(),
// 			RestApi:    testRestApi.ID(),
// 			StageName:  pulumi.String("prod"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testResource, err := apigateway.NewResource(ctx, "testResource", &apigateway.ResourceArgs{
// 			ParentId: testRestApi.RootResourceId,
// 			PathPart: pulumi.String("mytestresource"),
// 			RestApi:  testRestApi.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testMethod, err := apigateway.NewMethod(ctx, "testMethod", &apigateway.MethodArgs{
// 			Authorization: pulumi.String("NONE"),
// 			HttpMethod:    pulumi.String("GET"),
// 			ResourceId:    testResource.ID(),
// 			RestApi:       testRestApi.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewMethodSettings(ctx, "methodSettings", &apigateway.MethodSettingsArgs{
// 			MethodPath: pulumi.All(testResource.PathPart, testMethod.HttpMethod).ApplyT(func(_args []interface{}) (string, error) {
// 				pathPart := _args[0].(string)
// 				httpMethod := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v", pathPart, "/", httpMethod), nil
// 			}).(pulumi.StringOutput),
// 			RestApi: testRestApi.ID(),
// 			Settings: &apigateway.MethodSettingsSettingsArgs{
// 				LoggingLevel:   pulumi.String("INFO"),
// 				MetricsEnabled: pulumi.Bool(true),
// 			},
// 			StageName: testStage.StageName,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewIntegration(ctx, "testIntegration", &apigateway.IntegrationArgs{
// 			HttpMethod: testMethod.HttpMethod,
// 			RequestTemplates: pulumi.Map{
// 				"application/xml": pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "{\n", "   \"body\" : ", "$", "input.json('", "$", "')\n", "}\n", "\n")),
// 			},
// 			ResourceId: testResource.ID(),
// 			RestApi:    testRestApi.ID(),
// 			Type:       pulumi.String("MOCK"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MethodSettings struct {
	pulumi.CustomResourceState

	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath pulumi.StringOutput `pulumi:"methodPath"`
	// The ID of the REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The settings block, see below.
	Settings MethodSettingsSettingsOutput `pulumi:"settings"`
	// The name of the stage
	StageName pulumi.StringOutput `pulumi:"stageName"`
}

// NewMethodSettings registers a new resource with the given unique name, arguments, and options.
func NewMethodSettings(ctx *pulumi.Context,
	name string, args *MethodSettingsArgs, opts ...pulumi.ResourceOption) (*MethodSettings, error) {
	if args == nil || args.MethodPath == nil {
		return nil, errors.New("missing required argument 'MethodPath'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	if args == nil || args.StageName == nil {
		return nil, errors.New("missing required argument 'StageName'")
	}
	if args == nil {
		args = &MethodSettingsArgs{}
	}
	var resource MethodSettings
	err := ctx.RegisterResource("aws:apigateway/methodSettings:MethodSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethodSettings gets an existing MethodSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethodSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodSettingsState, opts ...pulumi.ResourceOption) (*MethodSettings, error) {
	var resource MethodSettings
	err := ctx.ReadResource("aws:apigateway/methodSettings:MethodSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MethodSettings resources.
type methodSettingsState struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath *string `pulumi:"methodPath"`
	// The ID of the REST API
	RestApi *string `pulumi:"restApi"`
	// The settings block, see below.
	Settings *MethodSettingsSettings `pulumi:"settings"`
	// The name of the stage
	StageName *string `pulumi:"stageName"`
}

type MethodSettingsState struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath pulumi.StringPtrInput
	// The ID of the REST API
	RestApi pulumi.StringPtrInput
	// The settings block, see below.
	Settings MethodSettingsSettingsPtrInput
	// The name of the stage
	StageName pulumi.StringPtrInput
}

func (MethodSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*methodSettingsState)(nil)).Elem()
}

type methodSettingsArgs struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath string `pulumi:"methodPath"`
	// The ID of the REST API
	RestApi interface{} `pulumi:"restApi"`
	// The settings block, see below.
	Settings MethodSettingsSettings `pulumi:"settings"`
	// The name of the stage
	StageName string `pulumi:"stageName"`
}

// The set of arguments for constructing a MethodSettings resource.
type MethodSettingsArgs struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath pulumi.StringInput
	// The ID of the REST API
	RestApi pulumi.Input
	// The settings block, see below.
	Settings MethodSettingsSettingsInput
	// The name of the stage
	StageName pulumi.StringInput
}

func (MethodSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*methodSettingsArgs)(nil)).Elem()
}
