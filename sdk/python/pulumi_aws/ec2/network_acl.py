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

__all__ = ['NetworkAcl']


class NetworkAcl(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 egress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclEgressArgs']]]]] = None,
                 ingress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclIngressArgs']]]]] = None,
                 subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a NetworkAcl resource with the given unique name, props, and options.
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

            __props__['egress'] = egress
            __props__['ingress'] = ingress
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
            __props__['owner_id'] = None
        super(NetworkAcl, __self__).__init__(
            'aws:ec2/networkAcl:NetworkAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            egress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclEgressArgs']]]]] = None,
            ingress: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NetworkAclIngressArgs']]]]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'NetworkAcl':
        """
        Get an existing NetworkAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["egress"] = egress
        __props__["ingress"] = ingress
        __props__["owner_id"] = owner_id
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return NetworkAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def egress(self) -> pulumi.Output[List['outputs.NetworkAclEgress']]:
        return pulumi.get(self, "egress")

    @property
    @pulumi.getter
    def ingress(self) -> pulumi.Output[List['outputs.NetworkAclIngress']]:
        return pulumi.get(self, "ingress")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

