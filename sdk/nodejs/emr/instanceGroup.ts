// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Elastic MapReduce Cluster Instance Group configuration.
 * See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.
 *
 * > **NOTE:** At this time, Instance Groups cannot be destroyed through the API nor
 * web interface. Instance Groups are destroyed when the EMR Cluster is destroyed.
 * this provider will resize any Instance Group to zero when destroying the resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const task = new aws.emr.InstanceGroup("task", {
 *     clusterId: aws_emr_cluster["tf-test-cluster"].id,
 *     instanceCount: 1,
 *     instanceType: "m5.xlarge",
 * });
 * ```
 */
export class InstanceGroup extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGroupState, opts?: pulumi.CustomResourceOptions): InstanceGroup {
        return new InstanceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:emr/instanceGroup:InstanceGroup';

    /**
     * Returns true if the given object is an instance of InstanceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceGroup.__pulumiType;
    }

    /**
     * The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
     */
    public readonly autoscalingPolicy!: pulumi.Output<string | undefined>;
    /**
     * If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
     */
    public readonly bidPrice!: pulumi.Output<string | undefined>;
    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
     */
    public readonly configurationsJson!: pulumi.Output<string | undefined>;
    /**
     * One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
     */
    public readonly ebsConfigs!: pulumi.Output<outputs.emr.InstanceGroupEbsConfig[]>;
    /**
     * Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
     */
    public readonly ebsOptimized!: pulumi.Output<boolean | undefined>;
    /**
     * target number of instances for the instance group. defaults to 0.
     */
    public readonly instanceCount!: pulumi.Output<number | undefined>;
    /**
     * The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Human friendly name given to the instance group. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly runningInstanceCount!: pulumi.Output<number>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a InstanceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceGroupArgs | InstanceGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceGroupState | undefined;
            inputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            inputs["bidPrice"] = state ? state.bidPrice : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["configurationsJson"] = state ? state.configurationsJson : undefined;
            inputs["ebsConfigs"] = state ? state.ebsConfigs : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["instanceCount"] = state ? state.instanceCount : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["runningInstanceCount"] = state ? state.runningInstanceCount : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as InstanceGroupArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["bidPrice"] = args ? args.bidPrice : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["configurationsJson"] = args ? args.configurationsJson : undefined;
            inputs["ebsConfigs"] = args ? args.ebsConfigs : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["instanceCount"] = args ? args.instanceCount : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["runningInstanceCount"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGroup resources.
 */
export interface InstanceGroupState {
    /**
     * The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
     */
    readonly autoscalingPolicy?: pulumi.Input<string>;
    /**
     * If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
     */
    readonly bidPrice?: pulumi.Input<string>;
    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
     */
    readonly configurationsJson?: pulumi.Input<string>;
    /**
     * One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly ebsConfigs?: pulumi.Input<pulumi.Input<inputs.emr.InstanceGroupEbsConfig>[]>;
    /**
     * Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
     */
    readonly ebsOptimized?: pulumi.Input<boolean>;
    /**
     * target number of instances for the instance group. defaults to 0.
     */
    readonly instanceCount?: pulumi.Input<number>;
    /**
     * The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
     */
    readonly instanceType?: pulumi.Input<string>;
    /**
     * Human friendly name given to the instance group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly runningInstanceCount?: pulumi.Input<number>;
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroup resource.
 */
export interface InstanceGroupArgs {
    /**
     * The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
     */
    readonly autoscalingPolicy?: pulumi.Input<string>;
    /**
     * If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
     */
    readonly bidPrice?: pulumi.Input<string>;
    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
     */
    readonly configurationsJson?: pulumi.Input<string>;
    /**
     * One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
     */
    readonly ebsConfigs?: pulumi.Input<pulumi.Input<inputs.emr.InstanceGroupEbsConfig>[]>;
    /**
     * Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
     */
    readonly ebsOptimized?: pulumi.Input<boolean>;
    /**
     * target number of instances for the instance group. defaults to 0.
     */
    readonly instanceCount?: pulumi.Input<number>;
    /**
     * The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
     */
    readonly instanceType: pulumi.Input<string>;
    /**
     * Human friendly name given to the instance group. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
}
