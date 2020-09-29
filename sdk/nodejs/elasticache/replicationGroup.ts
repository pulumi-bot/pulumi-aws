// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache Replication Group resource.
 * For working with Memcached or single primary Redis instances (Cluster Mode Disabled), see the
 * `aws.elasticache.Cluster` resource.
 *
 * > **Note:** When you change an attribute, such as `engineVersion`, by
 * default the ElastiCache API applies it in the next maintenance window. Because
 * of this, this provider may report a difference in its planning phase because the
 * actual modification has not yet taken place. You can use the
 * `applyImmediately` flag to instruct the service to apply the change
 * immediately. Using `applyImmediately` can result in a brief downtime as
 * servers reboots.
 *
 * ## Example Usage
 * ### Redis Cluster Mode Disabled
 *
 * To create a single shard primary with single read replica:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.elasticache.ReplicationGroup("example", {
 *     automaticFailoverEnabled: true,
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *     ],
 *     nodeType: "cache.m4.large",
 *     numberCacheClusters: 2,
 *     parameterGroupName: "default.redis3.2",
 *     port: 6379,
 *     replicationGroupDescription: "test description",
 * });
 * ```
 *
 * You have two options for adjusting the number of replicas:
 *
 * * Adjusting `numberCacheClusters` directly. This will attempt to automatically add or remove replicas, but provides no granular control (e.g. preferred availability zone, cache cluster ID) for the added or removed replicas. This also currently expects cache cluster IDs in the form of `replication_group_id-00#`.
 * * Otherwise for fine grained control of the underlying cache clusters, they can be added or removed with the `aws.elasticache.Cluster` resource and its `replicationGroupId` attribute. In this situation, you will need to utilize [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to prevent perpetual differences with the `numberCacheCluster` attribute.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.elasticache.ReplicationGroup("example", {
 *     automaticFailoverEnabled: true,
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *     ],
 *     replicationGroupDescription: "test description",
 *     nodeType: "cache.m4.large",
 *     numberCacheClusters: 2,
 *     parameterGroupName: "default.redis3.2",
 *     port: 6379,
 * });
 * let replica: aws.elasticache.Cluster | undefined;
 * if (1 == true) {
 *     replica = new aws.elasticache.Cluster("replica", {replicationGroupId: example.id});
 * }
 * ```
 * ### Redis Cluster Mode Enabled
 *
 * To create two shards with a primary and a single read replica each:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const baz = new aws.elasticache.ReplicationGroup("baz", {
 *     automaticFailoverEnabled: true,
 *     clusterMode: {
 *         numNodeGroups: 2,
 *         replicasPerNodeGroup: 1,
 *     },
 *     nodeType: "cache.t2.small",
 *     parameterGroupName: "default.redis3.2.cluster.on",
 *     port: 6379,
 *     replicationGroupDescription: "test description",
 * });
 * ```
 *
 * > **Note:** We currently do not support passing a `primaryClusterId` in order to create the Replication Group.
 *
 * > **Note:** Automatic Failover is unavailable for Redis versions earlier than 2.8.6,
 * and unavailable on T1 node types. For T2 node types, it is only available on Redis version 3.2.4 or later with cluster mode enabled. See the [High Availability Using Replication Groups](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Replication.html) guide
 * for full details on using Replication Groups.
 */
export class ReplicationGroup extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicationGroupState, opts?: pulumi.CustomResourceOptions): ReplicationGroup {
        return new ReplicationGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticache/replicationGroup:ReplicationGroup';

    /**
     * Returns true if the given object is an instance of ReplicationGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationGroup.__pulumiType;
    }

    /**
     * Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
     */
    public readonly applyImmediately!: pulumi.Output<boolean>;
    /**
     * Whether to enable encryption at rest.
     */
    public readonly atRestEncryptionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The password used to access a password protected server. Can be specified only if `transitEncryptionEnabled = true`.
     */
    public readonly authToken!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
     */
    public readonly automaticFailoverEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
     */
    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    /**
     * Create a native redis cluster. `automaticFailoverEnabled` must be set to true. Cluster Mode documented below. Only 1 `clusterMode` block is allowed.
     */
    public readonly clusterMode!: pulumi.Output<outputs.elasticache.ReplicationGroupClusterMode>;
    /**
     * The address of the replication group configuration endpoint when cluster mode is enabled.
     */
    public /*out*/ readonly configurationEndpointAddress!: pulumi.Output<string>;
    /**
     * The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * The version number of the cache engine to be used for the cache clusters in this replication group.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `atRestEncryptionEnabled = true`.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the weekly time range for when maintenance
     * on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
     * The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
     */
    public readonly maintenanceWindow!: pulumi.Output<string>;
    /**
     * The identifiers of all the nodes that are part of this replication group.
     */
    public /*out*/ readonly memberClusters!: pulumi.Output<string[]>;
    /**
     * The compute and memory capacity of the nodes in the node group.
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * An Amazon Resource Name (ARN) of an
     * SNS topic to send ElastiCache notifications to. Example:
     * `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
     */
    public readonly notificationTopicArn!: pulumi.Output<string | undefined>;
    /**
     * The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
     */
    public readonly numberCacheClusters!: pulumi.Output<number>;
    /**
     * The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
     */
    public readonly parameterGroupName!: pulumi.Output<string>;
    /**
     * The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * (Redis only) The address of the endpoint for the primary node in the replication group, if the cluster mode is disabled.
     */
    public /*out*/ readonly primaryEndpointAddress!: pulumi.Output<string>;
    /**
     * A user-created description for the replication group.
     */
    public readonly replicationGroupDescription!: pulumi.Output<string>;
    /**
     * The replication group identifier. This parameter is stored as a lowercase string.
     */
    public readonly replicationGroupId!: pulumi.Output<string>;
    /**
     * One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * A list of cache security group names to associate with this replication group.
     */
    public readonly securityGroupNames!: pulumi.Output<string[]>;
    /**
     * A single-element string list containing an
     * Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
     * Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
     */
    public readonly snapshotArns!: pulumi.Output<string[] | undefined>;
    /**
     * The name of a snapshot from which to restore data into the new node group. Changing the `snapshotName` forces a new resource.
     */
    public readonly snapshotName!: pulumi.Output<string | undefined>;
    /**
     * The number of days for which ElastiCache will
     * retain automatic cache cluster snapshots before deleting them. For example, if you set
     * SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
     * before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
     * Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
     */
    public readonly snapshotRetentionLimit!: pulumi.Output<number | undefined>;
    /**
     * The daily time range (in UTC) during which ElastiCache will
     * begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
     */
    public readonly snapshotWindow!: pulumi.Output<string>;
    /**
     * The name of the cache subnet group to be used for the replication group.
     */
    public readonly subnetGroupName!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. Adding tags to this resource will add or overwrite any existing tags on the clusters in the replication group and not to the group itself.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Whether to enable encryption in transit.
     */
    public readonly transitEncryptionEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ReplicationGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationGroupArgs | ReplicationGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReplicationGroupState | undefined;
            inputs["applyImmediately"] = state ? state.applyImmediately : undefined;
            inputs["atRestEncryptionEnabled"] = state ? state.atRestEncryptionEnabled : undefined;
            inputs["authToken"] = state ? state.authToken : undefined;
            inputs["autoMinorVersionUpgrade"] = state ? state.autoMinorVersionUpgrade : undefined;
            inputs["automaticFailoverEnabled"] = state ? state.automaticFailoverEnabled : undefined;
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["clusterMode"] = state ? state.clusterMode : undefined;
            inputs["configurationEndpointAddress"] = state ? state.configurationEndpointAddress : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["maintenanceWindow"] = state ? state.maintenanceWindow : undefined;
            inputs["memberClusters"] = state ? state.memberClusters : undefined;
            inputs["nodeType"] = state ? state.nodeType : undefined;
            inputs["notificationTopicArn"] = state ? state.notificationTopicArn : undefined;
            inputs["numberCacheClusters"] = state ? state.numberCacheClusters : undefined;
            inputs["parameterGroupName"] = state ? state.parameterGroupName : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["primaryEndpointAddress"] = state ? state.primaryEndpointAddress : undefined;
            inputs["replicationGroupDescription"] = state ? state.replicationGroupDescription : undefined;
            inputs["replicationGroupId"] = state ? state.replicationGroupId : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["securityGroupNames"] = state ? state.securityGroupNames : undefined;
            inputs["snapshotArns"] = state ? state.snapshotArns : undefined;
            inputs["snapshotName"] = state ? state.snapshotName : undefined;
            inputs["snapshotRetentionLimit"] = state ? state.snapshotRetentionLimit : undefined;
            inputs["snapshotWindow"] = state ? state.snapshotWindow : undefined;
            inputs["subnetGroupName"] = state ? state.subnetGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["transitEncryptionEnabled"] = state ? state.transitEncryptionEnabled : undefined;
        } else {
            const args = argsOrState as ReplicationGroupArgs | undefined;
            if (!args || args.replicationGroupDescription === undefined) {
                throw new Error("Missing required property 'replicationGroupDescription'");
            }
            inputs["applyImmediately"] = args ? args.applyImmediately : undefined;
            inputs["atRestEncryptionEnabled"] = args ? args.atRestEncryptionEnabled : undefined;
            inputs["authToken"] = args ? args.authToken : undefined;
            inputs["autoMinorVersionUpgrade"] = args ? args.autoMinorVersionUpgrade : undefined;
            inputs["automaticFailoverEnabled"] = args ? args.automaticFailoverEnabled : undefined;
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["clusterMode"] = args ? args.clusterMode : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            inputs["nodeType"] = args ? args.nodeType : undefined;
            inputs["notificationTopicArn"] = args ? args.notificationTopicArn : undefined;
            inputs["numberCacheClusters"] = args ? args.numberCacheClusters : undefined;
            inputs["parameterGroupName"] = args ? args.parameterGroupName : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["replicationGroupDescription"] = args ? args.replicationGroupDescription : undefined;
            inputs["replicationGroupId"] = args ? args.replicationGroupId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["securityGroupNames"] = args ? args.securityGroupNames : undefined;
            inputs["snapshotArns"] = args ? args.snapshotArns : undefined;
            inputs["snapshotName"] = args ? args.snapshotName : undefined;
            inputs["snapshotRetentionLimit"] = args ? args.snapshotRetentionLimit : undefined;
            inputs["snapshotWindow"] = args ? args.snapshotWindow : undefined;
            inputs["subnetGroupName"] = args ? args.subnetGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitEncryptionEnabled"] = args ? args.transitEncryptionEnabled : undefined;
            inputs["configurationEndpointAddress"] = undefined /*out*/;
            inputs["memberClusters"] = undefined /*out*/;
            inputs["primaryEndpointAddress"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReplicationGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicationGroup resources.
 */
export interface ReplicationGroupState {
    /**
     * Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * Whether to enable encryption at rest.
     */
    readonly atRestEncryptionEnabled?: pulumi.Input<boolean>;
    /**
     * The password used to access a password protected server. Can be specified only if `transitEncryptionEnabled = true`.
     */
    readonly authToken?: pulumi.Input<string>;
    /**
     * Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
     */
    readonly autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
     */
    readonly automaticFailoverEnabled?: pulumi.Input<boolean>;
    /**
     * A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Create a native redis cluster. `automaticFailoverEnabled` must be set to true. Cluster Mode documented below. Only 1 `clusterMode` block is allowed.
     */
    readonly clusterMode?: pulumi.Input<inputs.elasticache.ReplicationGroupClusterMode>;
    /**
     * The address of the replication group configuration endpoint when cluster mode is enabled.
     */
    readonly configurationEndpointAddress?: pulumi.Input<string>;
    /**
     * The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * The version number of the cache engine to be used for the cache clusters in this replication group.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `atRestEncryptionEnabled = true`.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the weekly time range for when maintenance
     * on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
     * The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
     */
    readonly maintenanceWindow?: pulumi.Input<string>;
    /**
     * The identifiers of all the nodes that are part of this replication group.
     */
    readonly memberClusters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The compute and memory capacity of the nodes in the node group.
     */
    readonly nodeType?: pulumi.Input<string>;
    /**
     * An Amazon Resource Name (ARN) of an
     * SNS topic to send ElastiCache notifications to. Example:
     * `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
     */
    readonly notificationTopicArn?: pulumi.Input<string>;
    /**
     * The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
     */
    readonly numberCacheClusters?: pulumi.Input<number>;
    /**
     * The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
     */
    readonly parameterGroupName?: pulumi.Input<string>;
    /**
     * The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * (Redis only) The address of the endpoint for the primary node in the replication group, if the cluster mode is disabled.
     */
    readonly primaryEndpointAddress?: pulumi.Input<string>;
    /**
     * A user-created description for the replication group.
     */
    readonly replicationGroupDescription?: pulumi.Input<string>;
    /**
     * The replication group identifier. This parameter is stored as a lowercase string.
     */
    readonly replicationGroupId?: pulumi.Input<string>;
    /**
     * One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of cache security group names to associate with this replication group.
     */
    readonly securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A single-element string list containing an
     * Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
     * Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
     */
    readonly snapshotArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of a snapshot from which to restore data into the new node group. Changing the `snapshotName` forces a new resource.
     */
    readonly snapshotName?: pulumi.Input<string>;
    /**
     * The number of days for which ElastiCache will
     * retain automatic cache cluster snapshots before deleting them. For example, if you set
     * SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
     * before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
     * Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
     */
    readonly snapshotRetentionLimit?: pulumi.Input<number>;
    /**
     * The daily time range (in UTC) during which ElastiCache will
     * begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
     */
    readonly snapshotWindow?: pulumi.Input<string>;
    /**
     * The name of the cache subnet group to be used for the replication group.
     */
    readonly subnetGroupName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. Adding tags to this resource will add or overwrite any existing tags on the clusters in the replication group and not to the group itself.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to enable encryption in transit.
     */
    readonly transitEncryptionEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ReplicationGroup resource.
 */
export interface ReplicationGroupArgs {
    /**
     * Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * Whether to enable encryption at rest.
     */
    readonly atRestEncryptionEnabled?: pulumi.Input<boolean>;
    /**
     * The password used to access a password protected server. Can be specified only if `transitEncryptionEnabled = true`.
     */
    readonly authToken?: pulumi.Input<string>;
    /**
     * Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
     */
    readonly autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
     */
    readonly automaticFailoverEnabled?: pulumi.Input<boolean>;
    /**
     * A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
     */
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Create a native redis cluster. `automaticFailoverEnabled` must be set to true. Cluster Mode documented below. Only 1 `clusterMode` block is allowed.
     */
    readonly clusterMode?: pulumi.Input<inputs.elasticache.ReplicationGroupClusterMode>;
    /**
     * The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * The version number of the cache engine to be used for the cache clusters in this replication group.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `atRestEncryptionEnabled = true`.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the weekly time range for when maintenance
     * on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
     * The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
     */
    readonly maintenanceWindow?: pulumi.Input<string>;
    /**
     * The compute and memory capacity of the nodes in the node group.
     */
    readonly nodeType?: pulumi.Input<string>;
    /**
     * An Amazon Resource Name (ARN) of an
     * SNS topic to send ElastiCache notifications to. Example:
     * `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
     */
    readonly notificationTopicArn?: pulumi.Input<string>;
    /**
     * The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
     */
    readonly numberCacheClusters?: pulumi.Input<number>;
    /**
     * The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
     */
    readonly parameterGroupName?: pulumi.Input<string>;
    /**
     * The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * A user-created description for the replication group.
     */
    readonly replicationGroupDescription: pulumi.Input<string>;
    /**
     * The replication group identifier. This parameter is stored as a lowercase string.
     */
    readonly replicationGroupId?: pulumi.Input<string>;
    /**
     * One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of cache security group names to associate with this replication group.
     */
    readonly securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A single-element string list containing an
     * Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
     * Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
     */
    readonly snapshotArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of a snapshot from which to restore data into the new node group. Changing the `snapshotName` forces a new resource.
     */
    readonly snapshotName?: pulumi.Input<string>;
    /**
     * The number of days for which ElastiCache will
     * retain automatic cache cluster snapshots before deleting them. For example, if you set
     * SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
     * before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
     * Please note that setting a `snapshotRetentionLimit` is not supported on cache.t1.micro or cache.t2.* cache nodes
     */
    readonly snapshotRetentionLimit?: pulumi.Input<number>;
    /**
     * The daily time range (in UTC) during which ElastiCache will
     * begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
     */
    readonly snapshotWindow?: pulumi.Input<string>;
    /**
     * The name of the cache subnet group to be used for the replication group.
     */
    readonly subnetGroupName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. Adding tags to this resource will add or overwrite any existing tags on the clusters in the replication group and not to the group itself.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to enable encryption in transit.
     */
    readonly transitEncryptionEnabled?: pulumi.Input<boolean>;
}
