# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Stage(pulumi.CustomResource):
    access_log_settings: pulumi.Output[dict]
    """
    Enables access logs for the API stage. Detailed below.
    """
    cache_cluster_enabled: pulumi.Output[bool]
    """
    Specifies whether a cache cluster is enabled for the stage
    """
    cache_cluster_size: pulumi.Output[str]
    """
    The size of the cache cluster for the stage, if enabled.
    Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
    """
    client_certificate_id: pulumi.Output[str]
    """
    The identifier of a client certificate for the stage.
    """
    deployment: pulumi.Output[str]
    """
    The ID of the deployment that the stage points to
    """
    description: pulumi.Output[str]
    """
    The description of the stage
    """
    documentation_version: pulumi.Output[str]
    """
    The version of the associated API documentation
    """
    execution_arn: pulumi.Output[str]
    """
    The execution ARN to be used in [`lambda_permission`](https://www.terraform.io/docs/providers/aws/r/lambda_permission.html)'s `source_arn`
    when allowing API Gateway to invoke a Lambda function,
    e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
    """
    invoke_url: pulumi.Output[str]
    """
    The URL to invoke the API pointing to the stage,
    e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
    """
    rest_api: pulumi.Output[str]
    """
    The ID of the associated REST API
    """
    stage_name: pulumi.Output[str]
    """
    The name of the stage
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    variables: pulumi.Output[dict]
    """
    A map that defines the stage variables
    """
    xray_tracing_enabled: pulumi.Output[bool]
    """
    Whether active tracing with X-ray is enabled. Defaults to `false`.
    """
    def __init__(__self__, resource_name, opts=None, access_log_settings=None, cache_cluster_enabled=None, cache_cluster_size=None, client_certificate_id=None, deployment=None, description=None, documentation_version=None, rest_api=None, stage_name=None, tags=None, variables=None, xray_tracing_enabled=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Stage.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_log_settings: Enables access logs for the API stage. Detailed below.
        :param pulumi.Input[bool] cache_cluster_enabled: Specifies whether a cache cluster is enabled for the stage
        :param pulumi.Input[str] cache_cluster_size: The size of the cache cluster for the stage, if enabled.
               Allowed values include `0.5`, `1.6`, `6.1`, `13.5`, `28.4`, `58.2`, `118` and `237`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage.
        :param pulumi.Input[str] deployment: The ID of the deployment that the stage points to
        :param pulumi.Input[str] description: The description of the stage
        :param pulumi.Input[str] documentation_version: The version of the associated API documentation
        :param pulumi.Input[str] rest_api: The ID of the associated REST API
        :param pulumi.Input[str] stage_name: The name of the stage
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] variables: A map that defines the stage variables
        :param pulumi.Input[bool] xray_tracing_enabled: Whether active tracing with X-ray is enabled. Defaults to `false`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['access_log_settings'] = access_log_settings

        __props__['cache_cluster_enabled'] = cache_cluster_enabled

        __props__['cache_cluster_size'] = cache_cluster_size

        __props__['client_certificate_id'] = client_certificate_id

        if deployment is None:
            raise TypeError("Missing required property 'deployment'")
        __props__['deployment'] = deployment

        __props__['description'] = description

        __props__['documentation_version'] = documentation_version

        if rest_api is None:
            raise TypeError("Missing required property 'rest_api'")
        __props__['rest_api'] = rest_api

        if stage_name is None:
            raise TypeError("Missing required property 'stage_name'")
        __props__['stage_name'] = stage_name

        __props__['tags'] = tags

        __props__['variables'] = variables

        __props__['xray_tracing_enabled'] = xray_tracing_enabled

        __props__['execution_arn'] = None
        __props__['invoke_url'] = None

        super(Stage, __self__).__init__(
            'aws:apigateway/stage:Stage',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

