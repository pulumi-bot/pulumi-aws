# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Account(pulumi.CustomResource):
    cloudwatch_role_arn: pulumi.Output[str]
    """
    The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
    See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
    Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
    """
    throttle_settings: pulumi.Output[dict]
    """
    Account-Level throttle settings. See exported fields below.

      * `burstLimit` (`float`) - The absolute maximum number of times API Gateway allows the API to be called per second (RPS).
      * `rate_limit` (`float`) - The number of times API Gateway allows the API to be called per second on average (RPS).
    """
    def __init__(__self__, resource_name, opts=None, cloudwatch_role_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a settings of an API Gateway Account. Settings is applied region-wide per `provider` block.

        > **Note:** As there is no API method for deleting account settings or resetting it to defaults, destroying this resource will keep your account settings intact

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_account.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloudwatch_role_arn: The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
               See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
               Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
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

            __props__['cloudwatch_role_arn'] = cloudwatch_role_arn
            __props__['throttle_settings'] = None
        super(Account, __self__).__init__(
            'aws:apigateway/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cloudwatch_role_arn=None, throttle_settings=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloudwatch_role_arn: The ARN of an IAM role for CloudWatch (to allow logging & monitoring).
               See more [in AWS Docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-stage-settings.html#how-to-stage-settings-console).
               Logging & monitoring can be enabled/disabled and otherwise tuned on the API Gateway Stage level.
        :param pulumi.Input[dict] throttle_settings: Account-Level throttle settings. See exported fields below.

        The **throttle_settings** object supports the following:

          * `burstLimit` (`pulumi.Input[float]`) - The absolute maximum number of times API Gateway allows the API to be called per second (RPS).
          * `rate_limit` (`pulumi.Input[float]`) - The number of times API Gateway allows the API to be called per second on average (RPS).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cloudwatch_role_arn"] = cloudwatch_role_arn
        __props__["throttle_settings"] = throttle_settings
        return Account(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

