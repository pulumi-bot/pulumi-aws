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

__all__ = ['StageArgs', 'Stage']

@pulumi.input_type
class StageArgs:
    def __init__(__self__, *,
                 deployment: pulumi.Input[str],
                 rest_api: pulumi.Input[str],
                 stage_name: pulumi.Input[str],
                 access_log_settings: Optional[pulumi.Input['StageAccessLogSettingsArgs']] = None,
                 cache_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 cache_cluster_size: Optional[pulumi.Input[str]] = None,
                 client_certificate_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 documentation_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 xray_tracing_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Stage resource.
        :param pulumi.Input[str] deployment: The ID of the deployment that the stage points to
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] stage_name: The name of the stage
        :param pulumi.Input['StageAccessLogSettingsArgs'] access_log_settings: Enables access logs for the API stage. Detailed below.
        :param pulumi.Input[bool] cache_cluster_enabled: Specifies whether a cache cluster is enabled for the stage
        :param pulumi.Input[str] cache_cluster_size: The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage.
        :param pulumi.Input[str] description: The description of the stage
        :param pulumi.Input[str] documentation_version: The version of the associated API documentation
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map that defines the stage variables
        :param pulumi.Input[bool] xray_tracing_enabled: Whether active tracing with X-ray is enabled. Defaults to `false`.
        """
        pulumi.set(__self__, "deployment", deployment)
        pulumi.set(__self__, "rest_api", rest_api)
        pulumi.set(__self__, "stage_name", stage_name)
        if access_log_settings is not None:
            pulumi.set(__self__, "access_log_settings", access_log_settings)
        if cache_cluster_enabled is not None:
            pulumi.set(__self__, "cache_cluster_enabled", cache_cluster_enabled)
        if cache_cluster_size is not None:
            pulumi.set(__self__, "cache_cluster_size", cache_cluster_size)
        if client_certificate_id is not None:
            pulumi.set(__self__, "client_certificate_id", client_certificate_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if documentation_version is not None:
            pulumi.set(__self__, "documentation_version", documentation_version)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)
        if xray_tracing_enabled is not None:
            pulumi.set(__self__, "xray_tracing_enabled", xray_tracing_enabled)

    @property
    @pulumi.getter
    def deployment(self) -> pulumi.Input[str]:
        """
        The ID of the deployment that the stage points to
        """
        return pulumi.get(self, "deployment")

    @deployment.setter
    def deployment(self, value: pulumi.Input[str]):
        pulumi.set(self, "deployment", value)

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
    @pulumi.getter(name="stageName")
    def stage_name(self) -> pulumi.Input[str]:
        """
        The name of the stage
        """
        return pulumi.get(self, "stage_name")

    @stage_name.setter
    def stage_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stage_name", value)

    @property
    @pulumi.getter(name="accessLogSettings")
    def access_log_settings(self) -> Optional[pulumi.Input['StageAccessLogSettingsArgs']]:
        """
        Enables access logs for the API stage. Detailed below.
        """
        return pulumi.get(self, "access_log_settings")

    @access_log_settings.setter
    def access_log_settings(self, value: Optional[pulumi.Input['StageAccessLogSettingsArgs']]):
        pulumi.set(self, "access_log_settings", value)

    @property
    @pulumi.getter(name="cacheClusterEnabled")
    def cache_cluster_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether a cache cluster is enabled for the stage
        """
        return pulumi.get(self, "cache_cluster_enabled")

    @cache_cluster_enabled.setter
    def cache_cluster_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cache_cluster_enabled", value)

    @property
    @pulumi.getter(name="cacheClusterSize")
    def cache_cluster_size(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        """
        return pulumi.get(self, "cache_cluster_size")

    @cache_cluster_size.setter
    def cache_cluster_size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cache_cluster_size", value)

    @property
    @pulumi.getter(name="clientCertificateId")
    def client_certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of a client certificate for the stage.
        """
        return pulumi.get(self, "client_certificate_id")

    @client_certificate_id.setter
    def client_certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_certificate_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the stage
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="documentationVersion")
    def documentation_version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the associated API documentation
        """
        return pulumi.get(self, "documentation_version")

    @documentation_version.setter
    def documentation_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "documentation_version", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map that defines the stage variables
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "variables", value)

    @property
    @pulumi.getter(name="xrayTracingEnabled")
    def xray_tracing_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether active tracing with X-ray is enabled. Defaults to `false`.
        """
        return pulumi.get(self, "xray_tracing_enabled")

    @xray_tracing_enabled.setter
    def xray_tracing_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_tracing_enabled", value)


class Stage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_log_settings: Optional[pulumi.Input[pulumi.InputType['StageAccessLogSettingsArgs']]] = None,
                 cache_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 cache_cluster_size: Optional[pulumi.Input[str]] = None,
                 client_certificate_id: Optional[pulumi.Input[str]] = None,
                 deployment: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 documentation_version: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 xray_tracing_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        `aws_api_gateway_stage` can be imported using `REST-API-ID/STAGE-NAME`, e.g.

        ```sh
         $ pulumi import aws:apigateway/stage:Stage example 12345abcde/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StageAccessLogSettingsArgs']] access_log_settings: Enables access logs for the API stage. Detailed below.
        :param pulumi.Input[bool] cache_cluster_enabled: Specifies whether a cache cluster is enabled for the stage
        :param pulumi.Input[str] cache_cluster_size: The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage.
        :param pulumi.Input[str] deployment: The ID of the deployment that the stage points to
        :param pulumi.Input[str] description: The description of the stage
        :param pulumi.Input[str] documentation_version: The version of the associated API documentation
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] stage_name: The name of the stage
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map that defines the stage variables
        :param pulumi.Input[bool] xray_tracing_enabled: Whether active tracing with X-ray is enabled. Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        `aws_api_gateway_stage` can be imported using `REST-API-ID/STAGE-NAME`, e.g.

        ```sh
         $ pulumi import aws:apigateway/stage:Stage example 12345abcde/example
        ```

        :param str resource_name: The name of the resource.
        :param StageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_log_settings: Optional[pulumi.Input[pulumi.InputType['StageAccessLogSettingsArgs']]] = None,
                 cache_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 cache_cluster_size: Optional[pulumi.Input[str]] = None,
                 client_certificate_id: Optional[pulumi.Input[str]] = None,
                 deployment: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 documentation_version: Optional[pulumi.Input[str]] = None,
                 rest_api: Optional[pulumi.Input[str]] = None,
                 stage_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 xray_tracing_enabled: Optional[pulumi.Input[bool]] = None,
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

            __props__['access_log_settings'] = access_log_settings
            __props__['cache_cluster_enabled'] = cache_cluster_enabled
            __props__['cache_cluster_size'] = cache_cluster_size
            __props__['client_certificate_id'] = client_certificate_id
            if deployment is None and not opts.urn:
                raise TypeError("Missing required property 'deployment'")
            __props__['deployment'] = deployment
            __props__['description'] = description
            __props__['documentation_version'] = documentation_version
            if rest_api is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api'")
            __props__['rest_api'] = rest_api
            if stage_name is None and not opts.urn:
                raise TypeError("Missing required property 'stage_name'")
            __props__['stage_name'] = stage_name
            __props__['tags'] = tags
            __props__['variables'] = variables
            __props__['xray_tracing_enabled'] = xray_tracing_enabled
            __props__['arn'] = None
            __props__['execution_arn'] = None
            __props__['invoke_url'] = None
        super(Stage, __self__).__init__(
            'aws:apigateway/stage:Stage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_log_settings: Optional[pulumi.Input[pulumi.InputType['StageAccessLogSettingsArgs']]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cache_cluster_enabled: Optional[pulumi.Input[bool]] = None,
            cache_cluster_size: Optional[pulumi.Input[str]] = None,
            client_certificate_id: Optional[pulumi.Input[str]] = None,
            deployment: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            documentation_version: Optional[pulumi.Input[str]] = None,
            execution_arn: Optional[pulumi.Input[str]] = None,
            invoke_url: Optional[pulumi.Input[str]] = None,
            rest_api: Optional[pulumi.Input[str]] = None,
            stage_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            xray_tracing_enabled: Optional[pulumi.Input[bool]] = None) -> 'Stage':
        """
        Get an existing Stage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['StageAccessLogSettingsArgs']] access_log_settings: Enables access logs for the API stage. Detailed below.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[bool] cache_cluster_enabled: Specifies whether a cache cluster is enabled for the stage
        :param pulumi.Input[str] cache_cluster_size: The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage.
        :param pulumi.Input[str] deployment: The ID of the deployment that the stage points to
        :param pulumi.Input[str] description: The description of the stage
        :param pulumi.Input[str] documentation_version: The version of the associated API documentation
        :param pulumi.Input[str] execution_arn: The execution ARN to be used in `lambda_permission`'s `source_arn`
               when allowing API Gateway to invoke a Lambda function,
               e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        :param pulumi.Input[str] invoke_url: The URL to invoke the API pointing to the stage,
               e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] stage_name: The name of the stage
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map that defines the stage variables
        :param pulumi.Input[bool] xray_tracing_enabled: Whether active tracing with X-ray is enabled. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_log_settings"] = access_log_settings
        __props__["arn"] = arn
        __props__["cache_cluster_enabled"] = cache_cluster_enabled
        __props__["cache_cluster_size"] = cache_cluster_size
        __props__["client_certificate_id"] = client_certificate_id
        __props__["deployment"] = deployment
        __props__["description"] = description
        __props__["documentation_version"] = documentation_version
        __props__["execution_arn"] = execution_arn
        __props__["invoke_url"] = invoke_url
        __props__["rest_api"] = rest_api
        __props__["stage_name"] = stage_name
        __props__["tags"] = tags
        __props__["variables"] = variables
        __props__["xray_tracing_enabled"] = xray_tracing_enabled
        return Stage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLogSettings")
    def access_log_settings(self) -> pulumi.Output[Optional['outputs.StageAccessLogSettings']]:
        """
        Enables access logs for the API stage. Detailed below.
        """
        return pulumi.get(self, "access_log_settings")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cacheClusterEnabled")
    def cache_cluster_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether a cache cluster is enabled for the stage
        """
        return pulumi.get(self, "cache_cluster_enabled")

    @property
    @pulumi.getter(name="cacheClusterSize")
    def cache_cluster_size(self) -> pulumi.Output[Optional[str]]:
        """
        The size of the cache cluster for the stage, if enabled. Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        """
        return pulumi.get(self, "cache_cluster_size")

    @property
    @pulumi.getter(name="clientCertificateId")
    def client_certificate_id(self) -> pulumi.Output[Optional[str]]:
        """
        The identifier of a client certificate for the stage.
        """
        return pulumi.get(self, "client_certificate_id")

    @property
    @pulumi.getter
    def deployment(self) -> pulumi.Output[str]:
        """
        The ID of the deployment that the stage points to
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the stage
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="documentationVersion")
    def documentation_version(self) -> pulumi.Output[Optional[str]]:
        """
        The version of the associated API documentation
        """
        return pulumi.get(self, "documentation_version")

    @property
    @pulumi.getter(name="executionArn")
    def execution_arn(self) -> pulumi.Output[str]:
        """
        The execution ARN to be used in `lambda_permission`'s `source_arn`
        when allowing API Gateway to invoke a Lambda function,
        e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        """
        return pulumi.get(self, "execution_arn")

    @property
    @pulumi.getter(name="invokeUrl")
    def invoke_url(self) -> pulumi.Output[str]:
        """
        The URL to invoke the API pointing to the stage,
        e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        """
        return pulumi.get(self, "invoke_url")

    @property
    @pulumi.getter(name="restApi")
    def rest_api(self) -> pulumi.Output[str]:
        """
        The ID of the associated REST API
        """
        return pulumi.get(self, "rest_api")

    @property
    @pulumi.getter(name="stageName")
    def stage_name(self) -> pulumi.Output[str]:
        """
        The name of the stage
        """
        return pulumi.get(self, "stage_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map that defines the stage variables
        """
        return pulumi.get(self, "variables")

    @property
    @pulumi.getter(name="xrayTracingEnabled")
    def xray_tracing_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether active tracing with X-ray is enabled. Defaults to `false`.
        """
        return pulumi.get(self, "xray_tracing_enabled")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

