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
    'GetLocalGatewayVirtualInterfaceResult',
    'AwaitableGetLocalGatewayVirtualInterfaceResult',
    'get_local_gateway_virtual_interface',
]



@pulumi.output_type
class GetLocalGatewayVirtualInterfaceResult:
    """
    A collection of values returned by getLocalGatewayVirtualInterface.
    """
    def __init__(__self__, filters=None, id=None, local_address=None, local_bgp_asn=None, local_gateway_id=None, local_gateway_virtual_interface_ids=None, peer_address=None, peer_bgp_asn=None, tags=None, vlan=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if local_address and not isinstance(local_address, str):
            raise TypeError("Expected argument 'local_address' to be a str")
        pulumi.set(__self__, "local_address", local_address)
        if local_bgp_asn and not isinstance(local_bgp_asn, float):
            raise TypeError("Expected argument 'local_bgp_asn' to be a float")
        pulumi.set(__self__, "local_bgp_asn", local_bgp_asn)
        if local_gateway_id and not isinstance(local_gateway_id, str):
            raise TypeError("Expected argument 'local_gateway_id' to be a str")
        pulumi.set(__self__, "local_gateway_id", local_gateway_id)
        if local_gateway_virtual_interface_ids and not isinstance(local_gateway_virtual_interface_ids, list):
            raise TypeError("Expected argument 'local_gateway_virtual_interface_ids' to be a list")
        pulumi.set(__self__, "local_gateway_virtual_interface_ids", local_gateway_virtual_interface_ids)
        if peer_address and not isinstance(peer_address, str):
            raise TypeError("Expected argument 'peer_address' to be a str")
        pulumi.set(__self__, "peer_address", peer_address)
        if peer_bgp_asn and not isinstance(peer_bgp_asn, float):
            raise TypeError("Expected argument 'peer_bgp_asn' to be a float")
        pulumi.set(__self__, "peer_bgp_asn", peer_bgp_asn)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vlan and not isinstance(vlan, float):
            raise TypeError("Expected argument 'vlan' to be a float")
        pulumi.set(__self__, "vlan", vlan)

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetLocalGatewayVirtualInterfaceFilterResult']]:
        ...

    @property
    @pulumi.getter
    def id(self) -> str:
        ...

    @property
    @pulumi.getter(name="localAddress")
    def local_address(self) -> str:
        """
        Local address.
        """
        ...

    @property
    @pulumi.getter(name="localBgpAsn")
    def local_bgp_asn(self) -> float:
        """
        Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the EC2 Local Gateway.
        """
        ...

    @property
    @pulumi.getter(name="localGatewayId")
    def local_gateway_id(self) -> str:
        """
        Identifier of the EC2 Local Gateway.
        """
        ...

    @property
    @pulumi.getter(name="localGatewayVirtualInterfaceIds")
    def local_gateway_virtual_interface_ids(self) -> List[str]:
        ...

    @property
    @pulumi.getter(name="peerAddress")
    def peer_address(self) -> str:
        """
        Peer address.
        """
        ...

    @property
    @pulumi.getter(name="peerBgpAsn")
    def peer_bgp_asn(self) -> float:
        """
        Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the peer.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        ...

    @property
    @pulumi.getter
    def vlan(self) -> float:
        """
        Virtual Local Area Network.
        """
        ...



class AwaitableGetLocalGatewayVirtualInterfaceResult(GetLocalGatewayVirtualInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocalGatewayVirtualInterfaceResult(
            filters=self.filters,
            id=self.id,
            local_address=self.local_address,
            local_bgp_asn=self.local_bgp_asn,
            local_gateway_id=self.local_gateway_id,
            local_gateway_virtual_interface_ids=self.local_gateway_virtual_interface_ids,
            peer_address=self.peer_address,
            peer_bgp_asn=self.peer_bgp_asn,
            tags=self.tags,
            vlan=self.vlan)


def get_local_gateway_virtual_interface(filters: Optional[List[pulumi.InputType['GetLocalGatewayVirtualInterfaceFilterArgs']]] = None,
                                        id: Optional[str] = None,
                                        tags: Optional[Mapping[str, str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocalGatewayVirtualInterfaceResult:
    """
    Provides details about an EC2 Local Gateway Virtual Interface. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = [aws.ec2.get_local_gateway_virtual_interface(id=__value) for __key, __value in data["aws_ec2_local_gateway_virtual_interface_group"]["example"]["local_gateway_virtual_interface_ids"]]
    ```


    :param List[pulumi.InputType['GetLocalGatewayVirtualInterfaceFilterArgs']] filters: One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGatewayVirtualInterfaces.html) for supported filters. Detailed below.
    :param str id: Identifier of EC2 Local Gateway Virtual Interface.
    :param Mapping[str, str] tags: Key-value map of resource tags, each pair of which must exactly match a pair on the desired local gateway route table.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLocalGatewayVirtualInterface:getLocalGatewayVirtualInterface', __args__, opts=opts, typ=GetLocalGatewayVirtualInterfaceResult).value

    return AwaitableGetLocalGatewayVirtualInterfaceResult(
        filters=__ret__.filters,
        id=__ret__.id,
        local_address=__ret__.local_address,
        local_bgp_asn=__ret__.local_bgp_asn,
        local_gateway_id=__ret__.local_gateway_id,
        local_gateway_virtual_interface_ids=__ret__.local_gateway_virtual_interface_ids,
        peer_address=__ret__.peer_address,
        peer_bgp_asn=__ret__.peer_bgp_asn,
        tags=__ret__.tags,
        vlan=__ret__.vlan)
