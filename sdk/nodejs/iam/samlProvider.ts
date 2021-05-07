// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM SAML provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * from "fs";
 *
 * const _default = new aws.iam.SamlProvider("default", {samlMetadataDocument: fs.readFileSync("saml-metadata.xml")});
 * ```
 *
 * ## Import
 *
 * IAM SAML Providers can be imported using the `arn`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:iam/samlProvider:SamlProvider default arn:aws:iam::123456789012:saml-provider/SAMLADFS
 * ```
 */
export class SamlProvider extends pulumi.CustomResource {
    /**
     * Get an existing SamlProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamlProviderState, opts?: pulumi.CustomResourceOptions): SamlProvider {
        return new SamlProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/samlProvider:SamlProvider';

    /**
     * Returns true if the given object is an instance of SamlProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamlProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SamlProvider.__pulumiType;
    }

    /**
     * The ARN assigned by AWS for this provider.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the provider to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An XML document generated by an identity provider that supports SAML 2.0.
     */
    public readonly samlMetadataDocument!: pulumi.Output<string>;
    /**
     * Map of resource tags for the IAM SAML provider. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The expiration date and time for the SAML provider in RFC1123 format, e.g. `Mon, 02 Jan 2006 15:04:05 MST`.
     */
    public /*out*/ readonly validUntil!: pulumi.Output<string>;

    /**
     * Create a SamlProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamlProviderArgs | SamlProviderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SamlProviderState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["samlMetadataDocument"] = state ? state.samlMetadataDocument : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["validUntil"] = state ? state.validUntil : undefined;
        } else {
            const args = argsOrState as SamlProviderArgs | undefined;
            if ((!args || args.samlMetadataDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlMetadataDocument'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["samlMetadataDocument"] = args ? args.samlMetadataDocument : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["validUntil"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SamlProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SamlProvider resources.
 */
export interface SamlProviderState {
    /**
     * The ARN assigned by AWS for this provider.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the provider to create.
     */
    name?: pulumi.Input<string>;
    /**
     * An XML document generated by an identity provider that supports SAML 2.0.
     */
    samlMetadataDocument?: pulumi.Input<string>;
    /**
     * Map of resource tags for the IAM SAML provider. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The expiration date and time for the SAML provider in RFC1123 format, e.g. `Mon, 02 Jan 2006 15:04:05 MST`.
     */
    validUntil?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SamlProvider resource.
 */
export interface SamlProviderArgs {
    /**
     * The name of the provider to create.
     */
    name?: pulumi.Input<string>;
    /**
     * An XML document generated by an identity provider that supports SAML 2.0.
     */
    samlMetadataDocument: pulumi.Input<string>;
    /**
     * Map of resource tags for the IAM SAML provider. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
