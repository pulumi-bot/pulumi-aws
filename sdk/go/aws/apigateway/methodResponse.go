// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an HTTP Method Response for an API Gateway Resource.
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
// 		myDemoMethod, err := apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
// 			RestApi:       myDemoAPI.ID(),
// 			ResourceId:    myDemoResource.ID(),
// 			HttpMethod:    pulumi.String("GET"),
// 			Authorization: pulumi.String("NONE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewIntegration(ctx, "myDemoIntegration", &apigateway.IntegrationArgs{
// 			RestApi:    myDemoAPI.ID(),
// 			ResourceId: myDemoResource.ID(),
// 			HttpMethod: myDemoMethod.HttpMethod,
// 			Type:       pulumi.String("MOCK"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewMethodResponse(ctx, "response200", &apigateway.MethodResponseArgs{
// 			RestApi:    myDemoAPI.ID(),
// 			ResourceId: myDemoResource.ID(),
// 			HttpMethod: myDemoMethod.HttpMethod,
// 			StatusCode: pulumi.String("200"),
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
// `aws_api_gateway_method_response` can be imported using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/methodResponse:MethodResponse example 12345abcde/67890fghij/GET/200
// ```
type MethodResponse struct {
	pulumi.CustomResourceState

	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// The API resource ID
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// A map of the API models used for the response's content type
	ResponseModels pulumi.StringMapOutput `pulumi:"responseModels"`
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapOutput `pulumi:"responseParameters"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The HTTP status code
	StatusCode pulumi.StringOutput `pulumi:"statusCode"`
}

// NewMethodResponse registers a new resource with the given unique name, arguments, and options.
func NewMethodResponse(ctx *pulumi.Context,
	name string, args *MethodResponseArgs, opts ...pulumi.ResourceOption) (*MethodResponse, error) {
	if args == nil || args.HttpMethod == nil {
		return nil, errors.New("missing required argument 'HttpMethod'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.StatusCode == nil {
		return nil, errors.New("missing required argument 'StatusCode'")
	}
	if args == nil {
		args = &MethodResponseArgs{}
	}
	var resource MethodResponse
	err := ctx.RegisterResource("aws:apigateway/methodResponse:MethodResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethodResponse gets an existing MethodResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethodResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodResponseState, opts ...pulumi.ResourceOption) (*MethodResponse, error) {
	var resource MethodResponse
	err := ctx.ReadResource("aws:apigateway/methodResponse:MethodResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MethodResponse resources.
type methodResponseState struct {
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod *string `pulumi:"httpMethod"`
	// The API resource ID
	ResourceId *string `pulumi:"resourceId"`
	// A map of the API models used for the response's content type
	ResponseModels map[string]string `pulumi:"responseModels"`
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters map[string]bool `pulumi:"responseParameters"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The HTTP status code
	StatusCode *string `pulumi:"statusCode"`
}

type MethodResponseState struct {
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringPtrInput
	// The API resource ID
	ResourceId pulumi.StringPtrInput
	// A map of the API models used for the response's content type
	ResponseModels pulumi.StringMapInput
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The HTTP status code
	StatusCode pulumi.StringPtrInput
}

func (MethodResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*methodResponseState)(nil)).Elem()
}

type methodResponseArgs struct {
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod string `pulumi:"httpMethod"`
	// The API resource ID
	ResourceId string `pulumi:"resourceId"`
	// A map of the API models used for the response's content type
	ResponseModels map[string]string `pulumi:"responseModels"`
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters map[string]bool `pulumi:"responseParameters"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The HTTP status code
	StatusCode string `pulumi:"statusCode"`
}

// The set of arguments for constructing a MethodResponse resource.
type MethodResponseArgs struct {
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput
	// The API resource ID
	ResourceId pulumi.StringInput
	// A map of the API models used for the response's content type
	ResponseModels pulumi.StringMapInput
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The HTTP status code
	StatusCode pulumi.StringInput
}

func (MethodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*methodResponseArgs)(nil)).Elem()
}

type MethodResponseInput interface {
	pulumi.Input

	ToMethodResponseOutput() MethodResponseOutput
	ToMethodResponseOutputWithContext(ctx context.Context) MethodResponseOutput
}

func (MethodResponse) ElementType() reflect.Type {
	return reflect.TypeOf((*MethodResponse)(nil)).Elem()
}

func (i MethodResponse) ToMethodResponseOutput() MethodResponseOutput {
	return i.ToMethodResponseOutputWithContext(context.Background())
}

func (i MethodResponse) ToMethodResponseOutputWithContext(ctx context.Context) MethodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodResponseOutput)
}

type MethodResponseOutput struct {
	*pulumi.OutputState
}

func (MethodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MethodResponseOutput)(nil)).Elem()
}

func (o MethodResponseOutput) ToMethodResponseOutput() MethodResponseOutput {
	return o
}

func (o MethodResponseOutput) ToMethodResponseOutputWithContext(ctx context.Context) MethodResponseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MethodResponseOutput{})
}
