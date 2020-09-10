// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Mesh extends pulumi.CustomResource {
    /**
     * Get an existing Mesh resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly spec!: pulumi.Output<outputs.appmesh.MeshSpec | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

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
    readonly arn?: pulumi.Input<string>;
    readonly createdDate?: pulumi.Input<string>;
    readonly lastUpdatedDate?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly spec?: pulumi.Input<inputs.appmesh.MeshSpec>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Mesh resource.
 */
export interface MeshArgs {
    readonly name?: pulumi.Input<string>;
    readonly spec?: pulumi.Input<inputs.appmesh.MeshSpec>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
