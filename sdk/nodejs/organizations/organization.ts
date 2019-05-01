// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an organization.
 * 
 * ## Example Usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const org = new aws.organizations.Organization("org", {
 *     awsServiceAccessPrincipals: [
 *         "cloudtrail.amazonaws.com",
 *         "config.amazonaws.com",
 *     ],
 *     featureSet: "ALL",
 * });
 * ```
 */
export class Organization extends pulumi.CustomResource {
    /**
     * Get an existing Organization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationState, opts?: pulumi.CustomResourceOptions): Organization {
        return new Organization(name, <any>state, { ...opts, id: id });
    }

    /**
     * ARN of the organization
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
     */
    public readonly awsServiceAccessPrincipals!: pulumi.Output<string[] | undefined>;
    /**
     * Specify "ALL" (default) or "CONSOLIDATED_BILLING".
     */
    public readonly featureSet!: pulumi.Output<string | undefined>;
    /**
     * ARN of the master account
     */
    public /*out*/ readonly masterAccountArn!: pulumi.Output<string>;
    /**
     * Email address of the master account
     */
    public /*out*/ readonly masterAccountEmail!: pulumi.Output<string>;
    /**
     * Identifier of the master account
     */
    public /*out*/ readonly masterAccountId!: pulumi.Output<string>;

    /**
     * Create a Organization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrganizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationArgs | OrganizationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsServiceAccessPrincipals"] = state ? state.awsServiceAccessPrincipals : undefined;
            inputs["featureSet"] = state ? state.featureSet : undefined;
            inputs["masterAccountArn"] = state ? state.masterAccountArn : undefined;
            inputs["masterAccountEmail"] = state ? state.masterAccountEmail : undefined;
            inputs["masterAccountId"] = state ? state.masterAccountId : undefined;
        } else {
            const args = argsOrState as OrganizationArgs | undefined;
            inputs["awsServiceAccessPrincipals"] = args ? args.awsServiceAccessPrincipals : undefined;
            inputs["featureSet"] = args ? args.featureSet : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["masterAccountArn"] = undefined /*out*/;
            inputs["masterAccountEmail"] = undefined /*out*/;
            inputs["masterAccountId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("aws:organizations/organization:Organization", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Organization resources.
 */
export interface OrganizationState {
    /**
     * ARN of the organization
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
     */
    readonly awsServiceAccessPrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify "ALL" (default) or "CONSOLIDATED_BILLING".
     */
    readonly featureSet?: pulumi.Input<string>;
    /**
     * ARN of the master account
     */
    readonly masterAccountArn?: pulumi.Input<string>;
    /**
     * Email address of the master account
     */
    readonly masterAccountEmail?: pulumi.Input<string>;
    /**
     * Identifier of the master account
     */
    readonly masterAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Organization resource.
 */
export interface OrganizationArgs {
    /**
     * List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `feature_set` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
     */
    readonly awsServiceAccessPrincipals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specify "ALL" (default) or "CONSOLIDATED_BILLING".
     */
    readonly featureSet?: pulumi.Input<string>;
}
