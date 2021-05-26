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
    'GetSubnetResult',
    'AwaitableGetSubnetResult',
    'get_subnet',
]

@pulumi.output_type
class GetSubnetResult:
    """
    A collection of values returned by getSubnet.
    """
    def __init__(__self__, arn=None, assign_ipv6_address_on_creation=None, availability_zone=None, availability_zone_id=None, available_ip_address_count=None, cidr_block=None, customer_owned_ipv4_pool=None, default_for_az=None, filters=None, id=None, ipv6_cidr_block=None, ipv6_cidr_block_association_id=None, map_customer_owned_ip_on_launch=None, map_public_ip_on_launch=None, outpost_arn=None, owner_id=None, state=None, tags=None, vpc_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if assign_ipv6_address_on_creation and not isinstance(assign_ipv6_address_on_creation, bool):
            raise TypeError("Expected argument 'assign_ipv6_address_on_creation' to be a bool")
        pulumi.set(__self__, "assign_ipv6_address_on_creation", assign_ipv6_address_on_creation)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if availability_zone_id and not isinstance(availability_zone_id, str):
            raise TypeError("Expected argument 'availability_zone_id' to be a str")
        pulumi.set(__self__, "availability_zone_id", availability_zone_id)
        if available_ip_address_count and not isinstance(available_ip_address_count, int):
            raise TypeError("Expected argument 'available_ip_address_count' to be a int")
        pulumi.set(__self__, "available_ip_address_count", available_ip_address_count)
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        pulumi.set(__self__, "cidr_block", cidr_block)
        if customer_owned_ipv4_pool and not isinstance(customer_owned_ipv4_pool, str):
            raise TypeError("Expected argument 'customer_owned_ipv4_pool' to be a str")
        pulumi.set(__self__, "customer_owned_ipv4_pool", customer_owned_ipv4_pool)
        if default_for_az and not isinstance(default_for_az, bool):
            raise TypeError("Expected argument 'default_for_az' to be a bool")
        pulumi.set(__self__, "default_for_az", default_for_az)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv6_cidr_block and not isinstance(ipv6_cidr_block, str):
            raise TypeError("Expected argument 'ipv6_cidr_block' to be a str")
        pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        if ipv6_cidr_block_association_id and not isinstance(ipv6_cidr_block_association_id, str):
            raise TypeError("Expected argument 'ipv6_cidr_block_association_id' to be a str")
        pulumi.set(__self__, "ipv6_cidr_block_association_id", ipv6_cidr_block_association_id)
        if map_customer_owned_ip_on_launch and not isinstance(map_customer_owned_ip_on_launch, bool):
            raise TypeError("Expected argument 'map_customer_owned_ip_on_launch' to be a bool")
        pulumi.set(__self__, "map_customer_owned_ip_on_launch", map_customer_owned_ip_on_launch)
        if map_public_ip_on_launch and not isinstance(map_public_ip_on_launch, bool):
            raise TypeError("Expected argument 'map_public_ip_on_launch' to be a bool")
        pulumi.set(__self__, "map_public_ip_on_launch", map_public_ip_on_launch)
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        pulumi.set(__self__, "outpost_arn", outpost_arn)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        ARN of the subnet.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assignIpv6AddressOnCreation")
    def assign_ipv6_address_on_creation(self) -> bool:
        """
        Whether an IPv6 address is assigned on creation.
        """
        return pulumi.get(self, "assign_ipv6_address_on_creation")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="availabilityZoneId")
    def availability_zone_id(self) -> str:
        return pulumi.get(self, "availability_zone_id")

    @property
    @pulumi.getter(name="availableIpAddressCount")
    def available_ip_address_count(self) -> int:
        """
        Available IP addresses of the subnet.
        """
        return pulumi.get(self, "available_ip_address_count")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> str:
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="customerOwnedIpv4Pool")
    def customer_owned_ipv4_pool(self) -> str:
        """
        Identifier of customer owned IPv4 address pool.
        """
        return pulumi.get(self, "customer_owned_ipv4_pool")

    @property
    @pulumi.getter(name="defaultForAz")
    def default_for_az(self) -> bool:
        return pulumi.get(self, "default_for_az")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetSubnetFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> str:
        return pulumi.get(self, "ipv6_cidr_block")

    @property
    @pulumi.getter(name="ipv6CidrBlockAssociationId")
    def ipv6_cidr_block_association_id(self) -> str:
        """
        Association ID of the IPv6 CIDR block.
        """
        return pulumi.get(self, "ipv6_cidr_block_association_id")

    @property
    @pulumi.getter(name="mapCustomerOwnedIpOnLaunch")
    def map_customer_owned_ip_on_launch(self) -> bool:
        """
        Whether customer owned IP addresses are assigned on network interface creation.
        """
        return pulumi.get(self, "map_customer_owned_ip_on_launch")

    @property
    @pulumi.getter(name="mapPublicIpOnLaunch")
    def map_public_ip_on_launch(self) -> bool:
        """
        Whether public IP addresses are assigned on instance launch.
        """
        return pulumi.get(self, "map_public_ip_on_launch")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> str:
        """
        ARN of the Outpost.
        """
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        ID of the AWS account that owns the subnet.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetSubnetResult(GetSubnetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetResult(
            arn=self.arn,
            assign_ipv6_address_on_creation=self.assign_ipv6_address_on_creation,
            availability_zone=self.availability_zone,
            availability_zone_id=self.availability_zone_id,
            available_ip_address_count=self.available_ip_address_count,
            cidr_block=self.cidr_block,
            customer_owned_ipv4_pool=self.customer_owned_ipv4_pool,
            default_for_az=self.default_for_az,
            filters=self.filters,
            id=self.id,
            ipv6_cidr_block=self.ipv6_cidr_block,
            ipv6_cidr_block_association_id=self.ipv6_cidr_block_association_id,
            map_customer_owned_ip_on_launch=self.map_customer_owned_ip_on_launch,
            map_public_ip_on_launch=self.map_public_ip_on_launch,
            outpost_arn=self.outpost_arn,
            owner_id=self.owner_id,
            state=self.state,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_subnet(availability_zone: Optional[str] = None,
               availability_zone_id: Optional[str] = None,
               cidr_block: Optional[str] = None,
               default_for_az: Optional[bool] = None,
               filters: Optional[Sequence[pulumi.InputType['GetSubnetFilterArgs']]] = None,
               id: Optional[str] = None,
               ipv6_cidr_block: Optional[str] = None,
               state: Optional[str] = None,
               tags: Optional[Mapping[str, str]] = None,
               vpc_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetResult:
    """
    `ec2.Subnet` provides details about a specific VPC subnet.

    This resource can prove useful when a module accepts a subnet ID as an input variable and needs to, for example, determine the ID of the VPC that the subnet belongs to.

    ## Example Usage

    The following example shows how one might accept a subnet ID as a variable and use this data source to obtain the data necessary to create a security group that allows connections from hosts in that subnet.

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    subnet_id = config.require_object("subnetId")
    selected = aws.ec2.get_subnet(id=subnet_id)
    subnet = aws.ec2.SecurityGroup("subnet",
        vpc_id=selected.vpc_id,
        ingress=[aws.ec2.SecurityGroupIngressArgs(
            cidr_blocks=[selected.cidr_block],
            from_port=80,
            to_port=80,
            protocol="tcp",
        )])
    ```
    ### Filter Example

    If you want to match against tag `Name`, use:

    ```python
    import pulumi
    import pulumi_aws as aws

    selected = aws.ec2.get_subnet(filters=[aws.ec2.GetSubnetFilterArgs(
        name="tag:Name",
        values=["yakdriver"],
    )])
    ```


    :param str availability_zone: Availability zone where the subnet must reside.
    :param str availability_zone_id: ID of the Availability Zone for the subnet.
    :param str cidr_block: CIDR block of the desired subnet.
    :param bool default_for_az: Whether the desired subnet must be the default subnet for its associated availability zone.
    :param Sequence[pulumi.InputType['GetSubnetFilterArgs']] filters: Configuration block. Detailed below.
    :param str id: ID of the specific subnet to retrieve.
    :param str ipv6_cidr_block: IPv6 CIDR block of the desired subnet.
    :param str state: State that the desired subnet must have.
    :param Mapping[str, str] tags: Map of tags, each pair of which must exactly match a pair on the desired subnet.
    :param str vpc_id: ID of the VPC that the desired subnet belongs to.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['availabilityZoneId'] = availability_zone_id
    __args__['cidrBlock'] = cidr_block
    __args__['defaultForAz'] = default_for_az
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['ipv6CidrBlock'] = ipv6_cidr_block
    __args__['state'] = state
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getSubnet:getSubnet', __args__, opts=opts, typ=GetSubnetResult).value

    return AwaitableGetSubnetResult(
        arn=__ret__.arn,
        assign_ipv6_address_on_creation=__ret__.assign_ipv6_address_on_creation,
        availability_zone=__ret__.availability_zone,
        availability_zone_id=__ret__.availability_zone_id,
        available_ip_address_count=__ret__.available_ip_address_count,
        cidr_block=__ret__.cidr_block,
        customer_owned_ipv4_pool=__ret__.customer_owned_ipv4_pool,
        default_for_az=__ret__.default_for_az,
        filters=__ret__.filters,
        id=__ret__.id,
        ipv6_cidr_block=__ret__.ipv6_cidr_block,
        ipv6_cidr_block_association_id=__ret__.ipv6_cidr_block_association_id,
        map_customer_owned_ip_on_launch=__ret__.map_customer_owned_ip_on_launch,
        map_public_ip_on_launch=__ret__.map_public_ip_on_launch,
        outpost_arn=__ret__.outpost_arn,
        owner_id=__ret__.owner_id,
        state=__ret__.state,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_subnet)
def get_subnet_apply(availability_zone: Optional[pulumi.Input[str]] = None,
                     availability_zone_id: Optional[pulumi.Input[str]] = None,
                     cidr_block: Optional[pulumi.Input[str]] = None,
                     default_for_az: Optional[pulumi.Input[bool]] = None,
                     filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GetSubnetFilterArgs']]]]] = None,
                     id: Optional[pulumi.Input[str]] = None,
                     ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                     state: Optional[pulumi.Input[str]] = None,
                     tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                     vpc_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetResult]:
    ...
