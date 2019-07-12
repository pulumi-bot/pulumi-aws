// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_router.html.markdown.
 */
export class VirtualRouter extends pulumi.CustomResource {
    /**
     * Get an existing VirtualRouter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualRouterState, opts?: pulumi.CustomResourceOptions): VirtualRouter {
        return new VirtualRouter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appmesh/virtualRouter:VirtualRouter';

    /**
     * Returns true if the given object is an instance of VirtualRouter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualRouter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualRouter.__pulumiType;
    }

    /**
     * The ARN of the virtual router.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The creation date of the virtual router.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The last update date of the virtual router.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * The name of the service mesh in which to create the virtual router.
     */
    public readonly meshName!: pulumi.Output<string>;
    /**
     * The name to use for the virtual router.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The virtual router specification to apply.
     */
    public readonly spec!: pulumi.Output<{ listener: { portMapping: { port: number, protocol: string } } }>;

    /**
     * Create a VirtualRouter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualRouterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualRouterArgs | VirtualRouterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VirtualRouterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["meshName"] = state ? state.meshName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["spec"] = state ? state.spec : undefined;
        } else {
            const args = argsOrState as VirtualRouterArgs | undefined;
            if (!args || args.meshName === undefined) {
                throw new Error("Missing required property 'meshName'");
            }
            if (!args || args.spec === undefined) {
                throw new Error("Missing required property 'spec'");
            }
            inputs["meshName"] = args ? args.meshName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["lastUpdatedDate"] = undefined /*out*/;
        }
        super(VirtualRouter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualRouter resources.
 */
export interface VirtualRouterState {
    /**
     * The ARN of the virtual router.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The creation date of the virtual router.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The last update date of the virtual router.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * The name of the service mesh in which to create the virtual router.
     */
    readonly meshName?: pulumi.Input<string>;
    /**
     * The name to use for the virtual router.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The virtual router specification to apply.
     */
    readonly spec?: pulumi.Input<{ listener: pulumi.Input<{ portMapping: pulumi.Input<{ port: pulumi.Input<number>, protocol: pulumi.Input<string> }> }> }>;
}

/**
 * The set of arguments for constructing a VirtualRouter resource.
 */
export interface VirtualRouterArgs {
    /**
     * The name of the service mesh in which to create the virtual router.
     */
    readonly meshName: pulumi.Input<string>;
    /**
     * The name to use for the virtual router.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The virtual router specification to apply.
     */
    readonly spec: pulumi.Input<{ listener: pulumi.Input<{ portMapping: pulumi.Input<{ port: pulumi.Input<number>, protocol: pulumi.Input<string> }> }> }>;
}
