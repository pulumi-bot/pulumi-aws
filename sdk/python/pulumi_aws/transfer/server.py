# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Server(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of Transfer Server
    """
    endpoint: pulumi.Output[str]
    """
    The endpoint of the Transfer Server (e.g. `s-12345678.server.transfer.REGION.amazonaws.com`)
    """
    endpoint_details: pulumi.Output[dict]
    """
    The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.

      * `vpc_endpoint_id` (`str`) - The ID of the VPC endpoint.
    """
    endpoint_type: pulumi.Output[str]
    """
    The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC_ENDPOINT`, your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
    """
    force_destroy: pulumi.Output[bool]
    """
    A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
    """
    host_key: pulumi.Output[str]
    """
    RSA private key (e.g. as generated by the `ssh-keygen -N "" -f my-new-server-key` command).
    """
    host_key_fingerprint: pulumi.Output[str]
    """
    This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
    """
    identity_provider_type: pulumi.Output[str]
    """
    The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
    """
    invocation_role: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
    """
    logging_role: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    url: pulumi.Output[str]
    """
    - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.
    """
    def __init__(__self__, resource_name, opts=None, endpoint_details=None, endpoint_type=None, force_destroy=None, host_key=None, identity_provider_type=None, invocation_role=None, logging_role=None, tags=None, url=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a AWS Transfer Server resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_server.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] endpoint_details: The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        :param pulumi.Input[str] endpoint_type: The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC_ENDPOINT`, your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] host_key: RSA private key (e.g. as generated by the `ssh-keygen -N "" -f my-new-server-key` command).
        :param pulumi.Input[str] identity_provider_type: The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        :param pulumi.Input[str] invocation_role: Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        :param pulumi.Input[str] logging_role: Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] url: - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.

        The **endpoint_details** object supports the following:

          * `vpc_endpoint_id` (`pulumi.Input[str]`) - The ID of the VPC endpoint.
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
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['endpoint_details'] = endpoint_details
            __props__['endpoint_type'] = endpoint_type
            __props__['force_destroy'] = force_destroy
            __props__['host_key'] = host_key
            __props__['identity_provider_type'] = identity_provider_type
            __props__['invocation_role'] = invocation_role
            __props__['logging_role'] = logging_role
            __props__['tags'] = tags
            __props__['url'] = url
            __props__['arn'] = None
            __props__['endpoint'] = None
            __props__['host_key_fingerprint'] = None
        super(Server, __self__).__init__(
            'aws:transfer/server:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, endpoint=None, endpoint_details=None, endpoint_type=None, force_destroy=None, host_key=None, host_key_fingerprint=None, identity_provider_type=None, invocation_role=None, logging_role=None, tags=None, url=None):
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of Transfer Server
        :param pulumi.Input[str] endpoint: The endpoint of the Transfer Server (e.g. `s-12345678.server.transfer.REGION.amazonaws.com`)
        :param pulumi.Input[dict] endpoint_details: The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        :param pulumi.Input[str] endpoint_type: The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC_ENDPOINT`, your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.  Defaults to `PUBLIC`.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] host_key: RSA private key (e.g. as generated by the `ssh-keygen -N "" -f my-new-server-key` command).
        :param pulumi.Input[str] host_key_fingerprint: This value contains the message-digest algorithm (MD5) hash of the server's host key. This value is equivalent to the output of the `ssh-keygen -l -E md5 -f my-new-server-key` command.
        :param pulumi.Input[str] identity_provider_type: The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        :param pulumi.Input[str] invocation_role: Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        :param pulumi.Input[str] logging_role: Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] url: - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.

        The **endpoint_details** object supports the following:

          * `vpc_endpoint_id` (`pulumi.Input[str]`) - The ID of the VPC endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["endpoint"] = endpoint
        __props__["endpoint_details"] = endpoint_details
        __props__["endpoint_type"] = endpoint_type
        __props__["force_destroy"] = force_destroy
        __props__["host_key"] = host_key
        __props__["host_key_fingerprint"] = host_key_fingerprint
        __props__["identity_provider_type"] = identity_provider_type
        __props__["invocation_role"] = invocation_role
        __props__["logging_role"] = logging_role
        __props__["tags"] = tags
        __props__["url"] = url
        return Server(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

