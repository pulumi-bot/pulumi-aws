# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetTaskDefinitionResult:
    """
    A collection of values returned by getTaskDefinition.
    """
    def __init__(__self__, family=None, network_mode=None, revision=None, status=None, task_definition=None, task_role_arn=None, id=None):
        if family and not isinstance(family, str):
            raise TypeError('Expected argument family to be a str')
        __self__.family = family
        """
        The family of this task definition
        """
        if network_mode and not isinstance(network_mode, str):
            raise TypeError('Expected argument network_mode to be a str')
        __self__.network_mode = network_mode
        """
        The Docker networking mode to use for the containers in this task.
        """
        if revision and not isinstance(revision, float):
            raise TypeError('Expected argument revision to be a float')
        __self__.revision = revision
        """
        The revision of this task definition
        """
        if status and not isinstance(status, str):
            raise TypeError('Expected argument status to be a str')
        __self__.status = status
        """
        The status of this task definition
        """
        if task_definition and not isinstance(task_definition, str):
            raise TypeError('Expected argument task_definition to be a str')
        __self__.task_definition = task_definition
        if task_role_arn and not isinstance(task_role_arn, str):
            raise TypeError('Expected argument task_role_arn to be a str')
        __self__.task_role_arn = task_role_arn
        """
        The ARN of the IAM role that containers in this task can assume
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_task_definition(task_definition=None,opts=None):
    """
    The ECS task definition data source allows access to details of
    a specific AWS ECS task definition.
    """
    __args__ = dict()

    __args__['taskDefinition'] = task_definition
    __ret__ = await pulumi.runtime.invoke('aws:ecs/getTaskDefinition:getTaskDefinition', __args__, opts=opts)

    return GetTaskDefinitionResult(
        family=__ret__.get('family'),
        network_mode=__ret__.get('networkMode'),
        revision=__ret__.get('revision'),
        status=__ret__.get('status'),
        task_definition=__ret__.get('taskDefinition'),
        task_role_arn=__ret__.get('taskRoleArn'),
        id=__ret__.get('id'))
