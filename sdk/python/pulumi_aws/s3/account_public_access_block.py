# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AccountPublicAccessBlock(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
    """
    block_public_acls: pulumi.Output[bool]
    """
    Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
    * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
    * PUT Object calls will fail if the request includes an object ACL.
    """
    block_public_policy: pulumi.Output[bool]
    """
    Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
    * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
    """
    ignore_public_acls: pulumi.Output[bool]
    """
    Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
    * Ignore all public ACLs on buckets in this account and any objects that they contain.
    """
    restrict_public_buckets: pulumi.Output[bool]
    """
    Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
    * Only the bucket owner and AWS Services can access buckets with public policies.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, block_public_acls=None, block_public_policy=None, ignore_public_acls=None, restrict_public_buckets=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages S3 account-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
        
        > **NOTE:** Each AWS account may only have one S3 Public Access Block configuration. Multiple configurations of the resource against the same AWS account will cause a perpetual difference.
        
        > Advanced usage: To use a custom API endpoint for this resource, use the [`s3control` endpoint provider configuration](https://www.terraform.io/docs/providers/aws/index.html#s3control), not the `s3` endpoint provider configuration.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls will fail if the request includes an object ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore all public ACLs on buckets in this account and any objects that they contain.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access buckets with public policies.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_account_public_access_block.html.markdown.
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

            __props__['account_id'] = account_id
            __props__['block_public_acls'] = block_public_acls
            __props__['block_public_policy'] = block_public_policy
            __props__['ignore_public_acls'] = ignore_public_acls
            __props__['restrict_public_buckets'] = restrict_public_buckets
        super(AccountPublicAccessBlock, __self__).__init__(
            'aws:s3/accountPublicAccessBlock:AccountPublicAccessBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, block_public_acls=None, block_public_policy=None, ignore_public_acls=None, restrict_public_buckets=None):
        """
        Get an existing AccountPublicAccessBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: AWS account ID to configure. Defaults to automatically determined account ID of the this provider AWS provider.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls will fail if the request includes an object ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect existing bucket policies. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore all public ACLs on buckets in this account and any objects that they contain.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to `false`. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access buckets with public policies.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_account_public_access_block.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["account_id"] = account_id
        __props__["block_public_acls"] = block_public_acls
        __props__["block_public_policy"] = block_public_policy
        __props__["ignore_public_acls"] = ignore_public_acls
        __props__["restrict_public_buckets"] = restrict_public_buckets
        return AccountPublicAccessBlock(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

