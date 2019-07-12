// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_instance.html.markdown.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.LayerIds == nil {
		return nil, errors.New("missing required argument 'LayerIds'")
	}
	if args == nil || args.StackId == nil {
		return nil, errors.New("missing required argument 'StackId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["agentVersion"] = nil
		inputs["amiId"] = nil
		inputs["architecture"] = nil
		inputs["autoScalingType"] = nil
		inputs["availabilityZone"] = nil
		inputs["createdAt"] = nil
		inputs["deleteEbs"] = nil
		inputs["deleteEip"] = nil
		inputs["ebsBlockDevices"] = nil
		inputs["ebsOptimized"] = nil
		inputs["ecsClusterArn"] = nil
		inputs["elasticIp"] = nil
		inputs["ephemeralBlockDevices"] = nil
		inputs["hostname"] = nil
		inputs["infrastructureClass"] = nil
		inputs["installUpdatesOnBoot"] = nil
		inputs["instanceProfileArn"] = nil
		inputs["instanceType"] = nil
		inputs["lastServiceErrorId"] = nil
		inputs["layerIds"] = nil
		inputs["os"] = nil
		inputs["platform"] = nil
		inputs["privateDns"] = nil
		inputs["privateIp"] = nil
		inputs["publicDns"] = nil
		inputs["publicIp"] = nil
		inputs["registeredBy"] = nil
		inputs["reportedAgentVersion"] = nil
		inputs["reportedOsFamily"] = nil
		inputs["reportedOsName"] = nil
		inputs["reportedOsVersion"] = nil
		inputs["rootBlockDevices"] = nil
		inputs["rootDeviceType"] = nil
		inputs["rootDeviceVolumeId"] = nil
		inputs["securityGroupIds"] = nil
		inputs["sshHostDsaKeyFingerprint"] = nil
		inputs["sshHostRsaKeyFingerprint"] = nil
		inputs["sshKeyName"] = nil
		inputs["stackId"] = nil
		inputs["state"] = nil
		inputs["status"] = nil
		inputs["subnetId"] = nil
		inputs["tenancy"] = nil
		inputs["virtualizationType"] = nil
	} else {
		inputs["agentVersion"] = args.AgentVersion
		inputs["amiId"] = args.AmiId
		inputs["architecture"] = args.Architecture
		inputs["autoScalingType"] = args.AutoScalingType
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["createdAt"] = args.CreatedAt
		inputs["deleteEbs"] = args.DeleteEbs
		inputs["deleteEip"] = args.DeleteEip
		inputs["ebsBlockDevices"] = args.EbsBlockDevices
		inputs["ebsOptimized"] = args.EbsOptimized
		inputs["ecsClusterArn"] = args.EcsClusterArn
		inputs["elasticIp"] = args.ElasticIp
		inputs["ephemeralBlockDevices"] = args.EphemeralBlockDevices
		inputs["hostname"] = args.Hostname
		inputs["infrastructureClass"] = args.InfrastructureClass
		inputs["installUpdatesOnBoot"] = args.InstallUpdatesOnBoot
		inputs["instanceProfileArn"] = args.InstanceProfileArn
		inputs["instanceType"] = args.InstanceType
		inputs["lastServiceErrorId"] = args.LastServiceErrorId
		inputs["layerIds"] = args.LayerIds
		inputs["os"] = args.Os
		inputs["platform"] = args.Platform
		inputs["privateDns"] = args.PrivateDns
		inputs["privateIp"] = args.PrivateIp
		inputs["publicDns"] = args.PublicDns
		inputs["publicIp"] = args.PublicIp
		inputs["registeredBy"] = args.RegisteredBy
		inputs["reportedAgentVersion"] = args.ReportedAgentVersion
		inputs["reportedOsFamily"] = args.ReportedOsFamily
		inputs["reportedOsName"] = args.ReportedOsName
		inputs["reportedOsVersion"] = args.ReportedOsVersion
		inputs["rootBlockDevices"] = args.RootBlockDevices
		inputs["rootDeviceType"] = args.RootDeviceType
		inputs["rootDeviceVolumeId"] = args.RootDeviceVolumeId
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["sshHostDsaKeyFingerprint"] = args.SshHostDsaKeyFingerprint
		inputs["sshHostRsaKeyFingerprint"] = args.SshHostRsaKeyFingerprint
		inputs["sshKeyName"] = args.SshKeyName
		inputs["stackId"] = args.StackId
		inputs["state"] = args.State
		inputs["status"] = args.Status
		inputs["subnetId"] = args.SubnetId
		inputs["tenancy"] = args.Tenancy
		inputs["virtualizationType"] = args.VirtualizationType
	}
	inputs["ec2InstanceId"] = nil
	s, err := ctx.RegisterResource("aws:opsworks/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["agentVersion"] = state.AgentVersion
		inputs["amiId"] = state.AmiId
		inputs["architecture"] = state.Architecture
		inputs["autoScalingType"] = state.AutoScalingType
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["createdAt"] = state.CreatedAt
		inputs["deleteEbs"] = state.DeleteEbs
		inputs["deleteEip"] = state.DeleteEip
		inputs["ebsBlockDevices"] = state.EbsBlockDevices
		inputs["ebsOptimized"] = state.EbsOptimized
		inputs["ec2InstanceId"] = state.Ec2InstanceId
		inputs["ecsClusterArn"] = state.EcsClusterArn
		inputs["elasticIp"] = state.ElasticIp
		inputs["ephemeralBlockDevices"] = state.EphemeralBlockDevices
		inputs["hostname"] = state.Hostname
		inputs["infrastructureClass"] = state.InfrastructureClass
		inputs["installUpdatesOnBoot"] = state.InstallUpdatesOnBoot
		inputs["instanceProfileArn"] = state.InstanceProfileArn
		inputs["instanceType"] = state.InstanceType
		inputs["lastServiceErrorId"] = state.LastServiceErrorId
		inputs["layerIds"] = state.LayerIds
		inputs["os"] = state.Os
		inputs["platform"] = state.Platform
		inputs["privateDns"] = state.PrivateDns
		inputs["privateIp"] = state.PrivateIp
		inputs["publicDns"] = state.PublicDns
		inputs["publicIp"] = state.PublicIp
		inputs["registeredBy"] = state.RegisteredBy
		inputs["reportedAgentVersion"] = state.ReportedAgentVersion
		inputs["reportedOsFamily"] = state.ReportedOsFamily
		inputs["reportedOsName"] = state.ReportedOsName
		inputs["reportedOsVersion"] = state.ReportedOsVersion
		inputs["rootBlockDevices"] = state.RootBlockDevices
		inputs["rootDeviceType"] = state.RootDeviceType
		inputs["rootDeviceVolumeId"] = state.RootDeviceVolumeId
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["sshHostDsaKeyFingerprint"] = state.SshHostDsaKeyFingerprint
		inputs["sshHostRsaKeyFingerprint"] = state.SshHostRsaKeyFingerprint
		inputs["sshKeyName"] = state.SshKeyName
		inputs["stackId"] = state.StackId
		inputs["state"] = state.State
		inputs["status"] = state.Status
		inputs["subnetId"] = state.SubnetId
		inputs["tenancy"] = state.Tenancy
		inputs["virtualizationType"] = state.VirtualizationType
	}
	s, err := ctx.ReadResource("aws:opsworks/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
func (r *Instance) AgentVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["agentVersion"])
}

// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
func (r *Instance) AmiId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["amiId"])
}

// Machine architecture for created instances.  Can be either `"x86_64"` (the default) or `"i386"`
func (r *Instance) Architecture() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["architecture"])
}

// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
func (r *Instance) AutoScalingType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["autoScalingType"])
}

// Name of the availability zone where instances will be created
// by default.
func (r *Instance) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

func (r *Instance) CreatedAt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createdAt"])
}

func (r *Instance) DeleteEbs() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["deleteEbs"])
}

func (r *Instance) DeleteEip() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["deleteEip"])
}

// Additional EBS block devices to attach to the
// instance.  See Block Devices below for details.
func (r *Instance) EbsBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ebsBlockDevices"])
}

// If true, the launched EC2 instance will be EBS-optimized.
func (r *Instance) EbsOptimized() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ebsOptimized"])
}

// EC2 instance ID
func (r *Instance) Ec2InstanceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ec2InstanceId"])
}

func (r *Instance) EcsClusterArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ecsClusterArn"])
}

func (r *Instance) ElasticIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["elasticIp"])
}

// Customize Ephemeral (also known as
// "Instance Store") volumes on the instance. See Block Devices below for details.
func (r *Instance) EphemeralBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ephemeralBlockDevices"])
}

// The instance's host name.
func (r *Instance) Hostname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["hostname"])
}

func (r *Instance) InfrastructureClass() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["infrastructureClass"])
}

// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
func (r *Instance) InstallUpdatesOnBoot() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["installUpdatesOnBoot"])
}

func (r *Instance) InstanceProfileArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceProfileArn"])
}

// The type of instance to start
func (r *Instance) InstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceType"])
}

func (r *Instance) LastServiceErrorId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastServiceErrorId"])
}

// The ids of the layers the instance will belong to.
func (r *Instance) LayerIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["layerIds"])
}

// Name of operating system that will be installed.
func (r *Instance) Os() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["os"])
}

func (r *Instance) Platform() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["platform"])
}

// The private DNS name assigned to the instance. Can only be
// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
// for your VPC
func (r *Instance) PrivateDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateDns"])
}

// The private IP address assigned to the instance
func (r *Instance) PrivateIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIp"])
}

// The public DNS name assigned to the instance. For EC2-VPC, this
// is only available if you've enabled DNS hostnames for your VPC
func (r *Instance) PublicDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicDns"])
}

// The public IP address assigned to the instance, if applicable.
func (r *Instance) PublicIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicIp"])
}

func (r *Instance) RegisteredBy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["registeredBy"])
}

func (r *Instance) ReportedAgentVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reportedAgentVersion"])
}

func (r *Instance) ReportedOsFamily() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reportedOsFamily"])
}

func (r *Instance) ReportedOsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reportedOsName"])
}

func (r *Instance) ReportedOsVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reportedOsVersion"])
}

// Customize details about the root block
// device of the instance. See Block Devices below for details.
func (r *Instance) RootBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["rootBlockDevices"])
}

// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
func (r *Instance) RootDeviceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["rootDeviceType"])
}

func (r *Instance) RootDeviceVolumeId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["rootDeviceVolumeId"])
}

// The associated security groups.
func (r *Instance) SecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroupIds"])
}

func (r *Instance) SshHostDsaKeyFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sshHostDsaKeyFingerprint"])
}

func (r *Instance) SshHostRsaKeyFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sshHostRsaKeyFingerprint"])
}

// Name of the SSH keypair that instances will have by default.
func (r *Instance) SshKeyName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sshKeyName"])
}

// The id of the stack the instance will belong to.
func (r *Instance) StackId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["stackId"])
}

// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
func (r *Instance) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

func (r *Instance) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Subnet ID to attach to
func (r *Instance) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
func (r *Instance) Tenancy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenancy"])
}

// Keyword to choose what virtualization mode created instances
// will use. Can be either `"paravirtual"` or `"hvm"`.
func (r *Instance) VirtualizationType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["virtualizationType"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion interface{}
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId interface{}
	// Machine architecture for created instances.  Can be either `"x86_64"` (the default) or `"i386"`
	Architecture interface{}
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType interface{}
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone interface{}
	CreatedAt interface{}
	DeleteEbs interface{}
	DeleteEip interface{}
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices interface{}
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized interface{}
	// EC2 instance ID
	Ec2InstanceId interface{}
	EcsClusterArn interface{}
	ElasticIp interface{}
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices interface{}
	// The instance's host name.
	Hostname interface{}
	InfrastructureClass interface{}
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot interface{}
	InstanceProfileArn interface{}
	// The type of instance to start
	InstanceType interface{}
	LastServiceErrorId interface{}
	// The ids of the layers the instance will belong to.
	LayerIds interface{}
	// Name of operating system that will be installed.
	Os interface{}
	Platform interface{}
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns interface{}
	// The private IP address assigned to the instance
	PrivateIp interface{}
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns interface{}
	// The public IP address assigned to the instance, if applicable.
	PublicIp interface{}
	RegisteredBy interface{}
	ReportedAgentVersion interface{}
	ReportedOsFamily interface{}
	ReportedOsName interface{}
	ReportedOsVersion interface{}
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices interface{}
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType interface{}
	RootDeviceVolumeId interface{}
	// The associated security groups.
	SecurityGroupIds interface{}
	SshHostDsaKeyFingerprint interface{}
	SshHostRsaKeyFingerprint interface{}
	// Name of the SSH keypair that instances will have by default.
	SshKeyName interface{}
	// The id of the stack the instance will belong to.
	StackId interface{}
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State interface{}
	Status interface{}
	// Subnet ID to attach to
	SubnetId interface{}
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy interface{}
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The AWS OpsWorks agent to install.  Defaults to `"INHERIT"`.
	AgentVersion interface{}
	// The AMI to use for the instance.  If an AMI is specified, `os` must be `"Custom"`.
	AmiId interface{}
	// Machine architecture for created instances.  Can be either `"x86_64"` (the default) or `"i386"`
	Architecture interface{}
	// Creates load-based or time-based instances.  If set, can be either: `"load"` or `"timer"`.
	AutoScalingType interface{}
	// Name of the availability zone where instances will be created
	// by default.
	AvailabilityZone interface{}
	CreatedAt interface{}
	DeleteEbs interface{}
	DeleteEip interface{}
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices interface{}
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized interface{}
	EcsClusterArn interface{}
	ElasticIp interface{}
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices interface{}
	// The instance's host name.
	Hostname interface{}
	InfrastructureClass interface{}
	// Controls where to install OS and package updates when the instance boots.  Defaults to `true`.
	InstallUpdatesOnBoot interface{}
	InstanceProfileArn interface{}
	// The type of instance to start
	InstanceType interface{}
	LastServiceErrorId interface{}
	// The ids of the layers the instance will belong to.
	LayerIds interface{}
	// Name of operating system that will be installed.
	Os interface{}
	Platform interface{}
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns interface{}
	// The private IP address assigned to the instance
	PrivateIp interface{}
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns interface{}
	// The public IP address assigned to the instance, if applicable.
	PublicIp interface{}
	RegisteredBy interface{}
	ReportedAgentVersion interface{}
	ReportedOsFamily interface{}
	ReportedOsName interface{}
	ReportedOsVersion interface{}
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevices interface{}
	// Name of the type of root device instances will have by default.  Can be either `"ebs"` or `"instance-store"`
	RootDeviceType interface{}
	RootDeviceVolumeId interface{}
	// The associated security groups.
	SecurityGroupIds interface{}
	SshHostDsaKeyFingerprint interface{}
	SshHostRsaKeyFingerprint interface{}
	// Name of the SSH keypair that instances will have by default.
	SshKeyName interface{}
	// The id of the stack the instance will belong to.
	StackId interface{}
	// The desired state of the instance.  Can be either `"running"` or `"stopped"`.
	State interface{}
	Status interface{}
	// Subnet ID to attach to
	SubnetId interface{}
	// Instance tenancy to use. Can be one of `"default"`, `"dedicated"` or `"host"`
	Tenancy interface{}
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either `"paravirtual"` or `"hvm"`.
	VirtualizationType interface{}
}
