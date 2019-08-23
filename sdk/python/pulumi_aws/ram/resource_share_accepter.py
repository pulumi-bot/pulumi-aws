# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ResourceShareAccepter(pulumi.CustomResource):
    invitation_arn: pulumi.Output[str]
    """
    The ARN of the resource share invitation.
    """
    receiver_account_id: pulumi.Output[str]
    """
    The account ID of the receiver account which accepts the invitation.
    """
    resources: pulumi.Output[list]
    """
    A list of the resource ARNs shared via the resource share.
    """
    sender_account_id: pulumi.Output[str]
    """
    The account ID of the sender account which extends the invitation.
    """
    share_arn: pulumi.Output[str]
    """
    The ARN of the resource share.
    """
    share_id: pulumi.Output[str]
    """
    The ID of the resource share as displayed in the console.
    """
    share_name: pulumi.Output[str]
    """
    The name of the resource share.
    """
    status: pulumi.Output[str]
    """
    The status of the invitation (e.g., ACCEPTED, REJECTED).
    """
    def __init__(__self__, resource_name, opts=None, share_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Manage accepting a Resource Access Manager (RAM) Resource Share invitation. From a _receiver_ AWS account, accept an invitation to share resources that were shared by a _sender_ AWS account. To create a resource share in the _sender_, see the [`ram.ResourceShare` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_share.html).
        
        > **Note:** If both AWS accounts are in the same Organization and [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), this resource is not necessary as RAM Resource Share invitations are not used.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_share_accepter.html.markdown.
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

            if share_arn is None:
                raise TypeError("Missing required property 'share_arn'")
            __props__['share_arn'] = share_arn
            __props__['invitation_arn'] = None
            __props__['receiver_account_id'] = None
            __props__['resources'] = None
            __props__['sender_account_id'] = None
            __props__['share_id'] = None
            __props__['share_name'] = None
            __props__['status'] = None
        super(ResourceShareAccepter, __self__).__init__(
            'aws:ram/resourceShareAccepter:ResourceShareAccepter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, invitation_arn=None, receiver_account_id=None, resources=None, sender_account_id=None, share_arn=None, share_id=None, share_name=None, status=None):
        """
        Get an existing ResourceShareAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] invitation_arn: The ARN of the resource share invitation.
        :param pulumi.Input[str] receiver_account_id: The account ID of the receiver account which accepts the invitation.
        :param pulumi.Input[list] resources: A list of the resource ARNs shared via the resource share.
        :param pulumi.Input[str] sender_account_id: The account ID of the sender account which extends the invitation.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.
        :param pulumi.Input[str] share_id: The ID of the resource share as displayed in the console.
        :param pulumi.Input[str] share_name: The name of the resource share.
        :param pulumi.Input[str] status: The status of the invitation (e.g., ACCEPTED, REJECTED).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_share_accepter.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["invitation_arn"] = invitation_arn
        __props__["receiver_account_id"] = receiver_account_id
        __props__["resources"] = resources
        __props__["sender_account_id"] = sender_account_id
        __props__["share_arn"] = share_arn
        __props__["share_id"] = share_id
        __props__["share_name"] = share_name
        __props__["status"] = status
        return ResourceShareAccepter(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

