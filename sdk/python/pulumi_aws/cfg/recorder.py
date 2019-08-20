# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Recorder(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the recorder. Defaults to `default`. Changing it recreates the resource.
    """
    recording_group: pulumi.Output[dict]
    """
    Recording group - see below.
    """
    role_arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the IAM role.
    used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
    See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
    """
    def __init__(__self__, resource_name, opts=None, name=None, recording_group=None, role_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Config Configuration Recorder. Please note that this resource **does not start** the created recorder automatically.
        
        > **Note:** _Starting_ the Configuration Recorder requires a [delivery channel](https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html) (while delivery channel creation requires Configuration Recorder). This is why [`cfg.RecorderStatus`](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder_status.html) is a separate resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the recorder. Defaults to `default`. Changing it recreates the resource.
        :param pulumi.Input[dict] recording_group: Recording group - see below.
        :param pulumi.Input[str] role_arn: Amazon Resource Name (ARN) of the IAM role.
               used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
               See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_configuration_recorder.html.markdown.
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

            __props__['name'] = name
            __props__['recording_group'] = recording_group
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
        super(Recorder, __self__).__init__(
            'aws:cfg/recorder:Recorder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, recording_group=None, role_arn=None):
        """
        Get an existing Recorder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the recorder. Defaults to `default`. Changing it recreates the resource.
        :param pulumi.Input[dict] recording_group: Recording group - see below.
        :param pulumi.Input[str] role_arn: Amazon Resource Name (ARN) of the IAM role.
               used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account.
               See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_configuration_recorder.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["name"] = name
        __props__["recording_group"] = recording_group
        __props__["role_arn"] = role_arn
        return Recorder(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

