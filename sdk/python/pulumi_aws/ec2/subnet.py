# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Subnet(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the subnet.
    """
    assign_ipv6_address_on_creation: pulumi.Output[bool]
    """
    Specify true to indicate
    that network interfaces created in the specified subnet should be
    assigned an IPv6 address. Default is `false`
    """
    availability_zone: pulumi.Output[str]
    """
    The AZ for the subnet.
    """
    availability_zone_id: pulumi.Output[str]
    """
    The AZ ID of the subnet.
    """
    cidr_block: pulumi.Output[str]
    """
    The CIDR block for the subnet.
    """
    ipv6_cidr_block: pulumi.Output[str]
    """
    The IPv6 network range for the subnet,
    in CIDR notation. The subnet size must use a /64 prefix length.
    """
    ipv6_cidr_block_association_id: pulumi.Output[str]
    """
    The association ID for the IPv6 CIDR block.
    """
    map_public_ip_on_launch: pulumi.Output[bool]
    """
    Specify true to indicate
    that instances launched into the subnet should be assigned
    a public IP address. Default is `false`.
    """
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the subnet.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The VPC ID.
    """
    def __init__(__self__, resource_name, opts=None, assign_ipv6_address_on_creation=None, availability_zone=None, availability_zone_id=None, cidr_block=None, ipv6_cidr_block=None, map_public_ip_on_launch=None, tags=None, vpc_id=None, __name__=None, __opts__=None):
        """
        Provides an VPC subnet resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] assign_ipv6_address_on_creation: Specify true to indicate
               that network interfaces created in the specified subnet should be
               assigned an IPv6 address. Default is `false`
        :param pulumi.Input[str] availability_zone: The AZ for the subnet.
        :param pulumi.Input[str] availability_zone_id: The AZ ID of the subnet.
        :param pulumi.Input[str] cidr_block: The CIDR block for the subnet.
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 network range for the subnet,
               in CIDR notation. The subnet size must use a /64 prefix length.
        :param pulumi.Input[bool] map_public_ip_on_launch: Specify true to indicate
               that instances launched into the subnet should be assigned
               a public IP address. Default is `false`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/subnet.html.markdown.
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

        __props__['assign_ipv6_address_on_creation'] = assign_ipv6_address_on_creation

        __props__['availability_zone'] = availability_zone

        __props__['availability_zone_id'] = availability_zone_id

        if cidr_block is None:
            raise TypeError("Missing required property 'cidr_block'")
        __props__['cidr_block'] = cidr_block

        __props__['ipv6_cidr_block'] = ipv6_cidr_block

        __props__['map_public_ip_on_launch'] = map_public_ip_on_launch

        __props__['tags'] = tags

        if vpc_id is None:
            raise TypeError("Missing required property 'vpc_id'")
        __props__['vpc_id'] = vpc_id

        __props__['arn'] = None
        __props__['ipv6_cidr_block_association_id'] = None
        __props__['owner_id'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Subnet, __self__).__init__(
            'aws:ec2/subnet:Subnet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

