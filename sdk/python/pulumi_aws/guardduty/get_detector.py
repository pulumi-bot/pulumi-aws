# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetDetectorResult',
    'AwaitableGetDetectorResult',
    'get_detector',
]

@pulumi.output_type
class GetDetectorResult:
    """
    A collection of values returned by getDetector.
    """
    def __init__(__self__, finding_publishing_frequency=None, id=None, service_role_arn=None, status=None):
        if finding_publishing_frequency and not isinstance(finding_publishing_frequency, str):
            raise TypeError("Expected argument 'finding_publishing_frequency' to be a str")
        pulumi.set(__self__, "finding_publishing_frequency", finding_publishing_frequency)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_role_arn and not isinstance(service_role_arn, str):
            raise TypeError("Expected argument 'service_role_arn' to be a str")
        pulumi.set(__self__, "service_role_arn", service_role_arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="findingPublishingFrequency")
    def finding_publishing_frequency(self) -> str:
        """
        The frequency of notifications sent about subsequent finding occurrences.
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> str:
        """
        The service-linked role that grants GuardDuty access to the resources in the AWS account.
        """
        return pulumi.get(self, "service_role_arn")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The current status of the detector.
        """
        return pulumi.get(self, "status")


class AwaitableGetDetectorResult(GetDetectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDetectorResult(
            finding_publishing_frequency=self.finding_publishing_frequency,
            id=self.id,
            service_role_arn=self.service_role_arn,
            status=self.status)


def get_detector(id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDetectorResult:
    """
    Retrieve information about a GuardDuty detector.


    :param str id: The ID of the detector.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:guardduty/getDetector:getDetector', __args__, opts=opts, typ=GetDetectorResult).value

    return AwaitableGetDetectorResult(
        finding_publishing_frequency=__ret__.finding_publishing_frequency,
        id=__ret__.id,
        service_role_arn=__ret__.service_role_arn,
        status=__ret__.status)
