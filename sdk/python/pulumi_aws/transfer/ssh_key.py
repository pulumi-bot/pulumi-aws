# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SshKey(pulumi.CustomResource):
    body: pulumi.Output[str]
    """
    The public key portion of an SSH key pair.
    """
    server_id: pulumi.Output[str]
    """
    The Server ID of the Transfer Server (e.g. `s-12345678`)
    """
    user_name: pulumi.Output[str]
    """
    The name of the user account that is assigned to one or more servers.
    """
    def __init__(__self__, resource_name, opts=None, body=None, server_id=None, user_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a AWS Transfer User SSH Key resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: The public key portion of an SSH key pair.
        :param pulumi.Input[str] server_id: The Server ID of the Transfer Server (e.g. `s-12345678`)
        :param pulumi.Input[str] user_name: The name of the user account that is assigned to one or more servers.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_ssh_key.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            if body is None:
                raise TypeError("Missing required property 'body'")
            __props__['body'] = body
            if server_id is None:
                raise TypeError("Missing required property 'server_id'")
            __props__['server_id'] = server_id
            if user_name is None:
                raise TypeError("Missing required property 'user_name'")
            __props__['user_name'] = user_name
        super(SshKey, __self__).__init__(
            'aws:transfer/sshKey:SshKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, body=None, server_id=None, user_name=None):
        """
        Get an existing SshKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: The public key portion of an SSH key pair.
        :param pulumi.Input[str] server_id: The Server ID of the Transfer Server (e.g. `s-12345678`)
        :param pulumi.Input[str] user_name: The name of the user account that is assigned to one or more servers.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_ssh_key.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["body"] = body
        __props__["server_id"] = server_id
        __props__["user_name"] = user_name
        return SshKey(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

