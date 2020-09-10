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

__all__ = ['Selection']


class Selection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 selection_tags: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SelectionSelectionTagArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Selection resource with the given unique name, props, and options.
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

            if iam_role_arn is None:
                raise TypeError("Missing required property 'iam_role_arn'")
            __props__['iam_role_arn'] = iam_role_arn
            __props__['name'] = name
            if plan_id is None:
                raise TypeError("Missing required property 'plan_id'")
            __props__['plan_id'] = plan_id
            __props__['resources'] = resources
            __props__['selection_tags'] = selection_tags
        super(Selection, __self__).__init__(
            'aws:backup/selection:Selection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            iam_role_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            plan_id: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            selection_tags: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SelectionSelectionTagArgs']]]]] = None) -> 'Selection':
        """
        Get an existing Selection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["iam_role_arn"] = iam_role_arn
        __props__["name"] = name
        __props__["plan_id"] = plan_id
        __props__["resources"] = resources
        __props__["selection_tags"] = selection_tags
        return Selection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="selectionTags")
    def selection_tags(self) -> pulumi.Output[Optional[List['outputs.SelectionSelectionTag']]]:
        return pulumi.get(self, "selection_tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

