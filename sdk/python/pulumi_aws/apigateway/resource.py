# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Resource(pulumi.CustomResource):
    parent_id: pulumi.Output[str]
    """
    The ID of the parent API resource
    """
    path: pulumi.Output[str]
    """
    The complete path for this API resource, including all parent paths.
    """
    path_part: pulumi.Output[str]
    """
    The last path segment of this API resource.
    """
    rest_api: pulumi.Output[str]
    """
    The ID of the associated REST API
    """
    def __init__(__self__, resource_name, opts=None, parent_id=None, path_part=None, rest_api=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent_id: The ID of the parent API resource
        :param pulumi.Input[str] path_part: The last path segment of this API resource.
        :param pulumi.Input[str] rest_api: The ID of the associated REST API

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_resource.html.markdown.
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

        if parent_id is None:
            raise TypeError("Missing required property 'parent_id'")
        __props__['parent_id'] = parent_id

        if path_part is None:
            raise TypeError("Missing required property 'path_part'")
        __props__['path_part'] = path_part

        if rest_api is None:
            raise TypeError("Missing required property 'rest_api'")
        __props__['rest_api'] = rest_api

        __props__['path'] = None

        super(Resource, __self__).__init__(
            'aws:apigateway/resource:Resource',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

