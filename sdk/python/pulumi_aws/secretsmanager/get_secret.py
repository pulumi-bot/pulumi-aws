# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetSecretResult',
    'AwaitableGetSecretResult',
    'get_secret',
]

@pulumi.output_type
class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, arn=None, description=None, id=None, kms_key_id=None, name=None, policy=None, rotation_enabled=None, rotation_lambda_arn=None, rotation_rules=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)
        if rotation_enabled and not isinstance(rotation_enabled, bool):
            raise TypeError("Expected argument 'rotation_enabled' to be a bool")
        if rotation_enabled is not None:
            warnings.warn("Use the aws_secretsmanager_secret_rotation data source instead", DeprecationWarning)
            pulumi.log.warn("rotation_enabled is deprecated: Use the aws_secretsmanager_secret_rotation data source instead")

        pulumi.set(__self__, "rotation_enabled", rotation_enabled)
        if rotation_lambda_arn and not isinstance(rotation_lambda_arn, str):
            raise TypeError("Expected argument 'rotation_lambda_arn' to be a str")
        if rotation_lambda_arn is not None:
            warnings.warn("Use the aws_secretsmanager_secret_rotation data source instead", DeprecationWarning)
            pulumi.log.warn("rotation_lambda_arn is deprecated: Use the aws_secretsmanager_secret_rotation data source instead")

        pulumi.set(__self__, "rotation_lambda_arn", rotation_lambda_arn)
        if rotation_rules and not isinstance(rotation_rules, list):
            raise TypeError("Expected argument 'rotation_rules' to be a list")
        if rotation_rules is not None:
            warnings.warn("Use the aws_secretsmanager_secret_rotation data source instead", DeprecationWarning)
            pulumi.log.warn("rotation_rules is deprecated: Use the aws_secretsmanager_secret_rotation data source instead")

        pulumi.set(__self__, "rotation_rules", rotation_rules)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the secret.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the secret.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        """
        The Key Management Service (KMS) Customer Master Key (CMK) associated with the secret.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policy(self) -> str:
        """
        The resource-based policy document that's attached to the secret.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="rotationEnabled")
    def rotation_enabled(self) -> bool:
        """
        Whether rotation is enabled or not.
        """
        return pulumi.get(self, "rotation_enabled")

    @property
    @pulumi.getter(name="rotationLambdaArn")
    def rotation_lambda_arn(self) -> str:
        """
        Rotation Lambda function Amazon Resource Name (ARN) if rotation is enabled.
        """
        return pulumi.get(self, "rotation_lambda_arn")

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> Sequence['outputs.GetSecretRotationRuleResult']:
        """
        Rotation rules if rotation is enabled.
        """
        return pulumi.get(self, "rotation_rules")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Tags of the secret.
        """
        return pulumi.get(self, "tags")


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


def get_secret(arn: Optional[str] = None,
               name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretResult:
    """
    Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the `secretsmanager.SecretVersion`.

    ## Example Usage
    ### ARN

    ```python
    import pulumi
    import pulumi_aws as aws

    by_arn = aws.secretsmanager.get_secret(arn="arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456")
    ```
    ### Name

    ```python
    import pulumi
    import pulumi_aws as aws

    by_name = aws.secretsmanager.get_secret(name="example")
    ```


    :param str arn: The Amazon Resource Name (ARN) of the secret to retrieve.
    :param str name: The name of the secret to retrieve.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:secretsmanager/getSecret:getSecret', __args__, opts=opts, typ=GetSecretResult).value

    return AwaitableGetSecretResult(
        arn=__ret__.arn,
        description=__ret__.description,
        id=__ret__.id,
        kms_key_id=__ret__.kms_key_id,
        name=__ret__.name,
        policy=__ret__.policy,
        rotation_enabled=__ret__.rotation_enabled,
        rotation_lambda_arn=__ret__.rotation_lambda_arn,
        rotation_rules=__ret__.rotation_rules,
        tags=__ret__.tags)
