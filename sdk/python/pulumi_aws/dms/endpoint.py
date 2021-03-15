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

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 endpoint_type: pulumi.Input[str],
                 engine_name: pulumi.Input[str],
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 elasticsearch_settings: Optional[pulumi.Input['EndpointElasticsearchSettingsArgs']] = None,
                 extra_connection_attributes: Optional[pulumi.Input[str]] = None,
                 kafka_settings: Optional[pulumi.Input['EndpointKafkaSettingsArgs']] = None,
                 kinesis_settings: Optional[pulumi.Input['EndpointKinesisSettingsArgs']] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 mongodb_settings: Optional[pulumi.Input['EndpointMongodbSettingsArgs']] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 s3_settings: Optional[pulumi.Input['EndpointS3SettingsArgs']] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 service_access_role: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Endpoint resource.
        :param pulumi.Input[str] endpoint_id: The database endpoint identifier.
        :param pulumi.Input[str] endpoint_type: The type of endpoint. Can be one of `source | target`.
        :param pulumi.Input[str] engine_name: The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        :param pulumi.Input[str] certificate_arn: The Amazon Resource Name (ARN) for the certificate.
        :param pulumi.Input[str] database_name: The name of the endpoint database.
        :param pulumi.Input['EndpointElasticsearchSettingsArgs'] elasticsearch_settings: Configuration block with Elasticsearch settings. Detailed below.
        :param pulumi.Input[str] extra_connection_attributes: Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        :param pulumi.Input['EndpointKafkaSettingsArgs'] kafka_settings: Configuration block with Kafka settings. Detailed below.
        :param pulumi.Input['EndpointKinesisSettingsArgs'] kinesis_settings: Configuration block with Kinesis settings. Detailed below.
        :param pulumi.Input[str] kms_key_arn: The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        :param pulumi.Input['EndpointMongodbSettingsArgs'] mongodb_settings: Configuration block with MongoDB settings. Detailed below.
        :param pulumi.Input[str] password: The password to be used to login to the endpoint database.
        :param pulumi.Input[int] port: The port used by the endpoint database.
        :param pulumi.Input['EndpointS3SettingsArgs'] s3_settings: Configuration block with S3 settings. Detailed below.
        :param pulumi.Input[str] server_name: The host name of the server.
        :param pulumi.Input[str] service_access_role: The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        :param pulumi.Input[str] ssl_mode: The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] username: The user name to be used to login to the endpoint database.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        pulumi.set(__self__, "engine_name", engine_name)
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if elasticsearch_settings is not None:
            pulumi.set(__self__, "elasticsearch_settings", elasticsearch_settings)
        if extra_connection_attributes is not None:
            pulumi.set(__self__, "extra_connection_attributes", extra_connection_attributes)
        if kafka_settings is not None:
            pulumi.set(__self__, "kafka_settings", kafka_settings)
        if kinesis_settings is not None:
            pulumi.set(__self__, "kinesis_settings", kinesis_settings)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if mongodb_settings is not None:
            pulumi.set(__self__, "mongodb_settings", mongodb_settings)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if s3_settings is not None:
            pulumi.set(__self__, "s3_settings", s3_settings)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if service_access_role is not None:
            pulumi.set(__self__, "service_access_role", service_access_role)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        """
        The database endpoint identifier.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Input[str]:
        """
        The type of endpoint. Can be one of `source | target`.
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter(name="engineName")
    def engine_name(self) -> pulumi.Input[str]:
        """
        The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        """
        return pulumi.get(self, "engine_name")

    @engine_name.setter
    def engine_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_name", value)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) for the certificate.
        """
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the endpoint database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="elasticsearchSettings")
    def elasticsearch_settings(self) -> Optional[pulumi.Input['EndpointElasticsearchSettingsArgs']]:
        """
        Configuration block with Elasticsearch settings. Detailed below.
        """
        return pulumi.get(self, "elasticsearch_settings")

    @elasticsearch_settings.setter
    def elasticsearch_settings(self, value: Optional[pulumi.Input['EndpointElasticsearchSettingsArgs']]):
        pulumi.set(self, "elasticsearch_settings", value)

    @property
    @pulumi.getter(name="extraConnectionAttributes")
    def extra_connection_attributes(self) -> Optional[pulumi.Input[str]]:
        """
        Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        """
        return pulumi.get(self, "extra_connection_attributes")

    @extra_connection_attributes.setter
    def extra_connection_attributes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extra_connection_attributes", value)

    @property
    @pulumi.getter(name="kafkaSettings")
    def kafka_settings(self) -> Optional[pulumi.Input['EndpointKafkaSettingsArgs']]:
        """
        Configuration block with Kafka settings. Detailed below.
        """
        return pulumi.get(self, "kafka_settings")

    @kafka_settings.setter
    def kafka_settings(self, value: Optional[pulumi.Input['EndpointKafkaSettingsArgs']]):
        pulumi.set(self, "kafka_settings", value)

    @property
    @pulumi.getter(name="kinesisSettings")
    def kinesis_settings(self) -> Optional[pulumi.Input['EndpointKinesisSettingsArgs']]:
        """
        Configuration block with Kinesis settings. Detailed below.
        """
        return pulumi.get(self, "kinesis_settings")

    @kinesis_settings.setter
    def kinesis_settings(self, value: Optional[pulumi.Input['EndpointKinesisSettingsArgs']]):
        pulumi.set(self, "kinesis_settings", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter(name="mongodbSettings")
    def mongodb_settings(self) -> Optional[pulumi.Input['EndpointMongodbSettingsArgs']]:
        """
        Configuration block with MongoDB settings. Detailed below.
        """
        return pulumi.get(self, "mongodb_settings")

    @mongodb_settings.setter
    def mongodb_settings(self, value: Optional[pulumi.Input['EndpointMongodbSettingsArgs']]):
        pulumi.set(self, "mongodb_settings", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password to be used to login to the endpoint database.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port used by the endpoint database.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="s3Settings")
    def s3_settings(self) -> Optional[pulumi.Input['EndpointS3SettingsArgs']]:
        """
        Configuration block with S3 settings. Detailed below.
        """
        return pulumi.get(self, "s3_settings")

    @s3_settings.setter
    def s3_settings(self, value: Optional[pulumi.Input['EndpointS3SettingsArgs']]):
        pulumi.set(self, "s3_settings", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        """
        The host name of the server.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="serviceAccessRole")
    def service_access_role(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        """
        return pulumi.get(self, "service_access_role")

    @service_access_role.setter
    def service_access_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_access_role", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        """
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_mode", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The user name to be used to login to the endpoint database.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Endpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.

        > **Note:** All arguments including the password will be stored in the raw state as plain-text.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new endpoint
        test = aws.dms.Endpoint("test",
            certificate_arn="arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
            database_name="test",
            endpoint_id="test-dms-endpoint-tf",
            endpoint_type="source",
            engine_name="aurora",
            extra_connection_attributes="",
            kms_key_arn="arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
            password="test",
            port=3306,
            server_name="test",
            ssl_mode="none",
            tags={
                "Name": "test",
            },
            username="test")
        ```

        ## Import

        Endpoints can be imported using the `endpoint_id`, e.g.

        ```sh
         $ pulumi import aws:dms/endpoint:Endpoint test test-dms-endpoint-tf
        ```

        :param str resource_name: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 elasticsearch_settings: Optional[pulumi.Input[pulumi.InputType['EndpointElasticsearchSettingsArgs']]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 engine_name: Optional[pulumi.Input[str]] = None,
                 extra_connection_attributes: Optional[pulumi.Input[str]] = None,
                 kafka_settings: Optional[pulumi.Input[pulumi.InputType['EndpointKafkaSettingsArgs']]] = None,
                 kinesis_settings: Optional[pulumi.Input[pulumi.InputType['EndpointKinesisSettingsArgs']]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 mongodb_settings: Optional[pulumi.Input[pulumi.InputType['EndpointMongodbSettingsArgs']]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 s3_settings: Optional[pulumi.Input[pulumi.InputType['EndpointS3SettingsArgs']]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 service_access_role: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.

        > **Note:** All arguments including the password will be stored in the raw state as plain-text.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new endpoint
        test = aws.dms.Endpoint("test",
            certificate_arn="arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
            database_name="test",
            endpoint_id="test-dms-endpoint-tf",
            endpoint_type="source",
            engine_name="aurora",
            extra_connection_attributes="",
            kms_key_arn="arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
            password="test",
            port=3306,
            server_name="test",
            ssl_mode="none",
            tags={
                "Name": "test",
            },
            username="test")
        ```

        ## Import

        Endpoints can be imported using the `endpoint_id`, e.g.

        ```sh
         $ pulumi import aws:dms/endpoint:Endpoint test test-dms-endpoint-tf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The Amazon Resource Name (ARN) for the certificate.
        :param pulumi.Input[str] database_name: The name of the endpoint database.
        :param pulumi.Input[pulumi.InputType['EndpointElasticsearchSettingsArgs']] elasticsearch_settings: Configuration block with Elasticsearch settings. Detailed below.
        :param pulumi.Input[str] endpoint_id: The database endpoint identifier.
        :param pulumi.Input[str] endpoint_type: The type of endpoint. Can be one of `source | target`.
        :param pulumi.Input[str] engine_name: The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        :param pulumi.Input[str] extra_connection_attributes: Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        :param pulumi.Input[pulumi.InputType['EndpointKafkaSettingsArgs']] kafka_settings: Configuration block with Kafka settings. Detailed below.
        :param pulumi.Input[pulumi.InputType['EndpointKinesisSettingsArgs']] kinesis_settings: Configuration block with Kinesis settings. Detailed below.
        :param pulumi.Input[str] kms_key_arn: The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        :param pulumi.Input[pulumi.InputType['EndpointMongodbSettingsArgs']] mongodb_settings: Configuration block with MongoDB settings. Detailed below.
        :param pulumi.Input[str] password: The password to be used to login to the endpoint database.
        :param pulumi.Input[int] port: The port used by the endpoint database.
        :param pulumi.Input[pulumi.InputType['EndpointS3SettingsArgs']] s3_settings: Configuration block with S3 settings. Detailed below.
        :param pulumi.Input[str] server_name: The host name of the server.
        :param pulumi.Input[str] service_access_role: The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        :param pulumi.Input[str] ssl_mode: The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] username: The user name to be used to login to the endpoint database.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 elasticsearch_settings: Optional[pulumi.Input[pulumi.InputType['EndpointElasticsearchSettingsArgs']]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 engine_name: Optional[pulumi.Input[str]] = None,
                 extra_connection_attributes: Optional[pulumi.Input[str]] = None,
                 kafka_settings: Optional[pulumi.Input[pulumi.InputType['EndpointKafkaSettingsArgs']]] = None,
                 kinesis_settings: Optional[pulumi.Input[pulumi.InputType['EndpointKinesisSettingsArgs']]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 mongodb_settings: Optional[pulumi.Input[pulumi.InputType['EndpointMongodbSettingsArgs']]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 s3_settings: Optional[pulumi.Input[pulumi.InputType['EndpointS3SettingsArgs']]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 service_access_role: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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

            __props__['certificate_arn'] = certificate_arn
            __props__['database_name'] = database_name
            __props__['elasticsearch_settings'] = elasticsearch_settings
            if endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__['endpoint_id'] = endpoint_id
            if endpoint_type is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_type'")
            __props__['endpoint_type'] = endpoint_type
            if engine_name is None and not opts.urn:
                raise TypeError("Missing required property 'engine_name'")
            __props__['engine_name'] = engine_name
            __props__['extra_connection_attributes'] = extra_connection_attributes
            __props__['kafka_settings'] = kafka_settings
            __props__['kinesis_settings'] = kinesis_settings
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
        super(Endpoint, __self__).__init__(
            'aws:dms/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            elasticsearch_settings: Optional[pulumi.Input[pulumi.InputType['EndpointElasticsearchSettingsArgs']]] = None,
            endpoint_arn: Optional[pulumi.Input[str]] = None,
            endpoint_id: Optional[pulumi.Input[str]] = None,
            endpoint_type: Optional[pulumi.Input[str]] = None,
            engine_name: Optional[pulumi.Input[str]] = None,
            extra_connection_attributes: Optional[pulumi.Input[str]] = None,
            kafka_settings: Optional[pulumi.Input[pulumi.InputType['EndpointKafkaSettingsArgs']]] = None,
            kinesis_settings: Optional[pulumi.Input[pulumi.InputType['EndpointKinesisSettingsArgs']]] = None,
            kms_key_arn: Optional[pulumi.Input[str]] = None,
            mongodb_settings: Optional[pulumi.Input[pulumi.InputType['EndpointMongodbSettingsArgs']]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            s3_settings: Optional[pulumi.Input[pulumi.InputType['EndpointS3SettingsArgs']]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            service_access_role: Optional[pulumi.Input[str]] = None,
            ssl_mode: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The Amazon Resource Name (ARN) for the certificate.
        :param pulumi.Input[str] database_name: The name of the endpoint database.
        :param pulumi.Input[pulumi.InputType['EndpointElasticsearchSettingsArgs']] elasticsearch_settings: Configuration block with Elasticsearch settings. Detailed below.
        :param pulumi.Input[str] endpoint_arn: The Amazon Resource Name (ARN) for the endpoint.
        :param pulumi.Input[str] endpoint_id: The database endpoint identifier.
        :param pulumi.Input[str] endpoint_type: The type of endpoint. Can be one of `source | target`.
        :param pulumi.Input[str] engine_name: The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        :param pulumi.Input[str] extra_connection_attributes: Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        :param pulumi.Input[pulumi.InputType['EndpointKafkaSettingsArgs']] kafka_settings: Configuration block with Kafka settings. Detailed below.
        :param pulumi.Input[pulumi.InputType['EndpointKinesisSettingsArgs']] kinesis_settings: Configuration block with Kinesis settings. Detailed below.
        :param pulumi.Input[str] kms_key_arn: The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        :param pulumi.Input[pulumi.InputType['EndpointMongodbSettingsArgs']] mongodb_settings: Configuration block with MongoDB settings. Detailed below.
        :param pulumi.Input[str] password: The password to be used to login to the endpoint database.
        :param pulumi.Input[int] port: The port used by the endpoint database.
        :param pulumi.Input[pulumi.InputType['EndpointS3SettingsArgs']] s3_settings: Configuration block with S3 settings. Detailed below.
        :param pulumi.Input[str] server_name: The host name of the server.
        :param pulumi.Input[str] service_access_role: The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        :param pulumi.Input[str] ssl_mode: The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] username: The user name to be used to login to the endpoint database.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_arn"] = certificate_arn
        __props__["database_name"] = database_name
        __props__["elasticsearch_settings"] = elasticsearch_settings
        __props__["endpoint_arn"] = endpoint_arn
        __props__["endpoint_id"] = endpoint_id
        __props__["endpoint_type"] = endpoint_type
        __props__["engine_name"] = engine_name
        __props__["extra_connection_attributes"] = extra_connection_attributes
        __props__["kafka_settings"] = kafka_settings
        __props__["kinesis_settings"] = kinesis_settings
        __props__["kms_key_arn"] = kms_key_arn
        __props__["mongodb_settings"] = mongodb_settings
        __props__["password"] = password
        __props__["port"] = port
        __props__["s3_settings"] = s3_settings
        __props__["server_name"] = server_name
        __props__["service_access_role"] = service_access_role
        __props__["ssl_mode"] = ssl_mode
        __props__["tags"] = tags
        __props__["username"] = username
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the certificate.
        """
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the endpoint database.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="elasticsearchSettings")
    def elasticsearch_settings(self) -> pulumi.Output[Optional['outputs.EndpointElasticsearchSettings']]:
        """
        Configuration block with Elasticsearch settings. Detailed below.
        """
        return pulumi.get(self, "elasticsearch_settings")

    @property
    @pulumi.getter(name="endpointArn")
    def endpoint_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the endpoint.
        """
        return pulumi.get(self, "endpoint_arn")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        """
        The database endpoint identifier.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> pulumi.Output[str]:
        """
        The type of endpoint. Can be one of `source | target`.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="engineName")
    def engine_name(self) -> pulumi.Output[str]:
        """
        The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        """
        return pulumi.get(self, "engine_name")

    @property
    @pulumi.getter(name="extraConnectionAttributes")
    def extra_connection_attributes(self) -> pulumi.Output[str]:
        """
        Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        """
        return pulumi.get(self, "extra_connection_attributes")

    @property
    @pulumi.getter(name="kafkaSettings")
    def kafka_settings(self) -> pulumi.Output[Optional['outputs.EndpointKafkaSettings']]:
        """
        Configuration block with Kafka settings. Detailed below.
        """
        return pulumi.get(self, "kafka_settings")

    @property
    @pulumi.getter(name="kinesisSettings")
    def kinesis_settings(self) -> pulumi.Output[Optional['outputs.EndpointKinesisSettings']]:
        """
        Configuration block with Kinesis settings. Detailed below.
        """
        return pulumi.get(self, "kinesis_settings")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="mongodbSettings")
    def mongodb_settings(self) -> pulumi.Output[Optional['outputs.EndpointMongodbSettings']]:
        """
        Configuration block with MongoDB settings. Detailed below.
        """
        return pulumi.get(self, "mongodb_settings")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password to be used to login to the endpoint database.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port used by the endpoint database.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="s3Settings")
    def s3_settings(self) -> pulumi.Output[Optional['outputs.EndpointS3Settings']]:
        """
        Configuration block with S3 settings. Detailed below.
        """
        return pulumi.get(self, "s3_settings")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Output[Optional[str]]:
        """
        The host name of the server.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="serviceAccessRole")
    def service_access_role(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        """
        return pulumi.get(self, "service_access_role")

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> pulumi.Output[str]:
        """
        The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        """
        return pulumi.get(self, "ssl_mode")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The user name to be used to login to the endpoint database.
        """
        return pulumi.get(self, "username")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

