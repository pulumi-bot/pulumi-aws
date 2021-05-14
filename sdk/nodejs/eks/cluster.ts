// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an EKS Cluster.
 *
 * ## Example Usage
 * ### Example IAM Role for EKS Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.iam.Role("example", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "eks.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `});
 * const example_AmazonEKSClusterPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
 *     role: example.name,
 * });
 * // Optionally, enable Security Groups for Pods
 * // Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
 * const example_AmazonEKSVPCResourceController = new aws.iam.RolePolicyAttachment("example-AmazonEKSVPCResourceController", {
 *     policyArn: "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController",
 *     role: example.name,
 * });
 * ```
 * ### Enabling Control Plane Logging
 *
 * [EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabledClusterLogTypes` argument. To manage the CloudWatch Log Group retention period, the `aws.cloudwatch.LogGroup` resource can be used.
 *
 * > The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const clusterName = config.get("clusterName") || "example";
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 7});
 * // ... potentially other configuration ...
 * const exampleCluster = new aws.eks.Cluster("exampleCluster", {enabledClusterLogTypes: [
 *     "api",
 *     "audit",
 * ]}, {
 *     dependsOn: [exampleLogGroup],
 * });
 * // ... other configuration ...
 * ```
 *
 * ## Import
 *
 * EKS Clusters can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:eks/cluster:Cluster my_cluster my_cluster
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:eks/cluster:Cluster';

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

    /**
     * ARN of the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
     */
    public /*out*/ readonly certificateAuthority!: pulumi.Output<outputs.eks.ClusterCertificateAuthority>;
    /**
     * Unix epoch timestamp in seconds for when the cluster was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     */
    public readonly enabledClusterLogTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.eks.ClusterEncryptionConfig | undefined>;
    /**
     * Endpoint for your Kubernetes API server.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
     */
    public /*out*/ readonly identities!: pulumi.Output<outputs.eks.ClusterIdentity[]>;
    /**
     * Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     */
    public readonly kubernetesNetworkConfig!: pulumi.Output<outputs.eks.ClusterKubernetesNetworkConfig>;
    /**
     * Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Platform version for the cluster.
     */
    public /*out*/ readonly platformVersion!: pulumi.Output<string>;
    /**
     * ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider.
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     */
    public readonly vpcConfig!: pulumi.Output<outputs.eks.ClusterVpcConfig>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificateAuthority"] = state ? state.certificateAuthority : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["enabledClusterLogTypes"] = state ? state.enabledClusterLogTypes : undefined;
            inputs["encryptionConfig"] = state ? state.encryptionConfig : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["identities"] = state ? state.identities : undefined;
            inputs["kubernetesNetworkConfig"] = state ? state.kubernetesNetworkConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformVersion"] = state ? state.platformVersion : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.vpcConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcConfig'");
            }
            inputs["enabledClusterLogTypes"] = args ? args.enabledClusterLogTypes : undefined;
            inputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            inputs["kubernetesNetworkConfig"] = args ? args.kubernetesNetworkConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificateAuthority"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["identities"] = undefined /*out*/;
            inputs["platformVersion"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * ARN of the cluster.
     */
    arn?: pulumi.Input<string>;
    /**
     * Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
     */
    certificateAuthority?: pulumi.Input<inputs.eks.ClusterCertificateAuthority>;
    /**
     * Unix epoch timestamp in seconds for when the cluster was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     */
    enabledClusterLogTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     */
    encryptionConfig?: pulumi.Input<inputs.eks.ClusterEncryptionConfig>;
    /**
     * Endpoint for your Kubernetes API server.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
     */
    identities?: pulumi.Input<pulumi.Input<inputs.eks.ClusterIdentity>[]>;
    /**
     * Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     */
    kubernetesNetworkConfig?: pulumi.Input<inputs.eks.ClusterKubernetesNetworkConfig>;
    /**
     * Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     */
    name?: pulumi.Input<string>;
    /**
     * Platform version for the cluster.
     */
    platformVersion?: pulumi.Input<string>;
    /**
     * ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     */
    status?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    version?: pulumi.Input<string>;
    /**
     * Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     */
    vpcConfig?: pulumi.Input<inputs.eks.ClusterVpcConfig>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     */
    enabledClusterLogTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     */
    encryptionConfig?: pulumi.Input<inputs.eks.ClusterEncryptionConfig>;
    /**
     * Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     */
    kubernetesNetworkConfig?: pulumi.Input<inputs.eks.ClusterKubernetesNetworkConfig>;
    /**
     * Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     */
    name?: pulumi.Input<string>;
    /**
     * ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    roleArn: pulumi.Input<string>;
    /**
     * Key-value map of resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    version?: pulumi.Input<string>;
    /**
     * Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     */
    vpcConfig: pulumi.Input<inputs.eks.ClusterVpcConfig>;
}
