// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages AWS Managed Streaming for Kafka cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kinesis"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/msk"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("192.168.0.0/22"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "available"
// 		azs, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
// 			State: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		subnetAz1, err := ec2.NewSubnet(ctx, "subnetAz1", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String(azs.Names[0]),
// 			CidrBlock:        pulumi.String("192.168.0.0/24"),
// 			VpcId:            vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnetAz2, err := ec2.NewSubnet(ctx, "subnetAz2", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String(azs.Names[1]),
// 			CidrBlock:        pulumi.String("192.168.1.0/24"),
// 			VpcId:            vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnetAz3, err := ec2.NewSubnet(ctx, "subnetAz3", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String(azs.Names[2]),
// 			CidrBlock:        pulumi.String("192.168.2.0/24"),
// 			VpcId:            vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
// 			VpcId: vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		kms, err := kms.NewKey(ctx, "kms", &kms.KeyArgs{
// 			Description: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test, err := cloudwatch.NewLogGroup(ctx, "test", nil)
// 		if err != nil {
// 			return err
// 		}
// 		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		firehoseRole, err := iam.NewRole(ctx, "firehoseRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "\"Version\": \"2012-10-17\",\n", "\"Statement\": [\n", "  {\n", "    \"Action\": \"sts:AssumeRole\",\n", "    \"Principal\": {\n", "      \"Service\": \"firehose.amazonaws.com\"\n", "    },\n", "    \"Effect\": \"Allow\",\n", "    \"Sid\": \"\"\n", "  }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testStream, err := kinesis.NewFirehoseDeliveryStream(ctx, "testStream", &kinesis.FirehoseDeliveryStreamArgs{
// 			Destination: pulumi.String("s3"),
// 			S3Configuration: &kinesis.FirehoseDeliveryStreamS3ConfigurationArgs{
// 				RoleArn:   firehoseRole.Arn,
// 				BucketArn: bucket.Arn,
// 			},
// 			Tags: pulumi.StringMap{
// 				"LogDeliveryEnabled": pulumi.String("placeholder"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := msk.NewCluster(ctx, "example", &msk.ClusterArgs{
// 			KafkaVersion:        pulumi.String("2.4.1"),
// 			NumberOfBrokerNodes: pulumi.Int(3),
// 			BrokerNodeGroupInfo: &msk.ClusterBrokerNodeGroupInfoArgs{
// 				InstanceType:  pulumi.String("kafka.m5.large"),
// 				EbsVolumeSize: pulumi.Int(1000),
// 				ClientSubnets: pulumi.StringArray{
// 					subnetAz1.ID(),
// 					subnetAz2.ID(),
// 					subnetAz3.ID(),
// 				},
// 				SecurityGroups: pulumi.StringArray{
// 					sg.ID(),
// 				},
// 			},
// 			EncryptionInfo: &msk.ClusterEncryptionInfoArgs{
// 				EncryptionAtRestKmsKeyArn: kms.Arn,
// 			},
// 			OpenMonitoring: &msk.ClusterOpenMonitoringArgs{
// 				Prometheus: &msk.ClusterOpenMonitoringPrometheusArgs{
// 					JmxExporter: &msk.ClusterOpenMonitoringPrometheusJmxExporterArgs{
// 						EnabledInBroker: pulumi.Bool(true),
// 					},
// 					NodeExporter: &msk.ClusterOpenMonitoringPrometheusNodeExporterArgs{
// 						EnabledInBroker: pulumi.Bool(true),
// 					},
// 				},
// 			},
// 			LoggingInfo: &msk.ClusterLoggingInfoArgs{
// 				BrokerLogs: &msk.ClusterLoggingInfoBrokerLogsArgs{
// 					CloudwatchLogs: &msk.ClusterLoggingInfoBrokerLogsCloudwatchLogsArgs{
// 						Enabled:  pulumi.Bool(true),
// 						LogGroup: test.Name,
// 					},
// 					Firehose: &msk.ClusterLoggingInfoBrokerLogsFirehoseArgs{
// 						Enabled:        pulumi.Bool(true),
// 						DeliveryStream: testStream.Name,
// 					},
// 					S3: &msk.ClusterLoggingInfoBrokerLogsS3Args{
// 						Enabled: pulumi.Bool(true),
// 						Bucket:  bucket.ID(),
// 						Prefix:  pulumi.String("logs/msk-"),
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("zookeeperConnectString", example.ZookeeperConnectString)
// 		ctx.Export("bootstrapBrokersTls", example.BootstrapBrokersTls)
// 		return nil
// 	})
// }
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
	BootstrapBrokers pulumi.StringOutput `pulumi:"bootstrapBrokers"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to bootstrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
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
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
	ZookeeperConnectString pulumi.StringOutput `pulumi:"zookeeperConnectString"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.BrokerNodeGroupInfo == nil {
		return nil, errors.New("missing required argument 'BrokerNodeGroupInfo'")
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
	// A comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
	BootstrapBrokers *string `pulumi:"bootstrapBrokers"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to bootstrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
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
	Tags map[string]string `pulumi:"tags"`
	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
	ZookeeperConnectString *string `pulumi:"zookeeperConnectString"`
}

type ClusterState struct {
	// Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
	Arn pulumi.StringPtrInput
	// A comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
	BootstrapBrokers pulumi.StringPtrInput
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to bootstrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
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
	Tags pulumi.StringMapInput
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
	ClusterName *string `pulumi:"clusterName"`
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
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Configuration block for the broker nodes of the Kafka cluster.
	BrokerNodeGroupInfo ClusterBrokerNodeGroupInfoInput
	// Configuration block for specifying a client authentication. See below.
	ClientAuthentication ClusterClientAuthenticationPtrInput
	// Name of the MSK cluster.
	ClusterName pulumi.StringPtrInput
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
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((*Cluster)(nil)).Elem()
}

func (i Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct {
	*pulumi.OutputState
}

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterOutput)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterOutput{})
}
