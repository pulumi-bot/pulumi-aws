# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ResourceShare']


class ResourceShare(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_external_principals: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Resource Access Manager (RAM) Resource Share. To associate principals with the share, see the `ram.PrincipalAssociation` resource. To associate resources with the share, see the `ram.ResourceAssociation` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ram.ResourceShare("example",
            allow_external_principals=True,
            tags={
                "Environment": "Production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_external_principals: Indicates whether principals outside your organization can be associated with a resource share.
        :param pulumi.Input[str] name: The name of the resource share.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource share.
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

            __props__['allow_external_principals'] = allow_external_principals
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(ResourceShare, __self__).__init__(
            'aws:ram/resourceShare:ResourceShare',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_external_principals: Optional[pulumi.Input[bool]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ResourceShare':
        """
        Get an existing ResourceShare resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_external_principals: Indicates whether principals outside your organization can be associated with a resource share.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the resource share.
        :param pulumi.Input[str] name: The name of the resource share.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource share.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_external_principals"] = allow_external_principals
        __props__["arn"] = arn
        __props__["name"] = name
        __props__["tags"] = tags
        return ResourceShare(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowExternalPrincipals")
    def allow_external_principals(self) -> Optional[bool]:
        """
        Indicates whether principals outside your organization can be associated with a resource share.
        """
        return pulumi.get(self, "allow_external_principals")

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the resource share.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource share.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource share.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

