# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetNetworkInterfacesResult:
    """
    A collection of values returned by getNetworkInterfaces.
    """
    def __init__(__self__, filters=None, id=None, ids=None, tags=None):
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
        """
        A list of all the network interface ids found. This data source will fail if none are found.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags


class AwaitableGetNetworkInterfacesResult(GetNetworkInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkInterfacesResult(
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            tags=self.tags)


def get_network_interfaces(filters=None, tags=None, opts=None):
    """
    ## Example Usage

    The following shows outputing all network interface ids in a region.

    ```python
    import pulumi
    import pulumi_aws as aws

    example_network_interfaces = aws.ec2.get_network_interfaces()
    pulumi.export("example", example_network_interfaces.ids)
    ```

    The following example retrieves a list of all network interface ids with a custom tag of `Name` set to a value of `test`.

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2.get_network_interfaces(tags={
        "Name": "test",
    })
    pulumi.export("example1", example.ids)
    ```

    The following example retrieves a network interface ids which associated
    with specific subnet.

    ```python
    import pulumi
    import pulumi_aws as aws

    example_network_interfaces = aws.ec2.get_network_interfaces(filters=[{
        "name": "subnet-id",
        "values": [aws_subnet["test"]["id"]],
    }])
    pulumi.export("example", example_network_interfaces.ids)
    ```


    :param list filters: Custom filter block as described below.
    :param dict tags: A map of tags, each pair of which must exactly match
           a pair on the desired network interfaces.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNetworkInterfaces.html).
      * `values` (`list`) - Set of values that are accepted for the given field.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getNetworkInterfaces:getNetworkInterfaces', __args__, opts=opts).value

    return AwaitableGetNetworkInterfacesResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        tags=__ret__.get('tags'))
