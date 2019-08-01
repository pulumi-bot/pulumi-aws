# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, secrets=None, id=None):
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        __self__.secrets = secrets
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_secret(secrets=None,opts=None):
    """
    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_secret.html.markdown.
    """
    __args__ = dict()

    __args__['secrets'] = secrets
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('aws:kms/getSecret:getSecret', __args__, opts=opts)

    return GetSecretResult(
        secrets=__ret__.get('secrets'),
        id=__ret__.get('id'))
