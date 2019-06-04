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
 */
export class AggregateAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing AggregateAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AggregateAuthorizationState, opts?: pulumi.CustomResourceOptions): AggregateAuthorization {
        return new AggregateAuthorization(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:cfg/aggregateAuthorization:AggregateAuthorization';

    /**
     * Returns true if the given object is an instance of AggregateAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AggregateAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === AggregateAuthorization.__pulumiType;
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
     * Create a AggregateAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AggregateAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AggregateAuthorizationArgs | AggregateAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AggregateAuthorizationState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as AggregateAuthorizationArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["arn"] = undefined /*out*/;
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
    readonly accountId?: pulumi.Input<string>;
    /**
     * The ARN of the authorization
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Region
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AggregateAuthorization resource.
 */
export interface AggregateAuthorizationArgs {
    /**
     * Account ID
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * Region
     */
    readonly region: pulumi.Input<string>;
}
