# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['LoadBalancer']

warnings.warn("aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer", DeprecationWarning)


class LoadBalancer(pulumi.CustomResource):
    warnings.warn("aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_logs: Optional[pulumi.Input[pulumi.InputType['LoadBalancerAccessLogsArgs']]] = None,
                 drop_invalid_header_fields: Optional[pulumi.Input[bool]] = None,
                 enable_cross_zone_load_balancing: Optional[pulumi.Input[bool]] = None,
                 enable_deletion_protection: Optional[pulumi.Input[bool]] = None,
                 enable_http2: Optional[pulumi.Input[bool]] = None,
                 idle_timeout: Optional[pulumi.Input[float]] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 ip_address_type: Optional[pulumi.Input[str]] = None,
                 load_balancer_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 subnet_mappings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancerSubnetMappingArgs']]]]] = None,
                 subnets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a LoadBalancer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        pulumi.log.warn("LoadBalancer is deprecated: aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer")
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

            __props__['access_logs'] = access_logs
            __props__['drop_invalid_header_fields'] = drop_invalid_header_fields
            __props__['enable_cross_zone_load_balancing'] = enable_cross_zone_load_balancing
            __props__['enable_deletion_protection'] = enable_deletion_protection
            __props__['enable_http2'] = enable_http2
            __props__['idle_timeout'] = idle_timeout
            __props__['internal'] = internal
            __props__['ip_address_type'] = ip_address_type
            __props__['load_balancer_type'] = load_balancer_type
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['security_groups'] = security_groups
            __props__['subnet_mappings'] = subnet_mappings
            __props__['subnets'] = subnets
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['arn_suffix'] = None
            __props__['dns_name'] = None
            __props__['vpc_id'] = None
            __props__['zone_id'] = None
        super(LoadBalancer, __self__).__init__(
            'aws:applicationloadbalancing/loadBalancer:LoadBalancer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_logs: Optional[pulumi.Input[pulumi.InputType['LoadBalancerAccessLogsArgs']]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            arn_suffix: Optional[pulumi.Input[str]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            drop_invalid_header_fields: Optional[pulumi.Input[bool]] = None,
            enable_cross_zone_load_balancing: Optional[pulumi.Input[bool]] = None,
            enable_deletion_protection: Optional[pulumi.Input[bool]] = None,
            enable_http2: Optional[pulumi.Input[bool]] = None,
            idle_timeout: Optional[pulumi.Input[float]] = None,
            internal: Optional[pulumi.Input[bool]] = None,
            ip_address_type: Optional[pulumi.Input[str]] = None,
            load_balancer_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            security_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            subnet_mappings: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['LoadBalancerSubnetMappingArgs']]]]] = None,
            subnets: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'LoadBalancer':
        """
        Get an existing LoadBalancer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_logs"] = access_logs
        __props__["arn"] = arn
        __props__["arn_suffix"] = arn_suffix
        __props__["dns_name"] = dns_name
        __props__["drop_invalid_header_fields"] = drop_invalid_header_fields
        __props__["enable_cross_zone_load_balancing"] = enable_cross_zone_load_balancing
        __props__["enable_deletion_protection"] = enable_deletion_protection
        __props__["enable_http2"] = enable_http2
        __props__["idle_timeout"] = idle_timeout
        __props__["internal"] = internal
        __props__["ip_address_type"] = ip_address_type
        __props__["load_balancer_type"] = load_balancer_type
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["security_groups"] = security_groups
        __props__["subnet_mappings"] = subnet_mappings
        __props__["subnets"] = subnets
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        __props__["zone_id"] = zone_id
        return LoadBalancer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLogs")
    def access_logs(self) -> pulumi.Output[Optional['outputs.LoadBalancerAccessLogs']]:
        return pulumi.get(self, "access_logs")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="arnSuffix")
    def arn_suffix(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn_suffix")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="dropInvalidHeaderFields")
    def drop_invalid_header_fields(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "drop_invalid_header_fields")

    @property
    @pulumi.getter(name="enableCrossZoneLoadBalancing")
    def enable_cross_zone_load_balancing(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_cross_zone_load_balancing")

    @property
    @pulumi.getter(name="enableDeletionProtection")
    def enable_deletion_protection(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_deletion_protection")

    @property
    @pulumi.getter(name="enableHttp2")
    def enable_http2(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_http2")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter
    def internal(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "internal")

    @property
    @pulumi.getter(name="ipAddressType")
    def ip_address_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip_address_type")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetMappings")
    def subnet_mappings(self) -> pulumi.Output[List['outputs.LoadBalancerSubnetMapping']]:
        return pulumi.get(self, "subnet_mappings")

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "zone_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

