# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ResourceShareAccepter']


class ResourceShareAccepter(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 share_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manage accepting a Resource Access Manager (RAM) Resource Share invitation. From a _receiver_ AWS account, accept an invitation to share resources that were shared by a _sender_ AWS account. To create a resource share in the _sender_, see the `ram.ResourceShare` resource.

        > **Note:** If both AWS accounts are in the same Organization and [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), this resource is not necessary as RAM Resource Share invitations are not used.

        ## Example Usage

        This configuration provides an example of using multiple AWS providers to configure two different AWS accounts. In the _sender_ account, the configuration creates a `ram.ResourceShare` and uses a data source in the _receiver_ account to create a `ram.PrincipalAssociation` resource with the _receiver's_ account ID. In the _receiver_ account, the configuration accepts the invitation to share resources with the `ram.ResourceShareAccepter`.

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        alternate = pulumi.providers.Aws("alternate", profile="profile1")
        sender_share = aws.ram.ResourceShare("senderShare",
            allow_external_principals=True,
            tags={
                "Name": "tf-test-resource-share",
            },
            opts=ResourceOptions(provider="aws.alternate"))
        receiver = aws.get_caller_identity()
        sender_invite = aws.ram.PrincipalAssociation("senderInvite",
            principal=receiver.account_id,
            resource_share_arn=sender_share.arn,
            opts=ResourceOptions(provider="aws.alternate"))
        receiver_accept = aws.ram.ResourceShareAccepter("receiverAccept", share_arn=sender_invite.resource_share_arn)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.
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
            opts.version = _utilities.get_version()
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            invitation_arn: Optional[pulumi.Input[str]] = None,
            receiver_account_id: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            sender_account_id: Optional[pulumi.Input[str]] = None,
            share_arn: Optional[pulumi.Input[str]] = None,
            share_id: Optional[pulumi.Input[str]] = None,
            share_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ResourceShareAccepter':
        """
        Get an existing ResourceShareAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] invitation_arn: The ARN of the resource share invitation.
        :param pulumi.Input[str] receiver_account_id: The account ID of the receiver account which accepts the invitation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resources: A list of the resource ARNs shared via the resource share.
        :param pulumi.Input[str] sender_account_id: The account ID of the sender account which submits the invitation.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.
        :param pulumi.Input[str] share_id: The ID of the resource share as displayed in the console.
        :param pulumi.Input[str] share_name: The name of the resource share.
        :param pulumi.Input[str] status: The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
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

    @property
    @pulumi.getter(name="invitationArn")
    def invitation_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the resource share invitation.
        """
        return pulumi.get(self, "invitation_arn")

    @property
    @pulumi.getter(name="receiverAccountId")
    def receiver_account_id(self) -> pulumi.Output[str]:
        """
        The account ID of the receiver account which accepts the invitation.
        """
        return pulumi.get(self, "receiver_account_id")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of the resource ARNs shared via the resource share.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="senderAccountId")
    def sender_account_id(self) -> pulumi.Output[str]:
        """
        The account ID of the sender account which submits the invitation.
        """
        return pulumi.get(self, "sender_account_id")

    @property
    @pulumi.getter(name="shareArn")
    def share_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the resource share.
        """
        return pulumi.get(self, "share_arn")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> pulumi.Output[str]:
        """
        The ID of the resource share as displayed in the console.
        """
        return pulumi.get(self, "share_id")

    @property
    @pulumi.getter(name="shareName")
    def share_name(self) -> pulumi.Output[str]:
        """
        The name of the resource share.
        """
        return pulumi.get(self, "share_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
        """
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

