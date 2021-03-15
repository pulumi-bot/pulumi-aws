# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ProxyTargetArgs', 'ProxyTarget']

@pulumi.input_type
class ProxyTargetArgs:
    def __init__(__self__, *,
                 db_proxy_name: pulumi.Input[str],
                 target_group_name: pulumi.Input[str],
                 db_cluster_identifier: Optional[pulumi.Input[str]] = None,
                 db_instance_identifier: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProxyTarget resource.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy.
        :param pulumi.Input[str] target_group_name: The name of the target group.
        :param pulumi.Input[str] db_cluster_identifier: DB cluster identifier.
        :param pulumi.Input[str] db_instance_identifier: DB instance identifier.
        """
        pulumi.set(__self__, "db_proxy_name", db_proxy_name)
        pulumi.set(__self__, "target_group_name", target_group_name)
        if db_cluster_identifier is not None:
            pulumi.set(__self__, "db_cluster_identifier", db_cluster_identifier)
        if db_instance_identifier is not None:
            pulumi.set(__self__, "db_instance_identifier", db_instance_identifier)

    @property
    @pulumi.getter(name="dbProxyName")
    def db_proxy_name(self) -> pulumi.Input[str]:
        """
        The name of the DB proxy.
        """
        return pulumi.get(self, "db_proxy_name")

    @db_proxy_name.setter
    def db_proxy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_proxy_name", value)

    @property
    @pulumi.getter(name="targetGroupName")
    def target_group_name(self) -> pulumi.Input[str]:
        """
        The name of the target group.
        """
        return pulumi.get(self, "target_group_name")

    @target_group_name.setter
    def target_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_group_name", value)

    @property
    @pulumi.getter(name="dbClusterIdentifier")
    def db_cluster_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        DB cluster identifier.
        """
        return pulumi.get(self, "db_cluster_identifier")

    @db_cluster_identifier.setter
    def db_cluster_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_cluster_identifier", value)

    @property
    @pulumi.getter(name="dbInstanceIdentifier")
    def db_instance_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        DB instance identifier.
        """
        return pulumi.get(self, "db_instance_identifier")

    @db_instance_identifier.setter
    def db_instance_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_identifier", value)


class ProxyTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProxyTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an RDS DB proxy target resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_proxy = aws.rds.Proxy("exampleProxy",
            debug_logging=False,
            engine_family="MYSQL",
            idle_client_timeout=1800,
            require_tls=True,
            role_arn=aws_iam_role["example"]["arn"],
            vpc_security_group_ids=[aws_security_group["example"]["id"]],
            vpc_subnet_ids=[aws_subnet["example"]["id"]],
            auths=[aws.rds.ProxyAuthArgs(
                auth_scheme="SECRETS",
                description="example",
                iam_auth="DISABLED",
                secret_arn=aws_secretsmanager_secret["example"]["arn"],
            )],
            tags={
                "Name": "example",
                "Key": "value",
            })
        example_proxy_default_target_group = aws.rds.ProxyDefaultTargetGroup("exampleProxyDefaultTargetGroup",
            db_proxy_name=example_proxy.name,
            connection_pool_config=aws.rds.ProxyDefaultTargetGroupConnectionPoolConfigArgs(
                connection_borrow_timeout=120,
                init_query="SET x=1, y=2",
                max_connections_percent=100,
                max_idle_connections_percent=50,
                session_pinning_filters=["EXCLUDE_VARIABLE_SETS"],
            ))
        example_proxy_target = aws.rds.ProxyTarget("exampleProxyTarget",
            db_instance_identifier=aws_db_instance["example"]["id"],
            db_proxy_name=example_proxy.name,
            target_group_name=example_proxy_default_target_group.name)
        ```

        ## Import

        RDS DB Proxy Targets can be imported using the `db_proxy_name`, `target_group_name`, target type (e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`), and resource identifier separated by forward slashes (`/`), e.g. Instances

        ```sh
         $ pulumi import aws:rds/proxyTarget:ProxyTarget example example-proxy/default/RDS_INSTANCE/example-instance
        ```

         Provisioned Clusters

        ```sh
         $ pulumi import aws:rds/proxyTarget:ProxyTarget example example-proxy/default/TRACKED_CLUSTER/example-cluster
        ```

        :param str resource_name: The name of the resource.
        :param ProxyTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_cluster_identifier: Optional[pulumi.Input[str]] = None,
                 db_instance_identifier: Optional[pulumi.Input[str]] = None,
                 db_proxy_name: Optional[pulumi.Input[str]] = None,
                 target_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an RDS DB proxy target resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_proxy = aws.rds.Proxy("exampleProxy",
            debug_logging=False,
            engine_family="MYSQL",
            idle_client_timeout=1800,
            require_tls=True,
            role_arn=aws_iam_role["example"]["arn"],
            vpc_security_group_ids=[aws_security_group["example"]["id"]],
            vpc_subnet_ids=[aws_subnet["example"]["id"]],
            auths=[aws.rds.ProxyAuthArgs(
                auth_scheme="SECRETS",
                description="example",
                iam_auth="DISABLED",
                secret_arn=aws_secretsmanager_secret["example"]["arn"],
            )],
            tags={
                "Name": "example",
                "Key": "value",
            })
        example_proxy_default_target_group = aws.rds.ProxyDefaultTargetGroup("exampleProxyDefaultTargetGroup",
            db_proxy_name=example_proxy.name,
            connection_pool_config=aws.rds.ProxyDefaultTargetGroupConnectionPoolConfigArgs(
                connection_borrow_timeout=120,
                init_query="SET x=1, y=2",
                max_connections_percent=100,
                max_idle_connections_percent=50,
                session_pinning_filters=["EXCLUDE_VARIABLE_SETS"],
            ))
        example_proxy_target = aws.rds.ProxyTarget("exampleProxyTarget",
            db_instance_identifier=aws_db_instance["example"]["id"],
            db_proxy_name=example_proxy.name,
            target_group_name=example_proxy_default_target_group.name)
        ```

        ## Import

        RDS DB Proxy Targets can be imported using the `db_proxy_name`, `target_group_name`, target type (e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`), and resource identifier separated by forward slashes (`/`), e.g. Instances

        ```sh
         $ pulumi import aws:rds/proxyTarget:ProxyTarget example example-proxy/default/RDS_INSTANCE/example-instance
        ```

         Provisioned Clusters

        ```sh
         $ pulumi import aws:rds/proxyTarget:ProxyTarget example example-proxy/default/TRACKED_CLUSTER/example-cluster
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_cluster_identifier: DB cluster identifier.
        :param pulumi.Input[str] db_instance_identifier: DB instance identifier.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy.
        :param pulumi.Input[str] target_group_name: The name of the target group.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProxyTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_cluster_identifier: Optional[pulumi.Input[str]] = None,
                 db_instance_identifier: Optional[pulumi.Input[str]] = None,
                 db_proxy_name: Optional[pulumi.Input[str]] = None,
                 target_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['db_cluster_identifier'] = db_cluster_identifier
            __props__['db_instance_identifier'] = db_instance_identifier
            if db_proxy_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_proxy_name'")
            __props__['db_proxy_name'] = db_proxy_name
            if target_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'target_group_name'")
            __props__['target_group_name'] = target_group_name
            __props__['endpoint'] = None
            __props__['port'] = None
            __props__['rds_resource_id'] = None
            __props__['target_arn'] = None
            __props__['tracked_cluster_id'] = None
            __props__['type'] = None
        super(ProxyTarget, __self__).__init__(
            'aws:rds/proxyTarget:ProxyTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_cluster_identifier: Optional[pulumi.Input[str]] = None,
            db_instance_identifier: Optional[pulumi.Input[str]] = None,
            db_proxy_name: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            rds_resource_id: Optional[pulumi.Input[str]] = None,
            target_arn: Optional[pulumi.Input[str]] = None,
            target_group_name: Optional[pulumi.Input[str]] = None,
            tracked_cluster_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ProxyTarget':
        """
        Get an existing ProxyTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_cluster_identifier: DB cluster identifier.
        :param pulumi.Input[str] db_instance_identifier: DB instance identifier.
        :param pulumi.Input[str] db_proxy_name: The name of the DB proxy.
        :param pulumi.Input[str] endpoint: Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type.
        :param pulumi.Input[int] port: Port for the target RDS DB Instance or Aurora DB Cluster.
        :param pulumi.Input[str] rds_resource_id: Identifier representing the DB Instance or DB Cluster target.
        :param pulumi.Input[str] target_arn: Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API.
        :param pulumi.Input[str] target_group_name: The name of the target group.
        :param pulumi.Input[str] tracked_cluster_id: DB Cluster identifier for the DB Instance target. Not returned unless manually importing an `RDS_INSTANCE` target that is part of a DB Cluster.
        :param pulumi.Input[str] type: Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["db_cluster_identifier"] = db_cluster_identifier
        __props__["db_instance_identifier"] = db_instance_identifier
        __props__["db_proxy_name"] = db_proxy_name
        __props__["endpoint"] = endpoint
        __props__["port"] = port
        __props__["rds_resource_id"] = rds_resource_id
        __props__["target_arn"] = target_arn
        __props__["target_group_name"] = target_group_name
        __props__["tracked_cluster_id"] = tracked_cluster_id
        __props__["type"] = type
        return ProxyTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbClusterIdentifier")
    def db_cluster_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        DB cluster identifier.
        """
        return pulumi.get(self, "db_cluster_identifier")

    @property
    @pulumi.getter(name="dbInstanceIdentifier")
    def db_instance_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        DB instance identifier.
        """
        return pulumi.get(self, "db_instance_identifier")

    @property
    @pulumi.getter(name="dbProxyName")
    def db_proxy_name(self) -> pulumi.Output[str]:
        """
        The name of the DB proxy.
        """
        return pulumi.get(self, "db_proxy_name")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        Hostname for the target RDS DB Instance. Only returned for `RDS_INSTANCE` type.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port for the target RDS DB Instance or Aurora DB Cluster.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="rdsResourceId")
    def rds_resource_id(self) -> pulumi.Output[str]:
        """
        Identifier representing the DB Instance or DB Cluster target.
        """
        return pulumi.get(self, "rds_resource_id")

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) for the DB instance or DB cluster. Currently not returned by the RDS API.
        """
        return pulumi.get(self, "target_arn")

    @property
    @pulumi.getter(name="targetGroupName")
    def target_group_name(self) -> pulumi.Output[str]:
        """
        The name of the target group.
        """
        return pulumi.get(self, "target_group_name")

    @property
    @pulumi.getter(name="trackedClusterId")
    def tracked_cluster_id(self) -> pulumi.Output[str]:
        """
        DB Cluster identifier for the DB Instance target. Not returned unless manually importing an `RDS_INSTANCE` target that is part of a DB Cluster.
        """
        return pulumi.get(self, "tracked_cluster_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of target. e.g. `RDS_INSTANCE` or `TRACKED_CLUSTER`
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

