# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Policy(pulumi.CustomResource):
    """
    Provides an IAM policy.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, name=None, name_prefix=None, path=None, policy=None):
        """Create a Policy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        Description of the IAM policy.
        """
        __props__['description'] = description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the policy. If omitted, Terraform will assign a random, unique name.
        """
        __props__['name'] = name

        if name_prefix and not isinstance(name_prefix, basestring):
            raise TypeError('Expected property name_prefix to be a basestring')
        __self__.name_prefix = name_prefix
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        __props__['namePrefix'] = name_prefix

        if path and not isinstance(path, basestring):
            raise TypeError('Expected property path to be a basestring')
        __self__.path = path
        """
        Path in which to create the policy.
        See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
        """
        __props__['path'] = path

        if not policy:
            raise TypeError('Missing required property policy')
        elif not isinstance(policy, basestring):
            raise TypeError('Expected property policy to be a basestring')
        __self__.policy = policy
        """
        The policy document. This is a JSON formatted string.
        The heredoc syntax, `file` function, or the [`aws_iam_policy_document` data
        source](https://www.terraform.io/docs/providers/aws/d/iam_policy_document.html)
        are all helpful here.
        """
        __props__['policy'] = policy

        __self__.arn = pulumi.runtime.UNKNOWN
        """
        The ARN assigned by AWS to this policy.
        """

        super(Policy, __self__).__init__(
            'aws:iam/policy:Policy',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'arn' in outs:
            self.arn = outs['arn']
        if 'description' in outs:
            self.description = outs['description']
        if 'name' in outs:
            self.name = outs['name']
        if 'namePrefix' in outs:
            self.name_prefix = outs['namePrefix']
        if 'path' in outs:
            self.path = outs['path']
        if 'policy' in outs:
            self.policy = outs['policy']
