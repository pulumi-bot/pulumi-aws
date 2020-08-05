# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetVpcAttachmentResult:
    """
    A collection of values returned by getVpcAttachment.
    """
    def __init__(__self__, dns_support=None, filters=None, id=None, ipv6_support=None, subnet_ids=None, tags=None, transit_gateway_id=None, vpc_id=None, vpc_owner_id=None):
        if dns_support and not isinstance(dns_support, str):
            raise TypeError("Expected argument 'dns_support' to be a str")
        __self__.dns_support = dns_support
        """
        Whether DNS support is enabled.
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        EC2 Transit Gateway VPC Attachment identifier
        """
        if ipv6_support and not isinstance(ipv6_support, str):
            raise TypeError("Expected argument 'ipv6_support' to be a str")
        __self__.ipv6_support = ipv6_support
        """
        Whether IPv6 support is enabled.
        """
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        __self__.subnet_ids = subnet_ids
        """
        Identifiers of EC2 Subnets.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Key-value tags for the EC2 Transit Gateway VPC Attachment
        """
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        __self__.transit_gateway_id = transit_gateway_id
        """
        EC2 Transit Gateway identifier
        """
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        Identifier of EC2 VPC.
        """
        if vpc_owner_id and not isinstance(vpc_owner_id, str):
            raise TypeError("Expected argument 'vpc_owner_id' to be a str")
        __self__.vpc_owner_id = vpc_owner_id
        """
        Identifier of the AWS account that owns the EC2 VPC.
        """


class AwaitableGetVpcAttachmentResult(GetVpcAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcAttachmentResult(
            dns_support=self.dns_support,
            filters=self.filters,
            id=self.id,
            ipv6_support=self.ipv6_support,
            subnet_ids=self.subnet_ids,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id,
            vpc_id=self.vpc_id,
            vpc_owner_id=self.vpc_owner_id)


def get_vpc_attachment(filters=None, id=None, tags=None, opts=None):
    """
    Get information on an EC2 Transit Gateway VPC Attachment.

    ## Example Usage
    ### By Filter

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_vpc_attachment(filters=[{
        "name": "vpc-id",
        "values": ["vpc-12345678"],
    }])
    ```
    ### By Identifier

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_vpc_attachment(id="tgw-attach-12345678")
    ```


    :param list filters: One or more configuration blocks containing name-values filters. Detailed below.
    :param str id: Identifier of the EC2 Transit Gateway VPC Attachment.
    :param dict tags: Key-value tags for the EC2 Transit Gateway VPC Attachment

    The **filters** object supports the following:

      * `name` (`str`) - Name of the filter.
      * `values` (`list`) - List of one or more values for the filter.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getVpcAttachment:getVpcAttachment', __args__, opts=opts).value

    return AwaitableGetVpcAttachmentResult(
        dns_support=__ret__.get('dnsSupport'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        ipv6_support=__ret__.get('ipv6Support'),
        subnet_ids=__ret__.get('subnetIds'),
        tags=__ret__.get('tags'),
        transit_gateway_id=__ret__.get('transitGatewayId'),
        vpc_id=__ret__.get('vpcId'),
        vpc_owner_id=__ret__.get('vpcOwnerId'))
