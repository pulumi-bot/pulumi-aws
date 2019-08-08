# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetCertificateResult:
    """
    A collection of values returned by getCertificate.
    """
    def __init__(__self__, arn=None, domain=None, key_types=None, most_recent=None, statuses=None, types=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        Set to the ARN of the found certificate, suitable for referencing in other resources that support ACM certificates.
        """
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        __self__.domain = domain
        if key_types and not isinstance(key_types, list):
            raise TypeError("Expected argument 'key_types' to be a list")
        __self__.key_types = key_types
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        __self__.most_recent = most_recent
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        __self__.statuses = statuses
        if types and not isinstance(types, list):
            raise TypeError("Expected argument 'types' to be a list")
        __self__.types = types
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        delattr(self, "__await__")
        delattr(self, "__iter__")
        return self

    __iter__ = __await__

def get_certificate(domain=None,key_types=None,most_recent=None,statuses=None,types=None,opts=None):
    """
    Use this data source to get the ARN of a certificate in AWS Certificate
    Manager (ACM), you can reference
    it by domain without having to hard code the ARNs as input.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/acm_certificate.html.markdown.
    """
    __args__ = dict()

    __args__['domain'] = domain
    __args__['keyTypes'] = key_types
    __args__['mostRecent'] = most_recent
    __args__['statuses'] = statuses
    __args__['types'] = types
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:acm/getCertificate:getCertificate', __args__, opts=opts).value

    return GetCertificateResult(
        arn=__ret__.get('arn'),
        domain=__ret__.get('domain'),
        key_types=__ret__.get('keyTypes'),
        most_recent=__ret__.get('mostRecent'),
        statuses=__ret__.get('statuses'),
        types=__ret__.get('types'),
        id=__ret__.get('id'))
