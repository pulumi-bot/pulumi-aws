# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Thing(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the thing.
    """
    attributes: pulumi.Output[dict]
    """
    Map of attributes of the thing.
    """
    default_client_id: pulumi.Output[str]
    """
    The default client ID.
    """
    name: pulumi.Output[str]
    """
    The name of the thing.
    """
    thing_type_name: pulumi.Output[str]
    """
    The thing type name.
    """
    version: pulumi.Output[float]
    """
    The current version of the thing record in the registry.
    """
    def __init__(__self__, resource_name, opts=None, attributes=None, name=None, thing_type_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates and manages an AWS IoT Thing.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attributes: Map of attributes of the thing.
        :param pulumi.Input[str] name: The name of the thing.
        :param pulumi.Input[str] thing_type_name: The thing type name.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing.html.markdown.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, arn=None, attributes=None, default_client_id=None, name=None, thing_type_name=None, version=None):
        """
        Get an existing Thing resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the thing.
        :param pulumi.Input[dict] attributes: Map of attributes of the thing.
        :param pulumi.Input[str] default_client_id: The default client ID.
        :param pulumi.Input[str] name: The name of the thing.
        :param pulumi.Input[str] thing_type_name: The thing type name.
        :param pulumi.Input[float] version: The current version of the thing record in the registry.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing.html.markdown.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

