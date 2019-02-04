# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IntegrationResponse(pulumi.CustomResource):
    content_handling: pulumi.Output[str]
    """
    Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
    """
    http_method: pulumi.Output[str]
    """
    The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
    """
    resource_id: pulumi.Output[str]
    """
    The API resource ID
    """
    response_parameters: pulumi.Output[dict]
    """
    A map of response parameters that can be read from the backend response.
    For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`,
    """
    response_parameters_in_json: pulumi.Output[str]
    """
    **Deprecated**, use `response_parameters` instead.
    """
    response_templates: pulumi.Output[dict]
    """
    A map specifying the templates used to transform the integration response body
    """
    rest_api: pulumi.Output[str]
    """
    The ID of the associated REST API
    """
    selection_pattern: pulumi.Output[str]
    """
    Specifies the regular expression pattern used to choose
    an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
    If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
    For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
    """
    status_code: pulumi.Output[str]
    """
    The HTTP status code
    """
    def __init__(__self__, __name__, __opts__=None, content_handling=None, http_method=None, resource_id=None, response_parameters=None, response_parameters_in_json=None, response_templates=None, rest_api=None, selection_pattern=None, status_code=None):
        """
        Provides an HTTP Method Integration Response for an API Gateway Resource.
        
        > **Note:** Depends on having `aws_api_gateway_integration` inside your rest api. To ensure this
        you might need to add an explicit `depends_on` for clean runs.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[dict] response_parameters: A map of response parameters that can be read from the backend response.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`,
        :param pulumi.Input[str] response_parameters_in_json: **Deprecated**, use `response_parameters` instead.
        :param pulumi.Input[dict] response_templates: A map specifying the templates used to transform the integration response body
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] selection_pattern: Specifies the regular expression pattern used to choose
               an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
               If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
               For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        :param pulumi.Input[str] status_code: The HTTP status code
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['content_handling'] = content_handling

        if http_method is None:
            raise TypeError('Missing required property http_method')
        __props__['http_method'] = http_method

        if resource_id is None:
            raise TypeError('Missing required property resource_id')
        __props__['resource_id'] = resource_id

        __props__['response_parameters'] = response_parameters

        __props__['response_parameters_in_json'] = response_parameters_in_json

        __props__['response_templates'] = response_templates

        if rest_api is None:
            raise TypeError('Missing required property rest_api')
        __props__['rest_api'] = rest_api

        __props__['selection_pattern'] = selection_pattern

        if status_code is None:
            raise TypeError('Missing required property status_code')
        __props__['status_code'] = status_code

        super(IntegrationResponse, __self__).__init__(
            'aws:apigateway/integrationResponse:IntegrationResponse',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

