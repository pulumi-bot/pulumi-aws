# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'RecordAliasArgs',
    'RecordFailoverRoutingPolicyArgs',
    'RecordGeolocationRoutingPolicyArgs',
    'RecordLatencyRoutingPolicyArgs',
    'RecordWeightedRoutingPolicyArgs',
    'ResolverEndpointIpAddressArgs',
    'ResolverRuleTargetIpArgs',
    'ZoneVpcArgs',
]

@pulumi.input_type
class RecordAliasArgs:
    def __init__(__self__, *,
                 evaluate_target_health: pulumi.Input[bool],
                 name: pulumi.Input[str],
                 zone_id: pulumi.Input[str]):
        pulumi.set(__self__, "evaluate_target_health", evaluate_target_health)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="evaluateTargetHealth")
    def evaluate_target_health(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "evaluate_target_health")

    @evaluate_target_health.setter
    def evaluate_target_health(self, value: pulumi.Input[bool]):
        pulumi.set(self, "evaluate_target_health", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class RecordFailoverRoutingPolicyArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class RecordGeolocationRoutingPolicyArgs:
    def __init__(__self__, *,
                 continent: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 subdivision: Optional[pulumi.Input[str]] = None):
        if continent is not None:
            pulumi.set(__self__, "continent", continent)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if subdivision is not None:
            pulumi.set(__self__, "subdivision", subdivision)

    @property
    @pulumi.getter
    def continent(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "continent")

    @continent.setter
    def continent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "continent", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter
    def subdivision(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subdivision")

    @subdivision.setter
    def subdivision(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdivision", value)


@pulumi.input_type
class RecordLatencyRoutingPolicyArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str]):
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class RecordWeightedRoutingPolicyArgs:
    def __init__(__self__, *,
                 weight: pulumi.Input[float]):
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[float]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[float]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class ResolverEndpointIpAddressArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "subnet_id", subnet_id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip_id is not None:
            pulumi.set(__self__, "ip_id", ip_id)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_id")

    @ip_id.setter
    def ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_id", value)


@pulumi.input_type
class ResolverRuleTargetIpArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 port: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class ZoneVpcArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 vpc_region: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_region is not None:
            pulumi.set(__self__, "vpc_region", vpc_region)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_region")

    @vpc_region.setter
    def vpc_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_region", value)


