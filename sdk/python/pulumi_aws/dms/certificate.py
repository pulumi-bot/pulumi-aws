# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Certificate(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) for the certificate.
    """
    certificate_id: pulumi.Output[str]
    """
    The certificate identifier.
    """
    certificate_pem: pulumi.Output[str]
    """
    The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
    """
    certificate_wallet: pulumi.Output[str]
    """
    The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
    """
    def __init__(__self__, resource_name, opts=None, certificate_id=None, certificate_pem=None, certificate_wallet=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DMS (Data Migration Service) certificate resource. DMS certificates can be created, deleted, and imported.

        > **Note:** All arguments including the PEM encoded certificate will be stored in the raw state as plain-text.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new certificate
        test = aws.dms.Certificate("test",
            certificate_id="test-dms-certificate-tf",
            certificate_pem="...")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: The certificate identifier.
        :param pulumi.Input[str] certificate_pem: The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
        :param pulumi.Input[str] certificate_wallet: The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
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

            if certificate_id is None:
                raise TypeError("Missing required property 'certificate_id'")
            __props__['certificate_id'] = certificate_id
            __props__['certificate_pem'] = certificate_pem
            __props__['certificate_wallet'] = certificate_wallet
            __props__['certificate_arn'] = None
        super(Certificate, __self__).__init__(
            'aws:dms/certificate:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, certificate_arn=None, certificate_id=None, certificate_pem=None, certificate_wallet=None):
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The Amazon Resource Name (ARN) for the certificate.
        :param pulumi.Input[str] certificate_id: The certificate identifier.
        :param pulumi.Input[str] certificate_pem: The contents of the .pem X.509 certificate file for the certificate. Either `certificate_pem` or `certificate_wallet` must be set.
        :param pulumi.Input[str] certificate_wallet: The contents of the Oracle Wallet certificate for use with SSL. Either `certificate_pem` or `certificate_wallet` must be set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_arn"] = certificate_arn
        __props__["certificate_id"] = certificate_id
        __props__["certificate_pem"] = certificate_pem
        __props__["certificate_wallet"] = certificate_wallet
        return Certificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
