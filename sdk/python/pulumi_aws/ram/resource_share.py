# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ResourceShare(pulumi.CustomResource):
    allow_external_principals: pulumi.Output[bool]
    """
    Indicates whether principals outside your organization can be associated with a resource share.
    """
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the resource share.
    """
    name: pulumi.Output[str]
    """
    The name of the resource share.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource share.
    """
    def __init__(__self__, resource_name, opts=None, allow_external_principals=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Resource Access Manager (RAM) Resource Share. To association principals with the share, see the [`ram.PrincipalAssociation` resource](https://www.terraform.io/docs/providers/aws/r/ram_principal_association.html). To associate resources with the share, see the [`ram.ResourceAssociation` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_association.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_external_principals: Indicates whether principals outside your organization can be associated with a resource share.
        :param pulumi.Input[str] name: The name of the resource share.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource share.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_share.html.markdown.
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

            __props__['allow_external_principals'] = allow_external_principals
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(ResourceShare, __self__).__init__(
            'aws:ram/resourceShare:ResourceShare',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allow_external_principals=None, arn=None, name=None, tags=None):
        """
        Get an existing ResourceShare resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_external_principals: Indicates whether principals outside your organization can be associated with a resource share.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the resource share.
        :param pulumi.Input[str] name: The name of the resource share.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource share.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_share.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["allow_external_principals"] = allow_external_principals
        __props__["arn"] = arn
        __props__["name"] = name
        __props__["tags"] = tags
        return ResourceShare(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

