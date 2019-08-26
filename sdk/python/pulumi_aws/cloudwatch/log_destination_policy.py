# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class LogDestinationPolicy(pulumi.CustomResource):
    access_policy: pulumi.Output[str]
    """
    The policy document. This is a JSON formatted string.
    """
    destination_name: pulumi.Output[str]
    """
    A name for the subscription filter
    """
    def __init__(__self__, resource_name, opts=None, access_policy=None, destination_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Logs destination policy resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] destination_name: A name for the subscription filter

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination_policy.html.markdown.
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

            if access_policy is None:
                raise TypeError("Missing required property 'access_policy'")
            __props__['access_policy'] = access_policy
            if destination_name is None:
                raise TypeError("Missing required property 'destination_name'")
            __props__['destination_name'] = destination_name
        super(LogDestinationPolicy, __self__).__init__(
            'aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_policy=None, destination_name=None):
        """
        Get an existing LogDestinationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] destination_name: A name for the subscription filter

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination_policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["access_policy"] = access_policy
        __props__["destination_name"] = destination_name
        return LogDestinationPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

