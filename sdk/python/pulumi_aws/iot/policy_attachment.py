# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class PolicyAttachment(pulumi.CustomResource):
    policy: pulumi.Output[str]
    """
    The name of the policy to attach.
    """
    target: pulumi.Output[str]
    """
    The identity to which the policy is attached.
    """
    def __init__(__self__, resource_name, opts=None, policy=None, target=None, __name__=None, __opts__=None):
        """
        Provides an IoT policy attachment.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: The name of the policy to attach.
        :param pulumi.Input[str] target: The identity to which the policy is attached.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if policy is None:
            raise TypeError("Missing required property 'policy'")
        __props__['policy'] = policy

        if target is None:
            raise TypeError("Missing required property 'target'")
        __props__['target'] = target

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(PolicyAttachment, __self__).__init__(
            'aws:iot/policyAttachment:PolicyAttachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

