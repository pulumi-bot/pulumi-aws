# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ThingArgs', 'Thing']

@pulumi.input_type
class ThingArgs:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 thing_type_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Thing resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attributes: Map of attributes of the thing.
        :param pulumi.Input[str] name: The name of the thing.
        :param pulumi.Input[str] thing_type_name: The thing type name.
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if thing_type_name is not None:
            pulumi.set(__self__, "thing_type_name", thing_type_name)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of attributes of the thing.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the thing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="thingTypeName")
    def thing_type_name(self) -> Optional[pulumi.Input[str]]:
        """
        The thing type name.
        """
        return pulumi.get(self, "thing_type_name")

    @thing_type_name.setter
    def thing_type_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thing_type_name", value)


class Thing(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 thing_type_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates and manages an AWS IoT Thing.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.iot.Thing("example", attributes={
            "First": "examplevalue",
        })
        ```

        ## Import

        IOT Things can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:iot/thing:Thing example example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attributes: Map of attributes of the thing.
        :param pulumi.Input[str] name: The name of the thing.
        :param pulumi.Input[str] thing_type_name: The thing type name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ThingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages an AWS IoT Thing.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.iot.Thing("example", attributes={
            "First": "examplevalue",
        })
        ```

        ## Import

        IOT Things can be imported using the name, e.g.

        ```sh
         $ pulumi import aws:iot/thing:Thing example example
        ```

        :param str resource_name: The name of the resource.
        :param ThingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ThingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 thing_type_name: Optional[pulumi.Input[str]] = None,
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

            __props__['attributes'] = attributes
            __props__['name'] = name
            __props__['thing_type_name'] = thing_type_name
            __props__['arn'] = None
            __props__['default_client_id'] = None
            __props__['version'] = None
        super(Thing, __self__).__init__(
            'aws:iot/thing:Thing',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            default_client_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            thing_type_name: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'Thing':
        """
        Get an existing Thing resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the thing.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attributes: Map of attributes of the thing.
        :param pulumi.Input[str] default_client_id: The default client ID.
        :param pulumi.Input[str] name: The name of the thing.
        :param pulumi.Input[str] thing_type_name: The thing type name.
        :param pulumi.Input[int] version: The current version of the thing record in the registry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["attributes"] = attributes
        __props__["default_client_id"] = default_client_id
        __props__["name"] = name
        __props__["thing_type_name"] = thing_type_name
        __props__["version"] = version
        return Thing(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the thing.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of attributes of the thing.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="defaultClientId")
    def default_client_id(self) -> pulumi.Output[str]:
        """
        The default client ID.
        """
        return pulumi.get(self, "default_client_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the thing.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="thingTypeName")
    def thing_type_name(self) -> pulumi.Output[Optional[str]]:
        """
        The thing type name.
        """
        return pulumi.get(self, "thing_type_name")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The current version of the thing record in the registry.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

