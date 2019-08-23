# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetCertificateAuthorityResult:
    """
    A collection of values returned by getCertificateAuthority.
    """
    def __init__(__self__, arn=None, certificate=None, certificate_chain=None, certificate_signing_request=None, not_after=None, not_before=None, revocation_configurations=None, serial=None, status=None, tags=None, type=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        __self__.certificate = certificate
        """
        Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        """
        if certificate_chain and not isinstance(certificate_chain, str):
            raise TypeError("Expected argument 'certificate_chain' to be a str")
        __self__.certificate_chain = certificate_chain
        """
        Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        """
        if certificate_signing_request and not isinstance(certificate_signing_request, str):
            raise TypeError("Expected argument 'certificate_signing_request' to be a str")
        __self__.certificate_signing_request = certificate_signing_request
        """
        The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        """
        if not_after and not isinstance(not_after, str):
            raise TypeError("Expected argument 'not_after' to be a str")
        __self__.not_after = not_after
        """
        Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        """
        if not_before and not isinstance(not_before, str):
            raise TypeError("Expected argument 'not_before' to be a str")
        __self__.not_before = not_before
        """
        Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        """
        if revocation_configurations and not isinstance(revocation_configurations, list):
            raise TypeError("Expected argument 'revocation_configurations' to be a list")
        __self__.revocation_configurations = revocation_configurations
        """
        Nested attribute containing revocation configuration.
        * `revocation_configuration.0.crl_configuration` - Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.
        * `revocation_configuration.0.crl_configuration.0.custom_cname` - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.
        * `revocation_configuration.0.crl_configuration.0.enabled` - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
        * `revocation_configuration.0.crl_configuration.0.expiration_in_days` - Number of days until a certificate expires.
        * `revocation_configuration.0.crl_configuration.0.s3_bucket_name` - Name of the S3 bucket that contains the CRL.
        """
        if serial and not isinstance(serial, str):
            raise TypeError("Expected argument 'serial' to be a str")
        __self__.serial = serial
        """
        Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        Status of the certificate authority.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the certificate authority.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetCertificateAuthorityResult(GetCertificateAuthorityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateAuthorityResult(
            arn=self.arn,
            certificate=self.certificate,
            certificate_chain=self.certificate_chain,
            certificate_signing_request=self.certificate_signing_request,
            not_after=self.not_after,
            not_before=self.not_before,
            revocation_configurations=self.revocation_configurations,
            serial=self.serial,
            status=self.status,
            tags=self.tags,
            type=self.type,
            id=self.id)

def get_certificate_authority(arn=None,revocation_configurations=None,tags=None,opts=None):
    """
    Get information on a AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/acmpca_certificate_authority.html.markdown.
    """
    __args__ = dict()

    __args__['arn'] = arn
    __args__['revocationConfigurations'] = revocation_configurations
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:acmpca/getCertificateAuthority:getCertificateAuthority', __args__, opts=opts).value

    return AwaitableGetCertificateAuthorityResult(
        arn=__ret__.get('arn'),
        certificate=__ret__.get('certificate'),
        certificate_chain=__ret__.get('certificateChain'),
        certificate_signing_request=__ret__.get('certificateSigningRequest'),
        not_after=__ret__.get('notAfter'),
        not_before=__ret__.get('notBefore'),
        revocation_configurations=__ret__.get('revocationConfigurations'),
        serial=__ret__.get('serial'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'),
        id=__ret__.get('id'))
