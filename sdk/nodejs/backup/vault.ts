// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an AWS Backup vault resource.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_vault.html.markdown.
 */
export class Vault extends pulumi.CustomResource {
    /**
     * Get an existing Vault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VaultState, opts?: pulumi.CustomResourceOptions): Vault {
        return new Vault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:backup/vault:Vault';

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

    /**
     * The ARN of the vault.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The server-side encryption key that is used to protect your backups.
     */
    public readonly kmsKeyArn!: pulumi.Output<string>;
    /**
     * Name of the backup vault to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of recovery points that are stored in a backup vault.
     */
    public /*out*/ readonly recoveryPoints!: pulumi.Output<number>;
    /**
     * Metadata that you can assign to help organize the resources that you create.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

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
            inputs["arn"] = state ? state.arn : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["recoveryPoints"] = state ? state.recoveryPoints : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VaultArgs | undefined;
            inputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["recoveryPoints"] = undefined /*out*/;
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
    /**
     * The ARN of the vault.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The server-side encryption key that is used to protect your backups.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
    /**
     * Name of the backup vault to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of recovery points that are stored in a backup vault.
     */
    readonly recoveryPoints?: pulumi.Input<number>;
    /**
     * Metadata that you can assign to help organize the resources that you create.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Vault resource.
 */
export interface VaultArgs {
    /**
     * The server-side encryption key that is used to protect your backups.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
    /**
     * Name of the backup vault to create.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Metadata that you can assign to help organize the resources that you create.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
