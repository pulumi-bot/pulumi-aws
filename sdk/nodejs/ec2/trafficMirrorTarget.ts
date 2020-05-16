// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Traffic mirror target.  
 * Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const nlb = new aws.ec2.TrafficMirrorTarget("nlb", {
 *     description: "NLB target",
 *     networkLoadBalancerArn: aws_lb_lb.arn,
 * });
 * const eni = new aws.ec2.TrafficMirrorTarget("eni", {
 *     description: "ENI target",
 *     networkInterfaceId: aws_instance_test.primaryNetworkInterfaceId,
 * });
 * ```
 */
export class TrafficMirrorTarget extends pulumi.CustomResource {
    /**
     * Get an existing TrafficMirrorTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficMirrorTargetState, opts?: pulumi.CustomResourceOptions): TrafficMirrorTarget {
        return new TrafficMirrorTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/trafficMirrorTarget:TrafficMirrorTarget';

    /**
     * Returns true if the given object is an instance of TrafficMirrorTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrafficMirrorTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrafficMirrorTarget.__pulumiType;
    }

    /**
     * A description of the traffic mirror session.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The network interface ID that is associated with the target.
     */
    public readonly networkInterfaceId!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
     */
    public readonly networkLoadBalancerArn!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a TrafficMirrorTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TrafficMirrorTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficMirrorTargetArgs | TrafficMirrorTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TrafficMirrorTargetState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["networkLoadBalancerArn"] = state ? state.networkLoadBalancerArn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as TrafficMirrorTargetArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            inputs["networkLoadBalancerArn"] = args ? args.networkLoadBalancerArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TrafficMirrorTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrafficMirrorTarget resources.
 */
export interface TrafficMirrorTargetState {
    /**
     * A description of the traffic mirror session.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The network interface ID that is associated with the target.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
     */
    readonly networkLoadBalancerArn?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a TrafficMirrorTarget resource.
 */
export interface TrafficMirrorTargetArgs {
    /**
     * A description of the traffic mirror session.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The network interface ID that is associated with the target.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
     */
    readonly networkLoadBalancerArn?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
