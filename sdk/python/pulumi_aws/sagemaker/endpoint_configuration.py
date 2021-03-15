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

__all__ = ['EndpointConfigurationArgs', 'EndpointConfiguration']

@pulumi.input_type
class EndpointConfigurationArgs:
    def __init__(__self__, *,
                 production_variants: pulumi.Input[Sequence[pulumi.Input['EndpointConfigurationProductionVariantArgs']]],
                 data_capture_config: Optional[pulumi.Input['EndpointConfigurationDataCaptureConfigArgs']] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a EndpointConfiguration resource.
        :param pulumi.Input[Sequence[pulumi.Input['EndpointConfigurationProductionVariantArgs']]] production_variants: Fields are documented below.
        :param pulumi.Input['EndpointConfigurationDataCaptureConfigArgs'] data_capture_config: Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
        :param pulumi.Input[str] kms_key_arn: Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        pulumi.set(__self__, "production_variants", production_variants)
        if data_capture_config is not None:
            pulumi.set(__self__, "data_capture_config", data_capture_config)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="productionVariants")
    def production_variants(self) -> pulumi.Input[Sequence[pulumi.Input['EndpointConfigurationProductionVariantArgs']]]:
        """
        Fields are documented below.
        """
        return pulumi.get(self, "production_variants")

    @production_variants.setter
    def production_variants(self, value: pulumi.Input[Sequence[pulumi.Input['EndpointConfigurationProductionVariantArgs']]]):
        pulumi.set(self, "production_variants", value)

    @property
    @pulumi.getter(name="dataCaptureConfig")
    def data_capture_config(self) -> Optional[pulumi.Input['EndpointConfigurationDataCaptureConfigArgs']]:
        """
        Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
        """
        return pulumi.get(self, "data_capture_config")

    @data_capture_config.setter
    def data_capture_config(self, value: Optional[pulumi.Input['EndpointConfigurationDataCaptureConfigArgs']]):
        pulumi.set(self, "data_capture_config", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class EndpointConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a SageMaker endpoint configuration resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        ec = aws.sagemaker.EndpointConfiguration("ec",
            production_variants=[aws.sagemaker.EndpointConfigurationProductionVariantArgs(
                variant_name="variant-1",
                model_name=aws_sagemaker_model["m"]["name"],
                initial_instance_count=1,
                instance_type="ml.t2.medium",
            )],
            tags={
                "Name": "foo",
            })
        ```

        ## Import

        Endpoint configurations can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:sagemaker/endpointConfiguration:EndpointConfiguration test_endpoint_config endpoint-config-foo
        ```

        :param str resource_name: The name of the resource.
        :param EndpointConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_capture_config: Optional[pulumi.Input[pulumi.InputType['EndpointConfigurationDataCaptureConfigArgs']]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 production_variants: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointConfigurationProductionVariantArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a SageMaker endpoint configuration resource.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        ec = aws.sagemaker.EndpointConfiguration("ec",
            production_variants=[aws.sagemaker.EndpointConfigurationProductionVariantArgs(
                variant_name="variant-1",
                model_name=aws_sagemaker_model["m"]["name"],
                initial_instance_count=1,
                instance_type="ml.t2.medium",
            )],
            tags={
                "Name": "foo",
            })
        ```

        ## Import

        Endpoint configurations can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:sagemaker/endpointConfiguration:EndpointConfiguration test_endpoint_config endpoint-config-foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['EndpointConfigurationDataCaptureConfigArgs']] data_capture_config: Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
        :param pulumi.Input[str] kms_key_arn: Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointConfigurationProductionVariantArgs']]]] production_variants: Fields are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_capture_config: Optional[pulumi.Input[pulumi.InputType['EndpointConfigurationDataCaptureConfigArgs']]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 production_variants: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointConfigurationProductionVariantArgs']]]]] = None,
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

            __props__['data_capture_config'] = data_capture_config
            __props__['kms_key_arn'] = kms_key_arn
            __props__['name'] = name
            if production_variants is None and not opts.urn:
                raise TypeError("Missing required property 'production_variants'")
            __props__['production_variants'] = production_variants
            __props__['tags'] = tags
            __props__['arn'] = None
        super(EndpointConfiguration, __self__).__init__(
            'aws:sagemaker/endpointConfiguration:EndpointConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            data_capture_config: Optional[pulumi.Input[pulumi.InputType['EndpointConfigurationDataCaptureConfigArgs']]] = None,
            kms_key_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            production_variants: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointConfigurationProductionVariantArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'EndpointConfiguration':
        """
        Get an existing EndpointConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
        :param pulumi.Input[pulumi.InputType['EndpointConfigurationDataCaptureConfigArgs']] data_capture_config: Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
        :param pulumi.Input[str] kms_key_arn: Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        :param pulumi.Input[str] name: The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointConfigurationProductionVariantArgs']]]] production_variants: Fields are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["data_capture_config"] = data_capture_config
        __props__["kms_key_arn"] = kms_key_arn
        __props__["name"] = name
        __props__["production_variants"] = production_variants
        __props__["tags"] = tags
        return EndpointConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) assigned by AWS to this endpoint configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dataCaptureConfig")
    def data_capture_config(self) -> pulumi.Output[Optional['outputs.EndpointConfigurationDataCaptureConfig']]:
        """
        Specifies the parameters to capture input/output of Sagemaker models endpoints. Fields are documented below.
        """
        return pulumi.get(self, "data_capture_config")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the endpoint configuration. If omitted, this provider will assign a random, unique name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="productionVariants")
    def production_variants(self) -> pulumi.Output[Sequence['outputs.EndpointConfigurationProductionVariant']]:
        """
        Fields are documented below.
        """
        return pulumi.get(self, "production_variants")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

