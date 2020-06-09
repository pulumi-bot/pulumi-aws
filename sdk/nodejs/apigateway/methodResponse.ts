// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Provides an HTTP Method Response for an API Gateway Resource.
 *
 * ## Example Usage
 *
 *
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
 * const response200 = new aws.apigateway.MethodResponse("response_200", {
 *     httpMethod: myDemoMethod.httpMethod,
 *     resourceId: myDemoResource.id,
 *     restApi: myDemoAPI.id,
 *     statusCode: "200",
 * });
 * ```
 */
export class MethodResponse extends pulumi.CustomResource {
    /**
     * Get an existing MethodResponse resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MethodResponseState, opts?: pulumi.CustomResourceOptions): MethodResponse {
        return new MethodResponse(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/methodResponse:MethodResponse';

    /**
     * Returns true if the given object is an instance of MethodResponse.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MethodResponse {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MethodResponse.__pulumiType;
    }

    /**
     * The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * The API resource ID
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * A map of the API models used for the response's content type
     */
    public readonly responseModels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of response parameters that can be sent to the caller.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
     * would define that the header `X-Some-Header` can be provided on the response.
     */
    public readonly responseParameters!: pulumi.Output<{[key: string]: boolean} | undefined>;
    /**
     * The ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * The HTTP status code
     */
    public readonly statusCode!: pulumi.Output<string>;

    /**
     * Create a MethodResponse resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MethodResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MethodResponseArgs | MethodResponseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MethodResponseState | undefined;
            inputs["httpMethod"] = state ? state.httpMethod : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["responseModels"] = state ? state.responseModels : undefined;
            inputs["responseParameters"] = state ? state.responseParameters : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["statusCode"] = state ? state.statusCode : undefined;
        } else {
            const args = argsOrState as MethodResponseArgs | undefined;
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
            inputs["httpMethod"] = args ? args.httpMethod : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["responseModels"] = args ? args.responseModels : undefined;
            inputs["responseParameters"] = args ? args.responseParameters : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["statusCode"] = args ? args.statusCode : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MethodResponse.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MethodResponse resources.
 */
export interface MethodResponseState {
    /**
     * The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * The API resource ID
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * A map of the API models used for the response's content type
     */
    readonly responseModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of response parameters that can be sent to the caller.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
     * would define that the header `X-Some-Header` can be provided on the response.
     */
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi?: pulumi.Input<string | RestApi>;
    /**
     * The HTTP status code
     */
    readonly statusCode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MethodResponse resource.
 */
export interface MethodResponseArgs {
    /**
     * The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
     */
    readonly httpMethod: pulumi.Input<string>;
    /**
     * The API resource ID
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * A map of the API models used for the response's content type
     */
    readonly responseModels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of response parameters that can be sent to the caller.
     * For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
     * would define that the header `X-Some-Header` can be provided on the response.
     */
    readonly responseParameters?: pulumi.Input<{[key: string]: pulumi.Input<boolean>}>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi: pulumi.Input<string | RestApi>;
    /**
     * The HTTP status code
     */
    readonly statusCode: pulumi.Input<string>;
}
