# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ApiArgs', 'Api']

@pulumi.input_type
class ApiArgs:
    def __init__(__self__, *,
                 protocol_type: pulumi.Input[str],
                 api_key_selection_expression: Optional[pulumi.Input[str]] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 cors_configuration: Optional[pulumi.Input['ApiCorsConfigurationArgs']] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_execute_api_endpoint: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 route_key: Optional[pulumi.Input[str]] = None,
                 route_selection_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Api resource.
        :param pulumi.Input[str] protocol_type: The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        :param pulumi.Input[str] api_key_selection_expression: An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
               Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
               Applicable for WebSocket APIs.
        :param pulumi.Input[str] body: An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        :param pulumi.Input['ApiCorsConfigurationArgs'] cors_configuration: The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] credentials_arn: Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        :param pulumi.Input[str] description: The description of the API. Must be less than or equal to 1024 characters in length.
        :param pulumi.Input[bool] disable_execute_api_endpoint: Whether clients can invoke the API by using the default `execute-api` endpoint.
               By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
               To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        :param pulumi.Input[str] name: The name of the API. Must be less than or equal to 128 characters in length.
        :param pulumi.Input[str] route_key: Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] route_selection_expression: The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
               Defaults to `$request.method $request.path`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the API.
        :param pulumi.Input[str] target: Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
               For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
               The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        :param pulumi.Input[str] version: A version identifier for the API. Must be between 1 and 64 characters in length.
        """
        pulumi.set(__self__, "protocol_type", protocol_type)
        if api_key_selection_expression is not None:
            pulumi.set(__self__, "api_key_selection_expression", api_key_selection_expression)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if cors_configuration is not None:
            pulumi.set(__self__, "cors_configuration", cors_configuration)
        if credentials_arn is not None:
            pulumi.set(__self__, "credentials_arn", credentials_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_execute_api_endpoint is not None:
            pulumi.set(__self__, "disable_execute_api_endpoint", disable_execute_api_endpoint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if route_key is not None:
            pulumi.set(__self__, "route_key", route_key)
        if route_selection_expression is not None:
            pulumi.set(__self__, "route_selection_expression", route_selection_expression)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Input[str]:
        """
        The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        """
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol_type", value)

    @property
    @pulumi.getter(name="apiKeySelectionExpression")
    def api_key_selection_expression(self) -> Optional[pulumi.Input[str]]:
        """
        An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
        Applicable for WebSocket APIs.
        """
        return pulumi.get(self, "api_key_selection_expression")

    @api_key_selection_expression.setter
    def api_key_selection_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key_selection_expression", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter(name="corsConfiguration")
    def cors_configuration(self) -> Optional[pulumi.Input['ApiCorsConfigurationArgs']]:
        """
        The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        """
        return pulumi.get(self, "cors_configuration")

    @cors_configuration.setter
    def cors_configuration(self, value: Optional[pulumi.Input['ApiCorsConfigurationArgs']]):
        pulumi.set(self, "cors_configuration", value)

    @property
    @pulumi.getter(name="credentialsArn")
    def credentials_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        """
        return pulumi.get(self, "credentials_arn")

    @credentials_arn.setter
    def credentials_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credentials_arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the API. Must be less than or equal to 1024 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableExecuteApiEndpoint")
    def disable_execute_api_endpoint(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether clients can invoke the API by using the default `execute-api` endpoint.
        By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
        To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        """
        return pulumi.get(self, "disable_execute_api_endpoint")

    @disable_execute_api_endpoint.setter
    def disable_execute_api_endpoint(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_execute_api_endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the API. Must be less than or equal to 128 characters in length.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="routeKey")
    def route_key(self) -> Optional[pulumi.Input[str]]:
        """
        Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        """
        return pulumi.get(self, "route_key")

    @route_key.setter
    def route_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_key", value)

    @property
    @pulumi.getter(name="routeSelectionExpression")
    def route_selection_expression(self) -> Optional[pulumi.Input[str]]:
        """
        The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        Defaults to `$request.method $request.path`.
        """
        return pulumi.get(self, "route_selection_expression")

    @route_selection_expression.setter
    def route_selection_expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_selection_expression", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the API.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
        For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
        The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        A version identifier for the API. Must be between 1 and 64 characters in length.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Api(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        `aws_apigatewayv2_api` can be imported by using the API identifier, e.g.

        ```sh
         $ pulumi import aws:apigatewayv2/api:Api example aabbccddee
        ```

        :param str resource_name: The name of the resource.
        :param ApiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key_selection_expression: Optional[pulumi.Input[str]] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 cors_configuration: Optional[pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']]] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_execute_api_endpoint: Optional[pulumi.Input[bool]] = None,
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

        ## Import

        `aws_apigatewayv2_api` can be imported by using the API identifier, e.g.

        ```sh
         $ pulumi import aws:apigatewayv2/api:Api example aabbccddee
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key_selection_expression: An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
               Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
               Applicable for WebSocket APIs.
        :param pulumi.Input[str] body: An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        :param pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']] cors_configuration: The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] credentials_arn: Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        :param pulumi.Input[str] description: The description of the API. Must be less than or equal to 1024 characters in length.
        :param pulumi.Input[bool] disable_execute_api_endpoint: Whether clients can invoke the API by using the default `execute-api` endpoint.
               By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
               To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        :param pulumi.Input[str] name: The name of the API. Must be less than or equal to 128 characters in length.
        :param pulumi.Input[str] protocol_type: The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        :param pulumi.Input[str] route_key: Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] route_selection_expression: The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
               Defaults to `$request.method $request.path`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the API.
        :param pulumi.Input[str] target: Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
               For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
               The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        :param pulumi.Input[str] version: A version identifier for the API. Must be between 1 and 64 characters in length.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key_selection_expression: Optional[pulumi.Input[str]] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 cors_configuration: Optional[pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']]] = None,
                 credentials_arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_execute_api_endpoint: Optional[pulumi.Input[bool]] = None,
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
            __props__['body'] = body
            __props__['cors_configuration'] = cors_configuration
            __props__['credentials_arn'] = credentials_arn
            __props__['description'] = description
            __props__['disable_execute_api_endpoint'] = disable_execute_api_endpoint
            __props__['name'] = name
            if protocol_type is None and not opts.urn:
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
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_endpoint: Optional[pulumi.Input[str]] = None,
            api_key_selection_expression: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            body: Optional[pulumi.Input[str]] = None,
            cors_configuration: Optional[pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']]] = None,
            credentials_arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_execute_api_endpoint: Optional[pulumi.Input[bool]] = None,
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
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_endpoint: The URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
        :param pulumi.Input[str] api_key_selection_expression: An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
               Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
               Applicable for WebSocket APIs.
        :param pulumi.Input[str] arn: The ARN of the API.
        :param pulumi.Input[str] body: An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        :param pulumi.Input[pulumi.InputType['ApiCorsConfigurationArgs']] cors_configuration: The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] credentials_arn: Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        :param pulumi.Input[str] description: The description of the API. Must be less than or equal to 1024 characters in length.
        :param pulumi.Input[bool] disable_execute_api_endpoint: Whether clients can invoke the API by using the default `execute-api` endpoint.
               By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
               To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        :param pulumi.Input[str] execution_arn: The ARN prefix to be used in an `lambda.Permission`'s `source_arn` attribute
               or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
               See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        :param pulumi.Input[str] name: The name of the API. Must be less than or equal to 128 characters in length.
        :param pulumi.Input[str] protocol_type: The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        :param pulumi.Input[str] route_key: Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        :param pulumi.Input[str] route_selection_expression: The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
               Defaults to `$request.method $request.path`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the API.
        :param pulumi.Input[str] target: Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
               For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
               The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        :param pulumi.Input[str] version: A version identifier for the API. Must be between 1 and 64 characters in length.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_endpoint"] = api_endpoint
        __props__["api_key_selection_expression"] = api_key_selection_expression
        __props__["arn"] = arn
        __props__["body"] = body
        __props__["cors_configuration"] = cors_configuration
        __props__["credentials_arn"] = credentials_arn
        __props__["description"] = description
        __props__["disable_execute_api_endpoint"] = disable_execute_api_endpoint
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
    def api_endpoint(self) -> pulumi.Output[str]:
        """
        The URI of the API, of the form `https://{api-id}.execute-api.{region}.amazonaws.com` for HTTP APIs and `wss://{api-id}.execute-api.{region}.amazonaws.com` for WebSocket APIs.
        """
        return pulumi.get(self, "api_endpoint")

    @property
    @pulumi.getter(name="apiKeySelectionExpression")
    def api_key_selection_expression(self) -> pulumi.Output[Optional[str]]:
        """
        An [API key selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
        Valid values: `$context.authorizer.usageIdentifierKey`, `$request.header.x-api-key`. Defaults to `$request.header.x-api-key`.
        Applicable for WebSocket APIs.
        """
        return pulumi.get(self, "api_key_selection_expression")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the API.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def body(self) -> pulumi.Output[Optional[str]]:
        """
        An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="corsConfiguration")
    def cors_configuration(self) -> pulumi.Output[Optional['outputs.ApiCorsConfiguration']]:
        """
        The cross-origin resource sharing (CORS) [configuration](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html). Applicable for HTTP APIs.
        """
        return pulumi.get(self, "cors_configuration")

    @property
    @pulumi.getter(name="credentialsArn")
    def credentials_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Part of _quick create_. Specifies any credentials required for the integration. Applicable for HTTP APIs.
        """
        return pulumi.get(self, "credentials_arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the API. Must be less than or equal to 1024 characters in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableExecuteApiEndpoint")
    def disable_execute_api_endpoint(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether clients can invoke the API by using the default `execute-api` endpoint.
        By default, clients can invoke the API with the default `{api_id}.execute-api.{region}.amazonaws.com endpoint`.
        To require that clients use a custom domain name to invoke the API, disable the default endpoint.
        """
        return pulumi.get(self, "disable_execute_api_endpoint")

    @property
    @pulumi.getter(name="executionArn")
    def execution_arn(self) -> pulumi.Output[str]:
        """
        The ARN prefix to be used in an `lambda.Permission`'s `source_arn` attribute
        or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        """
        return pulumi.get(self, "execution_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the API. Must be less than or equal to 128 characters in length.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Output[str]:
        """
        The API protocol. Valid values: `HTTP`, `WEBSOCKET`.
        """
        return pulumi.get(self, "protocol_type")

    @property
    @pulumi.getter(name="routeKey")
    def route_key(self) -> pulumi.Output[Optional[str]]:
        """
        Part of _quick create_. Specifies any [route key](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html). Applicable for HTTP APIs.
        """
        return pulumi.get(self, "route_key")

    @property
    @pulumi.getter(name="routeSelectionExpression")
    def route_selection_expression(self) -> pulumi.Output[Optional[str]]:
        """
        The [route selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-selection-expressions) for the API.
        Defaults to `$request.method $request.path`.
        """
        return pulumi.get(self, "route_selection_expression")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the API.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[Optional[str]]:
        """
        Part of _quick create_. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
        For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
        The type of the integration will be `HTTP_PROXY` or `AWS_PROXY`, respectively. Applicable for HTTP APIs.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        A version identifier for the API. Must be between 1 and 64 characters in length.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

