// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Route 53 Resolver DNSSEC config resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
 *     cidrBlock: "10.0.0.0/16",
 *     enableDnsSupport: true,
 *     enableDnsHostnames: true,
 * });
 * const exampleResolverDnsSecConfig = new aws.route53.ResolverDnsSecConfig("exampleResolverDnsSecConfig", {resourceId: exampleVpc.id});
 * ```
 *
 * ## Import
 *
 *  Route 53 Resolver DNSSEC configs can be imported using the Route 53 Resolver DNSSEC config ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig example rdsc-be1866ecc1683e95
 * ```
 */
export class ResolverDnsSecConfig extends pulumi.CustomResource {
    /**
     * Get an existing ResolverDnsSecConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResolverDnsSecConfigState, opts?: pulumi.CustomResourceOptions): ResolverDnsSecConfig {
        return new ResolverDnsSecConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig';

    /**
     * Returns true if the given object is an instance of ResolverDnsSecConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverDnsSecConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverDnsSecConfig.__pulumiType;
    }

    /**
     * The ARN for a configuration for DNSSEC validation.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
     */
    public /*out*/ readonly validationStatus!: pulumi.Output<string>;

    /**
     * Create a ResolverDnsSecConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResolverDnsSecConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResolverDnsSecConfigArgs | ResolverDnsSecConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResolverDnsSecConfigState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["validationStatus"] = state ? state.validationStatus : undefined;
        } else {
            const args = argsOrState as ResolverDnsSecConfigArgs | undefined;
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["validationStatus"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResolverDnsSecConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResolverDnsSecConfig resources.
 */
export interface ResolverDnsSecConfigState {
    /**
     * The ARN for a configuration for DNSSEC validation.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The owner account ID of the virtual private cloud (VPC) for a configuration for DNSSEC validation.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The validation status for a DNSSEC configuration. The status can be one of the following: `ENABLING`, `ENABLED`, `DISABLING` and `DISABLED`.
     */
    readonly validationStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResolverDnsSecConfig resource.
 */
export interface ResolverDnsSecConfigArgs {
    /**
     * The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.
     */
    readonly resourceId: pulumi.Input<string>;
}
