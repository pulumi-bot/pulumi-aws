# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetTopicResult:
    """
    A collection of values returned by getTopic.
    """
    def __init__(__self__, arn=None, name=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Set to the ARN of the found topic, suitable for referencing in other resources that support SNS topics.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_topic(name=None,opts=None):
    """
    Use this data source to get the ARN of a topic in AWS Simple Notification
    Service (SNS). By using this data source, you can reference SNS topics
    without having to hard code the ARNs as input.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sns_topic.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:sns/getTopic:getTopic', __args__, opts=opts).value

    return GetTopicResult(
        arn=__ret__.get('arn'),
        name=__ret__.get('name'),
        id=__ret__.get('id'))
