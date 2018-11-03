# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Trigger(pulumi.CustomResource):
    """
    Manages a Glue Trigger resource.
    """
    def __init__(__self__, __name__, __opts__=None, actions=None, description=None, enabled=None, name=None, predicate=None, schedule=None, type=None):
        """Create a Trigger resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not actions:
            raise TypeError('Missing required property actions')
        __props__['actions'] = actions

        __props__['description'] = description

        __props__['enabled'] = enabled

        __props__['name'] = name

        __props__['predicate'] = predicate

        __props__['schedule'] = schedule

        if not type:
            raise TypeError('Missing required property type')
        __props__['type'] = type

        super(Trigger, __self__).__init__(
            'aws:glue/trigger:Trigger',
            __name__,
            __props__,
            __opts__)

