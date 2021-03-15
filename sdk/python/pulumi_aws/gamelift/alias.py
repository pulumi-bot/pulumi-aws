# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['AliasArgs', 'Alias']

@pulumi.input_type
class AliasArgs:
    def __init__(__self__, *,
                 routing_strategy: pulumi.Input['AliasRoutingStrategyArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Alias resource.
        :param pulumi.Input['AliasRoutingStrategyArgs'] routing_strategy: Specifies the fleet and/or routing type to use for the alias.
        :param pulumi.Input[str] description: Description of the alias.
        :param pulumi.Input[str] name: Name of the alias.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        pulumi.set(__self__, "routing_strategy", routing_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="routingStrategy")
    def routing_strategy(self) -> pulumi.Input['AliasRoutingStrategyArgs']:
        """
        Specifies the fleet and/or routing type to use for the alias.
        """
        return pulumi.get(self, "routing_strategy")

    @routing_strategy.setter
    def routing_strategy(self, value: pulumi.Input['AliasRoutingStrategyArgs']):
        pulumi.set(self, "routing_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the alias.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the alias.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Alias(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AliasArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Gamelift Alias resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.gamelift.Alias("example",
            description="Example Description",
            routing_strategy=aws.gamelift.AliasRoutingStrategyArgs(
                message="Example Message",
                type="TERMINAL",
            ))
        ```

        ## Import

        Gamelift Aliases can be imported using the ID, e.g.

        ```sh
         $ pulumi import aws:gamelift/alias:Alias example <alias-id>
        ```

        :param str resource_name: The name of the resource.
        :param AliasArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 routing_strategy: Optional[pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Gamelift Alias resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.gamelift.Alias("example",
            description="Example Description",
            routing_strategy=aws.gamelift.AliasRoutingStrategyArgs(
                message="Example Message",
                type="TERMINAL",
            ))
        ```

        ## Import

        Gamelift Aliases can be imported using the ID, e.g.

        ```sh
         $ pulumi import aws:gamelift/alias:Alias example <alias-id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the alias.
        :param pulumi.Input[str] name: Name of the alias.
        :param pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']] routing_strategy: Specifies the fleet and/or routing type to use for the alias.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AliasArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 routing_strategy: Optional[pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if routing_strategy is None and not opts.urn:
                raise TypeError("Missing required property 'routing_strategy'")
            __props__['routing_strategy'] = routing_strategy
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Alias, __self__).__init__(
            'aws:gamelift/alias:Alias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            routing_strategy: Optional[pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Alias':
        """
        Get an existing Alias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Alias ARN.
        :param pulumi.Input[str] description: Description of the alias.
        :param pulumi.Input[str] name: Name of the alias.
        :param pulumi.Input[pulumi.InputType['AliasRoutingStrategyArgs']] routing_strategy: Specifies the fleet and/or routing type to use for the alias.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["name"] = name
        __props__["routing_strategy"] = routing_strategy
        __props__["tags"] = tags
        return Alias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Alias ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the alias.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the alias.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="routingStrategy")
    def routing_strategy(self) -> pulumi.Output['outputs.AliasRoutingStrategy']:
        """
        Specifies the fleet and/or routing type to use for the alias.
        """
        return pulumi.get(self, "routing_strategy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

