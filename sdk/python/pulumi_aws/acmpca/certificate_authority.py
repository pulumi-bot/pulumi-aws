# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
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
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] certificate_authority_configuration: Nested argument containing algorithms and certificate subject information. Defined below.
        :param pulumi.Input[bool] enabled: Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        :param pulumi.Input[float] permanent_deletion_time_in_days: The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        :param pulumi.Input[dict] revocation_configuration: Nested argument containing revocation configuration. Defined below.
        :param pulumi.Input[dict] tags: Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        :param pulumi.Input[str] type: The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acmpca_certificate_authority.html.markdown.
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

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acmpca_certificate_authority.html.markdown.
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

