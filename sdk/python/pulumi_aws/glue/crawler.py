# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Crawler']


class Crawler(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerCatalogTargetArgs']]]]] = None,
                 classifiers: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 configuration: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamodb_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerDynamodbTargetArgs']]]]] = None,
                 jdbc_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerJdbcTargetArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 s3_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerS3TargetArgs']]]]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 schema_change_policy: Optional[pulumi.Input[pulumi.InputType['CrawlerSchemaChangePolicyArgs']]] = None,
                 security_configuration: Optional[pulumi.Input[str]] = None,
                 table_prefix: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Crawler resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            __props__['catalog_targets'] = catalog_targets
            __props__['classifiers'] = classifiers
            __props__['configuration'] = configuration
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['description'] = description
            __props__['dynamodb_targets'] = dynamodb_targets
            __props__['jdbc_targets'] = jdbc_targets
            __props__['name'] = name
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['s3_targets'] = s3_targets
            __props__['schedule'] = schedule
            __props__['schema_change_policy'] = schema_change_policy
            __props__['security_configuration'] = security_configuration
            __props__['table_prefix'] = table_prefix
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Crawler, __self__).__init__(
            'aws:glue/crawler:Crawler',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            catalog_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerCatalogTargetArgs']]]]] = None,
            classifiers: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            configuration: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamodb_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerDynamodbTargetArgs']]]]] = None,
            jdbc_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerJdbcTargetArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            s3_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CrawlerS3TargetArgs']]]]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            schema_change_policy: Optional[pulumi.Input[pulumi.InputType['CrawlerSchemaChangePolicyArgs']]] = None,
            security_configuration: Optional[pulumi.Input[str]] = None,
            table_prefix: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Crawler':
        """
        Get an existing Crawler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["catalog_targets"] = catalog_targets
        __props__["classifiers"] = classifiers
        __props__["configuration"] = configuration
        __props__["database_name"] = database_name
        __props__["description"] = description
        __props__["dynamodb_targets"] = dynamodb_targets
        __props__["jdbc_targets"] = jdbc_targets
        __props__["name"] = name
        __props__["role"] = role
        __props__["s3_targets"] = s3_targets
        __props__["schedule"] = schedule
        __props__["schema_change_policy"] = schema_change_policy
        __props__["security_configuration"] = security_configuration
        __props__["table_prefix"] = table_prefix
        __props__["tags"] = tags
        return Crawler(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="catalogTargets")
    def catalog_targets(self) -> pulumi.Output[Optional[List['outputs.CrawlerCatalogTarget']]]:
        return pulumi.get(self, "catalog_targets")

    @property
    @pulumi.getter
    def classifiers(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "classifiers")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamodbTargets")
    def dynamodb_targets(self) -> pulumi.Output[Optional[List['outputs.CrawlerDynamodbTarget']]]:
        return pulumi.get(self, "dynamodb_targets")

    @property
    @pulumi.getter(name="jdbcTargets")
    def jdbc_targets(self) -> pulumi.Output[Optional[List['outputs.CrawlerJdbcTarget']]]:
        return pulumi.get(self, "jdbc_targets")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="s3Targets")
    def s3_targets(self) -> pulumi.Output[Optional[List['outputs.CrawlerS3Target']]]:
        return pulumi.get(self, "s3_targets")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="schemaChangePolicy")
    def schema_change_policy(self) -> pulumi.Output[Optional['outputs.CrawlerSchemaChangePolicy']]:
        return pulumi.get(self, "schema_change_policy")

    @property
    @pulumi.getter(name="securityConfiguration")
    def security_configuration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "security_configuration")

    @property
    @pulumi.getter(name="tablePrefix")
    def table_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "table_prefix")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

