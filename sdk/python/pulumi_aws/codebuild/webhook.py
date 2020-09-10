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
                 branch_filter: Optional[pulumi.Input[str]] = None,
                 filter_groups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebhookFilterGroupArgs']]]]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
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

            __props__['branch_filter'] = branch_filter
            __props__['filter_groups'] = filter_groups
            if project_name is None:
                raise TypeError("Missing required property 'project_name'")
            __props__['project_name'] = project_name
            __props__['payload_url'] = None
            __props__['secret'] = None
            __props__['url'] = None
        super(Webhook, __self__).__init__(
            'aws:codebuild/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch_filter: Optional[pulumi.Input[str]] = None,
            filter_groups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebhookFilterGroupArgs']]]]] = None,
            payload_url: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
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

        __props__["branch_filter"] = branch_filter
        __props__["filter_groups"] = filter_groups
        __props__["payload_url"] = payload_url
        __props__["project_name"] = project_name
        __props__["secret"] = secret
        __props__["url"] = url
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="branchFilter")
    def branch_filter(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "branch_filter")

    @property
    @pulumi.getter(name="filterGroups")
    def filter_groups(self) -> pulumi.Output[Optional[List['outputs.WebhookFilterGroup']]]:
        return pulumi.get(self, "filter_groups")

    @property
    @pulumi.getter(name="payloadUrl")
    def payload_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "payload_url")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

