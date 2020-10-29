# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ReportGroup']


class ReportGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 export_config: Optional[pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CodeBuild Report Groups Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']] export_config: Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        :param pulumi.Input[str] name: The name of a Report Group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] type: The export configuration type. Valid values are `S3` and `NO_EXPORT`.
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

            if export_config is None:
                raise TypeError("Missing required property 'export_config'")
            __props__['export_config'] = export_config
            __props__['name'] = name
            __props__['tags'] = tags
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['arn'] = None
            __props__['created'] = None
        super(ReportGroup, __self__).__init__(
            'aws:codebuild/reportGroup:ReportGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            created: Optional[pulumi.Input[str]] = None,
            export_config: Optional[pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ReportGroup':
        """
        Get an existing ReportGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of Report Group.
        :param pulumi.Input[str] created: The date and time this Report Group was created.
        :param pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']] export_config: Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        :param pulumi.Input[str] name: The name of a Report Group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] type: The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["created"] = created
        __props__["export_config"] = export_config
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["type"] = type
        return ReportGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of Report Group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        The date and time this Report Group was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="exportConfig")
    def export_config(self) -> pulumi.Output['outputs.ReportGroupExportConfig']:
        """
        Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        """
        return pulumi.get(self, "export_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of a Report Group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

