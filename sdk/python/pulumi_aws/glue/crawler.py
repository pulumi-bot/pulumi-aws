# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Crawler(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the crawler 
    """
    catalog_targets: pulumi.Output[list]
    classifiers: pulumi.Output[list]
    """
    List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
    """
    configuration: pulumi.Output[str]
    """
    JSON string of configuration information.
    """
    database_name: pulumi.Output[str]
    """
    Glue database where results are written.
    """
    description: pulumi.Output[str]
    """
    Description of the crawler.
    """
    dynamodb_targets: pulumi.Output[list]
    """
    List of nested DynamoDB target arguments. See below.

      * `path` (`str`) - The name of the DynamoDB table to crawl.
    """
    jdbc_targets: pulumi.Output[list]
    """
    List of nested JBDC target arguments. See below.

      * `connectionName` (`str`) - The name of the connection to use to connect to the JDBC target.
      * `exclusions` (`list`) - A list of glob patterns used to exclude from the crawl.
      * `path` (`str`) - The path of the JDBC target.
    """
    name: pulumi.Output[str]
    """
    Name of the crawler.
    """
    role: pulumi.Output[str]
    """
    The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
    """
    s3_targets: pulumi.Output[list]
    """
    List nested Amazon S3 target arguments. See below.

      * `exclusions` (`list`) - A list of glob patterns used to exclude from the crawl.
      * `path` (`str`) - The name of the DynamoDB table to crawl.
    """
    schedule: pulumi.Output[str]
    """
    A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
    """
    schema_change_policy: pulumi.Output[dict]
    """
    Policy for the crawler's update and deletion behavior.

      * `deleteBehavior` (`str`) - The deletion behavior when the crawler finds a deleted object. Valid values: `LOG`, `DELETE_FROM_DATABASE`, or `DEPRECATE_IN_DATABASE`. Defaults to `DEPRECATE_IN_DATABASE`.
      * `updateBehavior` (`str`) - The update behavior when the crawler finds a changed schema. Valid values: `LOG` or `UPDATE_IN_DATABASE`. Defaults to `UPDATE_IN_DATABASE`.
    """
    security_configuration: pulumi.Output[str]
    """
    The name of Security Configuration to be used by the crawler
    """
    table_prefix: pulumi.Output[str]
    """
    The table prefix used for catalog tables that are created.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    def __init__(__self__, resource_name, opts=None, catalog_targets=None, classifiers=None, configuration=None, database_name=None, description=None, dynamodb_targets=None, jdbc_targets=None, name=None, role=None, s3_targets=None, schedule=None, schema_change_policy=None, security_configuration=None, table_prefix=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Glue Crawler. More information can be found in the [AWS Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)

        ## Example Usage

        ### DynamoDB Target

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Crawler("example",
            database_name=aws_glue_catalog_database["example"]["name"],
            dynamodb_targets=[{
                "path": "table-name",
            }],
            role=aws_iam_role["example"]["arn"])
        ```

        ### JDBC Target

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Crawler("example",
            database_name=aws_glue_catalog_database["example"]["name"],
            jdbc_targets=[{
                "connectionName": aws_glue_connection["example"]["name"],
                "path": "database-name/%",
            }],
            role=aws_iam_role["example"]["arn"])
        ```

        ### S3 Target

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Crawler("example",
            database_name=aws_glue_catalog_database["example"]["name"],
            role=aws_iam_role["example"]["arn"],
            s3_targets=[{
                "path": f"s3://{aws_s3_bucket['example']['bucket']}",
            }])
        ```

        ### Catalog Target

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.Crawler("example",
            catalog_targets=[{
                "databaseName": aws_glue_catalog_database["example"]["name"],
                "tables": [aws_glue_catalog_table["example"]["name"]],
            }],
            configuration="""{
          "Version":1.0,
          "Grouping": {
            "TableGroupingPolicy": "CombineCompatibleSchemas"
          }
        }

        """,
            database_name=aws_glue_catalog_database["example"]["name"],
            role=aws_iam_role["example"]["arn"],
            schema_change_policy={
                "deleteBehavior": "LOG",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] classifiers: List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
        :param pulumi.Input[str] configuration: JSON string of configuration information.
        :param pulumi.Input[str] database_name: Glue database where results are written.
        :param pulumi.Input[str] description: Description of the crawler.
        :param pulumi.Input[list] dynamodb_targets: List of nested DynamoDB target arguments. See below.
        :param pulumi.Input[list] jdbc_targets: List of nested JBDC target arguments. See below.
        :param pulumi.Input[str] name: Name of the crawler.
        :param pulumi.Input[str] role: The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
        :param pulumi.Input[list] s3_targets: List nested Amazon S3 target arguments. See below.
        :param pulumi.Input[str] schedule: A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
        :param pulumi.Input[dict] schema_change_policy: Policy for the crawler's update and deletion behavior.
        :param pulumi.Input[str] security_configuration: The name of Security Configuration to be used by the crawler
        :param pulumi.Input[str] table_prefix: The table prefix used for catalog tables that are created.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **catalog_targets** object supports the following:

          * `database_name` (`pulumi.Input[str]`) - The name of the Glue database to be synchronized.
          * `tables` (`pulumi.Input[list]`) - A list of catalog tables to be synchronized.

        The **dynamodb_targets** object supports the following:

          * `path` (`pulumi.Input[str]`) - The name of the DynamoDB table to crawl.

        The **jdbc_targets** object supports the following:

          * `connectionName` (`pulumi.Input[str]`) - The name of the connection to use to connect to the JDBC target.
          * `exclusions` (`pulumi.Input[list]`) - A list of glob patterns used to exclude from the crawl.
          * `path` (`pulumi.Input[str]`) - The path of the JDBC target.

        The **s3_targets** object supports the following:

          * `exclusions` (`pulumi.Input[list]`) - A list of glob patterns used to exclude from the crawl.
          * `path` (`pulumi.Input[str]`) - The name of the DynamoDB table to crawl.

        The **schema_change_policy** object supports the following:

          * `deleteBehavior` (`pulumi.Input[str]`) - The deletion behavior when the crawler finds a deleted object. Valid values: `LOG`, `DELETE_FROM_DATABASE`, or `DEPRECATE_IN_DATABASE`. Defaults to `DEPRECATE_IN_DATABASE`.
          * `updateBehavior` (`pulumi.Input[str]`) - The update behavior when the crawler finds a changed schema. Valid values: `LOG` or `UPDATE_IN_DATABASE`. Defaults to `UPDATE_IN_DATABASE`.
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
    def get(resource_name, id, opts=None, arn=None, catalog_targets=None, classifiers=None, configuration=None, database_name=None, description=None, dynamodb_targets=None, jdbc_targets=None, name=None, role=None, s3_targets=None, schedule=None, schema_change_policy=None, security_configuration=None, table_prefix=None, tags=None):
        """
        Get an existing Crawler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the crawler 
        :param pulumi.Input[list] classifiers: List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
        :param pulumi.Input[str] configuration: JSON string of configuration information.
        :param pulumi.Input[str] database_name: Glue database where results are written.
        :param pulumi.Input[str] description: Description of the crawler.
        :param pulumi.Input[list] dynamodb_targets: List of nested DynamoDB target arguments. See below.
        :param pulumi.Input[list] jdbc_targets: List of nested JBDC target arguments. See below.
        :param pulumi.Input[str] name: Name of the crawler.
        :param pulumi.Input[str] role: The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
        :param pulumi.Input[list] s3_targets: List nested Amazon S3 target arguments. See below.
        :param pulumi.Input[str] schedule: A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
        :param pulumi.Input[dict] schema_change_policy: Policy for the crawler's update and deletion behavior.
        :param pulumi.Input[str] security_configuration: The name of Security Configuration to be used by the crawler
        :param pulumi.Input[str] table_prefix: The table prefix used for catalog tables that are created.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **catalog_targets** object supports the following:

          * `database_name` (`pulumi.Input[str]`) - The name of the Glue database to be synchronized.
          * `tables` (`pulumi.Input[list]`) - A list of catalog tables to be synchronized.

        The **dynamodb_targets** object supports the following:

          * `path` (`pulumi.Input[str]`) - The name of the DynamoDB table to crawl.

        The **jdbc_targets** object supports the following:

          * `connectionName` (`pulumi.Input[str]`) - The name of the connection to use to connect to the JDBC target.
          * `exclusions` (`pulumi.Input[list]`) - A list of glob patterns used to exclude from the crawl.
          * `path` (`pulumi.Input[str]`) - The path of the JDBC target.

        The **s3_targets** object supports the following:

          * `exclusions` (`pulumi.Input[list]`) - A list of glob patterns used to exclude from the crawl.
          * `path` (`pulumi.Input[str]`) - The name of the DynamoDB table to crawl.

        The **schema_change_policy** object supports the following:

          * `deleteBehavior` (`pulumi.Input[str]`) - The deletion behavior when the crawler finds a deleted object. Valid values: `LOG`, `DELETE_FROM_DATABASE`, or `DEPRECATE_IN_DATABASE`. Defaults to `DEPRECATE_IN_DATABASE`.
          * `updateBehavior` (`pulumi.Input[str]`) - The update behavior when the crawler finds a changed schema. Valid values: `LOG` or `UPDATE_IN_DATABASE`. Defaults to `UPDATE_IN_DATABASE`.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

