# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class VpcDhcpOptionsAssociation(pulumi.CustomResource):
    """
    Provides a VPC DHCP Options Association resource.
    * Removing the DHCP Options Association automatically sets AWS's `default` DHCP Options Set to the VPC.
    """
    def __init__(__self__, __name__, __opts__=None, dhcp_options_id=None, vpc_id=None):
        """Create a VpcDhcpOptionsAssociation resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not dhcp_options_id:
            raise TypeError('Missing required property dhcp_options_id')
        elif not isinstance(dhcp_options_id, basestring):
            raise TypeError('Expected property dhcp_options_id to be a basestring')
        __self__.dhcp_options_id = dhcp_options_id
        """
        The ID of the DHCP Options Set to associate to the VPC.
        """
        __props__['dhcpOptionsId'] = dhcp_options_id

        if not vpc_id:
            raise TypeError('Missing required property vpc_id')
        elif not isinstance(vpc_id, basestring):
            raise TypeError('Expected property vpc_id to be a basestring')
        __self__.vpc_id = vpc_id
        """
        The ID of the VPC to which we would like to associate a DHCP Options Set.
        """
        __props__['vpcId'] = vpc_id

        super(VpcDhcpOptionsAssociation, __self__).__init__(
            'aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'dhcpOptionsId' in outs:
            self.dhcp_options_id = outs['dhcpOptionsId']
        if 'vpcId' in outs:
            self.vpc_id = outs['vpcId']
