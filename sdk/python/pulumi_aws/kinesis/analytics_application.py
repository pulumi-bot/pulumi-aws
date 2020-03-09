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
      * `logStreamArn` (`str`)
      * `role_arn` (`str`)
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
      * `kinesisFirehose` (`dict`)
        * `resource_arn` (`str`)
        * `role_arn` (`str`)

      * `kinesisStream` (`dict`)
        * `resource_arn` (`str`)
        * `role_arn` (`str`)

      * `name_prefix` (`str`)
      * `parallelism` (`dict`)
        * `count` (`float`)

      * `processingConfiguration` (`dict`)
        * `lambda` (`dict`)
          * `resource_arn` (`str`)
          * `role_arn` (`str`)

      * `schema` (`dict`)
        * `recordColumns` (`list`)
          * `mapping` (`str`)
          * `name` (`str`) - Name of the Kinesis Analytics Application.
          * `sqlType` (`str`)

        * `recordEncoding` (`str`)
        * `recordFormat` (`dict`)
          * `mappingParameters` (`dict`)
            * `csv` (`dict`)
              * `recordColumnDelimiter` (`str`)
              * `recordRowDelimiter` (`str`)

            * `json` (`dict`)
              * `recordRowPath` (`str`)

          * `recordFormatType` (`str`)

      * `startingPositionConfigurations` (`list`)
        * `startingPosition` (`str`)

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
      * `kinesisFirehose` (`dict`)
        * `resource_arn` (`str`)
        * `role_arn` (`str`)

      * `kinesisStream` (`dict`)
        * `resource_arn` (`str`)
        * `role_arn` (`str`)

      * `lambda` (`dict`)
        * `resource_arn` (`str`)
        * `role_arn` (`str`)

      * `name` (`str`) - Name of the Kinesis Analytics Application.
      * `schema` (`dict`)
        * `recordFormatType` (`str`)
    """
    reference_data_sources: pulumi.Output[dict]
    """
    An S3 Reference Data Source for the application.
    See Reference Data Sources below for more details.

      * `id` (`str`) - The ARN of the Kinesis Analytics Application.
      * `s3` (`dict`)
        * `bucketArn` (`str`)
        * `fileKey` (`str`)
        * `role_arn` (`str`)

      * `schema` (`dict`)
        * `recordColumns` (`list`)
          * `mapping` (`str`)
          * `name` (`str`) - Name of the Kinesis Analytics Application.
          * `sqlType` (`str`)

        * `recordEncoding` (`str`)
        * `recordFormat` (`dict`)
          * `mappingParameters` (`dict`)
            * `csv` (`dict`)
              * `recordColumnDelimiter` (`str`)
              * `recordRowDelimiter` (`str`)

            * `json` (`dict`)
              * `recordRowPath` (`str`)

          * `recordFormatType` (`str`)

      * `table_name` (`str`)
    """
    status: pulumi.Output[str]
    """
    The Status of the application.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of tags for the Kinesis Analytics Application.
    """
    version: pulumi.Output[float]
    """
    The Version of the application.
    """
    def __init__(__self__, resource_name, opts=None, cloudwatch_logging_options=None, code=None, description=None, inputs=None, name=None, outputs=None, reference_data_sources=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
        allows processing and analyzing streaming data using standard SQL.

        For more details, see the [Amazon Kinesis Analytics Documentation][1].

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kinesis_analytics_application.html.markdown.

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
        :param pulumi.Input[dict] tags: Key-value mapping of tags for the Kinesis Analytics Application.

        The **cloudwatch_logging_options** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `logStreamArn` (`pulumi.Input[str]`)
          * `role_arn` (`pulumi.Input[str]`)

        The **inputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `kinesisStream` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `name_prefix` (`pulumi.Input[str]`)
          * `parallelism` (`pulumi.Input[dict]`)
            * `count` (`pulumi.Input[float]`)

          * `processingConfiguration` (`pulumi.Input[dict]`)
            * `lambda` (`pulumi.Input[dict]`)
              * `resource_arn` (`pulumi.Input[str]`)
              * `role_arn` (`pulumi.Input[str]`)

          * `schema` (`pulumi.Input[dict]`)
            * `recordColumns` (`pulumi.Input[list]`)
              * `mapping` (`pulumi.Input[str]`)
              * `name` (`pulumi.Input[str]`) - Name of the Kinesis Analytics Application.
              * `sqlType` (`pulumi.Input[str]`)

            * `recordEncoding` (`pulumi.Input[str]`)
            * `recordFormat` (`pulumi.Input[dict]`)
              * `mappingParameters` (`pulumi.Input[dict]`)
                * `csv` (`pulumi.Input[dict]`)
                  * `recordColumnDelimiter` (`pulumi.Input[str]`)
                  * `recordRowDelimiter` (`pulumi.Input[str]`)

                * `json` (`pulumi.Input[dict]`)
                  * `recordRowPath` (`pulumi.Input[str]`)

              * `recordFormatType` (`pulumi.Input[str]`)

          * `startingPositionConfigurations` (`pulumi.Input[list]`)
            * `startingPosition` (`pulumi.Input[str]`)

          * `streamNames` (`pulumi.Input[list]`)

        The **outputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `kinesisStream` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `lambda` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `name` (`pulumi.Input[str]`) - Name of the Kinesis Analytics Application.
          * `schema` (`pulumi.Input[dict]`)
            * `recordFormatType` (`pulumi.Input[str]`)

        The **reference_data_sources** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `s3` (`pulumi.Input[dict]`)
            * `bucketArn` (`pulumi.Input[str]`)
            * `fileKey` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `schema` (`pulumi.Input[dict]`)
            * `recordColumns` (`pulumi.Input[list]`)
              * `mapping` (`pulumi.Input[str]`)
              * `name` (`pulumi.Input[str]`) - Name of the Kinesis Analytics Application.
              * `sqlType` (`pulumi.Input[str]`)

            * `recordEncoding` (`pulumi.Input[str]`)
            * `recordFormat` (`pulumi.Input[dict]`)
              * `mappingParameters` (`pulumi.Input[dict]`)
                * `csv` (`pulumi.Input[dict]`)
                  * `recordColumnDelimiter` (`pulumi.Input[str]`)
                  * `recordRowDelimiter` (`pulumi.Input[str]`)

                * `json` (`pulumi.Input[dict]`)
                  * `recordRowPath` (`pulumi.Input[str]`)

              * `recordFormatType` (`pulumi.Input[str]`)

          * `table_name` (`pulumi.Input[str]`)
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
        :param pulumi.Input[dict] tags: Key-value mapping of tags for the Kinesis Analytics Application.
        :param pulumi.Input[float] version: The Version of the application.

        The **cloudwatch_logging_options** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `logStreamArn` (`pulumi.Input[str]`)
          * `role_arn` (`pulumi.Input[str]`)

        The **inputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `kinesisStream` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `name_prefix` (`pulumi.Input[str]`)
          * `parallelism` (`pulumi.Input[dict]`)
            * `count` (`pulumi.Input[float]`)

          * `processingConfiguration` (`pulumi.Input[dict]`)
            * `lambda` (`pulumi.Input[dict]`)
              * `resource_arn` (`pulumi.Input[str]`)
              * `role_arn` (`pulumi.Input[str]`)

          * `schema` (`pulumi.Input[dict]`)
            * `recordColumns` (`pulumi.Input[list]`)
              * `mapping` (`pulumi.Input[str]`)
              * `name` (`pulumi.Input[str]`) - Name of the Kinesis Analytics Application.
              * `sqlType` (`pulumi.Input[str]`)

            * `recordEncoding` (`pulumi.Input[str]`)
            * `recordFormat` (`pulumi.Input[dict]`)
              * `mappingParameters` (`pulumi.Input[dict]`)
                * `csv` (`pulumi.Input[dict]`)
                  * `recordColumnDelimiter` (`pulumi.Input[str]`)
                  * `recordRowDelimiter` (`pulumi.Input[str]`)

                * `json` (`pulumi.Input[dict]`)
                  * `recordRowPath` (`pulumi.Input[str]`)

              * `recordFormatType` (`pulumi.Input[str]`)

          * `startingPositionConfigurations` (`pulumi.Input[list]`)
            * `startingPosition` (`pulumi.Input[str]`)

          * `streamNames` (`pulumi.Input[list]`)

        The **outputs** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `kinesisFirehose` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `kinesisStream` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `lambda` (`pulumi.Input[dict]`)
            * `resource_arn` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `name` (`pulumi.Input[str]`) - Name of the Kinesis Analytics Application.
          * `schema` (`pulumi.Input[dict]`)
            * `recordFormatType` (`pulumi.Input[str]`)

        The **reference_data_sources** object supports the following:

          * `id` (`pulumi.Input[str]`) - The ARN of the Kinesis Analytics Application.
          * `s3` (`pulumi.Input[dict]`)
            * `bucketArn` (`pulumi.Input[str]`)
            * `fileKey` (`pulumi.Input[str]`)
            * `role_arn` (`pulumi.Input[str]`)

          * `schema` (`pulumi.Input[dict]`)
            * `recordColumns` (`pulumi.Input[list]`)
              * `mapping` (`pulumi.Input[str]`)
              * `name` (`pulumi.Input[str]`) - Name of the Kinesis Analytics Application.
              * `sqlType` (`pulumi.Input[str]`)

            * `recordEncoding` (`pulumi.Input[str]`)
            * `recordFormat` (`pulumi.Input[dict]`)
              * `mappingParameters` (`pulumi.Input[dict]`)
                * `csv` (`pulumi.Input[dict]`)
                  * `recordColumnDelimiter` (`pulumi.Input[str]`)
                  * `recordRowDelimiter` (`pulumi.Input[str]`)

                * `json` (`pulumi.Input[dict]`)
                  * `recordRowPath` (`pulumi.Input[str]`)

              * `recordFormatType` (`pulumi.Input[str]`)

          * `table_name` (`pulumi.Input[str]`)
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

