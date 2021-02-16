// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an API Gateway Gateway Response for a REST API Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.apigateway.RestApi("main", {});
 * const test = new aws.apigateway.Response("test", {
 *     restApiId: main.id,
 *     statusCode: "401",
 *     responseType: "UNAUTHORIZED",
 *     responseTemplates: {
 *         "application/json": `{"message":$context.error.messageString}`,
 *     },
 *     responseParameters: {
 *         "gatewayresponse.header.Authorization": "'Basic'",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * `aws_api_gateway_gateway_response` can be imported using `REST-API-ID/RESPONSE-TYPE`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:apigateway/response:Response example 12345abcde/UNAUTHORIZED
 * ```
 */
export class Response extends pulumi.CustomResource {
    /**
     * Get an existing Response resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResponseState, opts?: pulumi.CustomResourceOptions): Response {
        return new Response(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/response:Response';

    /**
     * Returns true if the given object is an instance of Response.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Response {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Response.__pulumiType;
    }

    /**
     * A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
     */
    public readonly responseParameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map specifying the templates used to transform the response body.
     */
    public readonly responseTemplates!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The response type of the associated GatewayResponse.
     */
    public readonly responseType!: pulumi.Output<string>;
    /**
     * The string identifier of the associated REST API.
     */
    public readonly restApiId!: pulumi.Output<string>;
    /**
     * The HTTP status code of the Gateway Response.
     */
    public readonly statusCode!: pulumi.Output<string | undefined>;

    /**
     * Create a Response resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResponseArgs | ResponseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResponseState | undefined;
            inputs["responseParameters"] = state ? state.responseParameters : undefined;
            inputs["responseTemplates"] = state ? state.responseTemplates : undefined;
            inputs["responseType"] = state ? state.responseType : undefined;
            inputs["restApiId"] = state ? state.restApiId : undefined;
            inputs["statusCode"] = state ? state.statusCode : undefined;
        } else {
            const args = argsOrState as ResponseArgs | undefined;
            if ((!args || args.responseType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'responseType'");
            }
            if ((!args || args.restApiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApiId'");
            }
            inputs["responseParameters"] = args ? args.responseParameters : undefined;
            inputs["responseTemplates"] = args ? args.responseTemplates : undefined;
            inputs["responseType"] = args ? args.responseType : undefined;
            inputs["restApiId"] = args ? args.restApiId : undefined;
            inputs["statusCode"] = args ? args.statusCode : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Response.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Response resources.
 */
export interface ResponseState {
    /**
     * A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
     */
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map specifying the templates used to transform the response body.
     */
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The response type of the associated GatewayResponse.
     */
    readonly responseType?: pulumi.Input<string>;
    /**
     * The string identifier of the associated REST API.
     */
    readonly restApiId?: pulumi.Input<string>;
    /**
     * The HTTP status code of the Gateway Response.
     */
    readonly statusCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Response resource.
 */
export interface ResponseArgs {
    /**
     * A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
     */
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map specifying the templates used to transform the response body.
     */
    readonly responseTemplates?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The response type of the associated GatewayResponse.
     */
    readonly responseType: pulumi.Input<string>;
    /**
     * The string identifier of the associated REST API.
     */
    readonly restApiId: pulumi.Input<string>;
    /**
     * The HTTP status code of the Gateway Response.
     */
    readonly statusCode?: pulumi.Input<string>;
}
