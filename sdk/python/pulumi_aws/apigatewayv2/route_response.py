# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['RouteResponse']


class RouteResponse(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 model_selection_expression: Optional[pulumi.Input[str]] = None,
                 response_models: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 route_response_key: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a RouteResponse resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
            __props__['model_selection_expression'] = model_selection_expression
            __props__['response_models'] = response_models
            if route_id is None:
                raise TypeError("Missing required property 'route_id'")
            __props__['route_id'] = route_id
            if route_response_key is None:
                raise TypeError("Missing required property 'route_response_key'")
            __props__['route_response_key'] = route_response_key
        super(RouteResponse, __self__).__init__(
            'aws:apigatewayv2/routeResponse:RouteResponse',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            model_selection_expression: Optional[pulumi.Input[str]] = None,
            response_models: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            route_id: Optional[pulumi.Input[str]] = None,
            route_response_key: Optional[pulumi.Input[str]] = None) -> 'RouteResponse':
        """
        Get an existing RouteResponse resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_id"] = api_id
        __props__["model_selection_expression"] = model_selection_expression
        __props__["response_models"] = response_models
        __props__["route_id"] = route_id
        __props__["route_response_key"] = route_response_key
        return RouteResponse(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="modelSelectionExpression")
    def model_selection_expression(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "model_selection_expression")

    @property
    @pulumi.getter(name="responseModels")
    def response_models(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "response_models")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "route_id")

    @property
    @pulumi.getter(name="routeResponseKey")
    def route_response_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "route_response_key")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

