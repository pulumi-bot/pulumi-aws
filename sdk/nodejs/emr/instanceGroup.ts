// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class InstanceGroup extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGroupState, opts?: pulumi.CustomResourceOptions): InstanceGroup {
        return new InstanceGroup(name, <any>state, { ...opts, id: id });
    }

    public readonly clusterId: pulumi.Output<string>;
    public readonly ebsConfigs: pulumi.Output<{ iops?: number, size: number, type: string, volumesPerInstance?: number }[] | undefined>;
    public readonly ebsOptimized: pulumi.Output<boolean | undefined>;
    public readonly instanceCount: pulumi.Output<number | undefined>;
    public readonly instanceType: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public /*out*/ readonly runningInstanceCount: pulumi.Output<number>;
    public /*out*/ readonly status: pulumi.Output<string>;

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
            const state: InstanceGroupState = argsOrState as InstanceGroupState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
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
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["ebsConfigs"] = args ? args.ebsConfigs : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["instanceCount"] = args ? args.instanceCount : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["runningInstanceCount"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        super("aws:emr/instanceGroup:InstanceGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGroup resources.
 */
export interface InstanceGroupState {
    readonly clusterId?: pulumi.Input<string>;
    readonly ebsConfigs?: pulumi.Input<pulumi.Input<{ iops?: pulumi.Input<number>, size: pulumi.Input<number>, type: pulumi.Input<string>, volumesPerInstance?: pulumi.Input<number> }>[]>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly instanceCount?: pulumi.Input<number>;
    readonly instanceType?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly runningInstanceCount?: pulumi.Input<number>;
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroup resource.
 */
export interface InstanceGroupArgs {
    readonly clusterId: pulumi.Input<string>;
    readonly ebsConfigs?: pulumi.Input<pulumi.Input<{ iops?: pulumi.Input<number>, size: pulumi.Input<number>, type: pulumi.Input<string>, volumesPerInstance?: pulumi.Input<number> }>[]>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly instanceCount?: pulumi.Input<number>;
    readonly instanceType: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
}
