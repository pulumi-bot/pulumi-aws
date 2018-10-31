# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Secret(pulumi.CustomResource):
    """
    Provides a resource to manage AWS Secrets Manager secret metadata. To manage a secret value, see the [`aws_secretsmanager_secret_version` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret_version.html).
    """
    def __init__(__self__, __name__, __opts__=None, description=None, kms_key_id=None, name=None, policy=None, recovery_window_in_days=None, rotation_lambda_arn=None, rotation_rules=None, tags=None):
        """Create a Secret resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['kmsKeyId'] = kms_key_id

        __props__['name'] = name

        __props__['policy'] = policy

        __props__['recoveryWindowInDays'] = recovery_window_in_days

        __props__['rotationLambdaArn'] = rotation_lambda_arn

        __props__['rotationRules'] = rotation_rules

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['rotation_enabled'] = None

        super(Secret, __self__).__init__(
            'aws:secretsmanager/secret:Secret',
            __name__,
            __props__,
            __opts__)

