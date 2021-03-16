# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['EmailChannelArgs', 'EmailChannel']

@pulumi.input_type
class EmailChannelArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 from_address: pulumi.Input[str],
                 identity: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a EmailChannel resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] from_address: The email address used to send emails from.
        :param pulumi.Input[str] identity: The ARN of an identity verified with SES.
        :param pulumi.Input[str] role_arn: The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "from_address", from_address)
        pulumi.set(__self__, "identity", identity)
        pulumi.set(__self__, "role_arn", role_arn)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="fromAddress")
    def from_address(self) -> pulumi.Input[str]:
        """
        The email address used to send emails from.
        """
        return pulumi.get(self, "from_address")

    @from_address.setter
    def from_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "from_address", value)

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Input[str]:
        """
        The ARN of an identity verified with SES.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the channel is enabled or disabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


class EmailChannel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 from_address: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Pinpoint Email Channel resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        app = aws.pinpoint.App("app")
        identity = aws.ses.DomainIdentity("identity", domain="example.com")
        role = aws.iam.Role("role", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "pinpoint.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        email = aws.pinpoint.EmailChannel("email",
            application_id=app.application_id,
            from_address="user@example.com",
            identity=identity.arn,
            role_arn=role.arn)
        role_policy = aws.iam.RolePolicy("rolePolicy",
            role=role.id,
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": {
            "Action": [
              "mobileanalytics:PutEvents",
              "mobileanalytics:PutItems"
            ],
            "Effect": "Allow",
            "Resource": [
              "*"
            ]
          }
        }
        \"\"\")
        ```

        ## Import

        Pinpoint Email Channel can be imported using the `application-id`, e.g.

        ```sh
         $ pulumi import aws:pinpoint/emailChannel:EmailChannel email application-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] from_address: The email address used to send emails from.
        :param pulumi.Input[str] identity: The ARN of an identity verified with SES.
        :param pulumi.Input[str] role_arn: The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EmailChannelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Pinpoint Email Channel resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        app = aws.pinpoint.App("app")
        identity = aws.ses.DomainIdentity("identity", domain="example.com")
        role = aws.iam.Role("role", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "pinpoint.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        email = aws.pinpoint.EmailChannel("email",
            application_id=app.application_id,
            from_address="user@example.com",
            identity=identity.arn,
            role_arn=role.arn)
        role_policy = aws.iam.RolePolicy("rolePolicy",
            role=role.id,
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": {
            "Action": [
              "mobileanalytics:PutEvents",
              "mobileanalytics:PutItems"
            ],
            "Effect": "Allow",
            "Resource": [
              "*"
            ]
          }
        }
        \"\"\")
        ```

        ## Import

        Pinpoint Email Channel can be imported using the `application-id`, e.g.

        ```sh
         $ pulumi import aws:pinpoint/emailChannel:EmailChannel email application-id
        ```

        :param str resource_name: The name of the resource.
        :param EmailChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmailChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 from_address: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['enabled'] = enabled
            if from_address is None and not opts.urn:
                raise TypeError("Missing required property 'from_address'")
            __props__['from_address'] = from_address
            if identity is None and not opts.urn:
                raise TypeError("Missing required property 'identity'")
            __props__['identity'] = identity
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['messages_per_second'] = None
        super(EmailChannel, __self__).__init__(
            'aws:pinpoint/emailChannel:EmailChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            from_address: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[str]] = None,
            messages_per_second: Optional[pulumi.Input[int]] = None,
            role_arn: Optional[pulumi.Input[str]] = None) -> 'EmailChannel':
        """
        Get an existing EmailChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] from_address: The email address used to send emails from.
        :param pulumi.Input[str] identity: The ARN of an identity verified with SES.
        :param pulumi.Input[int] messages_per_second: Messages per second that can be sent.
        :param pulumi.Input[str] role_arn: The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_id"] = application_id
        __props__["enabled"] = enabled
        __props__["from_address"] = from_address
        __props__["identity"] = identity
        __props__["messages_per_second"] = messages_per_second
        __props__["role_arn"] = role_arn
        return EmailChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the channel is enabled or disabled. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="fromAddress")
    def from_address(self) -> pulumi.Output[str]:
        """
        The email address used to send emails from.
        """
        return pulumi.get(self, "from_address")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[str]:
        """
        The ARN of an identity verified with SES.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="messagesPerSecond")
    def messages_per_second(self) -> pulumi.Output[int]:
        """
        Messages per second that can be sent.
        """
        return pulumi.get(self, "messages_per_second")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
        """
        return pulumi.get(self, "role_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

