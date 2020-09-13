# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Route']


class Route(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 api_key_required: Optional[pulumi.Input[bool]] = None,
                 authorization_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorization_type: Optional[pulumi.Input[str]] = None,
                 authorizer_id: Optional[pulumi.Input[str]] = None,
                 model_selection_expression: Optional[pulumi.Input[str]] = None,
                 operation_name: Optional[pulumi.Input[str]] = None,
                 request_models: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 route_key: Optional[pulumi.Input[str]] = None,
                 route_response_selection_expression: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Amazon API Gateway Version 2 route.
        More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Route("example",
            api_id=aws_apigatewayv2_api["example"]["id"],
            route_key="$default")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API identifier.
        :param pulumi.Input[bool] api_key_required: Boolean whether an API key is required for the route. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorization_scopes: The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
        :param pulumi.Input[str] authorization_type: The authorization type for the route.
               For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
               For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
               Defaults to `NONE`.
        :param pulumi.Input[str] authorizer_id: The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
        :param pulumi.Input[str] model_selection_expression: The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
        :param pulumi.Input[str] operation_name: The operation name for the route.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] request_models: The request models for the route.
        :param pulumi.Input[str] route_key: The route key for the route.
        :param pulumi.Input[str] route_response_selection_expression: The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
        :param pulumi.Input[str] target: The target for the route.
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

            if api_id is None:
                raise TypeError("Missing required property 'api_id'")
            __props__['api_id'] = api_id
            __props__['api_key_required'] = api_key_required
            __props__['authorization_scopes'] = authorization_scopes
            __props__['authorization_type'] = authorization_type
            __props__['authorizer_id'] = authorizer_id
            __props__['model_selection_expression'] = model_selection_expression
            __props__['operation_name'] = operation_name
            __props__['request_models'] = request_models
            if route_key is None:
                raise TypeError("Missing required property 'route_key'")
            __props__['route_key'] = route_key
            __props__['route_response_selection_expression'] = route_response_selection_expression
            __props__['target'] = target
        super(Route, __self__).__init__(
            'aws:apigatewayv2/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            api_key_required: Optional[pulumi.Input[bool]] = None,
            authorization_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            authorization_type: Optional[pulumi.Input[str]] = None,
            authorizer_id: Optional[pulumi.Input[str]] = None,
            model_selection_expression: Optional[pulumi.Input[str]] = None,
            operation_name: Optional[pulumi.Input[str]] = None,
            request_models: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            route_key: Optional[pulumi.Input[str]] = None,
            route_response_selection_expression: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: The API identifier.
        :param pulumi.Input[bool] api_key_required: Boolean whether an API key is required for the route. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorization_scopes: The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
        :param pulumi.Input[str] authorization_type: The authorization type for the route.
               For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
               For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
               Defaults to `NONE`.
        :param pulumi.Input[str] authorizer_id: The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
        :param pulumi.Input[str] model_selection_expression: The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
        :param pulumi.Input[str] operation_name: The operation name for the route.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] request_models: The request models for the route.
        :param pulumi.Input[str] route_key: The route key for the route.
        :param pulumi.Input[str] route_response_selection_expression: The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
        :param pulumi.Input[str] target: The target for the route.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_id"] = api_id
        __props__["api_key_required"] = api_key_required
        __props__["authorization_scopes"] = authorization_scopes
        __props__["authorization_type"] = authorization_type
        __props__["authorizer_id"] = authorizer_id
        __props__["model_selection_expression"] = model_selection_expression
        __props__["operation_name"] = operation_name
        __props__["request_models"] = request_models
        __props__["route_key"] = route_key
        __props__["route_response_selection_expression"] = route_response_selection_expression
        __props__["target"] = target
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        The API identifier.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="apiKeyRequired")
    def api_key_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean whether an API key is required for the route. Defaults to `false`.
        """
        return pulumi.get(self, "api_key_required")

    @property
    @pulumi.getter(name="authorizationScopes")
    def authorization_scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
        """
        return pulumi.get(self, "authorization_scopes")

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Output[Optional[str]]:
        """
        The authorization type for the route.
        For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
        For HTTP APIs, valid values are `NONE` for open access, or `JWT` for using JSON Web Tokens.
        Defaults to `NONE`.
        """
        return pulumi.get(self, "authorization_type")

    @property
    @pulumi.getter(name="authorizerId")
    def authorizer_id(self) -> pulumi.Output[Optional[str]]:
        """
        The identifier of the `apigatewayv2.Authorizer` resource to be associated with this route, if the authorizationType is `CUSTOM`.
        """
        return pulumi.get(self, "authorizer_id")

    @property
    @pulumi.getter(name="modelSelectionExpression")
    def model_selection_expression(self) -> pulumi.Output[Optional[str]]:
        """
        The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route.
        """
        return pulumi.get(self, "model_selection_expression")

    @property
    @pulumi.getter(name="operationName")
    def operation_name(self) -> pulumi.Output[Optional[str]]:
        """
        The operation name for the route.
        """
        return pulumi.get(self, "operation_name")

    @property
    @pulumi.getter(name="requestModels")
    def request_models(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The request models for the route.
        """
        return pulumi.get(self, "request_models")

    @property
    @pulumi.getter(name="routeKey")
    def route_key(self) -> pulumi.Output[str]:
        """
        The route key for the route.
        """
        return pulumi.get(self, "route_key")

    @property
    @pulumi.getter(name="routeResponseSelectionExpression")
    def route_response_selection_expression(self) -> pulumi.Output[Optional[str]]:
        """
        The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route.
        """
        return pulumi.get(self, "route_response_selection_expression")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[Optional[str]]:
        """
        The target for the route.
        """
        return pulumi.get(self, "target")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

