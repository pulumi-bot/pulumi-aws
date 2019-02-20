# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetServiceAccountResult:
    """
    A collection of values returned by getServiceAccount.
    """
    def __init__(__self__, arn=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        The ARN of the AWS ELB service account in the selected region.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_service_account(region=None,opts=None):
    """
    Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
    in a given region for the purpose of whitelisting in S3 bucket policy.
    """
    __args__ = dict()

    __args__['region'] = region
    __ret__ = await pulumi.runtime.invoke('aws:elasticloadbalancing/getServiceAccount:getServiceAccount', __args__, opts=opts)

    return GetServiceAccountResult(
        arn=__ret__.get('arn'),
        id=__ret__.get('id'))
