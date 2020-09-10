// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class LoadBalancerPolicy extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerPolicyState, opts?: pulumi.CustomResourceOptions): LoadBalancerPolicy {
        return new LoadBalancerPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elb/loadBalancerPolicy:LoadBalancerPolicy';

    /**
     * Returns true if the given object is an instance of LoadBalancerPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerPolicy.__pulumiType;
    }

    public readonly loadBalancerName!: pulumi.Output<string>;
    public readonly policyAttributes!: pulumi.Output<outputs.elb.LoadBalancerPolicyPolicyAttribute[] | undefined>;
    public readonly policyName!: pulumi.Output<string>;
    public readonly policyTypeName!: pulumi.Output<string>;

    /**
     * Create a LoadBalancerPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerPolicyArgs | LoadBalancerPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LoadBalancerPolicyState | undefined;
            inputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            inputs["policyAttributes"] = state ? state.policyAttributes : undefined;
            inputs["policyName"] = state ? state.policyName : undefined;
            inputs["policyTypeName"] = state ? state.policyTypeName : undefined;
        } else {
            const args = argsOrState as LoadBalancerPolicyArgs | undefined;
            if (!args || args.loadBalancerName === undefined) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            if (!args || args.policyName === undefined) {
                throw new Error("Missing required property 'policyName'");
            }
            if (!args || args.policyTypeName === undefined) {
                throw new Error("Missing required property 'policyTypeName'");
            }
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["policyAttributes"] = args ? args.policyAttributes : undefined;
            inputs["policyName"] = args ? args.policyName : undefined;
            inputs["policyTypeName"] = args ? args.policyTypeName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LoadBalancerPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerPolicy resources.
 */
export interface LoadBalancerPolicyState {
    readonly loadBalancerName?: pulumi.Input<string>;
    readonly policyAttributes?: pulumi.Input<pulumi.Input<inputs.elb.LoadBalancerPolicyPolicyAttribute>[]>;
    readonly policyName?: pulumi.Input<string>;
    readonly policyTypeName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancerPolicy resource.
 */
export interface LoadBalancerPolicyArgs {
    readonly loadBalancerName: pulumi.Input<string>;
    readonly policyAttributes?: pulumi.Input<pulumi.Input<inputs.elb.LoadBalancerPolicyPolicyAttribute>[]>;
    readonly policyName: pulumi.Input<string>;
    readonly policyTypeName: pulumi.Input<string>;
}
