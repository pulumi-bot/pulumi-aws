// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality
 *
 * > **NOTE:** When removing a Glacier Vault, the Vault must be empty.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const awsSnsTopic = new aws.sns.Topic("awsSnsTopic", {});
 * const myArchive = new aws.glacier.Vault("myArchive", {
 *     notification: {
 *         snsTopic: awsSnsTopic.arn,
 *         events: [
 *             "ArchiveRetrievalCompleted",
 *             "InventoryRetrievalCompleted",
 *         ],
 *     },
 *     accessPolicy: `{
 *     "Version":"2012-10-17",
 *     "Statement":[
 *        {
 *           "Sid": "add-read-only-perm",
 *           "Principal": "*",
 *           "Effect": "Allow",
 *           "Action": [
 *              "glacier:InitiateJob",
 *              "glacier:GetJobOutput"
 *           ],
 *           "Resource": "arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive"
 *        }
 *     ]
 * }
 * `,
 *     tags: {
 *         Test: "MyArchive",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Glacier Vaults can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glacier/vault:Vault archive my_archive
 * ```
 */
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

    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    public readonly accessPolicy!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the vault.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The URI of the vault that was created.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    public readonly notification!: pulumi.Output<outputs.glacier.VaultNotification | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VaultState | undefined;
            inputs["accessPolicy"] = state ? state.accessPolicy : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notification"] = state ? state.notification : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as VaultArgs | undefined;
            inputs["accessPolicy"] = args ? args.accessPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notification"] = args ? args.notification : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Vault.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vault resources.
 */
export interface VaultState {
    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    accessPolicy?: pulumi.Input<string>;
    /**
     * The ARN of the vault.
     */
    arn?: pulumi.Input<string>;
    /**
     * The URI of the vault that was created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    name?: pulumi.Input<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    notification?: pulumi.Input<inputs.glacier.VaultNotification>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Vault resource.
 */
export interface VaultArgs {
    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    accessPolicy?: pulumi.Input<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    name?: pulumi.Input<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    notification?: pulumi.Input<inputs.glacier.VaultNotification>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
