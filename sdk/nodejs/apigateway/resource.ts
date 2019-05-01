// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./restApi";

/**
 * Provides an API Gateway Resource.
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
 * ```
 */
export class Resource extends pulumi.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceState, opts?: pulumi.CustomResourceOptions): Resource {
        return new Resource(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ID of the parent API resource
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * The complete path for this API resource, including all parent paths.
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * The last path segment of this API resource.
     */
    public readonly pathPart!: pulumi.Output<string>;
    /**
     * The ID of the associated REST API
     */
    public readonly restApi!: pulumi.Output<RestApi>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceArgs | ResourceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceState | undefined;
            inputs["parentId"] = state ? state.parentId : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["pathPart"] = state ? state.pathPart : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
        } else {
            const args = argsOrState as ResourceArgs | undefined;
            if (!args || args.parentId === undefined) {
                throw new Error("Missing required property 'parentId'");
            }
            if (!args || args.pathPart === undefined) {
                throw new Error("Missing required property 'pathPart'");
            }
            if (!args || args.restApi === undefined) {
                throw new Error("Missing required property 'restApi'");
            }
            inputs["parentId"] = args ? args.parentId : undefined;
            inputs["pathPart"] = args ? args.pathPart : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["path"] = undefined /*out*/;
        }
        super("aws:apigateway/resource:Resource", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Resource resources.
 */
export interface ResourceState {
    /**
     * The ID of the parent API resource
     */
    readonly parentId?: pulumi.Input<string>;
    /**
     * The complete path for this API resource, including all parent paths.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The last path segment of this API resource.
     */
    readonly pathPart?: pulumi.Input<string>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi?: pulumi.Input<RestApi>;
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    /**
     * The ID of the parent API resource
     */
    readonly parentId: pulumi.Input<string>;
    /**
     * The last path segment of this API resource.
     */
    readonly pathPart: pulumi.Input<string>;
    /**
     * The ID of the associated REST API
     */
    readonly restApi: pulumi.Input<RestApi>;
}
