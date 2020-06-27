# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Model(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) assigned by AWS to this model.
    """
    containers: pulumi.Output[list]
    """
    Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.

      * `containerHostname` (`str`) - The DNS host name for the container.
      * `environment` (`dict`) - Environment variables for the Docker container.
        A list of key value pairs.
      * `image` (`str`) - The registry path where the inference code image is stored in Amazon ECR.
      * `modelDataUrl` (`str`) - The URL for the S3 location where model artifacts are stored.
    """
    enable_network_isolation: pulumi.Output[bool]
    """
    Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
    """
    execution_role_arn: pulumi.Output[str]
    """
    A role that SageMaker can assume to access model artifacts and docker images for deployment.
    """
    name: pulumi.Output[str]
    """
    The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
    """
    primary_container: pulumi.Output[dict]
    """
    The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.

      * `containerHostname` (`str`) - The DNS host name for the container.
      * `environment` (`dict`) - Environment variables for the Docker container.
        A list of key value pairs.
      * `image` (`str`) - The registry path where the inference code image is stored in Amazon ECR.
      * `modelDataUrl` (`str`) - The URL for the S3 location where model artifacts are stored.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    vpc_config: pulumi.Output[dict]
    """
    Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.

      * `security_group_ids` (`list`)
      * `subnets` (`list`)
    """
    def __init__(__self__, resource_name, opts=None, containers=None, enable_network_isolation=None, execution_role_arn=None, name=None, primary_container=None, tags=None, vpc_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a SageMaker model resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        model = aws.sagemaker.Model("model",
            execution_role_arn=aws_iam_role["foo"]["arn"],
            primary_container={
                "image": "174872318107.dkr.ecr.us-west-2.amazonaws.com/kmeans:1",
            })
        assume_role = aws.iam.get_policy_document(statements=[{
            "actions": ["sts:AssumeRole"],
            "principals": [{
                "identifiers": ["sagemaker.amazonaws.com"],
                "type": "Service",
            }],
        }])
        role = aws.iam.Role("role", assume_role_policy=assume_role.json)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] containers: Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
        :param pulumi.Input[bool] enable_network_isolation: Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
        :param pulumi.Input[str] execution_role_arn: A role that SageMaker can assume to access model artifacts and docker images for deployment.
        :param pulumi.Input[str] name: The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[dict] primary_container: The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[dict] vpc_config: Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.

        The **containers** object supports the following:

          * `containerHostname` (`pulumi.Input[str]`) - The DNS host name for the container.
          * `environment` (`pulumi.Input[dict]`) - Environment variables for the Docker container.
            A list of key value pairs.
          * `image` (`pulumi.Input[str]`) - The registry path where the inference code image is stored in Amazon ECR.
          * `modelDataUrl` (`pulumi.Input[str]`) - The URL for the S3 location where model artifacts are stored.

        The **primary_container** object supports the following:

          * `containerHostname` (`pulumi.Input[str]`) - The DNS host name for the container.
          * `environment` (`pulumi.Input[dict]`) - Environment variables for the Docker container.
            A list of key value pairs.
          * `image` (`pulumi.Input[str]`) - The registry path where the inference code image is stored in Amazon ECR.
          * `modelDataUrl` (`pulumi.Input[str]`) - The URL for the S3 location where model artifacts are stored.

        The **vpc_config** object supports the following:

          * `security_group_ids` (`pulumi.Input[list]`)
          * `subnets` (`pulumi.Input[list]`)
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
    def get(resource_name, id, opts=None, arn=None, containers=None, enable_network_isolation=None, execution_role_arn=None, name=None, primary_container=None, tags=None, vpc_config=None):
        """
        Get an existing Model resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this model.
        :param pulumi.Input[list] containers: Specifies containers in the inference pipeline. If not specified, the `primary_container` argument is required. Fields are documented below.
        :param pulumi.Input[bool] enable_network_isolation: Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
        :param pulumi.Input[str] execution_role_arn: A role that SageMaker can assume to access model artifacts and docker images for deployment.
        :param pulumi.Input[str] name: The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[dict] primary_container: The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[dict] vpc_config: Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.

        The **containers** object supports the following:

          * `containerHostname` (`pulumi.Input[str]`) - The DNS host name for the container.
          * `environment` (`pulumi.Input[dict]`) - Environment variables for the Docker container.
            A list of key value pairs.
          * `image` (`pulumi.Input[str]`) - The registry path where the inference code image is stored in Amazon ECR.
          * `modelDataUrl` (`pulumi.Input[str]`) - The URL for the S3 location where model artifacts are stored.

        The **primary_container** object supports the following:

          * `containerHostname` (`pulumi.Input[str]`) - The DNS host name for the container.
          * `environment` (`pulumi.Input[dict]`) - Environment variables for the Docker container.
            A list of key value pairs.
          * `image` (`pulumi.Input[str]`) - The registry path where the inference code image is stored in Amazon ECR.
          * `modelDataUrl` (`pulumi.Input[str]`) - The URL for the S3 location where model artifacts are stored.

        The **vpc_config** object supports the following:

          * `security_group_ids` (`pulumi.Input[list]`)
          * `subnets` (`pulumi.Input[list]`)
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
