# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Api']


class Api(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key_selection_expression: Optional[pulumi.Input[str]] = None,
                 cors_configuration: Optional[pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']]] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 route_key: Optional[pulumi.Input[str]] = None,
                 route_selection_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Amazon API Gateway Version 2 API.

        > **Note:** Amazon API Gateway Version 2 resources are used for creating and deploying WebSocket and HTTP APIs. To create and deploy REST APIs, use Amazon API Gateway Version 1.

        ## Example Usage
        ### Basic WebSocket API

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Api("example",
            protocol_type="WEBSOCKET",
            route_selection_expression="$request.body.action")
        ```
        ### Basic HTTP API

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Api("example", protocol_type="HTTP")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key_selection_expression: An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
               Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
               Applicable for WebSocket APIs.
        :param pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']] cors_configuration: The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] credentials_arn: Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        :param pulumi.Input[str] description: The description of the API.
        :param pulumi.Input[str] name: The name of the API.
        :param pulumi.Input[str] protocol_type: The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        :param pulumi.Input[str] route_key: Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] route_selection_expression: The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
               Defaults to `$request.method $request.path`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the API.
        :param pulumi.Input[str] target: Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
               For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
               The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        :param pulumi.Input[str] version: A version identifier for the API.
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

            __props__['api_key_selection_expression'] = api_key_selection_expression
            __props__['cors_configuration'] = cors_configuration
            __props__['credentials_arn'] = credentials_arn
            __props__['description'] = description
            __props__['name'] = name
            if protocol_type is None:
                raise TypeError("Missing required property 'protocol_type'")
            __props__['protocol_type'] = protocol_type
            __props__['route_key'] = route_key
            __props__['route_selection_expression'] = route_selection_expression
            __props__['tags'] = tags
            __props__['target'] = target
            __props__['version'] = version
            __props__['api_endpoint'] = None
            __props__['arn'] = None
            __props__['execution_arn'] = None
        super(Api, __self__).__init__(
            'aws:apigatewayv2/api:Api',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            api_endpoint: Optional[pulumi.Input[str]] = None,
            api_key_selection_expression: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cors_configuration: Optional[pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']]] = None,
            credentials_arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            execution_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol_type: Optional[pulumi.Input[str]] = None,
            route_key: Optional[pulumi.Input[str]] = None,
            route_selection_expression: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Api':
        """
        Get an existing Api resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_endpoint: The URI of the API, of the form `{api-id}.execute-api.{region}.amazonaws.com`.
        :param pulumi.Input[str] api_key_selection_expression: An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
               Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
               Applicable for WebSocket APIs.
        :param pulumi.Input[str] arn: The ARN of the API.
        :param pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']] cors_configuration: The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] credentials_arn: Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        :param pulumi.Input[str] description: The description of the API.
        :param pulumi.Input[str] execution_arn: The ARN prefix to be used in an `lambda.Permission`'s `source_arn` attribute
               or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
               See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        :param pulumi.Input[str] name: The name of the API.
        :param pulumi.Input[str] protocol_type: The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        :param pulumi.Input[str] route_key: Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] route_selection_expression: The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
               Defaults to `$request.method $request.path`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the API.
        :param pulumi.Input[str] target: Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
               For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
               The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        :param pulumi.Input[str] version: A version identifier for the API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_endpoint"] = api_endpoint
        __props__["api_key_selection_expression"] = api_key_selection_expression
        __props__["arn"] = arn
        __props__["cors_configuration"] = cors_configuration
        __props__["credentials_arn"] = credentials_arn
        __props__["description"] = description
        __props__["execution_arn"] = execution_arn
        __props__["name"] = name
        __props__["protocol_type"] = protocol_type
        __props__["route_key"] = route_key
        __props__["route_selection_expression"] = route_selection_expression
        __props__["tags"] = tags
        __props__["target"] = target
        __props__["version"] = version
        return Api(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiEndpoint")
    def api_endpoint(self) -> str:
        """
        The URI of the API, of the form `{api-id}.execute-api.{region}.amazonaws.com`.
        """
        ...

    @property
    @pulumi.getter(name="apiKeySelectionExpression")
    def api_key_selection_expression(self) -> Optional[str]:
        """
        An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
        Applicable for WebSocket APIs.
        """
        ...

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the API.
        """
        ...

    @property
    @pulumi.getter(name="corsConfiguration")
    def cors_configuration(self) -> Optional['outputs.ApiCorsConfiguration']:
        """
        The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        """
        ...

    @property
    @pulumi.getter(name="credentialsArn")
    def credentials_arn(self) -> Optional[str]:
        """
        Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        """
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the API.
        """
        ...

    @property
    @pulumi.getter(name="executionArn")
    def execution_arn(self) -> str:
        """
        The ARN prefix to be used in an `lambda.Permission`'s `source_arn` attribute
        or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the API.
        """
        ...

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> str:
        """
        The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        """
        ...

    @property
    @pulumi.getter(name="routeKey")
    def route_key(self) -> Optional[str]:
        """
        Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        """
        ...

    @property
    @pulumi.getter(name="routeSelectionExpression")
    def route_selection_expression(self) -> Optional[str]:
        """
        The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        Defaults to `$request.method $request.path`.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the API.
        """
        ...

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
        For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
        The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        """
        ...

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        A version identifier for the API.
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

