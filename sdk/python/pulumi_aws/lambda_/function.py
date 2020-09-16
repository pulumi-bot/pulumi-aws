# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Function']


class Function(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 code: Optional[pulumi.Input[pulumi.Archive]] = None,
                 dead_letter_config: Optional[pulumi.Input[pulumi.InputType['FunctionDeadLetterConfigArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[pulumi.InputType['FunctionEnvironmentArgs']]] = None,
                 file_system_config: Optional[pulumi.Input[pulumi.InputType['FunctionFileSystemConfigArgs']]] = None,
                 handler: Optional[pulumi.Input[str]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 layers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 memory_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 publish: Optional[pulumi.Input[bool]] = None,
                 reserved_concurrent_executions: Optional[pulumi.Input[int]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 runtime: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_key: Optional[pulumi.Input[str]] = None,
                 s3_object_version: Optional[pulumi.Input[str]] = None,
                 source_code_hash: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 tracing_config: Optional[pulumi.Input[pulumi.InputType['FunctionTracingConfigArgs']]] = None,
                 vpc_config: Optional[pulumi.Input[pulumi.InputType['FunctionVpcConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Lambda Function resource. Lambda allows you to trigger execution of code in response to events in AWS, enabling serverless backend solutions. The Lambda Function itself includes source code and runtime configuration.

        For information about Lambda and how to use it, see [What is AWS Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)

        > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), EC2 subnets and security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.

        ## Example Usage
        ### Lambda Layers

        ```python
        import pulumi
        import pulumi_aws as aws

        example_layer_version = aws.lambda_.LayerVersion("exampleLayerVersion")
        # ... other configuration ...
        example_function = aws.lambda_.Function("exampleFunction", layers=[example_layer_version.arn])
        ```
        ### Lambda File Systems

        Lambda File Systems allow you to connect an Amazon Elastic File System (EFS) file system to a Lambda function to share data across function invocations, access existing data including large files, and save function state.

        ```python
        import pulumi
        import pulumi_aws as aws

        # EFS file system
        efs_for_lambda = aws.efs.FileSystem("efsForLambda", tags={
            "Name": "efs_for_lambda",
        })
        # Mount target connects the file system to the subnet
        alpha = aws.efs.MountTarget("alpha",
            file_system_id=efs_for_lambda.id,
            subnet_id=aws_subnet["subnet_for_lambda"]["id"],
            security_groups=[aws_security_group["sg_for_lambda"]["id"]])
        # EFS access point used by lambda file system
        access_point_for_lambda = aws.efs.AccessPoint("accessPointForLambda",
            file_system_id=efs_for_lambda.id,
            root_directory=aws.efs.AccessPointRootDirectoryArgs(
                path="/lambda",
                creation_info=aws.efs.AccessPointRootDirectoryCreationInfoArgs(
                    owner_gid=1000,
                    owner_uid=1000,
                    permissions="777",
                ),
            ),
            posix_user=aws.efs.AccessPointPosixUserArgs(
                gid=1000,
                uid=1000,
            ))
        # A lambda function connected to an EFS file system
        # ... other configuration ...
        example = aws.lambda_.Function("example",
            file_system_config=aws.lambda..FunctionFileSystemConfigArgs(
                arn=access_point_for_lambda.arn,
                local_mount_path="/mnt/efs",
            ),
            vpc_config=aws.lambda..FunctionVpcConfigArgs(
                subnet_ids=[aws_subnet["subnet_for_lambda"]["id"]],
                security_group_ids=[aws_security_group["sg_for_lambda"]["id"]],
            ),
            opts=ResourceOptions(depends_on=[alpha]))
        ```
        ### CloudWatch Logging and Permissions

        For more information about CloudWatch Logs for Lambda, see the [Lambda User Guide](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions-logs.html).

        ```python
        import pulumi
        import pulumi_aws as aws

        # This is to optionally manage the CloudWatch Log Group for the Lambda Function.
        # If skipping this resource configuration, also add "logs:CreateLogGroup" to the IAM policy below.
        example = aws.cloudwatch.LogGroup("example", retention_in_days=14)
        # See also the following AWS managed policy: AWSLambdaBasicExecutionRole
        lambda_logging = aws.iam.Policy("lambdaLogging",
            path="/",
            description="IAM policy for logging from a lambda",
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents"
              ],
              "Resource": "arn:aws:logs:*:*:*",
              "Effect": "Allow"
            }
          ]
        }
        \"\"\")
        lambda_logs = aws.iam.RolePolicyAttachment("lambdaLogs",
            role=aws_iam_role["iam_for_lambda"]["name"],
            policy_arn=lambda_logging.arn)
        test_lambda = aws.lambda_.Function("testLambda", opts=ResourceOptions(depends_on=[
                lambda_logs,
                example,
            ]))
        ```
        ## Specifying the Deployment Package

        AWS Lambda expects source code to be provided as a deployment package whose structure varies depending on which `runtime` is in use.
        See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for the valid values of `runtime`. The expected structure of the deployment package can be found in
        [the AWS Lambda documentation for each runtime](https://docs.aws.amazon.com/lambda/latest/dg/deployment-package-v2.html).

        Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
        indirectly via Amazon S3 (using the `s3_bucket`, `s3_key` and `s3_object_version` arguments). When providing the deployment
        package via S3 it may be useful to use the `s3.BucketObject` resource to upload it.

        For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading
        large files efficiently.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.Archive] code: The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        :param pulumi.Input[pulumi.InputType['FunctionDeadLetterConfigArgs']] dead_letter_config: Nested block to configure the function's *dead letter queue*. See details below.
        :param pulumi.Input[str] description: Description of what your Lambda Function does.
        :param pulumi.Input[pulumi.InputType['FunctionEnvironmentArgs']] environment: The Lambda environment's configuration settings. Fields documented below.
        :param pulumi.Input[pulumi.InputType['FunctionFileSystemConfigArgs']] file_system_config: The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `file_system_config`, EFS mount targets much be in available lifecycle state. Use `depends_on` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
        :param pulumi.Input[str] handler: The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        :param pulumi.Input[str] kms_key_arn: Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and this provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] layers: List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
        :param pulumi.Input[int] memory_size: Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        :param pulumi.Input[str] name: A unique name for your Lambda Function.
        :param pulumi.Input[bool] publish: Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        :param pulumi.Input[int] reserved_concurrent_executions: The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
        :param pulumi.Input[str] role: IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
        :param pulumi.Input[str] runtime: See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
        :param pulumi.Input[str] s3_bucket: The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        :param pulumi.Input[str] s3_key: The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] s3_object_version: The object version containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] source_code_hash: Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the object.
        :param pulumi.Input[int] timeout: The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        :param pulumi.Input[pulumi.InputType['FunctionVpcConfigArgs']] vpc_config: Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
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

            __props__['code'] = code
            __props__['dead_letter_config'] = dead_letter_config
            __props__['description'] = description
            __props__['environment'] = environment
            __props__['file_system_config'] = file_system_config
            if handler is None:
                raise TypeError("Missing required property 'handler'")
            __props__['handler'] = handler
            __props__['kms_key_arn'] = kms_key_arn
            __props__['layers'] = layers
            __props__['memory_size'] = memory_size
            __props__['name'] = name
            __props__['publish'] = publish
            __props__['reserved_concurrent_executions'] = reserved_concurrent_executions
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if runtime is None:
                raise TypeError("Missing required property 'runtime'")
            __props__['runtime'] = runtime
            __props__['s3_bucket'] = s3_bucket
            __props__['s3_key'] = s3_key
            __props__['s3_object_version'] = s3_object_version
            __props__['source_code_hash'] = source_code_hash
            __props__['tags'] = tags
            __props__['timeout'] = timeout
            __props__['tracing_config'] = tracing_config
            __props__['vpc_config'] = vpc_config
            __props__['arn'] = None
            __props__['invoke_arn'] = None
            __props__['last_modified'] = None
            __props__['qualified_arn'] = None
            __props__['source_code_size'] = None
            __props__['version'] = None
        super(Function, __self__).__init__(
            'aws:lambda/function:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            code: Optional[pulumi.Input[pulumi.Archive]] = None,
            dead_letter_config: Optional[pulumi.Input[pulumi.InputType['FunctionDeadLetterConfigArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[pulumi.InputType['FunctionEnvironmentArgs']]] = None,
            file_system_config: Optional[pulumi.Input[pulumi.InputType['FunctionFileSystemConfigArgs']]] = None,
            handler: Optional[pulumi.Input[str]] = None,
            invoke_arn: Optional[pulumi.Input[str]] = None,
            kms_key_arn: Optional[pulumi.Input[str]] = None,
            last_modified: Optional[pulumi.Input[str]] = None,
            layers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            memory_size: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            publish: Optional[pulumi.Input[bool]] = None,
            qualified_arn: Optional[pulumi.Input[str]] = None,
            reserved_concurrent_executions: Optional[pulumi.Input[int]] = None,
            role: Optional[pulumi.Input[str]] = None,
            runtime: Optional[pulumi.Input[str]] = None,
            s3_bucket: Optional[pulumi.Input[str]] = None,
            s3_key: Optional[pulumi.Input[str]] = None,
            s3_object_version: Optional[pulumi.Input[str]] = None,
            source_code_hash: Optional[pulumi.Input[str]] = None,
            source_code_size: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            tracing_config: Optional[pulumi.Input[pulumi.InputType['FunctionTracingConfigArgs']]] = None,
            version: Optional[pulumi.Input[str]] = None,
            vpc_config: Optional[pulumi.Input[pulumi.InputType['FunctionVpcConfigArgs']]] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        :param pulumi.Input[pulumi.Archive] code: The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        :param pulumi.Input[pulumi.InputType['FunctionDeadLetterConfigArgs']] dead_letter_config: Nested block to configure the function's *dead letter queue*. See details below.
        :param pulumi.Input[str] description: Description of what your Lambda Function does.
        :param pulumi.Input[pulumi.InputType['FunctionEnvironmentArgs']] environment: The Lambda environment's configuration settings. Fields documented below.
        :param pulumi.Input[pulumi.InputType['FunctionFileSystemConfigArgs']] file_system_config: The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `file_system_config`, EFS mount targets much be in available lifecycle state. Use `depends_on` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
        :param pulumi.Input[str] handler: The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        :param pulumi.Input[str] invoke_arn: The ARN to be used for invoking Lambda Function from API Gateway - to be used in `apigateway.Integration`'s `uri`
        :param pulumi.Input[str] kms_key_arn: Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and this provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
        :param pulumi.Input[str] last_modified: The date this resource was last modified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] layers: List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
        :param pulumi.Input[int] memory_size: Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        :param pulumi.Input[str] name: A unique name for your Lambda Function.
        :param pulumi.Input[bool] publish: Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        :param pulumi.Input[str] qualified_arn: The Amazon Resource Name (ARN) identifying your Lambda Function Version
               (if versioning is enabled via `publish = true`).
        :param pulumi.Input[int] reserved_concurrent_executions: The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
        :param pulumi.Input[str] role: IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
        :param pulumi.Input[str] runtime: See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
        :param pulumi.Input[str] s3_bucket: The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        :param pulumi.Input[str] s3_key: The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] s3_object_version: The object version containing the function's deployment package. Conflicts with `filename`.
        :param pulumi.Input[str] source_code_hash: Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        :param pulumi.Input[int] source_code_size: The size in bytes of the function .zip file.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the object.
        :param pulumi.Input[int] timeout: The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        :param pulumi.Input[str] version: Latest published version of your Lambda Function.
        :param pulumi.Input[pulumi.InputType['FunctionVpcConfigArgs']] vpc_config: Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["code"] = code
        __props__["dead_letter_config"] = dead_letter_config
        __props__["description"] = description
        __props__["environment"] = environment
        __props__["file_system_config"] = file_system_config
        __props__["handler"] = handler
        __props__["invoke_arn"] = invoke_arn
        __props__["kms_key_arn"] = kms_key_arn
        __props__["last_modified"] = last_modified
        __props__["layers"] = layers
        __props__["memory_size"] = memory_size
        __props__["name"] = name
        __props__["publish"] = publish
        __props__["qualified_arn"] = qualified_arn
        __props__["reserved_concurrent_executions"] = reserved_concurrent_executions
        __props__["role"] = role
        __props__["runtime"] = runtime
        __props__["s3_bucket"] = s3_bucket
        __props__["s3_key"] = s3_key
        __props__["s3_object_version"] = s3_object_version
        __props__["source_code_hash"] = source_code_hash
        __props__["source_code_size"] = source_code_size
        __props__["tags"] = tags
        __props__["timeout"] = timeout
        __props__["tracing_config"] = tracing_config
        __props__["version"] = version
        __props__["vpc_config"] = vpc_config
        return Function(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def code(self) -> pulumi.Output[Optional[pulumi.Archive]]:
        """
        The path to the function's deployment package within the local filesystem. If defined, The `s3_`-prefixed options cannot be used.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="deadLetterConfig")
    def dead_letter_config(self) -> pulumi.Output[Optional['outputs.FunctionDeadLetterConfig']]:
        """
        Nested block to configure the function's *dead letter queue*. See details below.
        """
        return pulumi.get(self, "dead_letter_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of what your Lambda Function does.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional['outputs.FunctionEnvironment']]:
        """
        The Lambda environment's configuration settings. Fields documented below.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="fileSystemConfig")
    def file_system_config(self) -> pulumi.Output[Optional['outputs.FunctionFileSystemConfig']]:
        """
        The connection settings for an EFS file system. Fields documented below. Before creating or updating Lambda functions with `file_system_config`, EFS mount targets much be in available lifecycle state. Use `depends_on` to explicitly declare this dependency. See [Using Amazon EFS with Lambda](https://docs.aws.amazon.com/lambda/latest/dg/services-efs.html).
        """
        return pulumi.get(self, "file_system_config")

    @property
    @pulumi.getter
    def handler(self) -> pulumi.Output[str]:
        """
        The function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
        """
        return pulumi.get(self, "handler")

    @property
    @pulumi.getter(name="invokeArn")
    def invoke_arn(self) -> pulumi.Output[str]:
        """
        The ARN to be used for invoking Lambda Function from API Gateway - to be used in `apigateway.Integration`'s `uri`
        """
        return pulumi.get(self, "invoke_arn")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables. If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key. If this configuration is provided when environment variables are not in use, the AWS Lambda API does not save this configuration and this provider will show a perpetual difference of adding the key. To fix the perpetual difference, remove this configuration.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        The date this resource was last modified.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def layers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function. See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
        """
        return pulumi.get(self, "layers")

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> pulumi.Output[Optional[int]]:
        """
        Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        """
        return pulumi.get(self, "memory_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for your Lambda Function.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def publish(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
        """
        return pulumi.get(self, "publish")

    @property
    @pulumi.getter(name="qualifiedArn")
    def qualified_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) identifying your Lambda Function Version
        (if versioning is enabled via `publish = true`).
        """
        return pulumi.get(self, "qualified_arn")

    @property
    @pulumi.getter(name="reservedConcurrentExecutions")
    def reserved_concurrent_executions(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`. See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
        """
        return pulumi.get(self, "reserved_concurrent_executions")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        IAM role attached to the Lambda Function. This governs both who / what can invoke your Lambda Function, as well as what resources our Lambda Function has access to. See [Lambda Permission Model](https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html) for more details.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def runtime(self) -> pulumi.Output[str]:
        """
        See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Output[Optional[str]]:
        """
        The S3 bucket location containing the function's deployment package. Conflicts with `filename`. This bucket must reside in the same AWS region where you are creating the Lambda function.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Key")
    def s3_key(self) -> pulumi.Output[Optional[str]]:
        """
        The S3 key of an object containing the function's deployment package. Conflicts with `filename`.
        """
        return pulumi.get(self, "s3_key")

    @property
    @pulumi.getter(name="s3ObjectVersion")
    def s3_object_version(self) -> pulumi.Output[Optional[str]]:
        """
        The object version containing the function's deployment package. Conflicts with `filename`.
        """
        return pulumi.get(self, "s3_object_version")

    @property
    @pulumi.getter(name="sourceCodeHash")
    def source_code_hash(self) -> pulumi.Output[str]:
        """
        Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either `filename` or `s3_key`. The usual way to set this is `filebase64sha256("file.zip")` (this provider 0.11.12 and later) or `base64sha256(file("file.zip"))` (this provider 0.11.11 and earlier), where "file.zip" is the local filename of the lambda function source archive.
        """
        return pulumi.get(self, "source_code_hash")

    @property
    @pulumi.getter(name="sourceCodeSize")
    def source_code_size(self) -> pulumi.Output[int]:
        """
        The size in bytes of the function .zip file.
        """
        return pulumi.get(self, "source_code_size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the object.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of time your Lambda Function has to run in seconds. Defaults to `3`. See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="tracingConfig")
    def tracing_config(self) -> pulumi.Output['outputs.FunctionTracingConfig']:
        return pulumi.get(self, "tracing_config")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Latest published version of your Lambda Function.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> pulumi.Output[Optional['outputs.FunctionVpcConfig']]:
        """
        Provide this to allow your function to access your VPC. Fields documented below. See [Lambda in VPC](http://docs.aws.amazon.com/lambda/latest/dg/vpc.html)
        """
        return pulumi.get(self, "vpc_config")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

