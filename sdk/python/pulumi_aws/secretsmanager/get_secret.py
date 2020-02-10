# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, arn=None, description=None, id=None, kms_key_id=None, name=None, policy=None, rotation_enabled=None, rotation_lambda_arn=None, rotation_rules=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the secret.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A description of the secret.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        __self__.kms_key_id = kms_key_id
        """
        The Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        __self__.policy = policy
        """
        The resource-based policy document that's attached to the secret.
        """
        if rotation_enabled and not isinstance(rotation_enabled, bool):
            raise TypeError("Expected argument 'rotation_enabled' to be a bool")
        __self__.rotation_enabled = rotation_enabled
        """
        Whether rotation is enabled or not.
        """
        if rotation_lambda_arn and not isinstance(rotation_lambda_arn, str):
            raise TypeError("Expected argument 'rotation_lambda_arn' to be a str")
        __self__.rotation_lambda_arn = rotation_lambda_arn
        """
        Rotation Lambda function Amazon Resource Name (ARN) if rotation is enabled.
        """
        if rotation_rules and not isinstance(rotation_rules, list):
            raise TypeError("Expected argument 'rotation_rules' to be a list")
        __self__.rotation_rules = rotation_rules
        """
        Rotation rules if rotation is enabled.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Tags of the secret.
        """
class AwaitableGetSecretResult(GetSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretResult(
            arn=self.arn,
            description=self.description,
            id=self.id,
            kms_key_id=self.kms_key_id,
            name=self.name,
            policy=self.policy,
            rotation_enabled=self.rotation_enabled,
            rotation_lambda_arn=self.rotation_lambda_arn,
            rotation_rules=self.rotation_rules,
            tags=self.tags)

def get_secret(arn=None,name=None,opts=None):
    """
    Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the [`secretsmanager.SecretVersion` data source](https://www.terraform.io/docs/providers/aws/d/secretsmanager_secret_version.html).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/secretsmanager_secret.html.markdown.


    :param str arn: The Amazon Resource Name (ARN) of the secret to retrieve.
    :param str name: The name of the secret to retrieve.
    """
    __args__ = dict()


    __args__['arn'] = arn
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:secretsmanager/getSecret:getSecret', __args__, opts=opts).value

    return AwaitableGetSecretResult(
        arn=__ret__.get('arn'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        kms_key_id=__ret__.get('kmsKeyId'),
        name=__ret__.get('name'),
        policy=__ret__.get('policy'),
        rotation_enabled=__ret__.get('rotationEnabled'),
        rotation_lambda_arn=__ret__.get('rotationLambdaArn'),
        rotation_rules=__ret__.get('rotationRules'),
        tags=__ret__.get('tags'))
