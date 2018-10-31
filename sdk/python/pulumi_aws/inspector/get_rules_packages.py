# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetRulesPackagesResult(object):
    """
    A collection of values returned by getRulesPackages.
    """
    def __init__(__self__, arns=None, id=None):
        if arns and not isinstance(arns, list):
            raise TypeError('Expected argument arns to be a list')
        __self__.arns = arns
        """
        A list of the AWS Inspector Rules Packages arns available in the AWS region.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_rules_packages():
    """
    The AWS Inspector Rules Packages data source allows access to the list of AWS
    Inspector Rules Packages which can be used by AWS Inspector within the region
    configured in the provider.
    """
    __args__ = dict()

    __ret__ = pulumi.runtime.invoke('aws:inspector/getRulesPackages:getRulesPackages', __args__)

    return GetRulesPackagesResult(
        arns=__ret__.get('arns'),
        id=__ret__.get('id'))
