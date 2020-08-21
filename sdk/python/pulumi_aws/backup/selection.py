# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Selection']


class Selection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 selection_tags: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SelectionSelectionTagArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages selection conditions for AWS Backup plan resources.

        ## Example Usage
        ### IAM Role

        > For more information about creating and managing IAM Roles for backups and restores, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/iam-service-roles.html).

        The below example creates an IAM role with the default managed IAM Policy for allowing AWS Backup to create backups.

        ```python
        import pulumi
        import pulumi_aws as aws

        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": ["sts:AssumeRole"],
              "Effect": "allow",
              "Principal": {
                "Service": ["backup.amazonaws.com"]
              }
            }
          ]
        }
        \"\"\")
        example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
            role=example_role.name)
        # ... other configuration ...
        example_selection = aws.backup.Selection("exampleSelection", iam_role_arn=example_role.arn)
        ```
        ### Selecting Backups By Tag

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.backup.Selection("example",
            iam_role_arn=aws_iam_role["example"]["arn"],
            plan_id=aws_backup_plan["example"]["id"],
            selection_tags=[{
                "type": "STRINGEQUALS",
                "key": "foo",
                "value": "bar",
            }])
        ```
        ### Selecting Backups By Resource

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.backup.Selection("example",
            iam_role_arn=aws_iam_role["example"]["arn"],
            plan_id=aws_backup_plan["example"]["id"],
            resources=[
                aws_db_instance["example"]["arn"],
                aws_ebs_volume["example"]["arn"],
                aws_efs_file_system["example"]["arn"],
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] iam_role_arn: The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
        :param pulumi.Input[str] name: The display name of a resource selection document.
        :param pulumi.Input[str] plan_id: The backup plan ID to be associated with the selection of resources.
        :param pulumi.Input[List[pulumi.Input[str]]] resources: An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SelectionSelectionTagArgs']]]] selection_tags: Tag-based conditions used to specify a set of resources to assign to a backup plan.
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

            if iam_role_arn is None:
                raise TypeError("Missing required property 'iam_role_arn'")
            __props__['iam_role_arn'] = iam_role_arn
            __props__['name'] = name
            if plan_id is None:
                raise TypeError("Missing required property 'plan_id'")
            __props__['plan_id'] = plan_id
            __props__['resources'] = resources
            __props__['selection_tags'] = selection_tags
        super(Selection, __self__).__init__(
            'aws:backup/selection:Selection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            iam_role_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            plan_id: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            selection_tags: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SelectionSelectionTagArgs']]]]] = None) -> 'Selection':
        """
        Get an existing Selection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] iam_role_arn: The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
        :param pulumi.Input[str] name: The display name of a resource selection document.
        :param pulumi.Input[str] plan_id: The backup plan ID to be associated with the selection of resources.
        :param pulumi.Input[List[pulumi.Input[str]]] resources: An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SelectionSelectionTagArgs']]]] selection_tags: Tag-based conditions used to specify a set of resources to assign to a backup plan.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["iam_role_arn"] = iam_role_arn
        __props__["name"] = name
        __props__["plan_id"] = plan_id
        __props__["resources"] = resources
        __props__["selection_tags"] = selection_tags
        return Selection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> str:
        """
        The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
        """
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The display name of a resource selection document.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> str:
        """
        The backup plan ID to be associated with the selection of resources.
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter
    def resources(self) -> Optional[List[str]]:
        """
        An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan..
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="selectionTags")
    def selection_tags(self) -> Optional[List['outputs.SelectionSelectionTag']]:
        """
        Tag-based conditions used to specify a set of resources to assign to a backup plan.
        """
        return pulumi.get(self, "selection_tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

