// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication instance resource. DMS replication instances can be created, updated, deleted, and imported.
type ReplicationInstance struct {
	s *pulumi.ResourceState
}

// NewReplicationInstance registers a new resource with the given unique name, arguments, and options.
func NewReplicationInstance(ctx *pulumi.Context,
	name string, args *ReplicationInstanceArgs, opts ...pulumi.ResourceOpt) (*ReplicationInstance, error) {
	if args == nil || args.ReplicationInstanceClass == nil {
		return nil, errors.New("missing required argument 'ReplicationInstanceClass'")
	}
	if args == nil || args.ReplicationInstanceId == nil {
		return nil, errors.New("missing required argument 'ReplicationInstanceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allocatedStorage"] = nil
		inputs["applyImmediately"] = nil
		inputs["autoMinorVersionUpgrade"] = nil
		inputs["availabilityZone"] = nil
		inputs["engineVersion"] = nil
		inputs["kmsKeyArn"] = nil
		inputs["multiAz"] = nil
		inputs["preferredMaintenanceWindow"] = nil
		inputs["publiclyAccessible"] = nil
		inputs["replicationInstanceClass"] = nil
		inputs["replicationInstanceId"] = nil
		inputs["replicationSubnetGroupId"] = nil
		inputs["tags"] = nil
		inputs["vpcSecurityGroupIds"] = nil
	} else {
		inputs["allocatedStorage"] = args.AllocatedStorage
		inputs["applyImmediately"] = args.ApplyImmediately
		inputs["autoMinorVersionUpgrade"] = args.AutoMinorVersionUpgrade
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["engineVersion"] = args.EngineVersion
		inputs["kmsKeyArn"] = args.KmsKeyArn
		inputs["multiAz"] = args.MultiAz
		inputs["preferredMaintenanceWindow"] = args.PreferredMaintenanceWindow
		inputs["publiclyAccessible"] = args.PubliclyAccessible
		inputs["replicationInstanceClass"] = args.ReplicationInstanceClass
		inputs["replicationInstanceId"] = args.ReplicationInstanceId
		inputs["replicationSubnetGroupId"] = args.ReplicationSubnetGroupId
		inputs["tags"] = args.Tags
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
	}
	inputs["replicationInstanceArn"] = nil
	inputs["replicationInstancePrivateIps"] = nil
	inputs["replicationInstancePublicIps"] = nil
	s, err := ctx.RegisterResource("aws:dms/replicationInstance:ReplicationInstance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReplicationInstance{s: s}, nil
}

// GetReplicationInstance gets an existing ReplicationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ReplicationInstanceState, opts ...pulumi.ResourceOpt) (*ReplicationInstance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allocatedStorage"] = state.AllocatedStorage
		inputs["applyImmediately"] = state.ApplyImmediately
		inputs["autoMinorVersionUpgrade"] = state.AutoMinorVersionUpgrade
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["engineVersion"] = state.EngineVersion
		inputs["kmsKeyArn"] = state.KmsKeyArn
		inputs["multiAz"] = state.MultiAz
		inputs["preferredMaintenanceWindow"] = state.PreferredMaintenanceWindow
		inputs["publiclyAccessible"] = state.PubliclyAccessible
		inputs["replicationInstanceArn"] = state.ReplicationInstanceArn
		inputs["replicationInstanceClass"] = state.ReplicationInstanceClass
		inputs["replicationInstanceId"] = state.ReplicationInstanceId
		inputs["replicationInstancePrivateIps"] = state.ReplicationInstancePrivateIps
		inputs["replicationInstancePublicIps"] = state.ReplicationInstancePublicIps
		inputs["replicationSubnetGroupId"] = state.ReplicationSubnetGroupId
		inputs["tags"] = state.Tags
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
	}
	s, err := ctx.ReadResource("aws:dms/replicationInstance:ReplicationInstance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReplicationInstance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ReplicationInstance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ReplicationInstance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
func (r *ReplicationInstance) AllocatedStorage() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["allocatedStorage"])
}

// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
func (r *ReplicationInstance) ApplyImmediately() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["applyImmediately"])
}

// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
func (r *ReplicationInstance) AutoMinorVersionUpgrade() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoMinorVersionUpgrade"])
}

// The EC2 Availability Zone that the replication instance will be created in.
func (r *ReplicationInstance) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The engine version number of the replication instance.
func (r *ReplicationInstance) EngineVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["engineVersion"])
}

// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
func (r *ReplicationInstance) KmsKeyArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["kmsKeyArn"])
}

// Specifies if the replication instance is a multi-az deployment. You cannot set the `availability_zone` parameter if the `multi_az` parameter is set to `true`.
func (r *ReplicationInstance) MultiAz() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["multiAz"])
}

// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
func (r *ReplicationInstance) PreferredMaintenanceWindow() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["preferredMaintenanceWindow"])
}

// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
func (r *ReplicationInstance) PubliclyAccessible() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["publiclyAccessible"])
}

// The Amazon Resource Name (ARN) of the replication instance.
func (r *ReplicationInstance) ReplicationInstanceArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationInstanceArn"])
}

// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
func (r *ReplicationInstance) ReplicationInstanceClass() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationInstanceClass"])
}

// The replication instance identifier. This parameter is stored as a lowercase string.
func (r *ReplicationInstance) ReplicationInstanceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationInstanceId"])
}

// A list of the private IP addresses of the replication instance.
func (r *ReplicationInstance) ReplicationInstancePrivateIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["replicationInstancePrivateIps"])
}

// A list of the public IP addresses of the replication instance.
func (r *ReplicationInstance) ReplicationInstancePublicIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["replicationInstancePublicIps"])
}

// A subnet group to associate with the replication instance.
func (r *ReplicationInstance) ReplicationSubnetGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationSubnetGroupId"])
}

// A mapping of tags to assign to the resource.
func (r *ReplicationInstance) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
func (r *ReplicationInstance) VpcSecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vpcSecurityGroupIds"])
}

// Input properties used for looking up and filtering ReplicationInstance resources.
type ReplicationInstanceState struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage interface{}
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately interface{}
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade interface{}
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone interface{}
	// The engine version number of the replication instance.
	EngineVersion interface{}
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn interface{}
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availability_zone` parameter if the `multi_az` parameter is set to `true`.
	MultiAz interface{}
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow interface{}
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible interface{}
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn interface{}
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass interface{}
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId interface{}
	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps interface{}
	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps interface{}
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds interface{}
}

// The set of arguments for constructing a ReplicationInstance resource.
type ReplicationInstanceArgs struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage interface{}
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately interface{}
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade interface{}
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone interface{}
	// The engine version number of the replication instance.
	EngineVersion interface{}
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn interface{}
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availability_zone` parameter if the `multi_az` parameter is set to `true`.
	MultiAz interface{}
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow interface{}
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible interface{}
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass interface{}
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId interface{}
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds interface{}
}
