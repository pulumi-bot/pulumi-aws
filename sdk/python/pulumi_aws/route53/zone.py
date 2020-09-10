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

__all__ = ['Zone']


class Zone(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 delegation_set_id: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpcs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Zone resource with the given unique name, props, and options.
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

            if comment is None:
                comment = 'Managed by Pulumi'
            __props__['comment'] = comment
            __props__['delegation_set_id'] = delegation_set_id
            __props__['force_destroy'] = force_destroy
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['vpcs'] = vpcs
            __props__['name_servers'] = None
            __props__['zone_id'] = None
        super(Zone, __self__).__init__(
            'aws:route53/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            delegation_set_id: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_servers: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpcs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ZoneVpcArgs']]]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["comment"] = comment
        __props__["delegation_set_id"] = delegation_set_id
        __props__["force_destroy"] = force_destroy
        __props__["name"] = name
        __props__["name_servers"] = name_servers
        __props__["tags"] = tags
        __props__["vpcs"] = vpcs
        __props__["zone_id"] = zone_id
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="delegationSetId")
    def delegation_set_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "delegation_set_id")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vpcs(self) -> pulumi.Output[Optional[List['outputs.ZoneVpc']]]:
        return pulumi.get(self, "vpcs")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "zone_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

