// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface EventTargetBatchTarget {
    /**
     * The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
     */
    arraySize?: pulumi.Input<number>;
    /**
     * The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
     */
    jobAttempts?: pulumi.Input<number>;
    /**
     * The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
     */
    jobDefinition: pulumi.Input<string>;
    /**
     * The name to use for this execution of the job, if the target is an AWS Batch job.
     */
    jobName: pulumi.Input<string>;
}

export interface EventTargetEcsTarget {
    /**
     * Specifies an ECS task group for the task. The maximum length is 255 characters.
     */
    group?: pulumi.Input<string>;
    /**
     * Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values are EC2 or FARGATE.
     */
    launchType?: pulumi.Input<string>;
    /**
     * Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if launchType is FARGATE because the awsvpc mode is required for Fargate tasks.
     */
    networkConfiguration?: pulumi.Input<inputApi.cloudwatch.EventTargetEcsTargetNetworkConfiguration>;
    /**
     * Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see [AWS Fargate Platform Versions](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
     */
    platformVersion?: pulumi.Input<string>;
    /**
     * The number of tasks to create based on the TaskDefinition. The default is 1.
     */
    taskCount?: pulumi.Input<number>;
    /**
     * The ARN of the task definition to use if the event target is an Amazon ECS cluster.
     */
    taskDefinitionArn: pulumi.Input<string>;
}

export interface EventTargetEcsTargetNetworkConfiguration {
    /**
     * Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.
     */
    assignPublicIp?: pulumi.Input<boolean>;
    /**
     * The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The subnets associated with the task or service.
     */
    subnets: pulumi.Input<pulumi.Input<string>[]>;
}

export interface EventTargetInputTransformer {
    /**
     * Key value pairs specified in the form of JSONPath (for example, time = $.time)
     */
    inputPaths?: pulumi.Input<{[key: string]: any}>;
    /**
     * Structure containing the template body.
     */
    inputTemplate: pulumi.Input<string>;
}

export interface EventTargetKinesisTarget {
    /**
     * The JSON path to be extracted from the event and used as the partition key.
     */
    partitionKeyPath?: pulumi.Input<string>;
}

export interface EventTargetRunCommandTarget {
    /**
     * Can be either `tag:tag-key` or `InstanceIds`.
     */
    key: pulumi.Input<string>;
    /**
     * If Key is `tag:tag-key`, Values is a list of tag values. If Key is `InstanceIds`, Values is a list of Amazon EC2 instance IDs.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface EventTargetSqsTarget {
    /**
     * The FIFO message group ID to use as the target.
     */
    messageGroupId?: pulumi.Input<string>;
}
