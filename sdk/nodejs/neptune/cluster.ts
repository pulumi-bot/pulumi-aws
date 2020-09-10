// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:neptune/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public readonly applyImmediately!: pulumi.Output<boolean>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly availabilityZones!: pulumi.Output<string[]>;
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    public readonly clusterIdentifier!: pulumi.Output<string>;
    public readonly clusterIdentifierPrefix!: pulumi.Output<string>;
    public /*out*/ readonly clusterMembers!: pulumi.Output<string[]>;
    public /*out*/ readonly clusterResourceId!: pulumi.Output<string>;
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    public readonly enableCloudwatchLogsExports!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    public readonly engine!: pulumi.Output<string | undefined>;
    public readonly engineVersion!: pulumi.Output<string>;
    public readonly finalSnapshotIdentifier!: pulumi.Output<string | undefined>;
    public /*out*/ readonly hostedZoneId!: pulumi.Output<string>;
    public readonly iamDatabaseAuthenticationEnabled!: pulumi.Output<boolean | undefined>;
    public readonly iamRoles!: pulumi.Output<string[] | undefined>;
    public readonly kmsKeyArn!: pulumi.Output<string>;
    public readonly neptuneClusterParameterGroupName!: pulumi.Output<string | undefined>;
    public readonly neptuneSubnetGroupName!: pulumi.Output<string>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly preferredBackupWindow!: pulumi.Output<string>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string>;
    public /*out*/ readonly readerEndpoint!: pulumi.Output<string>;
    public readonly replicationSourceIdentifier!: pulumi.Output<string | undefined>;
    public readonly skipFinalSnapshot!: pulumi.Output<boolean | undefined>;
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
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
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["enableCloudwatchLogsExports"] = state ? state.enableCloudwatchLogsExports : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["finalSnapshotIdentifier"] = state ? state.finalSnapshotIdentifier : undefined;
            inputs["hostedZoneId"] = state ? state.hostedZoneId : undefined;
            inputs["iamDatabaseAuthenticationEnabled"] = state ? state.iamDatabaseAuthenticationEnabled : undefined;
            inputs["iamRoles"] = state ? state.iamRoles : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
            inputs["neptuneClusterParameterGroupName"] = state ? state.neptuneClusterParameterGroupName : undefined;
            inputs["neptuneSubnetGroupName"] = state ? state.neptuneSubnetGroupName : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["preferredBackupWindow"] = state ? state.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = state ? state.preferredMaintenanceWindow : undefined;
            inputs["readerEndpoint"] = state ? state.readerEndpoint : undefined;
            inputs["replicationSourceIdentifier"] = state ? state.replicationSourceIdentifier : undefined;
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
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["enableCloudwatchLogsExports"] = args ? args.enableCloudwatchLogsExports : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["finalSnapshotIdentifier"] = args ? args.finalSnapshotIdentifier : undefined;
            inputs["iamDatabaseAuthenticationEnabled"] = args ? args.iamDatabaseAuthenticationEnabled : undefined;
            inputs["iamRoles"] = args ? args.iamRoles : undefined;
            inputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            inputs["neptuneClusterParameterGroupName"] = args ? args.neptuneClusterParameterGroupName : undefined;
            inputs["neptuneSubnetGroupName"] = args ? args.neptuneSubnetGroupName : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            inputs["replicationSourceIdentifier"] = args ? args.replicationSourceIdentifier : undefined;
            inputs["skipFinalSnapshot"] = args ? args.skipFinalSnapshot : undefined;
            inputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            inputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["clusterMembers"] = undefined /*out*/;
            inputs["clusterResourceId"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["hostedZoneId"] = undefined /*out*/;
            inputs["readerEndpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    readonly applyImmediately?: pulumi.Input<boolean>;
    readonly arn?: pulumi.Input<string>;
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly backupRetentionPeriod?: pulumi.Input<number>;
    readonly clusterIdentifier?: pulumi.Input<string>;
    readonly clusterIdentifierPrefix?: pulumi.Input<string>;
    readonly clusterMembers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly clusterResourceId?: pulumi.Input<string>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly enableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    readonly endpoint?: pulumi.Input<string>;
    readonly engine?: pulumi.Input<string>;
    readonly engineVersion?: pulumi.Input<string>;
    readonly finalSnapshotIdentifier?: pulumi.Input<string>;
    readonly hostedZoneId?: pulumi.Input<string>;
    readonly iamDatabaseAuthenticationEnabled?: pulumi.Input<boolean>;
    readonly iamRoles?: pulumi.Input<pulumi.Input<string>[]>;
    readonly kmsKeyArn?: pulumi.Input<string>;
    readonly neptuneClusterParameterGroupName?: pulumi.Input<string>;
    readonly neptuneSubnetGroupName?: pulumi.Input<string>;
    readonly port?: pulumi.Input<number>;
    readonly preferredBackupWindow?: pulumi.Input<string>;
    readonly preferredMaintenanceWindow?: pulumi.Input<string>;
    readonly readerEndpoint?: pulumi.Input<string>;
    readonly replicationSourceIdentifier?: pulumi.Input<string>;
    readonly skipFinalSnapshot?: pulumi.Input<boolean>;
    readonly snapshotIdentifier?: pulumi.Input<string>;
    readonly storageEncrypted?: pulumi.Input<boolean>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    readonly applyImmediately?: pulumi.Input<boolean>;
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly backupRetentionPeriod?: pulumi.Input<number>;
    readonly clusterIdentifier?: pulumi.Input<string>;
    readonly clusterIdentifierPrefix?: pulumi.Input<string>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly enableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    readonly engine?: pulumi.Input<string>;
    readonly engineVersion?: pulumi.Input<string>;
    readonly finalSnapshotIdentifier?: pulumi.Input<string>;
    readonly iamDatabaseAuthenticationEnabled?: pulumi.Input<boolean>;
    readonly iamRoles?: pulumi.Input<pulumi.Input<string>[]>;
    readonly kmsKeyArn?: pulumi.Input<string>;
    readonly neptuneClusterParameterGroupName?: pulumi.Input<string>;
    readonly neptuneSubnetGroupName?: pulumi.Input<string>;
    readonly port?: pulumi.Input<number>;
    readonly preferredBackupWindow?: pulumi.Input<string>;
    readonly preferredMaintenanceWindow?: pulumi.Input<string>;
    readonly replicationSourceIdentifier?: pulumi.Input<string>;
    readonly skipFinalSnapshot?: pulumi.Input<boolean>;
    readonly snapshotIdentifier?: pulumi.Input<string>;
    readonly storageEncrypted?: pulumi.Input<boolean>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
