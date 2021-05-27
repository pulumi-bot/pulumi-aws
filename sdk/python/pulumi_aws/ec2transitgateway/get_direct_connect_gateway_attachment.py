# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetDirectConnectGatewayAttachmentResult',
    'AwaitableGetDirectConnectGatewayAttachmentResult',
    'get_direct_connect_gateway_attachment',
]

@pulumi.output_type
class GetDirectConnectGatewayAttachmentResult:
    """
    A collection of values returned by getDirectConnectGatewayAttachment.
    """
    def __init__(__self__, dx_gateway_id=None, filters=None, id=None, tags=None, transit_gateway_id=None):
        if dx_gateway_id and not isinstance(dx_gateway_id, str):
            raise TypeError("Expected argument 'dx_gateway_id' to be a str")
        pulumi.set(__self__, "dx_gateway_id", dx_gateway_id)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        pulumi.set(__self__, "transit_gateway_id", transit_gateway_id)

    @property
    @pulumi.getter(name="dxGatewayId")
    def dx_gateway_id(self) -> Optional[str]:
        return pulumi.get(self, "dx_gateway_id")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDirectConnectGatewayAttachmentFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Key-value tags for the EC2 Transit Gateway Attachment
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> Optional[str]:
        return pulumi.get(self, "transit_gateway_id")


class AwaitableGetDirectConnectGatewayAttachmentResult(GetDirectConnectGatewayAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectConnectGatewayAttachmentResult(
            dx_gateway_id=self.dx_gateway_id,
            filters=self.filters,
            id=self.id,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id)


def get_direct_connect_gateway_attachment(dx_gateway_id: Optional[str] = None,
                                          filters: Optional[Sequence[pulumi.InputType['GetDirectConnectGatewayAttachmentFilterArgs']]] = None,
                                          tags: Optional[Mapping[str, str]] = None,
                                          transit_gateway_id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectConnectGatewayAttachmentResult:
    """
    Get information on an EC2 Transit Gateway's attachment to a Direct Connect Gateway.

    ## Example Usage
    ### By Transit Gateway and Direct Connect Gateway Identifiers

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_direct_connect_gateway_attachment(transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
        dx_gateway_id=aws_dx_gateway["example"]["id"])
    ```


    :param str dx_gateway_id: Identifier of the Direct Connect Gateway.
    :param Sequence[pulumi.InputType['GetDirectConnectGatewayAttachmentFilterArgs']] filters: Configuration block(s) for filtering. Detailed below.
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match a pair on the desired Transit Gateway Direct Connect Gateway Attachment.
    :param str transit_gateway_id: Identifier of the EC2 Transit Gateway.
    """
    __args__ = dict()
    __args__['dxGatewayId'] = dx_gateway_id
    __args__['filters'] = filters
    __args__['tags'] = tags
    __args__['transitGatewayId'] = transit_gateway_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getDirectConnectGatewayAttachment:getDirectConnectGatewayAttachment', __args__, opts=opts, typ=GetDirectConnectGatewayAttachmentResult).value

    return AwaitableGetDirectConnectGatewayAttachmentResult(
        dx_gateway_id=__ret__.dx_gateway_id,
        filters=__ret__.filters,
        id=__ret__.id,
        tags=__ret__.tags,
        transit_gateway_id=__ret__.transit_gateway_id)


@_utilities.lift_output_func(get_direct_connect_gateway_attachment)
def get_direct_connect_gateway_attachment_apply(dx_gateway_id: Optional[pulumi.Input[str]] = None,
                                                filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GetDirectConnectGatewayAttachmentFilterArgs']]]]] = None,
                                                tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                                                transit_gateway_id: Optional[pulumi.Input[str]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDirectConnectGatewayAttachmentResult]:
    ...
