// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Retrieve information about an EKS Cluster.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = aws.eks.getCluster({
 *     name: "example",
 * });
 * 
 * export const endpoint = example.endpoint;
 * export const kubeconfigCertificateAuthorityData = example.certificateAuthority.data;
 * // Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
 * export const identityOidcIssuer = example.identities[0].oidcs[0].issuer;
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/eks_cluster.html.markdown.
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> & GetClusterResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClusterResult> = pulumi.runtime.invoke("aws:eks/getCluster:getCluster", {
        "name": args.name,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The name of the cluster
     */
    readonly name: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    readonly arn: string;
    /**
     * Nested attribute containing `certificate-authority-data` for your cluster.
     */
    readonly certificateAuthority: outputs.eks.GetClusterCertificateAuthority;
    /**
     * The Unix epoch time stamp in seconds for when the cluster was created.
     */
    readonly createdAt: string;
    /**
     * The enabled control plane logs.
     */
    readonly enabledClusterLogTypes: string[];
    /**
     * The endpoint for your Kubernetes API server.
     */
    readonly endpoint: string;
    /**
     * Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. For an example using this information to enable IAM Roles for Service Accounts, see the [`aws.eks.Cluster` resource documentation](https://www.terraform.io/docs/providers/aws/r/eks_cluster.html).
     */
    readonly identities: outputs.eks.GetClusterIdentity[];
    readonly name: string;
    /**
     * The platform version for the cluster.
     */
    readonly platformVersion: string;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    readonly roleArn: string;
    /**
     * The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     */
    readonly status: string;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags: {[key: string]: any};
    /**
     * The Kubernetes server version for the cluster.
     */
    readonly version: string;
    /**
     * Nested list containing VPC configuration for the cluster.
     */
    readonly vpcConfig: outputs.eks.GetClusterVpcConfig;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
