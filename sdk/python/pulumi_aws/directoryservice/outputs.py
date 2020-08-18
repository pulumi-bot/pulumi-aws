# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'DirectoryConnectSettings',
    'DirectoryVpcSettings',
    'GetDirectoryConnectSetting',
    'GetDirectoryVpcSetting',
]

@pulumi.output_type
class DirectoryConnectSettings(dict):
    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        ...

    @property
    @pulumi.getter(name="connectIps")
    def connect_ips(self) -> Optional[List[str]]:
        """
        The IP addresses of the AD Connector servers.
        """
        ...

    @property
    @pulumi.getter(name="customerDnsIps")
    def customer_dns_ips(self) -> List[str]:
        """
        The DNS IP addresses of the domain to connect to.
        """
        ...

    @property
    @pulumi.getter(name="customerUsername")
    def customer_username(self) -> str:
        """
        The username corresponding to the password provided.
        """
        ...

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        """
        The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        """
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The identifier of the VPC that the directory is in.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DirectoryVpcSettings(dict):
    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        ...

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        """
        The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        """
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The identifier of the VPC that the directory is in.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetDirectoryConnectSetting(dict):
    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> List[str]:
        ...

    @property
    @pulumi.getter(name="connectIps")
    def connect_ips(self) -> List[str]:
        """
        The IP addresses of the AD Connector servers.
        """
        ...

    @property
    @pulumi.getter(name="customerDnsIps")
    def customer_dns_ips(self) -> List[str]:
        """
        The DNS IP addresses of the domain to connect to.
        """
        ...

    @property
    @pulumi.getter(name="customerUsername")
    def customer_username(self) -> str:
        """
        The username corresponding to the password provided.
        """
        ...

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        """
        The identifiers of the subnets for the connector servers (2 subnets in 2 different AZs).
        """
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC that the connector is in.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetDirectoryVpcSetting(dict):
    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> List[str]:
        ...

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        """
        The identifiers of the subnets for the connector servers (2 subnets in 2 different AZs).
        """
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC that the connector is in.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


