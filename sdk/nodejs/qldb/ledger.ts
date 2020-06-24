// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AWS Quantum Ledger Database (QLDB) resource
 *
 * > **NOTE:** Deletion protection is enabled by default. To successfully delete this resource via this provider, `deletionProtection = false` must be applied before attempting deletion.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const sample_ledger = new aws.qldb.Ledger("sample-ledger", {});
 * ```
 */
export class Ledger extends pulumi.CustomResource {
    /**
     * Get an existing Ledger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LedgerState, opts?: pulumi.CustomResourceOptions): Ledger {
        return new Ledger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:qldb/ledger:Ledger';

    /**
     * Returns true if the given object is an instance of Ledger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ledger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ledger.__pulumiType;
    }

    /**
     * The ARN of the QLDB Ledger
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via this provider, this value must be configured to `false` and applied first before attempting deletion.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * The friendly name for the QLDB Ledger instance. This is atuo generated by default.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Ledger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LedgerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LedgerArgs | LedgerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LedgerState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as LedgerArgs | undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Ledger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ledger resources.
 */
export interface LedgerState {
    /**
     * The ARN of the QLDB Ledger
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via this provider, this value must be configured to `false` and applied first before attempting deletion.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * The friendly name for the QLDB Ledger instance. This is atuo generated by default.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Ledger resource.
 */
export interface LedgerArgs {
    /**
     * The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via this provider, this value must be configured to `false` and applied first before attempting deletion.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * The friendly name for the QLDB Ledger instance. This is atuo generated by default.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
