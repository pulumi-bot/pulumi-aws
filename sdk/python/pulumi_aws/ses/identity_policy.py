# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IdentityPolicy(pulumi.CustomResource):
    identity: pulumi.Output[str]
    """
    Name or Amazon Resource Name (ARN) of the SES Identity.
    """
    name: pulumi.Output[str]
    """
    Name of the policy.
    """
    policy: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, identity=None, name=None, policy=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a SES Identity Policy. More information about SES Sending Authorization Policies can be found in the [SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_identity_policy.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity: Name or Amazon Resource Name (ARN) of the SES Identity.
        :param pulumi.Input[str] name: Name of the policy.
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

            if identity is None:
                raise TypeError("Missing required property 'identity'")
            __props__['identity'] = identity
            __props__['name'] = name
            if policy is None:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
        super(IdentityPolicy, __self__).__init__(
            'aws:ses/identityPolicy:IdentityPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, identity=None, name=None, policy=None):
        """
        Get an existing IdentityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity: Name or Amazon Resource Name (ARN) of the SES Identity.
        :param pulumi.Input[str] name: Name of the policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["identity"] = identity
        __props__["name"] = name
        __props__["policy"] = policy
        return IdentityPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

