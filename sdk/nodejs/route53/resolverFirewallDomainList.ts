// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Route 53 Resolver DNS Firewall domain list resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.route53.ResolverFirewallDomainList("example", {});
 * ```
 *
 * ## Import
 *
 *  Route 53 Resolver DNS Firewall domain lists can be imported using the Route 53 Resolver DNS Firewall domain list ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList example rslvr-fdl-0123456789abcdef
 * ```
 */
export class ResolverFirewallDomainList extends pulumi.CustomResource {
    /**
     * Get an existing ResolverFirewallDomainList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResolverFirewallDomainListState, opts?: pulumi.CustomResourceOptions): ResolverFirewallDomainList {
        return new ResolverFirewallDomainList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/resolverFirewallDomainList:ResolverFirewallDomainList';

    /**
     * Returns true if the given object is an instance of ResolverFirewallDomainList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverFirewallDomainList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverFirewallDomainList.__pulumiType;
    }

    /**
     * The ARN (Amazon Resource Name) of the domain list.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A array of domains for the firewall domain list.
     */
    public readonly domains!: pulumi.Output<string[] | undefined>;
    /**
     * A name that lets you identify the domain list, to manage and use it.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. f configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a ResolverFirewallDomainList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ResolverFirewallDomainListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResolverFirewallDomainListArgs | ResolverFirewallDomainListState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResolverFirewallDomainListState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domains"] = state ? state.domains : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as ResolverFirewallDomainListArgs | undefined;
            inputs["domains"] = args ? args.domains : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResolverFirewallDomainList.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResolverFirewallDomainList resources.
 */
export interface ResolverFirewallDomainListState {
    /**
     * The ARN (Amazon Resource Name) of the domain list.
     */
    arn?: pulumi.Input<string>;
    /**
     * A array of domains for the firewall domain list.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A name that lets you identify the domain list, to manage and use it.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. f configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ResolverFirewallDomainList resource.
 */
export interface ResolverFirewallDomainListArgs {
    /**
     * A array of domains for the firewall domain list.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A name that lets you identify the domain list, to manage and use it.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. f configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
