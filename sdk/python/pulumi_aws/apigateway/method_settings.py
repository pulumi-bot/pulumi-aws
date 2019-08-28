# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class MethodSettings(pulumi.CustomResource):
    method_path: pulumi.Output[str]
    """
    Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
    """
    rest_api: pulumi.Output[str]
    """
    The ID of the REST API
    """
    settings: pulumi.Output[dict]
    """
    The settings block, see below.
    
      * `cacheDataEncrypted` (`bool`) - Specifies whether the cached responses are encrypted.
      * `cacheTtlInSeconds` (`float`) - Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
      * `cachingEnabled` (`bool`) - Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached. 
      * `dataTraceEnabled` (`bool`) - Specifies whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
      * `loggingLevel` (`str`) - Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are `OFF`, `ERROR`, and `INFO`.
      * `metricsEnabled` (`bool`) - Specifies whether Amazon CloudWatch metrics are enabled for this method.
      * `requireAuthorizationForCacheControl` (`bool`) - Specifies whether authorization is required for a cache invalidation request.
      * `throttlingBurstLimit` (`float`) - Specifies the throttling burst limit.
      * `throttlingRateLimit` (`float`) - Specifies the throttling rate limit.
      * `unauthorizedCacheControlHeaderStrategy` (`str`) - Specifies how to handle unauthorized requests for cache invalidation. The available values are `FAIL_WITH_403`, `SUCCEED_WITH_RESPONSE_HEADER`, `SUCCEED_WITHOUT_RESPONSE_HEADER`.
    """
    stage_name: pulumi.Output[str]
    """
    The name of the stage
    """
    def __init__(__self__, resource_name, opts=None, method_path=None, rest_api=None, settings=None, stage_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Method Settings, e.g. logging or monitoring.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] method_path: Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
        :param pulumi.Input[str] rest_api: The ID of the REST API
        :param pulumi.Input[dict] settings: The settings block, see below.
        :param pulumi.Input[str] stage_name: The name of the stage
        
        The **settings** object supports the following:
        
          * `cacheDataEncrypted` (`pulumi.Input[bool]`) - Specifies whether the cached responses are encrypted.
          * `cacheTtlInSeconds` (`pulumi.Input[float]`) - Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
          * `cachingEnabled` (`pulumi.Input[bool]`) - Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached. 
          * `dataTraceEnabled` (`pulumi.Input[bool]`) - Specifies whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
          * `loggingLevel` (`pulumi.Input[str]`) - Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are `OFF`, `ERROR`, and `INFO`.
          * `metricsEnabled` (`pulumi.Input[bool]`) - Specifies whether Amazon CloudWatch metrics are enabled for this method.
          * `requireAuthorizationForCacheControl` (`pulumi.Input[bool]`) - Specifies whether authorization is required for a cache invalidation request.
          * `throttlingBurstLimit` (`pulumi.Input[float]`) - Specifies the throttling burst limit.
          * `throttlingRateLimit` (`pulumi.Input[float]`) - Specifies the throttling rate limit.
          * `unauthorizedCacheControlHeaderStrategy` (`pulumi.Input[str]`) - Specifies how to handle unauthorized requests for cache invalidation. The available values are `FAIL_WITH_403`, `SUCCEED_WITH_RESPONSE_HEADER`, `SUCCEED_WITHOUT_RESPONSE_HEADER`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_method_settings.html.markdown.
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

            if method_path is None:
                raise TypeError("Missing required property 'method_path'")
            __props__['method_path'] = method_path
            if rest_api is None:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            if settings is None:
                raise TypeError("Missing required property 'settings'")
            __props__['settings'] = settings
            if stage_name is None:
                raise TypeError("Missing required property 'stage_name'")
            __props__['stage_name'] = stage_name
        super(MethodSettings, __self__).__init__(
            'aws:apigateway/methodSettings:MethodSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, method_path=None, rest_api=None, settings=None, stage_name=None):
        """
        Get an existing MethodSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] method_path: Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
        :param pulumi.Input[str] rest_api: The ID of the REST API
        :param pulumi.Input[dict] settings: The settings block, see below.
        :param pulumi.Input[str] stage_name: The name of the stage
        
        The **settings** object supports the following:
        
          * `cacheDataEncrypted` (`pulumi.Input[bool]`) - Specifies whether the cached responses are encrypted.
          * `cacheTtlInSeconds` (`pulumi.Input[float]`) - Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
          * `cachingEnabled` (`pulumi.Input[bool]`) - Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached. 
          * `dataTraceEnabled` (`pulumi.Input[bool]`) - Specifies whether data trace logging is enabled for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
          * `loggingLevel` (`pulumi.Input[str]`) - Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs. The available levels are `OFF`, `ERROR`, and `INFO`.
          * `metricsEnabled` (`pulumi.Input[bool]`) - Specifies whether Amazon CloudWatch metrics are enabled for this method.
          * `requireAuthorizationForCacheControl` (`pulumi.Input[bool]`) - Specifies whether authorization is required for a cache invalidation request.
          * `throttlingBurstLimit` (`pulumi.Input[float]`) - Specifies the throttling burst limit.
          * `throttlingRateLimit` (`pulumi.Input[float]`) - Specifies the throttling rate limit.
          * `unauthorizedCacheControlHeaderStrategy` (`pulumi.Input[str]`) - Specifies how to handle unauthorized requests for cache invalidation. The available values are `FAIL_WITH_403`, `SUCCEED_WITH_RESPONSE_HEADER`, `SUCCEED_WITHOUT_RESPONSE_HEADER`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_method_settings.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["method_path"] = method_path
        __props__["rest_api"] = rest_api
        __props__["settings"] = settings
        __props__["stage_name"] = stage_name
        return MethodSettings(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

