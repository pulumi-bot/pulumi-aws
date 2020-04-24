# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetRegionsResult:
    """
    A collection of values returned by getRegions.
    """
    def __init__(__self__, all_regions=None, filters=None, id=None, names=None):
        if all_regions and not isinstance(all_regions, bool):
            raise TypeError("Expected argument 'all_regions' to be a bool")
        __self__.all_regions = all_regions
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        Names of regions that meets the criteria.
        """
class AwaitableGetRegionsResult(GetRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionsResult(
            all_regions=self.all_regions,
            filters=self.filters,
            id=self.id,
            names=self.names)

def get_regions(all_regions=None,filters=None,opts=None):
    """
    Provides information about AWS Regions. Can be used to filter regions i.e. by Opt-In status or only regions enabled for current account. To get details like endpoint and description of each region the data source can be combined with the [`.getRegion` data source](https://www.terraform.io/docs/providers/aws/d/region.html).




    :param bool all_regions: If true the source will query all regions regardless of availability.
    :param list filters: Configuration block(s) to use as filters. Detailed below.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the filter field. Valid values can be found in the [describe-regions AWS CLI Reference][1].
      * `values` (`list`) - Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
    """
    __args__ = dict()


    __args__['allRegions'] = all_regions
    __args__['filters'] = filters
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getRegions:getRegions', __args__, opts=opts).value

    return AwaitableGetRegionsResult(
        all_regions=__ret__.get('allRegions'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        names=__ret__.get('names'))
