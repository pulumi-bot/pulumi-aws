# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ApnsSandboxChannel(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The application ID.
    """
    bundle_id: pulumi.Output[str]
    """
    The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
    """
    certificate: pulumi.Output[str]
    """
    The pem encoded TLS Certificate from Apple.
    """
    default_authentication_method: pulumi.Output[str]
    """
    The default authentication method used for APNs Sandbox. 
    __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
    You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
    If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the channel is enabled or disabled. Defaults to `true`.
    """
    private_key: pulumi.Output[str]
    """
    The Certificate Private Key file (ie. `.key` file).
    """
    team_id: pulumi.Output[str]
    """
    The ID assigned to your Apple developer account team. This value is provided on the Membership page.
    """
    token_key: pulumi.Output[str]
    """
    The `.p8` file that you download from your Apple developer account when you create an authentication key. 
    """
    token_key_id: pulumi.Output[str]
    """
    The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
    """
    def __init__(__self__, resource_name, opts=None, application_id=None, bundle_id=None, certificate=None, default_authentication_method=None, enabled=None, private_key=None, team_id=None, token_key=None, token_key_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Pinpoint APNs Sandbox Channel resource.
        
        > **Note:** All arguments, including certificates and tokens, will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] bundle_id: The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
        :param pulumi.Input[str] certificate: The pem encoded TLS Certificate from Apple.
        :param pulumi.Input[str] default_authentication_method: The default authentication method used for APNs Sandbox. 
               __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
               You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
               If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] private_key: The Certificate Private Key file (ie. `.key` file).
        :param pulumi.Input[str] team_id: The ID assigned to your Apple developer account team. This value is provided on the Membership page.
        :param pulumi.Input[str] token_key: The `.p8` file that you download from your Apple developer account when you create an authentication key. 
        :param pulumi.Input[str] token_key_id: The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_apns_sandbox_channel.html.markdown.
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

            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['bundle_id'] = bundle_id
            __props__['certificate'] = certificate
            __props__['default_authentication_method'] = default_authentication_method
            __props__['enabled'] = enabled
            __props__['private_key'] = private_key
            __props__['team_id'] = team_id
            __props__['token_key'] = token_key
            __props__['token_key_id'] = token_key_id
        super(ApnsSandboxChannel, __self__).__init__(
            'aws:pinpoint/apnsSandboxChannel:ApnsSandboxChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, application_id=None, bundle_id=None, certificate=None, default_authentication_method=None, enabled=None, private_key=None, team_id=None, token_key=None, token_key_id=None):
        """
        Get an existing ApnsSandboxChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] bundle_id: The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
        :param pulumi.Input[str] certificate: The pem encoded TLS Certificate from Apple.
        :param pulumi.Input[str] default_authentication_method: The default authentication method used for APNs Sandbox. 
               __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
               You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
               If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] private_key: The Certificate Private Key file (ie. `.key` file).
        :param pulumi.Input[str] team_id: The ID assigned to your Apple developer account team. This value is provided on the Membership page.
        :param pulumi.Input[str] token_key: The `.p8` file that you download from your Apple developer account when you create an authentication key. 
        :param pulumi.Input[str] token_key_id: The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_apns_sandbox_channel.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["application_id"] = application_id
        __props__["bundle_id"] = bundle_id
        __props__["certificate"] = certificate
        __props__["default_authentication_method"] = default_authentication_method
        __props__["enabled"] = enabled
        __props__["private_key"] = private_key
        __props__["team_id"] = team_id
        __props__["token_key"] = token_key
        __props__["token_key_id"] = token_key_id
        return ApnsSandboxChannel(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

