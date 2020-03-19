// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Provides an HTTP Method Integration Response for an API Gateway Resource.
 * 
 * > **Note:** Depends on having `aws.apigateway.Integration` inside your rest api. To ensure this
 * you might need to add an explicit `dependsOn` for clean runs.
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
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
 *     httpMethod: myDemoMethod.httpMethod,
 *     resourceId: myDemoResource.id,
 *     restApi: myDemoAPI.id,
 *     type: "MOCK",
 * });
 * const response200 = new aws.apigateway.MethodResponse("response200", {
 *     httpMethod: myDemoMethod.httpMethod,
 *     resourceId: myDemoResource.id,
 *     restApi: myDemoAPI.id,
 *     statusCode: "200",
 * });
 * const myDemoIntegrationResponse = new aws.apigateway.IntegrationResponse("MyDemoIntegrationResponse", {
 *     httpMethod: myDemoMethod.httpMethod,
 *     resourceId: myDemoResource.id,
 *     // Transforms the backend JSON response to XML
 *     responseTemplates: {
 *         "application/xml": `#set($inputRoot = $input.path('$'))
 * <?xml version="1.0" encoding="UTF-8"?>
 * <message>
 *     $inputRoot.body
 * </message>
 * `,
 *     },
 *     restApi: myDemoAPI.id,
 *     statusCode: response200.statusCode,
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_integration_response.html.markdown.
 */
export class IntegrationResponse extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationResponseState, opts?: pulumi.CustomResourceOptions): IntegrationResponse {
        return new IntegrationResponse(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/integrationResponse:IntegrationResponse';

    /**
     * Returns true if the given object is an instance of IntegrationResponse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationResponse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationResponse.__pulumiType;
    }

    /**
     * Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
     */
    public readonly contentHandling!: pulumi.Output<string | undefined>;
    /**
     * The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * The API resource ID
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * A map of response parameters that can be read from the backend response.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
     */
    public readonly responseParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map specifying the templates used to transform the integration response body
     */
    public readonly responseTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * Specifies the regular expression pattern used to choose
     * an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
     * If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
     * For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
     */
    public readonly selectionPattern!: pulumi.Output<string | undefined>;
    /**
     * The HTTP status code
     */
    public readonly statusCode!: pulumi.Output<string>;

    /**
     * Create a IntegrationResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationResponseArgs | IntegrationResponseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationResponseState | undefined;
            inputs["contentHandling"] = state ? state.contentHandling : undefined;
            inputs["httpMethod"] = state ? state.httpMethod : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["responseParameters"] = state ? state.responseParameters : undefined;
            inputs["responseTemplates"] = state ? state.responseTemplates : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["selectionPattern"] = state ? state.selectionPattern : undefined;
            inputs["statusCode"] = state ? state.statusCode : undefined;
        } else {
            const args = argsOrState as IntegrationResponseArgs | undefined;
            if (!args || args.httpMethod === undefined) {
                throw new Error("Missing required property 'httpMethod'");
            }
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            if (!args || args.statusCode === undefined) {
                throw new Error("Missing required property 'statusCode'");
            }
            inputs["contentHandling"] = args ? args.contentHandling : undefined;
            inputs["httpMethod"] = args ? args.httpMethod : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["responseParameters"] = args ? args.responseParameters : undefined;
            inputs["responseTemplates"] = args ? args.responseTemplates : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["selectionPattern"] = args ? args.selectionPattern : undefined;
            inputs["statusCode"] = args ? args.statusCode : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IntegrationResponse.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationResponse resources.
 */
export interface IntegrationResponseState {
    /**
     * Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
     */
    readonly contentHandling?: pulumi.Input<string>;
    /**
     * The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * The API resource ID
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * A map of response parameters that can be read from the backend response.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
     */
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map specifying the templates used to transform the integration response body
     */
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi?: pulumi.Input<string | RestApi>;
    /**
     * Specifies the regular expression pattern used to choose
     * an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
     * If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
     * For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
     */
    readonly selectionPattern?: pulumi.Input<string>;
    /**
     * The HTTP status code
     */
    readonly statusCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationResponse resource.
 */
export interface IntegrationResponseArgs {
    /**
     * Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
     */
    readonly contentHandling?: pulumi.Input<string>;
    /**
     * The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    readonly httpMethod: pulumi.Input<string>;
    /**
     * The API resource ID
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * A map of response parameters that can be read from the backend response.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
     */
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map specifying the templates used to transform the integration response body
     */
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi: pulumi.Input<string | RestApi>;
    /**
     * Specifies the regular expression pattern used to choose
     * an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
     * If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
     * For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
     */
    readonly selectionPattern?: pulumi.Input<string>;
    /**
     * The HTTP status code
     */
    readonly statusCode: pulumi.Input<string>;
}
