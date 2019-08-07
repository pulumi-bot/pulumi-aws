# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EgressOnlyInternetGateway(pulumi.CustomResource):
    vpc_id: pulumi.Output[str]
    """
    The VPC ID to create in.
    """
    def __init__(__self__, resource_name, opts=None, vpc_id=None, __name__=None, __opts__=None):
        """
        [IPv6 only] Creates an egress-only Internet gateway for your VPC.
        An egress-only Internet gateway is used to enable outbound communication
        over IPv6 from instances in your VPC to the Internet, and prevents hosts
        outside of your VPC from initiating an IPv6 connection with your instance.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID to create in.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/egress_only_internet_gateway.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if vpc_id is None:
            raise TypeError("Missing required property 'vpc_id'")
        __props__['vpc_id'] = vpc_id
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(EgressOnlyInternetGateway, __self__).__init__(
            'aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

