# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class MaintenanceWindowTask(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the maintenance window task.
    """
    logging_info: pulumi.Output[dict]
    """
    A structure containing information about an Amazon S3 bucket to write instance-level logs to. Use `task_invocation_parameters` configuration block `run_command_parameters` configuration block `output_s3_*` arguments instead. Conflicts with `task_invocation_parameters`. Documented below.

      * `s3_bucket_name` (`str`)
      * `s3BucketPrefix` (`str`)
      * `s3_region` (`str`)
    """
    max_concurrency: pulumi.Output[str]
    """
    The maximum number of targets this task can be run for in parallel.
    """
    max_errors: pulumi.Output[str]
    """
    The maximum number of errors allowed before this task stops being scheduled.
    """
    name: pulumi.Output[str]
    """
    The name of the maintenance window task.
    """
    priority: pulumi.Output[float]
    """
    The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
    """
    service_role_arn: pulumi.Output[str]
    """
    The role that should be assumed when executing the task.
    """
    targets: pulumi.Output[list]
    """
    The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.

      * `key` (`str`)
      * `values` (`list`)
    """
    task_arn: pulumi.Output[str]
    """
    The ARN of the task to execute.
    """
    task_invocation_parameters: pulumi.Output[dict]
    """
    The parameters for task execution. This argument is conflict with `task_parameters` and `logging_info`.

      * `automationParameters` (`dict`) - The parameters for an AUTOMATION task type. Documented below.
        * `document_version` (`str`) - The version of an Automation document to use during task execution.
        * `parameters` (`list`) - The parameters for the RUN_COMMAND task execution. Documented below.
          * `name` (`str`) - The parameter name.
          * `values` (`list`) - The array of strings.

      * `lambdaParameters` (`dict`) - The parameters for a LAMBDA task type. Documented below.
        * `clientContext` (`str`) - Pass client-specific information to the Lambda function that you are invoking.
        * `payload` (`str`) - JSON to provide to your Lambda function as input.
        * `qualifier` (`str`) - Specify a Lambda function version or alias name.

      * `runCommandParameters` (`dict`) - The parameters for a RUN_COMMAND task type. Documented below.
        * `comment` (`str`) - Information about the command(s) to execute.
        * `documentHash` (`str`) - The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
        * `documentHashType` (`str`) - SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
        * `notificationConfig` (`dict`) - Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
          * `notificationArn` (`str`) - An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
          * `notificationEvents` (`list`) - The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
          * `notification_type` (`str`) - When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`

        * `outputS3Bucket` (`str`) - The name of the Amazon S3 bucket.
        * `outputS3KeyPrefix` (`str`) - The Amazon S3 bucket subfolder.
        * `parameters` (`list`) - The parameters for the RUN_COMMAND task execution. Documented below.
          * `name` (`str`) - The parameter name.
          * `values` (`list`) - The array of strings.

        * `service_role_arn` (`str`) - The IAM service role to assume during task execution.
        * `timeoutSeconds` (`float`) - If this time is reached and the command has not already started executing, it doesn't run.

      * `stepFunctionsParameters` (`dict`) - The parameters for a STEP_FUNCTIONS task type. Documented below.
        * `input` (`str`) - The inputs for the STEP_FUNCTION task.
        * `name` (`str`) - The name of the STEP_FUNCTION task.
    """
    task_parameters: pulumi.Output[list]
    """
    A structure containing information about parameters required by the particular `task_arn`. Use `parameter` configuration blocks under the `task_invocation_parameters` configuration block instead. Conflicts with `task_invocation_parameters`. Documented below.

      * `name` (`str`) - The name of the maintenance window task.
      * `values` (`list`)
    """
    task_type: pulumi.Output[str]
    """
    The type of task being registered. The only allowed value is `RUN_COMMAND`.
    """
    window_id: pulumi.Output[str]
    """
    The Id of the maintenance window to register the task with.
    """
    def __init__(__self__, resource_name, opts=None, description=None, logging_info=None, max_concurrency=None, max_errors=None, name=None, priority=None, service_role_arn=None, targets=None, task_arn=None, task_invocation_parameters=None, task_parameters=None, task_type=None, window_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an SSM Maintenance Window Task resource

        ## Example Usage

        ### Automation Tasks

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssm.MaintenanceWindowTask("example",
            max_concurrency=2,
            max_errors=1,
            priority=1,
            service_role_arn=aws_iam_role["example"]["arn"],
            targets=[{
                "key": "InstanceIds",
                "values": [aws_instance["example"]["id"]],
            }],
            task_arn="AWS-RestartEC2Instance",
            task_invocation_parameters={
                "automationParameters": {
                    "document_version": "$$LATEST",
                    "parameter": [{
                        "name": "InstanceId",
                        "values": [aws_instance["example"]["id"]],
                    }],
                },
            },
            task_type="AUTOMATION",
            window_id=aws_ssm_maintenance_window["example"]["id"])
        ```

        ### Run Command Tasks

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssm.MaintenanceWindowTask("example",
            max_concurrency=2,
            max_errors=1,
            priority=1,
            service_role_arn=aws_iam_role["example"]["arn"],
            targets=[{
                "key": "InstanceIds",
                "values": [aws_instance["example"]["id"]],
            }],
            task_arn="AWS-RunShellScript",
            task_invocation_parameters={
                "runCommandParameters": {
                    "notificationConfig": {
                        "notificationArn": aws_sns_topic["example"]["arn"],
                        "notificationEvents": ["All"],
                        "notification_type": "Command",
                    },
                    "outputS3Bucket": aws_s3_bucket["example"]["bucket"],
                    "outputS3KeyPrefix": "output",
                    "parameter": [{
                        "name": "commands",
                        "values": ["date"],
                    }],
                    "service_role_arn": aws_iam_role["example"]["arn"],
                    "timeoutSeconds": 600,
                },
            },
            task_type="RUN_COMMAND",
            window_id=aws_ssm_maintenance_window["example"]["id"])
        ```

        ### Step Function Tasks

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssm.MaintenanceWindowTask("example",
            max_concurrency=2,
            max_errors=1,
            priority=1,
            service_role_arn=aws_iam_role["example"]["arn"],
            targets=[{
                "key": "InstanceIds",
                "values": [aws_instance["example"]["id"]],
            }],
            task_arn=aws_sfn_activity["example"]["id"],
            task_invocation_parameters={
                "stepFunctionsParameters": {
                    "input": "{\"key1\":\"value1\"}",
                    "name": "example",
                },
            },
            task_type="STEP_FUNCTIONS",
            window_id=aws_ssm_maintenance_window["example"]["id"])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window task.
        :param pulumi.Input[dict] logging_info: A structure containing information about an Amazon S3 bucket to write instance-level logs to. Use `task_invocation_parameters` configuration block `run_command_parameters` configuration block `output_s3_*` arguments instead. Conflicts with `task_invocation_parameters`. Documented below.
        :param pulumi.Input[str] max_concurrency: The maximum number of targets this task can be run for in parallel.
        :param pulumi.Input[str] max_errors: The maximum number of errors allowed before this task stops being scheduled.
        :param pulumi.Input[str] name: The name of the maintenance window task.
        :param pulumi.Input[float] priority: The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
        :param pulumi.Input[str] service_role_arn: The role that should be assumed when executing the task.
        :param pulumi.Input[list] targets: The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
        :param pulumi.Input[str] task_arn: The ARN of the task to execute.
        :param pulumi.Input[dict] task_invocation_parameters: The parameters for task execution. This argument is conflict with `task_parameters` and `logging_info`.
        :param pulumi.Input[list] task_parameters: A structure containing information about parameters required by the particular `task_arn`. Use `parameter` configuration blocks under the `task_invocation_parameters` configuration block instead. Conflicts with `task_invocation_parameters`. Documented below.
        :param pulumi.Input[str] task_type: The type of task being registered. The only allowed value is `RUN_COMMAND`.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the task with.

        The **logging_info** object supports the following:

          * `s3_bucket_name` (`pulumi.Input[str]`)
          * `s3BucketPrefix` (`pulumi.Input[str]`)
          * `s3_region` (`pulumi.Input[str]`)

        The **targets** object supports the following:

          * `key` (`pulumi.Input[str]`)
          * `values` (`pulumi.Input[list]`)

        The **task_invocation_parameters** object supports the following:

          * `automationParameters` (`pulumi.Input[dict]`) - The parameters for an AUTOMATION task type. Documented below.
            * `document_version` (`pulumi.Input[str]`) - The version of an Automation document to use during task execution.
            * `parameters` (`pulumi.Input[list]`) - The parameters for the RUN_COMMAND task execution. Documented below.
              * `name` (`pulumi.Input[str]`) - The parameter name.
              * `values` (`pulumi.Input[list]`) - The array of strings.

          * `lambdaParameters` (`pulumi.Input[dict]`) - The parameters for a LAMBDA task type. Documented below.
            * `clientContext` (`pulumi.Input[str]`) - Pass client-specific information to the Lambda function that you are invoking.
            * `payload` (`pulumi.Input[str]`) - JSON to provide to your Lambda function as input.
            * `qualifier` (`pulumi.Input[str]`) - Specify a Lambda function version or alias name.

          * `runCommandParameters` (`pulumi.Input[dict]`) - The parameters for a RUN_COMMAND task type. Documented below.
            * `comment` (`pulumi.Input[str]`) - Information about the command(s) to execute.
            * `documentHash` (`pulumi.Input[str]`) - The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
            * `documentHashType` (`pulumi.Input[str]`) - SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
            * `notificationConfig` (`pulumi.Input[dict]`) - Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
              * `notificationArn` (`pulumi.Input[str]`) - An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
              * `notificationEvents` (`pulumi.Input[list]`) - The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
              * `notification_type` (`pulumi.Input[str]`) - When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`

            * `outputS3Bucket` (`pulumi.Input[str]`) - The name of the Amazon S3 bucket.
            * `outputS3KeyPrefix` (`pulumi.Input[str]`) - The Amazon S3 bucket subfolder.
            * `parameters` (`pulumi.Input[list]`) - The parameters for the RUN_COMMAND task execution. Documented below.
              * `name` (`pulumi.Input[str]`) - The parameter name.
              * `values` (`pulumi.Input[list]`) - The array of strings.

            * `service_role_arn` (`pulumi.Input[str]`) - The IAM service role to assume during task execution.
            * `timeoutSeconds` (`pulumi.Input[float]`) - If this time is reached and the command has not already started executing, it doesn't run.

          * `stepFunctionsParameters` (`pulumi.Input[dict]`) - The parameters for a STEP_FUNCTIONS task type. Documented below.
            * `input` (`pulumi.Input[str]`) - The inputs for the STEP_FUNCTION task.
            * `name` (`pulumi.Input[str]`) - The name of the STEP_FUNCTION task.

        The **task_parameters** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the maintenance window task.
          * `values` (`pulumi.Input[list]`)
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

            __props__['description'] = description
            if logging_info is not None:
                warnings.warn("use 'task_invocation_parameters' argument instead", DeprecationWarning)
                pulumi.log.warn("logging_info is deprecated: use 'task_invocation_parameters' argument instead")
            __props__['logging_info'] = logging_info
            if max_concurrency is None:
                raise TypeError("Missing required property 'max_concurrency'")
            __props__['max_concurrency'] = max_concurrency
            if max_errors is None:
                raise TypeError("Missing required property 'max_errors'")
            __props__['max_errors'] = max_errors
            __props__['name'] = name
            __props__['priority'] = priority
            if service_role_arn is None:
                raise TypeError("Missing required property 'service_role_arn'")
            __props__['service_role_arn'] = service_role_arn
            if targets is None:
                raise TypeError("Missing required property 'targets'")
            __props__['targets'] = targets
            if task_arn is None:
                raise TypeError("Missing required property 'task_arn'")
            __props__['task_arn'] = task_arn
            __props__['task_invocation_parameters'] = task_invocation_parameters
            if task_parameters is not None:
                warnings.warn("use 'task_invocation_parameters' argument instead", DeprecationWarning)
                pulumi.log.warn("task_parameters is deprecated: use 'task_invocation_parameters' argument instead")
            __props__['task_parameters'] = task_parameters
            if task_type is None:
                raise TypeError("Missing required property 'task_type'")
            __props__['task_type'] = task_type
            if window_id is None:
                raise TypeError("Missing required property 'window_id'")
            __props__['window_id'] = window_id
        super(MaintenanceWindowTask, __self__).__init__(
            'aws:ssm/maintenanceWindowTask:MaintenanceWindowTask',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, logging_info=None, max_concurrency=None, max_errors=None, name=None, priority=None, service_role_arn=None, targets=None, task_arn=None, task_invocation_parameters=None, task_parameters=None, task_type=None, window_id=None):
        """
        Get an existing MaintenanceWindowTask resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window task.
        :param pulumi.Input[dict] logging_info: A structure containing information about an Amazon S3 bucket to write instance-level logs to. Use `task_invocation_parameters` configuration block `run_command_parameters` configuration block `output_s3_*` arguments instead. Conflicts with `task_invocation_parameters`. Documented below.
        :param pulumi.Input[str] max_concurrency: The maximum number of targets this task can be run for in parallel.
        :param pulumi.Input[str] max_errors: The maximum number of errors allowed before this task stops being scheduled.
        :param pulumi.Input[str] name: The name of the maintenance window task.
        :param pulumi.Input[float] priority: The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
        :param pulumi.Input[str] service_role_arn: The role that should be assumed when executing the task.
        :param pulumi.Input[list] targets: The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
        :param pulumi.Input[str] task_arn: The ARN of the task to execute.
        :param pulumi.Input[dict] task_invocation_parameters: The parameters for task execution. This argument is conflict with `task_parameters` and `logging_info`.
        :param pulumi.Input[list] task_parameters: A structure containing information about parameters required by the particular `task_arn`. Use `parameter` configuration blocks under the `task_invocation_parameters` configuration block instead. Conflicts with `task_invocation_parameters`. Documented below.
        :param pulumi.Input[str] task_type: The type of task being registered. The only allowed value is `RUN_COMMAND`.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the task with.

        The **logging_info** object supports the following:

          * `s3_bucket_name` (`pulumi.Input[str]`)
          * `s3BucketPrefix` (`pulumi.Input[str]`)
          * `s3_region` (`pulumi.Input[str]`)

        The **targets** object supports the following:

          * `key` (`pulumi.Input[str]`)
          * `values` (`pulumi.Input[list]`)

        The **task_invocation_parameters** object supports the following:

          * `automationParameters` (`pulumi.Input[dict]`) - The parameters for an AUTOMATION task type. Documented below.
            * `document_version` (`pulumi.Input[str]`) - The version of an Automation document to use during task execution.
            * `parameters` (`pulumi.Input[list]`) - The parameters for the RUN_COMMAND task execution. Documented below.
              * `name` (`pulumi.Input[str]`) - The parameter name.
              * `values` (`pulumi.Input[list]`) - The array of strings.

          * `lambdaParameters` (`pulumi.Input[dict]`) - The parameters for a LAMBDA task type. Documented below.
            * `clientContext` (`pulumi.Input[str]`) - Pass client-specific information to the Lambda function that you are invoking.
            * `payload` (`pulumi.Input[str]`) - JSON to provide to your Lambda function as input.
            * `qualifier` (`pulumi.Input[str]`) - Specify a Lambda function version or alias name.

          * `runCommandParameters` (`pulumi.Input[dict]`) - The parameters for a RUN_COMMAND task type. Documented below.
            * `comment` (`pulumi.Input[str]`) - Information about the command(s) to execute.
            * `documentHash` (`pulumi.Input[str]`) - The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
            * `documentHashType` (`pulumi.Input[str]`) - SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
            * `notificationConfig` (`pulumi.Input[dict]`) - Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
              * `notificationArn` (`pulumi.Input[str]`) - An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
              * `notificationEvents` (`pulumi.Input[list]`) - The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
              * `notification_type` (`pulumi.Input[str]`) - When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`

            * `outputS3Bucket` (`pulumi.Input[str]`) - The name of the Amazon S3 bucket.
            * `outputS3KeyPrefix` (`pulumi.Input[str]`) - The Amazon S3 bucket subfolder.
            * `parameters` (`pulumi.Input[list]`) - The parameters for the RUN_COMMAND task execution. Documented below.
              * `name` (`pulumi.Input[str]`) - The parameter name.
              * `values` (`pulumi.Input[list]`) - The array of strings.

            * `service_role_arn` (`pulumi.Input[str]`) - The IAM service role to assume during task execution.
            * `timeoutSeconds` (`pulumi.Input[float]`) - If this time is reached and the command has not already started executing, it doesn't run.

          * `stepFunctionsParameters` (`pulumi.Input[dict]`) - The parameters for a STEP_FUNCTIONS task type. Documented below.
            * `input` (`pulumi.Input[str]`) - The inputs for the STEP_FUNCTION task.
            * `name` (`pulumi.Input[str]`) - The name of the STEP_FUNCTION task.

        The **task_parameters** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the maintenance window task.
          * `values` (`pulumi.Input[list]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["logging_info"] = logging_info
        __props__["max_concurrency"] = max_concurrency
        __props__["max_errors"] = max_errors
        __props__["name"] = name
        __props__["priority"] = priority
        __props__["service_role_arn"] = service_role_arn
        __props__["targets"] = targets
        __props__["task_arn"] = task_arn
        __props__["task_invocation_parameters"] = task_invocation_parameters
        __props__["task_parameters"] = task_parameters
        __props__["task_type"] = task_type
        __props__["window_id"] = window_id
        return MaintenanceWindowTask(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
