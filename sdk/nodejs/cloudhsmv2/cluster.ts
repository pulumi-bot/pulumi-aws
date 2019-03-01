// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Amazon CloudHSM v2 cluster.
 * 
 * For information about CloudHSM v2, see the
 * [AWS CloudHSM User Guide][1] and the [Amazon
 * CloudHSM API Reference][2].
 * 
 * > **NOTE:** CloudHSM can take up to several minutes to be set up.
 * Practically no single attribute can be updated except TAGS.
 * If you need to delete a cluster, you have to remove its HSM modules first.
 * To initialize cluster you have to sign CSR and upload it.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /**
     * The list of cluster certificates.
     * * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
     * * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state.
     * * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
     * * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
     * * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
     */
    public /*out*/ readonly clusterCertificates: pulumi.Output<{ awsHardwareCertificate: string, clusterCertificate: string, clusterCsr: string, hsmCertificate: string, manufacturerHardwareCertificate: string }>;
    /**
     * The id of the CloudHSM cluster.
     */
    public /*out*/ readonly clusterId: pulumi.Output<string>;
    /**
     * The state of the cluster.
     */
    public /*out*/ readonly clusterState: pulumi.Output<string>;
    /**
     * The type of HSM module in the cluster. Currently, only hsm1.medium is supported.
     */
    public readonly hsmType: pulumi.Output<string>;
    /**
     * The ID of the security group associated with the CloudHSM cluster.
     */
    public /*out*/ readonly securityGroupId: pulumi.Output<string>;
    /**
     * The id of Cloud HSM v2 cluster backup to be restored.
     */
    public readonly sourceBackupIdentifier: pulumi.Output<string | undefined>;
    /**
     * The IDs of subnets in which cluster will operate.
     */
    public readonly subnetIds: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The id of the VPC that the CloudHSM cluster resides in.
     */
    public /*out*/ readonly vpcId: pulumi.Output<string>;

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
        if (opts && opts.id) {
            const state: ClusterState = argsOrState as ClusterState | undefined;
            inputs["clusterCertificates"] = state ? state.clusterCertificates : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["clusterState"] = state ? state.clusterState : undefined;
            inputs["hsmType"] = state ? state.hsmType : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["sourceBackupIdentifier"] = state ? state.sourceBackupIdentifier : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.hsmType === undefined) {
                throw new Error("Missing required property 'hsmType'");
            }
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            inputs["hsmType"] = args ? args.hsmType : undefined;
            inputs["sourceBackupIdentifier"] = args ? args.sourceBackupIdentifier : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["clusterCertificates"] = undefined /*out*/;
            inputs["clusterId"] = undefined /*out*/;
            inputs["clusterState"] = undefined /*out*/;
            inputs["securityGroupId"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        super("aws:cloudhsmv2/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The list of cluster certificates.
     * * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
     * * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in UNINITIALIZED state.
     * * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
     * * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
     * * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
     */
    readonly clusterCertificates?: pulumi.Input<{ awsHardwareCertificate?: pulumi.Input<string>, clusterCertificate?: pulumi.Input<string>, clusterCsr?: pulumi.Input<string>, hsmCertificate?: pulumi.Input<string>, manufacturerHardwareCertificate?: pulumi.Input<string> }>;
    /**
     * The id of the CloudHSM cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * The state of the cluster.
     */
    readonly clusterState?: pulumi.Input<string>;
    /**
     * The type of HSM module in the cluster. Currently, only hsm1.medium is supported.
     */
    readonly hsmType?: pulumi.Input<string>;
    /**
     * The ID of the security group associated with the CloudHSM cluster.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * The id of Cloud HSM v2 cluster backup to be restored.
     */
    readonly sourceBackupIdentifier?: pulumi.Input<string>;
    /**
     * The IDs of subnets in which cluster will operate.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The id of the VPC that the CloudHSM cluster resides in.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The type of HSM module in the cluster. Currently, only hsm1.medium is supported.
     */
    readonly hsmType: pulumi.Input<string>;
    /**
     * The id of Cloud HSM v2 cluster backup to be restored.
     */
    readonly sourceBackupIdentifier?: pulumi.Input<string>;
    /**
     * The IDs of subnets in which cluster will operate.
     */
    readonly subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
