// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS Config Aggregate Authorization
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cfg.AggregateAuthorization("example", {
 *     accountId: "123456789012",
 *     region: "eu-west-2",
 * });
 * ```
 *
 * ## Import
 *
 * Config aggregate authorizations can be imported using `account_id:region`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:cfg/aggregateAuthorization:AggregateAuthorization example 123456789012:us-east-1
 * ```
 */
export class AggregateAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing AggregateAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AggregateAuthorizationState, opts?: pulumi.CustomResourceOptions): AggregateAuthorization {
        return new AggregateAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/aggregateAuthorization:AggregateAuthorization';

    /**
     * Returns true if the given object is an instance of AggregateAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AggregateAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AggregateAuthorization.__pulumiType;
    }

    /**
     * Account ID
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The ARN of the authorization
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Region
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a AggregateAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AggregateAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AggregateAuthorizationArgs | AggregateAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AggregateAuthorizationState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as AggregateAuthorizationArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AggregateAuthorization.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AggregateAuthorization resources.
 */
export interface AggregateAuthorizationState {
    /**
     * Account ID
     */
    accountId?: pulumi.Input<string>;
    /**
     * The ARN of the authorization
     */
    arn?: pulumi.Input<string>;
    /**
     * Region
     */
    region?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a AggregateAuthorization resource.
 */
export interface AggregateAuthorizationArgs {
    /**
     * Account ID
     */
    accountId: pulumi.Input<string>;
    /**
     * Region
     */
    region: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
