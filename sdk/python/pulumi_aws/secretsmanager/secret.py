# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Secret(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the secret.
    """
    description: pulumi.Output[str]
    """
    A description of the secret.
    """
    kms_key_id: pulumi.Output[str]
    """
    Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
    """
    name: pulumi.Output[str]
    """
    Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified prefix. Conflicts with `name`.
    """
    policy: pulumi.Output[str]
    """
    A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
    """
    recovery_window_in_days: pulumi.Output[float]
    """
    Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
    """
    rotation_enabled: pulumi.Output[bool]
    """
    Specifies whether automatic rotation is enabled for this secret.
    """
    rotation_lambda_arn: pulumi.Output[str]
    """
    Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
    """
    rotation_rules: pulumi.Output[dict]
    """
    A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.

      * `automaticallyAfterDays` (`float`) - Specifies the number of days between automatic scheduled rotations of the secret.
    """
    tags: pulumi.Output[dict]
    """
    Specifies a key-value map of user-defined tags that are attached to the secret.
    """
    def __init__(__self__, resource_name, opts=None, description=None, kms_key_id=None, name=None, name_prefix=None, policy=None, recovery_window_in_days=None, rotation_lambda_arn=None, rotation_rules=None, tags=None, __props__=None, __name__=None, __opts__=None):
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
            rotation_rules={
                "automaticallyAfterDays": 7,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the secret.
        :param pulumi.Input[str] kms_key_id: Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        :param pulumi.Input[str] name: Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        :param pulumi.Input[float] recovery_window_in_days: Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[dict] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[dict] tags: Specifies a key-value map of user-defined tags that are attached to the secret.

        The **rotation_rules** object supports the following:

          * `automaticallyAfterDays` (`pulumi.Input[float]`) - Specifies the number of days between automatic scheduled rotations of the secret.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['kmsKeyId'] = kms_key_id
            __props__['name'] = name
            __props__['namePrefix'] = name_prefix
            __props__['policy'] = policy
            __props__['recoveryWindowInDays'] = recovery_window_in_days
            if rotation_lambda_arn is not None:
                warnings.warn("Use the aws_secretsmanager_secret_rotation resource instead", DeprecationWarning)
                pulumi.log.warn("rotation_lambda_arn is deprecated: Use the aws_secretsmanager_secret_rotation resource instead")
            __props__['rotationLambdaArn'] = rotation_lambda_arn
            if rotation_rules is not None:
                warnings.warn("Use the aws_secretsmanager_secret_rotation resource instead", DeprecationWarning)
                pulumi.log.warn("rotation_rules is deprecated: Use the aws_secretsmanager_secret_rotation resource instead")
            __props__['rotationRules'] = rotation_rules
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['rotation_enabled'] = None
        super(Secret, __self__).__init__(
            'aws:secretsmanager/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, kms_key_id=None, name=None, name_prefix=None, policy=None, recovery_window_in_days=None, rotation_enabled=None, rotation_lambda_arn=None, rotation_rules=None, tags=None):
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the secret.
        :param pulumi.Input[str] description: A description of the secret.
        :param pulumi.Input[str] kms_key_id: Specifies the ARN or Id of the AWS KMS customer master key (CMK) to be used to encrypt the secret values in the versions stored in this secret. If you don't specify this value, then Secrets Manager defaults to using the AWS account's default CMK (the one named `aws/secretsmanager`). If the default KMS CMK with that name doesn't yet exist, then AWS Secrets Manager creates it for you automatically the first time.
        :param pulumi.Input[str] name: Specifies the friendly name of the new secret. The secret name can consist of uppercase letters, lowercase letters, digits, and any of the following characters: `/_+=.@-` Conflicts with `name_prefix`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: A valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
        :param pulumi.Input[float] recovery_window_in_days: Specifies the number of days that AWS Secrets Manager waits before it can delete the secret. This value can be `0` to force deletion without recovery or range from `7` to `30` days. The default value is `30`.
        :param pulumi.Input[bool] rotation_enabled: Specifies whether automatic rotation is enabled for this secret.
        :param pulumi.Input[str] rotation_lambda_arn: Specifies the ARN of the Lambda function that can rotate the secret. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[dict] rotation_rules: A structure that defines the rotation configuration for this secret. Defined below. Use the `secretsmanager.SecretRotation` resource to manage this configuration instead. As of version 2.67.0, removal of this configuration will no longer remove rotation due to supporting the new resource. Either import the new resource and remove the configuration or manually remove rotation.
        :param pulumi.Input[dict] tags: Specifies a key-value map of user-defined tags that are attached to the secret.

        The **rotation_rules** object supports the following:

          * `automaticallyAfterDays` (`pulumi.Input[float]`) - Specifies the number of days between automatic scheduled rotations of the secret.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
