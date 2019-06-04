// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DocDB database cluster snapshot for DocDB clusters.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.docdb.ClusterSnapshot("example", {
 *     dbClusterIdentifier: aws_docdb_cluster_example.id,
 *     dbClusterSnapshotIdentifier: "resourcetestsnapshot1234",
 * });
 * ```
 */
export class ClusterSnapshot extends pulumi.CustomResource {
    /**
     * Get an existing ClusterSnapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterSnapshotState, opts?: pulumi.CustomResourceOptions): ClusterSnapshot {
        return new ClusterSnapshot(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:docdb/clusterSnapshot:ClusterSnapshot';

    /**
     * Returns true if the given object is an instance of ClusterSnapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterSnapshot {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === ClusterSnapshot.__pulumiType;
    }

    /**
     * List of EC2 Availability Zones that instances in the DocDB cluster snapshot can be restored in.
     */
    public /*out*/ readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * The DocDB Cluster Identifier from which to take the snapshot.
     */
    public readonly dbClusterIdentifier!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) for the DocDB Cluster Snapshot.
     */
    public /*out*/ readonly dbClusterSnapshotArn!: pulumi.Output<string>;
    /**
     * The Identifier for the snapshot.
     */
    public readonly dbClusterSnapshotIdentifier!: pulumi.Output<string>;
    /**
     * Specifies the name of the database engine.
     */
    public /*out*/ readonly engine!: pulumi.Output<string>;
    /**
     * Version of the database engine for this DocDB cluster snapshot.
     */
    public /*out*/ readonly engineVersion!: pulumi.Output<string>;
    /**
     * If storage_encrypted is true, the AWS KMS key identifier for the encrypted DocDB cluster snapshot.
     */
    public /*out*/ readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Port that the DocDB cluster was listening on at the time of the snapshot.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    public /*out*/ readonly snapshotType!: pulumi.Output<string>;
    public /*out*/ readonly sourceDbClusterSnapshotArn!: pulumi.Output<string>;
    /**
     * The status of this DocDB Cluster Snapshot.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies whether the DocDB cluster snapshot is encrypted.
     */
    public /*out*/ readonly storageEncrypted!: pulumi.Output<boolean>;
    /**
     * The VPC ID associated with the DocDB cluster snapshot.
     */
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
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["dbClusterIdentifier"] = state ? state.dbClusterIdentifier : undefined;
            inputs["dbClusterSnapshotArn"] = state ? state.dbClusterSnapshotArn : undefined;
            inputs["dbClusterSnapshotIdentifier"] = state ? state.dbClusterSnapshotIdentifier : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
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
            inputs["availabilityZones"] = undefined /*out*/;
            inputs["dbClusterSnapshotArn"] = undefined /*out*/;
            inputs["engine"] = undefined /*out*/;
            inputs["engineVersion"] = undefined /*out*/;
            inputs["kmsKeyId"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["snapshotType"] = undefined /*out*/;
            inputs["sourceDbClusterSnapshotArn"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["storageEncrypted"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        super(ClusterSnapshot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterSnapshot resources.
 */
export interface ClusterSnapshotState {
    /**
     * List of EC2 Availability Zones that instances in the DocDB cluster snapshot can be restored in.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The DocDB Cluster Identifier from which to take the snapshot.
     */
    readonly dbClusterIdentifier?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the DocDB Cluster Snapshot.
     */
    readonly dbClusterSnapshotArn?: pulumi.Input<string>;
    /**
     * The Identifier for the snapshot.
     */
    readonly dbClusterSnapshotIdentifier?: pulumi.Input<string>;
    /**
     * Specifies the name of the database engine.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * Version of the database engine for this DocDB cluster snapshot.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * If storage_encrypted is true, the AWS KMS key identifier for the encrypted DocDB cluster snapshot.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Port that the DocDB cluster was listening on at the time of the snapshot.
     */
    readonly port?: pulumi.Input<number>;
    readonly snapshotType?: pulumi.Input<string>;
    readonly sourceDbClusterSnapshotArn?: pulumi.Input<string>;
    /**
     * The status of this DocDB Cluster Snapshot.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Specifies whether the DocDB cluster snapshot is encrypted.
     */
    readonly storageEncrypted?: pulumi.Input<boolean>;
    /**
     * The VPC ID associated with the DocDB cluster snapshot.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterSnapshot resource.
 */
export interface ClusterSnapshotArgs {
    /**
     * The DocDB Cluster Identifier from which to take the snapshot.
     */
    readonly dbClusterIdentifier: pulumi.Input<string>;
    /**
     * The Identifier for the snapshot.
     */
    readonly dbClusterSnapshotIdentifier: pulumi.Input<string>;
}
