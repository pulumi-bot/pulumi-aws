// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a domain resource for the specified domain (e.g., example.com).
 * You cannot register a new domain name using Lightsail. You must register
 * a domain name using Amazon Route 53 or another domain name registrar.
 * If you have already registered your domain, you can enter its name in
 * this parameter to manage the DNS records for that domain.
 * 
 * > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
 * 
 * ## Example Usage, creating a new domain
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_lightsail_domain_domain_test = new aws.lightsail.Domain("domain_test", {
 *     domainName: "mydomain.com",
 * });
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the Lightsail domain
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The name of the Lightsail domain to manage
     */
    public readonly domainName: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DomainState = argsOrState as DomainState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if (!args || args.domainName === undefined) {
                throw new Error("Missing required property 'domainName'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:lightsail/domain:Domain", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * The ARN of the Lightsail domain
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the Lightsail domain to manage
     */
    readonly domainName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The name of the Lightsail domain to manage
     */
    readonly domainName: pulumi.Input<string>;
}
