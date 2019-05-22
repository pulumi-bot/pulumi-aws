// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Provides an HTTP Method Integration for an API Gateway Integration.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const myDemoAPI = new aws.apigateway.RestApi("MyDemoAPI", {
 *     description: "This is my API for demonstration purposes",
 * });
 * const myDemoResource = new aws.apigateway.Resource("MyDemoResource", {
 *     parentId: myDemoAPI.rootResourceId,
 *     pathPart: "mydemoresource",
 *     restApi: myDemoAPI.id,
 * });
 * const myDemoMethod = new aws.apigateway.Method("MyDemoMethod", {
 *     authorization: "NONE",
 *     httpMethod: "GET",
 *     resourceId: myDemoResource.id,
 *     restApi: myDemoAPI.id,
 * });
 * const myDemoIntegration = new aws.apigateway.Integration("MyDemoIntegration", {
 *     cacheKeyParameters: ["method.request.path.param"],
 *     cacheNamespace: "foobar",
 *     httpMethod: myDemoMethod.httpMethod,
 *     requestParameters: {
 *         "integration.request.header.X-Authorization": "'static'",
 *     },
 *     // Transforms the incoming XML request to JSON
 *     requestTemplates: {
 *         "application/xml": `{
 *    "body" : $input.json('$')
 * }
 * `,
 *     },
 *     resourceId: myDemoResource.id,
 *     restApi: myDemoAPI.id,
 *     timeoutMilliseconds: 29000,
 *     type: "MOCK",
 * });
 * ```
 * 
 * ## Lambda integration
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const accountId = config.require("accountId");
 * // Variables
 * const myregion = config.require("myregion");
 * 
 * // API Gateway
 * const api = new aws.apigateway.RestApi("api", {});
 * // IAM
 * const role = new aws.iam.Role("role", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "lambda.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const resource = new aws.apigateway.Resource("resource", {
 *     parentId: api.rootResourceId,
 *     pathPart: "resource",
 *     restApi: api.id,
 * });
 * const method = new aws.apigateway.Method("method", {
 *     authorization: "NONE",
 *     httpMethod: "GET",
 *     resourceId: resource.id,
 *     restApi: api.id,
 * });
 * const lambda = new aws.lambda.Function("lambda", {
 *     code: new pulumi.asset.FileArchive("lambda.zip"),
 *     handler: "lambda.lambda_handler",
 *     role: role.arn,
 *     runtime: "python2.7",
 * });
 * const integration = new aws.apigateway.Integration("integration", {
 *     httpMethod: method.httpMethod,
 *     integrationHttpMethod: "POST",
 *     resourceId: resource.id,
 *     restApi: api.id,
 *     type: "AWS_PROXY",
 *     uri: pulumi.interpolate`arn:aws:apigateway:${myregion}:lambda:path/2015-03-31/functions/${lambda.arn}/invocations`,
 * });
 * // Lambda
 * const apigwLambda = new aws.lambda.Permission("apigw_lambda", {
 *     action: "lambda:InvokeFunction",
 *     function: lambda.functionName,
 *     principal: "apigateway.amazonaws.com",
 *     sourceArn: pulumi.interpolate`arn:aws:execute-api:${myregion}:${accountId}:${api.id}/*&#47;${method.httpMethod}/${resource.path}`,
 * });
 * ```
 * 
 * ## VPC Link
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const name = config.require("name");
 * const subnetId = config.require("subnetId");
 * 
 * const testRestApi = new aws.apigateway.RestApi("test", {});
 * const testResource = new aws.apigateway.Resource("test", {
 *     parentId: testRestApi.rootResourceId,
 *     pathPart: "test",
 *     restApi: testRestApi.id,
 * });
 * const testMethod = new aws.apigateway.Method("test", {
 *     authorization: "NONE",
 *     httpMethod: "GET",
 *     requestModels: {
 *         "application/json": "Error",
 *     },
 *     resourceId: testResource.id,
 *     restApi: testRestApi.id,
 * });
 * const testLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("test", {
 *     internal: true,
 *     loadBalancerType: "network",
 *     subnets: [subnetId],
 * });
 * const testVpcLink = new aws.apigateway.VpcLink("test", {
 *     targetArn: testLoadBalancer.arn,
 * });
 * const testIntegration = new aws.apigateway.Integration("test", {
 *     connectionId: testVpcLink.id,
 *     connectionType: "VPC_LINK",
 *     contentHandling: "CONVERT_TO_TEXT",
 *     httpMethod: testMethod.httpMethod,
 *     integrationHttpMethod: "GET",
 *     passthroughBehavior: "WHEN_NO_MATCH",
 *     requestParameters: {
 *         "integration.request.header.X-Authorization": "'static'",
 *         "integration.request.header.X-Foo": "'Bar'",
 *     },
 *     requestTemplates: {
 *         "application/json": "",
 *         "application/xml": `#set($inputRoot = $input.path('$'))
 * { }`,
 *     },
 *     resourceId: testResource.id,
 *     restApi: testRestApi.id,
 *     type: "HTTP",
 *     uri: "https://www.google.de",
 * });
 * ```
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationState, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, <any>state, { ...opts, id: id });
    }

    /**
     * A list of cache key parameters for the integration.
     */
    public readonly cacheKeyParameters!: pulumi.Output<string[] | undefined>;
    /**
     * The integration's cache namespace.
     */
    public readonly cacheNamespace!: pulumi.Output<string>;
    /**
     * The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
     */
    public readonly connectionId!: pulumi.Output<string | undefined>;
    /**
     * The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
     */
    public readonly contentHandling!: pulumi.Output<string | undefined>;
    /**
     * The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
     */
    public readonly credentials!: pulumi.Output<string | undefined>;
    /**
     * The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
     * when calling the associated resource.
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * The integration HTTP method
     * (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`) specifying how API Gateway will interact with the back end.
     * **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
     * Not all methods are compatible with all `AWS` integrations.
     * e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
     */
    public readonly integrationHttpMethod!: pulumi.Output<string | undefined>;
    /**
     * The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
     */
    public readonly passthroughBehavior!: pulumi.Output<string>;
    /**
     * A map of request query string parameters and headers that should be passed to the backend responder.
     * For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
     */
    public readonly requestParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of the integration's request templates.
     */
    public readonly requestTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The API resource ID.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The ID of the associated REST API.
     */
    public readonly restApi!: pulumi.Output<RestApi>;
    /**
     * Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
     */
    public readonly timeoutMilliseconds!: pulumi.Output<number | undefined>;
    /**
     * The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
     * For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
     * e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`
     */
    public readonly uri!: pulumi.Output<string | undefined>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            inputs["cacheKeyParameters"] = state ? state.cacheKeyParameters : undefined;
            inputs["cacheNamespace"] = state ? state.cacheNamespace : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["connectionType"] = state ? state.connectionType : undefined;
            inputs["contentHandling"] = state ? state.contentHandling : undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["httpMethod"] = state ? state.httpMethod : undefined;
            inputs["integrationHttpMethod"] = state ? state.integrationHttpMethod : undefined;
            inputs["passthroughBehavior"] = state ? state.passthroughBehavior : undefined;
            inputs["requestParameters"] = state ? state.requestParameters : undefined;
            inputs["requestTemplates"] = state ? state.requestTemplates : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["timeoutMilliseconds"] = state ? state.timeoutMilliseconds : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if (!args || args.httpMethod === undefined) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["cacheKeyParameters"] = args ? args.cacheKeyParameters : undefined;
            inputs["cacheNamespace"] = args ? args.cacheNamespace : undefined;
            inputs["connectionId"] = args ? args.connectionId : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["contentHandling"] = args ? args.contentHandling : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["httpMethod"] = args ? args.httpMethod : undefined;
            inputs["integrationHttpMethod"] = args ? args.integrationHttpMethod : undefined;
            inputs["passthroughBehavior"] = args ? args.passthroughBehavior : undefined;
            inputs["requestParameters"] = args ? args.requestParameters : undefined;
            inputs["requestTemplates"] = args ? args.requestTemplates : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["timeoutMilliseconds"] = args ? args.timeoutMilliseconds : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uri"] = args ? args.uri : undefined;
        }
        super("aws:apigateway/integration:Integration", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * A list of cache key parameters for the integration.
     */
    readonly cacheKeyParameters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The integration's cache namespace.
     */
    readonly cacheNamespace?: pulumi.Input<string>;
    /**
     * The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
     */
    readonly connectionType?: pulumi.Input<string>;
    /**
     * Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
     */
    readonly contentHandling?: pulumi.Input<string>;
    /**
     * The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
     */
    readonly credentials?: pulumi.Input<string>;
    /**
     * The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
     * when calling the associated resource.
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * The integration HTTP method
     * (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`) specifying how API Gateway will interact with the back end.
     * **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
     * Not all methods are compatible with all `AWS` integrations.
     * e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
     */
    readonly integrationHttpMethod?: pulumi.Input<string>;
    /**
     * The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
     */
    readonly passthroughBehavior?: pulumi.Input<string>;
    /**
     * A map of request query string parameters and headers that should be passed to the backend responder.
     * For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
     */
    readonly requestParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of the integration's request templates.
     */
    readonly requestTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The API resource ID.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The ID of the associated REST API.
     */
    readonly restApi?: pulumi.Input<RestApi>;
    /**
     * Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
     */
    readonly timeoutMilliseconds?: pulumi.Input<number>;
    /**
     * The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
     * For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
     * e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`
     */
    readonly uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * A list of cache key parameters for the integration.
     */
    readonly cacheKeyParameters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The integration's cache namespace.
     */
    readonly cacheNamespace?: pulumi.Input<string>;
    /**
     * The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
     */
    readonly connectionType?: pulumi.Input<string>;
    /**
     * Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
     */
    readonly contentHandling?: pulumi.Input<string>;
    /**
     * The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
     */
    readonly credentials?: pulumi.Input<string>;
    /**
     * The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
     * when calling the associated resource.
     */
    readonly httpMethod: pulumi.Input<string>;
    /**
     * The integration HTTP method
     * (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`) specifying how API Gateway will interact with the back end.
     * **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
     * Not all methods are compatible with all `AWS` integrations.
     * e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
     */
    readonly integrationHttpMethod?: pulumi.Input<string>;
    /**
     * The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
     */
    readonly passthroughBehavior?: pulumi.Input<string>;
    /**
     * A map of request query string parameters and headers that should be passed to the backend responder.
     * For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
     */
    readonly requestParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of the integration's request templates.
     */
    readonly requestTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The API resource ID.
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The ID of the associated REST API.
     */
    readonly restApi: pulumi.Input<RestApi>;
    /**
     * Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
     */
    readonly timeoutMilliseconds?: pulumi.Input<number>;
    /**
     * The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
     */
    readonly type: pulumi.Input<string>;
    /**
     * The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
     * For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
     * e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`
     */
    readonly uri?: pulumi.Input<string>;
}
