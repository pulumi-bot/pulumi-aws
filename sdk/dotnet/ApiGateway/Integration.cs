// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an HTTP Method Integration for an API Gateway Integration.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///             Description = "This is my API for demonstration purposes",
    ///         });
    ///         var myDemoResource = new Aws.ApiGateway.Resource("myDemoResource", new Aws.ApiGateway.ResourceArgs
    ///         {
    ///             RestApi = myDemoAPI.Id,
    ///             ParentId = myDemoAPI.RootResourceId,
    ///             PathPart = "mydemoresource",
    ///         });
    ///         var myDemoMethod = new Aws.ApiGateway.Method("myDemoMethod", new Aws.ApiGateway.MethodArgs
    ///         {
    ///             RestApi = myDemoAPI.Id,
    ///             ResourceId = myDemoResource.Id,
    ///             HttpMethod = "GET",
    ///             Authorization = "NONE",
    ///         });
    ///         var myDemoIntegration = new Aws.ApiGateway.Integration("myDemoIntegration", new Aws.ApiGateway.IntegrationArgs
    ///         {
    ///             RestApi = myDemoAPI.Id,
    ///             ResourceId = myDemoResource.Id,
    ///             HttpMethod = myDemoMethod.HttpMethod,
    ///             Type = "MOCK",
    ///             CacheKeyParameters = 
    ///             {
    ///                 "method.request.path.param",
    ///             },
    ///             CacheNamespace = "foobar",
    ///             TimeoutMilliseconds = 29000,
    ///             RequestParameters = 
    ///             {
    ///                 { "integration.request.header.X-Authorization", "'static'" },
    ///             },
    ///             RequestTemplates = 
    ///             {
    ///                 { "application/xml", @"{
    ///    ""body"" : $input.json('$')
    /// }
    /// " },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## VPC Link
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var name = config.RequireObject&lt;dynamic&gt;("name");
    ///         var subnetId = config.RequireObject&lt;dynamic&gt;("subnetId");
    ///         var testLoadBalancer = new Aws.LB.LoadBalancer("testLoadBalancer", new Aws.LB.LoadBalancerArgs
    ///         {
    ///             Internal = true,
    ///             LoadBalancerType = "network",
    ///             Subnets = 
    ///             {
    ///                 subnetId,
    ///             },
    ///         });
    ///         var testVpcLink = new Aws.ApiGateway.VpcLink("testVpcLink", new Aws.ApiGateway.VpcLinkArgs
    ///         {
    ///             TargetArn = 
    ///             {
    ///                 testLoadBalancer.Arn,
    ///             },
    ///         });
    ///         var testRestApi = new Aws.ApiGateway.RestApi("testRestApi", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///         });
    ///         var testResource = new Aws.ApiGateway.Resource("testResource", new Aws.ApiGateway.ResourceArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             ParentId = testRestApi.RootResourceId,
    ///             PathPart = "test",
    ///         });
    ///         var testMethod = new Aws.ApiGateway.Method("testMethod", new Aws.ApiGateway.MethodArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             ResourceId = testResource.Id,
    ///             HttpMethod = "GET",
    ///             Authorization = "NONE",
    ///             RequestModels = 
    ///             {
    ///                 { "application/json", "Error" },
    ///             },
    ///         });
    ///         var testIntegration = new Aws.ApiGateway.Integration("testIntegration", new Aws.ApiGateway.IntegrationArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             ResourceId = testResource.Id,
    ///             HttpMethod = testMethod.HttpMethod,
    ///             RequestTemplates = 
    ///             {
    ///                 { "application/json", "" },
    ///                 { "application/xml", @"#set($inputRoot = $input.path('$'))
    /// { }" },
    ///             },
    ///             RequestParameters = 
    ///             {
    ///                 { "integration.request.header.X-Authorization", "'static'" },
    ///                 { "integration.request.header.X-Foo", "'Bar'" },
    ///             },
    ///             Type = "HTTP",
    ///             Uri = "https://www.google.de",
    ///             IntegrationHttpMethod = "GET",
    ///             PassthroughBehavior = "WHEN_NO_MATCH",
    ///             ContentHandling = "CONVERT_TO_TEXT",
    ///             ConnectionType = "VPC_LINK",
    ///             ConnectionId = testVpcLink.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_api_gateway_integration` can be imported using `REST-API-ID/RESOURCE-ID/HTTP-METHOD`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigateway/integration:Integration example 12345abcde/67890fghij/GET
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigateway/integration:Integration")]
    public partial class Integration : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of cache key parameters for the integration.
        /// </summary>
        [Output("cacheKeyParameters")]
        public Output<ImmutableArray<string>> CacheKeyParameters { get; private set; } = null!;

        /// <summary>
        /// The integration's cache namespace.
        /// </summary>
        [Output("cacheNamespace")]
        public Output<string> CacheNamespace { get; private set; } = null!;

        /// <summary>
        /// The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
        /// </summary>
        [Output("connectionId")]
        public Output<string?> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
        /// </summary>
        [Output("connectionType")]
        public Output<string?> ConnectionType { get; private set; } = null!;

        /// <summary>
        /// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
        /// </summary>
        [Output("contentHandling")]
        public Output<string?> ContentHandling { get; private set; } = null!;

        /// <summary>
        /// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
        /// </summary>
        [Output("credentials")]
        public Output<string?> Credentials { get; private set; } = null!;

        /// <summary>
        /// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
        /// when calling the associated resource.
        /// </summary>
        [Output("httpMethod")]
        public Output<string> HttpMethod { get; private set; } = null!;

        /// <summary>
        /// The integration HTTP method
        /// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
        /// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        /// Not all methods are compatible with all `AWS` integrations.
        /// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
        /// </summary>
        [Output("integrationHttpMethod")]
        public Output<string?> IntegrationHttpMethod { get; private set; } = null!;

        /// <summary>
        /// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
        /// </summary>
        [Output("passthroughBehavior")]
        public Output<string> PassthroughBehavior { get; private set; } = null!;

        /// <summary>
        /// A map of request query string parameters and headers that should be passed to the backend responder.
        /// For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
        /// </summary>
        [Output("requestParameters")]
        public Output<ImmutableDictionary<string, string>?> RequestParameters { get; private set; } = null!;

        /// <summary>
        /// A map of the integration's request templates.
        /// </summary>
        [Output("requestTemplates")]
        public Output<ImmutableDictionary<string, string>?> RequestTemplates { get; private set; } = null!;

        /// <summary>
        /// The API resource ID.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated REST API.
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;

        /// <summary>
        /// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
        /// </summary>
        [Output("timeoutMilliseconds")]
        public Output<int?> TimeoutMilliseconds { get; private set; } = null!;

        /// <summary>
        /// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        /// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
        /// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
        /// </summary>
        [Output("uri")]
        public Output<string?> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : Pulumi.ResourceArgs
    {
        [Input("cacheKeyParameters")]
        private InputList<string>? _cacheKeyParameters;

        /// <summary>
        /// A list of cache key parameters for the integration.
        /// </summary>
        public InputList<string> CacheKeyParameters
        {
            get => _cacheKeyParameters ?? (_cacheKeyParameters = new InputList<string>());
            set => _cacheKeyParameters = value;
        }

        /// <summary>
        /// The integration's cache namespace.
        /// </summary>
        [Input("cacheNamespace")]
        public Input<string>? CacheNamespace { get; set; }

        /// <summary>
        /// The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
        /// </summary>
        [Input("contentHandling")]
        public Input<string>? ContentHandling { get; set; }

        /// <summary>
        /// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        /// <summary>
        /// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
        /// when calling the associated resource.
        /// </summary>
        [Input("httpMethod", required: true)]
        public Input<string> HttpMethod { get; set; } = null!;

        /// <summary>
        /// The integration HTTP method
        /// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
        /// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        /// Not all methods are compatible with all `AWS` integrations.
        /// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
        /// </summary>
        [Input("integrationHttpMethod")]
        public Input<string>? IntegrationHttpMethod { get; set; }

        /// <summary>
        /// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
        /// </summary>
        [Input("passthroughBehavior")]
        public Input<string>? PassthroughBehavior { get; set; }

        [Input("requestParameters")]
        private InputMap<string>? _requestParameters;

        /// <summary>
        /// A map of request query string parameters and headers that should be passed to the backend responder.
        /// For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
        /// </summary>
        public InputMap<string> RequestParameters
        {
            get => _requestParameters ?? (_requestParameters = new InputMap<string>());
            set => _requestParameters = value;
        }

        [Input("requestTemplates")]
        private InputMap<string>? _requestTemplates;

        /// <summary>
        /// A map of the integration's request templates.
        /// </summary>
        public InputMap<string> RequestTemplates
        {
            get => _requestTemplates ?? (_requestTemplates = new InputMap<string>());
            set => _requestTemplates = value;
        }

        /// <summary>
        /// The API resource ID.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The ID of the associated REST API.
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        /// <summary>
        /// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
        /// </summary>
        [Input("timeoutMilliseconds")]
        public Input<int>? TimeoutMilliseconds { get; set; }

        /// <summary>
        /// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        /// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
        /// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public IntegrationArgs()
        {
        }
    }

    public sealed class IntegrationState : Pulumi.ResourceArgs
    {
        [Input("cacheKeyParameters")]
        private InputList<string>? _cacheKeyParameters;

        /// <summary>
        /// A list of cache key parameters for the integration.
        /// </summary>
        public InputList<string> CacheKeyParameters
        {
            get => _cacheKeyParameters ?? (_cacheKeyParameters = new InputList<string>());
            set => _cacheKeyParameters = value;
        }

        /// <summary>
        /// The integration's cache namespace.
        /// </summary>
        [Input("cacheNamespace")]
        public Input<string>? CacheNamespace { get; set; }

        /// <summary>
        /// The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
        /// </summary>
        [Input("connectionId")]
        public Input<string>? ConnectionId { get; set; }

        /// <summary>
        /// The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
        /// </summary>
        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        /// <summary>
        /// Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
        /// </summary>
        [Input("contentHandling")]
        public Input<string>? ContentHandling { get; set; }

        /// <summary>
        /// The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        /// <summary>
        /// The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
        /// when calling the associated resource.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// The integration HTTP method
        /// (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
        /// **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        /// Not all methods are compatible with all `AWS` integrations.
        /// e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
        /// </summary>
        [Input("integrationHttpMethod")]
        public Input<string>? IntegrationHttpMethod { get; set; }

        /// <summary>
        /// The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
        /// </summary>
        [Input("passthroughBehavior")]
        public Input<string>? PassthroughBehavior { get; set; }

        [Input("requestParameters")]
        private InputMap<string>? _requestParameters;

        /// <summary>
        /// A map of request query string parameters and headers that should be passed to the backend responder.
        /// For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
        /// </summary>
        public InputMap<string> RequestParameters
        {
            get => _requestParameters ?? (_requestParameters = new InputMap<string>());
            set => _requestParameters = value;
        }

        [Input("requestTemplates")]
        private InputMap<string>? _requestTemplates;

        /// <summary>
        /// A map of the integration's request templates.
        /// </summary>
        public InputMap<string> RequestTemplates
        {
            get => _requestTemplates ?? (_requestTemplates = new InputMap<string>());
            set => _requestTemplates = value;
        }

        /// <summary>
        /// The API resource ID.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The ID of the associated REST API.
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        /// <summary>
        /// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
        /// </summary>
        [Input("timeoutMilliseconds")]
        public Input<int>? TimeoutMilliseconds { get; set; }

        /// <summary>
        /// The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        /// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
        /// e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public IntegrationState()
        {
        }
    }
}
