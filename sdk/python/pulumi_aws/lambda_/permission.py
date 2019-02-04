# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Permission(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
    """
    event_source_token: pulumi.Output[str]
    """
    The Event Source Token to validate.  Used with [Alexa Skills][1].
    """
    function: pulumi.Output[str]
    """
    Name of the Lambda function whose resource policy you are updating
    """
    principal: pulumi.Output[str]
    """
    The principal who is getting this permission.
    e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
    such as `events.amazonaws.com` or `sns.amazonaws.com`.
    """
    qualifier: pulumi.Output[str]
    """
    Query parameter to specify function version or alias name.
    The permission will then apply to the specific qualified ARN.
    e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
    """
    source_account: pulumi.Output[str]
    """
    This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
    """
    source_arn: pulumi.Output[str]
    """
    When granting Amazon S3 or CloudWatch Events permission to
    invoke your function, you should specify this field with the Amazon Resource Name (ARN)
    for the S3 Bucket or CloudWatch Events Rule as its value.  This ensures that only events
    generated from the specified bucket or rule can invoke the function.
    API Gateway ARNs have a unique structure described
    [here](http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
    """
    statement_id: pulumi.Output[str]
    """
    A unique statement identifier. By default generated by Terraform.
    """
    statement_id_prefix: pulumi.Output[str]
    """
    A statement identifier prefix. Terraform will generate a unique suffix. Conflicts with `statement_id`.
    """
    def __init__(__self__, resource_name, opts=None, action=None, event_source_token=None, function=None, principal=None, qualifier=None, source_account=None, source_arn=None, statement_id=None, statement_id_prefix=None, __name__=None, __opts__=None):
        """
        Creates a Lambda permission to allow external sources invoking the Lambda function
        (e.g. CloudWatch Event Rule, SNS or S3).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        :param pulumi.Input[str] event_source_token: The Event Source Token to validate.  Used with [Alexa Skills][1].
        :param pulumi.Input[str] function: Name of the Lambda function whose resource policy you are updating
        :param pulumi.Input[str] principal: The principal who is getting this permission.
               e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal
               such as `events.amazonaws.com` or `sns.amazonaws.com`.
        :param pulumi.Input[str] qualifier: Query parameter to specify function version or alias name.
               The permission will then apply to the specific qualified ARN.
               e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        :param pulumi.Input[str] source_account: This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        :param pulumi.Input[str] source_arn: When granting Amazon S3 or CloudWatch Events permission to
               invoke your function, you should specify this field with the Amazon Resource Name (ARN)
               for the S3 Bucket or CloudWatch Events Rule as its value.  This ensures that only events
               generated from the specified bucket or rule can invoke the function.
               API Gateway ARNs have a unique structure described
               [here](http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        :param pulumi.Input[str] statement_id: A unique statement identifier. By default generated by Terraform.
        :param pulumi.Input[str] statement_id_prefix: A statement identifier prefix. Terraform will generate a unique suffix. Conflicts with `statement_id`.
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

        if not action:
            raise TypeError('Missing required property action')
        __props__['action'] = action

        __props__['event_source_token'] = event_source_token

        if not function:
            raise TypeError('Missing required property function')
        __props__['function'] = function

        if not principal:
            raise TypeError('Missing required property principal')
        __props__['principal'] = principal

        __props__['qualifier'] = qualifier

        __props__['source_account'] = source_account

        __props__['source_arn'] = source_arn

        __props__['statement_id'] = statement_id

        __props__['statement_id_prefix'] = statement_id_prefix

        super(Permission, __self__).__init__(
            'aws:lambda/permission:Permission',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

