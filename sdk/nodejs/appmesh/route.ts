// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an AWS App Mesh route resource.
 * 
 * ## Example Usage
 * 
 * ### HTTP Routing
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const serviceb = new aws.appmesh.Route("serviceb", {
 *     meshName: aws_appmesh_mesh_simple.id,
 *     name: "serviceB-route",
 *     spec: {
 *         httpRoute: {
 *             action: {
 *                 weightedTargets: [
 *                     {
 *                         virtualNode: aws_appmesh_virtual_node_serviceb1.name,
 *                         weight: 90,
 *                     },
 *                     {
 *                         virtualNode: aws_appmesh_virtual_node_serviceb2.name,
 *                         weight: 10,
 *                     },
 *                 ],
 *             },
 *             match: {
 *                 prefix: "/",
 *             },
 *         },
 *     },
 *     virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
 * });
 * ```
 * 
 * ### HTTP Header Routing
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const serviceb = new aws.appmesh.Route("serviceb", {
 *     meshName: aws_appmesh_mesh_simple.id,
 *     name: "serviceB-route",
 *     spec: {
 *         httpRoute: {
 *             action: {
 *                 weightedTargets: [{
 *                     virtualNode: aws_appmesh_virtual_node_serviceb.name,
 *                     weight: 100,
 *                 }],
 *             },
 *             match: {
 *                 headers: [{
 *                     match: {
 *                         prefix: "123",
 *                     },
 *                     name: "clientRequestId",
 *                 }],
 *                 method: "POST",
 *                 prefix: "/",
 *                 scheme: "https",
 *             },
 *         },
 *     },
 *     virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
 * });
 * ```
 * 
 * ### TCP Routing
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const serviceb = new aws.appmesh.Route("serviceb", {
 *     meshName: aws_appmesh_mesh_simple.id,
 *     name: "serviceB-route",
 *     spec: {
 *         tcpRoute: {
 *             action: {
 *                 weightedTargets: [{
 *                     virtualNode: aws_appmesh_virtual_node_serviceb1.name,
 *                     weight: 100,
 *                 }],
 *             },
 *         },
 *     },
 *     virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_route.html.markdown.
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appmesh/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * The ARN of the route.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The creation date of the route.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The last update date of the route.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * The name of the service mesh in which to create the route.
     */
    public readonly meshName!: pulumi.Output<string>;
    /**
     * The name to use for the route.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The route specification to apply.
     */
    public readonly spec!: pulumi.Output<outputs.appmesh.RouteSpec>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The name of the virtual router in which to create the route.
     */
    public readonly virtualRouterName!: pulumi.Output<string>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["meshName"] = state ? state.meshName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualRouterName"] = state ? state.virtualRouterName : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if (!args || args.meshName === undefined) {
                throw new Error("Missing required property 'meshName'");
            }
            if (!args || args.spec === undefined) {
                throw new Error("Missing required property 'spec'");
            }
            if (!args || args.virtualRouterName === undefined) {
                throw new Error("Missing required property 'virtualRouterName'");
            }
            inputs["meshName"] = args ? args.meshName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualRouterName"] = args ? args.virtualRouterName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["lastUpdatedDate"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    /**
     * The ARN of the route.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The creation date of the route.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The last update date of the route.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * The name of the service mesh in which to create the route.
     */
    readonly meshName?: pulumi.Input<string>;
    /**
     * The name to use for the route.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The route specification to apply.
     */
    readonly spec?: pulumi.Input<inputs.appmesh.RouteSpec>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the virtual router in which to create the route.
     */
    readonly virtualRouterName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * The name of the service mesh in which to create the route.
     */
    readonly meshName: pulumi.Input<string>;
    /**
     * The name to use for the route.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The route specification to apply.
     */
    readonly spec: pulumi.Input<inputs.appmesh.RouteSpec>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the virtual router in which to create the route.
     */
    readonly virtualRouterName: pulumi.Input<string>;
}
