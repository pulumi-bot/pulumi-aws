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

__all__ = ['ParameterGroupArgs', 'ParameterGroup']

@pulumi.input_type
class ParameterGroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]] = None):
        """
        The set of arguments for constructing a ParameterGroup resource.
        :param pulumi.Input[str] description: A description of the parameter group.
        :param pulumi.Input[str] name: The name of the parameter group.
        :param pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]] parameters: The parameters of the parameter group.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the parameter group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]]:
        """
        The parameters of the parameter group.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParameterArgs']]]]):
        pulumi.set(self, "parameters", value)


class ParameterGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DAX Parameter Group resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.dax.ParameterGroup("example", parameters=[
            aws.dax.ParameterGroupParameterArgs(
                name="query-ttl-millis",
                value="100000",
            ),
            aws.dax.ParameterGroupParameterArgs(
                name="record-ttl-millis",
                value="100000",
            ),
        ])
        ```

        ## Import

        DAX Parameter Group can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:dax/parameterGroup:ParameterGroup example my_dax_pg
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the parameter group.
        :param pulumi.Input[str] name: The name of the parameter group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]] parameters: The parameters of the parameter group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ParameterGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DAX Parameter Group resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.dax.ParameterGroup("example", parameters=[
            aws.dax.ParameterGroupParameterArgs(
                name="query-ttl-millis",
                value="100000",
            ),
            aws.dax.ParameterGroupParameterArgs(
                name="record-ttl-millis",
                value="100000",
            ),
        ])
        ```

        ## Import

        DAX Parameter Group can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:dax/parameterGroup:ParameterGroup example my_dax_pg
        ```

        :param str resource_name: The name of the resource.
        :param ParameterGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ParameterGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]]] = None,
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
            __props__['parameters'] = parameters
        super(ParameterGroup, __self__).__init__(
            'aws:dax/parameterGroup:ParameterGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]]] = None) -> 'ParameterGroup':
        """
        Get an existing ParameterGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the parameter group.
        :param pulumi.Input[str] name: The name of the parameter group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParameterArgs']]]] parameters: The parameters of the parameter group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["parameters"] = parameters
        return ParameterGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the parameter group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the parameter group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.ParameterGroupParameter']]:
        """
        The parameters of the parameter group.
        """
        return pulumi.get(self, "parameters")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

