// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks stack resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.opsworks.Stack("main", {
 *     region: "us-west-1",
 *     serviceRoleArn: aws_iam_role.opsworks.arn,
 *     defaultInstanceProfileArn: aws_iam_instance_profile.opsworks.arn,
 *     tags: {
 *         Name: "foobar-stack",
 *     },
 *     customJson: `{
 *  "foobar": {
 *     "version": "1.0.0"
 *   }
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * OpsWorks stacks can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:opsworks/stack:Stack bar 00000000-0000-0000-0000-000000000000
 * ```
 */
export class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackState, opts?: pulumi.CustomResourceOptions): Stack {
        return new Stack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/stack:Stack';

    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stack.__pulumiType;
    }

    /**
     * If set to `"LATEST"`, OpsWorks will automatically install the latest version.
     */
    public readonly agentVersion!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * If `manageBerkshelf` is enabled, the version of Berkshelf to use.
     */
    public readonly berkshelfVersion!: pulumi.Output<string | undefined>;
    /**
     * Color to paint next to the stack's resources in the OpsWorks console.
     */
    public readonly color!: pulumi.Output<string | undefined>;
    /**
     * Name of the configuration manager to use. Defaults to "Chef".
     */
    public readonly configurationManagerName!: pulumi.Output<string | undefined>;
    /**
     * Version of the configuration manager to use. Defaults to "11.4".
     */
    public readonly configurationManagerVersion!: pulumi.Output<string | undefined>;
    /**
     * When `useCustomCookbooks` is set, provide this sub-object as
     * described below.
     */
    public readonly customCookbooksSources!: pulumi.Output<outputs.opsworks.StackCustomCookbooksSource[]>;
    /**
     * Custom JSON attributes to apply to the entire stack.
     */
    public readonly customJson!: pulumi.Output<string | undefined>;
    /**
     * Name of the availability zone where instances will be created
     * by default. This is required unless you set `vpcId`.
     */
    public readonly defaultAvailabilityZone!: pulumi.Output<string>;
    /**
     * The ARN of an IAM Instance Profile that created instances
     * will have by default.
     */
    public readonly defaultInstanceProfileArn!: pulumi.Output<string>;
    /**
     * Name of OS that will be installed on instances by default.
     */
    public readonly defaultOs!: pulumi.Output<string | undefined>;
    /**
     * Name of the type of root device instances will have by default.
     */
    public readonly defaultRootDeviceType!: pulumi.Output<string | undefined>;
    /**
     * Name of the SSH keypair that instances will have by default.
     */
    public readonly defaultSshKeyName!: pulumi.Output<string | undefined>;
    /**
     * Id of the subnet in which instances will be created by default. Mandatory
     * if `vpcId` is set, and forbidden if it isn't.
     */
    public readonly defaultSubnetId!: pulumi.Output<string>;
    /**
     * Keyword representing the naming scheme that will be used for instance hostnames
     * within this stack.
     */
    public readonly hostnameTheme!: pulumi.Output<string | undefined>;
    /**
     * Boolean value controlling whether Opsworks will run Berkshelf for this stack.
     */
    public readonly manageBerkshelf!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the stack.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the region where the stack will exist.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The ARN of an IAM role that the OpsWorks service will act as.
     */
    public readonly serviceRoleArn!: pulumi.Output<string>;
    public /*out*/ readonly stackEndpoint!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Boolean value controlling whether the custom cookbook settings are
     * enabled.
     */
    public readonly useCustomCookbooks!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean value controlling whether the standard OpsWorks
     * security groups apply to created instances.
     */
    public readonly useOpsworksSecurityGroups!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the VPC that this stack belongs to.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackArgs | StackState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StackState | undefined;
            inputs["agentVersion"] = state ? state.agentVersion : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["berkshelfVersion"] = state ? state.berkshelfVersion : undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["configurationManagerName"] = state ? state.configurationManagerName : undefined;
            inputs["configurationManagerVersion"] = state ? state.configurationManagerVersion : undefined;
            inputs["customCookbooksSources"] = state ? state.customCookbooksSources : undefined;
            inputs["customJson"] = state ? state.customJson : undefined;
            inputs["defaultAvailabilityZone"] = state ? state.defaultAvailabilityZone : undefined;
            inputs["defaultInstanceProfileArn"] = state ? state.defaultInstanceProfileArn : undefined;
            inputs["defaultOs"] = state ? state.defaultOs : undefined;
            inputs["defaultRootDeviceType"] = state ? state.defaultRootDeviceType : undefined;
            inputs["defaultSshKeyName"] = state ? state.defaultSshKeyName : undefined;
            inputs["defaultSubnetId"] = state ? state.defaultSubnetId : undefined;
            inputs["hostnameTheme"] = state ? state.hostnameTheme : undefined;
            inputs["manageBerkshelf"] = state ? state.manageBerkshelf : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            inputs["stackEndpoint"] = state ? state.stackEndpoint : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["useCustomCookbooks"] = state ? state.useCustomCookbooks : undefined;
            inputs["useOpsworksSecurityGroups"] = state ? state.useOpsworksSecurityGroups : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as StackArgs | undefined;
            if (!args || args.defaultInstanceProfileArn === undefined) {
                throw new Error("Missing required property 'defaultInstanceProfileArn'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            if (!args || args.serviceRoleArn === undefined) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            inputs["agentVersion"] = args ? args.agentVersion : undefined;
            inputs["berkshelfVersion"] = args ? args.berkshelfVersion : undefined;
            inputs["color"] = args ? args.color : undefined;
            inputs["configurationManagerName"] = args ? args.configurationManagerName : undefined;
            inputs["configurationManagerVersion"] = args ? args.configurationManagerVersion : undefined;
            inputs["customCookbooksSources"] = args ? args.customCookbooksSources : undefined;
            inputs["customJson"] = args ? args.customJson : undefined;
            inputs["defaultAvailabilityZone"] = args ? args.defaultAvailabilityZone : undefined;
            inputs["defaultInstanceProfileArn"] = args ? args.defaultInstanceProfileArn : undefined;
            inputs["defaultOs"] = args ? args.defaultOs : undefined;
            inputs["defaultRootDeviceType"] = args ? args.defaultRootDeviceType : undefined;
            inputs["defaultSshKeyName"] = args ? args.defaultSshKeyName : undefined;
            inputs["defaultSubnetId"] = args ? args.defaultSubnetId : undefined;
            inputs["hostnameTheme"] = args ? args.hostnameTheme : undefined;
            inputs["manageBerkshelf"] = args ? args.manageBerkshelf : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["useCustomCookbooks"] = args ? args.useCustomCookbooks : undefined;
            inputs["useOpsworksSecurityGroups"] = args ? args.useOpsworksSecurityGroups : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["stackEndpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Stack.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stack resources.
 */
export interface StackState {
    /**
     * If set to `"LATEST"`, OpsWorks will automatically install the latest version.
     */
    readonly agentVersion?: pulumi.Input<string>;
    readonly arn?: pulumi.Input<string>;
    /**
     * If `manageBerkshelf` is enabled, the version of Berkshelf to use.
     */
    readonly berkshelfVersion?: pulumi.Input<string>;
    /**
     * Color to paint next to the stack's resources in the OpsWorks console.
     */
    readonly color?: pulumi.Input<string>;
    /**
     * Name of the configuration manager to use. Defaults to "Chef".
     */
    readonly configurationManagerName?: pulumi.Input<string>;
    /**
     * Version of the configuration manager to use. Defaults to "11.4".
     */
    readonly configurationManagerVersion?: pulumi.Input<string>;
    /**
     * When `useCustomCookbooks` is set, provide this sub-object as
     * described below.
     */
    readonly customCookbooksSources?: pulumi.Input<pulumi.Input<inputs.opsworks.StackCustomCookbooksSource>[]>;
    /**
     * Custom JSON attributes to apply to the entire stack.
     */
    readonly customJson?: pulumi.Input<string>;
    /**
     * Name of the availability zone where instances will be created
     * by default. This is required unless you set `vpcId`.
     */
    readonly defaultAvailabilityZone?: pulumi.Input<string>;
    /**
     * The ARN of an IAM Instance Profile that created instances
     * will have by default.
     */
    readonly defaultInstanceProfileArn?: pulumi.Input<string>;
    /**
     * Name of OS that will be installed on instances by default.
     */
    readonly defaultOs?: pulumi.Input<string>;
    /**
     * Name of the type of root device instances will have by default.
     */
    readonly defaultRootDeviceType?: pulumi.Input<string>;
    /**
     * Name of the SSH keypair that instances will have by default.
     */
    readonly defaultSshKeyName?: pulumi.Input<string>;
    /**
     * Id of the subnet in which instances will be created by default. Mandatory
     * if `vpcId` is set, and forbidden if it isn't.
     */
    readonly defaultSubnetId?: pulumi.Input<string>;
    /**
     * Keyword representing the naming scheme that will be used for instance hostnames
     * within this stack.
     */
    readonly hostnameTheme?: pulumi.Input<string>;
    /**
     * Boolean value controlling whether Opsworks will run Berkshelf for this stack.
     */
    readonly manageBerkshelf?: pulumi.Input<boolean>;
    /**
     * The name of the stack.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the region where the stack will exist.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that the OpsWorks service will act as.
     */
    readonly serviceRoleArn?: pulumi.Input<string>;
    readonly stackEndpoint?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Boolean value controlling whether the custom cookbook settings are
     * enabled.
     */
    readonly useCustomCookbooks?: pulumi.Input<boolean>;
    /**
     * Boolean value controlling whether the standard OpsWorks
     * security groups apply to created instances.
     */
    readonly useOpsworksSecurityGroups?: pulumi.Input<boolean>;
    /**
     * The id of the VPC that this stack belongs to.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    /**
     * If set to `"LATEST"`, OpsWorks will automatically install the latest version.
     */
    readonly agentVersion?: pulumi.Input<string>;
    /**
     * If `manageBerkshelf` is enabled, the version of Berkshelf to use.
     */
    readonly berkshelfVersion?: pulumi.Input<string>;
    /**
     * Color to paint next to the stack's resources in the OpsWorks console.
     */
    readonly color?: pulumi.Input<string>;
    /**
     * Name of the configuration manager to use. Defaults to "Chef".
     */
    readonly configurationManagerName?: pulumi.Input<string>;
    /**
     * Version of the configuration manager to use. Defaults to "11.4".
     */
    readonly configurationManagerVersion?: pulumi.Input<string>;
    /**
     * When `useCustomCookbooks` is set, provide this sub-object as
     * described below.
     */
    readonly customCookbooksSources?: pulumi.Input<pulumi.Input<inputs.opsworks.StackCustomCookbooksSource>[]>;
    /**
     * Custom JSON attributes to apply to the entire stack.
     */
    readonly customJson?: pulumi.Input<string>;
    /**
     * Name of the availability zone where instances will be created
     * by default. This is required unless you set `vpcId`.
     */
    readonly defaultAvailabilityZone?: pulumi.Input<string>;
    /**
     * The ARN of an IAM Instance Profile that created instances
     * will have by default.
     */
    readonly defaultInstanceProfileArn: pulumi.Input<string>;
    /**
     * Name of OS that will be installed on instances by default.
     */
    readonly defaultOs?: pulumi.Input<string>;
    /**
     * Name of the type of root device instances will have by default.
     */
    readonly defaultRootDeviceType?: pulumi.Input<string>;
    /**
     * Name of the SSH keypair that instances will have by default.
     */
    readonly defaultSshKeyName?: pulumi.Input<string>;
    /**
     * Id of the subnet in which instances will be created by default. Mandatory
     * if `vpcId` is set, and forbidden if it isn't.
     */
    readonly defaultSubnetId?: pulumi.Input<string>;
    /**
     * Keyword representing the naming scheme that will be used for instance hostnames
     * within this stack.
     */
    readonly hostnameTheme?: pulumi.Input<string>;
    /**
     * Boolean value controlling whether Opsworks will run Berkshelf for this stack.
     */
    readonly manageBerkshelf?: pulumi.Input<boolean>;
    /**
     * The name of the stack.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the region where the stack will exist.
     */
    readonly region: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that the OpsWorks service will act as.
     */
    readonly serviceRoleArn: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Boolean value controlling whether the custom cookbook settings are
     * enabled.
     */
    readonly useCustomCookbooks?: pulumi.Input<boolean>;
    /**
     * Boolean value controlling whether the standard OpsWorks
     * security groups apply to created instances.
     */
    readonly useOpsworksSecurityGroups?: pulumi.Input<boolean>;
    /**
     * The id of the VPC that this stack belongs to.
     */
    readonly vpcId?: pulumi.Input<string>;
}
