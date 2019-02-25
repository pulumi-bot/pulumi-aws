# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetLoadBalancerResult:
    """
    A collection of values returned by getLoadBalancer.
    """
    def __init__(__self__, access_logs=None, arn=None, arn_suffix=None, dns_name=None, enable_deletion_protection=None, idle_timeout=None, internal=None, load_balancer_type=None, name=None, security_groups=None, subnet_mappings=None, subnets=None, tags=None, vpc_id=None, zone_id=None, id=None):
        if access_logs and not isinstance(access_logs, dict):
            raise TypeError('Expected argument access_logs to be a dict')
        __self__.access_logs = access_logs
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        if arn_suffix and not isinstance(arn_suffix, str):
            raise TypeError('Expected argument arn_suffix to be a str')
        __self__.arn_suffix = arn_suffix
        if dns_name and not isinstance(dns_name, str):
            raise TypeError('Expected argument dns_name to be a str')
        __self__.dns_name = dns_name
        if enable_deletion_protection and not isinstance(enable_deletion_protection, bool):
            raise TypeError('Expected argument enable_deletion_protection to be a bool')
        __self__.enable_deletion_protection = enable_deletion_protection
        if idle_timeout and not isinstance(idle_timeout, int):
            raise TypeError('Expected argument idle_timeout to be a int')
        __self__.idle_timeout = idle_timeout
        if internal and not isinstance(internal, bool):
            raise TypeError('Expected argument internal to be a bool')
        __self__.internal = internal
        if load_balancer_type and not isinstance(load_balancer_type, str):
            raise TypeError('Expected argument load_balancer_type to be a str')
        __self__.load_balancer_type = load_balancer_type
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        if security_groups and not isinstance(security_groups, list):
            raise TypeError('Expected argument security_groups to be a list')
        __self__.security_groups = security_groups
        if subnet_mappings and not isinstance(subnet_mappings, list):
            raise TypeError('Expected argument subnet_mappings to be a list')
        __self__.subnet_mappings = subnet_mappings
        if subnets and not isinstance(subnets, list):
            raise TypeError('Expected argument subnets to be a list')
        __self__.subnets = subnets
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError('Expected argument vpc_id to be a str')
        __self__.vpc_id = vpc_id
        if zone_id and not isinstance(zone_id, str):
            raise TypeError('Expected argument zone_id to be a str')
        __self__.zone_id = zone_id
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_load_balancer(arn=None,name=None,tags=None,opts=None):
    """
    > **Note:** `aws_alb` is known as `aws_lb`. The functionality is identical.
    
    Provides information about a Load Balancer.
    
    This data source can prove useful when a module accepts an LB as an input
    variable and needs to, for example, determine the security groups associated
    with it, etc.
    """
    __args__ = dict()

    __args__['arn'] = arn
    __args__['name'] = name
    __args__['tags'] = tags
    __ret__ = await pulumi.runtime.invoke('aws:applicationloadbalancing/getLoadBalancer:getLoadBalancer', __args__, opts=opts)

    return GetLoadBalancerResult(
        access_logs=__ret__.get('accessLogs'),
        arn=__ret__.get('arn'),
        arn_suffix=__ret__.get('arnSuffix'),
        dns_name=__ret__.get('dnsName'),
        enable_deletion_protection=__ret__.get('enableDeletionProtection'),
        idle_timeout=__ret__.get('idleTimeout'),
        internal=__ret__.get('internal'),
        load_balancer_type=__ret__.get('loadBalancerType'),
        name=__ret__.get('name'),
        security_groups=__ret__.get('securityGroups'),
        subnet_mappings=__ret__.get('subnetMappings'),
        subnets=__ret__.get('subnets'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'),
        zone_id=__ret__.get('zoneId'),
        id=__ret__.get('id'))
