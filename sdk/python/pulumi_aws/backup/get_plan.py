# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetPlanResult:
    """
    A collection of values returned by getPlan.
    """
    def __init__(__self__, arn=None, id=None, name=None, plan_id=None, tags=None, version=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN of the backup plan.
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
        """
        The display name of a backup plan.
        """
        if plan_id and not isinstance(plan_id, str):
            raise TypeError("Expected argument 'plan_id' to be a str")
        __self__.plan_id = plan_id
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Metadata that you can assign to help organize the plans you create.
        """
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        __self__.version = version
        """
        Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
        """


class AwaitableGetPlanResult(GetPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlanResult(
            arn=self.arn,
            id=self.id,
            name=self.name,
            plan_id=self.plan_id,
            tags=self.tags,
            version=self.version)


def get_plan(plan_id=None,tags=None,opts=None):
    """
    Use this data source to get information on an existing backup plan.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.backup.get_plan(plan_id="tf_example_backup_plan_id")
    ```


    :param str plan_id: The backup plan ID.
    :param dict tags: Metadata that you can assign to help organize the plans you create.
    """
    __args__ = dict()

    __args__['planId'] = plan_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:backup/getPlan:getPlan', __args__, opts=opts).value

    return AwaitableGetPlanResult(
        arn=__ret__.get('arn'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        plan_id=__ret__.get('planId'),
        tags=__ret__.get('tags'),
        version=__ret__.get('version'))
