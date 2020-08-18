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
    'AssociationOutputLocation',
    'AssociationTarget',
    'DocumentAttachmentsSource',
    'DocumentParameter',
    'MaintenanceWindowTargetTarget',
    'MaintenanceWindowTaskTarget',
    'MaintenanceWindowTaskTaskInvocationParameters',
    'MaintenanceWindowTaskTaskInvocationParametersAutomationParameters',
    'MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter',
    'MaintenanceWindowTaskTaskInvocationParametersLambdaParameters',
    'MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters',
    'MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig',
    'MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter',
    'MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters',
    'PatchBaselineApprovalRule',
    'PatchBaselineApprovalRulePatchFilter',
    'PatchBaselineGlobalFilter',
    'ResourceDataSyncS3Destination',
]

@pulumi.output_type
class AssociationOutputLocation(dict):
    @property
    @pulumi.getter(name="s3BucketName")
    def s3_bucket_name(self) -> str:
        """
        The S3 bucket name.
        """
        ...

    @property
    @pulumi.getter(name="s3KeyPrefix")
    def s3_key_prefix(self) -> Optional[str]:
        """
        The S3 bucket prefix. Results stored in the root if not configured.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AssociationTarget(dict):
    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Either `InstanceIds` or `tag:Tag Name` to specify an EC2 tag.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        A list of instance IDs or tag values. AWS currently limits this list size to one value.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DocumentAttachmentsSource(dict):
    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The key describing the location of an attachment to a document. Valid key types include: `SourceUrl` and `S3FileUrl`
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the document attachment file
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        The value describing the location of an attachment to a document
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DocumentParameter(dict):
    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the document.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the document.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTargetTarget(dict):
    @property
    @pulumi.getter
    def key(self) -> str:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTarget(dict):
    @property
    @pulumi.getter
    def key(self) -> str:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        The array of strings.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParameters(dict):
    @property
    @pulumi.getter(name="automationParameters")
    def automation_parameters(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParametersAutomationParameters']:
        """
        The parameters for an AUTOMATION task type. Documented below.
        """
        ...

    @property
    @pulumi.getter(name="lambdaParameters")
    def lambda_parameters(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParametersLambdaParameters']:
        """
        The parameters for a LAMBDA task type. Documented below.
        """
        ...

    @property
    @pulumi.getter(name="runCommandParameters")
    def run_command_parameters(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters']:
        """
        The parameters for a RUN_COMMAND task type. Documented below.
        """
        ...

    @property
    @pulumi.getter(name="stepFunctionsParameters")
    def step_functions_parameters(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters']:
        """
        The parameters for a STEP_FUNCTIONS task type. Documented below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersAutomationParameters(dict):
    @property
    @pulumi.getter(name="documentVersion")
    def document_version(self) -> Optional[str]:
        """
        The version of an Automation document to use during task execution.
        """
        ...

    @property
    @pulumi.getter
    def parameters(self) -> Optional[List['outputs.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter']]:
        """
        The parameters for the RUN_COMMAND task execution. Documented below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The parameter name.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        The array of strings.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersLambdaParameters(dict):
    @property
    @pulumi.getter(name="clientContext")
    def client_context(self) -> Optional[str]:
        """
        Pass client-specific information to the Lambda function that you are invoking.
        """
        ...

    @property
    @pulumi.getter
    def payload(self) -> Optional[str]:
        """
        JSON to provide to your Lambda function as input.
        """
        ...

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[str]:
        """
        Specify a Lambda function version or alias name.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters(dict):
    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        """
        Information about the command(s) to execute.
        """
        ...

    @property
    @pulumi.getter(name="documentHash")
    def document_hash(self) -> Optional[str]:
        """
        The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
        """
        ...

    @property
    @pulumi.getter(name="documentHashType")
    def document_hash_type(self) -> Optional[str]:
        """
        SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
        """
        ...

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> Optional['outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig']:
        """
        Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
        """
        ...

    @property
    @pulumi.getter(name="outputS3Bucket")
    def output_s3_bucket(self) -> Optional[str]:
        """
        The name of the Amazon S3 bucket.
        """
        ...

    @property
    @pulumi.getter(name="outputS3KeyPrefix")
    def output_s3_key_prefix(self) -> Optional[str]:
        """
        The Amazon S3 bucket subfolder.
        """
        ...

    @property
    @pulumi.getter
    def parameters(self) -> Optional[List['outputs.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter']]:
        """
        The parameters for the RUN_COMMAND task execution. Documented below.
        """
        ...

    @property
    @pulumi.getter(name="serviceRoleArn")
    def service_role_arn(self) -> Optional[str]:
        """
        The IAM service role to assume during task execution.
        """
        ...

    @property
    @pulumi.getter(name="timeoutSeconds")
    def timeout_seconds(self) -> Optional[float]:
        """
        If this time is reached and the command has not already started executing, it doesn't run.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig(dict):
    @property
    @pulumi.getter(name="notificationArn")
    def notification_arn(self) -> Optional[str]:
        """
        An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
        """
        ...

    @property
    @pulumi.getter(name="notificationEvents")
    def notification_events(self) -> Optional[List[str]]:
        """
        The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
        """
        ...

    @property
    @pulumi.getter(name="notificationType")
    def notification_type(self) -> Optional[str]:
        """
        When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The parameter name.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        """
        The array of strings.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters(dict):
    @property
    @pulumi.getter
    def input(self) -> Optional[str]:
        """
        The inputs for the STEP_FUNCTION task.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the STEP_FUNCTION task.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PatchBaselineApprovalRule(dict):
    @property
    @pulumi.getter(name="approveAfterDays")
    def approve_after_days(self) -> float:
        """
        The number of days after the release date of each patch matched by the rule the patch is marked as approved in the patch baseline. Valid Range: 0 to 100.
        """
        ...

    @property
    @pulumi.getter(name="complianceLevel")
    def compliance_level(self) -> Optional[str]:
        """
        Defines the compliance level for patches approved by this rule. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
        """
        ...

    @property
    @pulumi.getter(name="enableNonSecurity")
    def enable_non_security(self) -> Optional[bool]:
        """
        Boolean enabling the application of non-security updates. The default value is 'false'. Valid for Linux instances only.
        """
        ...

    @property
    @pulumi.getter(name="patchFilters")
    def patch_filters(self) -> List['outputs.PatchBaselineApprovalRulePatchFilter']:
        """
        The patch filter group that defines the criteria for the rule. Up to 5 patch filters can be specified per approval rule using Key/Value pairs. Valid Keys are `PATCH_SET | PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PatchBaselineApprovalRulePatchFilter(dict):
    @property
    @pulumi.getter
    def key(self) -> str:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PatchBaselineGlobalFilter(dict):
    @property
    @pulumi.getter
    def key(self) -> str:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceDataSyncS3Destination(dict):
    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        """
        Name of S3 bucket where the aggregated data is stored.
        """
        ...

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        """
        ARN of an encryption key for a destination in Amazon S3.
        """
        ...

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        """
        Prefix for the bucket.
        """
        ...

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region with the bucket targeted by the Resource Data Sync.
        """
        ...

    @property
    @pulumi.getter(name="syncFormat")
    def sync_format(self) -> Optional[str]:
        """
        A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


