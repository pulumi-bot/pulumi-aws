# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['IntegrationResponseArgs', 'IntegrationResponse']

@pulumi.input_type
class IntegrationResponseArgs:
    def __init__(__self__, *,
                 http_method: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 rest_api: pulumi.Input[str],
                 status_code: pulumi.Input[str],
                 content_handling: Optional[pulumi.Input[str]] = None,
                 response_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 response_templates: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 selection_pattern: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IntegrationResponse resource.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] status_code: The HTTP status code
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_parameters: A map of response parameters that can be read from the backend response.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_templates: A map specifying the templates used to transform the integration response body
        :param pulumi.Input[str] selection_pattern: Specifies the regular expression pattern used to choose
               an integration response based on the response from the backend. Omit configuring this to make the integration the default one.
               If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
               For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        """
        pulumi.set(__self__, "http_method", http_method)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "rest_api", rest_api)
        pulumi.set(__self__, "status_code", status_code)
        if content_handling is not None:
            pulumi.set(__self__, "content_handling", content_handling)
        if response_parameters is not None:
            pulumi.set(__self__, "response_parameters", response_parameters)
        if response_templates is not None:
            pulumi.set(__self__, "response_templates", response_templates)
        if selection_pattern is not None:
            pulumi.set(__self__, "selection_pattern", selection_pattern)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Input[str]:
        """
        The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The API resource ID
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Input[str]:
        """
        The ID of the associated REST API
        """
        return pulumi.get(self, "rest_api")

    @rest_api.setter
    def rest_api(self, value: pulumi.Input[str]):
        pulumi.set(self, "rest_api", value)

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> pulumi.Input[str]:
        """
        The HTTP status code
        """
        return pulumi.get(self, "status_code")

    @status_code.setter
    def status_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "status_code", value)

    @property
    @pulumi.getter(name="contentHandling")
    def content_handling(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        """
        return pulumi.get(self, "content_handling")

    @content_handling.setter
    def content_handling(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_handling", value)

    @property
    @pulumi.getter(name="responseParameters")
    def response_parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of response parameters that can be read from the backend response.
        For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        """
        return pulumi.get(self, "response_parameters")

    @response_parameters.setter
    def response_parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "response_parameters", value)

    @property
    @pulumi.getter(name="responseTemplates")
    def response_templates(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map specifying the templates used to transform the integration response body
        """
        return pulumi.get(self, "response_templates")

    @response_templates.setter
    def response_templates(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "response_templates", value)

    @property
    @pulumi.getter(name="selectionPattern")
    def selection_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the regular expression pattern used to choose
        an integration response based on the response from the backend. Omit configuring this to make the integration the default one.
        If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
        For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        """
        return pulumi.get(self, "selection_pattern")

    @selection_pattern.setter
    def selection_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "selection_pattern", value)


class IntegrationResponse(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
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

        ## Import

        `aws_api_gateway_integration_response` can be imported using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`, e.g.

        ```sh
         $ pulumi import aws:apigateway/integrationResponse:IntegrationResponse example 12345abcde/67890fghij/GET/200
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
               an integration response based on the response from the backend. Omit configuring this to make the integration the default one.
               If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
               For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        :param pulumi.Input[str] status_code: The HTTP status code
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationResponseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        `aws_api_gateway_integration_response` can be imported using `REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE`, e.g.

        ```sh
         $ pulumi import aws:apigateway/integrationResponse:IntegrationResponse example 12345abcde/67890fghij/GET/200
        ```

        :param str resource_name: The name of the resource.
        :param IntegrationResponseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationResponseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
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
            if http_method is None and not opts.urn:
                raise TypeError("Missing required property 'http_method'")
            __props__['http_method'] = http_method
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['response_parameters'] = response_parameters
            __props__['response_templates'] = response_templates
            if rest_api is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            __props__['selection_pattern'] = selection_pattern
            if status_code is None and not opts.urn:
                raise TypeError("Missing required property 'status_code'")
            __props__['status_code'] = status_code
        super(IntegrationResponse, __self__).__init__(
            'aws:apigateway/integrationResponse:IntegrationResponse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
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
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        :param pulumi.Input[str] resource_id: The API resource ID
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_parameters: A map of response parameters that can be read from the backend response.
               For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] response_templates: A map specifying the templates used to transform the integration response body
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] selection_pattern: Specifies the regular expression pattern used to choose
               an integration response based on the response from the backend. Omit configuring this to make the integration the default one.
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

    @property
    @pulumi.getter(name="contentHandling")
    def content_handling(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
        """
        return pulumi.get(self, "content_handling")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Output[str]:
        """
        The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The API resource ID
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="responseParameters")
    def response_parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of response parameters that can be read from the backend response.
        For example: `response_parameters = { "method.response.header.X-Some-Header" = "integration.response.header.X-Some-Other-Header" }`
        """
        return pulumi.get(self, "response_parameters")

    @property
    @pulumi.getter(name="responseTemplates")
    def response_templates(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map specifying the templates used to transform the integration response body
        """
        return pulumi.get(self, "response_templates")

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Output[str]:
        """
        The ID of the associated REST API
        """
        return pulumi.get(self, "rest_api")

    @property
    @pulumi.getter(name="selectionPattern")
    def selection_pattern(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the regular expression pattern used to choose
        an integration response based on the response from the backend. Omit configuring this to make the integration the default one.
        If the backend is an `AWS` Lambda function, the AWS Lambda function error header is matched.
        For all other `HTTP` and `AWS` backends, the HTTP status code is matched.
        """
        return pulumi.get(self, "selection_pattern")

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> pulumi.Output[str]:
        """
        The HTTP status code
        """
        return pulumi.get(self, "status_code")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

