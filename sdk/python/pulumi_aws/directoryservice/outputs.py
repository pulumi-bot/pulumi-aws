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
    'GetDirectoryConnectSettingResult',
    'GetDirectoryVpcSettingResult',
]

@pulumi.output_type
class DirectoryConnectSettings(dict):
    def __init__(__self__, *,
                 customer_dns_ips: List[str],
                 customer_username: str,
                 subnet_ids: List[str],
                 vpc_id: str,
                 availability_zones: Optional[List[str]] = None,
                 connect_ips: Optional[List[str]] = None):
        pulumi.set(__self__, "customer_dns_ips", customer_dns_ips)
        pulumi.set(__self__, "customer_username", customer_username)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)
        if connect_ips is not None:
            pulumi.set(__self__, "connect_ips", connect_ips)

    @property
    @pulumi.getter(name="customerDnsIps")
    def customer_dns_ips(self) -> List[str]:
        return pulumi.get(self, "customer_dns_ips")

    @property
    @pulumi.getter(name="customerUsername")
    def customer_username(self) -> str:
        return pulumi.get(self, "customer_username")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="connectIps")
    def connect_ips(self) -> Optional[List[str]]:
        return pulumi.get(self, "connect_ips")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DirectoryVpcSettings(dict):
    def __init__(__self__, *,
                 subnet_ids: List[str],
                 vpc_id: str,
                 availability_zones: Optional[List[str]] = None):
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        return pulumi.get(self, "availability_zones")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetDirectoryConnectSettingResult(dict):
    def __init__(__self__, *,
                 availability_zones: List[str],
                 connect_ips: List[str],
                 customer_dns_ips: List[str],
                 customer_username: str,
                 subnet_ids: List[str],
                 vpc_id: str):
        pulumi.set(__self__, "availability_zones", availability_zones)
        pulumi.set(__self__, "connect_ips", connect_ips)
        pulumi.set(__self__, "customer_dns_ips", customer_dns_ips)
        pulumi.set(__self__, "customer_username", customer_username)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> List[str]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="connectIps")
    def connect_ips(self) -> List[str]:
        return pulumi.get(self, "connect_ips")

    @property
    @pulumi.getter(name="customerDnsIps")
    def customer_dns_ips(self) -> List[str]:
        return pulumi.get(self, "customer_dns_ips")

    @property
    @pulumi.getter(name="customerUsername")
    def customer_username(self) -> str:
        return pulumi.get(self, "customer_username")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GetDirectoryVpcSettingResult(dict):
    def __init__(__self__, *,
                 availability_zones: List[str],
                 subnet_ids: List[str],
                 vpc_id: str):
        pulumi.set(__self__, "availability_zones", availability_zones)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> List[str]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


