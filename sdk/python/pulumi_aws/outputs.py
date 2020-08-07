# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'GetAmiBlockDeviceMapping',
    'GetAmiFilter',
    'GetAmiIdsFilter',
    'GetAmiProductCode',
    'GetAutoscalingGroupsFilter',
    'GetAvailabilityZoneFilter',
    'GetAvailabilityZonesFilter',
    'GetElasticIpFilter',
    'GetPrefixListFilter',
    'GetRegionsFilter',
    'ProviderAssumeRole',
    'ProviderEndpoint',
    'ProviderIgnoreTags',
]

@pulumi.output_type
class GetAmiBlockDeviceMapping(dict):
    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> str:
        ...

    @property
    @pulumi.getter
    def ebs(self) -> Mapping[str, str]:
        ...

    @property
    @pulumi.getter(name="noDevice")
    def no_device(self) -> str:
        ...

    @property
    @pulumi.getter(name="virtualName")
    def virtual_name(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetAmiFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the AMI that was provided during image creation.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetAmiIdsFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetAmiProductCode(dict):
    @property
    @pulumi.getter(name="productCodeId")
    def product_code_id(self) -> str:
        ...

    @property
    @pulumi.getter(name="productCodeType")
    def product_code_type(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetAutoscalingGroupsFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter. The valid values are: `auto-scaling-group`, `key`, `value`, and `propagate-at-launch`.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        The value of the filter.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetAvailabilityZoneFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter field. Valid values can be found in the [EC2 DescribeAvailabilityZones API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html).
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetAvailabilityZonesFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter field. Valid values can be found in the [EC2 DescribeAvailabilityZones API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html).
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetElasticIpFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPrefixListFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetRegionsFilter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter field. Valid values can be found in the [describe-regions AWS CLI Reference][1].
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderAssumeRole(dict):
    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def policy(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderEndpoint(dict):
    @property
    @pulumi.getter
    def accessanalyzer(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def acm(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def acmpca(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def amplify(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def apigateway(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def applicationautoscaling(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def applicationinsights(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def appmesh(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def appstream(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def appsync(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def athena(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def autoscaling(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def autoscalingplans(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def backup(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def batch(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def budgets(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloud9(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudformation(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudfront(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudhsm(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudsearch(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudtrail(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudwatch(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudwatchevents(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cloudwatchlogs(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def codeartifact(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def codebuild(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def codecommit(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def codedeploy(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def codepipeline(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cognitoidentity(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cognitoidp(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def configservice(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def cur(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def dataexchange(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def datapipeline(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def datasync(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def dax(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def devicefarm(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def directconnect(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def dlm(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def dms(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def docdb(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ds(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def dynamodb(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ec2(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ecr(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ecs(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def efs(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def eks(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def elasticache(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def elasticbeanstalk(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def elastictranscoder(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def elb(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def emr(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def es(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def firehose(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def fms(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def forecast(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def fsx(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def gamelift(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def glacier(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def globalaccelerator(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def glue(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def greengrass(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def guardduty(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def iam(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def imagebuilder(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def inspector(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def iot(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def iotanalytics(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def iotevents(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def kafka(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def kinesis(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="kinesisAnalytics")
    def kinesis_analytics(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def kinesisanalytics(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def kinesisanalyticsv2(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def kinesisvideo(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def kms(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def lakeformation(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="lambda")
    def lambda_(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def lexmodels(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def licensemanager(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def lightsail(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def macie(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def managedblockchain(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def marketplacecatalog(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def mediaconnect(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def mediaconvert(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def medialive(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def mediapackage(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def mediastore(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def mediastoredata(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def mq(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def neptune(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def networkmanager(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def opsworks(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def organizations(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def outposts(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def personalize(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def pinpoint(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def pricing(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def qldb(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def quicksight(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def r53(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ram(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def rds(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def redshift(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def resourcegroups(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def resourcegroupstaggingapi(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def route53(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def route53domains(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def route53resolver(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def s3(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def s3control(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def sagemaker(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def sdb(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def secretsmanager(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def securityhub(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def serverlessrepo(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def servicecatalog(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def servicediscovery(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def servicequotas(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ses(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def shield(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def sns(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def sqs(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def ssm(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def stepfunctions(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def storagegateway(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def sts(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def swf(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def synthetics(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def transfer(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def waf(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def wafregional(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def wafv2(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def worklink(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def workmail(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def workspaces(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def xray(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderIgnoreTags(dict):
    @property
    @pulumi.getter(name="keyPrefixes")
    def key_prefixes(self) -> Optional[List[str]]:
        ...

    @property
    @pulumi.getter
    def keys(self) -> Optional[List[str]]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


