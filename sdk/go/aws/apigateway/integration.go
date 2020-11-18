// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an HTTP Method Integration for an API Gateway Integration.
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
// 			CacheKeyParameters: pulumi.StringArray{
// 				pulumi.String("method.request.path.param"),
// 			},
// 			CacheNamespace:      pulumi.String("foobar"),
// 			TimeoutMilliseconds: pulumi.Int(29000),
// 			RequestParameters: pulumi.StringMap{
// 				"integration.request.header.X-Authorization": pulumi.String("'static'"),
// 			},
// 			RequestTemplates: pulumi.StringMap{
// 				"application/xml": pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "{\n", "   \"body\" : ", "$", "input.json('", "$", "')\n", "}\n")),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Integration struct {
	pulumi.CustomResourceState

	// A list of cache key parameters for the integration.
	CacheKeyParameters pulumi.StringArrayOutput `pulumi:"cacheKeyParameters"`
	// The integration's cache namespace.
	CacheNamespace pulumi.StringOutput `pulumi:"cacheNamespace"`
	// The id of the VpcLink used for the integration. **Required** if `connectionType` is `VPC_LINK`
	ConnectionId pulumi.StringPtrOutput `pulumi:"connectionId"`
	// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType pulumi.StringPtrOutput `pulumi:"connectionType"`
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling pulumi.StringPtrOutput `pulumi:"contentHandling"`
	// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
	Credentials pulumi.StringPtrOutput `pulumi:"credentials"`
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
	// when calling the associated resource.
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// The integration HTTP method
	// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
	// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// Not all methods are compatible with all `AWS` integrations.
	// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
	IntegrationHttpMethod pulumi.StringPtrOutput `pulumi:"integrationHttpMethod"`
	// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `requestTemplates` is used.
	PassthroughBehavior pulumi.StringOutput `pulumi:"passthroughBehavior"`
	// A map of request query string parameters and headers that should be passed to the backend responder.
	// For example: `requestParameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
	RequestParameters pulumi.StringMapOutput `pulumi:"requestParameters"`
	// A map of the integration's request templates.
	RequestTemplates pulumi.StringMapOutput `pulumi:"requestTemplates"`
	// The API resource ID.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The ID of the associated REST API.
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds pulumi.IntPtrOutput `pulumi:"timeoutMilliseconds"`
	// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connectionType` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type pulumi.StringOutput `pulumi:"type"`
	// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
	// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	Uri pulumi.StringPtrOutput `pulumi:"uri"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HttpMethod == nil {
		return nil, errors.New("invalid value for required argument 'HttpMethod'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Integration
	err := ctx.RegisterResource("aws:apigateway/integration:Integration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:apigateway/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// A list of cache key parameters for the integration.
	CacheKeyParameters []string `pulumi:"cacheKeyParameters"`
	// The integration's cache namespace.
	CacheNamespace *string `pulumi:"cacheNamespace"`
	// The id of the VpcLink used for the integration. **Required** if `connectionType` is `VPC_LINK`
	ConnectionId *string `pulumi:"connectionId"`
	// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType *string `pulumi:"connectionType"`
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling *string `pulumi:"contentHandling"`
	// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
	Credentials *string `pulumi:"credentials"`
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
	// when calling the associated resource.
	HttpMethod *string `pulumi:"httpMethod"`
	// The integration HTTP method
	// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
	// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// Not all methods are compatible with all `AWS` integrations.
	// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
	IntegrationHttpMethod *string `pulumi:"integrationHttpMethod"`
	// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `requestTemplates` is used.
	PassthroughBehavior *string `pulumi:"passthroughBehavior"`
	// A map of request query string parameters and headers that should be passed to the backend responder.
	// For example: `requestParameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
	RequestParameters map[string]string `pulumi:"requestParameters"`
	// A map of the integration's request templates.
	RequestTemplates map[string]string `pulumi:"requestTemplates"`
	// The API resource ID.
	ResourceId *string `pulumi:"resourceId"`
	// The ID of the associated REST API.
	RestApi *string `pulumi:"restApi"`
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds *int `pulumi:"timeoutMilliseconds"`
	// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connectionType` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type *string `pulumi:"type"`
	// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
	// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	Uri *string `pulumi:"uri"`
}

type IntegrationState struct {
	// A list of cache key parameters for the integration.
	CacheKeyParameters pulumi.StringArrayInput
	// The integration's cache namespace.
	CacheNamespace pulumi.StringPtrInput
	// The id of the VpcLink used for the integration. **Required** if `connectionType` is `VPC_LINK`
	ConnectionId pulumi.StringPtrInput
	// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType pulumi.StringPtrInput
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling pulumi.StringPtrInput
	// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
	Credentials pulumi.StringPtrInput
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
	// when calling the associated resource.
	HttpMethod pulumi.StringPtrInput
	// The integration HTTP method
	// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
	// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// Not all methods are compatible with all `AWS` integrations.
	// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
	IntegrationHttpMethod pulumi.StringPtrInput
	// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `requestTemplates` is used.
	PassthroughBehavior pulumi.StringPtrInput
	// A map of request query string parameters and headers that should be passed to the backend responder.
	// For example: `requestParameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
	RequestParameters pulumi.StringMapInput
	// A map of the integration's request templates.
	RequestTemplates pulumi.StringMapInput
	// The API resource ID.
	ResourceId pulumi.StringPtrInput
	// The ID of the associated REST API.
	RestApi pulumi.StringPtrInput
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds pulumi.IntPtrInput
	// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connectionType` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type pulumi.StringPtrInput
	// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
	// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	Uri pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// A list of cache key parameters for the integration.
	CacheKeyParameters []string `pulumi:"cacheKeyParameters"`
	// The integration's cache namespace.
	CacheNamespace *string `pulumi:"cacheNamespace"`
	// The id of the VpcLink used for the integration. **Required** if `connectionType` is `VPC_LINK`
	ConnectionId *string `pulumi:"connectionId"`
	// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType *string `pulumi:"connectionType"`
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling *string `pulumi:"contentHandling"`
	// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
	Credentials *string `pulumi:"credentials"`
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
	// when calling the associated resource.
	HttpMethod string `pulumi:"httpMethod"`
	// The integration HTTP method
	// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
	// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// Not all methods are compatible with all `AWS` integrations.
	// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
	IntegrationHttpMethod *string `pulumi:"integrationHttpMethod"`
	// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `requestTemplates` is used.
	PassthroughBehavior *string `pulumi:"passthroughBehavior"`
	// A map of request query string parameters and headers that should be passed to the backend responder.
	// For example: `requestParameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
	RequestParameters map[string]string `pulumi:"requestParameters"`
	// A map of the integration's request templates.
	RequestTemplates map[string]string `pulumi:"requestTemplates"`
	// The API resource ID.
	ResourceId string `pulumi:"resourceId"`
	// The ID of the associated REST API.
	RestApi interface{} `pulumi:"restApi"`
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds *int `pulumi:"timeoutMilliseconds"`
	// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connectionType` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type string `pulumi:"type"`
	// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
	// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	Uri *string `pulumi:"uri"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// A list of cache key parameters for the integration.
	CacheKeyParameters pulumi.StringArrayInput
	// The integration's cache namespace.
	CacheNamespace pulumi.StringPtrInput
	// The id of the VpcLink used for the integration. **Required** if `connectionType` is `VPC_LINK`
	ConnectionId pulumi.StringPtrInput
	// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType pulumi.StringPtrInput
	// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling pulumi.StringPtrInput
	// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
	Credentials pulumi.StringPtrInput
	// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
	// when calling the associated resource.
	HttpMethod pulumi.StringInput
	// The integration HTTP method
	// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
	// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// Not all methods are compatible with all `AWS` integrations.
	// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
	IntegrationHttpMethod pulumi.StringPtrInput
	// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `requestTemplates` is used.
	PassthroughBehavior pulumi.StringPtrInput
	// A map of request query string parameters and headers that should be passed to the backend responder.
	// For example: `requestParameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
	RequestParameters pulumi.StringMapInput
	// A map of the integration's request templates.
	RequestTemplates pulumi.StringMapInput
	// The API resource ID.
	ResourceId pulumi.StringInput
	// The ID of the associated REST API.
	RestApi pulumi.Input
	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds pulumi.IntPtrInput
	// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connectionType` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type pulumi.StringInput
	// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
	// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	Uri pulumi.StringPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (Integration) ElementType() reflect.Type {
	return reflect.TypeOf((*Integration)(nil)).Elem()
}

func (i Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

type IntegrationOutput struct {
	*pulumi.OutputState
}

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationOutput)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationOutput{})
}
