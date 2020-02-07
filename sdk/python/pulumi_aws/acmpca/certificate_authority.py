# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CertificateAuthority(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the certificate authority.
    """
    certificate: pulumi.Output[str]
    """
    Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
    """
    certificate_authority_configuration: pulumi.Output[dict]
    """
    Nested argument containing algorithms and certificate subject information. Defined below.

      * `keyAlgorithm` (`str`) - Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
      * `signingAlgorithm` (`str`) - Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
      * `subject` (`dict`) - Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
        * `commonName` (`str`) - Fully qualified domain name (FQDN) associated with the certificate subject.
        * `country` (`str`) - Two digit code that specifies the country in which the certificate subject located.
        * `distinguishedNameQualifier` (`str`) - Disambiguating information for the certificate subject.
        * `generationQualifier` (`str`) - Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
        * `givenName` (`str`) - First name.
        * `initials` (`str`) - Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`.
        * `locality` (`str`) - The locality (such as a city or town) in which the certificate subject is located.
        * `organization` (`str`) - Legal name of the organization with which the certificate subject is affiliated.
        * `organizationalUnit` (`str`) - A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
        * `pseudonym` (`str`) - Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
        * `state` (`str`) - State in which the subject of the certificate is located.
        * `surname` (`str`) - Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
        * `title` (`str`) - A title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject.
    """
    certificate_chain: pulumi.Output[str]
    """
    Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
    """
    certificate_signing_request: pulumi.Output[str]
    """
    The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
    """
    enabled: pulumi.Output[bool]
    """
    Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
    """
    not_after: pulumi.Output[str]
    """
    Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
    """
    not_before: pulumi.Output[str]
    """
    Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
    """
    permanent_deletion_time_in_days: pulumi.Output[float]
    """
    The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
    """
    revocation_configuration: pulumi.Output[dict]
    """
    Nested argument containing revocation configuration. Defined below.

      * `crlConfiguration` (`dict`) - Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
        * `customCname` (`str`) - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public.
        * `enabled` (`bool`) - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        * `expirationInDays` (`float`) - Number of days until a certificate expires. Must be between 1 and 5000.
        * `s3_bucket_name` (`str`) - Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket.
    """
    serial: pulumi.Output[str]
    """
    Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
    """
    status: pulumi.Output[str]
    """
    Status of the certificate authority.
    """
    tags: pulumi.Output[dict]
    """
    Specifies a key-value map of user-defined tags that are attached to the certificate authority.
    """
    type: pulumi.Output[str]
    """
    The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
    """
    def __init__(__self__, resource_name, opts=None, certificate_authority_configuration=None, enabled=None, permanent_deletion_time_in_days=None, revocation_configuration=None, tags=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).

        > **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificate_signing_request` attribute and import the signed certificate using the AWS SDK, CLI or Console. This provider can support another resource to manage that workflow automatically in the future.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acmpca_certificate_authority.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] certificate_authority_configuration: Nested argument containing algorithms and certificate subject information. Defined below.
        :param pulumi.Input[bool] enabled: Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        :param pulumi.Input[float] permanent_deletion_time_in_days: The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        :param pulumi.Input[dict] revocation_configuration: Nested argument containing revocation configuration. Defined below.
        :param pulumi.Input[dict] tags: Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        :param pulumi.Input[str] type: The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.

        The **certificate_authority_configuration** object supports the following:

          * `keyAlgorithm` (`pulumi.Input[str]`) - Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
          * `signingAlgorithm` (`pulumi.Input[str]`) - Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
          * `subject` (`pulumi.Input[dict]`) - Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
            * `commonName` (`pulumi.Input[str]`) - Fully qualified domain name (FQDN) associated with the certificate subject.
            * `country` (`pulumi.Input[str]`) - Two digit code that specifies the country in which the certificate subject located.
            * `distinguishedNameQualifier` (`pulumi.Input[str]`) - Disambiguating information for the certificate subject.
            * `generationQualifier` (`pulumi.Input[str]`) - Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
            * `givenName` (`pulumi.Input[str]`) - First name.
            * `initials` (`pulumi.Input[str]`) - Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`.
            * `locality` (`pulumi.Input[str]`) - The locality (such as a city or town) in which the certificate subject is located.
            * `organization` (`pulumi.Input[str]`) - Legal name of the organization with which the certificate subject is affiliated.
            * `organizationalUnit` (`pulumi.Input[str]`) - A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
            * `pseudonym` (`pulumi.Input[str]`) - Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
            * `state` (`pulumi.Input[str]`) - State in which the subject of the certificate is located.
            * `surname` (`pulumi.Input[str]`) - Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
            * `title` (`pulumi.Input[str]`) - A title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject.

        The **revocation_configuration** object supports the following:

          * `crlConfiguration` (`pulumi.Input[dict]`) - Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
            * `customCname` (`pulumi.Input[str]`) - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public.
            * `enabled` (`pulumi.Input[bool]`) - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
            * `expirationInDays` (`pulumi.Input[float]`) - Number of days until a certificate expires. Must be between 1 and 5000.
            * `s3_bucket_name` (`pulumi.Input[str]`) - Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket.
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
    def get(resource_name, id, opts=None, arn=None, certificate=None, certificate_authority_configuration=None, certificate_chain=None, certificate_signing_request=None, enabled=None, not_after=None, not_before=None, permanent_deletion_time_in_days=None, revocation_configuration=None, serial=None, status=None, tags=None, type=None):
        """
        Get an existing CertificateAuthority resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the certificate authority.
        :param pulumi.Input[str] certificate: Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[dict] certificate_authority_configuration: Nested argument containing algorithms and certificate subject information. Defined below.
        :param pulumi.Input[str] certificate_chain: Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[str] certificate_signing_request: The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        :param pulumi.Input[bool] enabled: Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        :param pulumi.Input[str] not_after: Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[str] not_before: Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[float] permanent_deletion_time_in_days: The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        :param pulumi.Input[dict] revocation_configuration: Nested argument containing revocation configuration. Defined below.
        :param pulumi.Input[str] serial: Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        :param pulumi.Input[str] status: Status of the certificate authority.
        :param pulumi.Input[dict] tags: Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        :param pulumi.Input[str] type: The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.

        The **certificate_authority_configuration** object supports the following:

          * `keyAlgorithm` (`pulumi.Input[str]`) - Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
          * `signingAlgorithm` (`pulumi.Input[str]`) - Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
          * `subject` (`pulumi.Input[dict]`) - Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
            * `commonName` (`pulumi.Input[str]`) - Fully qualified domain name (FQDN) associated with the certificate subject.
            * `country` (`pulumi.Input[str]`) - Two digit code that specifies the country in which the certificate subject located.
            * `distinguishedNameQualifier` (`pulumi.Input[str]`) - Disambiguating information for the certificate subject.
            * `generationQualifier` (`pulumi.Input[str]`) - Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
            * `givenName` (`pulumi.Input[str]`) - First name.
            * `initials` (`pulumi.Input[str]`) - Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`.
            * `locality` (`pulumi.Input[str]`) - The locality (such as a city or town) in which the certificate subject is located.
            * `organization` (`pulumi.Input[str]`) - Legal name of the organization with which the certificate subject is affiliated.
            * `organizationalUnit` (`pulumi.Input[str]`) - A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
            * `pseudonym` (`pulumi.Input[str]`) - Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
            * `state` (`pulumi.Input[str]`) - State in which the subject of the certificate is located.
            * `surname` (`pulumi.Input[str]`) - Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
            * `title` (`pulumi.Input[str]`) - A title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject.

        The **revocation_configuration** object supports the following:

          * `crlConfiguration` (`pulumi.Input[dict]`) - Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
            * `customCname` (`pulumi.Input[str]`) - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public.
            * `enabled` (`pulumi.Input[bool]`) - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
            * `expirationInDays` (`pulumi.Input[float]`) - Number of days until a certificate expires. Must be between 1 and 5000.
            * `s3_bucket_name` (`pulumi.Input[str]`) - Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

