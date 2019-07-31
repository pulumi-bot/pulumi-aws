# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
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
    """
    endpoint_type: pulumi.Output[str]
    """
    The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC_ENDPOINT`, your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.
    """
    force_destroy: pulumi.Output[bool]
    """
    A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
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
    def __init__(__self__, resource_name, opts=None, endpoint_details=None, endpoint_type=None, force_destroy=None, identity_provider_type=None, invocation_role=None, logging_role=None, tags=None, url=None, __name__=None, __opts__=None):
        """
        Provides a AWS Transfer Server resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] endpoint_details: The virtual private cloud (VPC) endpoint settings that you want to configure for your SFTP server. Fields documented below.
        :param pulumi.Input[str] endpoint_type: The type of endpoint that you want your SFTP server connect to. If you connect to a `VPC_ENDPOINT`, your SFTP server isn't accessible over the public internet. If you want to connect your SFTP server via public internet, set `PUBLIC`.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all users associated with the server should be deleted so that the Server can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] identity_provider_type: The mode of authentication enabled for this service. The default value is `SERVICE_MANAGED`, which allows you to store and access SFTP user credentials within the service. `API_GATEWAY` indicates that user authentication requires a call to an API Gateway endpoint URL provided by you to integrate an identity provider of your choice.
        :param pulumi.Input[str] invocation_role: Amazon Resource Name (ARN) of the IAM role used to authenticate the user account with an `identity_provider_type` of `API_GATEWAY`.
        :param pulumi.Input[str] logging_role: Amazon Resource Name (ARN) of an IAM role that allows the service to write your SFTP users’ activity to your Amazon CloudWatch logs for monitoring and auditing purposes.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] url: - URL of the service endpoint used to authenticate users with an `identity_provider_type` of `API_GATEWAY`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_server.html.markdown.
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

        __props__['endpoint_details'] = endpoint_details

        __props__['endpoint_type'] = endpoint_type

        __props__['force_destroy'] = force_destroy

        __props__['identity_provider_type'] = identity_provider_type

        __props__['invocation_role'] = invocation_role

        __props__['logging_role'] = logging_role

        __props__['tags'] = tags

        __props__['url'] = url

        __props__['arn'] = None
        __props__['endpoint'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Server, __self__).__init__(
            'aws:transfer/server:Server',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

