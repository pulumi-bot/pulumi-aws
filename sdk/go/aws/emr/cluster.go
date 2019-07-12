// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/emr_cluster.html.markdown.
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	if args == nil || args.ReleaseLabel == nil {
		return nil, errors.New("missing required argument 'ReleaseLabel'")
	}
	if args == nil || args.ServiceRole == nil {
		return nil, errors.New("missing required argument 'ServiceRole'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["additionalInfo"] = nil
		inputs["applications"] = nil
		inputs["autoscalingRole"] = nil
		inputs["bootstrapActions"] = nil
		inputs["configurations"] = nil
		inputs["configurationsJson"] = nil
		inputs["coreInstanceCount"] = nil
		inputs["coreInstanceGroup"] = nil
		inputs["coreInstanceType"] = nil
		inputs["customAmiId"] = nil
		inputs["ebsRootVolumeSize"] = nil
		inputs["ec2Attributes"] = nil
		inputs["instanceGroups"] = nil
		inputs["keepJobFlowAliveWhenNoSteps"] = nil
		inputs["kerberosAttributes"] = nil
		inputs["logUri"] = nil
		inputs["masterInstanceGroup"] = nil
		inputs["masterInstanceType"] = nil
		inputs["name"] = nil
		inputs["releaseLabel"] = nil
		inputs["scaleDownBehavior"] = nil
		inputs["securityConfiguration"] = nil
		inputs["serviceRole"] = nil
		inputs["steps"] = nil
		inputs["tags"] = nil
		inputs["terminationProtection"] = nil
		inputs["visibleToAllUsers"] = nil
	} else {
		inputs["additionalInfo"] = args.AdditionalInfo
		inputs["applications"] = args.Applications
		inputs["autoscalingRole"] = args.AutoscalingRole
		inputs["bootstrapActions"] = args.BootstrapActions
		inputs["configurations"] = args.Configurations
		inputs["configurationsJson"] = args.ConfigurationsJson
		inputs["coreInstanceCount"] = args.CoreInstanceCount
		inputs["coreInstanceGroup"] = args.CoreInstanceGroup
		inputs["coreInstanceType"] = args.CoreInstanceType
		inputs["customAmiId"] = args.CustomAmiId
		inputs["ebsRootVolumeSize"] = args.EbsRootVolumeSize
		inputs["ec2Attributes"] = args.Ec2Attributes
		inputs["instanceGroups"] = args.InstanceGroups
		inputs["keepJobFlowAliveWhenNoSteps"] = args.KeepJobFlowAliveWhenNoSteps
		inputs["kerberosAttributes"] = args.KerberosAttributes
		inputs["logUri"] = args.LogUri
		inputs["masterInstanceGroup"] = args.MasterInstanceGroup
		inputs["masterInstanceType"] = args.MasterInstanceType
		inputs["name"] = args.Name
		inputs["releaseLabel"] = args.ReleaseLabel
		inputs["scaleDownBehavior"] = args.ScaleDownBehavior
		inputs["securityConfiguration"] = args.SecurityConfiguration
		inputs["serviceRole"] = args.ServiceRole
		inputs["steps"] = args.Steps
		inputs["tags"] = args.Tags
		inputs["terminationProtection"] = args.TerminationProtection
		inputs["visibleToAllUsers"] = args.VisibleToAllUsers
	}
	inputs["clusterState"] = nil
	inputs["masterPublicDns"] = nil
	s, err := ctx.RegisterResource("aws:emr/cluster:Cluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["additionalInfo"] = state.AdditionalInfo
		inputs["applications"] = state.Applications
		inputs["autoscalingRole"] = state.AutoscalingRole
		inputs["bootstrapActions"] = state.BootstrapActions
		inputs["clusterState"] = state.ClusterState
		inputs["configurations"] = state.Configurations
		inputs["configurationsJson"] = state.ConfigurationsJson
		inputs["coreInstanceCount"] = state.CoreInstanceCount
		inputs["coreInstanceGroup"] = state.CoreInstanceGroup
		inputs["coreInstanceType"] = state.CoreInstanceType
		inputs["customAmiId"] = state.CustomAmiId
		inputs["ebsRootVolumeSize"] = state.EbsRootVolumeSize
		inputs["ec2Attributes"] = state.Ec2Attributes
		inputs["instanceGroups"] = state.InstanceGroups
		inputs["keepJobFlowAliveWhenNoSteps"] = state.KeepJobFlowAliveWhenNoSteps
		inputs["kerberosAttributes"] = state.KerberosAttributes
		inputs["logUri"] = state.LogUri
		inputs["masterInstanceGroup"] = state.MasterInstanceGroup
		inputs["masterInstanceType"] = state.MasterInstanceType
		inputs["masterPublicDns"] = state.MasterPublicDns
		inputs["name"] = state.Name
		inputs["releaseLabel"] = state.ReleaseLabel
		inputs["scaleDownBehavior"] = state.ScaleDownBehavior
		inputs["securityConfiguration"] = state.SecurityConfiguration
		inputs["serviceRole"] = state.ServiceRole
		inputs["steps"] = state.Steps
		inputs["tags"] = state.Tags
		inputs["terminationProtection"] = state.TerminationProtection
		inputs["visibleToAllUsers"] = state.VisibleToAllUsers
	}
	s, err := ctx.ReadResource("aws:emr/cluster:Cluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Cluster) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Cluster) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Cluster) AdditionalInfo() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["additionalInfo"])
}

// A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, `Spark`, and `JupyterHub` (as of EMR 5.14.0). Case insensitive
func (r *Cluster) Applications() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["applications"])
}

// An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
func (r *Cluster) AutoscalingRole() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["autoscalingRole"])
}

// List of bootstrap actions that will be run before Hadoop is started on the cluster nodes. Defined below
func (r *Cluster) BootstrapActions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["bootstrapActions"])
}

func (r *Cluster) ClusterState() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterState"])
}

// List of configurations supplied for the EMR cluster you are creating
func (r *Cluster) Configurations() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["configurations"])
}

// A JSON string for supplying list of configurations for the EMR cluster.
func (r *Cluster) ConfigurationsJson() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["configurationsJson"])
}

// Use the `core_instance_group` configuration block `instance_count` argument instead. Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set. Default `1`
func (r *Cluster) CoreInstanceCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["coreInstanceCount"])
}

// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core). Cannot be specified if `core_instance_count` argument, `core_instance_type` argument, or `instance_group` configuration blocks are set. Detailed below.
func (r *Cluster) CoreInstanceGroup() *pulumi.Output {
	return r.s.State["coreInstanceGroup"]
}

// Use the `core_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the slave nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set.
func (r *Cluster) CoreInstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["coreInstanceType"])
}

// A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
func (r *Cluster) CustomAmiId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["customAmiId"])
}

// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
func (r *Cluster) EbsRootVolumeSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ebsRootVolumeSize"])
}

// Attributes for the EC2 instances running the job flow. Defined below
func (r *Cluster) Ec2Attributes() *pulumi.Output {
	return r.s.State["ec2Attributes"]
}

// Use the `master_instance_group` configuration block, `core_instance_group` configuration block and [`aws_emr_instance_group` resource(s)](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html) instead. A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Cannot be specified if `master_instance_group` or `core_instance_group` configuration blocks are set. Defined below
func (r *Cluster) InstanceGroups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["instanceGroups"])
}

// Switch on/off run cluster with no steps or when all steps are complete (default is on)
func (r *Cluster) KeepJobFlowAliveWhenNoSteps() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["keepJobFlowAliveWhenNoSteps"])
}

// Kerberos configuration for the cluster. Defined below
func (r *Cluster) KerberosAttributes() *pulumi.Output {
	return r.s.State["kerberosAttributes"]
}

// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created
func (r *Cluster) LogUri() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["logUri"])
}

// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master). Cannot be specified if `master_instance_type` argument or `instance_group` configuration blocks are set. Detailed below.
func (r *Cluster) MasterInstanceGroup() *pulumi.Output {
	return r.s.State["masterInstanceGroup"]
}

// Use the `master_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the master node. Cannot be specified if `master_instance_group` or `instance_group` configuration blocks are set.
func (r *Cluster) MasterInstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterInstanceType"])
}

// The public DNS name of the master EC2 instance.
// * `core_instance_group.0.id` - Core node type Instance Group ID, if using Instance Group for this node type.
func (r *Cluster) MasterPublicDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterPublicDns"])
}

// The name of the job flow
func (r *Cluster) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The release label for the Amazon EMR release
func (r *Cluster) ReleaseLabel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["releaseLabel"])
}

// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
func (r *Cluster) ScaleDownBehavior() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["scaleDownBehavior"])
}

// The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
func (r *Cluster) SecurityConfiguration() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["securityConfiguration"])
}

// IAM role that will be assumed by the Amazon EMR service to access AWS resources
func (r *Cluster) ServiceRole() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceRole"])
}

func (r *Cluster) Steps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["steps"])
}

// list of tags to apply to the EMR Cluster
func (r *Cluster) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Switch on/off termination protection (default is off)
func (r *Cluster) TerminationProtection() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["terminationProtection"])
}

// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
func (r *Cluster) VisibleToAllUsers() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["visibleToAllUsers"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	AdditionalInfo interface{}
	// A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, `Spark`, and `JupyterHub` (as of EMR 5.14.0). Case insensitive
	Applications interface{}
	// An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
	AutoscalingRole interface{}
	// List of bootstrap actions that will be run before Hadoop is started on the cluster nodes. Defined below
	BootstrapActions interface{}
	ClusterState interface{}
	// List of configurations supplied for the EMR cluster you are creating
	Configurations interface{}
	// A JSON string for supplying list of configurations for the EMR cluster.
	ConfigurationsJson interface{}
	// Use the `core_instance_group` configuration block `instance_count` argument instead. Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set. Default `1`
	CoreInstanceCount interface{}
	// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core). Cannot be specified if `core_instance_count` argument, `core_instance_type` argument, or `instance_group` configuration blocks are set. Detailed below.
	CoreInstanceGroup interface{}
	// Use the `core_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the slave nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set.
	CoreInstanceType interface{}
	// A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
	CustomAmiId interface{}
	// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
	EbsRootVolumeSize interface{}
	// Attributes for the EC2 instances running the job flow. Defined below
	Ec2Attributes interface{}
	// Use the `master_instance_group` configuration block, `core_instance_group` configuration block and [`aws_emr_instance_group` resource(s)](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html) instead. A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Cannot be specified if `master_instance_group` or `core_instance_group` configuration blocks are set. Defined below
	InstanceGroups interface{}
	// Switch on/off run cluster with no steps or when all steps are complete (default is on)
	KeepJobFlowAliveWhenNoSteps interface{}
	// Kerberos configuration for the cluster. Defined below
	KerberosAttributes interface{}
	// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created
	LogUri interface{}
	// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master). Cannot be specified if `master_instance_type` argument or `instance_group` configuration blocks are set. Detailed below.
	MasterInstanceGroup interface{}
	// Use the `master_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the master node. Cannot be specified if `master_instance_group` or `instance_group` configuration blocks are set.
	MasterInstanceType interface{}
	// The public DNS name of the master EC2 instance.
	// * `core_instance_group.0.id` - Core node type Instance Group ID, if using Instance Group for this node type.
	MasterPublicDns interface{}
	// The name of the job flow
	Name interface{}
	// The release label for the Amazon EMR release
	ReleaseLabel interface{}
	// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
	ScaleDownBehavior interface{}
	// The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
	SecurityConfiguration interface{}
	// IAM role that will be assumed by the Amazon EMR service to access AWS resources
	ServiceRole interface{}
	Steps interface{}
	// list of tags to apply to the EMR Cluster
	Tags interface{}
	// Switch on/off termination protection (default is off)
	TerminationProtection interface{}
	// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
	VisibleToAllUsers interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	AdditionalInfo interface{}
	// A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, `Spark`, and `JupyterHub` (as of EMR 5.14.0). Case insensitive
	Applications interface{}
	// An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
	AutoscalingRole interface{}
	// List of bootstrap actions that will be run before Hadoop is started on the cluster nodes. Defined below
	BootstrapActions interface{}
	// List of configurations supplied for the EMR cluster you are creating
	Configurations interface{}
	// A JSON string for supplying list of configurations for the EMR cluster.
	ConfigurationsJson interface{}
	// Use the `core_instance_group` configuration block `instance_count` argument instead. Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set. Default `1`
	CoreInstanceCount interface{}
	// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core). Cannot be specified if `core_instance_count` argument, `core_instance_type` argument, or `instance_group` configuration blocks are set. Detailed below.
	CoreInstanceGroup interface{}
	// Use the `core_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the slave nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set.
	CoreInstanceType interface{}
	// A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
	CustomAmiId interface{}
	// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
	EbsRootVolumeSize interface{}
	// Attributes for the EC2 instances running the job flow. Defined below
	Ec2Attributes interface{}
	// Use the `master_instance_group` configuration block, `core_instance_group` configuration block and [`aws_emr_instance_group` resource(s)](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html) instead. A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Cannot be specified if `master_instance_group` or `core_instance_group` configuration blocks are set. Defined below
	InstanceGroups interface{}
	// Switch on/off run cluster with no steps or when all steps are complete (default is on)
	KeepJobFlowAliveWhenNoSteps interface{}
	// Kerberos configuration for the cluster. Defined below
	KerberosAttributes interface{}
	// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created
	LogUri interface{}
	// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master). Cannot be specified if `master_instance_type` argument or `instance_group` configuration blocks are set. Detailed below.
	MasterInstanceGroup interface{}
	// Use the `master_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the master node. Cannot be specified if `master_instance_group` or `instance_group` configuration blocks are set.
	MasterInstanceType interface{}
	// The name of the job flow
	Name interface{}
	// The release label for the Amazon EMR release
	ReleaseLabel interface{}
	// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
	ScaleDownBehavior interface{}
	// The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
	SecurityConfiguration interface{}
	// IAM role that will be assumed by the Amazon EMR service to access AWS resources
	ServiceRole interface{}
	Steps interface{}
	// list of tags to apply to the EMR Cluster
	Tags interface{}
	// Switch on/off termination protection (default is off)
	TerminationProtection interface{}
	// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
	VisibleToAllUsers interface{}
}
