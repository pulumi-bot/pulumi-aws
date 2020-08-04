# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class LogResourcePolicy(pulumi.CustomResource):
    policy_document: pulumi.Output[str]
    """
    Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
    """
    policy_name: pulumi.Output[str]
    """
    Name of the resource policy.
    """
    def __init__(__self__, resource_name, opts=None, policy_document=None, policy_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage a CloudWatch log resource policy.

        ## Example Usage
        ### Elasticsearch Log Publishing

        ```python
        import pulumi
        import pulumi_aws as aws

        elasticsearch_log_publishing_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:PutLogEventsBatch",
            ],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["es.amazonaws.com"],
                type="Service",
            )],
            resources=["arn:aws:logs:*"],
        )])
        elasticsearch_log_publishing_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("elasticsearch-log-publishing-policyLogResourcePolicy",
            policy_document=elasticsearch_log_publishing_policy_policy_document.json,
            policy_name="elasticsearch-log-publishing-policy")
        ```
        ### Route53 Query Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        route53_query_logging_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "logs:CreateLogStream",
                "logs:PutLogEvents",
            ],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["route53.amazonaws.com"],
                type="Service",
            )],
            resources=["arn:aws:logs:*:*:log-group:/aws/route53/*"],
        )])
        route53_query_logging_policy_log_resource_policy = aws.cloudwatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy",
            policy_document=route53_query_logging_policy_policy_document.json,
            policy_name="route53-query-logging-policy")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        :param pulumi.Input[str] policy_name: Name of the resource policy.
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

            if policy_document is None:
                raise TypeError("Missing required property 'policy_document'")
            __props__['policy_document'] = policy_document
            if policy_name is None:
                raise TypeError("Missing required property 'policy_name'")
            __props__['policy_name'] = policy_name
        super(LogResourcePolicy, __self__).__init__(
            'aws:cloudwatch/logResourcePolicy:LogResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, policy_document=None, policy_name=None):
        """
        Get an existing LogResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_document: Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
        :param pulumi.Input[str] policy_name: Name of the resource policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["policy_document"] = policy_document
        __props__["policy_name"] = policy_name
        return LogResourcePolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
