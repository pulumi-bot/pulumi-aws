# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'ClusterParameterGroupParameter',
    'ClusterS3Import',
    'ClusterScalingConfiguration',
    'GlobalClusterGlobalClusterMember',
    'InstanceS3Import',
    'OptionGroupOption',
    'OptionGroupOptionOptionSetting',
    'ParameterGroupParameter',
    'SecurityGroupIngress',
]

@pulumi.output_type
class ClusterParameterGroupParameter(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str,
                 apply_method: Optional[str] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if apply_method is not None:
            pulumi.set(__self__, "apply_method", apply_method)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="applyMethod")
    def apply_method(self) -> Optional[str]:
        return pulumi.get(self, "apply_method")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterS3Import(dict):
    def __init__(__self__, *,
                 bucket_name: str,
                 ingestion_role: str,
                 source_engine: str,
                 source_engine_version: str,
                 bucket_prefix: Optional[str] = None):
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "ingestion_role", ingestion_role)
        pulumi.set(__self__, "source_engine", source_engine)
        pulumi.set(__self__, "source_engine_version", source_engine_version)
        if bucket_prefix is not None:
            pulumi.set(__self__, "bucket_prefix", bucket_prefix)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="ingestionRole")
    def ingestion_role(self) -> str:
        return pulumi.get(self, "ingestion_role")

    @property
    @pulumi.getter(name="sourceEngine")
    def source_engine(self) -> str:
        return pulumi.get(self, "source_engine")

    @property
    @pulumi.getter(name="sourceEngineVersion")
    def source_engine_version(self) -> str:
        return pulumi.get(self, "source_engine_version")

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> Optional[str]:
        return pulumi.get(self, "bucket_prefix")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterScalingConfiguration(dict):
    def __init__(__self__, *,
                 auto_pause: Optional[bool] = None,
                 max_capacity: Optional[float] = None,
                 min_capacity: Optional[float] = None,
                 seconds_until_auto_pause: Optional[float] = None,
                 timeout_action: Optional[str] = None):
        if auto_pause is not None:
            pulumi.set(__self__, "auto_pause", auto_pause)
        if max_capacity is not None:
            pulumi.set(__self__, "max_capacity", max_capacity)
        if min_capacity is not None:
            pulumi.set(__self__, "min_capacity", min_capacity)
        if seconds_until_auto_pause is not None:
            pulumi.set(__self__, "seconds_until_auto_pause", seconds_until_auto_pause)
        if timeout_action is not None:
            pulumi.set(__self__, "timeout_action", timeout_action)

    @property
    @pulumi.getter(name="autoPause")
    def auto_pause(self) -> Optional[bool]:
        return pulumi.get(self, "auto_pause")

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[float]:
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> Optional[float]:
        return pulumi.get(self, "min_capacity")

    @property
    @pulumi.getter(name="secondsUntilAutoPause")
    def seconds_until_auto_pause(self) -> Optional[float]:
        return pulumi.get(self, "seconds_until_auto_pause")

    @property
    @pulumi.getter(name="timeoutAction")
    def timeout_action(self) -> Optional[str]:
        return pulumi.get(self, "timeout_action")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GlobalClusterGlobalClusterMember(dict):
    def __init__(__self__, *,
                 db_cluster_arn: Optional[str] = None,
                 is_writer: Optional[bool] = None):
        if db_cluster_arn is not None:
            pulumi.set(__self__, "db_cluster_arn", db_cluster_arn)
        if is_writer is not None:
            pulumi.set(__self__, "is_writer", is_writer)

    @property
    @pulumi.getter(name="dbClusterArn")
    def db_cluster_arn(self) -> Optional[str]:
        return pulumi.get(self, "db_cluster_arn")

    @property
    @pulumi.getter(name="isWriter")
    def is_writer(self) -> Optional[bool]:
        return pulumi.get(self, "is_writer")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceS3Import(dict):
    def __init__(__self__, *,
                 bucket_name: str,
                 ingestion_role: str,
                 source_engine: str,
                 source_engine_version: str,
                 bucket_prefix: Optional[str] = None):
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "ingestion_role", ingestion_role)
        pulumi.set(__self__, "source_engine", source_engine)
        pulumi.set(__self__, "source_engine_version", source_engine_version)
        if bucket_prefix is not None:
            pulumi.set(__self__, "bucket_prefix", bucket_prefix)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="ingestionRole")
    def ingestion_role(self) -> str:
        return pulumi.get(self, "ingestion_role")

    @property
    @pulumi.getter(name="sourceEngine")
    def source_engine(self) -> str:
        return pulumi.get(self, "source_engine")

    @property
    @pulumi.getter(name="sourceEngineVersion")
    def source_engine_version(self) -> str:
        return pulumi.get(self, "source_engine_version")

    @property
    @pulumi.getter(name="bucketPrefix")
    def bucket_prefix(self) -> Optional[str]:
        return pulumi.get(self, "bucket_prefix")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OptionGroupOption(dict):
    def __init__(__self__, *,
                 option_name: str,
                 db_security_group_memberships: Optional[List[str]] = None,
                 option_settings: Optional[List['outputs.OptionGroupOptionOptionSetting']] = None,
                 port: Optional[float] = None,
                 version: Optional[str] = None,
                 vpc_security_group_memberships: Optional[List[str]] = None):
        pulumi.set(__self__, "option_name", option_name)
        if db_security_group_memberships is not None:
            pulumi.set(__self__, "db_security_group_memberships", db_security_group_memberships)
        if option_settings is not None:
            pulumi.set(__self__, "option_settings", option_settings)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if vpc_security_group_memberships is not None:
            pulumi.set(__self__, "vpc_security_group_memberships", vpc_security_group_memberships)

    @property
    @pulumi.getter(name="optionName")
    def option_name(self) -> str:
        return pulumi.get(self, "option_name")

    @property
    @pulumi.getter(name="dbSecurityGroupMemberships")
    def db_security_group_memberships(self) -> Optional[List[str]]:
        return pulumi.get(self, "db_security_group_memberships")

    @property
    @pulumi.getter(name="optionSettings")
    def option_settings(self) -> Optional[List['outputs.OptionGroupOptionOptionSetting']]:
        return pulumi.get(self, "option_settings")

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="vpcSecurityGroupMemberships")
    def vpc_security_group_memberships(self) -> Optional[List[str]]:
        return pulumi.get(self, "vpc_security_group_memberships")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OptionGroupOptionOptionSetting(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParameterGroupParameter(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str,
                 apply_method: Optional[str] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if apply_method is not None:
            pulumi.set(__self__, "apply_method", apply_method)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="applyMethod")
    def apply_method(self) -> Optional[str]:
        return pulumi.get(self, "apply_method")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SecurityGroupIngress(dict):
    def __init__(__self__, *,
                 cidr: Optional[str] = None,
                 security_group_id: Optional[str] = None,
                 security_group_name: Optional[str] = None,
                 security_group_owner_id: Optional[str] = None):
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if security_group_name is not None:
            pulumi.set(__self__, "security_group_name", security_group_name)
        if security_group_owner_id is not None:
            pulumi.set(__self__, "security_group_owner_id", security_group_owner_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[str]:
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="securityGroupName")
    def security_group_name(self) -> Optional[str]:
        return pulumi.get(self, "security_group_name")

    @property
    @pulumi.getter(name="securityGroupOwnerId")
    def security_group_owner_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_owner_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


