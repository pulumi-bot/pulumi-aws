# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VpcIpv4CidrBlockAssociation(pulumi.CustomResource):
    cidr_block: pulumi.Output[str]
    """
    The additional IPv4 CIDR block to associate with the VPC.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the VPC to make the association with.
    """
    def __init__(__self__, resource_name, opts=None, cidr_block=None, vpc_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to associate additional IPv4 CIDR blocks with a VPC.

        When a VPC is created, a primary IPv4 CIDR block for the VPC must be specified.
        The `ec2.VpcIpv4CidrBlockAssociation` resource allows further IPv4 CIDR blocks to be added to the VPC.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        secondary_cidr = aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr",
            cidr_block="172.2.0.0/16",
            vpc_id=main.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The additional IPv4 CIDR block to associate with the VPC.
        :param pulumi.Input[str] vpc_id: The ID of the VPC to make the association with.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cidr_block is None:
                raise TypeError("Missing required property 'cidr_block'")
            __props__['cidr_block'] = cidr_block
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
        super(VpcIpv4CidrBlockAssociation, __self__).__init__(
            'aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cidr_block=None, vpc_id=None):
        """
        Get an existing VpcIpv4CidrBlockAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The additional IPv4 CIDR block to associate with the VPC.
        :param pulumi.Input[str] vpc_id: The ID of the VPC to make the association with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cidr_block"] = cidr_block
        __props__["vpc_id"] = vpc_id
        return VpcIpv4CidrBlockAssociation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
