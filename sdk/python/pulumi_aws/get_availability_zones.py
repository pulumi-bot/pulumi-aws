# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities

class GetAvailabilityZonesResult(object):
    """
    A collection of values returned by getAvailabilityZones.
    """
    def __init__(__self__, names=None, id=None):
        if names and not isinstance(names, list):
            raise TypeError('Expected argument names to be a list')
        __self__.names = names
        """
        A list of the Availability Zone names available to the account.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_availability_zones(state=None):
    """
    The Availability Zones data source allows access to the list of AWS
    Availability Zones which can be accessed by an AWS account within the region
    configured in the provider.
    
    This is different from the `aws_availability_zone` (singular) data source,
    which provides some details about a specific availability zone.
    """
    __args__ = dict()

    __args__['state'] = state
    __ret__ = pulumi.runtime.invoke('aws:index/getAvailabilityZones:getAvailabilityZones', __args__)

    return GetAvailabilityZonesResult(
        names=__ret__.get('names'),
        id=__ret__.get('id'))
