# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrincipalPortfolioAssociationArgs', 'PrincipalPortfolioAssociation']

@pulumi.input_type
class PrincipalPortfolioAssociationArgs:
    def __init__(__self__, *,
                 portfolio_id: pulumi.Input[str],
                 principal_arn: pulumi.Input[str],
                 accept_language: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrincipalPortfolioAssociation resource.
        :param pulumi.Input[str] portfolio_id: Portfolio identifier.
        :param pulumi.Input[str] principal_arn: Principal ARN.
        :param pulumi.Input[str] accept_language: Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        :param pulumi.Input[str] principal_type: Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        pulumi.set(__self__, "portfolio_id", portfolio_id)
        pulumi.set(__self__, "principal_arn", principal_arn)
        if accept_language is not None:
            pulumi.set(__self__, "accept_language", accept_language)
        if principal_type is not None:
            pulumi.set(__self__, "principal_type", principal_type)

    @property
    @pulumi.getter(name="portfolioId")
    def portfolio_id(self) -> pulumi.Input[str]:
        """
        Portfolio identifier.
        """
        return pulumi.get(self, "portfolio_id")

    @portfolio_id.setter
    def portfolio_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "portfolio_id", value)

    @property
    @pulumi.getter(name="principalArn")
    def principal_arn(self) -> pulumi.Input[str]:
        """
        Principal ARN.
        """
        return pulumi.get(self, "principal_arn")

    @principal_arn.setter
    def principal_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_arn", value)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[pulumi.Input[str]]:
        """
        Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        """
        return pulumi.get(self, "accept_language")

    @accept_language.setter
    def accept_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accept_language", value)

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> Optional[pulumi.Input[str]]:
        """
        Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        return pulumi.get(self, "principal_type")

    @principal_type.setter
    def principal_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_type", value)


@pulumi.input_type
class _PrincipalPortfolioAssociationState:
    def __init__(__self__, *,
                 accept_language: Optional[pulumi.Input[str]] = None,
                 portfolio_id: Optional[pulumi.Input[str]] = None,
                 principal_arn: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrincipalPortfolioAssociation resources.
        :param pulumi.Input[str] accept_language: Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        :param pulumi.Input[str] portfolio_id: Portfolio identifier.
        :param pulumi.Input[str] principal_arn: Principal ARN.
        :param pulumi.Input[str] principal_type: Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        if accept_language is not None:
            pulumi.set(__self__, "accept_language", accept_language)
        if portfolio_id is not None:
            pulumi.set(__self__, "portfolio_id", portfolio_id)
        if principal_arn is not None:
            pulumi.set(__self__, "principal_arn", principal_arn)
        if principal_type is not None:
            pulumi.set(__self__, "principal_type", principal_type)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[pulumi.Input[str]]:
        """
        Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        """
        return pulumi.get(self, "accept_language")

    @accept_language.setter
    def accept_language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accept_language", value)

    @property
    @pulumi.getter(name="portfolioId")
    def portfolio_id(self) -> Optional[pulumi.Input[str]]:
        """
        Portfolio identifier.
        """
        return pulumi.get(self, "portfolio_id")

    @portfolio_id.setter
    def portfolio_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portfolio_id", value)

    @property
    @pulumi.getter(name="principalArn")
    def principal_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Principal ARN.
        """
        return pulumi.get(self, "principal_arn")

    @principal_arn.setter
    def principal_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_arn", value)

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> Optional[pulumi.Input[str]]:
        """
        Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        return pulumi.get(self, "principal_type")

    @principal_type.setter
    def principal_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_type", value)


class PrincipalPortfolioAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_language: Optional[pulumi.Input[str]] = None,
                 portfolio_id: Optional[pulumi.Input[str]] = None,
                 principal_arn: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Service Catalog Principal Portfolio Association.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.servicecatalog.PrincipalPortfolioAssociation("example",
            portfolio_id="port-68656c6c6f",
            principal_arn="arn:aws:iam::123456789012:user/Eleanor")
        ```

        ## Import

        `aws_servicecatalog_principal_portfolio_association` can be imported using the accept language, principal ARN, and portfolio ID, separated by a comma, e.g.

        ```sh
         $ pulumi import aws:servicecatalog/principalPortfolioAssociation:PrincipalPortfolioAssociation example en,arn:aws:iam::123456789012:user/Eleanor,port-68656c6c6f
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_language: Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        :param pulumi.Input[str] portfolio_id: Portfolio identifier.
        :param pulumi.Input[str] principal_arn: Principal ARN.
        :param pulumi.Input[str] principal_type: Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: PrincipalPortfolioAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Service Catalog Principal Portfolio Association.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.servicecatalog.PrincipalPortfolioAssociation("example",
            portfolio_id="port-68656c6c6f",
            principal_arn="arn:aws:iam::123456789012:user/Eleanor")
        ```

        ## Import

        `aws_servicecatalog_principal_portfolio_association` can be imported using the accept language, principal ARN, and portfolio ID, separated by a comma, e.g.

        ```sh
         $ pulumi import aws:servicecatalog/principalPortfolioAssociation:PrincipalPortfolioAssociation example en,arn:aws:iam::123456789012:user/Eleanor,port-68656c6c6f
        ```

        :param str resource_name_: The name of the resource.
        :param PrincipalPortfolioAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrincipalPortfolioAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_language: Optional[pulumi.Input[str]] = None,
                 portfolio_id: Optional[pulumi.Input[str]] = None,
                 principal_arn: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = PrincipalPortfolioAssociationArgs.__new__(PrincipalPortfolioAssociationArgs)

            __props__.__dict__["accept_language"] = accept_language
            if portfolio_id is None and not opts.urn:
                raise TypeError("Missing required property 'portfolio_id'")
            __props__.__dict__["portfolio_id"] = portfolio_id
            if principal_arn is None and not opts.urn:
                raise TypeError("Missing required property 'principal_arn'")
            __props__.__dict__["principal_arn"] = principal_arn
            __props__.__dict__["principal_type"] = principal_type
        super(PrincipalPortfolioAssociation, __self__).__init__(
            'aws:servicecatalog/principalPortfolioAssociation:PrincipalPortfolioAssociation',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accept_language: Optional[pulumi.Input[str]] = None,
            portfolio_id: Optional[pulumi.Input[str]] = None,
            principal_arn: Optional[pulumi.Input[str]] = None,
            principal_type: Optional[pulumi.Input[str]] = None) -> 'PrincipalPortfolioAssociation':
        """
        Get an existing PrincipalPortfolioAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_language: Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        :param pulumi.Input[str] portfolio_id: Portfolio identifier.
        :param pulumi.Input[str] principal_arn: Principal ARN.
        :param pulumi.Input[str] principal_type: Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrincipalPortfolioAssociationState.__new__(_PrincipalPortfolioAssociationState)

        __props__.__dict__["accept_language"] = accept_language
        __props__.__dict__["portfolio_id"] = portfolio_id
        __props__.__dict__["principal_arn"] = principal_arn
        __props__.__dict__["principal_type"] = principal_type
        return PrincipalPortfolioAssociation(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> pulumi.Output[Optional[str]]:
        """
        Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
        """
        return pulumi.get(self, "accept_language")

    @property
    @pulumi.getter(name="portfolioId")
    def portfolio_id(self) -> pulumi.Output[str]:
        """
        Portfolio identifier.
        """
        return pulumi.get(self, "portfolio_id")

    @property
    @pulumi.getter(name="principalArn")
    def principal_arn(self) -> pulumi.Output[str]:
        """
        Principal ARN.
        """
        return pulumi.get(self, "principal_arn")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Output[Optional[str]]:
        """
        Principal type. Setting this argument empty (e.g., `principal_type = ""`) will result in an error. Valid value is `IAM`. Default is `IAM`.
        """
        return pulumi.get(self, "principal_type")

