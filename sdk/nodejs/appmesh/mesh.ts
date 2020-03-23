// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an AWS App Mesh service mesh resource.
 * 
 * {{% examples %}}
 * ## Example Usage
 * 
 * {{% example %}}
 * ### Basic
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const simple = new aws.appmesh.Mesh("simple", {});
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### Egress Filter
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const simple = new aws.appmesh.Mesh("simple", {
 *     spec: {
 *         egressFilter: {
 *             type: "ALLOW_ALL",
 *         },
 *     },
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_mesh.html.markdown.
 */
export class Mesh extends pulumi.CustomResource {
    /**
     * Get an existing Mesh resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MeshState, opts?: pulumi.CustomResourceOptions): Mesh {
        return new Mesh(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appmesh/mesh:Mesh';

    /**
     * Returns true if the given object is an instance of Mesh.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Mesh {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Mesh.__pulumiType;
    }

    /**
     * The ARN of the service mesh.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The creation date of the service mesh.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The last update date of the service mesh.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * The name to use for the service mesh.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The service mesh specification to apply.
     */
    public readonly spec!: pulumi.Output<outputs.appmesh.MeshSpec | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Mesh resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MeshArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MeshArgs | MeshState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MeshState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as MeshArgs | undefined;
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
        super(Mesh.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Mesh resources.
 */
export interface MeshState {
    /**
     * The ARN of the service mesh.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The creation date of the service mesh.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The last update date of the service mesh.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * The name to use for the service mesh.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The service mesh specification to apply.
     */
    readonly spec?: pulumi.Input<inputs.appmesh.MeshSpec>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Mesh resource.
 */
export interface MeshArgs {
    /**
     * The name to use for the service mesh.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The service mesh specification to apply.
     */
    readonly spec?: pulumi.Input<inputs.appmesh.MeshSpec>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
