# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetGatewayResult',
    'AwaitableGetGatewayResult',
    'get_gateway',
]


class GetGatewayResult:
    """
    A collection of values returned by getGateway.
    """
    def __init__(__self__, amazon_side_asn=None, id=None, name=None, owner_account_id=None):
        if amazon_side_asn and not isinstance(amazon_side_asn, str):
            raise TypeError("Expected argument 'amazon_side_asn' to be a str")
        __self__.amazon_side_asn = amazon_side_asn
        """
        The ASN on the Amazon side of the connection.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if owner_account_id and not isinstance(owner_account_id, str):
            raise TypeError("Expected argument 'owner_account_id' to be a str")
        __self__.owner_account_id = owner_account_id
        """
        AWS Account ID of the gateway.
        """


class AwaitableGetGatewayResult(GetGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayResult(
            amazon_side_asn=self.amazon_side_asn,
            id=self.id,
            name=self.name,
            owner_account_id=self.owner_account_id)


def get_gateway(name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayResult:
    """
    Retrieve information about a Direct Connect Gateway.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.directconnect.get_gateway(name="example")
    ```


    :param str name: The name of the gateway to retrieve.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:directconnect/getGateway:getGateway', __args__, opts=opts).value

    return AwaitableGetGatewayResult(
        amazon_side_asn=__ret__.get('amazonSideAsn'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        owner_account_id=__ret__.get('ownerAccountId'))
