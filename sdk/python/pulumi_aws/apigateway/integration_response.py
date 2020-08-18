# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['IntegrationResponse']


class IntegrationResponse(pulumi.CustomResource):
    content_handling: pulumi.Output[Optional[str]] = pulumi.property("contentHandling")
    """
    Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
    """

    http_method: pulumi.Output[str] = pulumi.property("httpMethod")
    """
    The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
    """

    resource_id: pulumi.Output[str] = pulumi.property("resourceId")
    """
    The API resource ID
    """

    response_parameters: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("responseParameters")
    """
    A map of response parameters that can be read from the backend response.
    For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
    """

    response_templates: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("responseTemplates")
    """
    A map specifying the templates used to transform the integration response body
    """

    rest_api: pulumi.Output[str] = pulumi.property("restApi")
    """
    The ID of the associated REST API
    """

    selection_pattern: pulumi.Output[Optional[str]] = pulumi.property("selectionPattern")
    """
    Specifies the regular expression pattern used to choose
    an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
    If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
    For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
    """

    status_code: pulumi.Output[str] = pulumi.property("statusCode")
    """
    The HTTP status code
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_handling: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 response_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 response_templates: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 selection_pattern: Optional[pulumi.Input[str]] = None,
                 status_code: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an HTTP Method Integration Response for an API Gateway Resource.

        > **Note:** Depends on having `apigateway.Integration` inside your rest api. To ensure this
        you might need to add an explicit `depends_on` for clean runs.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
        my_demo_resource = aws.apigateway.Resource("myDemoResource",
            rest_api=my_demo_api.id,
            parent_id=my_demo_api.root_resource_id,
            path_part="mydemoresource")
        my_demo_method = aws.apigateway.Method("myDemoMethod",
            rest_api=my_demo_api.id,
            resource_id=my_demo_resource.id,
            http_method="GET",
            authorization="NONE")
        my_demo_integration = aws.apigateway.Integration("myDemoIntegration",
            rest_api=my_demo_api.id,
            resource_id=my_demo_resource.id,
            http_method=my_demo_method.http_method,
            type="MOCK")
        response200 = aws.apigateway.MethodResponse("response200",
            rest_api=my_demo_api.id,
            resource_id=my_demo_resource.id,
            http_method=my_demo_method.http_method,
            status_code="200")
        my_demo_integration_response = aws.apigateway.IntegrationResponse("myDemoIntegrationResponse",
            rest_api=my_demo_api.id,
            resource_id=my_demo_resource.id,
            http_method=my_demo_method.http_method,
            status_code=response200.status_code,
            response_templates={
                "application/xml": \"\"\"#set($inputRoot = $input.path('$'))
        <?xml version="1.0" encoding="UTF-8"?>
        <message>
            $inputRoot.body
        </message>
        \"\"\",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_parameters: A map of response parameters that can be read from the backend response.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_templates: A map specifying the templates used to transform the integration response body
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] selection_pattern: Specifies the regular expression pattern used to choose
               an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
               If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
               For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['content_handling'] = content_handling
            if http_method is None:
                raise TypeError("Missing required property 'http_method'")
            __props__['http_method'] = http_method
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['response_parameters'] = response_parameters
            __props__['response_templates'] = response_templates
            if rest_api is None:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            __props__['selection_pattern'] = selection_pattern
            if status_code is None:
                raise TypeError("Missing required property 'status_code'")
            __props__['status_code'] = status_code
        super(IntegrationResponse, __self__).__init__(
            'aws:apigateway/integrationResponse:IntegrationResponse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            content_handling: Optional[pulumi.Input[str]] = None,
            http_method: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            response_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            response_templates: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            rest_api: Optional[pulumi.Input[str]] = None,
            selection_pattern: Optional[pulumi.Input[str]] = None,
            status_code: Optional[pulumi.Input[str]] = None) -> 'IntegrationResponse':
        """
        Get an existing IntegrationResponse resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_parameters: A map of response parameters that can be read from the backend response.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_templates: A map specifying the templates used to transform the integration response body
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] selection_pattern: Specifies the regular expression pattern used to choose
               an integration response based on the response from the backend. Setting this to `-` makes the integration the default one.
               If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
               For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        :param pulumi.Input[str] status_code: The HTTP status code
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_handling"] = content_handling
        __props__["http_method"] = http_method
        __props__["resource_id"] = resource_id
        __props__["response_parameters"] = response_parameters
        __props__["response_templates"] = response_templates
        __props__["rest_api"] = rest_api
        __props__["selection_pattern"] = selection_pattern
        __props__["status_code"] = status_code
        return IntegrationResponse(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

