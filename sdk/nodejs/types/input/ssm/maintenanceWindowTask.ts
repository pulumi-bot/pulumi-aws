// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface MaintenanceWindowTaskLoggingInfo {
    s3BucketName: pulumi.Input<string>;
    s3BucketPrefix?: pulumi.Input<string>;
    s3Region: pulumi.Input<string>;
}

export interface MaintenanceWindowTaskTarget {
    key: pulumi.Input<string>;
    /**
     * The array of strings.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface MaintenanceWindowTaskTaskInvocationParameters {
    /**
     * The parameters for an AUTOMATION task type. Documented below.
     */
    automationParameters?: pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParameters>;
    /**
     * The parameters for a LAMBDA task type. Documented below.
     */
    lambdaParameters?: pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersLambdaParameters>;
    /**
     * The parameters for a RUN_COMMAND task type. Documented below.
     */
    runCommandParameters?: pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters>;
    /**
     * The parameters for a STEP_FUNCTIONS task type. Documented below.
     */
    stepFunctionsParameters?: pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersAutomationParameters {
    /**
     * The version of an Automation document to use during task execution.
     */
    documentVersion?: pulumi.Input<string>;
    /**
     * The parameters for the RUN_COMMAND task execution. Documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter>[]>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameter {
    /**
     * The parameter name.
     */
    name: pulumi.Input<string>;
    /**
     * The array of strings.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersLambdaParameters {
    /**
     * Pass client-specific information to the Lambda function that you are invoking.
     */
    clientContext?: pulumi.Input<string>;
    /**
     * JSON to provide to your Lambda function as input.
     */
    payload?: pulumi.Input<string>;
    /**
     * Specify a Lambda function version or alias name.
     */
    qualifier?: pulumi.Input<string>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters {
    /**
     * Information about the command(s) to execute.
     */
    comment?: pulumi.Input<string>;
    /**
     * The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
     */
    documentHash?: pulumi.Input<string>;
    /**
     * SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
     */
    documentHashType?: pulumi.Input<string>;
    /**
     * Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
     */
    notificationConfig?: pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig>;
    /**
     * The name of the Amazon S3 bucket.
     */
    outputS3Bucket?: pulumi.Input<string>;
    /**
     * The Amazon S3 bucket subfolder.
     */
    outputS3KeyPrefix?: pulumi.Input<string>;
    /**
     * The parameters for the RUN_COMMAND task execution. Documented below.
     */
    parameters?: pulumi.Input<pulumi.Input<inputApi.ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter>[]>;
    /**
     * The IAM service role to assume during task execution.
     */
    serviceRoleArn?: pulumi.Input<string>;
    /**
     * If this time is reached and the command has not already started executing, it doesn't run.
     */
    timeoutSeconds?: pulumi.Input<number>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig {
    /**
     * An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
     */
    notificationArn?: pulumi.Input<string>;
    /**
     * The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
     */
    notificationEvents?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`
     */
    notificationType?: pulumi.Input<string>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter {
    /**
     * The parameter name.
     */
    name: pulumi.Input<string>;
    /**
     * The array of strings.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters {
    /**
     * The inputs for the STEP_FUNCTION task.
     */
    input?: pulumi.Input<string>;
    /**
     * The parameter name.
     */
    name?: pulumi.Input<string>;
}

export interface MaintenanceWindowTaskTaskParameter {
    /**
     * The parameter name.
     */
    name: pulumi.Input<string>;
    /**
     * The array of strings.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}
