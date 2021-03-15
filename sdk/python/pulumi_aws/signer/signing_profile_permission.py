# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SigningProfilePermissionArgs', 'SigningProfilePermission']

@pulumi.input_type
class SigningProfilePermissionArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 principal: pulumi.Input[str],
                 profile_name: pulumi.Input[str],
                 profile_version: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 statement_id_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SigningProfilePermission resource.
        :param pulumi.Input[str] action: An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
        :param pulumi.Input[str] principal: The AWS principal to be granted a cross-account permission.
        :param pulumi.Input[str] profile_name: Name of the signing profile to add the cross-account permissions.
        :param pulumi.Input[str] profile_version: The signing profile version that a permission applies to.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "principal", principal)
        pulumi.set(__self__, "profile_name", profile_name)
        if profile_version is not None:
            pulumi.set(__self__, "profile_version", profile_version)
        if statement_id is not None:
            pulumi.set(__self__, "statement_id", statement_id)
        if statement_id_prefix is not None:
            pulumi.set(__self__, "statement_id_prefix", statement_id_prefix)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Input[str]:
        """
        The AWS principal to be granted a cross-account permission.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Input[str]:
        """
        Name of the signing profile to add the cross-account permissions.
        """
        return pulumi.get(self, "profile_name")

    @profile_name.setter
    def profile_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "profile_name", value)

    @property
    @pulumi.getter(name="profileVersion")
    def profile_version(self) -> Optional[pulumi.Input[str]]:
        """
        The signing profile version that a permission applies to.
        """
        return pulumi.get(self, "profile_version")

    @profile_version.setter
    def profile_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile_version", value)

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "statement_id")

    @statement_id.setter
    def statement_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "statement_id", value)

    @property
    @pulumi.getter(name="statementIdPrefix")
    def statement_id_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "statement_id_prefix")

    @statement_id_prefix.setter
    def statement_id_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "statement_id_prefix", value)


class SigningProfilePermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SigningProfilePermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Signer Signing Profile Permission. That is, a cross-account permission for a signing profile.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        prod_sp = aws.signer.SigningProfile("prodSp",
            platform_id="AWSLambda-SHA384-ECDSA",
            name_prefix="prod_sp_",
            signature_validity_period=aws.signer.SigningProfileSignatureValidityPeriodArgs(
                value=5,
                type="YEARS",
            ),
            tags={
                "tag1": "value1",
                "tag2": "value2",
            })
        sp_permission1 = aws.signer.SigningProfilePermission("spPermission1",
            profile_name=prod_sp.name,
            action="signer:StartSigningJob",
            principal=var["aws_account"])
        sp_permission2 = aws.signer.SigningProfilePermission("spPermission2",
            profile_name=prod_sp.name,
            action="signer:GetSigningProfile",
            principal=var["aws_team_role_arn"],
            statement_id="ProdAccountStartSigningJob_StatementId")
        sp_permission3 = aws.signer.SigningProfilePermission("spPermission3",
            profile_name=prod_sp.name,
            action="signer:RevokeSignature",
            principal="123456789012",
            profile_version=prod_sp.version,
            statement_id_prefix="version-permission-")
        ```

        ## Import

        Signer signing profile permission statements can be imported using profile_name/statement_id, e.g.

        ```sh
         $ pulumi import aws:signer/signingProfilePermission:SigningProfilePermission test_signer_signing_profile_permission prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK/ProdAccountStartSigningJobStatementId
        ```

        :param str resource_name: The name of the resource.
        :param SigningProfilePermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 profile_version: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 statement_id_prefix: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Signer Signing Profile Permission. That is, a cross-account permission for a signing profile.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        prod_sp = aws.signer.SigningProfile("prodSp",
            platform_id="AWSLambda-SHA384-ECDSA",
            name_prefix="prod_sp_",
            signature_validity_period=aws.signer.SigningProfileSignatureValidityPeriodArgs(
                value=5,
                type="YEARS",
            ),
            tags={
                "tag1": "value1",
                "tag2": "value2",
            })
        sp_permission1 = aws.signer.SigningProfilePermission("spPermission1",
            profile_name=prod_sp.name,
            action="signer:StartSigningJob",
            principal=var["aws_account"])
        sp_permission2 = aws.signer.SigningProfilePermission("spPermission2",
            profile_name=prod_sp.name,
            action="signer:GetSigningProfile",
            principal=var["aws_team_role_arn"],
            statement_id="ProdAccountStartSigningJob_StatementId")
        sp_permission3 = aws.signer.SigningProfilePermission("spPermission3",
            profile_name=prod_sp.name,
            action="signer:RevokeSignature",
            principal="123456789012",
            profile_version=prod_sp.version,
            statement_id_prefix="version-permission-")
        ```

        ## Import

        Signer signing profile permission statements can be imported using profile_name/statement_id, e.g.

        ```sh
         $ pulumi import aws:signer/signingProfilePermission:SigningProfilePermission test_signer_signing_profile_permission prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK/ProdAccountStartSigningJobStatementId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
        :param pulumi.Input[str] principal: The AWS principal to be granted a cross-account permission.
        :param pulumi.Input[str] profile_name: Name of the signing profile to add the cross-account permissions.
        :param pulumi.Input[str] profile_version: The signing profile version that a permission applies to.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SigningProfilePermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 profile_version: Optional[pulumi.Input[str]] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 statement_id_prefix: Optional[pulumi.Input[str]] = None,
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

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            if principal is None and not opts.urn:
                raise TypeError("Missing required property 'principal'")
            __props__['principal'] = principal
            if profile_name is None and not opts.urn:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['profile_version'] = profile_version
            __props__['statement_id'] = statement_id
            __props__['statement_id_prefix'] = statement_id_prefix
        super(SigningProfilePermission, __self__).__init__(
            'aws:signer/signingProfilePermission:SigningProfilePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            profile_name: Optional[pulumi.Input[str]] = None,
            profile_version: Optional[pulumi.Input[str]] = None,
            statement_id: Optional[pulumi.Input[str]] = None,
            statement_id_prefix: Optional[pulumi.Input[str]] = None) -> 'SigningProfilePermission':
        """
        Get an existing SigningProfilePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
        :param pulumi.Input[str] principal: The AWS principal to be granted a cross-account permission.
        :param pulumi.Input[str] profile_name: Name of the signing profile to add the cross-account permissions.
        :param pulumi.Input[str] profile_version: The signing profile version that a permission applies to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["principal"] = principal
        __props__["profile_name"] = profile_name
        __props__["profile_version"] = profile_version
        __props__["statement_id"] = statement_id
        __props__["statement_id_prefix"] = statement_id_prefix
        return SigningProfilePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        An AWS Signer action permitted as part of cross-account permissions. Valid values: `signer:StartSigningJob`, `signer:GetSigningProfile`, or `signer:RevokeSignature`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[str]:
        """
        The AWS principal to be granted a cross-account permission.
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> pulumi.Output[str]:
        """
        Name of the signing profile to add the cross-account permissions.
        """
        return pulumi.get(self, "profile_name")

    @property
    @pulumi.getter(name="profileVersion")
    def profile_version(self) -> pulumi.Output[str]:
        """
        The signing profile version that a permission applies to.
        """
        return pulumi.get(self, "profile_version")

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "statement_id")

    @property
    @pulumi.getter(name="statementIdPrefix")
    def statement_id_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "statement_id_prefix")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

