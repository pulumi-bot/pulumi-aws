// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Provides an API Gateway Method Settings, e.g. logging or monitoring.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const testRestApi = new aws.apigateway.RestApi("test", {
 *     description: "This is my API for demonstration purposes",
 * });
 * const testResource = new aws.apigateway.Resource("test", {
 *     parentId: testRestApi.rootResourceId,
 *     pathPart: "mytestresource",
 *     restApi: testRestApi.id,
 * });
 * const testMethod = new aws.apigateway.Method("test", {
 *     authorization: "NONE",
 *     httpMethod: "GET",
 *     resourceId: testResource.id,
 *     restApi: testRestApi.id,
 * });
 * const testIntegration = new aws.apigateway.Integration("test", {
 *     httpMethod: testMethod.httpMethod,
 *     requestTemplates: {
 *         "application/xml": `{
 *    "body" : $input.json('$')
 * }
 * `,
 *     },
 *     resourceId: testResource.id,
 *     restApi: testRestApi.id,
 *     type: "MOCK",
 * });
 * const testDeployment = new aws.apigateway.Deployment("test", {
 *     restApi: testRestApi.id,
 *     stageName: "dev",
 * }, {dependsOn: [testIntegration]});
 * const testStage = new aws.apigateway.Stage("test", {
 *     deployment: testDeployment.id,
 *     restApi: testRestApi.id,
 *     stageName: "prod",
 * });
 * const methodSettings = new aws.apigateway.MethodSettings("s", {
 *     methodPath: pulumi.interpolate`${testResource.pathPart}/${testMethod.httpMethod}`,
 *     restApi: testRestApi.id,
 *     settings: {
 *         loggingLevel: "INFO",
 *         metricsEnabled: true,
 *     },
 *     stageName: testStage.stageName,
 * });
 * ```
 */
export class MethodSettings extends pulumi.CustomResource {
    /**
     * Get an existing MethodSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MethodSettingsState, opts?: pulumi.CustomResourceOptions): MethodSettings {
        return new MethodSettings(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:apigateway/methodSettings:MethodSettings';

    /**
     * Returns true if the given object is an instance of MethodSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MethodSettings {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === MethodSettings.__pulumiType;
    }

    /**
     * Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*&#47;*` for overriding all methods in the stage.
     */
    public readonly methodPath!: pulumi.Output<string>;
    /**
     * The ID of the REST API
     */
    public readonly restApi!: pulumi.Output<RestApi>;
    /**
     * The settings block, see below.
     */
    public readonly settings!: pulumi.Output<{ cacheDataEncrypted?: boolean, cacheTtlInSeconds?: number, cachingEnabled?: boolean, dataTraceEnabled?: boolean, loggingLevel?: string, metricsEnabled?: boolean, requireAuthorizationForCacheControl?: boolean, throttlingBurstLimit?: number, throttlingRateLimit?: number, unauthorizedCacheControlHeaderStrategy?: string }>;
    /**
     * The name of the stage
     */
    public readonly stageName!: pulumi.Output<string>;

    /**
     * Create a MethodSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MethodSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MethodSettingsArgs | MethodSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MethodSettingsState | undefined;
            inputs["methodPath"] = state ? state.methodPath : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["settings"] = state ? state.settings : undefined;
            inputs["stageName"] = state ? state.stageName : undefined;
        } else {
            const args = argsOrState as MethodSettingsArgs | undefined;
            if (!args || args.methodPath === undefined) {
                throw new Error("Missing required property 'methodPath'");
            }
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            if (!args || args.settings === undefined) {
                throw new Error("Missing required property 'settings'");
            }
            if (!args || args.stageName === undefined) {
                throw new Error("Missing required property 'stageName'");
            }
            inputs["methodPath"] = args ? args.methodPath : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["stageName"] = args ? args.stageName : undefined;
        }
        super(MethodSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MethodSettings resources.
 */
export interface MethodSettingsState {
    /**
     * Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*&#47;*` for overriding all methods in the stage.
     */
    readonly methodPath?: pulumi.Input<string>;
    /**
     * The ID of the REST API
     */
    readonly restApi?: pulumi.Input<RestApi>;
    /**
     * The settings block, see below.
     */
    readonly settings?: pulumi.Input<{ cacheDataEncrypted?: pulumi.Input<boolean>, cacheTtlInSeconds?: pulumi.Input<number>, cachingEnabled?: pulumi.Input<boolean>, dataTraceEnabled?: pulumi.Input<boolean>, loggingLevel?: pulumi.Input<string>, metricsEnabled?: pulumi.Input<boolean>, requireAuthorizationForCacheControl?: pulumi.Input<boolean>, throttlingBurstLimit?: pulumi.Input<number>, throttlingRateLimit?: pulumi.Input<number>, unauthorizedCacheControlHeaderStrategy?: pulumi.Input<string> }>;
    /**
     * The name of the stage
     */
    readonly stageName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MethodSettings resource.
 */
export interface MethodSettingsArgs {
    /**
     * Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*&#47;*` for overriding all methods in the stage.
     */
    readonly methodPath: pulumi.Input<string>;
    /**
     * The ID of the REST API
     */
    readonly restApi: pulumi.Input<RestApi>;
    /**
     * The settings block, see below.
     */
    readonly settings: pulumi.Input<{ cacheDataEncrypted?: pulumi.Input<boolean>, cacheTtlInSeconds?: pulumi.Input<number>, cachingEnabled?: pulumi.Input<boolean>, dataTraceEnabled?: pulumi.Input<boolean>, loggingLevel?: pulumi.Input<string>, metricsEnabled?: pulumi.Input<boolean>, requireAuthorizationForCacheControl?: pulumi.Input<boolean>, throttlingBurstLimit?: pulumi.Input<number>, throttlingRateLimit?: pulumi.Input<number>, unauthorizedCacheControlHeaderStrategy?: pulumi.Input<string> }>;
    /**
     * The name of the stage
     */
    readonly stageName: pulumi.Input<string>;
}
