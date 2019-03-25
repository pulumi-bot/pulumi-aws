# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetGatewayResult:
    """
    A collection of values returned by getGateway.
    """
    def __init__(__self__, amazon_side_asn=None, name=None, id=None):
        if amazon_side_asn and not isinstance(amazon_side_asn, str):
            raise TypeError('Expected argument amazon_side_asn to be a str')
        __self__.amazon_side_asn = amazon_side_asn
        """
        The ASN on the Amazon side of the connection.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_gateway(name=None,opts=None):
    """
    Retrieve information about a Direct Connect Gateway.
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('aws:directconnect/getGateway:getGateway', __args__, opts=opts)

    return GetGatewayResult(
        amazon_side_asn=__ret__.get('amazonSideAsn'),
        name=__ret__.get('name'),
        id=__ret__.get('id'))
