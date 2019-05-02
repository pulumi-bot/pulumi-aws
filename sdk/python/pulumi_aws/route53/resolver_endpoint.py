# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ResolverEndpoint(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the Route 53 Resolver endpoint.
    """
    direction: pulumi.Output[str]
    """
    The direction of DNS queries to or from the Route 53 Resolver endpoint.
    Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
    or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
    """
    host_vpc_id: pulumi.Output[str]
    """
    The ID of the VPC that you want to create the resolver endpoint in.
    """
    ip_addresses: pulumi.Output[list]
    """
    The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
    to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
    """
    name: pulumi.Output[str]
    """
    The friendly name of the Route 53 Resolver endpoint.
    """
    security_group_ids: pulumi.Output[list]
    """
    The ID of one or more security groups that you want to use to control access to this VPC.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, direction=None, ip_addresses=None, name=None, security_group_ids=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a Route 53 Resolver endpoint resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] direction: The direction of DNS queries to or from the Route 53 Resolver endpoint.
               Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
               or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
        :param pulumi.Input[list] ip_addresses: The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
               to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
        :param pulumi.Input[str] name: The friendly name of the Route 53 Resolver endpoint.
        :param pulumi.Input[list] security_group_ids: The ID of one or more security groups that you want to use to control access to this VPC.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if direction is None:
            raise TypeError("Missing required property 'direction'")
        __props__['direction'] = direction

        if ip_addresses is None:
            raise TypeError("Missing required property 'ip_addresses'")
        __props__['ip_addresses'] = ip_addresses

        __props__['name'] = name

        if security_group_ids is None:
            raise TypeError("Missing required property 'security_group_ids'")
        __props__['security_group_ids'] = security_group_ids

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['host_vpc_id'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ResolverEndpoint, __self__).__init__(
            'aws:route53/resolverEndpoint:ResolverEndpoint',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

