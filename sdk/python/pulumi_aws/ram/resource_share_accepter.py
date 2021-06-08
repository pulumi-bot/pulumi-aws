# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourceShareAccepterArgs', 'ResourceShareAccepter']

@pulumi.input_type
class ResourceShareAccepterArgs:
    def __init__(__self__, *,
                 share_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a ResourceShareAccepter resource.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.
        """
        pulumi.set(__self__, "share_arn", share_arn)

    @property
    @pulumi.getter(name="shareArn")
    def share_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the resource share.
        """
        return pulumi.get(self, "share_arn")

    @share_arn.setter
    def share_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "share_arn", value)


@pulumi.input_type
class _ResourceShareAccepterState:
    def __init__(__self__, *,
                 invitation_arn: Optional[pulumi.Input[str]] = None,
                 receiver_account_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sender_account_id: Optional[pulumi.Input[str]] = None,
                 share_arn: Optional[pulumi.Input[str]] = None,
                 share_id: Optional[pulumi.Input[str]] = None,
                 share_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResourceShareAccepter resources.
        :param pulumi.Input[str] invitation_arn: The ARN of the resource share invitation.
        :param pulumi.Input[str] receiver_account_id: The account ID of the receiver account which accepts the invitation.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resources: A list of the resource ARNs shared via the resource share.
        :param pulumi.Input[str] sender_account_id: The account ID of the sender account which submits the invitation.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.
        :param pulumi.Input[str] share_id: The ID of the resource share as displayed in the console.
        :param pulumi.Input[str] share_name: The name of the resource share.
        :param pulumi.Input[str] status: The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
        """
        if invitation_arn is not None:
            pulumi.set(__self__, "invitation_arn", invitation_arn)
        if receiver_account_id is not None:
            pulumi.set(__self__, "receiver_account_id", receiver_account_id)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if sender_account_id is not None:
            pulumi.set(__self__, "sender_account_id", sender_account_id)
        if share_arn is not None:
            pulumi.set(__self__, "share_arn", share_arn)
        if share_id is not None:
            pulumi.set(__self__, "share_id", share_id)
        if share_name is not None:
            pulumi.set(__self__, "share_name", share_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="invitationArn")
    def invitation_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the resource share invitation.
        """
        return pulumi.get(self, "invitation_arn")

    @invitation_arn.setter
    def invitation_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invitation_arn", value)

    @property
    @pulumi.getter(name="receiverAccountId")
    def receiver_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The account ID of the receiver account which accepts the invitation.
        """
        return pulumi.get(self, "receiver_account_id")

    @receiver_account_id.setter
    def receiver_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "receiver_account_id", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the resource ARNs shared via the resource share.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter(name="senderAccountId")
    def sender_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The account ID of the sender account which submits the invitation.
        """
        return pulumi.get(self, "sender_account_id")

    @sender_account_id.setter
    def sender_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sender_account_id", value)

    @property
    @pulumi.getter(name="shareArn")
    def share_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the resource share.
        """
        return pulumi.get(self, "share_arn")

    @share_arn.setter
    def share_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_arn", value)

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource share as displayed in the console.
        """
        return pulumi.get(self, "share_id")

    @share_id.setter
    def share_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_id", value)

    @property
    @pulumi.getter(name="shareName")
    def share_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource share.
        """
        return pulumi.get(self, "share_name")

    @share_name.setter
    def share_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ResourceShareAccepter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 share_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
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
            opts=pulumi.ResourceOptions(provider=aws["alternate"]))
        receiver = aws.get_caller_identity()
        sender_invite = aws.ram.PrincipalAssociation("senderInvite",
            principal=receiver.account_id,
            resource_share_arn=sender_share.arn,
            opts=pulumi.ResourceOptions(provider=aws["alternate"]))
        receiver_accept = aws.ram.ResourceShareAccepter("receiverAccept", share_arn=sender_invite.resource_share_arn)
        ```

        ## Import

        Resource share accepters can be imported using the resource share ARN, e.g.

        ```sh
         $ pulumi import aws:ram/resourceShareAccepter:ResourceShareAccepter example arn:aws:ram:us-east-1:123456789012:resource-share/c4b56393-e8d9-89d9-6dc9-883752de4767
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] share_arn: The ARN of the resource share.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: ResourceShareAccepterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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
            opts=pulumi.ResourceOptions(provider=aws["alternate"]))
        receiver = aws.get_caller_identity()
        sender_invite = aws.ram.PrincipalAssociation("senderInvite",
            principal=receiver.account_id,
            resource_share_arn=sender_share.arn,
            opts=pulumi.ResourceOptions(provider=aws["alternate"]))
        receiver_accept = aws.ram.ResourceShareAccepter("receiverAccept", share_arn=sender_invite.resource_share_arn)
        ```

        ## Import

        Resource share accepters can be imported using the resource share ARN, e.g.

        ```sh
         $ pulumi import aws:ram/resourceShareAccepter:ResourceShareAccepter example arn:aws:ram:us-east-1:123456789012:resource-share/c4b56393-e8d9-89d9-6dc9-883752de4767
        ```

        :param str resource_name_: The name of the resource.
        :param ResourceShareAccepterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceShareAccepterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 share_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceShareAccepterArgs.__new__(ResourceShareAccepterArgs)

            if share_arn is None and not opts.urn:
                raise TypeError("Missing required property 'share_arn'")
            __props__.__dict__["share_arn"] = share_arn
            __props__.__dict__["invitation_arn"] = None
            __props__.__dict__["receiver_account_id"] = None
            __props__.__dict__["resources"] = None
            __props__.__dict__["sender_account_id"] = None
            __props__.__dict__["share_id"] = None
            __props__.__dict__["share_name"] = None
            __props__.__dict__["status"] = None
        super(ResourceShareAccepter, __self__).__init__(
            'aws:ram/resourceShareAccepter:ResourceShareAccepter',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
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

        :param str resource_name_: The unique name of the resulting resource.
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

        __props__ = _ResourceShareAccepterState.__new__(_ResourceShareAccepterState)

        __props__.__dict__["invitation_arn"] = invitation_arn
        __props__.__dict__["receiver_account_id"] = receiver_account_id
        __props__.__dict__["resources"] = resources
        __props__.__dict__["sender_account_id"] = sender_account_id
        __props__.__dict__["share_arn"] = share_arn
        __props__.__dict__["share_id"] = share_id
        __props__.__dict__["share_name"] = share_name
        __props__.__dict__["status"] = status
        return ResourceShareAccepter(resource_name_, opts=opts, __props__=__props__)

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

