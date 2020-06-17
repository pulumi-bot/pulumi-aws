// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages AWS Managed Streaming for Kafka cluster
//
// ## Example Usage
type Cluster struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
	BootstrapBrokers pulumi.StringOutput `pulumi:"bootstrapBrokers"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
	BootstrapBrokersTls pulumi.StringOutput `pulumi:"bootstrapBrokersTls"`
	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo ClusterBrokerNodeGroupInfoOutput `pulumi:"brokerNodeGroupInfo"`
	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication ClusterClientAuthenticationPtrOutput `pulumi:"clientAuthentication"`
	// Name of the MSK cluster.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	ConfigurationInfo ClusterConfigurationInfoPtrOutput `pulumi:"configurationInfo"`
	// Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
	// * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
	CurrentVersion pulumi.StringOutput `pulumi:"currentVersion"`
	// Configuration block for specifying encryption. See below.
	EncryptionInfo ClusterEncryptionInfoPtrOutput `pulumi:"encryptionInfo"`
	// Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	EnhancedMonitoring pulumi.StringPtrOutput `pulumi:"enhancedMonitoring"`
	// Specify the desired Kafka software version.
	KafkaVersion pulumi.StringOutput `pulumi:"kafkaVersion"`
	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	LoggingInfo ClusterLoggingInfoPtrOutput `pulumi:"loggingInfo"`
	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	NumberOfBrokerNodes pulumi.IntOutput `pulumi:"numberOfBrokerNodes"`
	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	OpenMonitoring ClusterOpenMonitoringPtrOutput `pulumi:"openMonitoring"`
	// A map of tags to assign to the resource
	Tags pulumi.MapOutput `pulumi:"tags"`
	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
	ZookeeperConnectString pulumi.StringOutput `pulumi:"zookeeperConnectString"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.BrokerNodeGroupInfo == nil {
		return nil, errors.New("missing required argument 'BrokerNodeGroupInfo'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.KafkaVersion == nil {
		return nil, errors.New("missing required argument 'KafkaVersion'")
	}
	if args == nil || args.NumberOfBrokerNodes == nil {
		return nil, errors.New("missing required argument 'NumberOfBrokerNodes'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("aws:msk/cluster:Cluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:msk/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn *string `pulumi:"arn"`
	// A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
	BootstrapBrokers *string `pulumi:"bootstrapBrokers"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
	BootstrapBrokersTls *string `pulumi:"bootstrapBrokersTls"`
	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo *ClusterBrokerNodeGroupInfo `pulumi:"brokerNodeGroupInfo"`
	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication *ClusterClientAuthentication `pulumi:"clientAuthentication"`
	// Name of the MSK cluster.
	ClusterName *string `pulumi:"clusterName"`
	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	ConfigurationInfo *ClusterConfigurationInfo `pulumi:"configurationInfo"`
	// Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
	// * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
	CurrentVersion *string `pulumi:"currentVersion"`
	// Configuration block for specifying encryption. See below.
	EncryptionInfo *ClusterEncryptionInfo `pulumi:"encryptionInfo"`
	// Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	EnhancedMonitoring *string `pulumi:"enhancedMonitoring"`
	// Specify the desired Kafka software version.
	KafkaVersion *string `pulumi:"kafkaVersion"`
	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	LoggingInfo *ClusterLoggingInfo `pulumi:"loggingInfo"`
	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	NumberOfBrokerNodes *int `pulumi:"numberOfBrokerNodes"`
	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	OpenMonitoring *ClusterOpenMonitoring `pulumi:"openMonitoring"`
	// A map of tags to assign to the resource
	Tags map[string]interface{} `pulumi:"tags"`
	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
	ZookeeperConnectString *string `pulumi:"zookeeperConnectString"`
}

type ClusterState struct {
	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn pulumi.StringPtrInput
	// A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
	BootstrapBrokers pulumi.StringPtrInput
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
	BootstrapBrokersTls pulumi.StringPtrInput
	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo ClusterBrokerNodeGroupInfoPtrInput
	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication ClusterClientAuthenticationPtrInput
	// Name of the MSK cluster.
	ClusterName pulumi.StringPtrInput
	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	ConfigurationInfo ClusterConfigurationInfoPtrInput
	// Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
	// * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
	CurrentVersion pulumi.StringPtrInput
	// Configuration block for specifying encryption. See below.
	EncryptionInfo ClusterEncryptionInfoPtrInput
	// Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	EnhancedMonitoring pulumi.StringPtrInput
	// Specify the desired Kafka software version.
	KafkaVersion pulumi.StringPtrInput
	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	LoggingInfo ClusterLoggingInfoPtrInput
	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	NumberOfBrokerNodes pulumi.IntPtrInput
	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	OpenMonitoring ClusterOpenMonitoringPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.MapInput
	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
	ZookeeperConnectString pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo ClusterBrokerNodeGroupInfo `pulumi:"brokerNodeGroupInfo"`
	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication *ClusterClientAuthentication `pulumi:"clientAuthentication"`
	// Name of the MSK cluster.
	ClusterName string `pulumi:"clusterName"`
	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	ConfigurationInfo *ClusterConfigurationInfo `pulumi:"configurationInfo"`
	// Configuration block for specifying encryption. See below.
	EncryptionInfo *ClusterEncryptionInfo `pulumi:"encryptionInfo"`
	// Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	EnhancedMonitoring *string `pulumi:"enhancedMonitoring"`
	// Specify the desired Kafka software version.
	KafkaVersion string `pulumi:"kafkaVersion"`
	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	LoggingInfo *ClusterLoggingInfo `pulumi:"loggingInfo"`
	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	NumberOfBrokerNodes int `pulumi:"numberOfBrokerNodes"`
	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	OpenMonitoring *ClusterOpenMonitoring `pulumi:"openMonitoring"`
	// A map of tags to assign to the resource
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo ClusterBrokerNodeGroupInfoInput
	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication ClusterClientAuthenticationPtrInput
	// Name of the MSK cluster.
	ClusterName pulumi.StringInput
	// Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
	ConfigurationInfo ClusterConfigurationInfoPtrInput
	// Configuration block for specifying encryption. See below.
	EncryptionInfo ClusterEncryptionInfoPtrInput
	// Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
	EnhancedMonitoring pulumi.StringPtrInput
	// Specify the desired Kafka software version.
	KafkaVersion pulumi.StringInput
	// Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
	LoggingInfo ClusterLoggingInfoPtrInput
	// The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
	NumberOfBrokerNodes pulumi.IntInput
	// Configuration block for JMX and Node monitoring for the MSK cluster. See below.
	OpenMonitoring ClusterOpenMonitoringPtrInput
	// A map of tags to assign to the resource
	Tags pulumi.MapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
