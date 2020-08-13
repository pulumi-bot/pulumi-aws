# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class AmiLaunchPermission(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    An AWS Account ID to add launch permissions.
    """
    image_id: pulumi.Output[str]
    """
    A region-unique name for the AMI.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, image_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Adds launch permission to Amazon Machine Image (AMI) from another AWS account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.AmiLaunchPermission("example",
            account_id="123456789012",
            image_id="ami-12345678")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: An AWS Account ID to add launch permissions.
        :param pulumi.Input[str] image_id: A region-unique name for the AMI.
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
            __props__['accountId'] = account_id
            if image_id is None:
                raise TypeError("Missing required property 'image_id'")
            __props__['imageId'] = image_id
        super(AmiLaunchPermission, __self__).__init__(
            'aws:ec2/amiLaunchPermission:AmiLaunchPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_id=None, image_id=None):
        """
        Get an existing AmiLaunchPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: An AWS Account ID to add launch permissions.
        :param pulumi.Input[str] image_id: A region-unique name for the AMI.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["image_id"] = image_id
        return AmiLaunchPermission(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
