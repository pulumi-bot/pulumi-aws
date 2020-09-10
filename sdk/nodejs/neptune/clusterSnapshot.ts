// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ClusterSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing ClusterSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterSnapshotState, opts?: pulumi.CustomResourceOptions): ClusterSnapshot {
        return new ClusterSnapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:neptune/clusterSnapshot:ClusterSnapshot';

    /**
     * Returns true if the given object is an instance of ClusterSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterSnapshot.__pulumiType;
    }

    public /*out*/ readonly allocatedStorage!: pulumi.Output<number>;
    public /*out*/ readonly availabilityZones!: pulumi.Output<string[]>;
    public readonly dbClusterIdentifier!: pulumi.Output<string>;
    public /*out*/ readonly dbClusterSnapshotArn!: pulumi.Output<string>;
    public readonly dbClusterSnapshotIdentifier!: pulumi.Output<string>;
    public /*out*/ readonly engine!: pulumi.Output<string>;
    public /*out*/ readonly engineVersion!: pulumi.Output<string>;
    public /*out*/ readonly kmsKeyId!: pulumi.Output<string>;
    public /*out*/ readonly licenseModel!: pulumi.Output<string>;
    public /*out*/ readonly port!: pulumi.Output<number>;
    public /*out*/ readonly snapshotType!: pulumi.Output<string>;
    public /*out*/ readonly sourceDbClusterSnapshotArn!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public /*out*/ readonly storageEncrypted!: pulumi.Output<boolean>;
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a ClusterSnapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterSnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterSnapshotArgs | ClusterSnapshotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterSnapshotState | undefined;
            inputs["allocatedStorage"] = state ? state.allocatedStorage : undefined;
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["dbClusterIdentifier"] = state ? state.dbClusterIdentifier : undefined;
            inputs["dbClusterSnapshotArn"] = state ? state.dbClusterSnapshotArn : undefined;
            inputs["dbClusterSnapshotIdentifier"] = state ? state.dbClusterSnapshotIdentifier : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["licenseModel"] = state ? state.licenseModel : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["snapshotType"] = state ? state.snapshotType : undefined;
            inputs["sourceDbClusterSnapshotArn"] = state ? state.sourceDbClusterSnapshotArn : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClusterSnapshotArgs | undefined;
            if (!args || args.dbClusterIdentifier === undefined) {
                throw new Error("Missing required property 'dbClusterIdentifier'");
            }
            if (!args || args.dbClusterSnapshotIdentifier === undefined) {
                throw new Error("Missing required property 'dbClusterSnapshotIdentifier'");
            }
            inputs["dbClusterIdentifier"] = args ? args.dbClusterIdentifier : undefined;
            inputs["dbClusterSnapshotIdentifier"] = args ? args.dbClusterSnapshotIdentifier : undefined;
            inputs["allocatedStorage"] = undefined /*out*/;
            inputs["availabilityZones"] = undefined /*out*/;
            inputs["dbClusterSnapshotArn"] = undefined /*out*/;
            inputs["engine"] = undefined /*out*/;
            inputs["engineVersion"] = undefined /*out*/;
            inputs["kmsKeyId"] = undefined /*out*/;
            inputs["licenseModel"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["snapshotType"] = undefined /*out*/;
            inputs["sourceDbClusterSnapshotArn"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["storageEncrypted"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterSnapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterSnapshot resources.
 */
export interface ClusterSnapshotState {
    readonly allocatedStorage?: pulumi.Input<number>;
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly dbClusterIdentifier?: pulumi.Input<string>;
    readonly dbClusterSnapshotArn?: pulumi.Input<string>;
    readonly dbClusterSnapshotIdentifier?: pulumi.Input<string>;
    readonly engine?: pulumi.Input<string>;
    readonly engineVersion?: pulumi.Input<string>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly licenseModel?: pulumi.Input<string>;
    readonly port?: pulumi.Input<number>;
    readonly snapshotType?: pulumi.Input<string>;
    readonly sourceDbClusterSnapshotArn?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly storageEncrypted?: pulumi.Input<boolean>;
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterSnapshot resource.
 */
export interface ClusterSnapshotArgs {
    readonly dbClusterIdentifier: pulumi.Input<string>;
    readonly dbClusterSnapshotIdentifier: pulumi.Input<string>;
}
