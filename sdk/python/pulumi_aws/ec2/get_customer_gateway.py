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

__all__ = [
    'GetCustomerGatewayResult',
    'AwaitableGetCustomerGatewayResult',
    'get_customer_gateway',
]



@pulumi.output_type
class GetCustomerGatewayResult:
    """
    A collection of values returned by getCustomerGateway.
    """
    def __init__(__self__, arn=None, bgp_asn=None, filters=None, id=None, ip_address=None, tags=None, type=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if bgp_asn and not isinstance(bgp_asn, float):
            raise TypeError("Expected argument 'bgp_asn' to be a float")
        pulumi.set(__self__, "bgp_asn", bgp_asn)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the customer gateway.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> float:
        """
        (Optional) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        """
        return pulumi.get(self, "bgp_asn")

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetCustomerGatewayFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        (Optional) The IP address of the gateway's Internet-routable external interface.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Map of key-value pairs assigned to the gateway.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        (Optional) The type of customer gateway. The only type AWS supports at this time is "ipsec.1".
        """
        return pulumi.get(self, "type")



class AwaitableGetCustomerGatewayResult(GetCustomerGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomerGatewayResult(
            arn=self.arn,
            bgp_asn=self.bgp_asn,
            filters=self.filters,
            id=self.id,
            ip_address=self.ip_address,
            tags=self.tags,
            type=self.type)


def get_customer_gateway(filters: Optional[List[pulumi.InputType['GetCustomerGatewayFilterArgs']]] = None,
                         id: Optional[str] = None,
                         tags: Optional[Mapping[str, str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomerGatewayResult:
    """
    Get an existing AWS Customer Gateway.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    foo = aws.ec2.get_customer_gateway(filters=[{
        "name": "tag:Name",
        "values": ["foo-prod"],
    }])
    main = aws.ec2.VpnGateway("main",
        vpc_id=aws_vpc["main"]["id"],
        amazon_side_asn="7224")
    transit = aws.ec2.VpnConnection("transit",
        vpn_gateway_id=main.id,
        customer_gateway_id=foo.id,
        type=foo.type,
        static_routes_only=False)
    ```


    :param List[pulumi.InputType['GetCustomerGatewayFilterArgs']] filters: One or more [name-value pairs][dcg-filters] to filter by.
    :param str id: The ID of the gateway.
    :param Mapping[str, str] tags: Map of key-value pairs assigned to the gateway.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getCustomerGateway:getCustomerGateway', __args__, opts=opts, typ=GetCustomerGatewayResult).value

    return AwaitableGetCustomerGatewayResult(
        arn=__ret__.arn,
        bgp_asn=__ret__.bgp_asn,
        filters=__ret__.filters,
        id=__ret__.id,
        ip_address=__ret__.ip_address,
        tags=__ret__.tags,
        type=__ret__.type)
