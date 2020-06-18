// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Gateway Response for a REST API Gateway.
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
// 		main, err := apigateway.NewRestApi(ctx, "main", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewResponse(ctx, "test", &apigateway.ResponseArgs{
// 			ResponseParameters: map[string]interface{}{
// 				"gatewayresponse.header.Authorization": "'Basic'",
// 			},
// 			ResponseTemplates: map[string]interface{}{
// 				"application/json": fmt.Sprintf("%v%v%v", "{'message':", "$", "context.error.messageString}"),
// 			},
// 			ResponseType: pulumi.String("UNAUTHORIZED"),
// 			RestApiId:    main.ID(),
// 			StatusCode:   pulumi.String("401"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Response struct {
	pulumi.CustomResourceState

	// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
	ResponseParameters pulumi.StringMapOutput `pulumi:"responseParameters"`
	// A map specifying the templates used to transform the response body.
	ResponseTemplates pulumi.StringMapOutput `pulumi:"responseTemplates"`
	// The response type of the associated GatewayResponse.
	ResponseType pulumi.StringOutput `pulumi:"responseType"`
	// The string identifier of the associated REST API.
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
	// The HTTP status code of the Gateway Response.
	StatusCode pulumi.StringPtrOutput `pulumi:"statusCode"`
}

// NewResponse registers a new resource with the given unique name, arguments, and options.
func NewResponse(ctx *pulumi.Context,
	name string, args *ResponseArgs, opts ...pulumi.ResourceOption) (*Response, error) {
	if args == nil || args.ResponseType == nil {
		return nil, errors.New("missing required argument 'ResponseType'")
	}
	if args == nil || args.RestApiId == nil {
		return nil, errors.New("missing required argument 'RestApiId'")
	}
	if args == nil {
		args = &ResponseArgs{}
	}
	var resource Response
	err := ctx.RegisterResource("aws:apigateway/response:Response", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResponse gets an existing Response resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResponseState, opts ...pulumi.ResourceOption) (*Response, error) {
	var resource Response
	err := ctx.ReadResource("aws:apigateway/response:Response", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Response resources.
type responseState struct {
	// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
	ResponseParameters map[string]string `pulumi:"responseParameters"`
	// A map specifying the templates used to transform the response body.
	ResponseTemplates map[string]string `pulumi:"responseTemplates"`
	// The response type of the associated GatewayResponse.
	ResponseType *string `pulumi:"responseType"`
	// The string identifier of the associated REST API.
	RestApiId *string `pulumi:"restApiId"`
	// The HTTP status code of the Gateway Response.
	StatusCode *string `pulumi:"statusCode"`
}

type ResponseState struct {
	// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
	ResponseParameters pulumi.StringMapInput
	// A map specifying the templates used to transform the response body.
	ResponseTemplates pulumi.StringMapInput
	// The response type of the associated GatewayResponse.
	ResponseType pulumi.StringPtrInput
	// The string identifier of the associated REST API.
	RestApiId pulumi.StringPtrInput
	// The HTTP status code of the Gateway Response.
	StatusCode pulumi.StringPtrInput
}

func (ResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*responseState)(nil)).Elem()
}

type responseArgs struct {
	// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
	ResponseParameters map[string]string `pulumi:"responseParameters"`
	// A map specifying the templates used to transform the response body.
	ResponseTemplates map[string]string `pulumi:"responseTemplates"`
	// The response type of the associated GatewayResponse.
	ResponseType string `pulumi:"responseType"`
	// The string identifier of the associated REST API.
	RestApiId string `pulumi:"restApiId"`
	// The HTTP status code of the Gateway Response.
	StatusCode *string `pulumi:"statusCode"`
}

// The set of arguments for constructing a Response resource.
type ResponseArgs struct {
	// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
	ResponseParameters pulumi.StringMapInput
	// A map specifying the templates used to transform the response body.
	ResponseTemplates pulumi.StringMapInput
	// The response type of the associated GatewayResponse.
	ResponseType pulumi.StringInput
	// The string identifier of the associated REST API.
	RestApiId pulumi.StringInput
	// The HTTP status code of the Gateway Response.
	StatusCode pulumi.StringPtrInput
}

func (ResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*responseArgs)(nil)).Elem()
}
