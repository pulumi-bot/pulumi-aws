// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a HTTP Method for an API Gateway Resource.
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
// 			PathPart: pulumi.String("mydemoresource"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
// 			RestApi:       myDemoAPI.ID(),
// 			ResourceId:    myDemoResource.ID(),
// 			HttpMethod:    pulumi.String("GET"),
// 			Authorization: pulumi.String("NONE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Usage with Cognito User Pool Authorizer
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cognito"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		cognitoUserPoolName := cfg.RequireObject("cognitoUserPoolName")
// 		thisUserPools, err := cognito.GetUserPools(ctx, &cognito.GetUserPoolsArgs{
// 			Name: cognitoUserPoolName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		thisRestApi, err := apigateway.NewRestApi(ctx, "thisRestApi", nil)
// 		if err != nil {
// 			return err
// 		}
// 		thisResource, err := apigateway.NewResource(ctx, "thisResource", &apigateway.ResourceArgs{
// 			RestApi:  thisRestApi.ID(),
// 			ParentId: thisRestApi.RootResourceId,
// 			PathPart: pulumi.String("{proxy+}"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		thisAuthorizer, err := apigateway.NewAuthorizer(ctx, "thisAuthorizer", &apigateway.AuthorizerArgs{
// 			Type:         pulumi.String("COGNITO_USER_POOLS"),
// 			RestApi:      thisRestApi.ID(),
// 			ProviderArns: toPulumiStringArray(thisUserPools.Arns),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewMethod(ctx, "any", &apigateway.MethodArgs{
// 			RestApi:       thisRestApi.ID(),
// 			ResourceId:    thisResource.ID(),
// 			HttpMethod:    pulumi.String("ANY"),
// 			Authorization: pulumi.String("COGNITO_USER_POOLS"),
// 			AuthorizerId:  thisAuthorizer.ID(),
// 			RequestParameters: pulumi.BoolMap{
// 				"method.request.path.proxy": pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func toPulumiStringArray(arr []string) pulumi.StringArray {
// 	var pulumiArr pulumi.StringArray
// 	for _, v := range arr {
// 		pulumiArr = append(pulumiArr, pulumi.String(v))
// 	}
// 	return pulumiArr
// }
// ```
type Method struct {
	pulumi.CustomResourceState

	// Specify if the method requires an API key
	ApiKeyRequired pulumi.BoolPtrOutput `pulumi:"apiKeyRequired"`
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization pulumi.StringOutput `pulumi:"authorization"`
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes pulumi.StringArrayOutput `pulumi:"authorizationScopes"`
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId pulumi.StringPtrOutput `pulumi:"authorizerId"`
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels pulumi.StringMapOutput `pulumi:"requestModels"`
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters pulumi.BoolMapOutput `pulumi:"requestParameters"`
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId pulumi.StringPtrOutput `pulumi:"requestValidatorId"`
	// The API resource ID
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
}

// NewMethod registers a new resource with the given unique name, arguments, and options.
func NewMethod(ctx *pulumi.Context,
	name string, args *MethodArgs, opts ...pulumi.ResourceOption) (*Method, error) {
	if args == nil || args.Authorization == nil {
		return nil, errors.New("missing required argument 'Authorization'")
	}
	if args == nil || args.HttpMethod == nil {
		return nil, errors.New("missing required argument 'HttpMethod'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil {
		args = &MethodArgs{}
	}
	var resource Method
	err := ctx.RegisterResource("aws:apigateway/method:Method", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethod gets an existing Method resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodState, opts ...pulumi.ResourceOption) (*Method, error) {
	var resource Method
	err := ctx.ReadResource("aws:apigateway/method:Method", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Method resources.
type methodState struct {
	// Specify if the method requires an API key
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization *string `pulumi:"authorization"`
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId *string `pulumi:"authorizerId"`
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod *string `pulumi:"httpMethod"`
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels map[string]string `pulumi:"requestModels"`
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters map[string]bool `pulumi:"requestParameters"`
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId *string `pulumi:"requestValidatorId"`
	// The API resource ID
	ResourceId *string `pulumi:"resourceId"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
}

type MethodState struct {
	// Specify if the method requires an API key
	ApiKeyRequired pulumi.BoolPtrInput
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization pulumi.StringPtrInput
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes pulumi.StringArrayInput
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId pulumi.StringPtrInput
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringPtrInput
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels pulumi.StringMapInput
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters pulumi.BoolMapInput
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId pulumi.StringPtrInput
	// The API resource ID
	ResourceId pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
}

func (MethodState) ElementType() reflect.Type {
	return reflect.TypeOf((*methodState)(nil)).Elem()
}

type methodArgs struct {
	// Specify if the method requires an API key
	ApiKeyRequired *bool `pulumi:"apiKeyRequired"`
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization string `pulumi:"authorization"`
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes []string `pulumi:"authorizationScopes"`
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId *string `pulumi:"authorizerId"`
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod string `pulumi:"httpMethod"`
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels map[string]string `pulumi:"requestModels"`
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters map[string]bool `pulumi:"requestParameters"`
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId *string `pulumi:"requestValidatorId"`
	// The API resource ID
	ResourceId string `pulumi:"resourceId"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
}

// The set of arguments for constructing a Method resource.
type MethodArgs struct {
	// Specify if the method requires an API key
	ApiKeyRequired pulumi.BoolPtrInput
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization pulumi.StringInput
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes pulumi.StringArrayInput
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId pulumi.StringPtrInput
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels pulumi.StringMapInput
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters pulumi.BoolMapInput
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId pulumi.StringPtrInput
	// The API resource ID
	ResourceId pulumi.StringInput
	// The ID of the associated REST API
	RestApi pulumi.Input
}

func (MethodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*methodArgs)(nil)).Elem()
}

type MethodInput interface {
	pulumi.Input

	ToMethodOutput() MethodOutput
	ToMethodOutputWithContext(ctx context.Context) MethodOutput
}

func (Method) ElementType() reflect.Type {
	return reflect.TypeOf((*Method)(nil)).Elem()
}

func (i Method) ToMethodOutput() MethodOutput {
	return i.ToMethodOutputWithContext(context.Background())
}

func (i Method) ToMethodOutputWithContext(ctx context.Context) MethodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodOutput)
}

type MethodOutput struct {
	*pulumi.OutputState
}

func (MethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MethodOutput)(nil)).Elem()
}

func (o MethodOutput) ToMethodOutput() MethodOutput {
	return o
}

func (o MethodOutput) ToMethodOutputWithContext(ctx context.Context) MethodOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MethodOutput{})
}
