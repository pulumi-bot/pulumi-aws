# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BasePathMapping(pulumi.CustomResource):
    rest_api: pulumi.Output[str]
    """
    The id of the API to connect.
    """
    base_path: pulumi.Output[str]
    """
    Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
    """
    domain_name: pulumi.Output[str]
    """
    The already-registered domain name to connect the API to.
    """
    stage_name: pulumi.Output[str]
    """
    The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
    """
    def __init__(__self__, __name__, __opts__=None, rest_api=None, base_path=None, domain_name=None, stage_name=None):
        """
        Connects a custom domain name registered via `aws_api_gateway_domain_name`
        with a deployed API so that its methods can be called via the
        custom domain name.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] rest_api: The id of the API to connect.
        :param pulumi.Input[str] base_path: Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
        :param pulumi.Input[str] domain_name: The already-registered domain name to connect the API to.
        :param pulumi.Input[str] stage_name: The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
        """
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

