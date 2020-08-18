# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['MaintenanceWindowTask']


class MaintenanceWindowTask(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 max_concurrency: Optional[pulumi.Input[str]] = None,
                 max_errors: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 service_role_arn: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTargetArgs']]]]] = None,
                 task_arn: Optional[pulumi.Input[str]] = None,
                 task_invocation_parameters: Optional[pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTaskInvocationParametersArgs']]] = None,
                 task_type: Optional[pulumi.Input[str]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an SSM Maintenance Window Task resource

        ## Example Usage
        ### Automation Tasks

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssm.MaintenanceWindowTask("example",
            max_concurrency="2",
            max_errors="1",
            priority=1,
            service_role_arn=aws_iam_role["example"]["arn"],
            task_arn="AWS-RestartEC2Instance",
            task_type="AUTOMATION",
            window_id=aws_ssm_maintenance_window["example"]["id"],
            targets=[{
                "key": "InstanceIds",
                "values": [aws_instance["example"]["id"]],
            }],
            task_invocation_parameters={
                "automationParameters": {
                    "document_version": "$LATEST",
                    "parameters": [{
                        "name": "InstanceId",
                        "values": [aws_instance["example"]["id"]],
                    }],
                },
            })
        ```
        ### Run Command Tasks

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssm.MaintenanceWindowTask("example",
            max_concurrency="2",
            max_errors="1",
            priority=1,
            service_role_arn=aws_iam_role["example"]["arn"],
            task_arn="AWS-RunShellScript",
            task_type="RUN_COMMAND",
            window_id=aws_ssm_maintenance_window["example"]["id"],
            targets=[{
                "key": "InstanceIds",
                "values": [aws_instance["example"]["id"]],
            }],
            task_invocation_parameters={
                "runCommandParameters": {
                    "outputS3Bucket": aws_s3_bucket["example"]["bucket"],
                    "outputS3KeyPrefix": "output",
                    "service_role_arn": aws_iam_role["example"]["arn"],
                    "timeoutSeconds": 600,
                    "notificationConfig": {
                        "notificationArn": aws_sns_topic["example"]["arn"],
                        "notificationEvents": ["All"],
                        "notification_type": "Command",
                    },
                    "parameters": [{
                        "name": "commands",
                        "values": ["date"],
                    }],
                },
            })
        ```
        ### Step Function Tasks

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ssm.MaintenanceWindowTask("example",
            max_concurrency="2",
            max_errors="1",
            priority=1,
            service_role_arn=aws_iam_role["example"]["arn"],
            task_arn=aws_sfn_activity["example"]["id"],
            task_type="STEP_FUNCTIONS",
            window_id=aws_ssm_maintenance_window["example"]["id"],
            targets=[{
                "key": "InstanceIds",
                "values": [aws_instance["example"]["id"]],
            }],
            task_invocation_parameters={
                "stepFunctionsParameters": {
                    "input": "{\"key1\":\"value1\"}",
                    "name": "example",
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window task.
        :param pulumi.Input[str] max_concurrency: The maximum number of targets this task can be run for in parallel.
        :param pulumi.Input[str] max_errors: The maximum number of errors allowed before this task stops being scheduled.
        :param pulumi.Input[str] name: The name of the maintenance window task.
        :param pulumi.Input[float] priority: The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
        :param pulumi.Input[str] service_role_arn: The role that should be assumed when executing the task.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTargetArgs']]]] targets: The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
        :param pulumi.Input[str] task_arn: The ARN of the task to execute.
        :param pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTaskInvocationParametersArgs']] task_invocation_parameters: Configuration block with parameters for task execution.
        :param pulumi.Input[str] task_type: The type of task being registered. The only allowed value is `RUN_COMMAND`.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the task with.
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
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            max_concurrency: Optional[pulumi.Input[str]] = None,
            max_errors: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            service_role_arn: Optional[pulumi.Input[str]] = None,
            targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTargetArgs']]]]] = None,
            task_arn: Optional[pulumi.Input[str]] = None,
            task_invocation_parameters: Optional[pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTaskInvocationParametersArgs']]] = None,
            task_type: Optional[pulumi.Input[str]] = None,
            window_id: Optional[pulumi.Input[str]] = None) -> 'MaintenanceWindowTask':
        """
        Get an existing MaintenanceWindowTask resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window task.
        :param pulumi.Input[str] max_concurrency: The maximum number of targets this task can be run for in parallel.
        :param pulumi.Input[str] max_errors: The maximum number of errors allowed before this task stops being scheduled.
        :param pulumi.Input[str] name: The name of the maintenance window task.
        :param pulumi.Input[float] priority: The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
        :param pulumi.Input[str] service_role_arn: The role that should be assumed when executing the task.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTargetArgs']]]] targets: The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
        :param pulumi.Input[str] task_arn: The ARN of the task to execute.
        :param pulumi.Input[pulumi.InputType['MaintenanceWindowTaskTaskInvocationParametersArgs']] task_invocation_parameters: Configuration block with parameters for task execution.
        :param pulumi.Input[str] task_type: The type of task being registered. The only allowed value is `RUN_COMMAND`.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the task with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["max_concurrency"] = max_concurrency
        __props__["max_errors"] = max_errors
        __props__["name"] = name
        __props__["priority"] = priority
        __props__["service_role_arn"] = service_role_arn
        __props__["targets"] = targets
        __props__["task_arn"] = task_arn
        __props__["task_invocation_parameters"] = task_invocation_parameters
        __props__["task_type"] = task_type
        __props__["window_id"] = window_id
        return MaintenanceWindowTask(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the maintenance window task.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="maxConcurrency")
    def max_concurrency(self) -> str:
        """
        The maximum number of targets this task can be run for in parallel.
        """
        return pulumi.get(self, "max_concurrency")

    @property
    @pulumi.getter(name="maxErrors")
    def max_errors(self) -> str:
        """
        The maximum number of errors allowed before this task stops being scheduled.
        """
        return pulumi.get(self, "max_errors")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the maintenance window task.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> Optional[float]:
        """
        The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> str:
        """
        The role that should be assumed when executing the task.
        """
        return pulumi.get(self, "service_role_arn")

    @property
    @pulumi.getter
    def targets(self) -> List['outputs.MaintenanceWindowTaskTarget']:
        """
        The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="taskArn")
    def task_arn(self) -> str:
        """
        The ARN of the task to execute.
        """
        return pulumi.get(self, "task_arn")

    @property
    @pulumi.getter(name="taskInvocationParameters")
    def task_invocation_parameters(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParameters']:
        """
        Configuration block with parameters for task execution.
        """
        return pulumi.get(self, "task_invocation_parameters")

    @property
    @pulumi.getter(name="taskType")
    def task_type(self) -> str:
        """
        The type of task being registered. The only allowed value is `RUN_COMMAND`.
        """
        return pulumi.get(self, "task_type")

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> str:
        """
        The Id of the maintenance window to register the task with.
        """
        return pulumi.get(self, "window_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

