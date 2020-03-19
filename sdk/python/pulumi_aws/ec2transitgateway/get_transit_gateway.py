# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetTransitGatewayResult:
    """
    A collection of values returned by getTransitGateway.
    """
    def __init__(__self__, amazon_side_asn=None, arn=None, association_default_route_table_id=None, auto_accept_shared_attachments=None, default_route_table_association=None, default_route_table_propagation=None, description=None, dns_support=None, filters=None, id=None, owner_id=None, propagation_default_route_table_id=None, tags=None, vpn_ecmp_support=None):
        if amazon_side_asn and not isinstance(amazon_side_asn, float):
            raise TypeError("Expected argument 'amazon_side_asn' to be a float")
        __self__.amazon_side_asn = amazon_side_asn
        """
        Private Autonomous System Number (ASN) for the Amazon side of a BGP session
        """
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        EC2 Transit Gateway Amazon Resource Name (ARN)
        """
        if association_default_route_table_id and not isinstance(association_default_route_table_id, str):
            raise TypeError("Expected argument 'association_default_route_table_id' to be a str")
        __self__.association_default_route_table_id = association_default_route_table_id
        """
        Identifier of the default association route table
        """
        if auto_accept_shared_attachments and not isinstance(auto_accept_shared_attachments, str):
            raise TypeError("Expected argument 'auto_accept_shared_attachments' to be a str")
        __self__.auto_accept_shared_attachments = auto_accept_shared_attachments
        """
        Whether resource attachment requests are automatically accepted.
        """
        if default_route_table_association and not isinstance(default_route_table_association, str):
            raise TypeError("Expected argument 'default_route_table_association' to be a str")
        __self__.default_route_table_association = default_route_table_association
        """
        Whether resource attachments are automatically associated with the default association route table.
        """
        if default_route_table_propagation and not isinstance(default_route_table_propagation, str):
            raise TypeError("Expected argument 'default_route_table_propagation' to be a str")
        __self__.default_route_table_propagation = default_route_table_propagation
        """
        Whether resource attachments automatically propagate routes to the default propagation route table.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of the EC2 Transit Gateway
        """
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
        EC2 Transit Gateway identifier
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        Identifier of the AWS account that owns the EC2 Transit Gateway
        """
        if propagation_default_route_table_id and not isinstance(propagation_default_route_table_id, str):
            raise TypeError("Expected argument 'propagation_default_route_table_id' to be a str")
        __self__.propagation_default_route_table_id = propagation_default_route_table_id
        """
        Identifier of the default propagation route table.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Key-value tags for the EC2 Transit Gateway
        """
        if vpn_ecmp_support and not isinstance(vpn_ecmp_support, str):
            raise TypeError("Expected argument 'vpn_ecmp_support' to be a str")
        __self__.vpn_ecmp_support = vpn_ecmp_support
        """
        Whether VPN Equal Cost Multipath Protocol support is enabled.
        """
class AwaitableGetTransitGatewayResult(GetTransitGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitGatewayResult(
            amazon_side_asn=self.amazon_side_asn,
            arn=self.arn,
            association_default_route_table_id=self.association_default_route_table_id,
            auto_accept_shared_attachments=self.auto_accept_shared_attachments,
            default_route_table_association=self.default_route_table_association,
            default_route_table_propagation=self.default_route_table_propagation,
            description=self.description,
            dns_support=self.dns_support,
            filters=self.filters,
            id=self.id,
            owner_id=self.owner_id,
            propagation_default_route_table_id=self.propagation_default_route_table_id,
            tags=self.tags,
            vpn_ecmp_support=self.vpn_ecmp_support)

def get_transit_gateway(filters=None,id=None,tags=None,opts=None):
    """
    Get information on an EC2 Transit Gateway.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ec2_transit_gateway.html.markdown.


    :param list filters: One or more configuration blocks containing name-values filters. Detailed below.
    :param str id: Identifier of the EC2 Transit Gateway.

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
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getTransitGateway:getTransitGateway', __args__, opts=opts).value

    return AwaitableGetTransitGatewayResult(
        amazon_side_asn=__ret__.get('amazonSideAsn'),
        arn=__ret__.get('arn'),
        association_default_route_table_id=__ret__.get('associationDefaultRouteTableId'),
        auto_accept_shared_attachments=__ret__.get('autoAcceptSharedAttachments'),
        default_route_table_association=__ret__.get('defaultRouteTableAssociation'),
        default_route_table_propagation=__ret__.get('defaultRouteTablePropagation'),
        description=__ret__.get('description'),
        dns_support=__ret__.get('dnsSupport'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        owner_id=__ret__.get('ownerId'),
        propagation_default_route_table_id=__ret__.get('propagationDefaultRouteTableId'),
        tags=__ret__.get('tags'),
        vpn_ecmp_support=__ret__.get('vpnEcmpSupport'))
