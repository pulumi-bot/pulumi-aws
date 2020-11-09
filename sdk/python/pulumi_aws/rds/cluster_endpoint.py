# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ClusterEndpoint']


class ClusterEndpoint(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_endpoint_identifier: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 custom_endpoint_type: Optional[pulumi.Input[str]] = None,
                 excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an RDS Aurora Cluster Endpoint.
        You can refer to the [User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html#Aurora.Endpoints.Cluster).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.rds.Cluster("default",
            cluster_identifier="aurora-cluster-demo",
            availability_zones=[
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            ],
            database_name="mydb",
            master_username="foo",
            master_password="bar",
            backup_retention_period=5,
            preferred_backup_window="07:00-09:00")
        test1 = aws.rds.ClusterInstance("test1",
            apply_immediately=True,
            cluster_identifier=default.id,
            identifier="test1",
            instance_class="db.t2.small",
            engine=default.engine,
            engine_version=default.engine_version)
        test2 = aws.rds.ClusterInstance("test2",
            apply_immediately=True,
            cluster_identifier=default.id,
            identifier="test2",
            instance_class="db.t2.small",
            engine=default.engine,
            engine_version=default.engine_version)
        test3 = aws.rds.ClusterInstance("test3",
            apply_immediately=True,
            cluster_identifier=default.id,
            identifier="test3",
            instance_class="db.t2.small",
            engine=default.engine,
            engine_version=default.engine_version)
        eligible = aws.rds.ClusterEndpoint("eligible",
            cluster_identifier=default.id,
            cluster_endpoint_identifier="reader",
            custom_endpoint_type="READER",
            excluded_members=[
                test1.id,
                test2.id,
            ])
        static = aws.rds.ClusterEndpoint("static",
            cluster_identifier=default.id,
            cluster_endpoint_identifier="static",
            custom_endpoint_type="READER",
            static_members=[
                test1.id,
                test3.id,
            ])
        ```

        ## Import

        RDS Clusters Endpoint can be imported using the `cluster_endpoint_identifier`, e.g.

        ```sh
         $ pulumi import aws:rds/clusterEndpoint:ClusterEndpoint custom_reader aurora-prod-cluster-custom-reader
        ```

         [1]https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html#Aurora.Endpoints.Cluster

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
        :param pulumi.Input[str] cluster_identifier: The cluster identifier.
        :param pulumi.Input[str] custom_endpoint_type: The type of the endpoint. One of: READER , ANY .
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] static_members: List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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

            if cluster_endpoint_identifier is None:
                raise TypeError("Missing required property 'cluster_endpoint_identifier'")
            __props__['cluster_endpoint_identifier'] = cluster_endpoint_identifier
            if cluster_identifier is None:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__['cluster_identifier'] = cluster_identifier
            if custom_endpoint_type is None:
                raise TypeError("Missing required property 'custom_endpoint_type'")
            __props__['custom_endpoint_type'] = custom_endpoint_type
            __props__['excluded_members'] = excluded_members
            __props__['static_members'] = static_members
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['endpoint'] = None
        super(ClusterEndpoint, __self__).__init__(
            'aws:rds/clusterEndpoint:ClusterEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cluster_endpoint_identifier: Optional[pulumi.Input[str]] = None,
            cluster_identifier: Optional[pulumi.Input[str]] = None,
            custom_endpoint_type: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            excluded_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            static_members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ClusterEndpoint':
        """
        Get an existing ClusterEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of cluster
        :param pulumi.Input[str] cluster_endpoint_identifier: The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
        :param pulumi.Input[str] cluster_identifier: The cluster identifier.
        :param pulumi.Input[str] custom_endpoint_type: The type of the endpoint. One of: READER , ANY .
        :param pulumi.Input[str] endpoint: A custom endpoint for the Aurora cluster
        :param pulumi.Input[Sequence[pulumi.Input[str]]] excluded_members: List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] static_members: List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["cluster_endpoint_identifier"] = cluster_endpoint_identifier
        __props__["cluster_identifier"] = cluster_identifier
        __props__["custom_endpoint_type"] = custom_endpoint_type
        __props__["endpoint"] = endpoint
        __props__["excluded_members"] = excluded_members
        __props__["static_members"] = static_members
        __props__["tags"] = tags
        return ClusterEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of cluster
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="clusterEndpointIdentifier")
    def cluster_endpoint_identifier(self) -> pulumi.Output[str]:
        """
        The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
        """
        return pulumi.get(self, "cluster_endpoint_identifier")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Output[str]:
        """
        The cluster identifier.
        """
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter(name="customEndpointType")
    def custom_endpoint_type(self) -> pulumi.Output[str]:
        """
        The type of the endpoint. One of: READER , ANY .
        """
        return pulumi.get(self, "custom_endpoint_type")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        A custom endpoint for the Aurora cluster
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="excludedMembers")
    def excluded_members(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `static_members`.
        """
        return pulumi.get(self, "excluded_members")

    @property
    @pulumi.getter(name="staticMembers")
    def static_members(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excluded_members`.
        """
        return pulumi.get(self, "static_members")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

