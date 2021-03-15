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

__all__ = ['InfrastructureConfigurationArgs', 'InfrastructureConfiguration']

@pulumi.input_type
class InfrastructureConfigurationArgs:
    def __init__(__self__, *,
                 instance_profile_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 instance_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key_pair: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input['InfrastructureConfigurationLoggingArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 terminate_instance_on_failure: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a InfrastructureConfiguration resource.
        :param pulumi.Input[str] instance_profile_name: Name of IAM Instance Profile.
        :param pulumi.Input[str] description: Description for the configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_types: Set of EC2 Instance Types.
        :param pulumi.Input[str] key_pair: Name of EC2 Key Pair.
        :param pulumi.Input['InfrastructureConfigurationLoggingArgs'] logging: Configuration block with logging settings. Detailed below.
        :param pulumi.Input[str] name: Name for the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] resource_tags: Key-value map of resource tags to assign to infrastructure created by the configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Set of EC2 Security Group identifiers.
        :param pulumi.Input[str] sns_topic_arn: Amazon Resource Name (ARN) of SNS Topic.
        :param pulumi.Input[str] subnet_id: EC2 Subnet identifier. Also requires `security_group_ids` argument.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags to assign to the configuration.
        :param pulumi.Input[bool] terminate_instance_on_failure: Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
        """
        pulumi.set(__self__, "instance_profile_name", instance_profile_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_types is not None:
            pulumi.set(__self__, "instance_types", instance_types)
        if key_pair is not None:
            pulumi.set(__self__, "key_pair", key_pair)
        if logging is not None:
            pulumi.set(__self__, "logging", logging)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource_tags is not None:
            pulumi.set(__self__, "resource_tags", resource_tags)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if sns_topic_arn is not None:
            pulumi.set(__self__, "sns_topic_arn", sns_topic_arn)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if terminate_instance_on_failure is not None:
            pulumi.set(__self__, "terminate_instance_on_failure", terminate_instance_on_failure)

    @property
    @pulumi.getter(name="instanceProfileName")
    def instance_profile_name(self) -> pulumi.Input[str]:
        """
        Name of IAM Instance Profile.
        """
        return pulumi.get(self, "instance_profile_name")

    @instance_profile_name.setter
    def instance_profile_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_profile_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the configuration.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceTypes")
    def instance_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of EC2 Instance Types.
        """
        return pulumi.get(self, "instance_types")

    @instance_types.setter
    def instance_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "instance_types", value)

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> Optional[pulumi.Input[str]]:
        """
        Name of EC2 Key Pair.
        """
        return pulumi.get(self, "key_pair")

    @key_pair.setter
    def key_pair(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_pair", value)

    @property
    @pulumi.getter
    def logging(self) -> Optional[pulumi.Input['InfrastructureConfigurationLoggingArgs']]:
        """
        Configuration block with logging settings. Detailed below.
        """
        return pulumi.get(self, "logging")

    @logging.setter
    def logging(self, value: Optional[pulumi.Input['InfrastructureConfigurationLoggingArgs']]):
        pulumi.set(self, "logging", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for the configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags to assign to infrastructure created by the configuration.
        """
        return pulumi.get(self, "resource_tags")

    @resource_tags.setter
    def resource_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "resource_tags", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of EC2 Security Group identifiers.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of SNS Topic.
        """
        return pulumi.get(self, "sns_topic_arn")

    @sns_topic_arn.setter
    def sns_topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_topic_arn", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        EC2 Subnet identifier. Also requires `security_group_ids` argument.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags to assign to the configuration.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="terminateInstanceOnFailure")
    def terminate_instance_on_failure(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
        """
        return pulumi.get(self, "terminate_instance_on_failure")

    @terminate_instance_on_failure.setter
    def terminate_instance_on_failure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "terminate_instance_on_failure", value)


class InfrastructureConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InfrastructureConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Image Builder Infrastructure Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.imagebuilder.InfrastructureConfiguration("example",
            description="example description",
            instance_profile_name=aws_iam_instance_profile["example"]["name"],
            instance_types=[
                "t2.nano",
                "t3.micro",
            ],
            key_pair=aws_key_pair["example"]["key_name"],
            security_group_ids=[aws_security_group["example"]["id"]],
            sns_topic_arn=aws_sns_topic["example"]["arn"],
            subnet_id=aws_subnet["main"]["id"],
            terminate_instance_on_failure=True,
            logging=aws.imagebuilder.InfrastructureConfigurationLoggingArgs(
                s3_logs=aws.imagebuilder.InfrastructureConfigurationLoggingS3LogsArgs(
                    s3_bucket_name=aws_s3_bucket["example"]["bucket"],
                    s3_key_prefix="logs",
                ),
            ),
            tags={
                "foo": "bar",
            })
        ```

        ## Import

        `aws_imagebuilder_infrastructure_configuration` can be imported using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-component/example
        ```

        :param str resource_name: The name of the resource.
        :param InfrastructureConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_profile_name: Optional[pulumi.Input[str]] = None,
                 instance_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key_pair: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[pulumi.InputType['InfrastructureConfigurationLoggingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 terminate_instance_on_failure: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an Image Builder Infrastructure Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.imagebuilder.InfrastructureConfiguration("example",
            description="example description",
            instance_profile_name=aws_iam_instance_profile["example"]["name"],
            instance_types=[
                "t2.nano",
                "t3.micro",
            ],
            key_pair=aws_key_pair["example"]["key_name"],
            security_group_ids=[aws_security_group["example"]["id"]],
            sns_topic_arn=aws_sns_topic["example"]["arn"],
            subnet_id=aws_subnet["main"]["id"],
            terminate_instance_on_failure=True,
            logging=aws.imagebuilder.InfrastructureConfigurationLoggingArgs(
                s3_logs=aws.imagebuilder.InfrastructureConfigurationLoggingS3LogsArgs(
                    s3_bucket_name=aws_s3_bucket["example"]["bucket"],
                    s3_key_prefix="logs",
                ),
            ),
            tags={
                "foo": "bar",
            })
        ```

        ## Import

        `aws_imagebuilder_infrastructure_configuration` can be imported using the Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-component/example
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the configuration.
        :param pulumi.Input[str] instance_profile_name: Name of IAM Instance Profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_types: Set of EC2 Instance Types.
        :param pulumi.Input[str] key_pair: Name of EC2 Key Pair.
        :param pulumi.Input[pulumi.InputType['InfrastructureConfigurationLoggingArgs']] logging: Configuration block with logging settings. Detailed below.
        :param pulumi.Input[str] name: Name for the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] resource_tags: Key-value map of resource tags to assign to infrastructure created by the configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Set of EC2 Security Group identifiers.
        :param pulumi.Input[str] sns_topic_arn: Amazon Resource Name (ARN) of SNS Topic.
        :param pulumi.Input[str] subnet_id: EC2 Subnet identifier. Also requires `security_group_ids` argument.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags to assign to the configuration.
        :param pulumi.Input[bool] terminate_instance_on_failure: Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InfrastructureConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_profile_name: Optional[pulumi.Input[str]] = None,
                 instance_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key_pair: Optional[pulumi.Input[str]] = None,
                 logging: Optional[pulumi.Input[pulumi.InputType['InfrastructureConfigurationLoggingArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sns_topic_arn: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 terminate_instance_on_failure: Optional[pulumi.Input[bool]] = None,
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

            __props__['description'] = description
            if instance_profile_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_profile_name'")
            __props__['instance_profile_name'] = instance_profile_name
            __props__['instance_types'] = instance_types
            __props__['key_pair'] = key_pair
            __props__['logging'] = logging
            __props__['name'] = name
            __props__['resource_tags'] = resource_tags
            __props__['security_group_ids'] = security_group_ids
            __props__['sns_topic_arn'] = sns_topic_arn
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            __props__['terminate_instance_on_failure'] = terminate_instance_on_failure
            __props__['arn'] = None
            __props__['date_created'] = None
            __props__['date_updated'] = None
        super(InfrastructureConfiguration, __self__).__init__(
            'aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            date_updated: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_profile_name: Optional[pulumi.Input[str]] = None,
            instance_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            key_pair: Optional[pulumi.Input[str]] = None,
            logging: Optional[pulumi.Input[pulumi.InputType['InfrastructureConfigurationLoggingArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            sns_topic_arn: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            terminate_instance_on_failure: Optional[pulumi.Input[bool]] = None) -> 'InfrastructureConfiguration':
        """
        Get an existing InfrastructureConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the configuration.
        :param pulumi.Input[str] date_created: Date when the configuration was created.
        :param pulumi.Input[str] date_updated: Date when the configuration was updated.
        :param pulumi.Input[str] description: Description for the configuration.
        :param pulumi.Input[str] instance_profile_name: Name of IAM Instance Profile.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] instance_types: Set of EC2 Instance Types.
        :param pulumi.Input[str] key_pair: Name of EC2 Key Pair.
        :param pulumi.Input[pulumi.InputType['InfrastructureConfigurationLoggingArgs']] logging: Configuration block with logging settings. Detailed below.
        :param pulumi.Input[str] name: Name for the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] resource_tags: Key-value map of resource tags to assign to infrastructure created by the configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Set of EC2 Security Group identifiers.
        :param pulumi.Input[str] sns_topic_arn: Amazon Resource Name (ARN) of SNS Topic.
        :param pulumi.Input[str] subnet_id: EC2 Subnet identifier. Also requires `security_group_ids` argument.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags to assign to the configuration.
        :param pulumi.Input[bool] terminate_instance_on_failure: Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["date_created"] = date_created
        __props__["date_updated"] = date_updated
        __props__["description"] = description
        __props__["instance_profile_name"] = instance_profile_name
        __props__["instance_types"] = instance_types
        __props__["key_pair"] = key_pair
        __props__["logging"] = logging
        __props__["name"] = name
        __props__["resource_tags"] = resource_tags
        __props__["security_group_ids"] = security_group_ids
        __props__["sns_topic_arn"] = sns_topic_arn
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["terminate_instance_on_failure"] = terminate_instance_on_failure
        return InfrastructureConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        Date when the configuration was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="dateUpdated")
    def date_updated(self) -> pulumi.Output[str]:
        """
        Date when the configuration was updated.
        """
        return pulumi.get(self, "date_updated")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description for the configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceProfileName")
    def instance_profile_name(self) -> pulumi.Output[str]:
        """
        Name of IAM Instance Profile.
        """
        return pulumi.get(self, "instance_profile_name")

    @property
    @pulumi.getter(name="instanceTypes")
    def instance_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Set of EC2 Instance Types.
        """
        return pulumi.get(self, "instance_types")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> pulumi.Output[Optional[str]]:
        """
        Name of EC2 Key Pair.
        """
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def logging(self) -> pulumi.Output[Optional['outputs.InfrastructureConfigurationLogging']]:
        """
        Configuration block with logging settings. Detailed below.
        """
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name for the configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceTags")
    def resource_tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags to assign to infrastructure created by the configuration.
        """
        return pulumi.get(self, "resource_tags")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Set of EC2 Security Group identifiers.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="snsTopicArn")
    def sns_topic_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Amazon Resource Name (ARN) of SNS Topic.
        """
        return pulumi.get(self, "sns_topic_arn")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        EC2 Subnet identifier. Also requires `security_group_ids` argument.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags to assign to the configuration.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="terminateInstanceOnFailure")
    def terminate_instance_on_failure(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
        """
        return pulumi.get(self, "terminate_instance_on_failure")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

