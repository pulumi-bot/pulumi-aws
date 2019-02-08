// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a CloudHSM v2 cluster
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    return pulumi.runtime.invoke("aws:cloudhsmv2/getCluster:getCluster", {
        "clusterId": args.clusterId,
        "clusterState": args.clusterState,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The id of Cloud HSM v2 cluster.
     */
    readonly clusterId: string;
    /**
     * The state of the cluster to be found.
     */
    readonly clusterState?: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    /**
     * The list of cluster certificates.
     * * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
     * * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state.
     * * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
     * * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
     * * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
     * The number of available cluster certificates may vary depending on state of the cluster.
     */
    readonly clusterCertificates: { awsHardwareCertificate: string, clusterCertificate: string, clusterCsr: string, hsmCertificate: string, manufacturerHardwareCertificate: string };
    readonly clusterState: string;
    /**
     * The ID of the security group associated with the CloudHSM cluster.
     */
    readonly securityGroupId: string;
    /**
     * The IDs of subnets in which cluster operates.
     */
    readonly subnetIds: string[];
    /**
     * The id of the VPC that the CloudHSM cluster resides in.
     */
    readonly vpcId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
