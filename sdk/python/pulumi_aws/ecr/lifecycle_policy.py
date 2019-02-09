# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class LifecyclePolicy(pulumi.CustomResource):
    policy: pulumi.Output[str]
    """
    The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
    """
    registry_id: pulumi.Output[str]
    """
    The registry ID where the repository was created.
    """
    repository: pulumi.Output[str]
    """
    Name of the repository to apply the policy.
    """
    def __init__(__self__, resource_name, opts=None, policy=None, repository=None, __name__=None, __opts__=None):
        """
        Manages an ECR repository lifecycle policy.
        
        > **NOTE:** Only one `aws_ecr_lifecycle_policy` resource can be used with the same ECR repository. To apply multiple rules, they must be combined in the `policy` JSON.
        
        > **NOTE:** The AWS ECR API seems to reorder rules based on `rulePriority`. If you define multiple rules that are not sorted in ascending `rulePriority` order in the Terraform code, the resource will be flagged for recreation every `terraform plan`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
        :param pulumi.Input[str] repository: Name of the repository to apply the policy.
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

        if policy is None:
            raise TypeError('Missing required property policy')
        __props__['policy'] = policy

        if repository is None:
            raise TypeError('Missing required property repository')
        __props__['repository'] = repository

        __props__['registry_id'] = None

        super(LifecyclePolicy, __self__).__init__(
            'aws:ecr/lifecyclePolicy:LifecyclePolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

