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

__all__ = [
    'GetSecurityGroupResult',
    'AwaitableGetSecurityGroupResult',
    'get_security_group',
]



@pulumi.output_type
class GetSecurityGroupResult:
    """
    A collection of values returned by getSecurityGroup.
    """
    def __init__(__self__, arn=None, description=None, filters=None, id=None, name=None, tags=None, vpc_id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The computed ARN of the security group.
        """
        ...

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the security group.
        """
        ...

    @property
    @pulumi.getter
    def filters(self) -> Optional[List['outputs.GetSecurityGroupFilterResult']]:
        ...

    @property
    @pulumi.getter
    def id(self) -> str:
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        ...

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        ...



class AwaitableGetSecurityGroupResult(GetSecurityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupResult(
            arn=self.arn,
            description=self.description,
            filters=self.filters,
            id=self.id,
            name=self.name,
            tags=self.tags,
            vpc_id=self.vpc_id)


def get_security_group(filters: Optional[List[pulumi.InputType['GetSecurityGroupFilterArgs']]] = None,
                       id: Optional[str] = None,
                       name: Optional[str] = None,
                       tags: Optional[Mapping[str, str]] = None,
                       vpc_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupResult:
    """
    `ec2.SecurityGroup` provides details about a specific Security Group.

    This resource can prove useful when a module accepts a Security Group id as
    an input variable and needs to, for example, determine the id of the
    VPC that the security group belongs to.

    ## Example Usage

    The following example shows how one might accept a Security Group id as a variable
    and use this data source to obtain the data necessary to create a subnet.

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    security_group_id = config.require_object("securityGroupId")
    selected = aws.ec2.get_security_group(id=security_group_id)
    subnet = aws.ec2.Subnet("subnet",
        vpc_id=selected.vpc_id,
        cidr_block="10.0.1.0/24")
    ```


    :param List[pulumi.InputType['GetSecurityGroupFilterArgs']] filters: Custom filter block as described below.
    :param str id: The id of the specific security group to retrieve.
    :param str name: The name of the field to filter by, as defined by
           [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html).
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match
           a pair on the desired security group.
    :param str vpc_id: The id of the VPC that the desired security group belongs to.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['name'] = name
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getSecurityGroup:getSecurityGroup', __args__, opts=opts, typ=GetSecurityGroupResult).value

    return AwaitableGetSecurityGroupResult(
        arn=__ret__.arn,
        description=__ret__.description,
        filters=__ret__.filters,
        id=__ret__.id,
        name=__ret__.name,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id)
