# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NatGateway(pulumi.CustomResource):
    allocation_id: pulumi.Output[str]
    """
    The Allocation ID of the Elastic IP address for the gateway.
    """
    network_interface_id: pulumi.Output[str]
    """
    The ENI ID of the network interface created by the NAT gateway.
    """
    private_ip: pulumi.Output[str]
    """
    The private IP address of the NAT Gateway.
    """
    public_ip: pulumi.Output[str]
    """
    The public IP address of the NAT Gateway.
    """
    subnet_id: pulumi.Output[str]
    """
    The Subnet ID of the subnet in which to place the gateway.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, allocation_id=None, subnet_id=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a VPC NAT Gateway.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_id: The Allocation ID of the Elastic IP address for the gateway.
        :param pulumi.Input[str] subnet_id: The Subnet ID of the subnet in which to place the gateway.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if allocation_id is None:
            raise TypeError("Missing required property 'allocation_id'")
        __props__['allocation_id'] = allocation_id

        if subnet_id is None:
            raise TypeError("Missing required property 'subnet_id'")
        __props__['subnet_id'] = subnet_id

        __props__['tags'] = tags

        __props__['network_interface_id'] = None
        __props__['private_ip'] = None
        __props__['public_ip'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(NatGateway, __self__).__init__(
            'aws:ec2/natGateway:NatGateway',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

