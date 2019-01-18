# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MaintenanceWindowTask(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the maintenance window task.
    """
    logging_info: pulumi.Output[dict]
    """
    A structure containing information about an Amazon S3 bucket to write instance-level logs to. Documented below.
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
    priority: pulumi.Output[int]
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
    """
    task_arn: pulumi.Output[str]
    """
    The ARN of the task to execute.
    """
    task_parameters: pulumi.Output[list]
    """
    A structure containing information about parameters required by the particular `task_arn`. Documented below.
    """
    task_type: pulumi.Output[str]
    """
    The type of task being registered. The only allowed value is `RUN_COMMAND`.
    """
    window_id: pulumi.Output[str]
    """
    The Id of the maintenance window to register the task with.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, logging_info=None, max_concurrency=None, max_errors=None, name=None, priority=None, service_role_arn=None, targets=None, task_arn=None, task_parameters=None, task_type=None, window_id=None):
        """
        Provides an SSM Maintenance Window Task resource
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] description: The description of the maintenance window task.
        :param pulumi.Input[dict] logging_info: A structure containing information about an Amazon S3 bucket to write instance-level logs to. Documented below.
        :param pulumi.Input[str] max_concurrency: The maximum number of targets this task can be run for in parallel.
        :param pulumi.Input[str] max_errors: The maximum number of errors allowed before this task stops being scheduled.
        :param pulumi.Input[str] name
        :param pulumi.Input[int] priority: The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
        :param pulumi.Input[str] service_role_arn: The role that should be assumed when executing the task.
        :param pulumi.Input[list] targets: The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
        :param pulumi.Input[str] task_arn: The ARN of the task to execute.
        :param pulumi.Input[list] task_parameters: A structure containing information about parameters required by the particular `task_arn`. Documented below.
        :param pulumi.Input[str] task_type: The type of task being registered. The only allowed value is `RUN_COMMAND`.
        :param pulumi.Input[str] window_id: The Id of the maintenance window to register the task with.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['logging_info'] = logging_info

        if not max_concurrency:
            raise TypeError('Missing required property max_concurrency')
        __props__['max_concurrency'] = max_concurrency

        if not max_errors:
            raise TypeError('Missing required property max_errors')
        __props__['max_errors'] = max_errors

        __props__['name'] = name

        __props__['priority'] = priority

        if not service_role_arn:
            raise TypeError('Missing required property service_role_arn')
        __props__['service_role_arn'] = service_role_arn

        if not targets:
            raise TypeError('Missing required property targets')
        __props__['targets'] = targets

        if not task_arn:
            raise TypeError('Missing required property task_arn')
        __props__['task_arn'] = task_arn

        __props__['task_parameters'] = task_parameters

        if not task_type:
            raise TypeError('Missing required property task_type')
        __props__['task_type'] = task_type

        if not window_id:
            raise TypeError('Missing required property window_id')
        __props__['window_id'] = window_id

        super(MaintenanceWindowTask, __self__).__init__(
            'aws:ssm/maintenanceWindowTask:MaintenanceWindowTask',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

