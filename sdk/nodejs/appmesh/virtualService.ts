// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AWS App Mesh virtual service resource.
 * 
 * ## Example Usage
 * 
 * ### Virtual Node Provider
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const servicea = new aws.appmesh.VirtualService("servicea", {
 *     meshName: aws_appmesh_mesh_simple.id,
 *     spec: {
 *         provider: {
 *             virtualNode: {
 *                 virtualNodeName: aws_appmesh_virtual_node_serviceb1.name,
 *             },
 *         },
 *     },
 * });
 * ```
 * 
 * ### Virtual Router Provider
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const servicea = new aws.appmesh.VirtualService("servicea", {
 *     meshName: aws_appmesh_mesh_simple.id,
 *     spec: {
 *         provider: {
 *             virtualRouter: {
 *                 virtualRouterName: aws_appmesh_virtual_router_serviceb.name,
 *             },
 *         },
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_service.html.markdown.
 */
export class VirtualService extends pulumi.CustomResource {
    /**
     * Get an existing VirtualService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualServiceState, opts?: pulumi.CustomResourceOptions): VirtualService {
        return new VirtualService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appmesh/virtualService:VirtualService';

    /**
     * Returns true if the given object is an instance of VirtualService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualService.__pulumiType;
    }

    /**
     * The ARN of the virtual service.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The creation date of the virtual service.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The last update date of the virtual service.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * The name of the service mesh in which to create the virtual service.
     */
    public readonly meshName!: pulumi.Output<string>;
    /**
     * The name to use for the virtual service.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The virtual service specification to apply.
     */
    public readonly spec!: pulumi.Output<{ provider?: { virtualNode?: { virtualNodeName: string }, virtualRouter?: { virtualRouterName: string } } }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a VirtualService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualServiceArgs | VirtualServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualServiceState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["meshName"] = state ? state.meshName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VirtualServiceArgs | undefined;
            if (!args || args.meshName === undefined) {
                throw new Error("Missing required property 'meshName'");
            }
            if (!args || args.spec === undefined) {
                throw new Error("Missing required property 'spec'");
            }
            inputs["meshName"] = args ? args.meshName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["tags"] = args ? args.tags : undefined;
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
        super(VirtualService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualService resources.
 */
export interface VirtualServiceState {
    /**
     * The ARN of the virtual service.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The creation date of the virtual service.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The last update date of the virtual service.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * The name of the service mesh in which to create the virtual service.
     */
    readonly meshName?: pulumi.Input<string>;
    /**
     * The name to use for the virtual service.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The virtual service specification to apply.
     */
    readonly spec?: pulumi.Input<{ provider?: pulumi.Input<{ virtualNode?: pulumi.Input<{ virtualNodeName: pulumi.Input<string> }>, virtualRouter?: pulumi.Input<{ virtualRouterName: pulumi.Input<string> }> }> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a VirtualService resource.
 */
export interface VirtualServiceArgs {
    /**
     * The name of the service mesh in which to create the virtual service.
     */
    readonly meshName: pulumi.Input<string>;
    /**
     * The name to use for the virtual service.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The virtual service specification to apply.
     */
    readonly spec: pulumi.Input<{ provider?: pulumi.Input<{ virtualNode?: pulumi.Input<{ virtualNodeName: pulumi.Input<string> }>, virtualRouter?: pulumi.Input<{ virtualRouterName: pulumi.Input<string> }> }> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
