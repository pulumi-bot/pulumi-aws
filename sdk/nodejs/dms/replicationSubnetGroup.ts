// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * // Create a new replication subnet group
 * const test = new aws.dms.ReplicationSubnetGroup("test", {
 *     replicationSubnetGroupDescription: "Test replication subnet group",
 *     replicationSubnetGroupId: "test-dms-replication-subnet-group-tf",
 *     subnetIds: ["subnet-12345678"],
 *     tags: {
 *         Name: "test",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dms_replication_subnet_group.html.markdown.
 */
export class ReplicationSubnetGroup extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationSubnetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicationSubnetGroupState, opts?: pulumi.CustomResourceOptions): ReplicationSubnetGroup {
        return new ReplicationSubnetGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:dms/replicationSubnetGroup:ReplicationSubnetGroup';

    /**
     * Returns true if the given object is an instance of ReplicationSubnetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationSubnetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationSubnetGroup.__pulumiType;
    }

    public /*out*/ readonly replicationSubnetGroupArn!: pulumi.Output<string>;
    /**
     * The description for the subnet group.
     */
    public readonly replicationSubnetGroupDescription!: pulumi.Output<string>;
    /**
     * The name for the replication subnet group. This value is stored as a lowercase string.
     */
    public readonly replicationSubnetGroupId!: pulumi.Output<string>;
    /**
     * A list of the EC2 subnet IDs for the subnet group.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the VPC the subnet group is in.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a ReplicationSubnetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationSubnetGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationSubnetGroupArgs | ReplicationSubnetGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReplicationSubnetGroupState | undefined;
            inputs["replicationSubnetGroupArn"] = state ? state.replicationSubnetGroupArn : undefined;
            inputs["replicationSubnetGroupDescription"] = state ? state.replicationSubnetGroupDescription : undefined;
            inputs["replicationSubnetGroupId"] = state ? state.replicationSubnetGroupId : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ReplicationSubnetGroupArgs | undefined;
            if (!args || args.replicationSubnetGroupDescription === undefined) {
                throw new Error("Missing required property 'replicationSubnetGroupDescription'");
            }
            if (!args || args.replicationSubnetGroupId === undefined) {
                throw new Error("Missing required property 'replicationSubnetGroupId'");
            }
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            inputs["replicationSubnetGroupDescription"] = args ? args.replicationSubnetGroupDescription : undefined;
            inputs["replicationSubnetGroupId"] = args ? args.replicationSubnetGroupId : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["replicationSubnetGroupArn"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReplicationSubnetGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicationSubnetGroup resources.
 */
export interface ReplicationSubnetGroupState {
    readonly replicationSubnetGroupArn?: pulumi.Input<string>;
    /**
     * The description for the subnet group.
     */
    readonly replicationSubnetGroupDescription?: pulumi.Input<string>;
    /**
     * The name for the replication subnet group. This value is stored as a lowercase string.
     */
    readonly replicationSubnetGroupId?: pulumi.Input<string>;
    /**
     * A list of the EC2 subnet IDs for the subnet group.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the VPC the subnet group is in.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicationSubnetGroup resource.
 */
export interface ReplicationSubnetGroupArgs {
    /**
     * The description for the subnet group.
     */
    readonly replicationSubnetGroupDescription: pulumi.Input<string>;
    /**
     * The name for the replication subnet group. This value is stored as a lowercase string.
     */
    readonly replicationSubnetGroupId: pulumi.Input<string>;
    /**
     * A list of the EC2 subnet IDs for the subnet group.
     */
    readonly subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
