// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about a specific redshift cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testCluster = aws.redshift.getCluster({
 *     clusterIdentifier: "test-cluster",
 * });
 * const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
 *     destination: "redshift",
 *     s3Configuration: {
 *         roleArn: aws_iam_role.firehose_role.arn,
 *         bucketArn: aws_s3_bucket.bucket.arn,
 *         bufferSize: 10,
 *         bufferInterval: 400,
 *         compressionFormat: "GZIP",
 *     },
 *     redshiftConfiguration: {
 *         roleArn: aws_iam_role.firehose_role.arn,
 *         clusterJdbcurl: Promise.all([testCluster, testCluster]).then(([testCluster, testCluster1]) => `jdbc:redshift://${testCluster.endpoint}/${testCluster1.databaseName}`),
 *         username: "testuser",
 *         password: "T3stPass",
 *         dataTableName: "test-table",
 *         copyOptions: "delimiter '|'",
 *         dataTableColumns: "test-col",
 *     },
 * });
 * ```
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:redshift/getCluster:getCluster", {
        "clusterIdentifier": args.clusterIdentifier,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The cluster identifier
     */
    readonly clusterIdentifier: string;
    /**
     * The tags associated to the cluster
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    /**
     * Whether major version upgrades can be applied during maintenance period
     */
    readonly allowVersionUpgrade: boolean;
    /**
     * The backup retention period
     */
    readonly automatedSnapshotRetentionPeriod: number;
    /**
     * The availability zone of the cluster
     */
    readonly availabilityZone: string;
    /**
     * The name of the S3 bucket where the log files are to be stored
     */
    readonly bucketName: string;
    /**
     * The cluster identifier
     */
    readonly clusterIdentifier: string;
    /**
     * The name of the parameter group to be associated with this cluster
     */
    readonly clusterParameterGroupName: string;
    /**
     * The public key for the cluster
     */
    readonly clusterPublicKey: string;
    /**
     * The cluster revision number
     */
    readonly clusterRevisionNumber: string;
    /**
     * The security groups associated with the cluster
     */
    readonly clusterSecurityGroups: string[];
    /**
     * The name of a cluster subnet group to be associated with this cluster
     */
    readonly clusterSubnetGroupName: string;
    /**
     * The cluster type
     */
    readonly clusterType: string;
    readonly clusterVersion: string;
    /**
     * The name of the default database in the cluster
     */
    readonly databaseName: string;
    /**
     * The Elastic IP of the cluster
     */
    readonly elasticIp: string;
    /**
     * Whether cluster logging is enabled
     */
    readonly enableLogging: boolean;
    /**
     * Whether the cluster data is encrypted
     */
    readonly encrypted: boolean;
    /**
     * The cluster endpoint
     */
    readonly endpoint: string;
    /**
     * Whether enhanced VPC routing is enabled
     */
    readonly enhancedVpcRouting: boolean;
    /**
     * The IAM roles associated to the cluster
     */
    readonly iamRoles: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The KMS encryption key associated to the cluster
     */
    readonly kmsKeyId: string;
    /**
     * Username for the master DB user
     */
    readonly masterUsername: string;
    /**
     * The cluster node type
     */
    readonly nodeType: string;
    /**
     * The number of nodes in the cluster
     */
    readonly numberOfNodes: number;
    /**
     * The port the cluster responds on
     */
    readonly port: number;
    /**
     * The maintenance window
     */
    readonly preferredMaintenanceWindow: string;
    /**
     * Whether the cluster is publicly accessible
     */
    readonly publiclyAccessible: boolean;
    /**
     * The folder inside the S3 bucket where the log files are stored
     */
    readonly s3KeyPrefix: string;
    /**
     * The tags associated to the cluster
     */
    readonly tags?: {[key: string]: string};
    /**
     * The VPC Id associated with the cluster
     */
    readonly vpcId: string;
    /**
     * The VPC security group Ids associated with the cluster
     */
    readonly vpcSecurityGroupIds: string[];
}
