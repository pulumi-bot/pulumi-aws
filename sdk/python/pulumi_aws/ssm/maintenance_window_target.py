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

__all__ = ['MaintenanceWindowTarget']


class MaintenanceWindowTarget(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a MaintenanceWindowTarget resource with the given unique name, props, and options.
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

            __props__['description'] = description
            __props__['name'] = name
            __props__['owner_information'] = owner_information
            if resource_type is None:
                raise TypeError("Missing required property 'resource_type'")
            __props__['resource_type'] = resource_type
            if targets is None:
                raise TypeError("Missing required property 'targets'")
            __props__['targets'] = targets
            if window_id is None:
                raise TypeError("Missing required property 'window_id'")
            __props__['window_id'] = window_id
        super(MaintenanceWindowTarget, __self__).__init__(
            'aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_information: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetArgs']]]]] = None,
            window_id: Optional[pulumi.Input[str]] = None) -> 'MaintenanceWindowTarget':
        """
        Get an existing MaintenanceWindowTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["owner_information"] = owner_information
        __props__["resource_type"] = resource_type
        __props__["targets"] = targets
        __props__["window_id"] = window_id
        return MaintenanceWindowTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "owner_information")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[List['outputs.MaintenanceWindowTargetTarget']]:
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "window_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

