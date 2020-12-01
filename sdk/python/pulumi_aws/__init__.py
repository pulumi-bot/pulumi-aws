# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from ._enums import *
from .get_ami import *
from .get_ami_ids import *
from .get_arn import *
from .get_autoscaling_groups import *
from .get_availability_zone import *
from .get_availability_zones import *
from .get_billing_service_account import *
from .get_caller_identity import *
from .get_canonical_user_id import *
from .get_elastic_ip import *
from .get_ip_ranges import *
from .get_partition import *
from .get_prefix_list import *
from .get_region import *
from .get_regions import *
from .provider import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    accessanalyzer,
    acm,
    acmpca,
    alb,
    apigateway,
    apigatewayv2,
    appautoscaling,
    applicationloadbalancing,
    appmesh,
    appsync,
    athena,
    autoscaling,
    autoscalingplans,
    backup,
    batch,
    budgets,
    cfg,
    cloud9,
    cloudformation,
    cloudfront,
    cloudhsmv2,
    cloudtrail,
    cloudwatch,
    codeartifact,
    codebuild,
    codecommit,
    codedeploy,
    codepipeline,
    codestarnotifications,
    cognito,
    config,
    cur,
    datapipeline,
    datasync,
    dax,
    devicefarm,
    directconnect,
    directoryservice,
    dlm,
    dms,
    docdb,
    dynamodb,
    ebs,
    ec2,
    ec2clientvpn,
    ec2transitgateway,
    ecr,
    ecs,
    efs,
    eks,
    elasticache,
    elasticbeanstalk,
    elasticloadbalancing,
    elasticloadbalancingv2,
    elasticsearch,
    elastictranscoder,
    elb,
    emr,
    fms,
    fsx,
    gamelift,
    glacier,
    globalaccelerator,
    glue,
    guardduty,
    iam,
    imagebuilder,
    inspector,
    iot,
    kinesis,
    kinesisanalyticsv2,
    kms,
    lambda_,
    lb,
    lex,
    licensemanager,
    lightsail,
    macie,
    mediaconvert,
    mediapackage,
    mediastore,
    mq,
    msk,
    neptune,
    networkfirewall,
    opsworks,
    organizations,
    outposts,
    pinpoint,
    pricing,
    qldb,
    quicksight,
    ram,
    rds,
    redshift,
    resourcegroups,
    route53,
    s3,
    s3control,
    s3outposts,
    sagemaker,
    secretsmanager,
    securityhub,
    servicecatalog,
    servicediscovery,
    servicequotas,
    ses,
    sfn,
    shield,
    signer,
    simpledb,
    sns,
    sqs,
    ssm,
    storagegateway,
    swf,
    transfer,
    waf,
    wafregional,
    wafv2,
    worklink,
    workspaces,
    xray,
)

def _register_module():
    import pulumi

    class Package(pulumi.runtime.ResourcePackage):
        def version(self):
            return None

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:aws":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("aws", Package())

_register_module()
