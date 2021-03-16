# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['LifecyclePolicyArgs', 'LifecyclePolicy']

@pulumi.input_type
class LifecyclePolicyArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 execution_role_arn: pulumi.Input[str],
                 policy_details: pulumi.Input['LifecyclePolicyPolicyDetailsArgs'],
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a LifecyclePolicy resource.
        :param pulumi.Input[str] description: A description for the DLM lifecycle policy.
        :param pulumi.Input[str] execution_role_arn: The ARN of an IAM role that is able to be assumed by the DLM service.
        :param pulumi.Input['LifecyclePolicyPolicyDetailsArgs'] policy_details: See the `policy_details` configuration block. Max of 1.
        :param pulumi.Input[str] state: Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "execution_role_arn", execution_role_arn)
        pulumi.set(__self__, "policy_details", policy_details)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        A description for the DLM lifecycle policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="executionRoleArn")
    def execution_role_arn(self) -> pulumi.Input[str]:
        """
        The ARN of an IAM role that is able to be assumed by the DLM service.
        """
        return pulumi.get(self, "execution_role_arn")

    @execution_role_arn.setter
    def execution_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "execution_role_arn", value)

    @property
    @pulumi.getter(name="policyDetails")
    def policy_details(self) -> pulumi.Input['LifecyclePolicyPolicyDetailsArgs']:
        """
        See the `policy_details` configuration block. Max of 1.
        """
        return pulumi.get(self, "policy_details")

    @policy_details.setter
    def policy_details(self, value: pulumi.Input['LifecyclePolicyPolicyDetailsArgs']):
        pulumi.set(self, "policy_details", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class LifecyclePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 policy_details: Optional[pulumi.Input[pulumi.InputType['LifecyclePolicyPolicyDetailsArgs']]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        dlm_lifecycle_role = aws.iam.Role("dlmLifecycleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "dlm.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        dlm_lifecycle = aws.iam.RolePolicy("dlmLifecycle",
            role=dlm_lifecycle_role.id,
            policy=\"\"\"{
           "Version": "2012-10-17",
           "Statement": [
              {
                 "Effect": "Allow",
                 "Action": [
                    "ec2:CreateSnapshot",
                    "ec2:CreateSnapshots",
                    "ec2:DeleteSnapshot",
                    "ec2:DescribeInstances",
                    "ec2:DescribeVolumes",
                    "ec2:DescribeSnapshots"
                 ],
                 "Resource": "*"
              },
              {
                 "Effect": "Allow",
                 "Action": [
                    "ec2:CreateTags"
                 ],
                 "Resource": "arn:aws:ec2:*::snapshot/*"
              }
           ]
        }
        \"\"\")
        example = aws.dlm.LifecyclePolicy("example",
            description="example DLM lifecycle policy",
            execution_role_arn=dlm_lifecycle_role.arn,
            state="ENABLED",
            policy_details=aws.dlm.LifecyclePolicyPolicyDetailsArgs(
                resource_types=["VOLUME"],
                schedules=[aws.dlm.LifecyclePolicyPolicyDetailsScheduleArgs(
                    name="2 weeks of daily snapshots",
                    create_rule=aws.dlm.LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs(
                        interval=24,
                        interval_unit="HOURS",
                        times=["23:45"],
                    ),
                    retain_rule=aws.dlm.LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs(
                        count=14,
                    ),
                    tags_to_add={
                        "SnapshotCreator": "DLM",
                    },
                    copy_tags=False,
                )],
                target_tags={
                    "Snapshot": "true",
                },
            ))
        ```

        ## Import

        DLM lifecyle policies can be imported by their policy ID

        ```sh
         $ pulumi import aws:dlm/lifecyclePolicy:LifecyclePolicy example policy-abcdef12345678901
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the DLM lifecycle policy.
        :param pulumi.Input[str] execution_role_arn: The ARN of an IAM role that is able to be assumed by the DLM service.
        :param pulumi.Input[pulumi.InputType['LifecyclePolicyPolicyDetailsArgs']] policy_details: See the `policy_details` configuration block. Max of 1.
        :param pulumi.Input[str] state: Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LifecyclePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        dlm_lifecycle_role = aws.iam.Role("dlmLifecycleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "dlm.amazonaws.com"
              },
              "Effect": "Allow",
              "Sid": ""
            }
          ]
        }
        \"\"\")
        dlm_lifecycle = aws.iam.RolePolicy("dlmLifecycle",
            role=dlm_lifecycle_role.id,
            policy=\"\"\"{
           "Version": "2012-10-17",
           "Statement": [
              {
                 "Effect": "Allow",
                 "Action": [
                    "ec2:CreateSnapshot",
                    "ec2:CreateSnapshots",
                    "ec2:DeleteSnapshot",
                    "ec2:DescribeInstances",
                    "ec2:DescribeVolumes",
                    "ec2:DescribeSnapshots"
                 ],
                 "Resource": "*"
              },
              {
                 "Effect": "Allow",
                 "Action": [
                    "ec2:CreateTags"
                 ],
                 "Resource": "arn:aws:ec2:*::snapshot/*"
              }
           ]
        }
        \"\"\")
        example = aws.dlm.LifecyclePolicy("example",
            description="example DLM lifecycle policy",
            execution_role_arn=dlm_lifecycle_role.arn,
            state="ENABLED",
            policy_details=aws.dlm.LifecyclePolicyPolicyDetailsArgs(
                resource_types=["VOLUME"],
                schedules=[aws.dlm.LifecyclePolicyPolicyDetailsScheduleArgs(
                    name="2 weeks of daily snapshots",
                    create_rule=aws.dlm.LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs(
                        interval=24,
                        interval_unit="HOURS",
                        times=["23:45"],
                    ),
                    retain_rule=aws.dlm.LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs(
                        count=14,
                    ),
                    tags_to_add={
                        "SnapshotCreator": "DLM",
                    },
                    copy_tags=False,
                )],
                target_tags={
                    "Snapshot": "true",
                },
            ))
        ```

        ## Import

        DLM lifecyle policies can be imported by their policy ID

        ```sh
         $ pulumi import aws:dlm/lifecyclePolicy:LifecyclePolicy example policy-abcdef12345678901
        ```

        :param str resource_name: The name of the resource.
        :param LifecyclePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LifecyclePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 policy_details: Optional[pulumi.Input[pulumi.InputType['LifecyclePolicyPolicyDetailsArgs']]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            if execution_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'execution_role_arn'")
            __props__['execution_role_arn'] = execution_role_arn
            if policy_details is None and not opts.urn:
                raise TypeError("Missing required property 'policy_details'")
            __props__['policy_details'] = policy_details
            __props__['state'] = state
            __props__['tags'] = tags
            __props__['arn'] = None
        super(LifecyclePolicy, __self__).__init__(
            'aws:dlm/lifecyclePolicy:LifecyclePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            execution_role_arn: Optional[pulumi.Input[str]] = None,
            policy_details: Optional[pulumi.Input[pulumi.InputType['LifecyclePolicyPolicyDetailsArgs']]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'LifecyclePolicy':
        """
        Get an existing LifecyclePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
        :param pulumi.Input[str] description: A description for the DLM lifecycle policy.
        :param pulumi.Input[str] execution_role_arn: The ARN of an IAM role that is able to be assumed by the DLM service.
        :param pulumi.Input[pulumi.InputType['LifecyclePolicyPolicyDetailsArgs']] policy_details: See the `policy_details` configuration block. Max of 1.
        :param pulumi.Input[str] state: Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["execution_role_arn"] = execution_role_arn
        __props__["policy_details"] = policy_details
        __props__["state"] = state
        __props__["tags"] = tags
        return LifecyclePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description for the DLM lifecycle policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="executionRoleArn")
    def execution_role_arn(self) -> pulumi.Output[str]:
        """
        The ARN of an IAM role that is able to be assumed by the DLM service.
        """
        return pulumi.get(self, "execution_role_arn")

    @property
    @pulumi.getter(name="policyDetails")
    def policy_details(self) -> pulumi.Output['outputs.LifecyclePolicyPolicyDetails']:
        """
        See the `policy_details` configuration block. Max of 1.
        """
        return pulumi.get(self, "policy_details")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

