// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Vault extends pulumi.CustomResource {
    /**
     * Get an existing Vault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VaultState, opts?: pulumi.CustomResourceOptions): Vault {
        return new Vault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glacier/vault:Vault';

    /**
     * Returns true if the given object is an instance of Vault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vault.__pulumiType;
    }

    public readonly accessPolicy!: pulumi.Output<string | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly location!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly notifications!: pulumi.Output<outputs.glacier.VaultNotification[] | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Vault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VaultArgs | VaultState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VaultState | undefined;
            inputs["accessPolicy"] = state ? state.accessPolicy : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifications"] = state ? state.notifications : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VaultArgs | undefined;
            inputs["accessPolicy"] = args ? args.accessPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Vault.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vault resources.
 */
export interface VaultState {
    readonly accessPolicy?: pulumi.Input<string>;
    readonly arn?: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.glacier.VaultNotification>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Vault resource.
 */
export interface VaultArgs {
    readonly accessPolicy?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.glacier.VaultNotification>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
