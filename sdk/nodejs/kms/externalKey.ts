// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a KMS Customer Master Key that uses external key material. To instead manage a KMS Customer Master Key where AWS automatically generates and potentially rotates key material, see the [`aws.kms.Key` resource](https://www.terraform.io/docs/providers/aws/r/kms_key.html).
 * 
 * > **Note:** All arguments including the key material will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.kms.ExternalKey("example", {
 *     description: "KMS EXTERNAL for AMI encryption",
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_external_key.html.markdown.
 */
export class ExternalKey extends pulumi.CustomResource {
    /**
     * Get an existing ExternalKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalKeyState, opts?: pulumi.CustomResourceOptions): ExternalKey {
        return new ExternalKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:kms/externalKey:ExternalKey';

    /**
     * Returns true if the given object is an instance of ExternalKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalKey.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the key.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
     */
    public readonly deletionWindowInDays!: pulumi.Output<number | undefined>;
    /**
     * Description of the key.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
     */
    public /*out*/ readonly expirationModel!: pulumi.Output<string>;
    /**
     * Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
     */
    public readonly keyMaterialBase64!: pulumi.Output<string | undefined>;
    /**
     * The state of the CMK.
     */
    public /*out*/ readonly keyState!: pulumi.Output<string>;
    /**
     * The cryptographic operations for which you can use the CMK.
     */
    public /*out*/ readonly keyUsage!: pulumi.Output<string>;
    /**
     * A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * A key-value map of tags to assign to the key.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    public readonly validTo!: pulumi.Output<string | undefined>;

    /**
     * Create a ExternalKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExternalKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalKeyArgs | ExternalKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ExternalKeyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["deletionWindowInDays"] = state ? state.deletionWindowInDays : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["expirationModel"] = state ? state.expirationModel : undefined;
            inputs["keyMaterialBase64"] = state ? state.keyMaterialBase64 : undefined;
            inputs["keyState"] = state ? state.keyState : undefined;
            inputs["keyUsage"] = state ? state.keyUsage : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["validTo"] = state ? state.validTo : undefined;
        } else {
            const args = argsOrState as ExternalKeyArgs | undefined;
            inputs["deletionWindowInDays"] = args ? args.deletionWindowInDays : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["keyMaterialBase64"] = args ? args.keyMaterialBase64 : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["validTo"] = args ? args.validTo : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["expirationModel"] = undefined /*out*/;
            inputs["keyState"] = undefined /*out*/;
            inputs["keyUsage"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ExternalKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalKey resources.
 */
export interface ExternalKeyState {
    /**
     * The Amazon Resource Name (ARN) of the key.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
     */
    readonly deletionWindowInDays?: pulumi.Input<number>;
    /**
     * Description of the key.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Whether the key material expires. Empty when pending key material import, otherwise `KEY_MATERIAL_EXPIRES` or `KEY_MATERIAL_DOES_NOT_EXPIRE`.
     */
    readonly expirationModel?: pulumi.Input<string>;
    /**
     * Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
     */
    readonly keyMaterialBase64?: pulumi.Input<string>;
    /**
     * The state of the CMK.
     */
    readonly keyState?: pulumi.Input<string>;
    /**
     * The cryptographic operations for which you can use the CMK.
     */
    readonly keyUsage?: pulumi.Input<string>;
    /**
     * A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * A key-value map of tags to assign to the key.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    readonly validTo?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalKey resource.
 */
export interface ExternalKeyArgs {
    /**
     * Duration in days after which the key is deleted after destruction of the resource. Must be between `7` and `30` days. Defaults to `30`.
     */
    readonly deletionWindowInDays?: pulumi.Input<number>;
    /**
     * Description of the key.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies whether the key is enabled. Keys pending import can only be `false`. Imported keys default to `true` unless expired.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Base64 encoded 256-bit symmetric encryption key material to import. The CMK is permanently associated with this key material. The same key material can be reimported, but you cannot import different key material.
     */
    readonly keyMaterialBase64?: pulumi.Input<string>;
    /**
     * A key policy JSON document. If you do not provide a key policy, AWS KMS attaches a default key policy to the CMK.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * A key-value map of tags to assign to the key.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Time at which the imported key material expires. When the key material expires, AWS KMS deletes the key material and the CMK becomes unusable. If not specified, key material does not expire. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
     */
    readonly validTo?: pulumi.Input<string>;
}
