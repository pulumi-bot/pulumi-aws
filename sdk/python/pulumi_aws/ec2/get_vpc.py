# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetVpcResult:
    """
    A collection of values returned by getVpc.
    """
    def __init__(__self__, arn=None, cidr_block=None, cidr_block_associations=None, default=None, dhcp_options_id=None, enable_dns_hostnames=None, enable_dns_support=None, filters=None, id=None, instance_tenancy=None, ipv6_association_id=None, ipv6_cidr_block=None, main_route_table_id=None, owner_id=None, state=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Amazon Resource Name (ARN) of VPC
        """
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        __self__.cidr_block = cidr_block
        """
        The CIDR block for the association.
        """
        if cidr_block_associations and not isinstance(cidr_block_associations, list):
            raise TypeError("Expected argument 'cidr_block_associations' to be a list")
        __self__.cidr_block_associations = cidr_block_associations
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        __self__.default = default
        if dhcp_options_id and not isinstance(dhcp_options_id, str):
            raise TypeError("Expected argument 'dhcp_options_id' to be a str")
        __self__.dhcp_options_id = dhcp_options_id
        if enable_dns_hostnames and not isinstance(enable_dns_hostnames, bool):
            raise TypeError("Expected argument 'enable_dns_hostnames' to be a bool")
        __self__.enable_dns_hostnames = enable_dns_hostnames
        """
        Whether or not the VPC has DNS hostname support
        """
        if enable_dns_support and not isinstance(enable_dns_support, bool):
            raise TypeError("Expected argument 'enable_dns_support' to be a bool")
        __self__.enable_dns_support = enable_dns_support
        """
        Whether or not the VPC has DNS support
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if instance_tenancy and not isinstance(instance_tenancy, str):
            raise TypeError("Expected argument 'instance_tenancy' to be a str")
        __self__.instance_tenancy = instance_tenancy
        """
        The allowed tenancy of instances launched into the
        selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
        """
        if ipv6_association_id and not isinstance(ipv6_association_id, str):
            raise TypeError("Expected argument 'ipv6_association_id' to be a str")
        __self__.ipv6_association_id = ipv6_association_id
        """
        The association ID for the IPv6 CIDR block.
        """
        if ipv6_cidr_block and not isinstance(ipv6_cidr_block, str):
            raise TypeError("Expected argument 'ipv6_cidr_block' to be a str")
        __self__.ipv6_cidr_block = ipv6_cidr_block
        """
        The IPv6 CIDR block.
        """
        if main_route_table_id and not isinstance(main_route_table_id, str):
            raise TypeError("Expected argument 'main_route_table_id' to be a str")
        __self__.main_route_table_id = main_route_table_id
        """
        The ID of the main route table associated with this VPC.
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        The ID of the AWS account that owns the VPC.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        The State of the association.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
class AwaitableGetVpcResult(GetVpcResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcResult(
            arn=self.arn,
            cidr_block=self.cidr_block,
            cidr_block_associations=self.cidr_block_associations,
            default=self.default,
            dhcp_options_id=self.dhcp_options_id,
            enable_dns_hostnames=self.enable_dns_hostnames,
            enable_dns_support=self.enable_dns_support,
            filters=self.filters,
            id=self.id,
            instance_tenancy=self.instance_tenancy,
            ipv6_association_id=self.ipv6_association_id,
            ipv6_cidr_block=self.ipv6_cidr_block,
            main_route_table_id=self.main_route_table_id,
            owner_id=self.owner_id,
            state=self.state,
            tags=self.tags)

def get_vpc(cidr_block=None,default=None,dhcp_options_id=None,filters=None,id=None,state=None,tags=None,opts=None):
    """
    `ec2.Vpc` provides details about a specific VPC.

    This resource can prove useful when a module accepts a vpc id as
    an input variable and needs to, for example, determine the CIDR block of that
    VPC.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpc.html.markdown.


    :param str cidr_block: The cidr block of the desired VPC.
    :param bool default: Boolean constraint on whether the desired VPC is
           the default VPC for the region.
    :param str dhcp_options_id: The DHCP options id of the desired VPC.
    :param list filters: Custom filter block as described below.
    :param str id: The id of the specific VPC to retrieve.
    :param str state: The current state of the desired VPC.
           Can be either `"pending"` or `"available"`.
    :param dict tags: A mapping of tags, each pair of which must exactly match
           a pair on the desired VPC.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html).
      * `values` (`list`) - Set of values that are accepted for the given field.
        A VPC will be selected if any one of the given values matches.
    """
    __args__ = dict()


    __args__['cidrBlock'] = cidr_block
    __args__['default'] = default
    __args__['dhcpOptionsId'] = dhcp_options_id
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getVpc:getVpc', __args__, opts=opts).value

    return AwaitableGetVpcResult(
        arn=__ret__.get('arn'),
        cidr_block=__ret__.get('cidrBlock'),
        cidr_block_associations=__ret__.get('cidrBlockAssociations'),
        default=__ret__.get('default'),
        dhcp_options_id=__ret__.get('dhcpOptionsId'),
        enable_dns_hostnames=__ret__.get('enableDnsHostnames'),
        enable_dns_support=__ret__.get('enableDnsSupport'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        instance_tenancy=__ret__.get('instanceTenancy'),
        ipv6_association_id=__ret__.get('ipv6AssociationId'),
        ipv6_cidr_block=__ret__.get('ipv6CidrBlock'),
        main_route_table_id=__ret__.get('mainRouteTableId'),
        owner_id=__ret__.get('ownerId'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'))
