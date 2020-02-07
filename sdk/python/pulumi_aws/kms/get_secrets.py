# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecretsResult:
    """
    A collection of values returned by getSecrets.
    """
    def __init__(__self__, id=None, plaintext=None, secrets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if plaintext and not isinstance(plaintext, dict):
            raise TypeError("Expected argument 'plaintext' to be a dict")
        __self__.plaintext = plaintext
        """
        Map containing each `secret` `name` as the key with its decrypted plaintext value
        """
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        __self__.secrets = secrets
class AwaitableGetSecretsResult(GetSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretsResult(
            id=self.id,
            plaintext=self.plaintext,
            secrets=self.secrets)

def get_secrets(secrets=None,opts=None):
    """
    Decrypt multiple secrets from data encrypted with the AWS KMS service.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_secrets.html.markdown.


    :param list secrets: One or more encrypted payload definitions from the KMS service. See the Secret Definitions below.

    The **secrets** object supports the following:

      * `context` (`dict`) - An optional mapping that makes up the Encryption Context for the secret.
      * `grantTokens` (`list`) - An optional list of Grant Tokens for the secret.
      * `name` (`str`) - The name to export this secret under in the attributes.
      * `payload` (`str`) - Base64 encoded payload, as returned from a KMS encrypt operation.
    """
    __args__ = dict()


    __args__['secrets'] = secrets
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:kms/getSecrets:getSecrets', __args__, opts=opts).value

    return AwaitableGetSecretsResult(
        id=__ret__.get('id'),
        plaintext=__ret__.get('plaintext'),
        secrets=__ret__.get('secrets'))
