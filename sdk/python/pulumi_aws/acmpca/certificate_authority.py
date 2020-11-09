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

__all__ = ['CertificateAuthority']


class CertificateAuthority(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_authority_configuration: Optional[pulumi.Input[pulumi.InputType['CertificateAuthorityCertificateAuthorityConfigurationArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 permanent_deletion_time_in_days: Optional[pulumi.Input[int]] = None,
                 revocation_configuration: Optional[pulumi.Input[pulumi.InputType['CertificateAuthorityRevocationConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).

        > **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificate_signing_request` attribute and import the signed certificate using the AWS SDK, CLI or Console. This provider can support another resource to manage that workflow automatically in the future.

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.acmpca.CertificateAuthority("example",
            certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
                    common_name="example.com",
                ),
            ),
            permanent_deletion_time_in_days=7)
        ```
        ### Enable Certificate Revocation List

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        acmpca_bucket_access = pulumi.Output.all(example_bucket.arn, example_bucket.arn).apply(lambda exampleBucketArn, exampleBucketArn1: aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
            actions=[
                "s3:GetBucketAcl",
                "s3:GetBucketLocation",
                "s3:PutObject",
                "s3:PutObjectAcl",
            ],
            resources=[
                example_bucket_arn,
                f"{example_bucket_arn1}/*",
            ],
            principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
                identifiers=["acm-pca.amazonaws.com"],
                type="Service",
            )],
        )]))
        example_bucket_policy = aws.s3.BucketPolicy("exampleBucketPolicy",
            bucket=example_bucket.id,
            policy=acmpca_bucket_access.json)
        example_certificate_authority = aws.acmpca.CertificateAuthority("exampleCertificateAuthority",
            certificate_authority_configuration=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationArgs(
                key_algorithm="RSA_4096",
                signing_algorithm="SHA512WITHRSA",
                subject=aws.acmpca.CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs(
                    common_name="example.com",
                ),
            ),
            revocation_configuration=aws.acmpca.CertificateAuthorityRevocationConfigurationArgs(
                crl_configuration=aws.acmpca.CertificateAuthorityRevocationConfigurationCrlConfigurationArgs(
                    custom_cname="crl.example.com",
                    enabled=True,
                    expiration_in_days=7,
                    s3_bucket_name=example_bucket.id,
                ),
            ),
            opts=ResourceOptions(depends_on=[example_bucket_policy]))
        ```

        ## Import

        `aws_acmpca_certificate_authority` can be imported by using the certificate authority Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:acmpca/certificateAuthority:CertificateAuthority example arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CertificateAuthorityCertificateAuthorityConfigurationArgs']] certificate_authority_configuration: Nested argument containing algorithms and certificate subject information. Defined below.
        :param pulumi.Input[bool] enabled: Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        :param pulumi.Input[int] permanent_deletion_time_in_days: The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        :param pulumi.Input[pulumi.InputType['CertificateAuthorityRevocationConfigurationArgs']] revocation_configuration: Nested argument containing revocation configuration. Defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        :param pulumi.Input[str] type: The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
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

            if certificate_authority_configuration is None:
                raise TypeError("Missing required property 'certificate_authority_configuration'")
            __props__['certificate_authority_configuration'] = certificate_authority_configuration
            __props__['enabled'] = enabled
            __props__['permanent_deletion_time_in_days'] = permanent_deletion_time_in_days
            __props__['revocation_configuration'] = revocation_configuration
            __props__['tags'] = tags
            __props__['type'] = type
            __props__['arn'] = None
            __props__['certificate'] = None
            __props__['certificate_chain'] = None
            __props__['certificate_signing_request'] = None
            __props__['not_after'] = None
            __props__['not_before'] = None
            __props__['serial'] = None
            __props__['status'] = None
        super(CertificateAuthority, __self__).__init__(
            'aws:acmpca/certificateAuthority:CertificateAuthority',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            certificate_authority_configuration: Optional[pulumi.Input[pulumi.InputType['CertificateAuthorityCertificateAuthorityConfigurationArgs']]] = None,
            certificate_chain: Optional[pulumi.Input[str]] = None,
            certificate_signing_request: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            not_after: Optional[pulumi.Input[str]] = None,
            not_before: Optional[pulumi.Input[str]] = None,
            permanent_deletion_time_in_days: Optional[pulumi.Input[int]] = None,
            revocation_configuration: Optional[pulumi.Input[pulumi.InputType['CertificateAuthorityRevocationConfigurationArgs']]] = None,
            serial: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'CertificateAuthority':
        """
        Get an existing CertificateAuthority resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the certificate authority.
        :param pulumi.Input[str] certificate: Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[pulumi.InputType['CertificateAuthorityCertificateAuthorityConfigurationArgs']] certificate_authority_configuration: Nested argument containing algorithms and certificate subject information. Defined below.
        :param pulumi.Input[str] certificate_chain: Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[str] certificate_signing_request: The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        :param pulumi.Input[bool] enabled: Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        :param pulumi.Input[str] not_after: Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[str] not_before: Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[int] permanent_deletion_time_in_days: The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        :param pulumi.Input[pulumi.InputType['CertificateAuthorityRevocationConfigurationArgs']] revocation_configuration: Nested argument containing revocation configuration. Defined below.
        :param pulumi.Input[str] serial: Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[str] status: Status of the certificate authority.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        :param pulumi.Input[str] type: The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["certificate"] = certificate
        __props__["certificate_authority_configuration"] = certificate_authority_configuration
        __props__["certificate_chain"] = certificate_chain
        __props__["certificate_signing_request"] = certificate_signing_request
        __props__["enabled"] = enabled
        __props__["not_after"] = not_after
        __props__["not_before"] = not_before
        __props__["permanent_deletion_time_in_days"] = permanent_deletion_time_in_days
        __props__["revocation_configuration"] = revocation_configuration
        __props__["serial"] = serial
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["type"] = type
        return CertificateAuthority(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the certificate authority.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateAuthorityConfiguration")
    def certificate_authority_configuration(self) -> pulumi.Output['outputs.CertificateAuthorityCertificateAuthorityConfiguration']:
        """
        Nested argument containing algorithms and certificate subject information. Defined below.
        """
        return pulumi.get(self, "certificate_authority_configuration")

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> pulumi.Output[str]:
        """
        Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        """
        return pulumi.get(self, "certificate_chain")

    @property
    @pulumi.getter(name="certificateSigningRequest")
    def certificate_signing_request(self) -> pulumi.Output[str]:
        """
        The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        """
        return pulumi.get(self, "certificate_signing_request")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> pulumi.Output[str]:
        """
        Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        """
        return pulumi.get(self, "not_after")

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> pulumi.Output[str]:
        """
        Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        """
        return pulumi.get(self, "not_before")

    @property
    @pulumi.getter(name="permanentDeletionTimeInDays")
    def permanent_deletion_time_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        """
        return pulumi.get(self, "permanent_deletion_time_in_days")

    @property
    @pulumi.getter(name="revocationConfiguration")
    def revocation_configuration(self) -> pulumi.Output[Optional['outputs.CertificateAuthorityRevocationConfiguration']]:
        """
        Nested argument containing revocation configuration. Defined below.
        """
        return pulumi.get(self, "revocation_configuration")

    @property
    @pulumi.getter
    def serial(self) -> pulumi.Output[str]:
        """
        Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        """
        return pulumi.get(self, "serial")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the certificate authority.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

