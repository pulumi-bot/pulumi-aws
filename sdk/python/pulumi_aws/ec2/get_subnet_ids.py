# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetSubnetIdsResult',
    'AwaitableGetSubnetIdsResult',
    'get_subnet_ids',
]

@pulumi.output_type
class GetSubnetIdsResult:
    """
    A collection of values returned by getSubnetIds.
    """
    def __init__(__self__, filters=None, id=None, ids=None, tags=None, vpc_id=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetSubnetIdsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        """
        A set of all the subnet ids found. This data source will fail if none are found.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")


class AwaitableGetSubnetIdsResult(GetSubnetIdsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetIdsResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_subnet_ids(filters: Optional[Sequence[pulumi.InputType['GetSubnetIdsFilterArgs']]] = None,
                   tags: Optional[Mapping[str, str]] = None,
                   vpc_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetIdsResult:
    """
    `ec2.getSubnetIds` provides a set of ids for a vpc_id

    This resource can be useful for getting back a set of subnet ids for a vpc.

    ## Example Usage

    The following shows outputing all cidr blocks for every subnet id in a vpc.

    ```python
    import pulumi
    import pulumi_aws as aws

    example_subnet_ids = aws.ec2.get_subnet_ids(vpc_id=var["vpc_id"])
    example_subnet = [aws.ec2.get_subnet(id=__value) for __key, __value in example_subnet_ids.ids]
    pulumi.export("subnetCidrBlocks", [s.cidr_block for s in example_subnet])
    ```

    The following example retrieves a set of all subnets in a VPC with a custom
    tag of `Tier` set to a value of "Private" so that the `ec2.Instance` resource
    can loop through the subnets, putting instances across availability zones.

    ```python
    import pulumi
    import pulumi_aws as aws

    private = aws.ec2.get_subnet_ids(vpc_id=var["vpc_id"],
        tags={
            "Tier": "Private",
        })
    app = []
    for range in [{"value": i} for i in range(0, data.aws_subnet_ids.example.ids)]:
        app.append(aws.ec2.Instance(f"app-{range['value']}",
            ami=var["ami"],
            instance_type="t2.micro",
            subnet_id=range["value"]))
    ```


    :param Sequence[pulumi.InputType['GetSubnetIdsFilterArgs']] filters: Custom filter block as described below.
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match
           a pair on the desired subnets.
    :param str vpc_id: The VPC ID that you want to filter from.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getSubnetIds:getSubnetIds', __args__, opts=opts, typ=GetSubnetIdsResult).value

    return AwaitableGetSubnetIdsResult(
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_subnet_ids)
def get_subnet_ids_apply(filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GetSubnetIdsFilterArgs']]]]] = None,
                         tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                         vpc_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetIdsResult]:
    ...
