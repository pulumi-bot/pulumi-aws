# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PatchBaseline(pulumi.CustomResource):
    approval_rules: pulumi.Output[list]
    """
    A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
    """
    approved_patches: pulumi.Output[list]
    """
    A list of explicitly approved patches for the baseline.
    """
    approved_patches_compliance_level: pulumi.Output[str]
    """
    Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
    """
    description: pulumi.Output[str]
    """
    The description of the patch baseline.
    """
    global_filters: pulumi.Output[list]
    """
    A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
    """
    name: pulumi.Output[str]
    """
    The name of the patch baseline.
    """
    operating_system: pulumi.Output[str]
    """
    Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `SUSE`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
    """
    rejected_patches: pulumi.Output[list]
    """
    A list of rejected patches.
    """
    tags: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, approval_rules=None, approved_patches=None, approved_patches_compliance_level=None, description=None, global_filters=None, name=None, operating_system=None, rejected_patches=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an SSM Patch Baseline resource
        
        > **NOTE on Patch Baselines:** The `approved_patches` and `approval_rule` are 
        both marked as optional fields, but the Patch Baseline requires that at least one
        of them is specified.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] approval_rules: A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
        :param pulumi.Input[list] approved_patches: A list of explicitly approved patches for the baseline.
        :param pulumi.Input[str] approved_patches_compliance_level: Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
        :param pulumi.Input[str] description: The description of the patch baseline.
        :param pulumi.Input[list] global_filters: A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
        :param pulumi.Input[str] name: The name of the patch baseline.
        :param pulumi.Input[str] operating_system: Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `SUSE`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
        :param pulumi.Input[list] rejected_patches: A list of rejected patches.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_patch_baseline.html.markdown.
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

            __props__['approval_rules'] = approval_rules
            __props__['approved_patches'] = approved_patches
            __props__['approved_patches_compliance_level'] = approved_patches_compliance_level
            __props__['description'] = description
            __props__['global_filters'] = global_filters
            __props__['name'] = name
            __props__['operating_system'] = operating_system
            __props__['rejected_patches'] = rejected_patches
            __props__['tags'] = tags
        super(PatchBaseline, __self__).__init__(
            'aws:ssm/patchBaseline:PatchBaseline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, approval_rules=None, approved_patches=None, approved_patches_compliance_level=None, description=None, global_filters=None, name=None, operating_system=None, rejected_patches=None, tags=None):
        """
        Get an existing PatchBaseline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] approval_rules: A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
        :param pulumi.Input[list] approved_patches: A list of explicitly approved patches for the baseline.
        :param pulumi.Input[str] approved_patches_compliance_level: Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
        :param pulumi.Input[str] description: The description of the patch baseline.
        :param pulumi.Input[list] global_filters: A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
        :param pulumi.Input[str] name: The name of the patch baseline.
        :param pulumi.Input[str] operating_system: Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `SUSE`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
        :param pulumi.Input[list] rejected_patches: A list of rejected patches.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_patch_baseline.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["approval_rules"] = approval_rules
        __props__["approved_patches"] = approved_patches
        __props__["approved_patches_compliance_level"] = approved_patches_compliance_level
        __props__["description"] = description
        __props__["global_filters"] = global_filters
        __props__["name"] = name
        __props__["operating_system"] = operating_system
        __props__["rejected_patches"] = rejected_patches
        __props__["tags"] = tags
        return PatchBaseline(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

