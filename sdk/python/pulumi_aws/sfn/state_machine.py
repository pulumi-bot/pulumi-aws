# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['StateMachine']


class StateMachine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Step Function State Machine resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # ...
        sfn_state_machine = aws.sfn.StateMachine("sfnStateMachine",
            role_arn=aws_iam_role["iam_for_sfn"]["arn"],
            definition=f\"\"\"{{
          "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
          "StartAt": "HelloWorld",
          "States": {{
            "HelloWorld": {{
              "Type": "Task",
              "Resource": "{aws_lambda_function["lambda"]["arn"]}",
              "End": true
            }}
          }}
        }}
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] definition: The Amazon States Language definition of the state machine.
        :param pulumi.Input[str] name: The name of the state machine.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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

            if definition is None:
                raise TypeError("Missing required property 'definition'")
            __props__['definition'] = definition
            __props__['name'] = name
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['creation_date'] = None
            __props__['status'] = None
        super(StateMachine, __self__).__init__(
            'aws:sfn/stateMachine:StateMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            definition: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'StateMachine':
        """
        Get an existing StateMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the state machine.
        :param pulumi.Input[str] creation_date: The date the state machine was created.
        :param pulumi.Input[str] definition: The Amazon States Language definition of the state machine.
        :param pulumi.Input[str] name: The name of the state machine.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
        :param pulumi.Input[str] status: The current status of the state machine. Either "ACTIVE" or "DELETING".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["creation_date"] = creation_date
        __props__["definition"] = definition
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        __props__["status"] = status
        __props__["tags"] = tags
        return StateMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the state machine.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        The date the state machine was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[str]:
        """
        The Amazon States Language definition of the state machine.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the state machine.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The current status of the state machine. Either "ACTIVE" or "DELETING".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

