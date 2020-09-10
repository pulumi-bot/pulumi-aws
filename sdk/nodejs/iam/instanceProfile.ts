// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Role} from "./index";

export class InstanceProfile extends pulumi.CustomResource {
    /**
     * Get an existing InstanceProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceProfileState, opts?: pulumi.CustomResourceOptions): InstanceProfile {
        return new InstanceProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/instanceProfile:InstanceProfile';

    /**
     * Returns true if the given object is an instance of InstanceProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceProfile.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createDate!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly path!: pulumi.Output<string | undefined>;
    public readonly role!: pulumi.Output<string | undefined>;
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a InstanceProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceProfileArgs | InstanceProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceProfileState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createDate"] = state ? state.createDate : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
        } else {
            const args = argsOrState as InstanceProfileArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createDate"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceProfile resources.
 */
export interface InstanceProfileState {
    readonly arn?: pulumi.Input<string>;
    readonly createDate?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    readonly role?: pulumi.Input<string | Role>;
    readonly uniqueId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceProfile resource.
 */
export interface InstanceProfileArgs {
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly path?: pulumi.Input<string>;
    readonly role?: pulumi.Input<string | Role>;
}
