// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Subscribes to a Security Hub standard.
 * 
 * > **NOTE:** This AWS service is in Preview and may change before General Availability release. Backwards compatibility is not guaranteed between Terraform AWS Provider releases.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleAccount = new aws.securityhub.Account("example", {});
 * const exampleStandardsSubscription = new aws.securityhub.StandardsSubscription("example", {
 *     standardsArn: "arn:aws:securityhub:::ruleset/cis-aws-foundations-benchmark/v/1.2.0",
 * }, {dependsOn: [exampleAccount]});
 * ```
 */
export class StandardsSubscription extends pulumi.CustomResource {
    /**
     * Get an existing StandardsSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StandardsSubscriptionState, opts?: pulumi.CustomResourceOptions): StandardsSubscription {
        return new StandardsSubscription(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of a standard - see below.
     */
    public readonly standardsArn!: pulumi.Output<string>;

    /**
     * Create a StandardsSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StandardsSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StandardsSubscriptionArgs | StandardsSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StandardsSubscriptionState | undefined;
            inputs["standardsArn"] = state ? state.standardsArn : undefined;
        } else {
            const args = argsOrState as StandardsSubscriptionArgs | undefined;
            if (!args || args.standardsArn === undefined) {
                throw new Error("Missing required property 'standardsArn'");
            }
            inputs["standardsArn"] = args ? args.standardsArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("aws:securityhub/standardsSubscription:StandardsSubscription", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StandardsSubscription resources.
 */
export interface StandardsSubscriptionState {
    /**
     * The ARN of a standard - see below.
     */
    readonly standardsArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StandardsSubscription resource.
 */
export interface StandardsSubscriptionArgs {
    /**
     * The ARN of a standard - see below.
     */
    readonly standardsArn: pulumi.Input<string>;
}
