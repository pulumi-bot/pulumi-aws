# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['VpcDhcpOptionsAssociation']


class VpcDhcpOptionsAssociation(pulumi.CustomResource):
    dhcp_options_id: pulumi.Output[str] = pulumi.property("dhcpOptionsId")
    """
    The ID of the DHCP Options Set to associate to the VPC.
    """

    vpc_id: pulumi.Output[str] = pulumi.property("vpcId")
    """
    The ID of the VPC to which we would like to associate a DHCP Options Set.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp_options_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a VPC DHCP Options Association resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        dns_resolver = aws.ec2.VpcDhcpOptionsAssociation("dnsResolver",
            vpc_id=aws_vpc["foo"]["id"],
            dhcp_options_id=aws_vpc_dhcp_options["foo"]["id"])
        ```
        ## Remarks

        * You can only associate one DHCP Options Set to a given VPC ID.
        * Removing the DHCP Options Association automatically sets AWS's `default` DHCP Options Set to the VPC.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dhcp_options_id: The ID of the DHCP Options Set to associate to the VPC.
        :param pulumi.Input[str] vpc_id: The ID of the VPC to which we would like to associate a DHCP Options Set.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if dhcp_options_id is None:
                raise TypeError("Missing required property 'dhcp_options_id'")
            __props__['dhcp_options_id'] = dhcp_options_id
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
        super(VpcDhcpOptionsAssociation, __self__).__init__(
            'aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            dhcp_options_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'VpcDhcpOptionsAssociation':
        """
        Get an existing VpcDhcpOptionsAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dhcp_options_id: The ID of the DHCP Options Set to associate to the VPC.
        :param pulumi.Input[str] vpc_id: The ID of the VPC to which we would like to associate a DHCP Options Set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dhcp_options_id"] = dhcp_options_id
        __props__["vpc_id"] = vpc_id
        return VpcDhcpOptionsAssociation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

