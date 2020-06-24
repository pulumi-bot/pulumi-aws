# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ClientCertificate(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN)
    """
    created_date: pulumi.Output[str]
    """
    The date when the client certificate was created.
    """
    description: pulumi.Output[str]
    """
    The description of the client certificate.
    """
    expiration_date: pulumi.Output[str]
    """
    The date when the client certificate will expire.
    """
    pem_encoded_certificate: pulumi.Output[str]
    """
    The PEM-encoded public key of the client certificate.
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of resource tags
    """
    def __init__(__self__, resource_name, opts=None, description=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Client Certificate.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        demo = aws.apigateway.ClientCertificate("demo", description="My client certificate")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the client certificate.
        :param pulumi.Input[dict] tags: Key-value map of resource tags
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

            __props__['description'] = description
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['expiration_date'] = None
            __props__['pem_encoded_certificate'] = None
        super(ClientCertificate, __self__).__init__(
            'aws:apigateway/clientCertificate:ClientCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, created_date=None, description=None, expiration_date=None, pem_encoded_certificate=None, tags=None):
        """
        Get an existing ClientCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] created_date: The date when the client certificate was created.
        :param pulumi.Input[str] description: The description of the client certificate.
        :param pulumi.Input[str] expiration_date: The date when the client certificate will expire.
        :param pulumi.Input[str] pem_encoded_certificate: The PEM-encoded public key of the client certificate.
        :param pulumi.Input[dict] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["created_date"] = created_date
        __props__["description"] = description
        __props__["expiration_date"] = expiration_date
        __props__["pem_encoded_certificate"] = pem_encoded_certificate
        __props__["tags"] = tags
        return ClientCertificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
