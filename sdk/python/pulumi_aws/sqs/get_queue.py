# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetQueueResult:
    """
    A collection of values returned by getQueue.
    """
    def __init__(__self__, arn=None, id=None, name=None, tags=None, url=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the queue.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags for the resource.
        """
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        __self__.url = url
        """
        The URL of the queue.
        """
class AwaitableGetQueueResult(GetQueueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQueueResult(
            arn=self.arn,
            id=self.id,
            name=self.name,
            tags=self.tags,
            url=self.url)

def get_queue(name=None,tags=None,opts=None):
    """
    Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
    By using this data source, you can reference SQS queues without having to hardcode
    the ARNs as input.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sqs_queue.html.markdown.


    :param str name: The name of the queue to match.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:sqs/getQueue:getQueue', __args__, opts=opts).value

    return AwaitableGetQueueResult(
        arn=__ret__.get('arn'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        tags=__ret__.get('tags'),
        url=__ret__.get('url'))
