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

__all__ = ['WebAcl']


class WebAcl(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 visibility_config: Optional[pulumi.Input[pulumi.InputType['WebAclVisibilityConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a WAFv2 Web ACL resource.

        ## Example Usage

        This resource is based on `wafv2.RuleGroup`, check the documentation of the `wafv2.RuleGroup` resource to see examples of the various available statements.
        ### Managed Rule

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.WebAcl("example",
            default_action=aws.wafv2.WebAclDefaultActionArgs(
                allow=aws.wafv2.WebAclDefaultActionAllowArgs(),
            ),
            description="Example of a managed rule.",
            rules=[aws.wafv2.WebAclRuleArgs(
                name="rule-1",
                override_action=aws.wafv2.WebAclRuleOverrideActionArgs(
                    count=aws.wafv2.WebAclRuleOverrideActionCountArgs(),
                ),
                priority=1,
                statement=aws.wafv2.WebAclRuleStatementArgs(
                    managed_rule_group_statement=aws.wafv2.WebAclRuleStatementManagedRuleGroupStatementArgs(
                        excluded_rule=[
                            {
                                "name": "SizeRestrictions_QUERYSTRING",
                            },
                            {
                                "name": "NoUserAgent_HEADER",
                            },
                        ],
                        name="AWSManagedRulesCommonRuleSet",
                        vendor_name="AWS",
                    ),
                ),
                visibility_config={
                    "cloudwatchMetricsEnabled": False,
                    "metric_name": "friendly-rule-metric-name",
                    "sampledRequestsEnabled": False,
                },
            )],
            scope="REGIONAL",
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            },
            visibility_config=aws.wafv2.WebAclVisibilityConfigArgs(
                cloudwatch_metrics_enabled=False,
                metric_name="friendly-metric-name",
                sampled_requests_enabled=False,
            ))
        ```
        ### Rate Based

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.WebAcl("example",
            default_action=aws.wafv2.WebAclDefaultActionArgs(
                block=aws.wafv2.WebAclDefaultActionBlockArgs(),
            ),
            description="Example of a rate based statement.",
            rules=[aws.wafv2.WebAclRuleArgs(
                action=aws.wafv2.WebAclRuleActionArgs(
                    count=aws.wafv2.WebAclRuleActionCountArgs(),
                ),
                name="rule-1",
                priority=1,
                statement=aws.wafv2.WebAclRuleStatementArgs(
                    rate_based_statement=aws.wafv2.WebAclRuleStatementRateBasedStatementArgs(
                        aggregate_key_type="IP",
                        limit=10000,
                        scope_down_statement=aws.wafv2.WebAclRuleStatementRateBasedStatementScopeDownStatementArgs(
                            geo_match_statement=aws.wafv2.WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatementArgs(
                                country_codes=[
                                    "US",
                                    "NL",
                                ],
                            ),
                        ),
                    ),
                ),
                visibility_config={
                    "cloudwatchMetricsEnabled": False,
                    "metric_name": "friendly-rule-metric-name",
                    "sampledRequestsEnabled": False,
                },
            )],
            scope="REGIONAL",
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            },
            visibility_config=aws.wafv2.WebAclVisibilityConfigArgs(
                cloudwatch_metrics_enabled=False,
                metric_name="friendly-metric-name",
                sampled_requests_enabled=False,
            ))
        ```
        ### Rule Group Reference

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.RuleGroup("example",
            capacity=10,
            scope="REGIONAL",
            rules=[
                aws.wafv2.RuleGroupRuleArgs(
                    name="rule-1",
                    priority=1,
                    action=aws.wafv2.RuleGroupRuleActionArgs(
                        count=aws.wafv2.RuleGroupRuleActionCountArgs(),
                    ),
                    statement=aws.wafv2.RuleGroupRuleStatementArgs(
                        geo_match_statement=aws.wafv2.RuleGroupRuleStatementGeoMatchStatementArgs(
                            country_codes=["NL"],
                        ),
                    ),
                    visibility_config={
                        "cloudwatchMetricsEnabled": False,
                        "metric_name": "friendly-rule-metric-name",
                        "sampledRequestsEnabled": False,
                    },
                ),
                aws.wafv2.RuleGroupRuleArgs(
                    name="rule-to-exclude-a",
                    priority=10,
                    action=aws.wafv2.RuleGroupRuleActionArgs(
                        allow=aws.wafv2.RuleGroupRuleActionAllowArgs(),
                    ),
                    statement=aws.wafv2.RuleGroupRuleStatementArgs(
                        geo_match_statement=aws.wafv2.RuleGroupRuleStatementGeoMatchStatementArgs(
                            country_codes=["US"],
                        ),
                    ),
                    visibility_config={
                        "cloudwatchMetricsEnabled": False,
                        "metric_name": "friendly-rule-metric-name",
                        "sampledRequestsEnabled": False,
                    },
                ),
                aws.wafv2.RuleGroupRuleArgs(
                    name="rule-to-exclude-b",
                    priority=15,
                    action=aws.wafv2.RuleGroupRuleActionArgs(
                        allow=aws.wafv2.RuleGroupRuleActionAllowArgs(),
                    ),
                    statement=aws.wafv2.RuleGroupRuleStatementArgs(
                        geo_match_statement=aws.wafv2.RuleGroupRuleStatementGeoMatchStatementArgs(
                            country_codes=["GB"],
                        ),
                    ),
                    visibility_config={
                        "cloudwatchMetricsEnabled": False,
                        "metric_name": "friendly-rule-metric-name",
                        "sampledRequestsEnabled": False,
                    },
                ),
            ],
            visibility_config=aws.wafv2.RuleGroupVisibilityConfigArgs(
                cloudwatch_metrics_enabled=False,
                metric_name="friendly-metric-name",
                sampled_requests_enabled=False,
            ))
        test = aws.wafv2.WebAcl("test",
            scope="REGIONAL",
            default_action=aws.wafv2.WebAclDefaultActionArgs(
                block=aws.wafv2.WebAclDefaultActionBlockArgs(),
            ),
            rules=[aws.wafv2.WebAclRuleArgs(
                name="rule-1",
                priority=1,
                override_action=aws.wafv2.WebAclRuleOverrideActionArgs(
                    count=aws.wafv2.WebAclRuleOverrideActionCountArgs(),
                ),
                statement=aws.wafv2.WebAclRuleStatementArgs(
                    rule_group_reference_statement=aws.wafv2.WebAclRuleStatementRuleGroupReferenceStatementArgs(
                        arn=example.arn,
                        excluded_rules=[
                            aws.wafv2.WebAclRuleStatementRuleGroupReferenceStatementExcludedRuleArgs(
                                name="rule-to-exclude-b",
                            ),
                            aws.wafv2.WebAclRuleStatementRuleGroupReferenceStatementExcludedRuleArgs(
                                name="rule-to-exclude-a",
                            ),
                        ],
                    ),
                ),
                visibility_config={
                    "cloudwatchMetricsEnabled": False,
                    "metric_name": "friendly-rule-metric-name",
                    "sampledRequestsEnabled": False,
                },
            )],
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            },
            visibility_config=aws.wafv2.WebAclVisibilityConfigArgs(
                cloudwatch_metrics_enabled=False,
                metric_name="friendly-metric-name",
                sampled_requests_enabled=False,
            ))
        ```

        ## Import

        WAFv2 Web ACLs can be imported using `ID/Name/Scope` e.g.

        ```sh
         $ pulumi import aws:wafv2/webAcl:WebAcl example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']] default_action: The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
        :param pulumi.Input[str] description: A friendly description of the WebACL.
        :param pulumi.Input[str] name: A friendly name of the WebACL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]] rules: The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
        :param pulumi.Input[pulumi.InputType['WebAclVisibilityConfigArgs']] visibility_config: Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
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

            if default_action is None:
                raise TypeError("Missing required property 'default_action'")
            __props__['default_action'] = default_action
            __props__['description'] = description
            __props__['name'] = name
            __props__['rules'] = rules
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['tags'] = tags
            if visibility_config is None:
                raise TypeError("Missing required property 'visibility_config'")
            __props__['visibility_config'] = visibility_config
            __props__['arn'] = None
            __props__['capacity'] = None
            __props__['lock_token'] = None
        super(WebAcl, __self__).__init__(
            'aws:wafv2/webAcl:WebAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            capacity: Optional[pulumi.Input[int]] = None,
            default_action: Optional[pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lock_token: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            visibility_config: Optional[pulumi.Input[pulumi.InputType['WebAclVisibilityConfigArgs']]] = None) -> 'WebAcl':
        """
        Get an existing WebAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the IP Set that this statement references.
        :param pulumi.Input[int] capacity: The web ACL capacity units (WCUs) currently being used by this web ACL.
        :param pulumi.Input[pulumi.InputType['WebAclDefaultActionArgs']] default_action: The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
        :param pulumi.Input[str] description: A friendly description of the WebACL.
        :param pulumi.Input[str] name: A friendly name of the WebACL.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebAclRuleArgs']]]] rules: The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
        :param pulumi.Input[pulumi.InputType['WebAclVisibilityConfigArgs']] visibility_config: Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["capacity"] = capacity
        __props__["default_action"] = default_action
        __props__["description"] = description
        __props__["lock_token"] = lock_token
        __props__["name"] = name
        __props__["rules"] = rules
        __props__["scope"] = scope
        __props__["tags"] = tags
        __props__["visibility_config"] = visibility_config
        return WebAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the IP Set that this statement references.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        """
        The web ACL capacity units (WCUs) currently being used by this web ACL.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output['outputs.WebAclDefaultAction']:
        """
        The action to perform if none of the `rules` contained in the WebACL match. See Default Action below for details.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly description of the WebACL.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lockToken")
    def lock_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "lock_token")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the WebACL.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.WebAclRule']]]:
        """
        The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An array of key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="visibilityConfig")
    def visibility_config(self) -> pulumi.Output['outputs.WebAclVisibilityConfig']:
        """
        Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
        """
        return pulumi.get(self, "visibility_config")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

