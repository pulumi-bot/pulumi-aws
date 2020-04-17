// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const lb = new aws.elb.LoadBalancer("lb", {
 *     availabilityZones: ["us-east-1a"],
 *     listeners: [{
 *         instancePort: 8000,
 *         instanceProtocol: "https",
 *         lbPort: 443,
 *         lbProtocol: "https",
 *         sslCertificateId: "arn:aws:iam::123456789012:server-certificate/certName",
 *     }],
 * });
 * const foo = new aws.elb.SslNegotiationPolicy("foo", {
 *     attributes: [
 *         {
 *             name: "Protocol-TLSv1",
 *             value: "false",
 *         },
 *         {
 *             name: "Protocol-TLSv1.1",
 *             value: "false",
 *         },
 *         {
 *             name: "Protocol-TLSv1.2",
 *             value: "true",
 *         },
 *         {
 *             name: "Server-Defined-Cipher-Order",
 *             value: "true",
 *         },
 *         {
 *             name: "ECDHE-RSA-AES128-GCM-SHA256",
 *             value: "true",
 *         },
 *         {
 *             name: "AES128-GCM-SHA256",
 *             value: "true",
 *         },
 *         {
 *             name: "EDH-RSA-DES-CBC3-SHA",
 *             value: "false",
 *         },
 *     ],
 *     lbPort: 443,
 *     loadBalancer: lb.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_ssl_negotiation_policy.html.markdown.
 */
export class SslNegotiationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SslNegotiationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SslNegotiationPolicyState, opts?: pulumi.CustomResourceOptions): SslNegotiationPolicy {
        return new SslNegotiationPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy';

    /**
     * Returns true if the given object is an instance of SslNegotiationPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SslNegotiationPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SslNegotiationPolicy.__pulumiType;
    }

    /**
     * An SSL Negotiation policy attribute. Each has two properties:
     */
    public readonly attributes!: pulumi.Output<outputs.elasticloadbalancing.SslNegotiationPolicyAttribute[] | undefined>;
    /**
     * The load balancer port to which the policy
     * should be applied. This must be an active listener on the load
     * balancer.
     */
    public readonly lbPort!: pulumi.Output<number>;
    /**
     * The load balancer to which the policy
     * should be attached.
     */
    public readonly loadBalancer!: pulumi.Output<string>;
    /**
     * The name of the SSL negotiation policy.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SslNegotiationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslNegotiationPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SslNegotiationPolicyArgs | SslNegotiationPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SslNegotiationPolicyState | undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["lbPort"] = state ? state.lbPort : undefined;
            inputs["loadBalancer"] = state ? state.loadBalancer : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SslNegotiationPolicyArgs | undefined;
            if (!args || args.lbPort === undefined) {
                throw new Error("Missing required property 'lbPort'");
            }
            if (!args || args.loadBalancer === undefined) {
                throw new Error("Missing required property 'loadBalancer'");
            }
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["lbPort"] = args ? args.lbPort : undefined;
            inputs["loadBalancer"] = args ? args.loadBalancer : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SslNegotiationPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SslNegotiationPolicy resources.
 */
export interface SslNegotiationPolicyState {
    /**
     * An SSL Negotiation policy attribute. Each has two properties:
     */
    readonly attributes?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancing.SslNegotiationPolicyAttribute>[]>;
    /**
     * The load balancer port to which the policy
     * should be applied. This must be an active listener on the load
     * balancer.
     */
    readonly lbPort?: pulumi.Input<number>;
    /**
     * The load balancer to which the policy
     * should be attached.
     */
    readonly loadBalancer?: pulumi.Input<string>;
    /**
     * The name of the SSL negotiation policy.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SslNegotiationPolicy resource.
 */
export interface SslNegotiationPolicyArgs {
    /**
     * An SSL Negotiation policy attribute. Each has two properties:
     */
    readonly attributes?: pulumi.Input<pulumi.Input<inputs.elasticloadbalancing.SslNegotiationPolicyAttribute>[]>;
    /**
     * The load balancer port to which the policy
     * should be applied. This must be an active listener on the load
     * balancer.
     */
    readonly lbPort: pulumi.Input<number>;
    /**
     * The load balancer to which the policy
     * should be attached.
     */
    readonly loadBalancer: pulumi.Input<string>;
    /**
     * The name of the SSL negotiation policy.
     */
    readonly name?: pulumi.Input<string>;
}
