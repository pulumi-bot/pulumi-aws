// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Adds permission to create volumes off of a given EBS Snapshot.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/snapshot_create_volume_permission.html.markdown.
 */
export class SnapshotCreateVolumePermission extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotCreateVolumePermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotCreateVolumePermissionState, opts?: pulumi.CustomResourceOptions): SnapshotCreateVolumePermission {
        return new SnapshotCreateVolumePermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission';

    /**
     * Returns true if the given object is an instance of SnapshotCreateVolumePermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotCreateVolumePermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotCreateVolumePermission.__pulumiType;
    }

    /**
     * An AWS Account ID to add create volume permissions
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * A snapshot ID
     */
    public readonly snapshotId!: pulumi.Output<string>;

    /**
     * Create a SnapshotCreateVolumePermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotCreateVolumePermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotCreateVolumePermissionArgs | SnapshotCreateVolumePermissionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SnapshotCreateVolumePermissionState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
        } else {
            const args = argsOrState as SnapshotCreateVolumePermissionArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            if (!args || args.snapshotId === undefined) {
                throw new Error("Missing required property 'snapshotId'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SnapshotCreateVolumePermission.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotCreateVolumePermission resources.
 */
export interface SnapshotCreateVolumePermissionState {
    /**
     * An AWS Account ID to add create volume permissions
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * A snapshot ID
     */
    readonly snapshotId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnapshotCreateVolumePermission resource.
 */
export interface SnapshotCreateVolumePermissionArgs {
    /**
     * An AWS Account ID to add create volume permissions
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * A snapshot ID
     */
    readonly snapshotId: pulumi.Input<string>;
}
