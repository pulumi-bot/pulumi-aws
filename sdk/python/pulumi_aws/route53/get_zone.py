# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetZoneResult',
    'AwaitableGetZoneResult',
    'get_zone',
]

@pulumi.output_type
class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, caller_reference=None, comment=None, id=None, linked_service_description=None, linked_service_principal=None, name=None, name_servers=None, private_zone=None, resource_record_set_count=None, tags=None, vpc_id=None, zone_id=None):
        if caller_reference and not isinstance(caller_reference, str):
            raise TypeError("Expected argument 'caller_reference' to be a str")
        pulumi.set(__self__, "caller_reference", caller_reference)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if linked_service_description and not isinstance(linked_service_description, str):
            raise TypeError("Expected argument 'linked_service_description' to be a str")
        pulumi.set(__self__, "linked_service_description", linked_service_description)
        if linked_service_principal and not isinstance(linked_service_principal, str):
            raise TypeError("Expected argument 'linked_service_principal' to be a str")
        pulumi.set(__self__, "linked_service_principal", linked_service_principal)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_servers and not isinstance(name_servers, list):
            raise TypeError("Expected argument 'name_servers' to be a list")
        pulumi.set(__self__, "name_servers", name_servers)
        if private_zone and not isinstance(private_zone, bool):
            raise TypeError("Expected argument 'private_zone' to be a bool")
        pulumi.set(__self__, "private_zone", private_zone)
        if resource_record_set_count and not isinstance(resource_record_set_count, int):
            raise TypeError("Expected argument 'resource_record_set_count' to be a int")
        pulumi.set(__self__, "resource_record_set_count", resource_record_set_count)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="callerReference")
    def caller_reference(self) -> str:
        """
        Caller Reference of the Hosted Zone.
        """
        return pulumi.get(self, "caller_reference")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        The comment field of the Hosted Zone.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="linkedServiceDescription")
    def linked_service_description(self) -> str:
        """
        The description provided by the service that created the Hosted Zone (e.g. `arn:aws:servicediscovery:us-east-1:1234567890:namespace/ns-xxxxxxxxxxxxxxxx`).
        """
        return pulumi.get(self, "linked_service_description")

    @property
    @pulumi.getter(name="linkedServicePrincipal")
    def linked_service_principal(self) -> str:
        """
        The service that created the Hosted Zone (e.g. `servicediscovery.amazonaws.com`).
        """
        return pulumi.get(self, "linked_service_principal")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> Sequence[str]:
        """
        The list of DNS name servers for the Hosted Zone.
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter(name="privateZone")
    def private_zone(self) -> Optional[bool]:
        return pulumi.get(self, "private_zone")

    @property
    @pulumi.getter(name="resourceRecordSetCount")
    def resource_record_set_count(self) -> int:
        """
        The number of Record Set in the Hosted Zone.
        """
        return pulumi.get(self, "resource_record_set_count")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        return pulumi.get(self, "zone_id")


class AwaitableGetZoneResult(GetZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneResult(
            caller_reference=self.caller_reference,
            comment=self.comment,
            id=self.id,
            linked_service_description=self.linked_service_description,
            linked_service_principal=self.linked_service_principal,
            name=self.name,
            name_servers=self.name_servers,
            private_zone=self.private_zone,
            resource_record_set_count=self.resource_record_set_count,
            tags=self.tags,
            vpc_id=self.vpc_id,
            zone_id=self.zone_id)


def get_zone(name: Optional[str] = None,
             private_zone: Optional[bool] = None,
             resource_record_set_count: Optional[int] = None,
             tags: Optional[Mapping[str, str]] = None,
             vpc_id: Optional[str] = None,
             zone_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneResult:
    """
    `route53.Zone` provides details about a specific Route 53 Hosted Zone.

    This data source allows to find a Hosted Zone ID given Hosted Zone name and certain search criteria.

    ## Example Usage

    The following example shows how to get a Hosted Zone from its name and from this data how to create a Record Set.

    ```python
    import pulumi
    import pulumi_aws as aws

    selected = aws.route53.get_zone(name="test.com.",
        private_zone=True)
    www = aws.route53.Record("www",
        zone_id=selected.zone_id,
        name=f"www.{selected.name}",
        type="A",
        ttl=300,
        records=["10.0.0.1"])
    ```


    :param str name: The Hosted Zone name of the desired Hosted Zone.
    :param bool private_zone: Used with `name` field to get a private Hosted Zone.
    :param int resource_record_set_count: The number of Record Set in the Hosted Zone.
    :param Mapping[str, str] tags: Used with `name` field. A map of tags, each pair of which must exactly match a pair on the desired Hosted Zone.
    :param str vpc_id: Used with `name` field to get a private Hosted Zone associated with the vpc_id (in this case, private_zone is not mandatory).
    :param str zone_id: The Hosted Zone id of the desired Hosted Zone.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['privateZone'] = private_zone
    __args__['resourceRecordSetCount'] = resource_record_set_count
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:route53/getZone:getZone', __args__, opts=opts, typ=GetZoneResult).value

    return AwaitableGetZoneResult(
        caller_reference=__ret__.caller_reference,
        comment=__ret__.comment,
        id=__ret__.id,
        linked_service_description=__ret__.linked_service_description,
        linked_service_principal=__ret__.linked_service_principal,
        name=__ret__.name,
        name_servers=__ret__.name_servers,
        private_zone=__ret__.private_zone,
        resource_record_set_count=__ret__.resource_record_set_count,
        tags=__ret__.tags,
        vpc_id=__ret__.vpc_id,
        zone_id=__ret__.zone_id)
