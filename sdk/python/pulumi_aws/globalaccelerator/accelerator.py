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

__all__ = ['Accelerator']


class Accelerator(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[pulumi.InputType['AcceleratorAttributesArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_address_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Accelerator resource with the given unique name, props, and options.
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

            __props__['attributes'] = attributes
            __props__['enabled'] = enabled
            __props__['ip_address_type'] = ip_address_type
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['dns_name'] = None
            __props__['hosted_zone_id'] = None
            __props__['ip_sets'] = None
        super(Accelerator, __self__).__init__(
            'aws:globalaccelerator/accelerator:Accelerator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[pulumi.InputType['AcceleratorAttributesArgs']]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            hosted_zone_id: Optional[pulumi.Input[str]] = None,
            ip_address_type: Optional[pulumi.Input[str]] = None,
            ip_sets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AcceleratorIpSetArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Accelerator':
        """
        Get an existing Accelerator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["dns_name"] = dns_name
        __props__["enabled"] = enabled
        __props__["hosted_zone_id"] = hosted_zone_id
        __props__["ip_address_type"] = ip_address_type
        __props__["ip_sets"] = ip_sets
        __props__["name"] = name
        __props__["tags"] = tags
        return Accelerator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional['outputs.AcceleratorAttributes']]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter(name="ipAddressType")
    def ip_address_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ip_address_type")

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> pulumi.Output[List['outputs.AcceleratorIpSet']]:
        return pulumi.get(self, "ip_sets")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

