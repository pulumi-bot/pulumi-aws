# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
    """
    bootstrap_brokers: pulumi.Output[str]
    """
    A comma separated list of one or more hostname:port pairs of kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`.
    """
    bootstrap_brokers_tls: pulumi.Output[str]
    """
    A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `client_broker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`.
    """
    broker_node_group_info: pulumi.Output[dict]
    """
    Configuration block for the broker nodes of the Kafka cluster.
    """
    client_authentication: pulumi.Output[dict]
    """
    Configuration block for specifying a client authentication. See below.
    """
    cluster_name: pulumi.Output[str]
    """
    Name of the MSK cluster.
    """
    configuration_info: pulumi.Output[dict]
    """
    Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
    """
    current_version: pulumi.Output[str]
    """
    Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
    * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
    """
    encryption_info: pulumi.Output[dict]
    """
    Configuration block for specifying encryption. See below.
    """
    enhanced_monitoring: pulumi.Output[str]
    """
    Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
    """
    kafka_version: pulumi.Output[str]
    """
    Specify the desired Kafka software version.
    """
    number_of_broker_nodes: pulumi.Output[float]
    """
    The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource
    """
    zookeeper_connect_string: pulumi.Output[str]
    """
    A comma separated list of one or more IP:port pairs to use to connect to the Apache Zookeeper cluster.
    """
    def __init__(__self__, resource_name, opts=None, broker_node_group_info=None, client_authentication=None, cluster_name=None, configuration_info=None, encryption_info=None, enhanced_monitoring=None, kafka_version=None, number_of_broker_nodes=None, tags=None, __name__=None, __opts__=None):
        """
        Manages AWS Managed Streaming for Kafka cluster
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] broker_node_group_info: Configuration block for the broker nodes of the Kafka cluster.
        :param pulumi.Input[dict] client_authentication: Configuration block for specifying a client authentication. See below.
        :param pulumi.Input[str] cluster_name: Name of the MSK cluster.
        :param pulumi.Input[dict] configuration_info: Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
        :param pulumi.Input[dict] encryption_info: Configuration block for specifying encryption. See below.
        :param pulumi.Input[str] enhanced_monitoring: Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
        :param pulumi.Input[str] kafka_version: Specify the desired Kafka software version.
        :param pulumi.Input[float] number_of_broker_nodes: The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/msk_cluster.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

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

        if number_of_broker_nodes is None:
            raise TypeError("Missing required property 'number_of_broker_nodes'")
        __props__['number_of_broker_nodes'] = number_of_broker_nodes

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['bootstrap_brokers'] = None
        __props__['bootstrap_brokers_tls'] = None
        __props__['current_version'] = None
        __props__['zookeeper_connect_string'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Cluster, __self__).__init__(
            'aws:msk/cluster:Cluster',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

