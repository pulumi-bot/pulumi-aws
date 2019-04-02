# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, caller_reference=None, comment=None, name=None, name_servers=None, private_zone=None, resource_record_set_count=None, tags=None, vpc_id=None, zone_id=None, id=None):
        if caller_reference and not isinstance(caller_reference, str):
            raise TypeError('Expected argument caller_reference to be a str')
        __self__.caller_reference = caller_reference
        """
        Caller Reference of the Hosted Zone.
        """
        if comment and not isinstance(comment, str):
            raise TypeError('Expected argument comment to be a str')
        __self__.comment = comment
        """
        The comment field of the Hosted Zone.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        if name_servers and not isinstance(name_servers, list):
            raise TypeError('Expected argument name_servers to be a list')
        __self__.name_servers = name_servers
        """
        The list of DNS name servers for the Hosted Zone.
        """
        if private_zone and not isinstance(private_zone, bool):
            raise TypeError('Expected argument private_zone to be a bool')
        __self__.private_zone = private_zone
        if resource_record_set_count and not isinstance(resource_record_set_count, float):
            raise TypeError('Expected argument resource_record_set_count to be a float')
        __self__.resource_record_set_count = resource_record_set_count
        """
        the number of Record Set in the Hosted Zone
        """
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError('Expected argument vpc_id to be a str')
        __self__.vpc_id = vpc_id
        if zone_id and not isinstance(zone_id, str):
            raise TypeError('Expected argument zone_id to be a str')
        __self__.zone_id = zone_id
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_zone(caller_reference=None,comment=None,name=None,private_zone=None,resource_record_set_count=None,tags=None,vpc_id=None,zone_id=None,opts=None):
    """
    `aws_route53_zone` provides details about a specific Route 53 Hosted Zone.
    
    This data source allows to find a Hosted Zone ID given Hosted Zone name and certain search criteria.
    """
    __args__ = dict()

    __args__['callerReference'] = caller_reference
    __args__['comment'] = comment
    __args__['name'] = name
    __args__['privateZone'] = private_zone
    __args__['resourceRecordSetCount'] = resource_record_set_count
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __args__['zoneId'] = zone_id
    __ret__ = await pulumi.runtime.invoke('aws:route53/getZone:getZone', __args__, opts=opts)

    return GetZoneResult(
        caller_reference=__ret__.get('callerReference'),
        comment=__ret__.get('comment'),
        name=__ret__.get('name'),
        name_servers=__ret__.get('nameServers'),
        private_zone=__ret__.get('privateZone'),
        resource_record_set_count=__ret__.get('resourceRecordSetCount'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'),
        zone_id=__ret__.get('zoneId'),
        id=__ret__.get('id'))
