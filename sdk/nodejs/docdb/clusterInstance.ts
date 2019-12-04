// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
 * attributes that are specific to a single instance in a [DocDB Cluster][1].
 * 
 * You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
 * Instances and DocDB manages the replication. You can use the [count][3]
 * meta-parameter to make multiple instances and join them all to the same DocDB
 * Cluster, or you may specify different Cluster Instance resources with various
 * `instanceClass` sizes.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const defaultCluster = new aws.docdb.Cluster("default", {
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *         "us-west-2c",
 *     ],
 *     clusterIdentifier: "docdb-cluster-demo",
 *     masterPassword: "barbut8chars",
 *     masterUsername: "foo",
 * });
 * const clusterInstances: aws.docdb.ClusterInstance[] = [];
 * for (let i = 0; i < 2; i++) {
 *     clusterInstances.push(new aws.docdb.ClusterInstance(`cluster_instances-${i}`, {
 *         clusterIdentifier: defaultCluster.id,
 *         identifier: `docdb-cluster-demo-${i}`,
 *         instanceClass: "db.r4.large",
 *     }));
 * }
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/docdb_cluster_instance.html.markdown.
 */
export class ClusterInstance extends pulumi.CustomResource {
    /**
     * Get an existing ClusterInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterInstanceState, opts?: pulumi.CustomResourceOptions): ClusterInstance {
        return new ClusterInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:docdb/clusterInstance:ClusterInstance';

    /**
     * Returns true if the given object is an instance of ClusterInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterInstance.__pulumiType;
    }

    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    public readonly applyImmediately!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of cluster instance
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     */
    public readonly caCertIdentifier!: pulumi.Output<string>;
    /**
     * The identifier of the [`aws.docdb.Cluster`](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html) in which to launch this instance.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * The DB subnet group to associate with this DB instance.
     */
    public /*out*/ readonly dbSubnetGroupName!: pulumi.Output<string>;
    /**
     * The region-unique, immutable identifier for the DB instance.
     */
    public /*out*/ readonly dbiResourceId!: pulumi.Output<string>;
    /**
     * The DNS address for this instance. May not be writable
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * The database engine version
     */
    public /*out*/ readonly engineVersion!: pulumi.Output<string>;
    /**
     * The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
     */
    public readonly identifierPrefix!: pulumi.Output<string>;
    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances][2]. DocDB currently
     * supports the below instance classes. Please see [AWS Documentation][4] for complete details.
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the cluster.
     */
    public /*out*/ readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * The database port
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled.
     */
    public /*out*/ readonly preferredBackupWindow!: pulumi.Output<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    public readonly promotionTier!: pulumi.Output<number | undefined>;
    public /*out*/ readonly publiclyAccessible!: pulumi.Output<boolean>;
    /**
     * Specifies whether the DB cluster is encrypted.
     */
    public /*out*/ readonly storageEncrypted!: pulumi.Output<boolean>;
    /**
     * A mapping of tags to assign to the instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     */
    public /*out*/ readonly writer!: pulumi.Output<boolean>;

    /**
     * Create a ClusterInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterInstanceArgs | ClusterInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterInstanceState | undefined;
            inputs["applyImmediately"] = state ? state.applyImmediately : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoMinorVersionUpgrade"] = state ? state.autoMinorVersionUpgrade : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["caCertIdentifier"] = state ? state.caCertIdentifier : undefined;
            inputs["clusterIdentifier"] = state ? state.clusterIdentifier : undefined;
            inputs["dbSubnetGroupName"] = state ? state.dbSubnetGroupName : undefined;
            inputs["dbiResourceId"] = state ? state.dbiResourceId : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["identifierPrefix"] = state ? state.identifierPrefix : undefined;
            inputs["instanceClass"] = state ? state.instanceClass : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["preferredBackupWindow"] = state ? state.preferredBackupWindow : undefined;
            inputs["preferredMaintenanceWindow"] = state ? state.preferredMaintenanceWindow : undefined;
            inputs["promotionTier"] = state ? state.promotionTier : undefined;
            inputs["publiclyAccessible"] = state ? state.publiclyAccessible : undefined;
            inputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["writer"] = state ? state.writer : undefined;
        } else {
            const args = argsOrState as ClusterInstanceArgs | undefined;
            if (!args || args.clusterIdentifier === undefined) {
                throw new Error("Missing required property 'clusterIdentifier'");
            }
            if (!args || args.instanceClass === undefined) {
                throw new Error("Missing required property 'instanceClass'");
            }
            inputs["applyImmediately"] = args ? args.applyImmediately : undefined;
            inputs["autoMinorVersionUpgrade"] = args ? args.autoMinorVersionUpgrade : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["caCertIdentifier"] = args ? args.caCertIdentifier : undefined;
            inputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["identifierPrefix"] = args ? args.identifierPrefix : undefined;
            inputs["instanceClass"] = args ? args.instanceClass : undefined;
            inputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            inputs["promotionTier"] = args ? args.promotionTier : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dbSubnetGroupName"] = undefined /*out*/;
            inputs["dbiResourceId"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["engineVersion"] = undefined /*out*/;
            inputs["kmsKeyId"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["preferredBackupWindow"] = undefined /*out*/;
            inputs["publiclyAccessible"] = undefined /*out*/;
            inputs["storageEncrypted"] = undefined /*out*/;
            inputs["writer"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterInstance resources.
 */
export interface ClusterInstanceState {
    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of cluster instance
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
     */
    readonly autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     */
    readonly caCertIdentifier?: pulumi.Input<string>;
    /**
     * The identifier of the [`aws.docdb.Cluster`](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html) in which to launch this instance.
     */
    readonly clusterIdentifier?: pulumi.Input<string>;
    /**
     * The DB subnet group to associate with this DB instance.
     */
    readonly dbSubnetGroupName?: pulumi.Input<string>;
    /**
     * The region-unique, immutable identifier for the DB instance.
     */
    readonly dbiResourceId?: pulumi.Input<string>;
    /**
     * The DNS address for this instance. May not be writable
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * The database engine version
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
     */
    readonly identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances][2]. DocDB currently
     * supports the below instance classes. Please see [AWS Documentation][4] for complete details.
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     */
    readonly instanceClass?: pulumi.Input<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the cluster.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * The database port
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled.
     */
    readonly preferredBackupWindow?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    readonly preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    readonly promotionTier?: pulumi.Input<number>;
    readonly publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * Specifies whether the DB cluster is encrypted.
     */
    readonly storageEncrypted?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the instance.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     */
    readonly writer?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ClusterInstance resource.
 */
export interface ClusterInstanceArgs {
    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    readonly applyImmediately?: pulumi.Input<boolean>;
    /**
     * Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
     */
    readonly autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     */
    readonly caCertIdentifier?: pulumi.Input<string>;
    /**
     * The identifier of the [`aws.docdb.Cluster`](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html) in which to launch this instance.
     */
    readonly clusterIdentifier: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
     */
    readonly identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances][2]. DocDB currently
     * supports the below instance classes. Please see [AWS Documentation][4] for complete details.
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     */
    readonly instanceClass: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    readonly preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    readonly promotionTier?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the instance.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
