// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Route 53 Resolver endpoint resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.route53.ResolverEndpoint("foo", {
 *     direction: "INBOUND",
 *     ipAddresses: [
 *         {
 *             subnetId: aws_subnet_sn1.id,
 *         },
 *         {
 *             ip: "10.0.64.4",
 *             subnetId: aws_subnet_sn2.id,
 *         },
 *     ],
 *     securityGroupIds: [
 *         aws_security_group_sg1.id,
 *         aws_security_group_sg2.id,
 *     ],
 *     tags: {
 *         Environment: "Prod",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_endpoint.html.markdown.
 */
export class ResolverEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing ResolverEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResolverEndpointState, opts?: pulumi.CustomResourceOptions): ResolverEndpoint {
        return new ResolverEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/resolverEndpoint:ResolverEndpoint';

    /**
     * Returns true if the given object is an instance of ResolverEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverEndpoint.__pulumiType;
    }

    /**
     * The ARN of the Route 53 Resolver endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The direction of DNS queries to or from the Route 53 Resolver endpoint.
     * Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
     * or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The ID of the VPC that you want to create the resolver endpoint in.
     */
    public /*out*/ readonly hostVpcId!: pulumi.Output<string>;
    /**
     * The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
     * to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
     */
    public readonly ipAddresses!: pulumi.Output<{ ip: string, ipId: string, subnetId: string }[]>;
    /**
     * The friendly name of the Route 53 Resolver endpoint.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of one or more security groups that you want to use to control access to this VPC.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a ResolverEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResolverEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResolverEndpointArgs | ResolverEndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResolverEndpointState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["hostVpcId"] = state ? state.hostVpcId : undefined;
            inputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ResolverEndpointArgs | undefined;
            if (!args || args.direction === undefined) {
                throw new Error("Missing required property 'direction'");
            }
            if (!args || args.ipAddresses === undefined) {
                throw new Error("Missing required property 'ipAddresses'");
            }
            if (!args || args.securityGroupIds === undefined) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            inputs["direction"] = args ? args.direction : undefined;
            inputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["hostVpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResolverEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResolverEndpoint resources.
 */
export interface ResolverEndpointState {
    /**
     * The ARN of the Route 53 Resolver endpoint.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The direction of DNS queries to or from the Route 53 Resolver endpoint.
     * Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
     * or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * The ID of the VPC that you want to create the resolver endpoint in.
     */
    readonly hostVpcId?: pulumi.Input<string>;
    /**
     * The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
     * to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
     */
    readonly ipAddresses?: pulumi.Input<pulumi.Input<{ ip?: pulumi.Input<string>, ipId?: pulumi.Input<string>, subnetId: pulumi.Input<string> }>[]>;
    /**
     * The friendly name of the Route 53 Resolver endpoint.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of one or more security groups that you want to use to control access to this VPC.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ResolverEndpoint resource.
 */
export interface ResolverEndpointArgs {
    /**
     * The direction of DNS queries to or from the Route 53 Resolver endpoint.
     * Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
     * or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
     */
    readonly direction: pulumi.Input<string>;
    /**
     * The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
     * to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
     */
    readonly ipAddresses: pulumi.Input<pulumi.Input<{ ip?: pulumi.Input<string>, ipId?: pulumi.Input<string>, subnetId: pulumi.Input<string> }>[]>;
    /**
     * The friendly name of the Route 53 Resolver endpoint.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of one or more security groups that you want to use to control access to this VPC.
     */
    readonly securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
