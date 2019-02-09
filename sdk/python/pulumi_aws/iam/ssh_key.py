# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SshKey(pulumi.CustomResource):
    encoding: pulumi.Output[str]
    """
    Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
    """
    fingerprint: pulumi.Output[str]
    """
    The MD5 message digest of the SSH public key.
    """
    public_key: pulumi.Output[str]
    """
    The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
    """
    ssh_public_key_id: pulumi.Output[str]
    """
    The unique identifier for the SSH public key.
    """
    status: pulumi.Output[str]
    """
    The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
    """
    username: pulumi.Output[str]
    """
    The name of the IAM user to associate the SSH public key with.
    """
    def __init__(__self__, resource_name, opts=None, encoding=None, public_key=None, status=None, username=None, __name__=None, __opts__=None):
        """
        Uploads an SSH public key and associates it with the specified IAM user.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] encoding: Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
        :param pulumi.Input[str] public_key: The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
        :param pulumi.Input[str] status: The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
        :param pulumi.Input[str] username: The name of the IAM user to associate the SSH public key with.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if encoding is None:
            raise TypeError('Missing required property encoding')
        __props__['encoding'] = encoding

        if public_key is None:
            raise TypeError('Missing required property public_key')
        __props__['public_key'] = public_key

        __props__['status'] = status

        if username is None:
            raise TypeError('Missing required property username')
        __props__['username'] = username

        __props__['fingerprint'] = None
        __props__['ssh_public_key_id'] = None

        super(SshKey, __self__).__init__(
            'aws:iam/sshKey:SshKey',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

