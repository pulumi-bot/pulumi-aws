// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an SSM Maintenance Window Target resource
 *
 * ## Example Usage
 * ### Instance Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const window = new aws.ssm.MaintenanceWindow("window", {
 *     schedule: "cron(0 16 ? * TUE *)",
 *     duration: 3,
 *     cutoff: 1,
 * });
 * const target1 = new aws.ssm.MaintenanceWindowTarget("target1", {
 *     windowId: window.id,
 *     description: "This is a maintenance window target",
 *     resourceType: "INSTANCE",
 *     targets: [{
 *         key: "tag:Name",
 *         values: ["acceptance_test"],
 *     }],
 * });
 * ```
 * ### Resource Group Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const window = new aws.ssm.MaintenanceWindow("window", {
 *     schedule: "cron(0 16 ? * TUE *)",
 *     duration: 3,
 *     cutoff: 1,
 * });
 * const target1 = new aws.ssm.MaintenanceWindowTarget("target1", {
 *     windowId: window.id,
 *     description: "This is a maintenance window target",
 *     resourceType: "RESOURCE_GROUP",
 *     targets: [{
 *         key: "resource-groups:ResourceTypeFilters",
 *         values: ["AWS::EC2::Instance"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * SSM Maintenance Window targets can be imported using `WINDOW_ID/WINDOW_TARGET_ID`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget example mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE
 * ```
 */
export class MaintenanceWindowTarget extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindowTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaintenanceWindowTargetState, opts?: pulumi.CustomResourceOptions): MaintenanceWindowTarget {
        return new MaintenanceWindowTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget';

    /**
     * Returns true if the given object is an instance of MaintenanceWindowTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindowTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindowTarget.__pulumiType;
    }

    /**
     * The description of the maintenance window target.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the maintenance window target.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
     */
    public readonly ownerInformation!: pulumi.Output<string | undefined>;
    /**
     * The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
     * (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
     */
    public readonly targets!: pulumi.Output<outputs.ssm.MaintenanceWindowTargetTarget[]>;
    /**
     * The Id of the maintenance window to register the target with.
     */
    public readonly windowId!: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindowTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaintenanceWindowTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaintenanceWindowTargetArgs | MaintenanceWindowTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MaintenanceWindowTargetState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ownerInformation"] = state ? state.ownerInformation : undefined;
            inputs["resourceType"] = state ? state.resourceType : undefined;
            inputs["targets"] = state ? state.targets : undefined;
            inputs["windowId"] = state ? state.windowId : undefined;
        } else {
            const args = argsOrState as MaintenanceWindowTargetArgs | undefined;
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.targets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targets'");
            }
            if ((!args || args.windowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'windowId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ownerInformation"] = args ? args.ownerInformation : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["windowId"] = args ? args.windowId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MaintenanceWindowTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaintenanceWindowTarget resources.
 */
export interface MaintenanceWindowTargetState {
    /**
     * The description of the maintenance window target.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the maintenance window target.
     */
    name?: pulumi.Input<string>;
    /**
     * User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
     */
    ownerInformation?: pulumi.Input<string>;
    /**
     * The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
     * (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
     */
    targets?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTargetTarget>[]>;
    /**
     * The Id of the maintenance window to register the target with.
     */
    windowId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MaintenanceWindowTarget resource.
 */
export interface MaintenanceWindowTargetArgs {
    /**
     * The description of the maintenance window target.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the maintenance window target.
     */
    name?: pulumi.Input<string>;
    /**
     * User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
     */
    ownerInformation?: pulumi.Input<string>;
    /**
     * The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
     * (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
     */
    targets: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTargetTarget>[]>;
    /**
     * The Id of the maintenance window to register the target with.
     */
    windowId: pulumi.Input<string>;
}
