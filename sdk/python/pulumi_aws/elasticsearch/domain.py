# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Domain(pulumi.CustomResource):
    access_policies: pulumi.Output[str]
    """
    IAM policy document specifying the access policies for the domain
    """
    advanced_options: pulumi.Output[dict]
    """
    Key-value string pairs to specify advanced configuration options.
    Note that the values for these configuration options must be strings (wrapped in quotes) or they
    may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
    domain on every apply.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the domain.
    """
    cluster_config: pulumi.Output[dict]
    """
    Cluster configuration of the domain, see below.

      * `dedicatedMasterCount` (`float`) - Number of dedicated master nodes in the cluster
      * `dedicatedMasterEnabled` (`bool`) - Indicates whether dedicated master nodes are enabled for the cluster.
      * `dedicatedMasterType` (`str`) - Instance type of the dedicated master nodes in the cluster.
      * `instanceCount` (`float`) - Number of instances in the cluster.
      * `instance_type` (`str`) - Instance type of data nodes in the cluster.
      * `zoneAwarenessConfig` (`dict`) - Configuration block containing zone awareness settings. Documented below.
        * `availabilityZoneCount` (`float`) - Number of Availability Zones for the domain to use with `zone_awareness_enabled`. Defaults to `2`. Valid values: `2` or `3`.

      * `zoneAwarenessEnabled` (`bool`) - Indicates whether zone awareness is enabled. To enable awareness with three Availability Zones, the `availability_zone_count` within the `zone_awareness_config` must be set to `3`.
    """
    cognito_options: pulumi.Output[dict]
    domain_endpoint_options: pulumi.Output[dict]
    """
    Domain endpoint HTTP(S) related options. See below.

      * `enforceHttps` (`bool`) - Whether or not to require HTTPS
      * `tlsSecurityPolicy` (`str`)
    """
    domain_id: pulumi.Output[str]
    """
    Unique identifier for the domain.
    """
    domain_name: pulumi.Output[str]
    """
    Name of the domain.
    """
    ebs_options: pulumi.Output[dict]
    """
    EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.

      * `ebsEnabled` (`bool`) - Whether EBS volumes are attached to data nodes in the domain.
      * `iops` (`float`) - The baseline input/output (I/O) performance of EBS volumes
        attached to data nodes. Applicable only for the Provisioned IOPS EBS volume type.
      * `volume_size` (`float`) - The size of EBS volumes attached to data nodes (in GB).
        **Required** if `ebs_enabled` is set to `true`.
      * `volumeType` (`str`) - The type of EBS volumes attached to data nodes.
    """
    elasticsearch_version: pulumi.Output[str]
    """
    The version of Elasticsearch to deploy. Defaults to `1.5`
    """
    encrypt_at_rest: pulumi.Output[dict]
    """
    Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.

      * `enabled` (`bool`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
      * `kms_key_id` (`str`) - The KMS key id to encrypt the Elasticsearch domain with. If not specified then it defaults to using the `aws/es` service KMS key.
    """
    endpoint: pulumi.Output[str]
    """
    Domain-specific endpoint used to submit index, search, and data upload requests.
    """
    kibana_endpoint: pulumi.Output[str]
    """
    Domain-specific endpoint for kibana without https scheme.
    * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
    * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
    """
    log_publishing_options: pulumi.Output[list]
    """
    Options for publishing slow logs to CloudWatch Logs.

      * `cloudwatch_log_group_arn` (`str`) - ARN of the Cloudwatch log group to which log needs to be published.
      * `enabled` (`bool`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
      * `logType` (`str`) - A type of Elasticsearch log. Valid values: INDEX_SLOW_LOGS, SEARCH_SLOW_LOGS, ES_APPLICATION_LOGS
    """
    node_to_node_encryption: pulumi.Output[dict]
    """
    Node-to-node encryption options. See below.

      * `enabled` (`bool`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
    """
    snapshot_options: pulumi.Output[dict]
    """
    Snapshot related options, see below.

      * `automatedSnapshotStartHour` (`float`) - Hour during which the service takes an automated daily
        snapshot of the indices in the domain.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource
    """
    vpc_options: pulumi.Output[dict]
    """
    VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).

      * `availability_zones` (`list`)
      * `security_group_ids` (`list`) - List of VPC Security Group IDs to be applied to the Elasticsearch domain endpoints. If omitted, the default Security Group for the VPC will be used.
      * `subnet_ids` (`list`) - List of VPC Subnet IDs for the Elasticsearch domain endpoints to be created in.
      * `vpc_id` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, access_policies=None, advanced_options=None, cluster_config=None, cognito_options=None, domain_endpoint_options=None, domain_name=None, ebs_options=None, elasticsearch_version=None, encrypt_at_rest=None, log_publishing_options=None, node_to_node_encryption=None, snapshot_options=None, tags=None, vpc_options=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an AWS Elasticsearch Domain.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticsearch_domain.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_policies: IAM policy document specifying the access policies for the domain
        :param pulumi.Input[dict] advanced_options: Key-value string pairs to specify advanced configuration options.
               Note that the values for these configuration options must be strings (wrapped in quotes) or they
               may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
               domain on every apply.
        :param pulumi.Input[dict] cluster_config: Cluster configuration of the domain, see below.
        :param pulumi.Input[dict] domain_endpoint_options: Domain endpoint HTTP(S) related options. See below.
        :param pulumi.Input[str] domain_name: Name of the domain.
        :param pulumi.Input[dict] ebs_options: EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
        :param pulumi.Input[str] elasticsearch_version: The version of Elasticsearch to deploy. Defaults to `1.5`
        :param pulumi.Input[dict] encrypt_at_rest: Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
        :param pulumi.Input[list] log_publishing_options: Options for publishing slow logs to CloudWatch Logs.
        :param pulumi.Input[dict] node_to_node_encryption: Node-to-node encryption options. See below.
        :param pulumi.Input[dict] snapshot_options: Snapshot related options, see below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource
        :param pulumi.Input[dict] vpc_options: VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).

        The **cluster_config** object supports the following:

          * `dedicatedMasterCount` (`pulumi.Input[float]`) - Number of dedicated master nodes in the cluster
          * `dedicatedMasterEnabled` (`pulumi.Input[bool]`) - Indicates whether dedicated master nodes are enabled for the cluster.
          * `dedicatedMasterType` (`pulumi.Input[str]`) - Instance type of the dedicated master nodes in the cluster.
          * `instanceCount` (`pulumi.Input[float]`) - Number of instances in the cluster.
          * `instance_type` (`pulumi.Input[str]`) - Instance type of data nodes in the cluster.
          * `zoneAwarenessConfig` (`pulumi.Input[dict]`) - Configuration block containing zone awareness settings. Documented below.
            * `availabilityZoneCount` (`pulumi.Input[float]`) - Number of Availability Zones for the domain to use with `zone_awareness_enabled`. Defaults to `2`. Valid values: `2` or `3`.

          * `zoneAwarenessEnabled` (`pulumi.Input[bool]`) - Indicates whether zone awareness is enabled. To enable awareness with three Availability Zones, the `availability_zone_count` within the `zone_awareness_config` must be set to `3`.

        The **cognito_options** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
          * `identity_pool_id` (`pulumi.Input[str]`) - ID of the Cognito Identity Pool to use
          * `role_arn` (`pulumi.Input[str]`) - ARN of the IAM role that has the AmazonESCognitoAccess policy attached
          * `user_pool_id` (`pulumi.Input[str]`) - ID of the Cognito User Pool to use

        The **domain_endpoint_options** object supports the following:

          * `enforceHttps` (`pulumi.Input[bool]`) - Whether or not to require HTTPS
          * `tlsSecurityPolicy` (`pulumi.Input[str]`)

        The **ebs_options** object supports the following:

          * `ebsEnabled` (`pulumi.Input[bool]`) - Whether EBS volumes are attached to data nodes in the domain.
          * `iops` (`pulumi.Input[float]`) - The baseline input/output (I/O) performance of EBS volumes
            attached to data nodes. Applicable only for the Provisioned IOPS EBS volume type.
          * `volume_size` (`pulumi.Input[float]`) - The size of EBS volumes attached to data nodes (in GB).
            **Required** if `ebs_enabled` is set to `true`.
          * `volumeType` (`pulumi.Input[str]`) - The type of EBS volumes attached to data nodes.

        The **encrypt_at_rest** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
          * `kms_key_id` (`pulumi.Input[str]`) - The KMS key id to encrypt the Elasticsearch domain with. If not specified then it defaults to using the `aws/es` service KMS key.

        The **log_publishing_options** object supports the following:

          * `cloudwatch_log_group_arn` (`pulumi.Input[str]`) - ARN of the Cloudwatch log group to which log needs to be published.
          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
          * `logType` (`pulumi.Input[str]`) - A type of Elasticsearch log. Valid values: INDEX_SLOW_LOGS, SEARCH_SLOW_LOGS, ES_APPLICATION_LOGS

        The **node_to_node_encryption** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not

        The **snapshot_options** object supports the following:

          * `automatedSnapshotStartHour` (`pulumi.Input[float]`) - Hour during which the service takes an automated daily
            snapshot of the indices in the domain.

        The **vpc_options** object supports the following:

          * `availability_zones` (`pulumi.Input[list]`)
          * `security_group_ids` (`pulumi.Input[list]`) - List of VPC Security Group IDs to be applied to the Elasticsearch domain endpoints. If omitted, the default Security Group for the VPC will be used.
          * `subnet_ids` (`pulumi.Input[list]`) - List of VPC Subnet IDs for the Elasticsearch domain endpoints to be created in.
          * `vpc_id` (`pulumi.Input[str]`)
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

            __props__['access_policies'] = access_policies
            __props__['advanced_options'] = advanced_options
            __props__['cluster_config'] = cluster_config
            __props__['cognito_options'] = cognito_options
            __props__['domain_endpoint_options'] = domain_endpoint_options
            __props__['domain_name'] = domain_name
            __props__['ebs_options'] = ebs_options
            __props__['elasticsearch_version'] = elasticsearch_version
            __props__['encrypt_at_rest'] = encrypt_at_rest
            __props__['log_publishing_options'] = log_publishing_options
            __props__['node_to_node_encryption'] = node_to_node_encryption
            __props__['snapshot_options'] = snapshot_options
            __props__['tags'] = tags
            __props__['vpc_options'] = vpc_options
            __props__['arn'] = None
            __props__['domain_id'] = None
            __props__['endpoint'] = None
            __props__['kibana_endpoint'] = None
        super(Domain, __self__).__init__(
            'aws:elasticsearch/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_policies=None, advanced_options=None, arn=None, cluster_config=None, cognito_options=None, domain_endpoint_options=None, domain_id=None, domain_name=None, ebs_options=None, elasticsearch_version=None, encrypt_at_rest=None, endpoint=None, kibana_endpoint=None, log_publishing_options=None, node_to_node_encryption=None, snapshot_options=None, tags=None, vpc_options=None):
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_policies: IAM policy document specifying the access policies for the domain
        :param pulumi.Input[dict] advanced_options: Key-value string pairs to specify advanced configuration options.
               Note that the values for these configuration options must be strings (wrapped in quotes) or they
               may be wrong and cause a perpetual diff, causing this provider to want to recreate your Elasticsearch
               domain on every apply.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the domain.
        :param pulumi.Input[dict] cluster_config: Cluster configuration of the domain, see below.
        :param pulumi.Input[dict] domain_endpoint_options: Domain endpoint HTTP(S) related options. See below.
        :param pulumi.Input[str] domain_id: Unique identifier for the domain.
        :param pulumi.Input[str] domain_name: Name of the domain.
        :param pulumi.Input[dict] ebs_options: EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
        :param pulumi.Input[str] elasticsearch_version: The version of Elasticsearch to deploy. Defaults to `1.5`
        :param pulumi.Input[dict] encrypt_at_rest: Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
        :param pulumi.Input[str] endpoint: Domain-specific endpoint used to submit index, search, and data upload requests.
        :param pulumi.Input[str] kibana_endpoint: Domain-specific endpoint for kibana without https scheme.
               * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
               * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
        :param pulumi.Input[list] log_publishing_options: Options for publishing slow logs to CloudWatch Logs.
        :param pulumi.Input[dict] node_to_node_encryption: Node-to-node encryption options. See below.
        :param pulumi.Input[dict] snapshot_options: Snapshot related options, see below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource
        :param pulumi.Input[dict] vpc_options: VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).

        The **cluster_config** object supports the following:

          * `dedicatedMasterCount` (`pulumi.Input[float]`) - Number of dedicated master nodes in the cluster
          * `dedicatedMasterEnabled` (`pulumi.Input[bool]`) - Indicates whether dedicated master nodes are enabled for the cluster.
          * `dedicatedMasterType` (`pulumi.Input[str]`) - Instance type of the dedicated master nodes in the cluster.
          * `instanceCount` (`pulumi.Input[float]`) - Number of instances in the cluster.
          * `instance_type` (`pulumi.Input[str]`) - Instance type of data nodes in the cluster.
          * `zoneAwarenessConfig` (`pulumi.Input[dict]`) - Configuration block containing zone awareness settings. Documented below.
            * `availabilityZoneCount` (`pulumi.Input[float]`) - Number of Availability Zones for the domain to use with `zone_awareness_enabled`. Defaults to `2`. Valid values: `2` or `3`.

          * `zoneAwarenessEnabled` (`pulumi.Input[bool]`) - Indicates whether zone awareness is enabled. To enable awareness with three Availability Zones, the `availability_zone_count` within the `zone_awareness_config` must be set to `3`.

        The **cognito_options** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
          * `identity_pool_id` (`pulumi.Input[str]`) - ID of the Cognito Identity Pool to use
          * `role_arn` (`pulumi.Input[str]`) - ARN of the IAM role that has the AmazonESCognitoAccess policy attached
          * `user_pool_id` (`pulumi.Input[str]`) - ID of the Cognito User Pool to use

        The **domain_endpoint_options** object supports the following:

          * `enforceHttps` (`pulumi.Input[bool]`) - Whether or not to require HTTPS
          * `tlsSecurityPolicy` (`pulumi.Input[str]`)

        The **ebs_options** object supports the following:

          * `ebsEnabled` (`pulumi.Input[bool]`) - Whether EBS volumes are attached to data nodes in the domain.
          * `iops` (`pulumi.Input[float]`) - The baseline input/output (I/O) performance of EBS volumes
            attached to data nodes. Applicable only for the Provisioned IOPS EBS volume type.
          * `volume_size` (`pulumi.Input[float]`) - The size of EBS volumes attached to data nodes (in GB).
            **Required** if `ebs_enabled` is set to `true`.
          * `volumeType` (`pulumi.Input[str]`) - The type of EBS volumes attached to data nodes.

        The **encrypt_at_rest** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
          * `kms_key_id` (`pulumi.Input[str]`) - The KMS key id to encrypt the Elasticsearch domain with. If not specified then it defaults to using the `aws/es` service KMS key.

        The **log_publishing_options** object supports the following:

          * `cloudwatch_log_group_arn` (`pulumi.Input[str]`) - ARN of the Cloudwatch log group to which log needs to be published.
          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not
          * `logType` (`pulumi.Input[str]`) - A type of Elasticsearch log. Valid values: INDEX_SLOW_LOGS, SEARCH_SLOW_LOGS, ES_APPLICATION_LOGS

        The **node_to_node_encryption** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Specifies whether Amazon Cognito authentication with Kibana is enabled or not

        The **snapshot_options** object supports the following:

          * `automatedSnapshotStartHour` (`pulumi.Input[float]`) - Hour during which the service takes an automated daily
            snapshot of the indices in the domain.

        The **vpc_options** object supports the following:

          * `availability_zones` (`pulumi.Input[list]`)
          * `security_group_ids` (`pulumi.Input[list]`) - List of VPC Security Group IDs to be applied to the Elasticsearch domain endpoints. If omitted, the default Security Group for the VPC will be used.
          * `subnet_ids` (`pulumi.Input[list]`) - List of VPC Subnet IDs for the Elasticsearch domain endpoints to be created in.
          * `vpc_id` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_policies"] = access_policies
        __props__["advanced_options"] = advanced_options
        __props__["arn"] = arn
        __props__["cluster_config"] = cluster_config
        __props__["cognito_options"] = cognito_options
        __props__["domain_endpoint_options"] = domain_endpoint_options
        __props__["domain_id"] = domain_id
        __props__["domain_name"] = domain_name
        __props__["ebs_options"] = ebs_options
        __props__["elasticsearch_version"] = elasticsearch_version
        __props__["encrypt_at_rest"] = encrypt_at_rest
        __props__["endpoint"] = endpoint
        __props__["kibana_endpoint"] = kibana_endpoint
        __props__["log_publishing_options"] = log_publishing_options
        __props__["node_to_node_encryption"] = node_to_node_encryption
        __props__["snapshot_options"] = snapshot_options
        __props__["tags"] = tags
        __props__["vpc_options"] = vpc_options
        return Domain(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

