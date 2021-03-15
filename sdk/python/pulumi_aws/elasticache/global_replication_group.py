# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['GlobalReplicationGroupArgs', 'GlobalReplicationGroup']

@pulumi.input_type
class GlobalReplicationGroupArgs:
    def __init__(__self__, *,
                 global_replication_group_id_suffix: pulumi.Input[str],
                 primary_replication_group_id: pulumi.Input[str],
                 global_replication_group_description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GlobalReplicationGroup resource.
        :param pulumi.Input[str] global_replication_group_id_suffix: The suffix name of a Global Datastore. If `global_replication_group_id_suffix` is changed, creates a new resource.
        :param pulumi.Input[str] primary_replication_group_id: The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primary_replication_group_id` is changed, creates a new resource.
        :param pulumi.Input[str] global_replication_group_description: A user-created description for the global replication group.
        """
        pulumi.set(__self__, "global_replication_group_id_suffix", global_replication_group_id_suffix)
        pulumi.set(__self__, "primary_replication_group_id", primary_replication_group_id)
        if global_replication_group_description is not None:
            pulumi.set(__self__, "global_replication_group_description", global_replication_group_description)

    @property
    @pulumi.getter(name="globalReplicationGroupIdSuffix")
    def global_replication_group_id_suffix(self) -> pulumi.Input[str]:
        """
        The suffix name of a Global Datastore. If `global_replication_group_id_suffix` is changed, creates a new resource.
        """
        return pulumi.get(self, "global_replication_group_id_suffix")

    @global_replication_group_id_suffix.setter
    def global_replication_group_id_suffix(self, value: pulumi.Input[str]):
        pulumi.set(self, "global_replication_group_id_suffix", value)

    @property
    @pulumi.getter(name="primaryReplicationGroupId")
    def primary_replication_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primary_replication_group_id` is changed, creates a new resource.
        """
        return pulumi.get(self, "primary_replication_group_id")

    @primary_replication_group_id.setter
    def primary_replication_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "primary_replication_group_id", value)

    @property
    @pulumi.getter(name="globalReplicationGroupDescription")
    def global_replication_group_description(self) -> Optional[pulumi.Input[str]]:
        """
        A user-created description for the global replication group.
        """
        return pulumi.get(self, "global_replication_group_description")

    @global_replication_group_description.setter
    def global_replication_group_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_replication_group_description", value)


class GlobalReplicationGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalReplicationGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        ElastiCache Global Replication Groups can be imported using the `global_replication_group_id`, e.g.

        ```sh
         $ pulumi import aws:elasticache/globalReplicationGroup:GlobalReplicationGroup my_global_replication_group okuqm-global-replication-group-1
        ```

        :param str resource_name: The name of the resource.
        :param GlobalReplicationGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 global_replication_group_description: Optional[pulumi.Input[str]] = None,
                 global_replication_group_id_suffix: Optional[pulumi.Input[str]] = None,
                 primary_replication_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        ElastiCache Global Replication Groups can be imported using the `global_replication_group_id`, e.g.

        ```sh
         $ pulumi import aws:elasticache/globalReplicationGroup:GlobalReplicationGroup my_global_replication_group okuqm-global-replication-group-1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] global_replication_group_description: A user-created description for the global replication group.
        :param pulumi.Input[str] global_replication_group_id_suffix: The suffix name of a Global Datastore. If `global_replication_group_id_suffix` is changed, creates a new resource.
        :param pulumi.Input[str] primary_replication_group_id: The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primary_replication_group_id` is changed, creates a new resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalReplicationGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 global_replication_group_description: Optional[pulumi.Input[str]] = None,
                 global_replication_group_id_suffix: Optional[pulumi.Input[str]] = None,
                 primary_replication_group_id: Optional[pulumi.Input[str]] = None,
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

            __props__['global_replication_group_description'] = global_replication_group_description
            if global_replication_group_id_suffix is None and not opts.urn:
                raise TypeError("Missing required property 'global_replication_group_id_suffix'")
            __props__['global_replication_group_id_suffix'] = global_replication_group_id_suffix
            if primary_replication_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'primary_replication_group_id'")
            __props__['primary_replication_group_id'] = primary_replication_group_id
            __props__['actual_engine_version'] = None
            __props__['arn'] = None
            __props__['at_rest_encryption_enabled'] = None
            __props__['auth_token_enabled'] = None
            __props__['cache_node_type'] = None
            __props__['cluster_enabled'] = None
            __props__['engine'] = None
            __props__['global_replication_group_id'] = None
            __props__['transit_encryption_enabled'] = None
        super(GlobalReplicationGroup, __self__).__init__(
            'aws:elasticache/globalReplicationGroup:GlobalReplicationGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actual_engine_version: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            at_rest_encryption_enabled: Optional[pulumi.Input[bool]] = None,
            auth_token_enabled: Optional[pulumi.Input[bool]] = None,
            cache_node_type: Optional[pulumi.Input[str]] = None,
            cluster_enabled: Optional[pulumi.Input[bool]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            global_replication_group_description: Optional[pulumi.Input[str]] = None,
            global_replication_group_id: Optional[pulumi.Input[str]] = None,
            global_replication_group_id_suffix: Optional[pulumi.Input[str]] = None,
            primary_replication_group_id: Optional[pulumi.Input[str]] = None,
            transit_encryption_enabled: Optional[pulumi.Input[bool]] = None) -> 'GlobalReplicationGroup':
        """
        Get an existing GlobalReplicationGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] actual_engine_version: The full version number of the cache engine running on the members of this global replication group.
        :param pulumi.Input[str] arn: The ARN of the ElastiCache Global Replication Group.
        :param pulumi.Input[bool] at_rest_encryption_enabled: A flag that indicate whether the encryption at rest is enabled.
        :param pulumi.Input[bool] auth_token_enabled: A flag that indicate whether AuthToken (password) is enabled.
        :param pulumi.Input[str] cache_node_type: The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
        :param pulumi.Input[bool] cluster_enabled: Indicates whether the Global Datastore is cluster enabled.
        :param pulumi.Input[str] engine: The name of the cache engine to be used for the clusters in this global replication group.
        :param pulumi.Input[str] global_replication_group_description: A user-created description for the global replication group.
        :param pulumi.Input[str] global_replication_group_id: The full ID of the global replication group.
        :param pulumi.Input[str] global_replication_group_id_suffix: The suffix name of a Global Datastore. If `global_replication_group_id_suffix` is changed, creates a new resource.
        :param pulumi.Input[str] primary_replication_group_id: The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primary_replication_group_id` is changed, creates a new resource.
        :param pulumi.Input[bool] transit_encryption_enabled: A flag that indicates whether the encryption in transit is enabled.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actual_engine_version"] = actual_engine_version
        __props__["arn"] = arn
        __props__["at_rest_encryption_enabled"] = at_rest_encryption_enabled
        __props__["auth_token_enabled"] = auth_token_enabled
        __props__["cache_node_type"] = cache_node_type
        __props__["cluster_enabled"] = cluster_enabled
        __props__["engine"] = engine
        __props__["global_replication_group_description"] = global_replication_group_description
        __props__["global_replication_group_id"] = global_replication_group_id
        __props__["global_replication_group_id_suffix"] = global_replication_group_id_suffix
        __props__["primary_replication_group_id"] = primary_replication_group_id
        __props__["transit_encryption_enabled"] = transit_encryption_enabled
        return GlobalReplicationGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actualEngineVersion")
    def actual_engine_version(self) -> pulumi.Output[str]:
        """
        The full version number of the cache engine running on the members of this global replication group.
        """
        return pulumi.get(self, "actual_engine_version")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the ElastiCache Global Replication Group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="atRestEncryptionEnabled")
    def at_rest_encryption_enabled(self) -> pulumi.Output[bool]:
        """
        A flag that indicate whether the encryption at rest is enabled.
        """
        return pulumi.get(self, "at_rest_encryption_enabled")

    @property
    @pulumi.getter(name="authTokenEnabled")
    def auth_token_enabled(self) -> pulumi.Output[bool]:
        """
        A flag that indicate whether AuthToken (password) is enabled.
        """
        return pulumi.get(self, "auth_token_enabled")

    @property
    @pulumi.getter(name="cacheNodeType")
    def cache_node_type(self) -> pulumi.Output[str]:
        """
        The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
        """
        return pulumi.get(self, "cache_node_type")

    @property
    @pulumi.getter(name="clusterEnabled")
    def cluster_enabled(self) -> pulumi.Output[bool]:
        """
        Indicates whether the Global Datastore is cluster enabled.
        """
        return pulumi.get(self, "cluster_enabled")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The name of the cache engine to be used for the clusters in this global replication group.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="globalReplicationGroupDescription")
    def global_replication_group_description(self) -> pulumi.Output[Optional[str]]:
        """
        A user-created description for the global replication group.
        """
        return pulumi.get(self, "global_replication_group_description")

    @property
    @pulumi.getter(name="globalReplicationGroupId")
    def global_replication_group_id(self) -> pulumi.Output[str]:
        """
        The full ID of the global replication group.
        """
        return pulumi.get(self, "global_replication_group_id")

    @property
    @pulumi.getter(name="globalReplicationGroupIdSuffix")
    def global_replication_group_id_suffix(self) -> pulumi.Output[str]:
        """
        The suffix name of a Global Datastore. If `global_replication_group_id_suffix` is changed, creates a new resource.
        """
        return pulumi.get(self, "global_replication_group_id_suffix")

    @property
    @pulumi.getter(name="primaryReplicationGroupId")
    def primary_replication_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primary_replication_group_id` is changed, creates a new resource.
        """
        return pulumi.get(self, "primary_replication_group_id")

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> pulumi.Output[bool]:
        """
        A flag that indicates whether the encryption in transit is enabled.
        """
        return pulumi.get(self, "transit_encryption_enabled")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

