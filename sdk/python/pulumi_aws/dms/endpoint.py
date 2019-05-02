# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Endpoint(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) for the certificate.
    """
    database_name: pulumi.Output[str]
    """
    The name of the endpoint database.
    """
    endpoint_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) for the endpoint.
    """
    endpoint_id: pulumi.Output[str]
    """
    The database endpoint identifier.
    """
    endpoint_type: pulumi.Output[str]
    """
    The type of endpoint. Can be one of `source | target`.
    """
    engine_name: pulumi.Output[str]
    """
    The type of engine for the endpoint. Can be one of `aurora | azuredb | docdb | dynamodb | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
    """
    extra_connection_attributes: pulumi.Output[str]
    """
    Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
    """
    kms_key_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
    """
    mongodb_settings: pulumi.Output[dict]
    """
    Settings for the source MongoDB endpoint. Available settings are `auth_type` (default: `PASSWORD`), `auth_mechanism` (default: `DEFAULT`), `nesting_level` (default: `NONE`), `extract_doc_id` (default: `false`), `docs_to_investigate` (default: `1000`) and `auth_source` (default: `admin`). For more details, see [Using MongoDB as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html).
    """
    password: pulumi.Output[str]
    """
    The password to be used to login to the endpoint database.
    """
    port: pulumi.Output[float]
    """
    The port used by the endpoint database.
    """
    s3_settings: pulumi.Output[dict]
    """
    Settings for the target S3 endpoint. Available settings are `service_access_role_arn`, `external_table_definition`, `csv_row_delimiter` (default: `\\n`), `csv_delimiter` (default: `,`), `bucket_folder`, `bucket_name` and `compression_type` (default: `NONE`). For more details, see [Using Amazon S3 as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html).
    """
    server_name: pulumi.Output[str]
    """
    The host name of the server.
    """
    service_access_role: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
    """
    ssl_mode: pulumi.Output[str]
    """
    The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    username: pulumi.Output[str]
    """
    The user name to be used to login to the endpoint database.
    """
    def __init__(__self__, resource_name, opts=None, certificate_arn=None, database_name=None, endpoint_id=None, endpoint_type=None, engine_name=None, extra_connection_attributes=None, kms_key_arn=None, mongodb_settings=None, password=None, port=None, s3_settings=None, server_name=None, service_access_role=None, ssl_mode=None, tags=None, username=None, __name__=None, __opts__=None):
        """
        Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.
        
        > **Note:** All arguments including the password will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The Amazon Resource Name (ARN) for the certificate.
        :param pulumi.Input[str] database_name: The name of the endpoint database.
        :param pulumi.Input[str] endpoint_id: The database endpoint identifier.
        :param pulumi.Input[str] endpoint_type: The type of endpoint. Can be one of `source | target`.
        :param pulumi.Input[str] engine_name: The type of engine for the endpoint. Can be one of `aurora | azuredb | docdb | dynamodb | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        :param pulumi.Input[str] extra_connection_attributes: Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        :param pulumi.Input[str] kms_key_arn: The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        :param pulumi.Input[dict] mongodb_settings: Settings for the source MongoDB endpoint. Available settings are `auth_type` (default: `PASSWORD`), `auth_mechanism` (default: `DEFAULT`), `nesting_level` (default: `NONE`), `extract_doc_id` (default: `false`), `docs_to_investigate` (default: `1000`) and `auth_source` (default: `admin`). For more details, see [Using MongoDB as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html).
        :param pulumi.Input[str] password: The password to be used to login to the endpoint database.
        :param pulumi.Input[float] port: The port used by the endpoint database.
        :param pulumi.Input[dict] s3_settings: Settings for the target S3 endpoint. Available settings are `service_access_role_arn`, `external_table_definition`, `csv_row_delimiter` (default: `\\n`), `csv_delimiter` (default: `,`), `bucket_folder`, `bucket_name` and `compression_type` (default: `NONE`). For more details, see [Using Amazon S3 as a Target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html).
        :param pulumi.Input[str] server_name: The host name of the server.
        :param pulumi.Input[str] service_access_role: The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        :param pulumi.Input[str] ssl_mode: The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] username: The user name to be used to login to the endpoint database.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['certificate_arn'] = certificate_arn

        __props__['database_name'] = database_name

        if endpoint_id is None:
            raise TypeError("Missing required property 'endpoint_id'")
        __props__['endpoint_id'] = endpoint_id

        if endpoint_type is None:
            raise TypeError("Missing required property 'endpoint_type'")
        __props__['endpoint_type'] = endpoint_type

        if engine_name is None:
            raise TypeError("Missing required property 'engine_name'")
        __props__['engine_name'] = engine_name

        __props__['extra_connection_attributes'] = extra_connection_attributes

        __props__['kms_key_arn'] = kms_key_arn

        __props__['mongodb_settings'] = mongodb_settings

        __props__['password'] = password

        __props__['port'] = port

        __props__['s3_settings'] = s3_settings

        __props__['server_name'] = server_name

        __props__['service_access_role'] = service_access_role

        __props__['ssl_mode'] = ssl_mode

        __props__['tags'] = tags

        __props__['username'] = username

        __props__['endpoint_arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Endpoint, __self__).__init__(
            'aws:dms/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

