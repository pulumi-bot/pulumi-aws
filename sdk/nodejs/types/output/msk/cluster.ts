// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ClusterBrokerNodeGroupInfo {
    /**
     * The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
     */
    azDistribution?: string;
    /**
     * A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
     */
    clientSubnets: string[];
    /**
     * The size in GiB of the EBS volume for the data drive on each broker node.
     */
    ebsVolumeSize: number;
    /**
     * Specify the instance type to use for the kafka brokers. e.g. kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
     */
    instanceType: string;
    /**
     * A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
     */
    securityGroups: string[];
}

export interface ClusterClientAuthentication {
    /**
     * Configuration block for specifying TLS client authentication. See below.
     */
    tls?: outputApi.msk.ClusterClientAuthenticationTls;
}

export interface ClusterClientAuthenticationTls {
    /**
     * List of ACM Certificate Authority Amazon Resource Names (ARNs).
     */
    certificateAuthorityArns?: string[];
}

export interface ClusterConfigurationInfo {
    /**
     * Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
     */
    arn: string;
    /**
     * Revision of the MSK Configuration to use in the cluster.
     */
    revision: number;
}

export interface ClusterEncryptionInfo {
    /**
     * You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
     */
    encryptionAtRestKmsKeyArn: string;
    /**
     * Configuration block to specify encryption in transit. See below.
     */
    encryptionInTransit?: outputApi.msk.ClusterEncryptionInfoEncryptionInTransit;
}

export interface ClusterEncryptionInfoEncryptionInTransit {
    /**
     * Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value: `TLS_PLAINTEXT`.
     */
    clientBroker?: string;
    /**
     * Whether data communication among broker nodes is encrypted. Default value: `true`.
     */
    inCluster?: boolean;
}
