# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Integration']


class Integration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_key_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cache_namespace: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 connection_type: Optional[pulumi.Input[str]] = None,
                 content_handling: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 integration_http_method: Optional[pulumi.Input[str]] = None,
                 passthrough_behavior: Optional[pulumi.Input[str]] = None,
                 request_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_templates: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 timeout_milliseconds: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an HTTP Method Integration for an API Gateway Integration.

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
            type="MOCK",
            cache_key_parameters=["method.request.path.param"],
            cache_namespace="foobar",
            timeout_milliseconds=29000,
            request_parameters={
                "integration.request.header.X-Authorization": "'static'",
            },
            request_templates={
                "application/xml": \"\"\"{
           "body" : $input.json('$')
        }
        \"\"\",
            })
        ```
        ## VPC Link

        ```python
        import pulumi
        import pulumi_aws as aws

        config = pulumi.Config()
        name = config.require_object("name")
        subnet_id = config.require_object("subnetId")
        test_load_balancer = aws.lb.LoadBalancer("testLoadBalancer",
            internal=True,
            load_balancer_type="network",
            subnets=[subnet_id])
        test_vpc_link = aws.apigateway.VpcLink("testVpcLink", target_arn=[test_load_balancer.arn])
        test_rest_api = aws.apigateway.RestApi("testRestApi")
        test_resource = aws.apigateway.Resource("testResource",
            rest_api=test_rest_api.id,
            parent_id=test_rest_api.root_resource_id,
            path_part="test")
        test_method = aws.apigateway.Method("testMethod",
            rest_api=test_rest_api.id,
            resource_id=test_resource.id,
            http_method="GET",
            authorization="NONE",
            request_models={
                "application/json": "Error",
            })
        test_integration = aws.apigateway.Integration("testIntegration",
            rest_api=test_rest_api.id,
            resource_id=test_resource.id,
            http_method=test_method.http_method,
            request_templates={
                "application/json": "",
                "application/xml": \"\"\"#set($inputRoot = $input.path('$'))
        { }\"\"\",
            },
            request_parameters={
                "integration.request.header.X-Authorization": "'static'",
                "integration.request.header.X-Foo": "'Bar'",
            },
            type="HTTP",
            uri="https://www.google.de",
            integration_http_method="GET",
            passthrough_behavior="WHEN_NO_MATCH",
            content_handling="CONVERT_TO_TEXT",
            connection_type="VPC_LINK",
            connection_id=test_vpc_link.id)
        ```

        ## Import

        `aws_api_gateway_integration` can be imported using `REST-API-ID/RESOURCE-ID/HTTP-METHOD`, e.g.

        ```sh
         $ pulumi import aws:apigateway/integration:Integration example 12345abcde/67890fghij/GET
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cache_key_parameters: A list of cache key parameters for the integration.
        :param pulumi.Input[str] cache_namespace: The integration's cache namespace.
        :param pulumi.Input[str] connection_id: The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
        :param pulumi.Input[str] connection_type: The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
        :param pulumi.Input[str] credentials: The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
               when calling the associated resource.
        :param pulumi.Input[str] integration_http_method: The integration HTTP method
               (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
               **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
               Not all methods are compatible with all `AWS` integrations.
               e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
        :param pulumi.Input[str] passthrough_behavior: The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] request_parameters: A map of request query string parameters and headers that should be passed to the backend responder.
               For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] request_templates: A map of the integration's request templates.
        :param pulumi.Input[str] resource_id: The API resource ID.
        :param pulumi.Input[str] rest_api: The ID of the associated REST API.
        :param pulumi.Input[int] timeout_milliseconds: Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
        :param pulumi.Input[str] type: The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
        :param pulumi.Input[str] uri: The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
               For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
               e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
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

            __props__['cache_key_parameters'] = cache_key_parameters
            __props__['cache_namespace'] = cache_namespace
            __props__['connection_id'] = connection_id
            __props__['connection_type'] = connection_type
            __props__['content_handling'] = content_handling
            __props__['credentials'] = credentials
            if http_method is None:
                raise TypeError("Missing required property 'http_method'")
            __props__['http_method'] = http_method
            __props__['integration_http_method'] = integration_http_method
            __props__['passthrough_behavior'] = passthrough_behavior
            __props__['request_parameters'] = request_parameters
            __props__['request_templates'] = request_templates
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            if rest_api is None:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            __props__['timeout_milliseconds'] = timeout_milliseconds
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['uri'] = uri
        super(Integration, __self__).__init__(
            'aws:apigateway/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cache_key_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cache_namespace: Optional[pulumi.Input[str]] = None,
            connection_id: Optional[pulumi.Input[str]] = None,
            connection_type: Optional[pulumi.Input[str]] = None,
            content_handling: Optional[pulumi.Input[str]] = None,
            credentials: Optional[pulumi.Input[str]] = None,
            http_method: Optional[pulumi.Input[str]] = None,
            integration_http_method: Optional[pulumi.Input[str]] = None,
            passthrough_behavior: Optional[pulumi.Input[str]] = None,
            request_parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            request_templates: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            rest_api: Optional[pulumi.Input[str]] = None,
            timeout_milliseconds: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cache_key_parameters: A list of cache key parameters for the integration.
        :param pulumi.Input[str] cache_namespace: The integration's cache namespace.
        :param pulumi.Input[str] connection_id: The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
        :param pulumi.Input[str] connection_type: The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
        :param pulumi.Input[str] content_handling: Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
        :param pulumi.Input[str] credentials: The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
        :param pulumi.Input[str] http_method: The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
               when calling the associated resource.
        :param pulumi.Input[str] integration_http_method: The integration HTTP method
               (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
               **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
               Not all methods are compatible with all `AWS` integrations.
               e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
        :param pulumi.Input[str] passthrough_behavior: The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] request_parameters: A map of request query string parameters and headers that should be passed to the backend responder.
               For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] request_templates: A map of the integration's request templates.
        :param pulumi.Input[str] resource_id: The API resource ID.
        :param pulumi.Input[str] rest_api: The ID of the associated REST API.
        :param pulumi.Input[int] timeout_milliseconds: Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
        :param pulumi.Input[str] type: The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
        :param pulumi.Input[str] uri: The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
               For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
               e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cache_key_parameters"] = cache_key_parameters
        __props__["cache_namespace"] = cache_namespace
        __props__["connection_id"] = connection_id
        __props__["connection_type"] = connection_type
        __props__["content_handling"] = content_handling
        __props__["credentials"] = credentials
        __props__["http_method"] = http_method
        __props__["integration_http_method"] = integration_http_method
        __props__["passthrough_behavior"] = passthrough_behavior
        __props__["request_parameters"] = request_parameters
        __props__["request_templates"] = request_templates
        __props__["resource_id"] = resource_id
        __props__["rest_api"] = rest_api
        __props__["timeout_milliseconds"] = timeout_milliseconds
        __props__["type"] = type
        __props__["uri"] = uri
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cacheKeyParameters")
    def cache_key_parameters(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of cache key parameters for the integration.
        """
        return pulumi.get(self, "cache_key_parameters")

    @property
    @pulumi.getter(name="cacheNamespace")
    def cache_namespace(self) -> pulumi.Output[str]:
        """
        The integration's cache namespace.
        """
        return pulumi.get(self, "cache_namespace")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of the VpcLink used for the integration. **Required** if `connection_type` is `VPC_LINK`
        """
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="connectionType")
    def connection_type(self) -> pulumi.Output[Optional[str]]:
        """
        The integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
        """
        return pulumi.get(self, "connection_type")

    @property
    @pulumi.getter(name="contentHandling")
    def content_handling(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies how to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
        """
        return pulumi.get(self, "content_handling")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional[str]]:
        """
        The credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\*:user/\*`.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Output[str]:
        """
        The HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
        when calling the associated resource.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter(name="integrationHttpMethod")
    def integration_http_method(self) -> pulumi.Output[Optional[str]]:
        """
        The integration HTTP method
        (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
        **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        Not all methods are compatible with all `AWS` integrations.
        e.g. Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
        """
        return pulumi.get(self, "integration_http_method")

    @property
    @pulumi.getter(name="passthroughBehavior")
    def passthrough_behavior(self) -> pulumi.Output[str]:
        """
        The integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  **Required** if `request_templates` is used.
        """
        return pulumi.get(self, "passthrough_behavior")

    @property
    @pulumi.getter(name="requestParameters")
    def request_parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of request query string parameters and headers that should be passed to the backend responder.
        For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
        """
        return pulumi.get(self, "request_parameters")

    @property
    @pulumi.getter(name="requestTemplates")
    def request_templates(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of the integration's request templates.
        """
        return pulumi.get(self, "request_templates")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The API resource ID.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Output[str]:
        """
        The ID of the associated REST API.
        """
        return pulumi.get(self, "rest_api")

    @property
    @pulumi.getter(name="timeoutMilliseconds")
    def timeout_milliseconds(self) -> pulumi.Output[Optional[int]]:
        """
        Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
        """
        return pulumi.get(self, "timeout_milliseconds")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[Optional[str]]:
        """
        The input's URI. **Required** if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
        For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
        e.g. `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
        """
        return pulumi.get(self, "uri")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

