// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 integration.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
//
// ## Example Usage
type Integration struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs.
	ConnectionId pulumi.StringPtrOutput `pulumi:"connectionId"`
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType pulumi.StringPtrOutput `pulumi:"connectionType"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy pulumi.StringPtrOutput `pulumi:"contentHandlingStrategy"`
	// The credentials required for the integration, if any.
	CredentialsArn pulumi.StringPtrOutput `pulumi:"credentialsArn"`
	// The description of the integration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod pulumi.StringPtrOutput `pulumi:"integrationMethod"`
	// The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
	IntegrationResponseSelectionExpression pulumi.StringOutput `pulumi:"integrationResponseSelectionExpression"`
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values.
	IntegrationSubtype pulumi.StringPtrOutput `pulumi:"integrationSubtype"`
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType pulumi.StringOutput `pulumi:"integrationType"`
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri pulumi.StringPtrOutput `pulumi:"integrationUri"`
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior pulumi.StringPtrOutput `pulumi:"passthroughBehavior"`
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion pulumi.StringPtrOutput `pulumi:"payloadFormatVersion"`
	// A key-value map specifying request parameters that are passed from the method request to the backend.
	// Supported only for WebSocket APIs.
	RequestParameters pulumi.StringMapOutput `pulumi:"requestParameters"`
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates pulumi.StringMapOutput `pulumi:"requestTemplates"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression pulumi.StringPtrOutput `pulumi:"templateSelectionExpression"`
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
	TimeoutMilliseconds pulumi.IntPtrOutput `pulumi:"timeoutMilliseconds"`
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig IntegrationTlsConfigPtrOutput `pulumi:"tlsConfig"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	if args == nil || args.IntegrationType == nil {
		return nil, errors.New("missing required argument 'IntegrationType'")
	}
	if args == nil {
		args = &IntegrationArgs{}
	}
	var resource Integration
	err := ctx.RegisterResource("aws:apigatewayv2/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("aws:apigatewayv2/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs.
	ConnectionId *string `pulumi:"connectionId"`
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType *string `pulumi:"connectionType"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy *string `pulumi:"contentHandlingStrategy"`
	// The credentials required for the integration, if any.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// The description of the integration.
	Description *string `pulumi:"description"`
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod *string `pulumi:"integrationMethod"`
	// The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
	IntegrationResponseSelectionExpression *string `pulumi:"integrationResponseSelectionExpression"`
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values.
	IntegrationSubtype *string `pulumi:"integrationSubtype"`
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType *string `pulumi:"integrationType"`
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri *string `pulumi:"integrationUri"`
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior *string `pulumi:"passthroughBehavior"`
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion *string `pulumi:"payloadFormatVersion"`
	// A key-value map specifying request parameters that are passed from the method request to the backend.
	// Supported only for WebSocket APIs.
	RequestParameters map[string]string `pulumi:"requestParameters"`
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]string `pulumi:"requestTemplates"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression *string `pulumi:"templateSelectionExpression"`
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
	TimeoutMilliseconds *int `pulumi:"timeoutMilliseconds"`
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig *IntegrationTlsConfig `pulumi:"tlsConfig"`
}

type IntegrationState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs.
	ConnectionId pulumi.StringPtrInput
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType pulumi.StringPtrInput
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy pulumi.StringPtrInput
	// The credentials required for the integration, if any.
	CredentialsArn pulumi.StringPtrInput
	// The description of the integration.
	Description pulumi.StringPtrInput
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod pulumi.StringPtrInput
	// The [integration response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions) for the integration.
	IntegrationResponseSelectionExpression pulumi.StringPtrInput
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values.
	IntegrationSubtype pulumi.StringPtrInput
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType pulumi.StringPtrInput
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri pulumi.StringPtrInput
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior pulumi.StringPtrInput
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion pulumi.StringPtrInput
	// A key-value map specifying request parameters that are passed from the method request to the backend.
	// Supported only for WebSocket APIs.
	RequestParameters pulumi.StringMapInput
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates pulumi.StringMapInput
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression pulumi.StringPtrInput
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
	TimeoutMilliseconds pulumi.IntPtrInput
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig IntegrationTlsConfigPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs.
	ConnectionId *string `pulumi:"connectionId"`
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType *string `pulumi:"connectionType"`
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy *string `pulumi:"contentHandlingStrategy"`
	// The credentials required for the integration, if any.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// The description of the integration.
	Description *string `pulumi:"description"`
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod *string `pulumi:"integrationMethod"`
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values.
	IntegrationSubtype *string `pulumi:"integrationSubtype"`
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType string `pulumi:"integrationType"`
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri *string `pulumi:"integrationUri"`
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior *string `pulumi:"passthroughBehavior"`
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion *string `pulumi:"payloadFormatVersion"`
	// A key-value map specifying request parameters that are passed from the method request to the backend.
	// Supported only for WebSocket APIs.
	RequestParameters map[string]string `pulumi:"requestParameters"`
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates map[string]string `pulumi:"requestTemplates"`
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression *string `pulumi:"templateSelectionExpression"`
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
	TimeoutMilliseconds *int `pulumi:"timeoutMilliseconds"`
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig *IntegrationTlsConfig `pulumi:"tlsConfig"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The ID of the VPC link for a private integration. Supported only for HTTP APIs.
	ConnectionId pulumi.StringPtrInput
	// The type of the network connection to the integration endpoint. Valid values: `INTERNET`, `VPC_LINK`. Default is `INTERNET`.
	ConnectionType pulumi.StringPtrInput
	// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`. Supported only for WebSocket APIs.
	ContentHandlingStrategy pulumi.StringPtrInput
	// The credentials required for the integration, if any.
	CredentialsArn pulumi.StringPtrInput
	// The description of the integration.
	Description pulumi.StringPtrInput
	// The integration's HTTP method. Must be specified if `integrationType` is not `MOCK`.
	IntegrationMethod pulumi.StringPtrInput
	// Specifies the AWS service action to invoke. Supported only for HTTP APIs when `integrationType` is `AWS_PROXY`. See the [AWS service integration reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) documentation for supported values.
	IntegrationSubtype pulumi.StringPtrInput
	// The integration type of an integration.
	// Valid values: `AWS`, `AWS_PROXY`, `HTTP`, `HTTP_PROXY`, `MOCK`.
	IntegrationType pulumi.StringInput
	// The URI of the Lambda function for a Lambda proxy integration, when `integrationType` is `AWS_PROXY`.
	// For an `HTTP` integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.
	IntegrationUri pulumi.StringPtrInput
	// The pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` attribute.
	// Valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`. Default is `WHEN_NO_MATCH`. Supported only for WebSocket APIs.
	PassthroughBehavior pulumi.StringPtrInput
	// The [format of the payload](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html#http-api-develop-integrations-lambda.proxy-format) sent to an integration. Valid values: `1.0`, `2.0`. Default is `1.0`.
	PayloadFormatVersion pulumi.StringPtrInput
	// A key-value map specifying request parameters that are passed from the method request to the backend.
	// Supported only for WebSocket APIs.
	RequestParameters pulumi.StringMapInput
	// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. Supported only for WebSocket APIs.
	RequestTemplates pulumi.StringMapInput
	// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration.
	TemplateSelectionExpression pulumi.StringPtrInput
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.
	TimeoutMilliseconds pulumi.IntPtrInput
	// The TLS configuration for a private integration. Supported only for HTTP APIs.
	TlsConfig IntegrationTlsConfigPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}
