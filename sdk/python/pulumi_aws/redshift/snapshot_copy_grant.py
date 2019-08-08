# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SnapshotCopyGrant(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of snapshot copy grant
    """
    kms_key_id: pulumi.Output[str]
    """
    The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
    """
    snapshot_copy_grant_name: pulumi.Output[str]
    """
    A friendly name for identifying the grant.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, kms_key_id=None, snapshot_copy_grant_name=None, tags=None, __name__=None, __opts__=None):
        """
        Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.
        
        Note that the grant must exist in the destination region, and not in the region of the cluster.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_id: The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
        :param pulumi.Input[str] snapshot_copy_grant_name: A friendly name for identifying the grant.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_snapshot_copy_grant.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['kms_key_id'] = kms_key_id
        if snapshot_copy_grant_name is None:
            raise TypeError("Missing required property 'snapshot_copy_grant_name'")
        __props__['snapshot_copy_grant_name'] = snapshot_copy_grant_name
        __props__['tags'] = tags
        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(SnapshotCopyGrant, __self__).__init__(
            'aws:redshift/snapshotCopyGrant:SnapshotCopyGrant',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

