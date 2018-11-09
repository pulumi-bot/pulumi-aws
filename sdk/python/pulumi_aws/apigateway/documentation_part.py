# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class DocumentationPart(pulumi.CustomResource):
    """
    Provides a settings of an API Gateway Documentation Part.
    """
    def __init__(__self__, __name__, __opts__=None, location=None, properties=None, rest_api_id=None):
        """Create a DocumentationPart resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        if not properties:
            raise TypeError('Missing required property properties')
        __props__['properties'] = properties

        if not rest_api_id:
            raise TypeError('Missing required property rest_api_id')
        __props__['rest_api_id'] = rest_api_id

        super(DocumentationPart, __self__).__init__(
            'aws:apigateway/documentationPart:DocumentationPart',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

