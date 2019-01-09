# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpcPeeringConnectionAccepter(pulumi.CustomResource):
    """
    Provides a resource to manage the accepter's side of a VPC Peering Connection.
    
    When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
    VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
    accepter's account.
    The requester can use the `aws_vpc_peering_connection` resource to manage its side of the connection
    and the accepter can use the `aws_vpc_peering_connection_accepter` resource to "adopt" its side of the
    connection into management.
    """
    def __init__(__self__, __name__, __opts__=None, accepter=None, auto_accept=None, requester=None, tags=None, vpc_peering_connection_id=None):
        """Create a VpcPeeringConnectionAccepter resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['accepter'] = accepter

        __props__['auto_accept'] = auto_accept

        __props__['requester'] = requester

        __props__['tags'] = tags

        if not vpc_peering_connection_id:
            raise TypeError('Missing required property vpc_peering_connection_id')
        __props__['vpc_peering_connection_id'] = vpc_peering_connection_id

        __props__['accept_status'] = None
        __props__['peer_owner_id'] = None
        __props__['peer_region'] = None
        __props__['peer_vpc_id'] = None
        __props__['vpc_id'] = None

        super(VpcPeeringConnectionAccepter, __self__).__init__(
            'aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

