# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SecretArgs', 'Secret']

@pulumi.input_type
class SecretArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input['SecretRotationRulesArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Secret resource.
        :param pulumi.Input[str] description: A description of the secret.
        :param pulumi.Input[str] kms_key_id: Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        :param pulumi.Input[str] name: Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        :param pulumi.Input[int] recovery_window_in_days: Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input['SecretRotationRulesArgs'] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies a key-value map of user-defined tags that are attached to the secret.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if recovery_window_in_days is not None:
            pulumi.set(__self__, "recovery_window_in_days", recovery_window_in_days)
        if rotation_lambda_arn is not None:
            warnings.warn("""Use the aws_secretsmanager_secret_rotation resource instead""", DeprecationWarning)
            pulumi.log.warn("""rotation_lambda_arn is deprecated: Use the aws_secretsmanager_secret_rotation resource instead""")
        if rotation_lambda_arn is not None:
            pulumi.set(__self__, "rotation_lambda_arn", rotation_lambda_arn)
        if rotation_rules is not None:
            warnings.warn("""Use the aws_secretsmanager_secret_rotation resource instead""", DeprecationWarning)
            pulumi.log.warn("""rotation_rules is deprecated: Use the aws_secretsmanager_secret_rotation resource instead""")
        if rotation_rules is not None:
            pulumi.set(__self__, "rotation_rules", rotation_rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the secret.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="recoveryWindowInDays")
    def recovery_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        """
        return pulumi.get(self, "recovery_window_in_days")

    @recovery_window_in_days.setter
    def recovery_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "recovery_window_in_days", value)

    @property
    @pulumi.getter(name="rotationLambdaArn")
    def rotation_lambda_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        """
        return pulumi.get(self, "rotation_lambda_arn")

    @rotation_lambda_arn.setter
    def rotation_lambda_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rotation_lambda_arn", value)

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> Optional[pulumi.Input['SecretRotationRulesArgs']]:
        """
        A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        """
        return pulumi.get(self, "rotation_rules")

    @rotation_rules.setter
    def rotation_rules(self, value: Optional[pulumi.Input['SecretRotationRulesArgs']]):
        pulumi.set(self, "rotation_rules", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies a key-value map of user-defined tags that are attached to the secret.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Secret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage AWS Secrets Manager secret metadata. To manage secret rotation, see the `secretsmanager.SecretRotation` resource. To manage a secret value, see the `secretsmanager.SecretVersion` resource.

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.secretsmanager.Secret("example")
        ```
        ### Rotation Configuration

        To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g. RDS) or deploying a custom Lambda function.

        > **NOTE:** Configuring rotation causes the secret to rotate once as soon as you store the secret. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.

        > **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.

        ```python
        import pulumi
        import pulumi_aws as aws

        rotation_example = aws.secretsmanager.Secret("rotation-example",
            rotation_lambda_arn=aws_lambda_function["example"]["arn"],
            rotation_rules=aws.secretsmanager.SecretRotationRulesArgs(
                automatically_after_days=7,
            ))
        ```

        ## Import

        `aws_secretsmanager_secret` can be imported by using the secret Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:secretsmanager/secret:Secret example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
        ```

        :param str resource_name: The name of the resource.
        :param SecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input[pulumi.InputType['SecretRotationRulesArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage AWS Secrets Manager secret metadata. To manage secret rotation, see the `secretsmanager.SecretRotation` resource. To manage a secret value, see the `secretsmanager.SecretVersion` resource.

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.secretsmanager.Secret("example")
        ```
        ### Rotation Configuration

        To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g. RDS) or deploying a custom Lambda function.

        > **NOTE:** Configuring rotation causes the secret to rotate once as soon as you store the secret. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.

        > **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.

        ```python
        import pulumi
        import pulumi_aws as aws

        rotation_example = aws.secretsmanager.Secret("rotation-example",
            rotation_lambda_arn=aws_lambda_function["example"]["arn"],
            rotation_rules=aws.secretsmanager.SecretRotationRulesArgs(
                automatically_after_days=7,
            ))
        ```

        ## Import

        `aws_secretsmanager_secret` can be imported by using the secret Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:secretsmanager/secret:Secret example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the secret.
        :param pulumi.Input[str] kms_key_id: Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        :param pulumi.Input[str] name: Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        :param pulumi.Input[int] recovery_window_in_days: Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[pulumi.InputType['SecretRotationRulesArgs']] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies a key-value map of user-defined tags that are attached to the secret.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
                 rotation_rules: Optional[pulumi.Input[pulumi.InputType['SecretRotationRulesArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['description'] = description
            __props__['kms_key_id'] = kms_key_id
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['policy'] = policy
            __props__['recovery_window_in_days'] = recovery_window_in_days
            if rotation_lambda_arn is not None and not opts.urn:
                warnings.warn("""Use the aws_secretsmanager_secret_rotation resource instead""", DeprecationWarning)
                pulumi.log.warn("""rotation_lambda_arn is deprecated: Use the aws_secretsmanager_secret_rotation resource instead""")
            __props__['rotation_lambda_arn'] = rotation_lambda_arn
            if rotation_rules is not None and not opts.urn:
                warnings.warn("""Use the aws_secretsmanager_secret_rotation resource instead""", DeprecationWarning)
                pulumi.log.warn("""rotation_rules is deprecated: Use the aws_secretsmanager_secret_rotation resource instead""")
            __props__['rotation_rules'] = rotation_rules
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['rotation_enabled'] = None
        super(Secret, __self__).__init__(
            'aws:secretsmanager/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            recovery_window_in_days: Optional[pulumi.Input[int]] = None,
            rotation_enabled: Optional[pulumi.Input[bool]] = None,
            rotation_lambda_arn: Optional[pulumi.Input[str]] = None,
            rotation_rules: Optional[pulumi.Input[pulumi.InputType['SecretRotationRulesArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Secret':
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the secret.
        :param pulumi.Input[str] description: A description of the secret.
        :param pulumi.Input[str] kms_key_id: Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        :param pulumi.Input[str] name: Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        :param pulumi.Input[int] recovery_window_in_days: Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        :param pulumi.Input[bool] rotation_enabled: Specifies whether automatic rotation is enabled for this secret.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[pulumi.InputType['SecretRotationRulesArgs']] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies a key-value map of user-defined tags that are attached to the secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["kms_key_id"] = kms_key_id
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["policy"] = policy
        __props__["recovery_window_in_days"] = recovery_window_in_days
        __props__["rotation_enabled"] = rotation_enabled
        __props__["rotation_lambda_arn"] = rotation_lambda_arn
        __props__["rotation_rules"] = rotation_rules
        __props__["tags"] = tags
        return Secret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the secret.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the secret.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="recoveryWindowInDays")
    def recovery_window_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        """
        return pulumi.get(self, "recovery_window_in_days")

    @property
    @pulumi.getter(name="rotationEnabled")
    def rotation_enabled(self) -> pulumi.Output[bool]:
        """
        Specifies whether automatic rotation is enabled for this secret.
        """
        return pulumi.get(self, "rotation_enabled")

    @property
    @pulumi.getter(name="rotationLambdaArn")
    def rotation_lambda_arn(self) -> pulumi.Output[str]:
        """
        Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        """
        return pulumi.get(self, "rotation_lambda_arn")

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> pulumi.Output['outputs.SecretRotationRules']:
        """
        A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        """
        return pulumi.get(self, "rotation_rules")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies a key-value map of user-defined tags that are attached to the secret.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

