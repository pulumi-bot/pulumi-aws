// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Attaches a load balancer policy to an ELB backend server.
 * 
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/load_balancer_backend_server_policy.html.markdown.
 */
/** @deprecated aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favour of aws.elb.LoadBalancerBackendServerPolicy */
export class LoadBalancerBackendServerPolicy extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerBackendServerPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerBackendServerPolicyState, opts?: pulumi.CustomResourceOptions): LoadBalancerBackendServerPolicy {
        pulumi.log.warn("LoadBalancerBackendServerPolicy is deprecated: aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favour of aws.elb.LoadBalancerBackendServerPolicy")
        return new LoadBalancerBackendServerPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy';

    /**
     * Returns true if the given object is an instance of LoadBalancerBackendServerPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerBackendServerPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerBackendServerPolicy.__pulumiType;
    }

    /**
     * The instance port to apply the policy to.
     */
    public readonly instancePort!: pulumi.Output<number>;
    /**
     * The load balancer to attach the policy to.
     */
    public readonly loadBalancerName!: pulumi.Output<string>;
    /**
     * List of Policy Names to apply to the backend server.
     */
    public readonly policyNames!: pulumi.Output<string[] | undefined>;

    /**
     * Create a LoadBalancerBackendServerPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favour of aws.elb.LoadBalancerBackendServerPolicy */
    constructor(name: string, args: LoadBalancerBackendServerPolicyArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favour of aws.elb.LoadBalancerBackendServerPolicy */
    constructor(name: string, argsOrState?: LoadBalancerBackendServerPolicyArgs | LoadBalancerBackendServerPolicyState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoadBalancerBackendServerPolicy is deprecated: aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favour of aws.elb.LoadBalancerBackendServerPolicy")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LoadBalancerBackendServerPolicyState | undefined;
            inputs["instancePort"] = state ? state.instancePort : undefined;
            inputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            inputs["policyNames"] = state ? state.policyNames : undefined;
        } else {
            const args = argsOrState as LoadBalancerBackendServerPolicyArgs | undefined;
            if (!args || args.instancePort === undefined) {
                throw new Error("Missing required property 'instancePort'");
            }
            if (!args || args.loadBalancerName === undefined) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            inputs["instancePort"] = args ? args.instancePort : undefined;
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["policyNames"] = args ? args.policyNames : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LoadBalancerBackendServerPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancerBackendServerPolicy resources.
 */
export interface LoadBalancerBackendServerPolicyState {
    /**
     * The instance port to apply the policy to.
     */
    readonly instancePort?: pulumi.Input<number>;
    /**
     * The load balancer to attach the policy to.
     */
    readonly loadBalancerName?: pulumi.Input<string>;
    /**
     * List of Policy Names to apply to the backend server.
     */
    readonly policyNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a LoadBalancerBackendServerPolicy resource.
 */
export interface LoadBalancerBackendServerPolicyArgs {
    /**
     * The instance port to apply the policy to.
     */
    readonly instancePort: pulumi.Input<number>;
    /**
     * The load balancer to attach the policy to.
     */
    readonly loadBalancerName: pulumi.Input<string>;
    /**
     * List of Policy Names to apply to the backend server.
     */
    readonly policyNames?: pulumi.Input<pulumi.Input<string>[]>;
}
