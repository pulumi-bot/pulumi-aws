// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 integration response.
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
// 		_, err := apigatewayv2.NewIntegrationResponse(ctx, "example", &apigatewayv2.IntegrationResponseArgs{
// 			ApiId:                  pulumi.Any(aws_apigatewayv2_api.Example.Id),
// 			IntegrationId:          pulumi.Any(aws_apigatewayv2_integration.Example.Id),
// 			IntegrationResponseKey: pulumi.String("/200/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IntegrationResponse struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
	ContentHandlingStrategy pulumi.StringPtrOutput `pulumi:"contentHandlingStrategy"`
	// The identifier of the `apigatewayv2.Integration`.
	IntegrationId pulumi.StringOutput `pulumi:"integrationId"`
	// The integration response key.
	IntegrationResponseKey pulumi.StringOutput `pulumi:"integrationResponseKey"`
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	ResponseTemplates pulumi.StringMapOutput `pulumi:"responseTemplates"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
	TemplateSelectionExpression pulumi.StringPtrOutput `pulumi:"templateSelectionExpression"`
}

// NewIntegrationResponse registers a new resource with the given unique name, arguments, and options.
func NewIntegrationResponse(ctx *pulumi.Context,
	name string, args *IntegrationResponseArgs, opts ...pulumi.ResourceOption) (*IntegrationResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.IntegrationId == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationId'")
	}
	if args.IntegrationResponseKey == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationResponseKey'")
	}
	var resource IntegrationResponse
	err := ctx.RegisterResource("aws:apigatewayv2/integrationResponse:IntegrationResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationResponse gets an existing IntegrationResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationResponseState, opts ...pulumi.ResourceOption) (*IntegrationResponse, error) {
	var resource IntegrationResponse
	err := ctx.ReadResource("aws:apigatewayv2/integrationResponse:IntegrationResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationResponse resources.
type integrationResponseState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
	ContentHandlingStrategy *string `pulumi:"contentHandlingStrategy"`
	// The identifier of the `apigatewayv2.Integration`.
	IntegrationId *string `pulumi:"integrationId"`
	// The integration response key.
	IntegrationResponseKey *string `pulumi:"integrationResponseKey"`
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	ResponseTemplates map[string]string `pulumi:"responseTemplates"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
	TemplateSelectionExpression *string `pulumi:"templateSelectionExpression"`
}

type IntegrationResponseState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
	ContentHandlingStrategy pulumi.StringPtrInput
	// The identifier of the `apigatewayv2.Integration`.
	IntegrationId pulumi.StringPtrInput
	// The integration response key.
	IntegrationResponseKey pulumi.StringPtrInput
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	ResponseTemplates pulumi.StringMapInput
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
	TemplateSelectionExpression pulumi.StringPtrInput
}

func (IntegrationResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationResponseState)(nil)).Elem()
}

type integrationResponseArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
	ContentHandlingStrategy *string `pulumi:"contentHandlingStrategy"`
	// The identifier of the `apigatewayv2.Integration`.
	IntegrationId string `pulumi:"integrationId"`
	// The integration response key.
	IntegrationResponseKey string `pulumi:"integrationResponseKey"`
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	ResponseTemplates map[string]string `pulumi:"responseTemplates"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
	TemplateSelectionExpression *string `pulumi:"templateSelectionExpression"`
}

// The set of arguments for constructing a IntegrationResponse resource.
type IntegrationResponseArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
	ContentHandlingStrategy pulumi.StringPtrInput
	// The identifier of the `apigatewayv2.Integration`.
	IntegrationId pulumi.StringInput
	// The integration response key.
	IntegrationResponseKey pulumi.StringInput
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	ResponseTemplates pulumi.StringMapInput
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
	TemplateSelectionExpression pulumi.StringPtrInput
}

func (IntegrationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationResponseArgs)(nil)).Elem()
}
