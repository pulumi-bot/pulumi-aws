# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourcePolicyArgs', 'ResourcePolicy']

@pulumi.input_type
class ResourcePolicyArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input[str]):
        """
        The set of arguments for constructing a ResourcePolicy resource.
        :param pulumi.Input[str] policy: The policy to be applied to the aws glue data catalog.
        """
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The policy to be applied to the aws glue data catalog.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)


@pulumi.input_type
class _ResourcePolicyState:
    def __init__(__self__, *,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResourcePolicy resources.
        :param pulumi.Input[str] policy: The policy to be applied to the aws glue data catalog.
        """
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        The policy to be applied to the aws glue data catalog.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class ResourcePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Glue resource policy. Only one can exist per region.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        current_caller_identity = aws.get_caller_identity()
        current_partition = aws.get_partition()
        current_region = aws.get_region()
        glue_example_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=["glue:CreateTable"],
            resources=[f"arn:{current_partition.partition}:glue:{current_region.name}:{current_caller_identity.account_id}:*"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["*"],
                type="AWS",
            )],
        )])
        example = aws.glue.ResourcePolicy("example", policy=glue_example_policy.json)
        ```

        ## Import

        Glue Resource Policy can be imported using the account ID

        ```sh
         $ pulumi import aws:glue/resourcePolicy:ResourcePolicy Test 12356789012
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: The policy to be applied to the aws glue data catalog.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: ResourcePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Glue resource policy. Only one can exist per region.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        current_caller_identity = aws.get_caller_identity()
        current_partition = aws.get_partition()
        current_region = aws.get_region()
        glue_example_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=["glue:CreateTable"],
            resources=[f"arn:{current_partition.partition}:glue:{current_region.name}:{current_caller_identity.account_id}:*"],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["*"],
                type="AWS",
            )],
        )])
        example = aws.glue.ResourcePolicy("example", policy=glue_example_policy.json)
        ```

        ## Import

        Glue Resource Policy can be imported using the account ID

        ```sh
         $ pulumi import aws:glue/resourcePolicy:ResourcePolicy Test 12356789012
        ```

        :param str resource_name_: The name of the resource.
        :param ResourcePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourcePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourcePolicyArgs.__new__(ResourcePolicyArgs)

            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
        super(ResourcePolicy, __self__).__init__(
            'aws:glue/resourcePolicy:ResourcePolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'ResourcePolicy':
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: The policy to be applied to the aws glue data catalog.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourcePolicyState.__new__(_ResourcePolicyState)

        __props__.__dict__["policy"] = policy
        return ResourcePolicy(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The policy to be applied to the aws glue data catalog.
        """
        return pulumi.get(self, "policy")

