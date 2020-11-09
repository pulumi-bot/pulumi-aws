# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PatchBaseline']


class PatchBaseline(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineApprovalRuleArgs']]]]] = None,
                 approved_patches: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 approved_patches_compliance_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineGlobalFilterArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operating_system: Optional[pulumi.Input[str]] = None,
                 rejected_patches: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an SSM Patch Baseline resource

        > **NOTE on Patch Baselines:** The `approved_patches` and `approval_rule` are
        both marked as optional fields, but the Patch Baseline requires that at least one
        of them is specified.

        ## Example Usage

        Basic usage using `approved_patches` only

        ```python
        import pulumi
        import pulumi_aws as aws

        production = aws.ssm.PatchBaseline("production", approved_patches=["KB123456"])
        ```

        Advanced usage, specifying patch filters

        ```python
        import pulumi
        import pulumi_aws as aws

        production = aws.ssm.PatchBaseline("production",
            approval_rules=[
                aws.ssm.PatchBaselineApprovalRuleArgs(
                    approve_after_days=7,
                    compliance_level="HIGH",
                    patch_filters=[
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="PRODUCT",
                            values=["WindowsServer2016"],
                        ),
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="CLASSIFICATION",
                            values=[
                                "CriticalUpdates",
                                "SecurityUpdates",
                                "Updates",
                            ],
                        ),
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="MSRC_SEVERITY",
                            values=[
                                "Critical",
                                "Important",
                                "Moderate",
                            ],
                        ),
                    ],
                ),
                aws.ssm.PatchBaselineApprovalRuleArgs(
                    approve_after_days=7,
                    patch_filters=[aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                        key="PRODUCT",
                        values=["WindowsServer2012"],
                    )],
                ),
            ],
            approved_patches=[
                "KB123456",
                "KB456789",
            ],
            description="Patch Baseline Description",
            global_filters=[
                aws.ssm.PatchBaselineGlobalFilterArgs(
                    key="PRODUCT",
                    values=["WindowsServer2008"],
                ),
                aws.ssm.PatchBaselineGlobalFilterArgs(
                    key="CLASSIFICATION",
                    values=["ServicePacks"],
                ),
                aws.ssm.PatchBaselineGlobalFilterArgs(
                    key="MSRC_SEVERITY",
                    values=["Low"],
                ),
            ],
            rejected_patches=["KB987654"])
        ```

        Advanced usage, specifying Microsoft application and Windows patch rules

        ```python
        import pulumi
        import pulumi_aws as aws

        windows_os_apps = aws.ssm.PatchBaseline("windowsOsApps",
            approval_rules=[
                aws.ssm.PatchBaselineApprovalRuleArgs(
                    approve_after_days=7,
                    patch_filters=[
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="CLASSIFICATION",
                            values=[
                                "CriticalUpdates",
                                "SecurityUpdates",
                            ],
                        ),
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="MSRC_SEVERITY",
                            values=[
                                "Critical",
                                "Important",
                            ],
                        ),
                    ],
                ),
                aws.ssm.PatchBaselineApprovalRuleArgs(
                    approve_after_days=7,
                    patch_filters=[
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="PATCH_SET",
                            values=["APPLICATION"],
                        ),
                        aws.ssm.PatchBaselineApprovalRulePatchFilterArgs(
                            key="PRODUCT",
                            values=[
                                "Office 2013",
                                "Office 2016",
                            ],
                        ),
                    ],
                ),
            ],
            description="Patch both Windows and Microsoft apps",
            operating_system="WINDOWS")
        ```

        ## Import

        SSM Patch Baselines can be imported by their baseline ID, e.g.

        ```sh
         $ pulumi import aws:ssm/patchBaseline:PatchBaseline example pb-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineApprovalRuleArgs']]]] approval_rules: A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_patches: A list of explicitly approved patches for the baseline.
        :param pulumi.Input[str] approved_patches_compliance_level: Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
        :param pulumi.Input[str] description: The description of the patch baseline.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineGlobalFilterArgs']]]] global_filters: A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
        :param pulumi.Input[str] name: The name of the patch baseline.
        :param pulumi.Input[str] operating_system: Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `SUSE`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rejected_patches: A list of rejected patches.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            approval_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineApprovalRuleArgs']]]]] = None,
            approved_patches: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            approved_patches_compliance_level: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            global_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineGlobalFilterArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            operating_system: Optional[pulumi.Input[str]] = None,
            rejected_patches: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'PatchBaseline':
        """
        Get an existing PatchBaseline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineApprovalRuleArgs']]]] approval_rules: A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] approved_patches: A list of explicitly approved patches for the baseline.
        :param pulumi.Input[str] approved_patches_compliance_level: Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
        :param pulumi.Input[str] description: The description of the patch baseline.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PatchBaselineGlobalFilterArgs']]]] global_filters: A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
        :param pulumi.Input[str] name: The name of the patch baseline.
        :param pulumi.Input[str] operating_system: Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `SUSE`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rejected_patches: A list of rejected patches.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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

    @property
    @pulumi.getter(name="approvalRules")
    def approval_rules(self) -> pulumi.Output[Optional[Sequence['outputs.PatchBaselineApprovalRule']]]:
        """
        A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
        """
        return pulumi.get(self, "approval_rules")

    @property
    @pulumi.getter(name="approvedPatches")
    def approved_patches(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of explicitly approved patches for the baseline.
        """
        return pulumi.get(self, "approved_patches")

    @property
    @pulumi.getter(name="approvedPatchesComplianceLevel")
    def approved_patches_compliance_level(self) -> pulumi.Output[Optional[str]]:
        """
        Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
        """
        return pulumi.get(self, "approved_patches_compliance_level")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the patch baseline.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalFilters")
    def global_filters(self) -> pulumi.Output[Optional[Sequence['outputs.PatchBaselineGlobalFilter']]]:
        """
        A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
        """
        return pulumi.get(self, "global_filters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the patch baseline.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operatingSystem")
    def operating_system(self) -> pulumi.Output[Optional[str]]:
        """
        Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `AMAZON_LINUX_2`, `SUSE`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
        """
        return pulumi.get(self, "operating_system")

    @property
    @pulumi.getter(name="rejectedPatches")
    def rejected_patches(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of rejected patches.
        """
        return pulumi.get(self, "rejected_patches")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

