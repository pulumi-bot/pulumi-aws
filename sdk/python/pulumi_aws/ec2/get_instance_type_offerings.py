# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetInstanceTypeOfferingsResult:
    """
    A collection of values returned by getInstanceTypeOfferings.
    """
    def __init__(__self__, filters=None, id=None, instance_types=None, location_type=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instance_types and not isinstance(instance_types, list):
            raise TypeError("Expected argument 'instance_types' to be a list")
        __self__.instance_types = instance_types
        """
        Set of EC2 Instance Types.
        """
        if location_type and not isinstance(location_type, str):
            raise TypeError("Expected argument 'location_type' to be a str")
        __self__.location_type = location_type


class AwaitableGetInstanceTypeOfferingsResult(GetInstanceTypeOfferingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceTypeOfferingsResult(
            filters=self.filters,
            id=self.id,
            instance_types=self.instance_types,
            location_type=self.location_type)


def get_instance_type_offerings(filters=None, location_type=None, opts=None):
    """
    Information about EC2 Instance Type Offerings.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2.get_instance_type_offerings(filters=[
            {
                "name": "instance-type",
                "values": [
                    "t2.micro",
                    "t3.micro",
                ],
            },
            {
                "name": "location",
                "values": ["usw2-az4"],
            },
        ],
        location_type="availability-zone-id")
    ```


    :param list filters: One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
    :param str location_type: Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.

    The **filters** object supports the following:

      * `name` (`str`) - Name of the filter. The `location` filter depends on the top-level `location_type` argument and if not specified, defaults to the current region.
      * `values` (`list`) - List of one or more values for the filter.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['locationType'] = location_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getInstanceTypeOfferings:getInstanceTypeOfferings', __args__, opts=opts).value

    return AwaitableGetInstanceTypeOfferingsResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        instance_types=__ret__.get('instanceTypes'),
        location_type=__ret__.get('locationType'))
