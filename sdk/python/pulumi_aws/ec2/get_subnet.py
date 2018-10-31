# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetSubnetResult(object):
    """
    A collection of values returned by getSubnet.
    """
    def __init__(__self__, arn=None, assign_ipv6_address_on_creation=None, availability_zone=None, cidr_block=None, default_for_az=None, id=None, ipv6_cidr_block=None, ipv6_cidr_block_association_id=None, map_public_ip_on_launch=None, state=None, tags=None, vpc_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        The ARN of the subnet.
        """
        if assign_ipv6_address_on_creation and not isinstance(assign_ipv6_address_on_creation, bool):
            raise TypeError('Expected argument assign_ipv6_address_on_creation to be a bool')
        __self__.assign_ipv6_address_on_creation = assign_ipv6_address_on_creation
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError('Expected argument availability_zone to be a str')
        __self__.availability_zone = availability_zone
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError('Expected argument cidr_block to be a str')
        __self__.cidr_block = cidr_block
        if default_for_az and not isinstance(default_for_az, bool):
            raise TypeError('Expected argument default_for_az to be a bool')
        __self__.default_for_az = default_for_az
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        if ipv6_cidr_block and not isinstance(ipv6_cidr_block, str):
            raise TypeError('Expected argument ipv6_cidr_block to be a str')
        __self__.ipv6_cidr_block = ipv6_cidr_block
        if ipv6_cidr_block_association_id and not isinstance(ipv6_cidr_block_association_id, str):
            raise TypeError('Expected argument ipv6_cidr_block_association_id to be a str')
        __self__.ipv6_cidr_block_association_id = ipv6_cidr_block_association_id
        if map_public_ip_on_launch and not isinstance(map_public_ip_on_launch, bool):
            raise TypeError('Expected argument map_public_ip_on_launch to be a bool')
        __self__.map_public_ip_on_launch = map_public_ip_on_launch
        if state and not isinstance(state, str):
            raise TypeError('Expected argument state to be a str')
        __self__.state = state
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError('Expected argument vpc_id to be a str')
        __self__.vpc_id = vpc_id

def get_subnet(availability_zone=None, cidr_block=None, default_for_az=None, filters=None, id=None, ipv6_cidr_block=None, state=None, tags=None, vpc_id=None):
    """
    `aws_subnet` provides details about a specific VPC subnet.
    
    This resource can prove useful when a module accepts a subnet id as
    an input variable and needs to, for example, determine the id of the
    VPC that the subnet belongs to.
    """
    __args__ = dict()

    __args__['availabilityZone'] = availability_zone
    __args__['cidrBlock'] = cidr_block
    __args__['defaultForAz'] = default_for_az
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['ipv6CidrBlock'] = ipv6_cidr_block
    __args__['state'] = state
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __ret__ = pulumi.runtime.invoke('aws:ec2/getSubnet:getSubnet', __args__)

    return GetSubnetResult(
        arn=__ret__.get('arn'),
        assign_ipv6_address_on_creation=__ret__.get('assignIpv6AddressOnCreation'),
        availability_zone=__ret__.get('availabilityZone'),
        cidr_block=__ret__.get('cidrBlock'),
        default_for_az=__ret__.get('defaultForAz'),
        id=__ret__.get('id'),
        ipv6_cidr_block=__ret__.get('ipv6CidrBlock'),
        ipv6_cidr_block_association_id=__ret__.get('ipv6CidrBlockAssociationId'),
        map_public_ip_on_launch=__ret__.get('mapPublicIpOnLaunch'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'))
