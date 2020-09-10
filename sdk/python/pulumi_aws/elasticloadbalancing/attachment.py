# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Attachment']

warnings.warn("aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment", DeprecationWarning)


class Attachment(pulumi.CustomResource):
    warnings.warn("aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 elb: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Attachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        pulumi.log.warn("Attachment is deprecated: aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment")
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

            if elb is None:
                raise TypeError("Missing required property 'elb'")
            __props__['elb'] = elb
            if instance is None:
                raise TypeError("Missing required property 'instance'")
            __props__['instance'] = instance
        super(Attachment, __self__).__init__(
            'aws:elasticloadbalancing/attachment:Attachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            elb: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None) -> 'Attachment':
        """
        Get an existing Attachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["elb"] = elb
        __props__["instance"] = instance
        return Attachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def elb(self) -> pulumi.Output[str]:
        return pulumi.get(self, "elb")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

