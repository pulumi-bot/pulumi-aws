# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Application']


class Application(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_configuration: Optional[pulumi.Input[pulumi.InputType['ApplicationApplicationConfigurationArgs']]] = None,
                 cloudwatch_logging_options: Optional[pulumi.Input[pulumi.InputType['ApplicationCloudwatchLoggingOptionsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 runtime_environment: Optional[pulumi.Input[str]] = None,
                 service_execution_role: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Kinesis Analytics v2 Application.
        This resource can be used to manage both Kinesis Data Analytics for SQL applications and Kinesis Data Analytics for Apache Flink applications.

        > **Note:** Kinesis Data Analytics for SQL applications created using this resource cannot currently be viewed in the AWS Console. To manage Kinesis Data Analytics for SQL applications that can also be viewed in the AWS Console, use the [`kinesis.AnalyticsApplication`](https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application.html) resource.

        ## Example Usage
        ### Apache Flink Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        example_bucket_object = aws.s3.BucketObject("exampleBucketObject",
            bucket=example_bucket.bucket,
            key="example-flink-application",
            source=pulumi.FileAsset("flink-app.jar"))
        example_application = aws.kinesisanalyticsv2.Application("exampleApplication",
            runtime_environment="FLINK-1_8",
            service_execution_role=aws_iam_role["example"]["arn"],
            application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationArgs(
                application_code_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs(
                    code_content=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs(
                        s3_content_location=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs(
                            bucket_arn=example_bucket.arn,
                            file_key=example_bucket_object.key,
                        ),
                    ),
                    code_content_type="ZIPFILE",
                ),
                environment_properties=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesArgs(
                    property_groups=[
                        aws.kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs(
                            property_group_id="PROPERTY-GROUP-1",
                            property_map={
                                "Key1": "Value1",
                            },
                        ),
                        aws.kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs(
                            property_group_id="PROPERTY-GROUP-2",
                            property_map={
                                "KeyA": "ValueA",
                                "KeyB": "ValueB",
                            },
                        ),
                    ],
                ),
                flink_application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs(
                    checkpoint_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs(
                        configuration_type="DEFAULT",
                    ),
                    monitoring_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs(
                        configuration_type="CUSTOM",
                        log_level="DEBUG",
                        metrics_level="TASK",
                    ),
                    parallelism_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs(
                        auto_scaling_enabled=True,
                        configuration_type="CUSTOM",
                        parallelism=10,
                        parallelism_per_kpu=4,
                    ),
                ),
            ),
            tags={
                "Environment": "test",
            })
        ```
        ### SQL Application

        ```python
        import pulumi
        import pulumi_aws as aws

        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
        example_log_stream = aws.cloudwatch.LogStream("exampleLogStream", log_group_name=example_log_group.name)
        example_application = aws.kinesisanalyticsv2.Application("exampleApplication",
            runtime_environment="SQL-1.0",
            service_execution_role=aws_iam_role["example"]["arn"],
            application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationArgs(
                application_code_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs(
                    code_content=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs(
                        text_content="SELECT 1;\n",
                    ),
                    code_content_type="PLAINTEXT",
                ),
                sql_application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationArgs(
                    input=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputArgs(
                        name_prefix="PREFIX_1",
                        input_parallelism=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs(
                            count=3,
                        ),
                        input_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs(
                            record_columns=[
                                aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs(
                                    name="COLUMN_1",
                                    sql_type="VARCHAR(8)",
                                    mapping="MAPPING-1",
                                ),
                                aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs(
                                    name="COLUMN_2",
                                    sql_type="DOUBLE",
                                ),
                            ],
                            record_encoding="UTF-8",
                            record_format=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs(
                                record_format_type="CSV",
                                mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs(
                                    csv_mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs(
                                        record_column_delimiter=",",
                                        record_row_delimiter="\n",
                                    ),
                                ),
                            ),
                        ),
                        kinesis_streams_input=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs(
                            resource_arn=aws_kinesis_stream["example"]["arn"],
                        ),
                    ),
                    outputs=[
                        aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs(
                            name="OUTPUT_1",
                            destination_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs(
                                record_format_type="JSON",
                            ),
                            lambda_output=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs(
                                resource_arn=aws_lambda_function["example"]["arn"],
                            ),
                        ),
                        aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs(
                            name="OUTPUT_2",
                            destination_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs(
                                record_format_type="CSV",
                            ),
                            kinesis_firehose_output=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs(
                                resource_arn=aws_kinesis_firehose_delivery_stream["example"]["arn"],
                            ),
                        ),
                    ],
                    reference_data_source=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs(
                        table_name="TABLE-1",
                        reference_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs(
                            record_columns=[aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArgs(
                                name="COLUMN_1",
                                sql_type="INTEGER",
                            )],
                            record_format=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs(
                                record_format_type="JSON",
                                mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs(
                                    json_mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs(
                                        record_row_path="$",
                                    ),
                                ),
                            ),
                        ),
                        s3_reference_data_source=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs(
                            bucket_arn=aws_s3_bucket["example"]["arn"],
                            file_key="KEY-1",
                        ),
                    ),
                ),
            ),
            cloudwatch_logging_options=aws.kinesisanalyticsv2.ApplicationCloudwatchLoggingOptionsArgs(
                log_stream_arn=example_log_stream.arn,
            ))
        ```
        ### VPC Configuration

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        example_bucket_object = aws.s3.BucketObject("exampleBucketObject",
            bucket=example_bucket.bucket,
            key="example-flink-application",
            source=pulumi.FileAsset("flink-app.jar"))
        example_application = aws.kinesisanalyticsv2.Application("exampleApplication",
            runtime_environment="FLINK-1_8",
            service_execution_role=aws_iam_role["example"]["arn"],
            application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationArgs(
                application_code_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs(
                    code_content=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs(
                        s3_content_location=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs(
                            bucket_arn=example_bucket.arn,
                            file_key=example_bucket_object.key,
                        ),
                    ),
                    code_content_type="ZIPFILE",
                ),
                vpc_configuration={
                    "security_group_ids": [
                        aws_security_group["example"][0]["id"],
                        aws_security_group["example"][1]["id"],
                    ],
                    "subnet_ids": [aws_subnet["example"]["id"]],
                },
            ))
        ```

        ## Import

        `aws_kinesisanalyticsv2_application` can be imported by using the application ARN, e.g.

        ```sh
         $ pulumi import aws:kinesisanalyticsv2/application:Application example arn:aws:kinesisanalytics:us-west-2:123456789012:application/example-sql-application
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ApplicationApplicationConfigurationArgs']] application_configuration: The application's configuration
        :param pulumi.Input[pulumi.InputType['ApplicationCloudwatchLoggingOptionsArgs']] cloudwatch_logging_options: A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
        :param pulumi.Input[str] description: A summary description of the application.
        :param pulumi.Input[str] name: The name of the application.
        :param pulumi.Input[str] runtime_environment: The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
        :param pulumi.Input[str] service_execution_role: The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the application.
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

            __props__['application_configuration'] = application_configuration
            __props__['cloudwatch_logging_options'] = cloudwatch_logging_options
            __props__['description'] = description
            __props__['name'] = name
            if runtime_environment is None and not opts.urn:
                raise TypeError("Missing required property 'runtime_environment'")
            __props__['runtime_environment'] = runtime_environment
            if service_execution_role is None and not opts.urn:
                raise TypeError("Missing required property 'service_execution_role'")
            __props__['service_execution_role'] = service_execution_role
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['create_timestamp'] = None
            __props__['last_update_timestamp'] = None
            __props__['status'] = None
            __props__['version_id'] = None
        super(Application, __self__).__init__(
            'aws:kinesisanalyticsv2/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_configuration: Optional[pulumi.Input[pulumi.InputType['ApplicationApplicationConfigurationArgs']]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cloudwatch_logging_options: Optional[pulumi.Input[pulumi.InputType['ApplicationCloudwatchLoggingOptionsArgs']]] = None,
            create_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            last_update_timestamp: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            runtime_environment: Optional[pulumi.Input[str]] = None,
            service_execution_role: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version_id: Optional[pulumi.Input[int]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ApplicationApplicationConfigurationArgs']] application_configuration: The application's configuration
        :param pulumi.Input[str] arn: The ARN of the application.
        :param pulumi.Input[pulumi.InputType['ApplicationCloudwatchLoggingOptionsArgs']] cloudwatch_logging_options: A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
        :param pulumi.Input[str] create_timestamp: The current timestamp when the application was created.
        :param pulumi.Input[str] description: A summary description of the application.
        :param pulumi.Input[str] last_update_timestamp: The current timestamp when the application was last updated.
        :param pulumi.Input[str] name: The name of the application.
        :param pulumi.Input[str] runtime_environment: The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
        :param pulumi.Input[str] service_execution_role: The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
        :param pulumi.Input[str] status: The status of the application.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the application.
        :param pulumi.Input[int] version_id: The current application version. Kinesis Data Analytics updates the `version_id` each time the application is updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_configuration"] = application_configuration
        __props__["arn"] = arn
        __props__["cloudwatch_logging_options"] = cloudwatch_logging_options
        __props__["create_timestamp"] = create_timestamp
        __props__["description"] = description
        __props__["last_update_timestamp"] = last_update_timestamp
        __props__["name"] = name
        __props__["runtime_environment"] = runtime_environment
        __props__["service_execution_role"] = service_execution_role
        __props__["status"] = status
        __props__["tags"] = tags
        __props__["version_id"] = version_id
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationConfiguration")
    def application_configuration(self) -> pulumi.Output['outputs.ApplicationApplicationConfiguration']:
        """
        The application's configuration
        """
        return pulumi.get(self, "application_configuration")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the application.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cloudwatchLoggingOptions")
    def cloudwatch_logging_options(self) -> pulumi.Output[Optional['outputs.ApplicationCloudwatchLoggingOptions']]:
        """
        A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
        """
        return pulumi.get(self, "cloudwatch_logging_options")

    @property
    @pulumi.getter(name="createTimestamp")
    def create_timestamp(self) -> pulumi.Output[str]:
        """
        The current timestamp when the application was created.
        """
        return pulumi.get(self, "create_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A summary description of the application.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastUpdateTimestamp")
    def last_update_timestamp(self) -> pulumi.Output[str]:
        """
        The current timestamp when the application was last updated.
        """
        return pulumi.get(self, "last_update_timestamp")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="runtimeEnvironment")
    def runtime_environment(self) -> pulumi.Output[str]:
        """
        The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
        """
        return pulumi.get(self, "runtime_environment")

    @property
    @pulumi.getter(name="serviceExecutionRole")
    def service_execution_role(self) -> pulumi.Output[str]:
        """
        The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
        """
        return pulumi.get(self, "service_execution_role")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the application.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the application.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[int]:
        """
        The current application version. Kinesis Data Analytics updates the `version_id` each time the application is updated.
        """
        return pulumi.get(self, "version_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

