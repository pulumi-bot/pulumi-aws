# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class DocumentationPart(pulumi.CustomResource):
    location: pulumi.Output[dict]
    """
    The location of the targeted API entity of the to-be-created documentation part. See below.

      * `method` (`str`) - The HTTP verb of a method. The default value is `*` for any method.
      * `name` (`str`) - The name of the targeted API entity.
      * `path` (`str`) - The URL path of the target. The default value is `/` for the root resource.
      * `status_code` (`str`) - The HTTP status code of a response. The default value is `*` for any status code.
      * `type` (`str`) - The type of API entity to which the documentation content applies. e.g. `API`, `METHOD` or `REQUEST_BODY`
    """
    properties: pulumi.Output[str]
    """
    A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
    """
    rest_api_id: pulumi.Output[str]
    """
    The ID of the associated Rest API
    """
    def __init__(__self__, resource_name, opts=None, location=None, properties=None, rest_api_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a settings of an API Gateway Documentation Part.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rest_api = aws.apigateway.RestApi("exampleRestApi")
        example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
            location={
                "method": "GET",
                "path": "/example",
                "type": "METHOD",
            },
            properties="{\"description\":\"Example description\"}",
            rest_api_id=example_rest_api.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] location: The location of the targeted API entity of the to-be-created documentation part. See below.
        :param pulumi.Input[str] properties: A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API

        The **location** object supports the following:

          * `method` (`pulumi.Input[str]`) - The HTTP verb of a method. The default value is `*` for any method.
          * `name` (`pulumi.Input[str]`) - The name of the targeted API entity.
          * `path` (`pulumi.Input[str]`) - The URL path of the target. The default value is `/` for the root resource.
          * `status_code` (`pulumi.Input[str]`) - The HTTP status code of a response. The default value is `*` for any status code.
          * `type` (`pulumi.Input[str]`) - The type of API entity to which the documentation content applies. e.g. `API`, `METHOD` or `REQUEST_BODY`
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if rest_api_id is None:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__['rest_api_id'] = rest_api_id
        super(DocumentationPart, __self__).__init__(
            'aws:apigateway/documentationPart:DocumentationPart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location=None, properties=None, rest_api_id=None):
        """
        Get an existing DocumentationPart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] location: The location of the targeted API entity of the to-be-created documentation part. See below.
        :param pulumi.Input[str] properties: A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        :param pulumi.Input[str] rest_api_id: The ID of the associated Rest API

        The **location** object supports the following:

          * `method` (`pulumi.Input[str]`) - The HTTP verb of a method. The default value is `*` for any method.
          * `name` (`pulumi.Input[str]`) - The name of the targeted API entity.
          * `path` (`pulumi.Input[str]`) - The URL path of the target. The default value is `/` for the root resource.
          * `status_code` (`pulumi.Input[str]`) - The HTTP status code of a response. The default value is `*` for any status code.
          * `type` (`pulumi.Input[str]`) - The type of API entity to which the documentation content applies. e.g. `API`, `METHOD` or `REQUEST_BODY`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location"] = location
        __props__["properties"] = properties
        __props__["rest_api_id"] = rest_api_id
        return DocumentationPart(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
