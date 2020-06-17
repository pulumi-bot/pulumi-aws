# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AnalyticsApplication(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the Kinesis Analytics Appliation.
    """
    cloudwatch_logging_options: pulumi.Output[dict]
    """
    The CloudWatch log stream options to monitor application errors.
    See CloudWatch Logging Options below for more details.

      * `id` (`str`) - The ARN of the Kinesis Analytics Application.
      * `logStreamArn` (`str`) - The ARN of the CloudWatch Log Stream.
      * `role_arn` (`str`) - The ARN of the IAM Role used to send application messages.
    """
    code: pulumi.Output[str]
    """
    SQL Code to transform input data, and generate output.
    """
    create_timestamp: pulumi.Output[str]
    """
    The Timestamp when the application version was created.
    """
    description: pulumi.Output[str]
    """
    Description of the application.
    """
    inputs: pulumi.Output[dict]
    """
    Input configuration of the application. See Inputs below for more details.

      * `id` (`str`) - The ARN of the Kinesis Analytics Application.
      * `kinesisFirehose` (`dict`) - The Kinesis Firehose configuration for the streaming source. Conflicts with `kinesis_stream`.
        See Kinesis Firehose below for more details.
        * `resource_arn` (`str`) - The ARN of the Kinesis Firehose delivery stream.
        * `role_arn` (`str`) - The ARN of the IAM Role used to access the stream.

      * `kinesisStream` (`dict`) - The Kinesis Stream configuration for the streaming source. Conflicts with `kinesis_firehose`.
        See Kinesis Stream below for more details.
        * `resource_arn` (`str`) - The ARN of the Kinesis Stream.
        * `role_arn` (`str`) - The ARN of the IAM Role used to access the stream.

      * `name_prefix` (`str`) - The Name Prefix to use when creating an in-application stream.
      * `parallelism` (`dict`) - The number of Parallel in-application streams to create.
        See Parallelism below for more details.
        * `count` (`float`) - The Count of streams.

      * `processingConfiguration` (`dict`) - The Processing Configuration to transform records as they are received from the stream.
        See Processing Configuration below for more details.
        * `lambda_` (`dict`) - The Lambda function configuration. See Lambda below for more details.
          * `resource_arn` (`str`) - The ARN of the Lambda function.
          * `role_arn` (`str`) - The ARN of the IAM Role used to access the Lambda function.

      * `schema` (`dict`) - The Schema format of the data in the streaming source. See Source Schema below for more details.
        * `recordColumns` (`list`) - The Record Column mapping for the streaming source data element.
          See Record Columns below for more details.
          * `mapping` (`str`) - The Mapping reference to the data element.
          * `name` (`str`) - Name of the column.
          * `sqlType` (`str`) - The SQL Type of the column.

        * `recordEncoding` (`str`) - The Encoding of the record in the streaming source.
        * `recordFormat` (`dict`) - The Record Format and mapping information to schematize a record.
          See Record Format below for more details.
          * `mappingParameters` (`dict`) - The Mapping Information for the record format.
            See Mapping Parameters below for more details.
            * `csv` (`dict`) - Mapping information when the record format uses delimiters.
              See CSV Mapping Parameters below for more details.
              * `recordColumnDelimiter` (`str`) - The Column Delimiter.
              * `recordRowDelimiter` (`str`) - The Row Delimiter.

            * `json` (`dict`) - Mapping information when JSON is the record format on the streaming source.
              See JSON Mapping Parameters below for more details.
              * `recordRowPath` (`str`) - Path to the top-level parent that contains the records.

          * `recordFormatType` (`str`) - The type of Record Format. Can be `CSV` or `JSON`.

      * `startingPositionConfigurations` (`list`)
        * `starting_position` (`str`)

      * `streamNames` (`list`)
    """
    last_update_timestamp: pulumi.Output[str]
    """
    The Timestamp when the application was last updated.
    """
    name: pulumi.Output[str]
    """
    Name of the Kinesis Analytics Application.
    """
    outputs: pulumi.Output[list]
    """
    Output destination configuration of the application. See Outputs below for more details.

      * `id` (`str`) - The ARN of the Kinesis Analytics Application.
      * `kinesisFirehose` (`dict`) - The Kinesis Firehose configuration for the destination stream. Conflicts with `kinesis_stream`.
        See Kinesis Firehose below for more details.
        * `resource_arn` (`str`) - The ARN of the Kinesis Firehose delivery stream.
        * `role_arn` (`str`) - The ARN of the IAM Role used to access the stream.

      * `kinesisStream` (`dict`) - The Kinesis Stream configuration for the destination stream. Conflicts with `kinesis_firehose`.
        See Kinesis Stream below for more details.
        * `resource_arn` (`str`) - The ARN of the Kinesis Stream.
        * `role_arn` (`str`) - The ARN of the IAM Role used to access the stream.

      * `lambda_` (`dict`) - The Lambda function destination. See Lambda below for more details.
        * `resource_arn` (`str`) - The ARN of the Lambda function.
        * `role_arn` (`str`) - The ARN of the IAM Role used to access the Lambda function.

      * `name` (`str`) - The Name of the in-application stream.
      * `schema` (`dict`) - The Schema format of the data written to the destination. See Destination Schema below for more details.
        * `recordFormatType` (`str`) - The Format Type of the records on the output stream. Can be `CSV` or `JSON`.
    """
    reference_data_sources: pulumi.Output[dict]
    """
    An S3 Reference Data Source for the application.
    See Reference Data Sources below for more details.

      * `id` (`str`) - The ARN of the Kinesis Analytics Application.
      * `s3` (`dict`) - The S3 configuration for the reference data source. See S3 Reference below for more details.
        * `bucketArn` (`str`) - The S3 Bucket ARN.
        * `fileKey` (`str`) - The File Key name containing reference data.
        * `role_arn` (`str`) - The ARN of the IAM Role used to send application messages.

      * `schema` (`dict`) - The Schema format of the data in the streaming source. See Source Schema below for more details.
        * `recordColumns` (`list`) - The Record Column mapping for the streaming source data element.
          See Record Columns below for more details.
          * `mapping` (`str`) - The Mapping reference to the data element.
          * `name` (`str`) - Name of the column.
          * `sqlType` (`str`) - The SQL Type of the column.

        * `recordEncoding` (`str`) - The Encoding of the record in the streaming source.
        * `recordFormat` (`dict`) - The Record Format and mapping information to schematize a record.
          See Record Format below for more details.
          * `mappingParameters` (`dict`) - The Mapping Information for the record format.
            See Mapping Parameters below for more details.
            * `csv` (`dict`) - Mapping information when the record format uses delimiters.
              See CSV Mapping Parameters below for more details.
              * `recordColumnDelimiter` (`str`) - The Column Delimiter.
              * `recordRowDelimiter` (`str`) - The Row Delimiter.

            * `json` (`dict`) - Mapping information when JSON is the record format on the streaming source.
              See JSON Mapping Parameters below for more details.
              * `recordRowPath` (`str`) - Path to the top-level parent that contains the records.

          * `recordFormatType` (`str`) - The type of Record Format. Can be `CSV` or `JSON`.

      * `table_name` (`str`) - The in-application Table Name.
    """
    status: pulumi.Output[str]
    """
    The Status of the application.
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of tags for the Kinesis Analytics Application.
    """
    version: pulumi.Output[float]
    """
    The Version of the application.
    """
    def __init__(__self__, resource_name, opts=None, cloudwatch_logging_options=None, code=None, description=None, inputs=None, name=None, outputs=None, reference_data_sources=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
        allows processing and analyzing streaming data using standard SQL.

        For more details, see the [Amazon Kinesis Analytics Documentation](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/what-is.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_stream = aws.kinesis.Stream("testStream", shard_count=1)
        test_application = aws.kinesis.AnalyticsApplication("testApplication", inputs={
            "kinesisStream": {
                "resource_arn": test_stream.arn,
                "role_arn": aws_iam_role["test"]["arn"],
            },
            "name_prefix": "test_prefix",
            "parallelism": {
                "count": 1,
            },
            "schema": {
                "recordColumns": [{
                    "mapping": "$.test",
                    "name": "test",
                    "sqlType": "VARCHAR(8)",
                }],
                "recordEncoding": "UTF-8",
                "recordFormat": {
                    "mappingParameters": {
                        "json": {
                            "recordRowPath": "$",
                        },
                    },
                },
            },
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] cloudwatch_logging_options: The CloudWatch log stream options to monitor application errors.
               See CloudWatch Logging Options below for more details.
        :param pulumi.Input[str] code: SQL Code to transform input data, and generate output.
        :param pulumi.Input[str] description: Description of the application.
        :param pulumi.Input[dict] inputs: Input configuration of the application. See Inputs below for more details.
        :param pulumi.Input[str] name: Name of the Kinesis Analytics Application.
        :param pulumi.Input[list] outputs: Output destination configuration of the application. See Outputs below for more details.
        :param pulumi.Input[dict] reference_data_sources: An S3 Reference Data Source for the application.
               See Reference Data Sources below for more details.
        :param pulumi.Input[dict] tags: Key-value map of tags for the Kinesis Analytics Application.

        The **cloudwatch_logging_options** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `logStreamArn` (`pulumi.Input[str]`) - The ARN of the CloudWatch Log Stream.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to send application messages.

        The **inputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`) - The Kinesis Firehose configuration for the streaming source. Conflicts with `kinesis_stream`.
            See Kinesis Firehose below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Firehose delivery stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `kinesisStream` (`pulumi.Input[dict]`) - The Kinesis Stream configuration for the streaming source. Conflicts with `kinesis_firehose`.
            See Kinesis Stream below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `name_prefix` (`pulumi.Input[str]`) - The Name Prefix to use when creating an in-application stream.
          * `parallelism` (`pulumi.Input[dict]`) - The number of Parallel in-application streams to create.
            See Parallelism below for more details.
            * `count` (`pulumi.Input[float]`) - The Count of streams.

          * `processingConfiguration` (`pulumi.Input[dict]`) - The Processing Configuration to transform records as they are received from the stream.
            See Processing Configuration below for more details.
            * `lambda_` (`pulumi.Input[dict]`) - The Lambda function configuration. See Lambda below for more details.
              * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Lambda function.
              * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the Lambda function.

          * `schema` (`pulumi.Input[dict]`) - The Schema format of the data in the streaming source. See Source Schema below for more details.
            * `recordColumns` (`pulumi.Input[list]`) - The Record Column mapping for the streaming source data element.
              See Record Columns below for more details.
              * `mapping` (`pulumi.Input[str]`) - The Mapping reference to the data element.
              * `name` (`pulumi.Input[str]`) - Name of the column.
              * `sqlType` (`pulumi.Input[str]`) - The SQL Type of the column.

            * `recordEncoding` (`pulumi.Input[str]`) - The Encoding of the record in the streaming source.
            * `recordFormat` (`pulumi.Input[dict]`) - The Record Format and mapping information to schematize a record.
              See Record Format below for more details.
              * `mappingParameters` (`pulumi.Input[dict]`) - The Mapping Information for the record format.
                See Mapping Parameters below for more details.
                * `csv` (`pulumi.Input[dict]`) - Mapping information when the record format uses delimiters.
                  See CSV Mapping Parameters below for more details.
                  * `recordColumnDelimiter` (`pulumi.Input[str]`) - The Column Delimiter.
                  * `recordRowDelimiter` (`pulumi.Input[str]`) - The Row Delimiter.

                * `json` (`pulumi.Input[dict]`) - Mapping information when JSON is the record format on the streaming source.
                  See JSON Mapping Parameters below for more details.
                  * `recordRowPath` (`pulumi.Input[str]`) - Path to the top-level parent that contains the records.

              * `recordFormatType` (`pulumi.Input[str]`) - The type of Record Format. Can be `CSV` or `JSON`.

          * `startingPositionConfigurations` (`pulumi.Input[list]`)
            * `starting_position` (`pulumi.Input[str]`)

          * `streamNames` (`pulumi.Input[list]`)

        The **outputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`) - The Kinesis Firehose configuration for the destination stream. Conflicts with `kinesis_stream`.
            See Kinesis Firehose below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Firehose delivery stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `kinesisStream` (`pulumi.Input[dict]`) - The Kinesis Stream configuration for the destination stream. Conflicts with `kinesis_firehose`.
            See Kinesis Stream below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `lambda_` (`pulumi.Input[dict]`) - The Lambda function destination. See Lambda below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Lambda function.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the Lambda function.

          * `name` (`pulumi.Input[str]`) - The Name of the in-application stream.
          * `schema` (`pulumi.Input[dict]`) - The Schema format of the data written to the destination. See Destination Schema below for more details.
            * `recordFormatType` (`pulumi.Input[str]`) - The Format Type of the records on the output stream. Can be `CSV` or `JSON`.

        The **reference_data_sources** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `s3` (`pulumi.Input[dict]`) - The S3 configuration for the reference data source. See S3 Reference below for more details.
            * `bucketArn` (`pulumi.Input[str]`) - The S3 Bucket ARN.
            * `fileKey` (`pulumi.Input[str]`) - The File Key name containing reference data.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to send application messages.

          * `schema` (`pulumi.Input[dict]`) - The Schema format of the data in the streaming source. See Source Schema below for more details.
            * `recordColumns` (`pulumi.Input[list]`) - The Record Column mapping for the streaming source data element.
              See Record Columns below for more details.
              * `mapping` (`pulumi.Input[str]`) - The Mapping reference to the data element.
              * `name` (`pulumi.Input[str]`) - Name of the column.
              * `sqlType` (`pulumi.Input[str]`) - The SQL Type of the column.

            * `recordEncoding` (`pulumi.Input[str]`) - The Encoding of the record in the streaming source.
            * `recordFormat` (`pulumi.Input[dict]`) - The Record Format and mapping information to schematize a record.
              See Record Format below for more details.
              * `mappingParameters` (`pulumi.Input[dict]`) - The Mapping Information for the record format.
                See Mapping Parameters below for more details.
                * `csv` (`pulumi.Input[dict]`) - Mapping information when the record format uses delimiters.
                  See CSV Mapping Parameters below for more details.
                  * `recordColumnDelimiter` (`pulumi.Input[str]`) - The Column Delimiter.
                  * `recordRowDelimiter` (`pulumi.Input[str]`) - The Row Delimiter.

                * `json` (`pulumi.Input[dict]`) - Mapping information when JSON is the record format on the streaming source.
                  See JSON Mapping Parameters below for more details.
                  * `recordRowPath` (`pulumi.Input[str]`) - Path to the top-level parent that contains the records.

              * `recordFormatType` (`pulumi.Input[str]`) - The type of Record Format. Can be `CSV` or `JSON`.

          * `table_name` (`pulumi.Input[str]`) - The in-application Table Name.
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

            __props__['cloudwatch_logging_options'] = cloudwatch_logging_options
            __props__['code'] = code
            __props__['description'] = description
            __props__['inputs'] = inputs
            __props__['name'] = name
            __props__['outputs'] = outputs
            __props__['reference_data_sources'] = reference_data_sources
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['create_timestamp'] = None
            __props__['last_update_timestamp'] = None
            __props__['status'] = None
            __props__['version'] = None
        super(AnalyticsApplication, __self__).__init__(
            'aws:kinesis/analyticsApplication:AnalyticsApplication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, cloudwatch_logging_options=None, code=None, create_timestamp=None, description=None, inputs=None, last_update_timestamp=None, name=None, outputs=None, reference_data_sources=None, status=None, tags=None, version=None):
        """
        Get an existing AnalyticsApplication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Kinesis Analytics Appliation.
        :param pulumi.Input[dict] cloudwatch_logging_options: The CloudWatch log stream options to monitor application errors.
               See CloudWatch Logging Options below for more details.
        :param pulumi.Input[str] code: SQL Code to transform input data, and generate output.
        :param pulumi.Input[str] create_timestamp: The Timestamp when the application version was created.
        :param pulumi.Input[str] description: Description of the application.
        :param pulumi.Input[dict] inputs: Input configuration of the application. See Inputs below for more details.
        :param pulumi.Input[str] last_update_timestamp: The Timestamp when the application was last updated.
        :param pulumi.Input[str] name: Name of the Kinesis Analytics Application.
        :param pulumi.Input[list] outputs: Output destination configuration of the application. See Outputs below for more details.
        :param pulumi.Input[dict] reference_data_sources: An S3 Reference Data Source for the application.
               See Reference Data Sources below for more details.
        :param pulumi.Input[str] status: The Status of the application.
        :param pulumi.Input[dict] tags: Key-value map of tags for the Kinesis Analytics Application.
        :param pulumi.Input[float] version: The Version of the application.

        The **cloudwatch_logging_options** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `logStreamArn` (`pulumi.Input[str]`) - The ARN of the CloudWatch Log Stream.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to send application messages.

        The **inputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`) - The Kinesis Firehose configuration for the streaming source. Conflicts with `kinesis_stream`.
            See Kinesis Firehose below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Firehose delivery stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `kinesisStream` (`pulumi.Input[dict]`) - The Kinesis Stream configuration for the streaming source. Conflicts with `kinesis_firehose`.
            See Kinesis Stream below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `name_prefix` (`pulumi.Input[str]`) - The Name Prefix to use when creating an in-application stream.
          * `parallelism` (`pulumi.Input[dict]`) - The number of Parallel in-application streams to create.
            See Parallelism below for more details.
            * `count` (`pulumi.Input[float]`) - The Count of streams.

          * `processingConfiguration` (`pulumi.Input[dict]`) - The Processing Configuration to transform records as they are received from the stream.
            See Processing Configuration below for more details.
            * `lambda_` (`pulumi.Input[dict]`) - The Lambda function configuration. See Lambda below for more details.
              * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Lambda function.
              * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the Lambda function.

          * `schema` (`pulumi.Input[dict]`) - The Schema format of the data in the streaming source. See Source Schema below for more details.
            * `recordColumns` (`pulumi.Input[list]`) - The Record Column mapping for the streaming source data element.
              See Record Columns below for more details.
              * `mapping` (`pulumi.Input[str]`) - The Mapping reference to the data element.
              * `name` (`pulumi.Input[str]`) - Name of the column.
              * `sqlType` (`pulumi.Input[str]`) - The SQL Type of the column.

            * `recordEncoding` (`pulumi.Input[str]`) - The Encoding of the record in the streaming source.
            * `recordFormat` (`pulumi.Input[dict]`) - The Record Format and mapping information to schematize a record.
              See Record Format below for more details.
              * `mappingParameters` (`pulumi.Input[dict]`) - The Mapping Information for the record format.
                See Mapping Parameters below for more details.
                * `csv` (`pulumi.Input[dict]`) - Mapping information when the record format uses delimiters.
                  See CSV Mapping Parameters below for more details.
                  * `recordColumnDelimiter` (`pulumi.Input[str]`) - The Column Delimiter.
                  * `recordRowDelimiter` (`pulumi.Input[str]`) - The Row Delimiter.

                * `json` (`pulumi.Input[dict]`) - Mapping information when JSON is the record format on the streaming source.
                  See JSON Mapping Parameters below for more details.
                  * `recordRowPath` (`pulumi.Input[str]`) - Path to the top-level parent that contains the records.

              * `recordFormatType` (`pulumi.Input[str]`) - The type of Record Format. Can be `CSV` or `JSON`.

          * `startingPositionConfigurations` (`pulumi.Input[list]`)
            * `starting_position` (`pulumi.Input[str]`)

          * `streamNames` (`pulumi.Input[list]`)

        The **outputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`) - The Kinesis Firehose configuration for the destination stream. Conflicts with `kinesis_stream`.
            See Kinesis Firehose below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Firehose delivery stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `kinesisStream` (`pulumi.Input[dict]`) - The Kinesis Stream configuration for the destination stream. Conflicts with `kinesis_firehose`.
            See Kinesis Stream below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Kinesis Stream.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the stream.

          * `lambda_` (`pulumi.Input[dict]`) - The Lambda function destination. See Lambda below for more details.
            * `resource_arn` (`pulumi.Input[str]`) - The ARN of the Lambda function.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to access the Lambda function.

          * `name` (`pulumi.Input[str]`) - The Name of the in-application stream.
          * `schema` (`pulumi.Input[dict]`) - The Schema format of the data written to the destination. See Destination Schema below for more details.
            * `recordFormatType` (`pulumi.Input[str]`) - The Format Type of the records on the output stream. Can be `CSV` or `JSON`.

        The **reference_data_sources** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `s3` (`pulumi.Input[dict]`) - The S3 configuration for the reference data source. See S3 Reference below for more details.
            * `bucketArn` (`pulumi.Input[str]`) - The S3 Bucket ARN.
            * `fileKey` (`pulumi.Input[str]`) - The File Key name containing reference data.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM Role used to send application messages.

          * `schema` (`pulumi.Input[dict]`) - The Schema format of the data in the streaming source. See Source Schema below for more details.
            * `recordColumns` (`pulumi.Input[list]`) - The Record Column mapping for the streaming source data element.
              See Record Columns below for more details.
              * `mapping` (`pulumi.Input[str]`) - The Mapping reference to the data element.
              * `name` (`pulumi.Input[str]`) - Name of the column.
              * `sqlType` (`pulumi.Input[str]`) - The SQL Type of the column.

            * `recordEncoding` (`pulumi.Input[str]`) - The Encoding of the record in the streaming source.
            * `recordFormat` (`pulumi.Input[dict]`) - The Record Format and mapping information to schematize a record.
              See Record Format below for more details.
              * `mappingParameters` (`pulumi.Input[dict]`) - The Mapping Information for the record format.
                See Mapping Parameters below for more details.
                * `csv` (`pulumi.Input[dict]`) - Mapping information when the record format uses delimiters.
                  See CSV Mapping Parameters below for more details.
                  * `recordColumnDelimiter` (`pulumi.Input[str]`) - The Column Delimiter.
                  * `recordRowDelimiter` (`pulumi.Input[str]`) - The Row Delimiter.

                * `json` (`pulumi.Input[dict]`) - Mapping information when JSON is the record format on the streaming source.
                  See JSON Mapping Parameters below for more details.
                  * `recordRowPath` (`pulumi.Input[str]`) - Path to the top-level parent that contains the records.

              * `recordFormatType` (`pulumi.Input[str]`) - The type of Record Format. Can be `CSV` or `JSON`.

          * `table_name` (`pulumi.Input[str]`) - The in-application Table Name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["cloudwatch_logging_options"] = cloudwatch_logging_options
        __props__["code"] = code
        __props__["create_timestamp"] = create_timestamp
        __props__["description"] = description
        __props__["inputs"] = inputs
        __props__["last_update_timestamp"] = last_update_timestamp
        __props__["name"] = name
        __props__["outputs"] = outputs
        __props__["reference_data_sources"] = reference_data_sources
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["version"] = version
        return AnalyticsApplication(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
