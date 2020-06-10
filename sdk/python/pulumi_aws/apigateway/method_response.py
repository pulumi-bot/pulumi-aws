# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class MethodResponse(pulumi.CustomResource):
    http_method: pulumi.Output[str]
    """
    The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
    """
    resource_id: pulumi.Output[str]
    """
    The API resource ID
    """
    response_models: pulumi.Output[dict]
    """
    A map of the API models used for the response's content type
    """
    response_parameters: pulumi.Output[dict]
    """
    A map of response parameters that can be sent to the caller.
    For example: `response_parameters = { "method.response.header.X-Some-Header" = true }`
    would define that the header `X-Some-Header` can be provided on the response.
    """
    rest_api: pulumi.Output[str]
    """
    The ID of the associated REST API
    """
    status_code: pulumi.Output[str]
    """
    The HTTP status code
    """
    def __init__(__self__, resource_name, opts=None, http_method=None, resource_id=None, response_models=None, response_parameters=None, rest_api=None, status_code=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an HTTP Method Response for an API Gateway Resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
        my_demo_resource = aws.apigateway.Resource("myDemoResource",
            parent_id=my_demo_api.root_resource_id,
            path_part="mydemoresource",
            rest_api=my_demo_api.id)
        my_demo_method = aws.apigateway.Method("myDemoMethod",
            authorization="NONE",
            http_method="GET",
            resource_id=my_demo_resource.id,
            rest_api=my_demo_api.id)
        my_demo_integration = aws.apigateway.Integration("myDemoIntegration",
            http_method=my_demo_method.http_method,
            resource_id=my_demo_resource.id,
            rest_api=my_demo_api.id,
            type="MOCK")
        response200 = aws.apigateway.MethodResponse("response200",
            http_method=my_demo_method.http_method,
            resource_id=my_demo_resource.id,
            rest_api=my_demo_api.id,
            status_code="200")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_method: The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[dict] response_models: A map of the API models used for the response's content type
        :param pulumi.Input[dict] response_parameters: A map of response parameters that can be sent to the caller.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = true }`
               would define that the header `X-Some-Header` can be provided on the response.
        :param pulumi.Input[dict] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] status_code: The HTTP status code
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if http_method is None:
                raise TypeError("Missing required property 'http_method'")
            __props__['http_method'] = http_method
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['response_models'] = response_models
            __props__['response_parameters'] = response_parameters
            if rest_api is None:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            if status_code is None:
                raise TypeError("Missing required property 'status_code'")
            __props__['status_code'] = status_code
        super(MethodResponse, __self__).__init__(
            'aws:apigateway/methodResponse:MethodResponse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, http_method=None, resource_id=None, response_models=None, response_parameters=None, rest_api=None, status_code=None):
        """
        Get an existing MethodResponse resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] http_method: The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[dict] response_models: A map of the API models used for the response's content type
        :param pulumi.Input[dict] response_parameters: A map of response parameters that can be sent to the caller.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = true }`
               would define that the header `X-Some-Header` can be provided on the response.
        :param pulumi.Input[dict] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] status_code: The HTTP status code
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["http_method"] = http_method
        __props__["resource_id"] = resource_id
        __props__["response_models"] = response_models
        __props__["response_parameters"] = response_parameters
        __props__["rest_api"] = rest_api
        __props__["status_code"] = status_code
        return MethodResponse(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
