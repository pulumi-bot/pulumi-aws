# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Cluster(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the DAX cluster
    """
    availability_zones: pulumi.Output[list]
    """
    List of Availability Zones in which the
    nodes will be created
    """
    cluster_address: pulumi.Output[str]
    """
    The DNS name of the DAX cluster without the port appended
    """
    cluster_name: pulumi.Output[str]
    """
    Group identifier. DAX converts this name to
    lowercase
    """
    configuration_endpoint: pulumi.Output[str]
    """
    The configuration endpoint for this DAX cluster,
    consisting of a DNS name and a port number
    """
    description: pulumi.Output[str]
    """
    Description for the cluster
    """
    iam_role_arn: pulumi.Output[str]
    """
    A valid Amazon Resource Name (ARN) that identifies
    an IAM role. At runtime, DAX will assume this role and use the role's
    permissions to access DynamoDB on your behalf
    """
    maintenance_window: pulumi.Output[str]
    """
    Specifies the weekly time range for when
    maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
    (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
    `sun:05:00-sun:09:00`
    """
    node_type: pulumi.Output[str]
    """
    The compute and memory capacity of the nodes. See
    [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
    """
    nodes: pulumi.Output[list]
    """
    List of node objects including `id`, `address`, `port` and
    `availability_zone`. Referenceable e.g. as
    `${aws_dax_cluster.test.nodes.0.address}`

      * `address` (`str`)
      * `availability_zone` (`str`)
      * `id` (`str`)
      * `port` (`float`) - The port used by the configuration endpoint
    """
    notification_topic_arn: pulumi.Output[str]
    """
    An Amazon Resource Name (ARN) of an
    SNS topic to send DAX notifications to. Example:
    `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
    """
    parameter_group_name: pulumi.Output[str]
    """
    Name of the parameter group to associate
    with this DAX cluster
    """
    port: pulumi.Output[float]
    """
    The port used by the configuration endpoint
    """
    replication_factor: pulumi.Output[float]
    """
    The number of nodes in the DAX cluster. A
    replication factor of 1 will create a single-node cluster, without any read
    replicas
    """
    security_group_ids: pulumi.Output[list]
    """
    One or more VPC security groups associated
    with the cluster
    """
    server_side_encryption: pulumi.Output[dict]
    """
    Encrypt at rest options

      * `enabled` (`bool`) - Whether to enable encryption at rest. Defaults to `false`.
    """
    subnet_group_name: pulumi.Output[str]
    """
    Name of the subnet group to be used for the
    cluster
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource
    """
    def __init__(__self__, resource_name, opts=None, availability_zones=None, cluster_name=None, description=None, iam_role_arn=None, maintenance_window=None, node_type=None, notification_topic_arn=None, parameter_group_name=None, replication_factor=None, security_group_ids=None, server_side_encryption=None, subnet_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DAX Cluster resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar = aws.dax.Cluster("bar",
            cluster_name="cluster-example",
            iam_role_arn=data["aws_iam_role"]["example"]["arn"],
            node_type="dax.r4.large",
            replication_factor=1)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] availability_zones: List of Availability Zones in which the
               nodes will be created
        :param pulumi.Input[str] cluster_name: Group identifier. DAX converts this name to
               lowercase
        :param pulumi.Input[str] description: Description for the cluster
        :param pulumi.Input[str] iam_role_arn: A valid Amazon Resource Name (ARN) that identifies
               an IAM role. At runtime, DAX will assume this role and use the role's
               permissions to access DynamoDB on your behalf
        :param pulumi.Input[str] maintenance_window: Specifies the weekly time range for when
               maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
               (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
               `sun:05:00-sun:09:00`
        :param pulumi.Input[str] node_type: The compute and memory capacity of the nodes. See
               [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
        :param pulumi.Input[str] notification_topic_arn: An Amazon Resource Name (ARN) of an
               SNS topic to send DAX notifications to. Example:
               `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        :param pulumi.Input[str] parameter_group_name: Name of the parameter group to associate
               with this DAX cluster
        :param pulumi.Input[float] replication_factor: The number of nodes in the DAX cluster. A
               replication factor of 1 will create a single-node cluster, without any read
               replicas
        :param pulumi.Input[list] security_group_ids: One or more VPC security groups associated
               with the cluster
        :param pulumi.Input[dict] server_side_encryption: Encrypt at rest options
        :param pulumi.Input[str] subnet_group_name: Name of the subnet group to be used for the
               cluster
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource

        The **server_side_encryption** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Whether to enable encryption at rest. Defaults to `false`.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['availability_zones'] = availability_zones
            if cluster_name is None:
                raise TypeError("Missing required property 'cluster_name'")
            __props__['cluster_name'] = cluster_name
            __props__['description'] = description
            if iam_role_arn is None:
                raise TypeError("Missing required property 'iam_role_arn'")
            __props__['iam_role_arn'] = iam_role_arn
            __props__['maintenance_window'] = maintenance_window
            if node_type is None:
                raise TypeError("Missing required property 'node_type'")
            __props__['node_type'] = node_type
            __props__['notification_topic_arn'] = notification_topic_arn
            __props__['parameter_group_name'] = parameter_group_name
            if replication_factor is None:
                raise TypeError("Missing required property 'replication_factor'")
            __props__['replication_factor'] = replication_factor
            __props__['security_group_ids'] = security_group_ids
            __props__['server_side_encryption'] = server_side_encryption
            __props__['subnet_group_name'] = subnet_group_name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['cluster_address'] = None
            __props__['configuration_endpoint'] = None
            __props__['nodes'] = None
            __props__['port'] = None
        super(Cluster, __self__).__init__(
            'aws:dax/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, availability_zones=None, cluster_address=None, cluster_name=None, configuration_endpoint=None, description=None, iam_role_arn=None, maintenance_window=None, node_type=None, nodes=None, notification_topic_arn=None, parameter_group_name=None, port=None, replication_factor=None, security_group_ids=None, server_side_encryption=None, subnet_group_name=None, tags=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the DAX cluster
        :param pulumi.Input[list] availability_zones: List of Availability Zones in which the
               nodes will be created
        :param pulumi.Input[str] cluster_address: The DNS name of the DAX cluster without the port appended
        :param pulumi.Input[str] cluster_name: Group identifier. DAX converts this name to
               lowercase
        :param pulumi.Input[str] configuration_endpoint: The configuration endpoint for this DAX cluster,
               consisting of a DNS name and a port number
        :param pulumi.Input[str] description: Description for the cluster
        :param pulumi.Input[str] iam_role_arn: A valid Amazon Resource Name (ARN) that identifies
               an IAM role. At runtime, DAX will assume this role and use the role's
               permissions to access DynamoDB on your behalf
        :param pulumi.Input[str] maintenance_window: Specifies the weekly time range for when
               maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
               (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
               `sun:05:00-sun:09:00`
        :param pulumi.Input[str] node_type: The compute and memory capacity of the nodes. See
               [Nodes](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DAX.concepts.cluster.html#DAX.concepts.nodes) for supported node types
        :param pulumi.Input[list] nodes: List of node objects including `id`, `address`, `port` and
               `availability_zone`. Referenceable e.g. as
               `${aws_dax_cluster.test.nodes.0.address}`
        :param pulumi.Input[str] notification_topic_arn: An Amazon Resource Name (ARN) of an
               SNS topic to send DAX notifications to. Example:
               `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        :param pulumi.Input[str] parameter_group_name: Name of the parameter group to associate
               with this DAX cluster
        :param pulumi.Input[float] port: The port used by the configuration endpoint
        :param pulumi.Input[float] replication_factor: The number of nodes in the DAX cluster. A
               replication factor of 1 will create a single-node cluster, without any read
               replicas
        :param pulumi.Input[list] security_group_ids: One or more VPC security groups associated
               with the cluster
        :param pulumi.Input[dict] server_side_encryption: Encrypt at rest options
        :param pulumi.Input[str] subnet_group_name: Name of the subnet group to be used for the
               cluster
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource

        The **nodes** object supports the following:

          * `address` (`pulumi.Input[str]`)
          * `availability_zone` (`pulumi.Input[str]`)
          * `id` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`) - The port used by the configuration endpoint

        The **server_side_encryption** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Whether to enable encryption at rest. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["availability_zones"] = availability_zones
        __props__["cluster_address"] = cluster_address
        __props__["cluster_name"] = cluster_name
        __props__["configuration_endpoint"] = configuration_endpoint
        __props__["description"] = description
        __props__["iam_role_arn"] = iam_role_arn
        __props__["maintenance_window"] = maintenance_window
        __props__["node_type"] = node_type
        __props__["nodes"] = nodes
        __props__["notification_topic_arn"] = notification_topic_arn
        __props__["parameter_group_name"] = parameter_group_name
        __props__["port"] = port
        __props__["replication_factor"] = replication_factor
        __props__["security_group_ids"] = security_group_ids
        __props__["server_side_encryption"] = server_side_encryption
        __props__["subnet_group_name"] = subnet_group_name
        __props__["tags"] = tags
        return Cluster(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
