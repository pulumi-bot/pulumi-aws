// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DAX Cluster resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dax"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = dax.NewCluster(ctx, "bar", &dax.ClusterArgs{
// 			ClusterName:       pulumi.String("cluster-example"),
// 			IamRoleArn:        pulumi.String(data.Aws_iam_role.Example.Arn),
// 			NodeType:          pulumi.String("dax.r4.large"),
// 			ReplicationFactor: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// The ARN of the DAX cluster
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The DNS name of the DAX cluster without the port appended
	ClusterAddress pulumi.StringOutput `pulumi:"clusterAddress"`
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint pulumi.StringOutput `pulumi:"configurationEndpoint"`
	// Description for the cluster
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringOutput `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// List of node objects including `id`, `address`, `port` and
	// `availabilityZone`. Referenceable e.g. as
	// `${aws_dax_cluster.test.nodes.0.address}`
	Nodes ClusterNodeArrayOutput `pulumi:"nodes"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrOutput `pulumi:"notificationTopicArn"`
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
	// The port used by the configuration endpoint
	Port pulumi.IntOutput `pulumi:"port"`
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor pulumi.IntOutput `pulumi:"replicationFactor"`
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Encrypt at rest options
	ServerSideEncryption ClusterServerSideEncryptionPtrOutput `pulumi:"serverSideEncryption"`
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName pulumi.StringOutput `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.IamRoleArn == nil {
		return nil, errors.New("missing required argument 'IamRoleArn'")
	}
	if args == nil || args.NodeType == nil {
		return nil, errors.New("missing required argument 'NodeType'")
	}
	if args == nil || args.ReplicationFactor == nil {
		return nil, errors.New("missing required argument 'ReplicationFactor'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:dax/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws:dax/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The ARN of the DAX cluster
	Arn *string `pulumi:"arn"`
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The DNS name of the DAX cluster without the port appended
	ClusterAddress *string `pulumi:"clusterAddress"`
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName *string `pulumi:"clusterName"`
	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint *string `pulumi:"configurationEndpoint"`
	// Description for the cluster
	Description *string `pulumi:"description"`
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType *string `pulumi:"nodeType"`
	// List of node objects including `id`, `address`, `port` and
	// `availabilityZone`. Referenceable e.g. as
	// `${aws_dax_cluster.test.nodes.0.address}`
	Nodes []ClusterNode `pulumi:"nodes"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn *string `pulumi:"notificationTopicArn"`
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The port used by the configuration endpoint
	Port *int `pulumi:"port"`
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor *int `pulumi:"replicationFactor"`
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Encrypt at rest options
	ServerSideEncryption *ClusterServerSideEncryption `pulumi:"serverSideEncryption"`
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource
	Tags map[string]interface{} `pulumi:"tags"`
}

type ClusterState struct {
	// The ARN of the DAX cluster
	Arn pulumi.StringPtrInput
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones pulumi.StringArrayInput
	// The DNS name of the DAX cluster without the port appended
	ClusterAddress pulumi.StringPtrInput
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName pulumi.StringPtrInput
	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint pulumi.StringPtrInput
	// Description for the cluster
	Description pulumi.StringPtrInput
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn pulumi.StringPtrInput
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType pulumi.StringPtrInput
	// List of node objects including `id`, `address`, `port` and
	// `availabilityZone`. Referenceable e.g. as
	// `${aws_dax_cluster.test.nodes.0.address}`
	Nodes ClusterNodeArrayInput
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrInput
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName pulumi.StringPtrInput
	// The port used by the configuration endpoint
	Port pulumi.IntPtrInput
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor pulumi.IntPtrInput
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds pulumi.StringArrayInput
	// Encrypt at rest options
	ServerSideEncryption ClusterServerSideEncryptionPtrInput
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName pulumi.StringPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.MapInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName string `pulumi:"clusterName"`
	// Description for the cluster
	Description *string `pulumi:"description"`
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn string `pulumi:"iamRoleArn"`
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType string `pulumi:"nodeType"`
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn *string `pulumi:"notificationTopicArn"`
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor int `pulumi:"replicationFactor"`
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Encrypt at rest options
	ServerSideEncryption *ClusterServerSideEncryption `pulumi:"serverSideEncryption"`
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A map of tags to assign to the resource
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// List of Availability Zones in which the
	// nodes will be created
	AvailabilityZones pulumi.StringArrayInput
	// Group identifier. DAX converts this name to
	// lowercase
	ClusterName pulumi.StringInput
	// Description for the cluster
	Description pulumi.StringPtrInput
	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	IamRoleArn pulumi.StringInput
	// Specifies the weekly time range for when
	// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// `sun:05:00-sun:09:00`
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes. See
	// [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
	NodeType pulumi.StringInput
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send DAX notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn pulumi.StringPtrInput
	// Name of the parameter group to associate
	// with this DAX cluster
	ParameterGroupName pulumi.StringPtrInput
	// The number of nodes in the DAX cluster. A
	// replication factor of 1 will create a single-node cluster, without any read
	// replicas
	ReplicationFactor pulumi.IntInput
	// One or more VPC security groups associated
	// with the cluster
	SecurityGroupIds pulumi.StringArrayInput
	// Encrypt at rest options
	ServerSideEncryption ClusterServerSideEncryptionPtrInput
	// Name of the subnet group to be used for the
	// cluster
	SubnetGroupName pulumi.StringPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.MapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
