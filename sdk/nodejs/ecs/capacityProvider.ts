// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).
 * 
 * > **NOTE:** The AWS API does not currently support deleting ECS cluster capacity providers. Removing this resource will only remove the state for it.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const test = new aws.ecs.CapacityProvider("test", {auto_scaling_group_provider: {
 *     autoScalingGroupArn: aws_autoscaling_group.test.arn,
 *     managedTerminationProtection: "ENABLED",
 *     managed_scaling: {
 *         maximumScalingStepSize: 1000,
 *         minimumScalingStepSize: 1,
 *         status: "ENABLED",
 *         targetCapacity: 10,
 *     },
 * }});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_capacity_provider.html.markdown.
 */
export class CapacityProvider extends pulumi.CustomResource {
    /**
     * Get an existing CapacityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CapacityProviderState, opts?: pulumi.CustomResourceOptions): CapacityProvider {
        return new CapacityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecs/capacityProvider:CapacityProvider';

    /**
     * Returns true if the given object is an instance of CapacityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CapacityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CapacityProvider.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) that identifies the capacity provider.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Nested argument defining the provider for the ECS auto scaling group. Defined below.
     */
    public readonly autoScalingGroupProvider!: pulumi.Output<outputs.ecs.CapacityProviderAutoScalingGroupProvider>;
    /**
     * The name of the capacity provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a CapacityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CapacityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CapacityProviderArgs | CapacityProviderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CapacityProviderState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoScalingGroupProvider"] = state ? state.autoScalingGroupProvider : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as CapacityProviderArgs | undefined;
            if (!args || args.autoScalingGroupProvider === undefined) {
                throw new Error("Missing required property 'autoScalingGroupProvider'");
            }
            inputs["autoScalingGroupProvider"] = args ? args.autoScalingGroupProvider : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CapacityProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CapacityProvider resources.
 */
export interface CapacityProviderState {
    /**
     * The Amazon Resource Name (ARN) that identifies the capacity provider.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Nested argument defining the provider for the ECS auto scaling group. Defined below.
     */
    readonly autoScalingGroupProvider?: pulumi.Input<inputs.ecs.CapacityProviderAutoScalingGroupProvider>;
    /**
     * The name of the capacity provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a CapacityProvider resource.
 */
export interface CapacityProviderArgs {
    /**
     * Nested argument defining the provider for the ECS auto scaling group. Defined below.
     */
    readonly autoScalingGroupProvider: pulumi.Input<inputs.ecs.CapacityProviderAutoScalingGroupProvider>;
    /**
     * The name of the capacity provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
