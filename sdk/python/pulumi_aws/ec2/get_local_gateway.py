# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLocalGatewayResult:
    """
    A collection of values returned by getLocalGateway.
    """
    def __init__(__self__, filters=None, id=None, outpost_arn=None, owner_id=None, state=None, tags=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        __self__.outpost_arn = outpost_arn
        """
        Amazon Resource Name (ARN) of Outpost
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        AWS account identifier that owns the Local Gateway.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        State of the local gateway.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
class AwaitableGetLocalGatewayResult(GetLocalGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocalGatewayResult(
            filters=self.filters,
            id=self.id,
            outpost_arn=self.outpost_arn,
            owner_id=self.owner_id,
            state=self.state,
            tags=self.tags)

def get_local_gateway(filters=None,id=None,state=None,tags=None,opts=None):
    """
    Provides details about an EC2 Local Gateway.

    ## Example Usage

    The following example shows how one might accept a local gateway id as a variable.

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    local_gateway_id = config.require_object("localGatewayId")
    selected = aws.ec2.get_local_gateway(id=local_gateway_id)
    ```


    :param list filters: Custom filter block as described below.
    :param str id: The id of the specific Local Gateway to retrieve.
    :param str state: The current state of the desired Local Gateway.
           Can be either `"pending"` or `"available"`.
    :param dict tags: A mapping of tags, each pair of which must exactly match
           a pair on the desired Local Gateway.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGateways.html).
      * `values` (`list`) - Set of values that are accepted for the given field.
        A Local Gateway will be selected if any one of the given values matches.
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLocalGateway:getLocalGateway', __args__, opts=opts).value

    return AwaitableGetLocalGatewayResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        outpost_arn=__ret__.get('outpostArn'),
        owner_id=__ret__.get('ownerId'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'))
