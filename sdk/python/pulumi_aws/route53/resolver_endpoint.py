# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
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

      * `ip` (`str`) - The IP address in the subnet that you want to use for DNS queries.
      * `ipId` (`str`)
      * `subnet_id` (`str`) - The ID of the subnet that contains the IP address.
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
    def __init__(__self__, resource_name, opts=None, direction=None, ip_addresses=None, name=None, security_group_ids=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Route 53 Resolver endpoint resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_endpoint.html.markdown.

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

        The **ip_addresses** object supports the following:

          * `ip` (`pulumi.Input[str]`) - The IP address in the subnet that you want to use for DNS queries.
          * `ipId` (`pulumi.Input[str]`)
          * `subnet_id` (`pulumi.Input[str]`) - The ID of the subnet that contains the IP address.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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
        super(ResolverEndpoint, __self__).__init__(
            'aws:route53/resolverEndpoint:ResolverEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, direction=None, host_vpc_id=None, ip_addresses=None, name=None, security_group_ids=None, tags=None):
        """
        Get an existing ResolverEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Route 53 Resolver endpoint.
        :param pulumi.Input[str] direction: The direction of DNS queries to or from the Route 53 Resolver endpoint.
               Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
               or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
        :param pulumi.Input[str] host_vpc_id: The ID of the VPC that you want to create the resolver endpoint in.
        :param pulumi.Input[list] ip_addresses: The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
               to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
        :param pulumi.Input[str] name: The friendly name of the Route 53 Resolver endpoint.
        :param pulumi.Input[list] security_group_ids: The ID of one or more security groups that you want to use to control access to this VPC.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **ip_addresses** object supports the following:

          * `ip` (`pulumi.Input[str]`) - The IP address in the subnet that you want to use for DNS queries.
          * `ipId` (`pulumi.Input[str]`)
          * `subnet_id` (`pulumi.Input[str]`) - The ID of the subnet that contains the IP address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["direction"] = direction
        __props__["host_vpc_id"] = host_vpc_id
        __props__["ip_addresses"] = ip_addresses
        __props__["name"] = name
        __props__["security_group_ids"] = security_group_ids
        __props__["tags"] = tags
        return ResolverEndpoint(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

