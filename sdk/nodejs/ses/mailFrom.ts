// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SES domain MAIL FROM resource.
 *
 * > **NOTE:** For the MAIL FROM domain to be fully usable, this resource should be paired with the `aws.ses.DomainIdentity` resource. To validate the MAIL FROM domain, a DNS MX record is required. To pass SPF checks, a DNS TXT record may also be required. See the [Amazon SES MAIL FROM documentation](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-set.html) for more information.
 */
export class MailFrom extends pulumi.CustomResource {
    /**
     * Get an existing MailFrom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MailFromState, opts?: pulumi.CustomResourceOptions): MailFrom {
        return new MailFrom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ses/mailFrom:MailFrom';

    /**
     * Returns true if the given object is an instance of MailFrom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MailFrom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MailFrom.__pulumiType;
    }

    /**
     * The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
     */
    public readonly behaviorOnMxFailure!: pulumi.Output<string | undefined>;
    /**
     * Verified domain name to generate DKIM tokens for.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
     */
    public readonly mailFromDomain!: pulumi.Output<string>;

    /**
     * Create a MailFrom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MailFromArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MailFromArgs | MailFromState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MailFromState | undefined;
            inputs["behaviorOnMxFailure"] = state ? state.behaviorOnMxFailure : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["mailFromDomain"] = state ? state.mailFromDomain : undefined;
        } else {
            const args = argsOrState as MailFromArgs | undefined;
            if (!args || args.domain === undefined) {
                throw new Error("Missing required property 'domain'");
            }
            if (!args || args.mailFromDomain === undefined) {
                throw new Error("Missing required property 'mailFromDomain'");
            }
            inputs["behaviorOnMxFailure"] = args ? args.behaviorOnMxFailure : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["mailFromDomain"] = args ? args.mailFromDomain : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MailFrom.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MailFrom resources.
 */
export interface MailFromState {
    /**
     * The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
     */
    readonly behaviorOnMxFailure?: pulumi.Input<string>;
    /**
     * Verified domain name to generate DKIM tokens for.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
     */
    readonly mailFromDomain?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MailFrom resource.
 */
export interface MailFromArgs {
    /**
     * The action that you want Amazon SES to take if it cannot successfully read the required MX record when you send an email. Defaults to `UseDefaultValue`. See the [SES API documentation](https://docs.aws.amazon.com/ses/latest/APIReference/API_SetIdentityMailFromDomain.html) for more information.
     */
    readonly behaviorOnMxFailure?: pulumi.Input<string>;
    /**
     * Verified domain name to generate DKIM tokens for.
     */
    readonly domain: pulumi.Input<string>;
    /**
     * Subdomain (of above domain) which is to be used as MAIL FROM address (Required for DMARC validation)
     */
    readonly mailFromDomain: pulumi.Input<string>;
}
