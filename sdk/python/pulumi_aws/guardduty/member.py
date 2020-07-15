# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Member(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    AWS account ID for member account.
    """
    detector_id: pulumi.Output[str]
    """
    The detector ID of the GuardDuty account where you want to create member accounts.
    """
    disable_email_notification: pulumi.Output[bool]
    """
    Boolean whether an email notification is sent to the accounts. Defaults to `false`.
    """
    email: pulumi.Output[str]
    """
    Email address for member account.
    """
    invitation_message: pulumi.Output[str]
    """
    Message for invitation.
    """
    invite: pulumi.Output[bool]
    """
    Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
    """
    relationship_status: pulumi.Output[str]
    """
    The status of the relationship between the member account and its master account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, detector_id=None, disable_email_notification=None, email=None, invitation_message=None, invite=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage a GuardDuty member. To accept invitations in member accounts, see the `guardduty.InviteAccepter` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        master = aws.guardduty.Detector("master", enable=True)
        member_detector = aws.guardduty.Detector("memberDetector", enable=True,
        opts=ResourceOptions(provider="aws.dev"))
        member_member = aws.guardduty.Member("memberMember",
            account_id=member_detector.account_id,
            detector_id=master.id,
            email="required@example.com",
            invite=True,
            invitation_message="please accept guardduty invitation")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID for member account.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty account where you want to create member accounts.
        :param pulumi.Input[bool] disable_email_notification: Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        :param pulumi.Input[str] email: Email address for member account.
        :param pulumi.Input[str] invitation_message: Message for invitation.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
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

            if account_id is None:
                raise TypeError("Missing required property 'account_id'")
            __props__['account_id'] = account_id
            if detector_id is None:
                raise TypeError("Missing required property 'detector_id'")
            __props__['detector_id'] = detector_id
            __props__['disable_email_notification'] = disable_email_notification
            if email is None:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['invitation_message'] = invitation_message
            __props__['invite'] = invite
            __props__['relationship_status'] = None
        super(Member, __self__).__init__(
            'aws:guardduty/member:Member',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, detector_id=None, disable_email_notification=None, email=None, invitation_message=None, invite=None, relationship_status=None):
        """
        Get an existing Member resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID for member account.
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty account where you want to create member accounts.
        :param pulumi.Input[bool] disable_email_notification: Boolean whether an email notification is sent to the accounts. Defaults to `false`.
        :param pulumi.Input[str] email: Email address for member account.
        :param pulumi.Input[str] invitation_message: Message for invitation.
        :param pulumi.Input[bool] invite: Boolean whether to invite the account to GuardDuty as a member. Defaults to `false`. To detect if an invitation needs to be (re-)sent, the this provider state value is `true` based on a `relationship_status` of `Disabled`, `Enabled`, `Invited`, or `EmailVerificationInProgress`.
        :param pulumi.Input[str] relationship_status: The status of the relationship between the member account and its master account. More information can be found in [Amazon GuardDuty API Reference](https://docs.aws.amazon.com/guardduty/latest/ug/get-members.html).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["detector_id"] = detector_id
        __props__["disable_email_notification"] = disable_email_notification
        __props__["email"] = email
        __props__["invitation_message"] = invitation_message
        __props__["invite"] = invite
        __props__["relationship_status"] = relationship_status
        return Member(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
