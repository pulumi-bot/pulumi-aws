# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetAmiIdsResult:
    """
    A collection of values returned by getAmiIds.
    """
    def __init__(__self__, executable_users=None, filters=None, id=None, ids=None, name_regex=None, owners=None, sort_ascending=None):
        if executable_users and not isinstance(executable_users, list):
            raise TypeError("Expected argument 'executable_users' to be a list")
        __self__.executable_users = executable_users
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        __self__.owners = owners
        if sort_ascending and not isinstance(sort_ascending, bool):
            raise TypeError("Expected argument 'sort_ascending' to be a bool")
        __self__.sort_ascending = sort_ascending
class AwaitableGetAmiIdsResult(GetAmiIdsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAmiIdsResult(
            executable_users=self.executable_users,
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            owners=self.owners,
            sort_ascending=self.sort_ascending)

def get_ami_ids(executable_users=None,filters=None,name_regex=None,owners=None,sort_ascending=None,opts=None):
    """
    Use this data source to get a list of AMI IDs matching the specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    ubuntu = aws.get_ami_ids(filters=[{
            "name": "name",
            "values": ["ubuntu/images/ubuntu-*-*-amd64-server-*"],
        }],
        owners=["099720109477"])
    ```


    :param list executable_users: Limit search to users with *explicit* launch
           permission on  the image. Valid items are the numeric account ID or `self`.
    :param list filters: One or more name/value pairs to filter off of. There
           are several valid keys, for a full reference, check out
           [describe-images in the AWS CLI reference][1].
    :param str name_regex: A regex string to apply to the AMI list returned
           by AWS. This allows more advanced filtering not supported from the AWS API.
           This filtering is done locally on what AWS returns, and could have a performance
           impact if the result is large. It is recommended to combine this with other
           options to narrow down the list AWS returns.
    :param list owners: List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
    :param bool sort_ascending: Used to sort AMIs by creation time.

    The **filters** object supports the following:

      * `name` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()


    __args__['executableUsers'] = executable_users
    __args__['filters'] = filters
    __args__['nameRegex'] = name_regex
    __args__['owners'] = owners
    __args__['sortAscending'] = sort_ascending
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getAmiIds:getAmiIds', __args__, opts=opts).value

    return AwaitableGetAmiIdsResult(
        executable_users=__ret__.get('executableUsers'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        name_regex=__ret__.get('nameRegex'),
        owners=__ret__.get('owners'),
        sort_ascending=__ret__.get('sortAscending'))
