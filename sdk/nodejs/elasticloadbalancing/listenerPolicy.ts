// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a load balancer policy to an ELB Listener.
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_elb_wu_tang = new aws.elasticloadbalancing.LoadBalancer("wu-tang", {
 *     availabilityZones: ["us-east-1a"],
 *     listeners: [{
 *         instancePort: 443,
 *         instanceProtocol: "http",
 *         lbPort: 443,
 *         lbProtocol: "https",
 *         sslCertificateId: "arn:aws:iam::000000000000:server-certificate/wu-tang.net",
 *     }],
 *     name: "wu-tang",
 *     tags: {
 *         Name: "wu-tang",
 *     },
 * });
 * const aws_load_balancer_policy_wu_tang_ssl = new aws.elasticloadbalancing.LoadBalancerPolicy("wu-tang-ssl", {
 *     loadBalancerName: aws_elb_wu_tang.name,
 *     policyAttributes: [
 *         {
 *             name: "ECDHE-ECDSA-AES128-GCM-SHA256",
 *             value: "true",
 *         },
 *         {
 *             name: "Protocol-TLSv1.2",
 *             value: "true",
 *         },
 *     ],
 *     policyName: "wu-tang-ssl",
 *     policyTypeName: "SSLNegotiationPolicyType",
 * });
 * const aws_load_balancer_listener_policy_wu_tang_listener_policies_443 = new aws.elasticloadbalancing.ListenerPolicy("wu-tang-listener-policies-443", {
 *     loadBalancerName: aws_elb_wu_tang.name,
 *     loadBalancerPort: 443,
 *     policyNames: [aws_load_balancer_policy_wu_tang_ssl.policyName],
 * });
 * ```
 * This example shows how to customize the TLS settings of an HTTPS listener.
 * 
 */
export class ListenerPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ListenerPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerPolicyState, opts?: pulumi.CustomResourceOptions): ListenerPolicy {
        return new ListenerPolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * The load balancer to attach the policy to.
     */
    public readonly loadBalancerName: pulumi.Output<string>;
    /**
     * The load balancer listener port to apply the policy to.
     */
    public readonly loadBalancerPort: pulumi.Output<number>;
    /**
     * List of Policy Names to apply to the backend server.
     */
    public readonly policyNames: pulumi.Output<string[] | undefined>;

    /**
     * Create a ListenerPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerPolicyArgs | ListenerPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ListenerPolicyState = argsOrState as ListenerPolicyState | undefined;
            inputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            inputs["loadBalancerPort"] = state ? state.loadBalancerPort : undefined;
            inputs["policyNames"] = state ? state.policyNames : undefined;
        } else {
            const args = argsOrState as ListenerPolicyArgs | undefined;
            if (!args || args.loadBalancerName === undefined) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            if (!args || args.loadBalancerPort === undefined) {
                throw new Error("Missing required property 'loadBalancerPort'");
            }
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["loadBalancerPort"] = args ? args.loadBalancerPort : undefined;
            inputs["policyNames"] = args ? args.policyNames : undefined;
        }
        super("aws:elasticloadbalancing/listenerPolicy:ListenerPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerPolicy resources.
 */
export interface ListenerPolicyState {
    /**
     * The load balancer to attach the policy to.
     */
    readonly loadBalancerName?: pulumi.Input<string>;
    /**
     * The load balancer listener port to apply the policy to.
     */
    readonly loadBalancerPort?: pulumi.Input<number>;
    /**
     * List of Policy Names to apply to the backend server.
     */
    readonly policyNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ListenerPolicy resource.
 */
export interface ListenerPolicyArgs {
    /**
     * The load balancer to attach the policy to.
     */
    readonly loadBalancerName: pulumi.Input<string>;
    /**
     * The load balancer listener port to apply the policy to.
     */
    readonly loadBalancerPort: pulumi.Input<number>;
    /**
     * List of Policy Names to apply to the backend server.
     */
    readonly policyNames?: pulumi.Input<pulumi.Input<string>[]>;
}
