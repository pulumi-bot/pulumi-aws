# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class BasePathMapping(pulumi.CustomResource):
    """
    Connects a custom domain name registered via `aws_api_gateway_domain_name`
    with a deployed API so that its methods can be called via the
    custom domain name.
    """
    def __init__(__self__, __name__, __opts__=None, rest_api=None, base_path=None, domain_name=None, stage_name=None):
        """Create a BasePathMapping resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not rest_api:
            raise TypeError('Missing required property rest_api')
        __props__['rest_api'] = rest_api

        __props__['base_path'] = base_path

        if not domain_name:
            raise TypeError('Missing required property domain_name')
        __props__['domain_name'] = domain_name

        __props__['stage_name'] = stage_name

        super(BasePathMapping, __self__).__init__(
            'aws:apigateway/basePathMapping:BasePathMapping',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

