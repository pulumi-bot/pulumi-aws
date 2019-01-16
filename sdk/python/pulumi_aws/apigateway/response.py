# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Response(pulumi.CustomResource):
    """
    Provides an API Gateway Gateway Response for a REST API Gateway.
    """
    def __init__(__self__, __name__, __opts__=None, response_parameters=None, response_templates=None, response_type=None, rest_api_id=None, status_code=None):
        """Create a Response resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['response_parameters'] = response_parameters

        __props__['response_templates'] = response_templates

        if not response_type:
            raise TypeError('Missing required property response_type')
        __props__['response_type'] = response_type

        if not rest_api_id:
            raise TypeError('Missing required property rest_api_id')
        __props__['rest_api_id'] = rest_api_id

        __props__['status_code'] = status_code

        super(Response, __self__).__init__(
            'aws:apigateway/response:Response',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

