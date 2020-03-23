# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLogGroupResult:
    """
    A collection of values returned by getLogGroup.
    """
    def __init__(__self__, arn=None, creation_time=None, id=None, kms_key_id=None, name=None, retention_in_days=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN of the Cloudwatch log group
        """
        if creation_time and not isinstance(creation_time, float):
            raise TypeError("Expected argument 'creation_time' to be a float")
        __self__.creation_time = creation_time
        """
        The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        __self__.kms_key_id = kms_key_id
        """
        The ARN of the KMS Key to use when encrypting log data.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if retention_in_days and not isinstance(retention_in_days, float):
            raise TypeError("Expected argument 'retention_in_days' to be a float")
        __self__.retention_in_days = retention_in_days
        """
        The number of days log events retained in the specified log group.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
class AwaitableGetLogGroupResult(GetLogGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLogGroupResult(
            arn=self.arn,
            creation_time=self.creation_time,
            id=self.id,
            kms_key_id=self.kms_key_id,
            name=self.name,
            retention_in_days=self.retention_in_days,
            tags=self.tags)

def get_log_group(name=None,tags=None,opts=None):
    """
    Use this data source to get information about an AWS Cloudwatch Log Group

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudwatch_log_group.html.markdown.


    :param str name: The name of the Cloudwatch log group
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:cloudwatch/getLogGroup:getLogGroup', __args__, opts=opts).value

    return AwaitableGetLogGroupResult(
        arn=__ret__.get('arn'),
        creation_time=__ret__.get('creationTime'),
        id=__ret__.get('id'),
        kms_key_id=__ret__.get('kmsKeyId'),
        name=__ret__.get('name'),
        retention_in_days=__ret__.get('retentionInDays'),
        tags=__ret__.get('tags'))
