// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

import {Tags} from "../index";

/**
 * Provides an Elastic MapReduce Cluster, a web service that makes it easy to
 * process large amounts of data efficiently. See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/elastic-mapreduce/)
 * for more information.
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
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState): Cluster {
        return new Cluster(name, <any>state, { id });
    }

    /**
     * A JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore Terraform cannot detect drift from the actual EMR cluster if its value is changed outside Terraform.
     */
    public readonly additionalInfo: pulumi.Output<string | undefined>;
    /**
     * A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, and `Spark`. Case insensitive
     */
    public readonly applications: pulumi.Output<string[] | undefined>;
    /**
     * An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
     */
    public readonly autoscalingRole: pulumi.Output<string | undefined>;
    /**
     * List of bootstrap actions that will be run before Hadoop is started on
     * the cluster nodes. Defined below
     */
    public readonly bootstrapActions: pulumi.Output<{ args?: string[], name: string, path: string }[] | undefined>;
    public /*out*/ readonly clusterState: pulumi.Output<string>;
    /**
     * List of configurations supplied for the EMR cluster you are creating
     */
    public readonly configurations: pulumi.Output<string | undefined>;
    /**
     * A JSON string for supplying list of configurations for the EMR cluster.
     */
    public readonly configurationsJson: pulumi.Output<string | undefined>;
    /**
     * Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `instance_groups` is set. Default `1`
     */
    public readonly coreInstanceCount: pulumi.Output<number | undefined>;
    /**
     * The EC2 instance type of the slave nodes. Cannot be specified if `instance_groups` is set
     */
    public readonly coreInstanceType: pulumi.Output<string>;
    /**
     * A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
     */
    public readonly customAmiId: pulumi.Output<string | undefined>;
    /**
     * Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
     */
    public readonly ebsRootVolumeSize: pulumi.Output<number | undefined>;
    /**
     * Attributes for the EC2 instances running the job
     * flow. Defined below
     */
    public readonly ec2Attributes: pulumi.Output<{ additionalMasterSecurityGroups?: string, additionalSlaveSecurityGroups?: string, emrManagedMasterSecurityGroup?: string, emrManagedSlaveSecurityGroup?: string, instanceProfile: string, keyName?: string, serviceAccessSecurityGroup?: string, subnetId?: string } | undefined>;
    /**
     * A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Defined below
     */
    public readonly instanceGroups: pulumi.Output<{ autoscalingPolicy?: string, bidPrice?: string, ebsConfigs?: { iops?: number, size: number, type: string, volumesPerInstance?: number }[], instanceCount?: number, instanceRole: string, instanceType: string, name?: string }[] | undefined>;
    /**
     * Switch on/off run cluster with no steps or when all steps are complete (default is on)
     */
    public readonly keepJobFlowAliveWhenNoSteps: pulumi.Output<boolean>;
    /**
     * Kerberos configuration for the cluster. Defined below
     */
    public readonly kerberosAttributes: pulumi.Output<{ adDomainJoinPassword?: string, adDomainJoinUser?: string, crossRealmTrustPrincipalPassword?: string, kdcAdminPassword: string, realm: string } | undefined>;
    /**
     * S3 bucket to write the log files of the job flow. If a value
     * is not provided, logs are not created
     */
    public readonly logUri: pulumi.Output<string | undefined>;
    /**
     * The EC2 instance type of the master node. Exactly one of `master_instance_type` and `instance_group` must be specified.
     */
    public readonly masterInstanceType: pulumi.Output<string | undefined>;
    /**
     * The public DNS name of the master EC2 instance.
     */
    public /*out*/ readonly masterPublicDns: pulumi.Output<string>;
    /**
     * The name of the job flow
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The release label for the Amazon EMR release
     */
    public readonly releaseLabel: pulumi.Output<string>;
    /**
     * The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized. 
     */
    public readonly scaleDownBehavior: pulumi.Output<string>;
    /**
     * The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
     */
    public readonly securityConfiguration: pulumi.Output<string | undefined>;
    /**
     * IAM role that will be assumed by the Amazon EMR service to access AWS resources
     */
    public readonly serviceRole: pulumi.Output<string>;
    /**
     * List of steps to run when creating the cluster. Defined below. It is highly recommended to utilize the [lifecycle configuration block](https://www.terraform.io/docs/configuration/resources.html) with `ignore_changes` if other steps are being managed outside of Terraform.
     */
    public readonly steps: pulumi.Output<{ actionOnFailure: string, hadoopJarStep: { args?: string[], jar: string, mainClass?: string, properties?: {[key: string]: any} }, name: string }[]>;
    /**
     * list of tags to apply to the EMR Cluster
     */
    public readonly tags: pulumi.Output<Tags | undefined>;
    /**
     * Switch on/off termination protection (default is off)
     */
    public readonly terminationProtection: pulumi.Output<boolean>;
    /**
     * Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
     */
    public readonly visibleToAllUsers: pulumi.Output<boolean | undefined>;

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
            inputs["additionalInfo"] = state ? state.additionalInfo : undefined;
            inputs["applications"] = state ? state.applications : undefined;
            inputs["autoscalingRole"] = state ? state.autoscalingRole : undefined;
            inputs["bootstrapActions"] = state ? state.bootstrapActions : undefined;
            inputs["clusterState"] = state ? state.clusterState : undefined;
            inputs["configurations"] = state ? state.configurations : undefined;
            inputs["configurationsJson"] = state ? state.configurationsJson : undefined;
            inputs["coreInstanceCount"] = state ? state.coreInstanceCount : undefined;
            inputs["coreInstanceType"] = state ? state.coreInstanceType : undefined;
            inputs["customAmiId"] = state ? state.customAmiId : undefined;
            inputs["ebsRootVolumeSize"] = state ? state.ebsRootVolumeSize : undefined;
            inputs["ec2Attributes"] = state ? state.ec2Attributes : undefined;
            inputs["instanceGroups"] = state ? state.instanceGroups : undefined;
            inputs["keepJobFlowAliveWhenNoSteps"] = state ? state.keepJobFlowAliveWhenNoSteps : undefined;
            inputs["kerberosAttributes"] = state ? state.kerberosAttributes : undefined;
            inputs["logUri"] = state ? state.logUri : undefined;
            inputs["masterInstanceType"] = state ? state.masterInstanceType : undefined;
            inputs["masterPublicDns"] = state ? state.masterPublicDns : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["releaseLabel"] = state ? state.releaseLabel : undefined;
            inputs["scaleDownBehavior"] = state ? state.scaleDownBehavior : undefined;
            inputs["securityConfiguration"] = state ? state.securityConfiguration : undefined;
            inputs["serviceRole"] = state ? state.serviceRole : undefined;
            inputs["steps"] = state ? state.steps : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["terminationProtection"] = state ? state.terminationProtection : undefined;
            inputs["visibleToAllUsers"] = state ? state.visibleToAllUsers : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.releaseLabel === undefined) {
                throw new Error("Missing required property 'releaseLabel'");
            }
            if (!args || args.serviceRole === undefined) {
                throw new Error("Missing required property 'serviceRole'");
            }
            inputs["additionalInfo"] = args ? args.additionalInfo : undefined;
            inputs["applications"] = args ? args.applications : undefined;
            inputs["autoscalingRole"] = args ? args.autoscalingRole : undefined;
            inputs["bootstrapActions"] = args ? args.bootstrapActions : undefined;
            inputs["configurations"] = args ? args.configurations : undefined;
            inputs["configurationsJson"] = args ? args.configurationsJson : undefined;
            inputs["coreInstanceCount"] = args ? args.coreInstanceCount : undefined;
            inputs["coreInstanceType"] = args ? args.coreInstanceType : undefined;
            inputs["customAmiId"] = args ? args.customAmiId : undefined;
            inputs["ebsRootVolumeSize"] = args ? args.ebsRootVolumeSize : undefined;
            inputs["ec2Attributes"] = args ? args.ec2Attributes : undefined;
            inputs["instanceGroups"] = args ? args.instanceGroups : undefined;
            inputs["keepJobFlowAliveWhenNoSteps"] = args ? args.keepJobFlowAliveWhenNoSteps : undefined;
            inputs["kerberosAttributes"] = args ? args.kerberosAttributes : undefined;
            inputs["logUri"] = args ? args.logUri : undefined;
            inputs["masterInstanceType"] = args ? args.masterInstanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["releaseLabel"] = args ? args.releaseLabel : undefined;
            inputs["scaleDownBehavior"] = args ? args.scaleDownBehavior : undefined;
            inputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            inputs["serviceRole"] = args ? args.serviceRole : undefined;
            inputs["steps"] = args ? args.steps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["terminationProtection"] = args ? args.terminationProtection : undefined;
            inputs["visibleToAllUsers"] = args ? args.visibleToAllUsers : undefined;
            inputs["clusterState"] = undefined /*out*/;
            inputs["masterPublicDns"] = undefined /*out*/;
        }
        super("aws:emr/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * A JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore Terraform cannot detect drift from the actual EMR cluster if its value is changed outside Terraform.
     */
    readonly additionalInfo?: pulumi.Input<string>;
    /**
     * A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, and `Spark`. Case insensitive
     */
    readonly applications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
     */
    readonly autoscalingRole?: pulumi.Input<string>;
    /**
     * List of bootstrap actions that will be run before Hadoop is started on
     * the cluster nodes. Defined below
     */
    readonly bootstrapActions?: pulumi.Input<pulumi.Input<{ args?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, path: pulumi.Input<string> }>[]>;
    readonly clusterState?: pulumi.Input<string>;
    /**
     * List of configurations supplied for the EMR cluster you are creating
     */
    readonly configurations?: pulumi.Input<string>;
    /**
     * A JSON string for supplying list of configurations for the EMR cluster.
     */
    readonly configurationsJson?: pulumi.Input<string>;
    /**
     * Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `instance_groups` is set. Default `1`
     */
    readonly coreInstanceCount?: pulumi.Input<number>;
    /**
     * The EC2 instance type of the slave nodes. Cannot be specified if `instance_groups` is set
     */
    readonly coreInstanceType?: pulumi.Input<string>;
    /**
     * A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
     */
    readonly customAmiId?: pulumi.Input<string>;
    /**
     * Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
     */
    readonly ebsRootVolumeSize?: pulumi.Input<number>;
    /**
     * Attributes for the EC2 instances running the job
     * flow. Defined below
     */
    readonly ec2Attributes?: pulumi.Input<{ additionalMasterSecurityGroups?: pulumi.Input<string>, additionalSlaveSecurityGroups?: pulumi.Input<string>, emrManagedMasterSecurityGroup?: pulumi.Input<string>, emrManagedSlaveSecurityGroup?: pulumi.Input<string>, instanceProfile: pulumi.Input<string>, keyName?: pulumi.Input<string>, serviceAccessSecurityGroup?: pulumi.Input<string>, subnetId?: pulumi.Input<string> }>;
    /**
     * A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Defined below
     */
    readonly instanceGroups?: pulumi.Input<pulumi.Input<{ autoscalingPolicy?: pulumi.Input<string>, bidPrice?: pulumi.Input<string>, ebsConfigs?: pulumi.Input<pulumi.Input<{ iops?: pulumi.Input<number>, size: pulumi.Input<number>, type: pulumi.Input<string>, volumesPerInstance?: pulumi.Input<number> }>[]>, instanceCount?: pulumi.Input<number>, instanceRole: pulumi.Input<string>, instanceType: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    /**
     * Switch on/off run cluster with no steps or when all steps are complete (default is on)
     */
    readonly keepJobFlowAliveWhenNoSteps?: pulumi.Input<boolean>;
    /**
     * Kerberos configuration for the cluster. Defined below
     */
    readonly kerberosAttributes?: pulumi.Input<{ adDomainJoinPassword?: pulumi.Input<string>, adDomainJoinUser?: pulumi.Input<string>, crossRealmTrustPrincipalPassword?: pulumi.Input<string>, kdcAdminPassword: pulumi.Input<string>, realm: pulumi.Input<string> }>;
    /**
     * S3 bucket to write the log files of the job flow. If a value
     * is not provided, logs are not created
     */
    readonly logUri?: pulumi.Input<string>;
    /**
     * The EC2 instance type of the master node. Exactly one of `master_instance_type` and `instance_group` must be specified.
     */
    readonly masterInstanceType?: pulumi.Input<string>;
    /**
     * The public DNS name of the master EC2 instance.
     */
    readonly masterPublicDns?: pulumi.Input<string>;
    /**
     * The name of the job flow
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The release label for the Amazon EMR release
     */
    readonly releaseLabel?: pulumi.Input<string>;
    /**
     * The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized. 
     */
    readonly scaleDownBehavior?: pulumi.Input<string>;
    /**
     * The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * IAM role that will be assumed by the Amazon EMR service to access AWS resources
     */
    readonly serviceRole?: pulumi.Input<string>;
    /**
     * List of steps to run when creating the cluster. Defined below. It is highly recommended to utilize the [lifecycle configuration block](https://www.terraform.io/docs/configuration/resources.html) with `ignore_changes` if other steps are being managed outside of Terraform.
     */
    readonly steps?: pulumi.Input<pulumi.Input<{ actionOnFailure: pulumi.Input<string>, hadoopJarStep: pulumi.Input<{ args?: pulumi.Input<pulumi.Input<string>[]>, jar: pulumi.Input<string>, mainClass?: pulumi.Input<string>, properties?: pulumi.Input<{[key: string]: any}> }>, name: pulumi.Input<string> }>[]>;
    /**
     * list of tags to apply to the EMR Cluster
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * Switch on/off termination protection (default is off)
     */
    readonly terminationProtection?: pulumi.Input<boolean>;
    /**
     * Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
     */
    readonly visibleToAllUsers?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * A JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore Terraform cannot detect drift from the actual EMR cluster if its value is changed outside Terraform.
     */
    readonly additionalInfo?: pulumi.Input<string>;
    /**
     * A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, and `Spark`. Case insensitive
     */
    readonly applications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
     */
    readonly autoscalingRole?: pulumi.Input<string>;
    /**
     * List of bootstrap actions that will be run before Hadoop is started on
     * the cluster nodes. Defined below
     */
    readonly bootstrapActions?: pulumi.Input<pulumi.Input<{ args?: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string>, path: pulumi.Input<string> }>[]>;
    /**
     * List of configurations supplied for the EMR cluster you are creating
     */
    readonly configurations?: pulumi.Input<string>;
    /**
     * A JSON string for supplying list of configurations for the EMR cluster.
     */
    readonly configurationsJson?: pulumi.Input<string>;
    /**
     * Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `instance_groups` is set. Default `1`
     */
    readonly coreInstanceCount?: pulumi.Input<number>;
    /**
     * The EC2 instance type of the slave nodes. Cannot be specified if `instance_groups` is set
     */
    readonly coreInstanceType?: pulumi.Input<string>;
    /**
     * A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
     */
    readonly customAmiId?: pulumi.Input<string>;
    /**
     * Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
     */
    readonly ebsRootVolumeSize?: pulumi.Input<number>;
    /**
     * Attributes for the EC2 instances running the job
     * flow. Defined below
     */
    readonly ec2Attributes?: pulumi.Input<{ additionalMasterSecurityGroups?: pulumi.Input<string>, additionalSlaveSecurityGroups?: pulumi.Input<string>, emrManagedMasterSecurityGroup?: pulumi.Input<string>, emrManagedSlaveSecurityGroup?: pulumi.Input<string>, instanceProfile: pulumi.Input<string>, keyName?: pulumi.Input<string>, serviceAccessSecurityGroup?: pulumi.Input<string>, subnetId?: pulumi.Input<string> }>;
    /**
     * A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Defined below
     */
    readonly instanceGroups?: pulumi.Input<pulumi.Input<{ autoscalingPolicy?: pulumi.Input<string>, bidPrice?: pulumi.Input<string>, ebsConfigs?: pulumi.Input<pulumi.Input<{ iops?: pulumi.Input<number>, size: pulumi.Input<number>, type: pulumi.Input<string>, volumesPerInstance?: pulumi.Input<number> }>[]>, instanceCount?: pulumi.Input<number>, instanceRole: pulumi.Input<string>, instanceType: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    /**
     * Switch on/off run cluster with no steps or when all steps are complete (default is on)
     */
    readonly keepJobFlowAliveWhenNoSteps?: pulumi.Input<boolean>;
    /**
     * Kerberos configuration for the cluster. Defined below
     */
    readonly kerberosAttributes?: pulumi.Input<{ adDomainJoinPassword?: pulumi.Input<string>, adDomainJoinUser?: pulumi.Input<string>, crossRealmTrustPrincipalPassword?: pulumi.Input<string>, kdcAdminPassword: pulumi.Input<string>, realm: pulumi.Input<string> }>;
    /**
     * S3 bucket to write the log files of the job flow. If a value
     * is not provided, logs are not created
     */
    readonly logUri?: pulumi.Input<string>;
    /**
     * The EC2 instance type of the master node. Exactly one of `master_instance_type` and `instance_group` must be specified.
     */
    readonly masterInstanceType?: pulumi.Input<string>;
    /**
     * The name of the job flow
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The release label for the Amazon EMR release
     */
    readonly releaseLabel: pulumi.Input<string>;
    /**
     * The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized. 
     */
    readonly scaleDownBehavior?: pulumi.Input<string>;
    /**
     * The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * IAM role that will be assumed by the Amazon EMR service to access AWS resources
     */
    readonly serviceRole: pulumi.Input<string>;
    /**
     * List of steps to run when creating the cluster. Defined below. It is highly recommended to utilize the [lifecycle configuration block](https://www.terraform.io/docs/configuration/resources.html) with `ignore_changes` if other steps are being managed outside of Terraform.
     */
    readonly steps?: pulumi.Input<pulumi.Input<{ actionOnFailure: pulumi.Input<string>, hadoopJarStep: pulumi.Input<{ args?: pulumi.Input<pulumi.Input<string>[]>, jar: pulumi.Input<string>, mainClass?: pulumi.Input<string>, properties?: pulumi.Input<{[key: string]: any}> }>, name: pulumi.Input<string> }>[]>;
    /**
     * list of tags to apply to the EMR Cluster
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * Switch on/off termination protection (default is off)
     */
    readonly terminationProtection?: pulumi.Input<boolean>;
    /**
     * Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
     */
    readonly visibleToAllUsers?: pulumi.Input<boolean>;
}
