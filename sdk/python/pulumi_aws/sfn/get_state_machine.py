# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetStateMachineResult:
    """
    A collection of values returned by getStateMachine.
    """
    def __init__(__self__, arn=None, creation_date=None, definition=None, id=None, name=None, role_arn=None, status=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Set to the arn of the state function.
        """
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        __self__.creation_date = creation_date
        """
        The date the state machine was created.
        """
        if definition and not isinstance(definition, str):
            raise TypeError("Expected argument 'definition' to be a str")
        __self__.definition = definition
        """
        Set to the state machine definition.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        __self__.role_arn = role_arn
        """
        Set to the role_arn used by the state function.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        Set to the current status of the state machine.
        """
class AwaitableGetStateMachineResult(GetStateMachineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStateMachineResult(
            arn=self.arn,
            creation_date=self.creation_date,
            definition=self.definition,
            id=self.id,
            name=self.name,
            role_arn=self.role_arn,
            status=self.status)

def get_state_machine(name=None,opts=None):
    """
    Use this data source to get the ARN of a State Machine in AWS Step
    Function (SFN). By using this data source, you can reference a
    state machine without having to hard code the ARNs as input.




    :param str name: The friendly name of the state machine to match.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:sfn/getStateMachine:getStateMachine', __args__, opts=opts).value

    return AwaitableGetStateMachineResult(
        arn=__ret__.get('arn'),
        creation_date=__ret__.get('creationDate'),
        definition=__ret__.get('definition'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        role_arn=__ret__.get('roleArn'),
        status=__ret__.get('status'))
