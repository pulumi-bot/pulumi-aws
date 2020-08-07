# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'ClusterBrokerNodeGroupInfo',
    'ClusterClientAuthentication',
    'ClusterClientAuthenticationTls',
    'ClusterConfigurationInfo',
    'ClusterEncryptionInfo',
    'ClusterEncryptionInfoEncryptionInTransit',
    'ClusterLoggingInfo',
    'ClusterLoggingInfoBrokerLogs',
    'ClusterLoggingInfoBrokerLogsCloudwatchLogs',
    'ClusterLoggingInfoBrokerLogsFirehose',
    'ClusterLoggingInfoBrokerLogsS3',
    'ClusterOpenMonitoring',
    'ClusterOpenMonitoringPrometheus',
    'ClusterOpenMonitoringPrometheusJmxExporter',
    'ClusterOpenMonitoringPrometheusNodeExporter',
]

@pulumi.output_type
class ClusterBrokerNodeGroupInfo(dict):
    @property
    @pulumi.getter(name="azDistribution")
    def az_distribution(self) -> Optional[str]:
        """
        The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
        """
        ...

    @property
    @pulumi.getter(name="clientSubnets")
    def client_subnets(self) -> List[str]:
        """
        A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
        """
        ...

    @property
    @pulumi.getter(name="ebsVolumeSize")
    def ebs_volume_size(self) -> float:
        """
        The size in GiB of the EBS volume for the data drive on each broker node.
        """
        ...

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        Specify the instance type to use for the kafka brokers. e.g. kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
        """
        ...

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> List[str]:
        """
        A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterClientAuthentication(dict):
    @property
    @pulumi.getter
    def tls(self) -> Optional['outputs.ClusterClientAuthenticationTls']:
        """
        Configuration block for specifying TLS client authentication. See below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterClientAuthenticationTls(dict):
    @property
    @pulumi.getter(name="certificateAuthorityArns")
    def certificate_authority_arns(self) -> Optional[List[str]]:
        """
        List of ACM Certificate Authority Amazon Resource Names (ARNs).
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterConfigurationInfo(dict):
    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
        """
        ...

    @property
    @pulumi.getter
    def revision(self) -> float:
        """
        Revision of the MSK Configuration to use in the cluster.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterEncryptionInfo(dict):
    @property
    @pulumi.getter(name="encryptionAtRestKmsKeyArn")
    def encryption_at_rest_kms_key_arn(self) -> Optional[str]:
        """
        You may specify a KMS key short ID or ARN (it will always output an ARN) to use for encrypting your data at rest.  If no key is specified, an AWS managed KMS ('aws/msk' managed service) key will be used for encrypting the data at rest.
        """
        ...

    @property
    @pulumi.getter(name="encryptionInTransit")
    def encryption_in_transit(self) -> Optional['outputs.ClusterEncryptionInfoEncryptionInTransit']:
        """
        Configuration block to specify encryption in transit. See below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterEncryptionInfoEncryptionInTransit(dict):
    @property
    @pulumi.getter(name="clientBroker")
    def client_broker(self) -> Optional[str]:
        """
        Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value is `TLS_PLAINTEXT` when `encryption_in_transit` block defined, but `TLS` when `encryption_in_transit` block omitted.
        """
        ...

    @property
    @pulumi.getter(name="inCluster")
    def in_cluster(self) -> Optional[bool]:
        """
        Whether data communication among broker nodes is encrypted. Default value: `true`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterLoggingInfo(dict):
    @property
    @pulumi.getter(name="brokerLogs")
    def broker_logs(self) -> 'outputs.ClusterLoggingInfoBrokerLogs':
        """
        Configuration block for Broker Logs settings for logging info. See below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterLoggingInfoBrokerLogs(dict):
    @property
    @pulumi.getter(name="cloudwatchLogs")
    def cloudwatch_logs(self) -> Optional['outputs.ClusterLoggingInfoBrokerLogsCloudwatchLogs']:
        ...

    @property
    @pulumi.getter
    def firehose(self) -> Optional['outputs.ClusterLoggingInfoBrokerLogsFirehose']:
        ...

    @property
    @pulumi.getter
    def s3(self) -> Optional['outputs.ClusterLoggingInfoBrokerLogsS3']:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterLoggingInfoBrokerLogsCloudwatchLogs(dict):
    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Indicates whether you want to enable or disable streaming broker logs to Cloudwatch Logs.
        """
        ...

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> Optional[str]:
        """
        Name of the Cloudwatch Log Group to deliver logs to.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterLoggingInfoBrokerLogsFirehose(dict):
    @property
    @pulumi.getter(name="deliveryStream")
    def delivery_stream(self) -> Optional[str]:
        """
        Name of the Kinesis Data Firehose delivery stream to deliver logs to.
        """
        ...

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Indicates whether you want to enable or disable streaming broker logs to Cloudwatch Logs.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterLoggingInfoBrokerLogsS3(dict):
    @property
    @pulumi.getter
    def bucket(self) -> Optional[str]:
        """
        Name of the S3 bucket to deliver logs to.
        """
        ...

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Indicates whether you want to enable or disable streaming broker logs to Cloudwatch Logs.
        """
        ...

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        """
        Prefix to append to the folder name.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterOpenMonitoring(dict):
    @property
    @pulumi.getter
    def prometheus(self) -> 'outputs.ClusterOpenMonitoringPrometheus':
        """
        Configuration block for Prometheus settings for open monitoring. See below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterOpenMonitoringPrometheus(dict):
    @property
    @pulumi.getter(name="jmxExporter")
    def jmx_exporter(self) -> Optional['outputs.ClusterOpenMonitoringPrometheusJmxExporter']:
        """
        Configuration block for JMX Exporter. See below.
        """
        ...

    @property
    @pulumi.getter(name="nodeExporter")
    def node_exporter(self) -> Optional['outputs.ClusterOpenMonitoringPrometheusNodeExporter']:
        """
        Configuration block for Node Exporter. See below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterOpenMonitoringPrometheusJmxExporter(dict):
    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> bool:
        """
        Indicates whether you want to enable or disable the JMX Exporter.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterOpenMonitoringPrometheusNodeExporter(dict):
    @property
    @pulumi.getter(name="enabledInBroker")
    def enabled_in_broker(self) -> bool:
        """
        Indicates whether you want to enable or disable the JMX Exporter.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


