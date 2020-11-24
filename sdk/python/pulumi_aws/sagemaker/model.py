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

__all__ = ['Model']


class Model(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 containers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelContainerArgs']]]]] = None,
                 enable_network_isolation: Optional[pulumi.Input[bool]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 primary_container: Optional[pulumi.Input[pulumi.InputType['ModelPrimaryContainerArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_config: Optional[pulumi.Input[pulumi.InputType['ModelVpcConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a SageMaker model resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        assume_role = aws.iam.get_policy_document(statements=[{
            "actions": ["sts:AssumeRole"],
            "principals": [{
                "type": "Service",
                "identifiers": ["sagemaker.amazonaws.com"],
            }],
        }])
        example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
        example_model = aws.sagemaker.Model("exampleModel",
            execution_role_arn=example_role.arn,
            primary_container={
                "image": "174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1",
            })
        ```

        ## Import

        Models can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:sagemaker/model:Model test_model model-foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelContainerArgs']]]] containers: Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
        :param pulumi.Input[bool] enable_network_isolation: Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
        :param pulumi.Input[str] execution_role_arn: A role that SageMaker can assume to access model artifacts and docker images for deployment.
        :param pulumi.Input[str] name: The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[pulumi.InputType['ModelPrimaryContainerArgs']] primary_container: The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['ModelVpcConfigArgs']] vpc_config: Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
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

            __props__['containers'] = containers
            __props__['enable_network_isolation'] = enable_network_isolation
            if execution_role_arn is None:
                raise TypeError("Missing required property 'execution_role_arn'")
            __props__['execution_role_arn'] = execution_role_arn
            __props__['name'] = name
            __props__['primary_container'] = primary_container
            __props__['tags'] = tags
            __props__['vpc_config'] = vpc_config
            __props__['arn'] = None
        super(Model, __self__).__init__(
            'aws:sagemaker/model:Model',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            containers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelContainerArgs']]]]] = None,
            enable_network_isolation: Optional[pulumi.Input[bool]] = None,
            execution_role_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            primary_container: Optional[pulumi.Input[pulumi.InputType['ModelPrimaryContainerArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_config: Optional[pulumi.Input[pulumi.InputType['ModelVpcConfigArgs']]] = None) -> 'Model':
        """
        Get an existing Model resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this model.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelContainerArgs']]]] containers: Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
        :param pulumi.Input[bool] enable_network_isolation: Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
        :param pulumi.Input[str] execution_role_arn: A role that SageMaker can assume to access model artifacts and docker images for deployment.
        :param pulumi.Input[str] name: The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[pulumi.InputType['ModelPrimaryContainerArgs']] primary_container: The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['ModelVpcConfigArgs']] vpc_config: Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["containers"] = containers
        __props__["enable_network_isolation"] = enable_network_isolation
        __props__["execution_role_arn"] = execution_role_arn
        __props__["name"] = name
        __props__["primary_container"] = primary_container
        __props__["tags"] = tags
        __props__["vpc_config"] = vpc_config
        return Model(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this model.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def containers(self) -> pulumi.Output[Optional[Sequence['outputs.ModelContainer']]]:
        """
        Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
        """
        return pulumi.get(self, "containers")

    @property
    @pulumi.getter(name="enableNetworkIsolation")
    def enable_network_isolation(self) -> pulumi.Output[Optional[bool]]:
        """
        Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
        """
        return pulumi.get(self, "enable_network_isolation")

    @property
    @pulumi.getter(name="executionRoleArn")
    def execution_role_arn(self) -> pulumi.Output[str]:
        """
        A role that SageMaker can assume to access model artifacts and docker images for deployment.
        """
        return pulumi.get(self, "execution_role_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryContainer")
    def primary_container(self) -> pulumi.Output[Optional['outputs.ModelPrimaryContainer']]:
        """
        The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
        """
        return pulumi.get(self, "primary_container")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> pulumi.Output[Optional['outputs.ModelVpcConfig']]:
        """
        Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
        """
        return pulumi.get(self, "vpc_config")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

