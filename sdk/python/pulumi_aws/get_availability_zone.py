# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetAvailabilityZoneResult:
    """
    A collection of values returned by getAvailabilityZone.
    """
    def __init__(__self__, name=None, name_suffix=None, region=None, state=None, zone_id=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the selected availability zone.
        """
        if name_suffix and not isinstance(name_suffix, str):
            raise TypeError("Expected argument 'name_suffix' to be a str")
        __self__.name_suffix = name_suffix
        """
        The part of the AZ name that appears after the region name,
        uniquely identifying the AZ within its region.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        The region where the selected availability zone resides.
        This is always the region selected on the provider, since this data source
        searches only within that region.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        The current state of the AZ.
        """
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        __self__.zone_id = zone_id
        """
        (Optional) The zone ID of the selected availability zone.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_availability_zone(name=None,state=None,zone_id=None,opts=None):
    """
    `aws_availability_zone` provides details about a specific availability zone (AZ)
    in the current region.
    
    This can be used both to validate an availability zone given in a variable
    and to split the AZ name into its component parts of an AWS region and an
    AZ identifier letter. The latter may be useful e.g. for implementing a
    consistent subnet numbering scheme across several regions by mapping both
    the region and the subnet letter to network numbers.
    
    This is different from the `aws_availability_zones` (plural) data source,
    which provides a list of the available zones.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/availability_zone.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['state'] = state
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('aws:index/getAvailabilityZone:getAvailabilityZone', __args__, opts=opts)

    return GetAvailabilityZoneResult(
        name=__ret__.get('name'),
        name_suffix=__ret__.get('nameSuffix'),
        region=__ret__.get('region'),
        state=__ret__.get('state'),
        zone_id=__ret__.get('zoneId'),
        id=__ret__.get('id'))
