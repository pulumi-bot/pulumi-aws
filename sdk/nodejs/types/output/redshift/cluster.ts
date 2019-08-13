// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ClusterLogging {
    /**
     * The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
     * For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
     */
    bucketName: string;
    /**
     * Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
     */
    enable: boolean;
    /**
     * The prefix applied to the log file names.
     */
    s3KeyPrefix: string;
}

export interface ClusterSnapshotCopy {
    /**
     * The destination region that you want to copy snapshots to.
     */
    destinationRegion: string;
    /**
     * The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
     */
    grantName?: string;
    /**
     * The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
     */
    retentionPeriod?: number;
}
