# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Cluster']


class Cluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broker_node_group_info: Optional[pulumi.Input[pulumi.InputType['ClusterBrokerNodeGroupInfoArgs']]] = None,
                 client_authentication: Optional[pulumi.Input[pulumi.InputType['ClusterClientAuthenticationArgs']]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 configuration_info: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationInfoArgs']]] = None,
                 encryption_info: Optional[pulumi.Input[pulumi.InputType['ClusterEncryptionInfoArgs']]] = None,
                 enhanced_monitoring: Optional[pulumi.Input[str]] = None,
                 kafka_version: Optional[pulumi.Input[str]] = None,
                 logging_info: Optional[pulumi.Input[pulumi.InputType['ClusterLoggingInfoArgs']]] = None,
                 number_of_broker_nodes: Optional[pulumi.Input[float]] = None,
                 open_monitoring: Optional[pulumi.Input[pulumi.InputType['ClusterOpenMonitoringArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages AWS Managed Streaming for Kafka cluster

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        vpc = aws.ec2.Vpc("vpc", cidr_block="192.168.0.0/22")
        azs = aws.get_availability_zones(state="available")
        subnet_az1 = aws.ec2.Subnet("subnetAz1",
            availability_zone=azs.names[0],
            cidr_block="192.168.0.0/24",
            vpc_id=vpc.id)
        subnet_az2 = aws.ec2.Subnet("subnetAz2",
            availability_zone=azs.names[1],
            cidr_block="192.168.1.0/24",
            vpc_id=vpc.id)
        subnet_az3 = aws.ec2.Subnet("subnetAz3",
            availability_zone=azs.names[2],
            cidr_block="192.168.2.0/24",
            vpc_id=vpc.id)
        sg = aws.ec2.SecurityGroup("sg", vpc_id=vpc.id)
        kms = aws.kms.Key("kms", description="example")
        test = aws.cloudwatch.LogGroup("test")
        bucket = aws.s3.Bucket("bucket", acl="private")
        firehose_role = aws.iam.Role("firehoseRole", assume_role_policy=\"\"\"{
        "Version": "2012-10-17",
        "Statement": [
          {
            "Action": "sts:AssumeRole",
            "Principal": {
              "Service": "firehose.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
          }
          ]
        }
        \"\"\")
        test_stream = aws.kinesis.FirehoseDeliveryStream("testStream",
            destination="s3",
            s3_configuration={
                "role_arn": firehose_role.arn,
                "bucketArn": bucket.arn,
            },
            tags={
                "LogDeliveryEnabled": "placeholder",
            })
        example = aws.msk.Cluster("example",
            cluster_name="example",
            kafka_version="2.4.1",
            number_of_broker_nodes=3,
            broker_node_group_info={
                "instance_type": "kafka.m5.large",
                "ebsVolumeSize": 1000,
                "clientSubnets": [
                    subnet_az1.id,
                    subnet_az2.id,
                    subnet_az3.id,
                ],
                "security_groups": [sg.id],
            },
            encryption_info={
                "encryptionAtRestKmsKeyArn": kms.arn,
            },
            open_monitoring={
                "prometheus": {
                    "jmxExporter": {
                        "enabledInBroker": True,
                    },
                    "nodeExporter": {
                        "enabledInBroker": True,
                    },
                },
            },
            logging_info={
                "brokerLogs": {
                    "cloudwatchLogs": {
                        "enabled": True,
                        "log_group": test.name,
                    },
                    "firehose": {
                        "enabled": True,
                        "deliveryStream": test_stream.name,
                    },
                    "s3": {
                        "enabled": True,
                        "bucket": bucket.id,
                        "prefix": "logs/msk-",
                    },
                },
            },
            tags={
                "foo": "bar",
            })
        pulumi.export("zookeeperConnectString", example.zookeeper_connect_string)
        pulumi.export("bootstrapBrokersTls", example.bootstrap_brokers_tls)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ClusterBrokerNodeGroupInfoArgs']] broker_node_group_info: Configuration block for the broker nodes of the Kafka cluster.
        :param pulumi.Input[pulumi.InputType['ClusterClientAuthenticationArgs']] client_authentication: Configuration block for specifying a client authentication. See below.
        :param pulumi.Input[str] cluster_name: Name of the MSK cluster.
        :param pulumi.Input[pulumi.InputType['ClusterConfigurationInfoArgs']] configuration_info: Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        :param pulumi.Input[pulumi.InputType['ClusterEncryptionInfoArgs']] encryption_info: Configuration block for specifying encryption. See below.
        :param pulumi.Input[str] enhanced_monitoring: Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        :param pulumi.Input[str] kafka_version: Specify the desired Kafka software version.
        :param pulumi.Input[pulumi.InputType['ClusterLoggingInfoArgs']] logging_info: Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
        :param pulumi.Input[float] number_of_broker_nodes: The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        :param pulumi.Input[pulumi.InputType['ClusterOpenMonitoringArgs']] open_monitoring: Configuration block for JMX and Node monitoring for the MSK cluster. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if broker_node_group_info is None:
                raise TypeError("Missing required property 'broker_node_group_info'")
            __props__['broker_node_group_info'] = broker_node_group_info
            __props__['client_authentication'] = client_authentication
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['configuration_info'] = configuration_info
            __props__['encryption_info'] = encryption_info
            __props__['enhanced_monitoring'] = enhanced_monitoring
            if kafka_version is None:
                raise TypeError("Missing required property 'kafka_version'")
            __props__['kafka_version'] = kafka_version
            __props__['logging_info'] = logging_info
            if number_of_broker_nodes is None:
                raise TypeError("Missing required property 'number_of_broker_nodes'")
            __props__['number_of_broker_nodes'] = number_of_broker_nodes
            __props__['open_monitoring'] = open_monitoring
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['bootstrap_brokers'] = None
            __props__['bootstrap_brokers_tls'] = None
            __props__['current_version'] = None
            __props__['zookeeper_connect_string'] = None
        super(Cluster, __self__).__init__(
            'aws:msk/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bootstrap_brokers: Optional[pulumi.Input[str]] = None,
            bootstrap_brokers_tls: Optional[pulumi.Input[str]] = None,
            broker_node_group_info: Optional[pulumi.Input[pulumi.InputType['ClusterBrokerNodeGroupInfoArgs']]] = None,
            client_authentication: Optional[pulumi.Input[pulumi.InputType['ClusterClientAuthenticationArgs']]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            configuration_info: Optional[pulumi.Input[pulumi.InputType['ClusterConfigurationInfoArgs']]] = None,
            current_version: Optional[pulumi.Input[str]] = None,
            encryption_info: Optional[pulumi.Input[pulumi.InputType['ClusterEncryptionInfoArgs']]] = None,
            enhanced_monitoring: Optional[pulumi.Input[str]] = None,
            kafka_version: Optional[pulumi.Input[str]] = None,
            logging_info: Optional[pulumi.Input[pulumi.InputType['ClusterLoggingInfoArgs']]] = None,
            number_of_broker_nodes: Optional[pulumi.Input[float]] = None,
            open_monitoring: Optional[pulumi.Input[pulumi.InputType['ClusterOpenMonitoringArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            zookeeper_connect_string: Optional[pulumi.Input[str]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        :param pulumi.Input[str] bootstrap_brokers: A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
        :param pulumi.Input[str] bootstrap_brokers_tls: A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
        :param pulumi.Input[pulumi.InputType['ClusterBrokerNodeGroupInfoArgs']] broker_node_group_info: Configuration block for the broker nodes of the Kafka cluster.
        :param pulumi.Input[pulumi.InputType['ClusterClientAuthenticationArgs']] client_authentication: Configuration block for specifying a client authentication. See below.
        :param pulumi.Input[str] cluster_name: Name of the MSK cluster.
        :param pulumi.Input[pulumi.InputType['ClusterConfigurationInfoArgs']] configuration_info: Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        :param pulumi.Input[str] current_version: Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
               * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
        :param pulumi.Input[pulumi.InputType['ClusterEncryptionInfoArgs']] encryption_info: Configuration block for specifying encryption. See below.
        :param pulumi.Input[str] enhanced_monitoring: Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        :param pulumi.Input[str] kafka_version: Specify the desired Kafka software version.
        :param pulumi.Input[pulumi.InputType['ClusterLoggingInfoArgs']] logging_info: Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
        :param pulumi.Input[float] number_of_broker_nodes: The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        :param pulumi.Input[pulumi.InputType['ClusterOpenMonitoringArgs']] open_monitoring: Configuration block for JMX and Node monitoring for the MSK cluster. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource
        :param pulumi.Input[str] zookeeper_connect_string: A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["bootstrap_brokers"] = bootstrap_brokers
        __props__["bootstrap_brokers_tls"] = bootstrap_brokers_tls
        __props__["broker_node_group_info"] = broker_node_group_info
        __props__["client_authentication"] = client_authentication
        __props__["cluster_name"] = cluster_name
        __props__["configuration_info"] = configuration_info
        __props__["current_version"] = current_version
        __props__["encryption_info"] = encryption_info
        __props__["enhanced_monitoring"] = enhanced_monitoring
        __props__["kafka_version"] = kafka_version
        __props__["logging_info"] = logging_info
        __props__["number_of_broker_nodes"] = number_of_broker_nodes
        __props__["open_monitoring"] = open_monitoring
        __props__["tags"] = tags
        __props__["zookeeper_connect_string"] = zookeeper_connect_string
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        """
        ...

    @property
    @pulumi.getter(name="bootstrapBrokers")
    def bootstrap_brokers(self) -> str:
        """
        A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
        """
        ...

    @property
    @pulumi.getter(name="bootstrapBrokersTls")
    def bootstrap_brokers_tls(self) -> str:
        """
        A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
        """
        ...

    @property
    @pulumi.getter(name="brokerNodeGroupInfo")
    def broker_node_group_info(self) -> 'outputs.ClusterBrokerNodeGroupInfo':
        """
        Configuration block for the broker nodes of the Kafka cluster.
        """
        ...

    @property
    @pulumi.getter(name="clientAuthentication")
    def client_authentication(self) -> Optional['outputs.ClusterClientAuthentication']:
        """
        Configuration block for specifying a client authentication. See below.
        """
        ...

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        """
        Name of the MSK cluster.
        """
        ...

    @property
    @pulumi.getter(name="configurationInfo")
    def configuration_info(self) -> Optional['outputs.ClusterConfigurationInfo']:
        """
        Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        """
        ...

    @property
    @pulumi.getter(name="currentVersion")
    def current_version(self) -> str:
        """
        Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
        * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
        """
        ...

    @property
    @pulumi.getter(name="encryptionInfo")
    def encryption_info(self) -> Optional['outputs.ClusterEncryptionInfo']:
        """
        Configuration block for specifying encryption. See below.
        """
        ...

    @property
    @pulumi.getter(name="enhancedMonitoring")
    def enhanced_monitoring(self) -> Optional[str]:
        """
        Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        """
        ...

    @property
    @pulumi.getter(name="kafkaVersion")
    def kafka_version(self) -> str:
        """
        Specify the desired Kafka software version.
        """
        ...

    @property
    @pulumi.getter(name="loggingInfo")
    def logging_info(self) -> Optional['outputs.ClusterLoggingInfo']:
        """
        Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
        """
        ...

    @property
    @pulumi.getter(name="numberOfBrokerNodes")
    def number_of_broker_nodes(self) -> float:
        """
        The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        """
        ...

    @property
    @pulumi.getter(name="openMonitoring")
    def open_monitoring(self) -> Optional['outputs.ClusterOpenMonitoring']:
        """
        Configuration block for JMX and Node monitoring for the MSK cluster. See below.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource
        """
        ...

    @property
    @pulumi.getter(name="zookeeperConnectString")
    def zookeeper_connect_string(self) -> str:
        """
        A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster.
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

