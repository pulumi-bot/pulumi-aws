# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Activation(pulumi.CustomResource):
    activation_code: pulumi.Output[str]
    """
    The code the system generates when it processes the activation.
    """
    description: pulumi.Output[str]
    """
    The description of the resource that you want to register.
    """
    expiration_date: pulumi.Output[str]
    """
    A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time.
    """
    expired: pulumi.Output[str]
    """
    If the current activation has expired.
    """
    iam_role: pulumi.Output[str]
    """
    The IAM Role to attach to the managed instance.
    """
    name: pulumi.Output[str]
    """
    The default name of the registered managed instance.
    """
    registration_count: pulumi.Output[float]
    """
    The number of managed instances that are currently registered using this activation.
    """
    registration_limit: pulumi.Output[float]
    """
    The maximum number of managed instances you want to register. The default value is 1 instance.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the object.
    """
    def __init__(__self__, resource_name, opts=None, description=None, expiration_date=None, iam_role=None, name=None, registration_limit=None, tags=None, __name__=None, __opts__=None):
        """
        Registers an on-premises server or virtual machine with Amazon EC2 so that it can be managed using Run Command.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the resource that you want to register.
        :param pulumi.Input[str] expiration_date: A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time.
        :param pulumi.Input[str] iam_role: The IAM Role to attach to the managed instance.
        :param pulumi.Input[str] name: The default name of the registered managed instance.
        :param pulumi.Input[float] registration_limit: The maximum number of managed instances you want to register. The default value is 1 instance.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the object.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_activation.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description
        __props__['expiration_date'] = expiration_date
        if iam_role is None:
            raise TypeError("Missing required property 'iam_role'")
        __props__['iam_role'] = iam_role
        __props__['name'] = name
        __props__['registration_limit'] = registration_limit
        __props__['tags'] = tags
        __props__['activation_code'] = None
        __props__['expired'] = None
        __props__['registration_count'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Activation, __self__).__init__(
            'aws:ssm/activation:Activation',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

