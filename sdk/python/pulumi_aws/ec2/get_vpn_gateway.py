# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetVpnGatewayResult:
    """
    A collection of values returned by getVpnGateway.
    """
    def __init__(__self__, amazon_side_asn=None, attached_vpc_id=None, availability_zone=None, filters=None, id=None, state=None, tags=None):
        if amazon_side_asn and not isinstance(amazon_side_asn, str):
            raise TypeError("Expected argument 'amazon_side_asn' to be a str")
        __self__.amazon_side_asn = amazon_side_asn
        if attached_vpc_id and not isinstance(attached_vpc_id, str):
            raise TypeError("Expected argument 'attached_vpc_id' to be a str")
        __self__.attached_vpc_id = attached_vpc_id
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
class AwaitableGetVpnGatewayResult(GetVpnGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnGatewayResult(
            amazon_side_asn=self.amazon_side_asn,
            attached_vpc_id=self.attached_vpc_id,
            availability_zone=self.availability_zone,
            filters=self.filters,
            id=self.id,
            state=self.state,
            tags=self.tags)

def get_vpn_gateway(amazon_side_asn=None,attached_vpc_id=None,availability_zone=None,filters=None,id=None,state=None,tags=None,opts=None):
    """
    The VPN Gateway data source provides details about
    a specific VPN gateway.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    selected = aws.ec2.get_vpn_gateway(filters=[{
        "name": "tag:Name",
        "values": ["vpn-gw"],
    }])
    pulumi.export("vpnGatewayId", selected.id)
    ```



    :param str amazon_side_asn: The Autonomous System Number (ASN) for the Amazon side of the specific VPN Gateway to retrieve.
    :param str attached_vpc_id: The ID of a VPC attached to the specific VPN Gateway to retrieve.
    :param str availability_zone: The Availability Zone of the specific VPN Gateway to retrieve.
    :param list filters: Custom filter block as described below.
    :param str id: The ID of the specific VPN Gateway to retrieve.
    :param str state: The state of the specific VPN Gateway to retrieve.
    :param dict tags: A map of tags, each pair of which must exactly match
           a pair on the desired VPN Gateway.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnGateways.html).
      * `values` (`list`) - Set of values that are accepted for the given field.
        A VPN Gateway will be selected if any one of the given values matches.
    """
    __args__ = dict()


    __args__['amazonSideAsn'] = amazon_side_asn
    __args__['attachedVpcId'] = attached_vpc_id
    __args__['availabilityZone'] = availability_zone
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getVpnGateway:getVpnGateway', __args__, opts=opts).value

    return AwaitableGetVpnGatewayResult(
        amazon_side_asn=__ret__.get('amazonSideAsn'),
        attached_vpc_id=__ret__.get('attachedVpcId'),
        availability_zone=__ret__.get('availabilityZone'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'))
