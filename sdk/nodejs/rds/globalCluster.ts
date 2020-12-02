// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
 *
 * More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
 *
 * ## Example Usage
 * ### New Global Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const primary = new aws.Provider("primary", {region: "us-east-2"});
 * const secondary = new aws.Provider("secondary", {region: "us-west-2"});
 * const example = new aws.rds.GlobalCluster("example", {globalClusterIdentifier: "example"}, {
 *     provider: aws.primary,
 * });
 * const primaryCluster = new aws.rds.Cluster("primaryCluster", {globalClusterIdentifier: example.id}, {
 *     provider: aws.primary,
 * });
 * const primaryClusterInstance = new aws.rds.ClusterInstance("primaryClusterInstance", {clusterIdentifier: primaryCluster.id}, {
 *     provider: aws.primary,
 * });
 * const secondaryCluster = new aws.rds.Cluster("secondaryCluster", {globalClusterIdentifier: example.id}, {
 *     provider: aws.secondary,
 *     dependsOn: [primaryClusterInstance],
 * });
 * const secondaryClusterInstance = new aws.rds.ClusterInstance("secondaryClusterInstance", {clusterIdentifier: secondaryCluster.id}, {
 *     provider: aws.secondary,
 * });
 * ```
 * ### New Global Cluster From Existing DB Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ... other configuration ...
 * const exampleCluster = new aws.rds.Cluster("exampleCluster", {});
 * const exampleGlobalCluster = new aws.rds.GlobalCluster("exampleGlobalCluster", {
 *     forceDestroy: true,
 *     globalClusterIdentifier: "example",
 *     sourceDbClusterIdentifier: exampleCluster.arn,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_rds_global_cluster` can be imported by using the RDS Global Cluster identifier, e.g.
 *
 * ```sh
 *  $ pulumi import aws:rds/globalCluster:GlobalCluster example example
 * ```
 *
 *  Certain resource arguments, like `force_destroy`, only exist within Terraform. If the argument is set in the Terraform configuration on an imported resource, Terraform will show a difference on the first plan after import to update the state value. This change is safe to apply immediately so the state matches the desired configuration. Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the Terraform configuration on an imported resource, Terraform will always show a difference. To workaround this behavior, either omit the argument from the Terraform configuration or use [`ignore_changes`](/docs/configuration/resources.html#ignore_changes) to hide the difference, e.g. hcl resource "aws_rds_global_cluster" "example" {
 *
 * # ... other configuration ...
 *
 * # There is no API for reading source_db_cluster_identifier
 *
 *  lifecycle {
 *
 *  ignore_changes = [source_db_cluster_identifier]
 *
 *  } }
 */
export class GlobalCluster extends pulumi.CustomResource {
    /**
     * Get an existing GlobalCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalClusterState, opts?: pulumi.CustomResourceOptions): GlobalCluster {
        return new GlobalCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/globalCluster:GlobalCluster';

    /**
     * Returns true if the given object is an instance of GlobalCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalCluster.__pulumiType;
    }

    /**
     * RDS Global Cluster Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Name for an automatically created database on cluster creation.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    public readonly engine!: pulumi.Output<string>;
    /**
     * Engine version of the Aurora global database.
     * * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The global cluster identifier.
     */
    public readonly globalClusterIdentifier!: pulumi.Output<string>;
    /**
     * Set of objects containing Global Cluster members.
     */
    public /*out*/ readonly globalClusterMembers!: pulumi.Output<outputs.rds.GlobalClusterGlobalClusterMember[]>;
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
     */
    public /*out*/ readonly globalClusterResourceId!: pulumi.Output<string>;
    public readonly sourceDbClusterIdentifier!: pulumi.Output<string>;
    public readonly storageEncrypted!: pulumi.Output<boolean>;

    /**
     * Create a GlobalCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalClusterArgs | GlobalClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GlobalClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["globalClusterIdentifier"] = state ? state.globalClusterIdentifier : undefined;
            inputs["globalClusterMembers"] = state ? state.globalClusterMembers : undefined;
            inputs["globalClusterResourceId"] = state ? state.globalClusterResourceId : undefined;
            inputs["sourceDbClusterIdentifier"] = state ? state.sourceDbClusterIdentifier : undefined;
            inputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
        } else {
            const args = argsOrState as GlobalClusterArgs | undefined;
            if ((!args || args.globalClusterIdentifier === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'globalClusterIdentifier'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["globalClusterIdentifier"] = args ? args.globalClusterIdentifier : undefined;
            inputs["sourceDbClusterIdentifier"] = args ? args.sourceDbClusterIdentifier : undefined;
            inputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["globalClusterMembers"] = undefined /*out*/;
            inputs["globalClusterResourceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GlobalCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalCluster resources.
 */
export interface GlobalClusterState {
    /**
     * RDS Global Cluster Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Name for an automatically created database on cluster creation.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly engine?: pulumi.Input<string>;
    /**
     * Engine version of the Aurora global database.
     * * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The global cluster identifier.
     */
    readonly globalClusterIdentifier?: pulumi.Input<string>;
    /**
     * Set of objects containing Global Cluster members.
     */
    readonly globalClusterMembers?: pulumi.Input<pulumi.Input<inputs.rds.GlobalClusterGlobalClusterMember>[]>;
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
     */
    readonly globalClusterResourceId?: pulumi.Input<string>;
    readonly sourceDbClusterIdentifier?: pulumi.Input<string>;
    readonly storageEncrypted?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GlobalCluster resource.
 */
export interface GlobalClusterArgs {
    /**
     * Name for an automatically created database on cluster creation.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly engine?: pulumi.Input<string>;
    /**
     * Engine version of the Aurora global database.
     * * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The global cluster identifier.
     */
    readonly globalClusterIdentifier: pulumi.Input<string>;
    readonly sourceDbClusterIdentifier?: pulumi.Input<string>;
    readonly storageEncrypted?: pulumi.Input<boolean>;
}
