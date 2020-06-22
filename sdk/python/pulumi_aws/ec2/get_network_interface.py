# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNetworkInterfaceResult:
    """
    A collection of values returned by getNetworkInterface.
    """
    def __init__(__self__, associations=None, attachments=None, availability_zone=None, description=None, filters=None, id=None, interface_type=None, ipv6_addresses=None, mac_address=None, outpost_arn=None, owner_id=None, private_dns_name=None, private_ip=None, private_ips=None, requester_id=None, security_groups=None, subnet_id=None, tags=None, vpc_id=None):
        if associations and not isinstance(associations, list):
            raise TypeError("Expected argument 'associations' to be a list")
        __self__.associations = associations
        """
        The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.
        """
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        __self__.attachments = attachments
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        """
        The Availability Zone.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of the network interface.
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if interface_type and not isinstance(interface_type, str):
            raise TypeError("Expected argument 'interface_type' to be a str")
        __self__.interface_type = interface_type
        """
        The type of interface.
        """
        if ipv6_addresses and not isinstance(ipv6_addresses, list):
            raise TypeError("Expected argument 'ipv6_addresses' to be a list")
        __self__.ipv6_addresses = ipv6_addresses
        """
        List of IPv6 addresses to assign to the ENI.
        """
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        __self__.mac_address = mac_address
        """
        The MAC address.
        """
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        __self__.outpost_arn = outpost_arn
        """
        The Amazon Resource Name (ARN) of the Outpost.
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        The AWS account ID of the owner of the network interface.
        """
        if private_dns_name and not isinstance(private_dns_name, str):
            raise TypeError("Expected argument 'private_dns_name' to be a str")
        __self__.private_dns_name = private_dns_name
        """
        The private DNS name.
        """
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        __self__.private_ip = private_ip
        """
        The private IPv4 address of the network interface within the subnet.
        """
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        __self__.private_ips = private_ips
        """
        The private IPv4 addresses associated with the network interface.
        """
        if requester_id and not isinstance(requester_id, str):
            raise TypeError("Expected argument 'requester_id' to be a str")
        __self__.requester_id = requester_id
        """
        The ID of the entity that launched the instance on your behalf.
        """
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        __self__.security_groups = security_groups
        """
        The list of security groups for the network interface.
        """
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        __self__.subnet_id = subnet_id
        """
        The ID of the subnet.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Any tags assigned to the network interface.
        """
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        __self__.vpc_id = vpc_id
        """
        The ID of the VPC.
        """
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

def get_network_interface(filters=None,id=None,tags=None,opts=None):
    """
    Use this data source to get information about a Network Interface.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    bar = aws.ec2.get_network_interface(id="eni-01234567")
    ```


    :param list filters: One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-network-interfaces](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-network-interfaces.html) in the AWS CLI reference.
    :param str id: The identifier for the network interface.
    :param dict tags: Any tags assigned to the network interface.

    The **filters** object supports the following:

      * `name` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getNetworkInterface:getNetworkInterface', __args__, opts=opts).value

    return AwaitableGetNetworkInterfaceResult(
        associations=__ret__.get('associations'),
        attachments=__ret__.get('attachments'),
        availability_zone=__ret__.get('availabilityZone'),
        description=__ret__.get('description'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        interface_type=__ret__.get('interfaceType'),
        ipv6_addresses=__ret__.get('ipv6Addresses'),
        mac_address=__ret__.get('macAddress'),
        outpost_arn=__ret__.get('outpostArn'),
        owner_id=__ret__.get('ownerId'),
        private_dns_name=__ret__.get('privateDnsName'),
        private_ip=__ret__.get('privateIp'),
        private_ips=__ret__.get('privateIps'),
        requester_id=__ret__.get('requesterId'),
        security_groups=__ret__.get('securityGroups'),
        subnet_id=__ret__.get('subnetId'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'))
