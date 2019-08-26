# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetAvailabilityZonesResult:
    """
    A collection of values returned by getAvailabilityZones.
    """
    def __init__(__self__, blacklisted_names=None, blacklisted_zone_ids=None, names=None, state=None, zone_ids=None, id=None):
        if blacklisted_names and not isinstance(blacklisted_names, list):
            raise TypeError("Expected argument 'blacklisted_names' to be a list")
        __self__.blacklisted_names = blacklisted_names
        if blacklisted_zone_ids and not isinstance(blacklisted_zone_ids, list):
            raise TypeError("Expected argument 'blacklisted_zone_ids' to be a list")
        __self__.blacklisted_zone_ids = blacklisted_zone_ids
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of the Availability Zone names available to the account.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        if zone_ids and not isinstance(zone_ids, list):
            raise TypeError("Expected argument 'zone_ids' to be a list")
        __self__.zone_ids = zone_ids
        """
        A list of the Availability Zone IDs available to the account.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetAvailabilityZonesResult(GetAvailabilityZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAvailabilityZonesResult(
            blacklisted_names=self.blacklisted_names,
            blacklisted_zone_ids=self.blacklisted_zone_ids,
            names=self.names,
            state=self.state,
            zone_ids=self.zone_ids,
            id=self.id)

def get_availability_zones(blacklisted_names=None,blacklisted_zone_ids=None,state=None,opts=None):
    """
    The Availability Zones data source allows access to the list of AWS
    Availability Zones which can be accessed by an AWS account within the region
    configured in the provider.
    
    This is different from the `.getAvailabilityZone` (singular) data source,
    which provides some details about a specific availability zone.
    
    :param list blacklisted_names: List of blacklisted Availability Zone names.
    :param list blacklisted_zone_ids: List of blacklisted Availability Zone IDs.
    :param str state: Allows to filter list of Availability Zones based on their
           current state. Can be either `"available"`, `"information"`, `"impaired"` or
           `"unavailable"`. By default the list includes a complete set of Availability Zones
           to which the underlying AWS account has access, regardless of their state.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/availability_zones.html.markdown.
    """
    __args__ = dict()

    __args__['blacklistedNames'] = blacklisted_names
    __args__['blacklistedZoneIds'] = blacklisted_zone_ids
    __args__['state'] = state
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getAvailabilityZones:getAvailabilityZones', __args__, opts=opts).value

    return AwaitableGetAvailabilityZonesResult(
        blacklisted_names=__ret__.get('blacklistedNames'),
        blacklisted_zone_ids=__ret__.get('blacklistedZoneIds'),
        names=__ret__.get('names'),
        state=__ret__.get('state'),
        zone_ids=__ret__.get('zoneIds'),
        id=__ret__.get('id'))
