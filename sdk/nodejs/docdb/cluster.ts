// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DocDB Cluster.
 * 
 * Changes to a DocDB Cluster can occur when you manually change a
 * parameter, such as `port`, and are reflected in the next maintenance
 * window. Because of this, Terraform may report a difference in its planning
 * phase because a modification has not yet taken place. You can use the
 * `apply_immediately` flag to instruct the service to apply the change immediately
 * (see documentation below).
 * 
 * > **Note:** using `apply_immediately` can result in a brief downtime as the server reboots.
 * > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const docdb = new aws.docdb.Cluster("docdb", {
 *     backupRetentionPeriod: 5,
 *     clusterIdentifier: "my-docdb-cluster",
 *     engine: "docdb",
 *     masterPassword: "mustbeeightchars",
 *     masterUsername: "foo",
 *     preferredBackupWindow: "07:00-09:00",
 *     skipFinalSnapshot: true,
 * });
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:docdb/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === Cluster.__pulumiType;
    }

    /**
     * Specifies whether any cluster modifications
     * are applied immediately, or during the next maintenance window. Default is
     * `false`.
     */
    public readonly applyImmediately!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of cluster
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of EC2 Availability Zones that
     * instances in the DB cluster can be created in.
     */
    public readonly availabilityZones!: pulumi.Output<string[]>;
    /**
     * The days to retain backups for. Default `1`
     */
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * The cluster identifier. If omitted, Terraform will assign a random, unique identifier.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifer`.
     */
    public readonly clusterIdentifierPrefix!: pulumi.Output<string>;
    /**
     * List of DocDB Instances that are a part of this cluster
     */
    public readonly clusterMembers!: pulumi.Output<string[]>;
    /**
     * The DocDB Cluster Resource ID
     */
    public /*out*/ readonly clusterResourceId!: pulumi.Output<string>;
    /**
     * A cluster parameter group to associate with the cluster.
     */
    public readonly dbClusterParameterGroupName!: pulumi.Output<string>;
    /**
     * A DB subnet group to associate with this DB instance.
     */
    public readonly dbSubnetGroupName!: pulumi.Output<string>;
    /**
     * List of log types to export to cloudwatch. If omitted, no logs will be exported.
     * The following log types are supported: `audit`.
     */
    public readonly enabledCloudwatchLogsExports!: pulumi.Output<string[] | undefined>;
    /**
     * The DNS address of the DocDB instance
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid Values: `docdb`
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * The database engine version. Updating this argument results in an outage.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The name of your final DB snapshot
     * when this DB cluster is deleted. If omitted, no final snapshot will be
     * made.
     */
    public readonly finalSnapshotIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The Route53 Hosted Zone ID of the endpoint
     */
    public /*out*/ readonly hostedZoneId!: pulumi.Output<string>;
    /**
     * The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Password for the master DB user. Note that this may
     * show up in logs, and it will be stored in the state file. Please refer to the DocDB Naming Constraints.
     */
    public readonly masterPassword!: pulumi.Output<string | undefined>;
    /**
     * Username for the master DB user. 
     */
    public readonly masterUsername!: pulumi.Output<string>;
    /**
     * The port on which the DB accepts connections
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
     * Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
     */
    public readonly preferredBackupWindow!: pulumi.Output<string>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string>;
    /**
     * A read-only endpoint for the DocDB cluster, automatically load-balanced across replicas
     */
    public /*out*/ readonly readerEndpoint!: pulumi.Output<string>;
    /**
     * Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
     */
    public readonly skipFinalSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
     */
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false`.
     */
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the DB cluster.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * List of VPC security groups to associate
     * with the Cluster
     */
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[]>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["applyImmediately"] = state ? state.applyImmediately : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["backupRetentionPeriod"] = state ? state.backupRetentionPeriod : undefined;
            inputs["clusterIdentifier"] = state ? state.clusterIdentifier : undefined;
            inputs["clusterIdentifierPrefix"] = state ? state.clusterIdentifierPrefix : undefined;
            inputs["clusterMembers"] = state ? state.clusterMembers : undefined;
            inputs["clusterResourceId"] = state ? state.clusterResourceId : undefined;
            inputs["dbClusterParameterGroupName"] = state ? state.dbClusterParameterGroupName : undefined;
            inputs["dbSubnetGroupName"] = state ? state.dbSubnetGroupName : undefined;
            inputs["enabledCloudwatchLogsExports"] = state ? state.enabledCloudwatchLogsExports : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["finalSnapshotIdentifier"] = state ? state.finalSnapshotIdentifier : undefined;
            inputs["hostedZoneId"] = state ? state.hostedZoneId : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["masterPassword"] = state ? state.masterPassword : undefined;
            inputs["masterUsername"] = state ? state.masterUsername : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["preferredBackupWindow"] = state ? state.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = state ? state.preferredMaintenanceWindow : undefined;
            inputs["readerEndpoint"] = state ? state.readerEndpoint : undefined;
            inputs["skipFinalSnapshot"] = state ? state.skipFinalSnapshot : undefined;
            inputs["snapshotIdentifier"] = state ? state.snapshotIdentifier : undefined;
            inputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcSecurityGroupIds"] = state ? state.vpcSecurityGroupIds : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["applyImmediately"] = args ? args.applyImmediately : undefined;
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["backupRetentionPeriod"] = args ? args.backupRetentionPeriod : undefined;
            inputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            inputs["clusterIdentifierPrefix"] = args ? args.clusterIdentifierPrefix : undefined;
            inputs["clusterMembers"] = args ? args.clusterMembers : undefined;
            inputs["dbClusterParameterGroupName"] = args ? args.dbClusterParameterGroupName : undefined;
            inputs["dbSubnetGroupName"] = args ? args.dbSubnetGroupName : undefined;
            inputs["enabledCloudwatchLogsExports"] = args ? args.enabledCloudwatchLogsExports : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["finalSnapshotIdentifier"] = args ? args.finalSnapshotIdentifier : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["masterPassword"] = args ? args.masterPassword : undefined;
            inputs["masterUsername"] = args ? args.masterUsername : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            inputs["skipFinalSnapshot"] = args ? args.skipFinalSnapshot : undefined;
            inputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            inputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["clusterResourceId"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["hostedZoneId"] = undefined /*out*/;
            inputs["readerEndpoint"] = undefined /*out*/;
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Specifies whether any cluster modifications
     * are applied immediately, or during the next maintenance window. Default is
     * `false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of cluster
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A list of EC2 Availability Zones that
     * instances in the DB cluster can be created in.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The days to retain backups for. Default `1`
     */
    readonly backupRetentionPeriod?: pulumi.Input<number>;
    /**
     * The cluster identifier. If omitted, Terraform will assign a random, unique identifier.
     */
    readonly clusterIdentifier?: pulumi.Input<string>;
    /**
     * Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifer`.
     */
    readonly clusterIdentifierPrefix?: pulumi.Input<string>;
    /**
     * List of DocDB Instances that are a part of this cluster
     */
    readonly clusterMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The DocDB Cluster Resource ID
     */
    readonly clusterResourceId?: pulumi.Input<string>;
    /**
     * A cluster parameter group to associate with the cluster.
     */
    readonly dbClusterParameterGroupName?: pulumi.Input<string>;
    /**
     * A DB subnet group to associate with this DB instance.
     */
    readonly dbSubnetGroupName?: pulumi.Input<string>;
    /**
     * List of log types to export to cloudwatch. If omitted, no logs will be exported.
     * The following log types are supported: `audit`.
     */
    readonly enabledCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The DNS address of the DocDB instance
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid Values: `docdb`
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * The database engine version. Updating this argument results in an outage.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The name of your final DB snapshot
     * when this DB cluster is deleted. If omitted, no final snapshot will be
     * made.
     */
    readonly finalSnapshotIdentifier?: pulumi.Input<string>;
    /**
     * The Route53 Hosted Zone ID of the endpoint
     */
    readonly hostedZoneId?: pulumi.Input<string>;
    /**
     * The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Password for the master DB user. Note that this may
     * show up in logs, and it will be stored in the state file. Please refer to the DocDB Naming Constraints.
     */
    readonly masterPassword?: pulumi.Input<string>;
    /**
     * Username for the master DB user. 
     */
    readonly masterUsername?: pulumi.Input<string>;
    /**
     * The port on which the DB accepts connections
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
     * Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
     */
    readonly preferredBackupWindow?: pulumi.Input<string>;
    readonly preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * A read-only endpoint for the DocDB cluster, automatically load-balanced across replicas
     */
    readonly readerEndpoint?: pulumi.Input<string>;
    /**
     * Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
     */
    readonly skipFinalSnapshot?: pulumi.Input<boolean>;
    /**
     * Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
     */
    readonly snapshotIdentifier?: pulumi.Input<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false`.
     */
    readonly storageEncrypted?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the DB cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * List of VPC security groups to associate
     * with the Cluster
     */
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Specifies whether any cluster modifications
     * are applied immediately, or during the next maintenance window. Default is
     * `false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * A list of EC2 Availability Zones that
     * instances in the DB cluster can be created in.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The days to retain backups for. Default `1`
     */
    readonly backupRetentionPeriod?: pulumi.Input<number>;
    /**
     * The cluster identifier. If omitted, Terraform will assign a random, unique identifier.
     */
    readonly clusterIdentifier?: pulumi.Input<string>;
    /**
     * Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifer`.
     */
    readonly clusterIdentifierPrefix?: pulumi.Input<string>;
    /**
     * List of DocDB Instances that are a part of this cluster
     */
    readonly clusterMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A cluster parameter group to associate with the cluster.
     */
    readonly dbClusterParameterGroupName?: pulumi.Input<string>;
    /**
     * A DB subnet group to associate with this DB instance.
     */
    readonly dbSubnetGroupName?: pulumi.Input<string>;
    /**
     * List of log types to export to cloudwatch. If omitted, no logs will be exported.
     * The following log types are supported: `audit`.
     */
    readonly enabledCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the database engine to be used for this DB cluster. Defaults to `docdb`. Valid Values: `docdb`
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * The database engine version. Updating this argument results in an outage.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The name of your final DB snapshot
     * when this DB cluster is deleted. If omitted, no final snapshot will be
     * made.
     */
    readonly finalSnapshotIdentifier?: pulumi.Input<string>;
    /**
     * The ARN for the KMS encryption key. When specifying `kms_key_id`, `storage_encrypted` needs to be set to true.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Password for the master DB user. Note that this may
     * show up in logs, and it will be stored in the state file. Please refer to the DocDB Naming Constraints.
     */
    readonly masterPassword?: pulumi.Input<string>;
    /**
     * Username for the master DB user. 
     */
    readonly masterUsername?: pulumi.Input<string>;
    /**
     * The port on which the DB accepts connections
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.Time in UTC
     * Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
     */
    readonly preferredBackupWindow?: pulumi.Input<string>;
    readonly preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Determines whether a final DB snapshot is created before the DB cluster is deleted. If true is specified, no DB snapshot is created. If false is specified, a DB snapshot is created before the DB cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
     */
    readonly skipFinalSnapshot?: pulumi.Input<boolean>;
    /**
     * Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a DB cluster snapshot, or the ARN when specifying a DB snapshot.
     */
    readonly snapshotIdentifier?: pulumi.Input<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false`.
     */
    readonly storageEncrypted?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the DB cluster.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * List of VPC security groups to associate
     * with the Cluster
     */
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
