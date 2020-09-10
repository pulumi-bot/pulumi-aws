# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ClusterInstance']


class ClusterInstance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_immediately: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 neptune_parameter_group_name: Optional[pulumi.Input[str]] = None,
                 neptune_subnet_group_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 preferred_backup_window: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 promotion_tier: Optional[pulumi.Input[float]] = None,
                 publicly_accessible: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ClusterInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            __props__['apply_immediately'] = apply_immediately
            __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade
            __props__['availability_zone'] = availability_zone
            if cluster_identifier is None:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__['cluster_identifier'] = cluster_identifier
            __props__['engine'] = engine
            __props__['engine_version'] = engine_version
            __props__['identifier'] = identifier
            __props__['identifier_prefix'] = identifier_prefix
            if instance_class is None:
                raise TypeError("Missing required property 'instance_class'")
            __props__['instance_class'] = instance_class
            __props__['neptune_parameter_group_name'] = neptune_parameter_group_name
            __props__['neptune_subnet_group_name'] = neptune_subnet_group_name
            __props__['port'] = port
            __props__['preferred_backup_window'] = preferred_backup_window
            __props__['preferred_maintenance_window'] = preferred_maintenance_window
            __props__['promotion_tier'] = promotion_tier
            __props__['publicly_accessible'] = publicly_accessible
            __props__['tags'] = tags
            __props__['address'] = None
            __props__['arn'] = None
            __props__['dbi_resource_id'] = None
            __props__['endpoint'] = None
            __props__['kms_key_arn'] = None
            __props__['storage_encrypted'] = None
            __props__['writer'] = None
        super(ClusterInstance, __self__).__init__(
            'aws:neptune/clusterInstance:ClusterInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            apply_immediately: Optional[pulumi.Input[bool]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            cluster_identifier: Optional[pulumi.Input[str]] = None,
            dbi_resource_id: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            identifier_prefix: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            kms_key_arn: Optional[pulumi.Input[str]] = None,
            neptune_parameter_group_name: Optional[pulumi.Input[str]] = None,
            neptune_subnet_group_name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[float]] = None,
            preferred_backup_window: Optional[pulumi.Input[str]] = None,
            preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
            promotion_tier: Optional[pulumi.Input[float]] = None,
            publicly_accessible: Optional[pulumi.Input[bool]] = None,
            storage_encrypted: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            writer: Optional[pulumi.Input[bool]] = None) -> 'ClusterInstance':
        """
        Get an existing ClusterInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["apply_immediately"] = apply_immediately
        __props__["arn"] = arn
        __props__["auto_minor_version_upgrade"] = auto_minor_version_upgrade
        __props__["availability_zone"] = availability_zone
        __props__["cluster_identifier"] = cluster_identifier
        __props__["dbi_resource_id"] = dbi_resource_id
        __props__["endpoint"] = endpoint
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["identifier"] = identifier
        __props__["identifier_prefix"] = identifier_prefix
        __props__["instance_class"] = instance_class
        __props__["kms_key_arn"] = kms_key_arn
        __props__["neptune_parameter_group_name"] = neptune_parameter_group_name
        __props__["neptune_subnet_group_name"] = neptune_subnet_group_name
        __props__["port"] = port
        __props__["preferred_backup_window"] = preferred_backup_window
        __props__["preferred_maintenance_window"] = preferred_maintenance_window
        __props__["promotion_tier"] = promotion_tier
        __props__["publicly_accessible"] = publicly_accessible
        __props__["storage_encrypted"] = storage_encrypted
        __props__["tags"] = tags
        __props__["writer"] = writer
        return ClusterInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="applyImmediately")
    def apply_immediately(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "apply_immediately")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "auto_minor_version_upgrade")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter(name="dbiResourceId")
    def dbi_resource_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dbi_resource_id")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="identifierPrefix")
    def identifier_prefix(self) -> pulumi.Output[str]:
        return pulumi.get(self, "identifier_prefix")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="neptuneParameterGroupName")
    def neptune_parameter_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "neptune_parameter_group_name")

    @property
    @pulumi.getter(name="neptuneSubnetGroupName")
    def neptune_subnet_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "neptune_subnet_group_name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> pulumi.Output[str]:
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> pulumi.Output[str]:
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter(name="promotionTier")
    def promotion_tier(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "promotion_tier")

    @property
    @pulumi.getter(name="publiclyAccessible")
    def publicly_accessible(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "publicly_accessible")

    @property
    @pulumi.getter(name="storageEncrypted")
    def storage_encrypted(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "storage_encrypted")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def writer(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "writer")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

