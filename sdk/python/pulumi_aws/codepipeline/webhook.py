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

__all__ = ['Webhook']


class Webhook(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[str]] = None,
                 authentication_configuration: Optional[pulumi.Input[pulumi.InputType['WebhookAuthenticationConfigurationArgs']]] = None,
                 filters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebhookFilterArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_action: Optional[pulumi.Input[str]] = None,
                 target_pipeline: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Webhook resource with the given unique name, props, and options.
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

            if authentication is None:
                raise TypeError("Missing required property 'authentication'")
            __props__['authentication'] = authentication
            __props__['authentication_configuration'] = authentication_configuration
            if filters is None:
                raise TypeError("Missing required property 'filters'")
            __props__['filters'] = filters
            __props__['name'] = name
            __props__['tags'] = tags
            if target_action is None:
                raise TypeError("Missing required property 'target_action'")
            __props__['target_action'] = target_action
            if target_pipeline is None:
                raise TypeError("Missing required property 'target_pipeline'")
            __props__['target_pipeline'] = target_pipeline
            __props__['url'] = None
        super(Webhook, __self__).__init__(
            'aws:codepipeline/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication: Optional[pulumi.Input[str]] = None,
            authentication_configuration: Optional[pulumi.Input[pulumi.InputType['WebhookAuthenticationConfigurationArgs']]] = None,
            filters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebhookFilterArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_action: Optional[pulumi.Input[str]] = None,
            target_pipeline: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authentication"] = authentication
        __props__["authentication_configuration"] = authentication_configuration
        __props__["filters"] = filters
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["target_action"] = target_action
        __props__["target_pipeline"] = target_pipeline
        __props__["url"] = url
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output[str]:
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="authenticationConfiguration")
    def authentication_configuration(self) -> pulumi.Output[Optional['outputs.WebhookAuthenticationConfiguration']]:
        return pulumi.get(self, "authentication_configuration")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[List['outputs.WebhookFilter']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetAction")
    def target_action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_action")

    @property
    @pulumi.getter(name="targetPipeline")
    def target_pipeline(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_pipeline")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

