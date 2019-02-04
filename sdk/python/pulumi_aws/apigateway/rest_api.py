# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RestApi(pulumi.CustomResource):
    api_key_source: pulumi.Output[str]
    """
    The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
    """
    binary_media_types: pulumi.Output[list]
    """
    The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
    """
    body: pulumi.Output[str]
    """
    An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
    """
    created_date: pulumi.Output[str]
    """
    The creation date of the REST API
    """
    description: pulumi.Output[str]
    """
    The description of the REST API
    """
    endpoint_configuration: pulumi.Output[dict]
    """
    Nested argument defining API endpoint configuration including endpoint type. Defined below.
    """
    execution_arn: pulumi.Output[str]
    """
    The execution ARN part to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn`
    when allowing API Gateway to invoke a Lambda function,
    e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
    """
    minimum_compression_size: pulumi.Output[int]
    """
    Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
    """
    name: pulumi.Output[str]
    """
    The name of the REST API
    """
    policy: pulumi.Output[str]
    """
    JSON formatted policy document that controls access to the API Gateway. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
    """
    root_resource_id: pulumi.Output[str]
    """
    The resource ID of the REST API's root
    """
    def __init__(__self__, __name__, __opts__=None, api_key_source=None, binary_media_types=None, body=None, description=None, endpoint_configuration=None, minimum_compression_size=None, name=None, policy=None):
        """
        Provides an API Gateway REST API.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] api_key_source: The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
        :param pulumi.Input[list] binary_media_types: The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
        :param pulumi.Input[str] body: An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
        :param pulumi.Input[str] description: The description of the REST API
        :param pulumi.Input[dict] endpoint_configuration: Nested argument defining API endpoint configuration including endpoint type. Defined below.
        :param pulumi.Input[int] minimum_compression_size: Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
        :param pulumi.Input[str] name: The name of the REST API
        :param pulumi.Input[str] policy: JSON formatted policy document that controls access to the API Gateway. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['api_key_source'] = api_key_source

        __props__['binary_media_types'] = binary_media_types

        __props__['body'] = body

        __props__['description'] = description

        __props__['endpoint_configuration'] = endpoint_configuration

        __props__['minimum_compression_size'] = minimum_compression_size

        __props__['name'] = name

        __props__['policy'] = policy

        __props__['created_date'] = None
        __props__['execution_arn'] = None
        __props__['root_resource_id'] = None

        super(RestApi, __self__).__init__(
            'aws:apigateway/restApi:RestApi',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

