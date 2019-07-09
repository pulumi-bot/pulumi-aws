# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSecretsResult:
    """
    A collection of values returned by getSecrets.
    """
    def __init__(__self__, plaintext=None, secrets=None, id=None):
        if plaintext and not isinstance(plaintext, dict):
            raise TypeError("Expected argument 'plaintext' to be a dict")
        __self__.plaintext = plaintext
        """
        Map containing each `secret` `name` as the key with its decrypted plaintext value
        """
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        __self__.secrets = secrets
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_secrets(secrets=None,opts=None):
    __args__ = dict()

    __args__['secrets'] = secrets
    __ret__ = await pulumi.runtime.invoke('aws:kms/getSecrets:getSecrets', __args__, opts=opts)

    return GetSecretsResult(
        plaintext=__ret__.get('plaintext'),
        secrets=__ret__.get('secrets'),
        id=__ret__.get('id'))
