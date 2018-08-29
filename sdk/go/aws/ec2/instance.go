// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an EC2 instance resource. This allows instances to be created, updated,
// and deleted. Instances also support [provisioning](https://www.terraform.io/docs/provisioners/index.html).
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.Ami == nil {
		return nil, errors.New("missing required argument 'Ami'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["ami"] = nil
		inputs["associatePublicIpAddress"] = nil
		inputs["availabilityZone"] = nil
		inputs["cpuCoreCount"] = nil
		inputs["cpuThreadsPerCore"] = nil
		inputs["creditSpecification"] = nil
		inputs["disableApiTermination"] = nil
		inputs["ebsBlockDevices"] = nil
		inputs["ebsOptimized"] = nil
		inputs["ephemeralBlockDevices"] = nil
		inputs["getPasswordData"] = nil
		inputs["iamInstanceProfile"] = nil
		inputs["instanceInitiatedShutdownBehavior"] = nil
		inputs["instanceType"] = nil
		inputs["ipv6AddressCount"] = nil
		inputs["ipv6Addresses"] = nil
		inputs["keyName"] = nil
		inputs["monitoring"] = nil
		inputs["networkInterfaces"] = nil
		inputs["placementGroup"] = nil
		inputs["privateIp"] = nil
		inputs["rootBlockDevice"] = nil
		inputs["securityGroups"] = nil
		inputs["sourceDestCheck"] = nil
		inputs["subnetId"] = nil
		inputs["tags"] = nil
		inputs["tenancy"] = nil
		inputs["userData"] = nil
		inputs["userDataBase64"] = nil
		inputs["volumeTags"] = nil
		inputs["vpcSecurityGroupIds"] = nil
	} else {
		inputs["ami"] = args.Ami
		inputs["associatePublicIpAddress"] = args.AssociatePublicIpAddress
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["cpuCoreCount"] = args.CpuCoreCount
		inputs["cpuThreadsPerCore"] = args.CpuThreadsPerCore
		inputs["creditSpecification"] = args.CreditSpecification
		inputs["disableApiTermination"] = args.DisableApiTermination
		inputs["ebsBlockDevices"] = args.EbsBlockDevices
		inputs["ebsOptimized"] = args.EbsOptimized
		inputs["ephemeralBlockDevices"] = args.EphemeralBlockDevices
		inputs["getPasswordData"] = args.GetPasswordData
		inputs["iamInstanceProfile"] = args.IamInstanceProfile
		inputs["instanceInitiatedShutdownBehavior"] = args.InstanceInitiatedShutdownBehavior
		inputs["instanceType"] = args.InstanceType
		inputs["ipv6AddressCount"] = args.Ipv6AddressCount
		inputs["ipv6Addresses"] = args.Ipv6Addresses
		inputs["keyName"] = args.KeyName
		inputs["monitoring"] = args.Monitoring
		inputs["networkInterfaces"] = args.NetworkInterfaces
		inputs["placementGroup"] = args.PlacementGroup
		inputs["privateIp"] = args.PrivateIp
		inputs["rootBlockDevice"] = args.RootBlockDevice
		inputs["securityGroups"] = args.SecurityGroups
		inputs["sourceDestCheck"] = args.SourceDestCheck
		inputs["subnetId"] = args.SubnetId
		inputs["tags"] = args.Tags
		inputs["tenancy"] = args.Tenancy
		inputs["userData"] = args.UserData
		inputs["userDataBase64"] = args.UserDataBase64
		inputs["volumeTags"] = args.VolumeTags
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
	}
	inputs["arn"] = nil
	inputs["instanceState"] = nil
	inputs["networkInterfaceId"] = nil
	inputs["passwordData"] = nil
	inputs["primaryNetworkInterfaceId"] = nil
	inputs["privateDns"] = nil
	inputs["publicDns"] = nil
	inputs["publicIp"] = nil
	s, err := ctx.RegisterResource("aws:ec2/instance:Instance", name, true, inputs, opts...)
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
		inputs["ami"] = state.Ami
		inputs["arn"] = state.Arn
		inputs["associatePublicIpAddress"] = state.AssociatePublicIpAddress
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["cpuCoreCount"] = state.CpuCoreCount
		inputs["cpuThreadsPerCore"] = state.CpuThreadsPerCore
		inputs["creditSpecification"] = state.CreditSpecification
		inputs["disableApiTermination"] = state.DisableApiTermination
		inputs["ebsBlockDevices"] = state.EbsBlockDevices
		inputs["ebsOptimized"] = state.EbsOptimized
		inputs["ephemeralBlockDevices"] = state.EphemeralBlockDevices
		inputs["getPasswordData"] = state.GetPasswordData
		inputs["iamInstanceProfile"] = state.IamInstanceProfile
		inputs["instanceInitiatedShutdownBehavior"] = state.InstanceInitiatedShutdownBehavior
		inputs["instanceState"] = state.InstanceState
		inputs["instanceType"] = state.InstanceType
		inputs["ipv6AddressCount"] = state.Ipv6AddressCount
		inputs["ipv6Addresses"] = state.Ipv6Addresses
		inputs["keyName"] = state.KeyName
		inputs["monitoring"] = state.Monitoring
		inputs["networkInterfaces"] = state.NetworkInterfaces
		inputs["networkInterfaceId"] = state.NetworkInterfaceId
		inputs["passwordData"] = state.PasswordData
		inputs["placementGroup"] = state.PlacementGroup
		inputs["primaryNetworkInterfaceId"] = state.PrimaryNetworkInterfaceId
		inputs["privateDns"] = state.PrivateDns
		inputs["privateIp"] = state.PrivateIp
		inputs["publicDns"] = state.PublicDns
		inputs["publicIp"] = state.PublicIp
		inputs["rootBlockDevice"] = state.RootBlockDevice
		inputs["securityGroups"] = state.SecurityGroups
		inputs["sourceDestCheck"] = state.SourceDestCheck
		inputs["subnetId"] = state.SubnetId
		inputs["tags"] = state.Tags
		inputs["tenancy"] = state.Tenancy
		inputs["userData"] = state.UserData
		inputs["userDataBase64"] = state.UserDataBase64
		inputs["volumeTags"] = state.VolumeTags
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
	}
	s, err := ctx.ReadResource("aws:ec2/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The AMI to use for the instance.
func (r *Instance) Ami() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ami"])
}

// The ARN of the instance.
func (r *Instance) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Associate a public ip address with an instance in a VPC.  Boolean value.
func (r *Instance) AssociatePublicIpAddress() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["associatePublicIpAddress"])
}

// The AZ to start the instance in.
func (r *Instance) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// Sets the number of CPU cores for an instance. This option is 
// only supported on creation of instance type that support CPU Options
// [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
func (r *Instance) CpuCoreCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuCoreCount"])
}

// If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
func (r *Instance) CpuThreadsPerCore() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuThreadsPerCore"])
}

// Customize the credit specification of the instance. See Credit Specification below for more details.
func (r *Instance) CreditSpecification() *pulumi.Output {
	return r.s.State["creditSpecification"]
}

// If true, enables [EC2 Instance
// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
func (r *Instance) DisableApiTermination() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disableApiTermination"])
}

// Additional EBS block devices to attach to the
// instance.  See Block Devices below for details.
func (r *Instance) EbsBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ebsBlockDevices"])
}

// If true, the launched EC2 instance will be EBS-optimized.
// Note that if this is not set on an instance type that is optimized by default then
// this will show as disabled but if the instance type is optimized by default then
// there is no need to set this and there is no effect to disabling it.
// See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
func (r *Instance) EbsOptimized() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["ebsOptimized"])
}

// Customize Ephemeral (also known as
// "Instance Store") volumes on the instance. See Block Devices below for details.
func (r *Instance) EphemeralBlockDevices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ephemeralBlockDevices"])
}

// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
func (r *Instance) GetPasswordData() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["getPasswordData"])
}

// The IAM Instance Profile to
// launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
// * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
func (r *Instance) IamInstanceProfile() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["iamInstanceProfile"])
}

// Shutdown behavior for the
// instance. Amazon defaults this to `stop` for EBS-backed instances and
// `terminate` for instance-store instances. Cannot be set on instance-store
// instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
func (r *Instance) InstanceInitiatedShutdownBehavior() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceInitiatedShutdownBehavior"])
}

func (r *Instance) InstanceState() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceState"])
}

// The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
func (r *Instance) InstanceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instanceType"])
}

func (r *Instance) Ipv6AddressCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ipv6AddressCount"])
}

// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
func (r *Instance) Ipv6Addresses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipv6Addresses"])
}

// The key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource.
func (r *Instance) KeyName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyName"])
}

// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
func (r *Instance) Monitoring() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["monitoring"])
}

// Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
func (r *Instance) NetworkInterfaces() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkInterfaces"])
}

// The ID of the network interface that was created with the instance.
func (r *Instance) NetworkInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkInterfaceId"])
}

// Base-64 encoded encrypted password data for the instance.
// Useful for getting the administrator password for instances running Microsoft Windows.
// This attribute is only exported if `get_password_data` is true.
// Note that this encrypted value will be stored in the state file, as with all exported attributes.
// See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
func (r *Instance) PasswordData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordData"])
}

// The Placement Group to start the instance in.
func (r *Instance) PlacementGroup() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["placementGroup"])
}

// The ID of the instance's primary network interface.
func (r *Instance) PrimaryNetworkInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["primaryNetworkInterfaceId"])
}

// The private DNS name assigned to the instance. Can only be
// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
// for your VPC
func (r *Instance) PrivateDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateDns"])
}

// Private IP address to associate with the
// instance in a VPC.
func (r *Instance) PrivateIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIp"])
}

// The public DNS name assigned to the instance. For EC2-VPC, this
// is only available if you've enabled DNS hostnames for your VPC
func (r *Instance) PublicDns() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicDns"])
}

// The public IP address assigned to the instance, if applicable. **NOTE**: If you are using an [`aws_eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
func (r *Instance) PublicIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicIp"])
}

// Customize details about the root block
// device of the instance. See Block Devices below for details.
func (r *Instance) RootBlockDevice() *pulumi.Output {
	return r.s.State["rootBlockDevice"]
}

// A list of security group names to associate with.
func (r *Instance) SecurityGroups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroups"])
}

// Controls if traffic is routed to the instance when
// the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
func (r *Instance) SourceDestCheck() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["sourceDestCheck"])
}

// The VPC Subnet ID to launch in.
func (r *Instance) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// A mapping of tags to assign to the resource.
func (r *Instance) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
func (r *Instance) Tenancy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenancy"])
}

// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
func (r *Instance) UserData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userData"])
}

// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
func (r *Instance) UserDataBase64() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userDataBase64"])
}

// A mapping of tags to assign to the devices created by the instance at launch time.
func (r *Instance) VolumeTags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["volumeTags"])
}

// A list of security group IDs to associate with.
func (r *Instance) VpcSecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vpcSecurityGroupIds"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	// The AMI to use for the instance.
	Ami interface{}
	// The ARN of the instance.
	Arn interface{}
	// Associate a public ip address with an instance in a VPC.  Boolean value.
	AssociatePublicIpAddress interface{}
	// The AZ to start the instance in.
	AvailabilityZone interface{}
	// Sets the number of CPU cores for an instance. This option is 
	// only supported on creation of instance type that support CPU Options
	// [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
	CpuCoreCount interface{}
	// If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
	CpuThreadsPerCore interface{}
	// Customize the credit specification of the instance. See Credit Specification below for more details.
	CreditSpecification interface{}
	// If true, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination interface{}
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices interface{}
	// If true, the launched EC2 instance will be EBS-optimized.
	// Note that if this is not set on an instance type that is optimized by default then
	// this will show as disabled but if the instance type is optimized by default then
	// there is no need to set this and there is no effect to disabling it.
	// See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
	EbsOptimized interface{}
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices interface{}
	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData interface{}
	// The IAM Instance Profile to
	// launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
	// * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	IamInstanceProfile interface{}
	// Shutdown behavior for the
	// instance. Amazon defaults this to `stop` for EBS-backed instances and
	// `terminate` for instance-store instances. Cannot be set on instance-store
	// instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
	InstanceInitiatedShutdownBehavior interface{}
	InstanceState interface{}
	// The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
	InstanceType interface{}
	Ipv6AddressCount interface{}
	// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
	Ipv6Addresses interface{}
	// The key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource.
	KeyName interface{}
	// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
	Monitoring interface{}
	// Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
	NetworkInterfaces interface{}
	// The ID of the network interface that was created with the instance.
	NetworkInterfaceId interface{}
	// Base-64 encoded encrypted password data for the instance.
	// Useful for getting the administrator password for instances running Microsoft Windows.
	// This attribute is only exported if `get_password_data` is true.
	// Note that this encrypted value will be stored in the state file, as with all exported attributes.
	// See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	PasswordData interface{}
	// The Placement Group to start the instance in.
	PlacementGroup interface{}
	// The ID of the instance's primary network interface.
	PrimaryNetworkInterfaceId interface{}
	// The private DNS name assigned to the instance. Can only be
	// used inside the Amazon EC2, and only available if you've enabled DNS hostnames
	// for your VPC
	PrivateDns interface{}
	// Private IP address to associate with the
	// instance in a VPC.
	PrivateIp interface{}
	// The public DNS name assigned to the instance. For EC2-VPC, this
	// is only available if you've enabled DNS hostnames for your VPC
	PublicDns interface{}
	// The public IP address assigned to the instance, if applicable. **NOTE**: If you are using an [`aws_eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
	PublicIp interface{}
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice interface{}
	// A list of security group names to associate with.
	SecurityGroups interface{}
	// Controls if traffic is routed to the instance when
	// the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
	SourceDestCheck interface{}
	// The VPC Subnet ID to launch in.
	SubnetId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
	Tenancy interface{}
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
	UserData interface{}
	// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 interface{}
	// A mapping of tags to assign to the devices created by the instance at launch time.
	VolumeTags interface{}
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The AMI to use for the instance.
	Ami interface{}
	// Associate a public ip address with an instance in a VPC.  Boolean value.
	AssociatePublicIpAddress interface{}
	// The AZ to start the instance in.
	AvailabilityZone interface{}
	// Sets the number of CPU cores for an instance. This option is 
	// only supported on creation of instance type that support CPU Options
	// [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
	CpuCoreCount interface{}
	// If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
	CpuThreadsPerCore interface{}
	// Customize the credit specification of the instance. See Credit Specification below for more details.
	CreditSpecification interface{}
	// If true, enables [EC2 Instance
	// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
	DisableApiTermination interface{}
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices interface{}
	// If true, the launched EC2 instance will be EBS-optimized.
	// Note that if this is not set on an instance type that is optimized by default then
	// this will show as disabled but if the instance type is optimized by default then
	// there is no need to set this and there is no effect to disabling it.
	// See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
	EbsOptimized interface{}
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices interface{}
	// If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
	GetPasswordData interface{}
	// The IAM Instance Profile to
	// launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
	// * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
	IamInstanceProfile interface{}
	// Shutdown behavior for the
	// instance. Amazon defaults this to `stop` for EBS-backed instances and
	// `terminate` for instance-store instances. Cannot be set on instance-store
	// instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
	InstanceInitiatedShutdownBehavior interface{}
	// The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
	InstanceType interface{}
	Ipv6AddressCount interface{}
	// Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
	Ipv6Addresses interface{}
	// The key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource.
	KeyName interface{}
	// If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
	Monitoring interface{}
	// Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
	NetworkInterfaces interface{}
	// The Placement Group to start the instance in.
	PlacementGroup interface{}
	// Private IP address to associate with the
	// instance in a VPC.
	PrivateIp interface{}
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice interface{}
	// A list of security group names to associate with.
	SecurityGroups interface{}
	// Controls if traffic is routed to the instance when
	// the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
	SourceDestCheck interface{}
	// The VPC Subnet ID to launch in.
	SubnetId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
	Tenancy interface{}
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
	UserData interface{}
	// Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 interface{}
	// A mapping of tags to assign to the devices created by the instance at launch time.
	VolumeTags interface{}
	// A list of security group IDs to associate with.
	VpcSecurityGroupIds interface{}
}
