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
    'GetNetworkInterfaceResult',
    'AwaitableGetNetworkInterfaceResult',
    'get_network_interface',
]

@pulumi.output_type
class GetNetworkInterfaceResult:
    """
    A collection of values returned by getNetworkInterface.
    """
    def __init__(__self__, associations=None, attachments=None, availability_zone=None, description=None, filters=None, id=None, interface_type=None, ipv6_addresses=None, mac_address=None, outpost_arn=None, owner_id=None, private_dns_name=None, private_ip=None, private_ips=None, requester_id=None, security_groups=None, subnet_id=None, tags=None, vpc_id=None):
        if associations and not isinstance(associations, list):
            raise TypeError("Expected argument 'associations' to be a list")
        pulumi.set(__self__, "associations", associations)
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        pulumi.set(__self__, "attachments", attachments)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_type and not isinstance(interface_type, str):
            raise TypeError("Expected argument 'interface_type' to be a str")
        pulumi.set(__self__, "interface_type", interface_type)
        if ipv6_addresses and not isinstance(ipv6_addresses, list):
            raise TypeError("Expected argument 'ipv6_addresses' to be a list")
        pulumi.set(__self__, "ipv6_addresses", ipv6_addresses)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        pulumi.set(__self__, "outpost_arn", outpost_arn)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if private_dns_name and not isinstance(private_dns_name, str):
            raise TypeError("Expected argument 'private_dns_name' to be a str")
        pulumi.set(__self__, "private_dns_name", private_dns_name)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        pulumi.set(__self__, "private_ips", private_ips)
        if requester_id and not isinstance(requester_id, str):
            raise TypeError("Expected argument 'requester_id' to be a str")
        pulumi.set(__self__, "requester_id", requester_id)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def associations(self) -> Sequence['outputs.GetNetworkInterfaceAssociationResult']:
        """
        The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.
        """
        return pulumi.get(self, "associations")

    @property
    @pulumi.getter
    def attachments(self) -> Sequence['outputs.GetNetworkInterfaceAttachmentResult']:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The Availability Zone.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the network interface.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetNetworkInterfaceFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceType")
    def interface_type(self) -> str:
        """
        The type of interface.
        """
        return pulumi.get(self, "interface_type")

    @property
    @pulumi.getter(name="ipv6Addresses")
    def ipv6_addresses(self) -> Sequence[str]:
        """
        List of IPv6 addresses to assign to the ENI.
        """
        return pulumi.get(self, "ipv6_addresses")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        """
        The MAC address.
        """
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the Outpost.
        """
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        The AWS account ID of the owner of the network interface.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="privateDnsName")
    def private_dns_name(self) -> str:
        """
        The private DNS name.
        """
        return pulumi.get(self, "private_dns_name")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        The private IPv4 address of the network interface within the subnet.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Sequence[str]:
        """
        The private IPv4 addresses associated with the network interface.
        """
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter(name="requesterId")
    def requester_id(self) -> str:
        """
        The ID of the entity that launched the instance on your behalf.
        """
        return pulumi.get(self, "requester_id")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        """
        The list of security groups for the network interface.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Any tags assigned to the network interface.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetNetworkInterfaceResult(GetNetworkInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfaceResult(
            associations=self.associations,
            attachments=self.attachments,
            availability_zone=self.availability_zone,
            description=self.description,
            filters=self.filters,
            id=self.id,
            interface_type=self.interface_type,
            ipv6_addresses=self.ipv6_addresses,
            mac_address=self.mac_address,
            outpost_arn=self.outpost_arn,
            owner_id=self.owner_id,
            private_dns_name=self.private_dns_name,
            private_ip=self.private_ip,
            private_ips=self.private_ips,
            requester_id=self.requester_id,
            security_groups=self.security_groups,
            subnet_id=self.subnet_id,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_network_interface(filters: Optional[Sequence[pulumi.InputType['GetNetworkInterfaceFilterArgs']]] = None,
                          id: Optional[str] = None,
                          tags: Optional[Mapping[str, str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkInterfaceResult:
    """
    Use this data source to get information about a Network Interface.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    bar = aws.ec2.get_network_interface(id="eni-01234567")
    ```


    :param Sequence[pulumi.InputType['GetNetworkInterfaceFilterArgs']] filters: One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-network-interfaces](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-network-interfaces.html) in the AWS CLI reference.
    :param str id: The identifier for the network interface.
    :param Mapping[str, str] tags: Any tags assigned to the network interface.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getNetworkInterface:getNetworkInterface', __args__, opts=opts, typ=GetNetworkInterfaceResult).value

    return AwaitableGetNetworkInterfaceResult(
        associations=__ret__.associations,
        attachments=__ret__.attachments,
        availability_zone=__ret__.availability_zone,
        description=__ret__.description,
        filters=__ret__.filters,
        id=__ret__.id,
        interface_type=__ret__.interface_type,
        ipv6_addresses=__ret__.ipv6_addresses,
        mac_address=__ret__.mac_address,
        outpost_arn=__ret__.outpost_arn,
        owner_id=__ret__.owner_id,
        private_dns_name=__ret__.private_dns_name,
        private_ip=__ret__.private_ip,
        private_ips=__ret__.private_ips,
        requester_id=__ret__.requester_id,
        security_groups=__ret__.security_groups,
        subnet_id=__ret__.subnet_id,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_network_interface)
def get_network_interface_apply(filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GetNetworkInterfaceFilterArgs']]]]] = None,
                                id: Optional[pulumi.Input[str]] = None,
                                tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkInterfaceResult]:
    ...
