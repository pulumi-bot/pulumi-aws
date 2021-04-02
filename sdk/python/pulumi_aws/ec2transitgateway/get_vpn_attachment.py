# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetVpnAttachmentResult',
    'AwaitableGetVpnAttachmentResult',
    'get_vpn_attachment',
]

@pulumi.output_type
class GetVpnAttachmentResult:
    """
    A collection of values returned by getVpnAttachment.
    """
    def __init__(__self__, filters=None, id=None, tags=None, transit_gateway_id=None, vpn_connection_id=None):
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
        if vpn_connection_id and not isinstance(vpn_connection_id, str):
            raise TypeError("Expected argument 'vpn_connection_id' to be a str")
        pulumi.set(__self__, "vpn_connection_id", vpn_connection_id)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetVpnAttachmentFilterResult']]:
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
        Key-value tags for the EC2 Transit Gateway VPN Attachment
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> Optional[str]:
        return pulumi.get(self, "transit_gateway_id")

    @property
    @pulumi.getter(name="vpnConnectionId")
    def vpn_connection_id(self) -> Optional[str]:
        return pulumi.get(self, "vpn_connection_id")


class AwaitableGetVpnAttachmentResult(GetVpnAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpnAttachmentResult(
            filters=self.filters,
            id=self.id,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id,
            vpn_connection_id=self.vpn_connection_id)


def get_vpn_attachment(filters: Optional[Sequence[pulumi.InputType['GetVpnAttachmentFilterArgs']]] = None,
                       tags: Optional[Mapping[str, str]] = None,
                       transit_gateway_id: Optional[str] = None,
                       vpn_connection_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpnAttachmentResult:
    """
    Get information on an EC2 Transit Gateway VPN Attachment.

    > EC2 Transit Gateway VPN Attachments are implicitly created by VPN Connections referencing an EC2 Transit Gateway so there is no managed resource. For ease, the `ec2.VpnConnection` resource includes a `transit_gateway_attachment_id` attribute which can replace some usage of this data source. For tagging the attachment, see the `ec2.Tag` resource.

    ## Example Usage
    ### By Transit Gateway and VPN Connection Identifiers

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_vpn_attachment(transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
        vpn_connection_id=aws_vpn_connection["example"]["id"])
    ```
    ### Filter

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.ec2transitgateway.get_vpn_attachment(filters=[aws.ec2transitgateway.GetVpnAttachmentFilterArgs(
        name="resource-id",
        values=["some-resource"],
    )])
    ```


    :param Sequence[pulumi.InputType['GetVpnAttachmentFilterArgs']] filters: Configuration block(s) for filtering. Detailed below.
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match a pair on the desired Transit Gateway VPN Attachment.
    :param str transit_gateway_id: Identifier of the EC2 Transit Gateway.
    :param str vpn_connection_id: Identifier of the EC2 VPN Connection.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['tags'] = tags
    __args__['transitGatewayId'] = transit_gateway_id
    __args__['vpnConnectionId'] = vpn_connection_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getVpnAttachment:getVpnAttachment', __args__, opts=opts, typ=GetVpnAttachmentResult).value

    return AwaitableGetVpnAttachmentResult(
        filters=__ret__.filters,
        id=__ret__.id,
        tags=__ret__.tags,
        transit_gateway_id=__ret__.transit_gateway_id,
        vpn_connection_id=__ret__.vpn_connection_id)
