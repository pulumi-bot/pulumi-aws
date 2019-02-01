# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Eip(pulumi.CustomResource):
    allocation_id: pulumi.Output[str]
    associate_with_private_ip: pulumi.Output[str]
    """
    A user specified primary or secondary private IP address to
    associate with the Elastic IP address. If no private IP address is specified,
    the Elastic IP address is associated with the primary private IP address.
    """
    association_id: pulumi.Output[str]
    domain: pulumi.Output[str]
    instance: pulumi.Output[str]
    """
    EC2 instance ID.
    """
    network_interface: pulumi.Output[str]
    """
    Network interface ID to associate with.
    """
    private_ip: pulumi.Output[str]
    """
    Contains the private IP address (if in VPC).
    """
    public_ip: pulumi.Output[str]
    """
    Contains the public IP address.
    """
    public_ipv4_pool: pulumi.Output[str]
    """
    EC2 IPv4 address pool identifier or `amazon`. This option is only available for VPC EIPs.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc: pulumi.Output[bool]
    """
    Boolean if the EIP is in a VPC or not.
    """
    def __init__(__self__, __name__, __opts__=None, associate_with_private_ip=None, instance=None, network_interface=None, public_ipv4_pool=None, tags=None, vpc=None):
        """
        Provides an Elastic IP resource.
        
        > **Note:** EIP may require IGW to exist prior to association. Use `depends_on` to set an explicit dependency on the IGW.
        
        > **Note:** Do not use `network_interface` to associate the EIP to `aws_lb` or `aws_nat_gateway` resources. Instead use the `allocation_id` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] associate_with_private_ip: A user specified primary or secondary private IP address to
               associate with the Elastic IP address. If no private IP address is specified,
               the Elastic IP address is associated with the primary private IP address.
        :param pulumi.Input[str] instance: EC2 instance ID.
        :param pulumi.Input[str] network_interface: Network interface ID to associate with.
        :param pulumi.Input[str] public_ipv4_pool: EC2 IPv4 address pool identifier or `amazon`. This option is only available for VPC EIPs.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[bool] vpc: Boolean if the EIP is in a VPC or not.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['associate_with_private_ip'] = associate_with_private_ip

        __props__['instance'] = instance

        __props__['network_interface'] = network_interface

        __props__['public_ipv4_pool'] = public_ipv4_pool

        __props__['tags'] = tags

        __props__['vpc'] = vpc

        __props__['allocation_id'] = None
        __props__['association_id'] = None
        __props__['domain'] = None
        __props__['private_ip'] = None
        __props__['public_ip'] = None

        super(Eip, __self__).__init__(
            'aws:ec2/eip:Eip',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

