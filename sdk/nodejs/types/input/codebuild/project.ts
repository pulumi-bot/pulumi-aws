// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ProjectArtifacts {
    /**
     * If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
     */
    encryptionDisabled?: pulumi.Input<boolean>;
    /**
     * The location of the source code from git or s3.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to use in storing build artifacts. If `type` is set to `S3`, then valid values for this parameter are: `BUILD_ID` or `NONE`.
     */
    namespaceType?: pulumi.Input<string>;
    /**
     * If set to true, a name specified in the build spec file overrides the artifact name.
     */
    overrideArtifactName?: pulumi.Input<boolean>;
    /**
     * The type of build output artifact to create. If `type` is set to `S3`, valid values for this parameter are: `NONE` or `ZIP`
     */
    packaging?: pulumi.Input<string>;
    /**
     * If `type` is set to `S3`, this is the path to the output artifact
     */
    path?: pulumi.Input<string>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectCache {
    /**
     * The location of the source code from git or s3.
     */
    location?: pulumi.Input<string>;
    /**
     * Specifies settings that AWS CodeBuild uses to store and reuse build dependencies. Valid values:  `LOCAL_SOURCE_CACHE`, `LOCAL_DOCKER_LAYER_CACHE`, and `LOCAL_CUSTOM_CACHE`
     */
    modes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type?: pulumi.Input<string>;
}

export interface ProjectEnvironment {
    /**
     * The ARN of the S3 bucket, path prefix and object key that contains the PEM-encoded certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Information about the compute resources the build project will use. Available values for this parameter are: `BUILD_GENERAL1_SMALL`, `BUILD_GENERAL1_MEDIUM` or `BUILD_GENERAL1_LARGE`. `BUILD_GENERAL1_SMALL` is only valid if `type` is set to `LINUX_CONTAINER`
     */
    computeType: pulumi.Input<string>;
    /**
     * A set of environment variables to make available to builds for this build project.
     */
    environmentVariables?: pulumi.Input<pulumi.Input<inputApi.codebuild.ProjectEnvironmentEnvironmentVariable>[]>;
    /**
     * The Docker image to use for this build project. Valid values include [Docker images provided by CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html) (e.g `aws/codebuild/standard:2.0`), [Docker Hub images](https://hub.docker.com/) (e.g. `nginx:latest`), and full Docker repository URIs such as those for ECR (e.g. `137112412989.dkr.ecr.us-west-2.amazonaws.com/amazonlinux:latest`).
     */
    image: pulumi.Input<string>;
    /**
     * The type of credentials AWS CodeBuild uses to pull images in your build. Available values for this parameter are `CODEBUILD` or `SERVICE_ROLE`. When you use a cross-account or private registry image, you must use SERVICE_ROLE credentials. When you use an AWS CodeBuild curated image, you must use CODEBUILD credentials. Default to `CODEBUILD`
     */
    imagePullCredentialsType?: pulumi.Input<string>;
    /**
     * If set to true, enables running the Docker daemon inside a Docker container. Defaults to `false`.
     */
    privilegedMode?: pulumi.Input<boolean>;
    /**
     * Information about credentials for access to a private Docker registry. Registry Credential config blocks are documented below.
     */
    registryCredential?: pulumi.Input<inputApi.codebuild.ProjectEnvironmentRegistryCredential>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectEnvironmentEnvironmentVariable {
    /**
     * The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
     */
    name: pulumi.Input<string>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type?: pulumi.Input<string>;
    /**
     * The environment variable's value.
     */
    value: pulumi.Input<string>;
}

export interface ProjectEnvironmentRegistryCredential {
    /**
     * The Amazon Resource Name (ARN) or name of credentials created using AWS Secrets Manager.
     */
    credential: pulumi.Input<string>;
    /**
     * The service that created the credentials to access a private Docker registry. The valid value, SECRETS_MANAGER, is for AWS Secrets Manager.
     */
    credentialProvider: pulumi.Input<string>;
}

export interface ProjectLogsConfig {
    /**
     * Configuration for the builds to store logs to CloudWatch
     */
    cloudwatchLogs?: pulumi.Input<inputApi.codebuild.ProjectLogsConfigCloudwatchLogs>;
    /**
     * Configuration for the builds to store logs to S3.
     */
    s3Logs?: pulumi.Input<inputApi.codebuild.ProjectLogsConfigS3Logs>;
}

export interface ProjectLogsConfigCloudwatchLogs {
    /**
     * The group name of the logs in CloudWatch Logs.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Current status of logs in S3 for a build project. Valid values: `ENABLED`, `DISABLED`. Defaults to `DISABLED`.
     */
    status?: pulumi.Input<string>;
    /**
     * The stream name of the logs in CloudWatch Logs.
     */
    streamName?: pulumi.Input<string>;
}

export interface ProjectLogsConfigS3Logs {
    /**
     * If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
     */
    encryptionDisabled?: pulumi.Input<boolean>;
    /**
     * The location of the source code from git or s3.
     */
    location?: pulumi.Input<string>;
    /**
     * Current status of logs in S3 for a build project. Valid values: `ENABLED`, `DISABLED`. Defaults to `DISABLED`.
     */
    status?: pulumi.Input<string>;
}

export interface ProjectSecondaryArtifact {
    /**
     * The artifact identifier. Must be the same specified inside AWS CodeBuild buildspec.
     */
    artifactIdentifier: pulumi.Input<string>;
    /**
     * If set to true, output artifacts will not be encrypted. If `type` is set to `NO_ARTIFACTS` then this value will be ignored. Defaults to `false`.
     */
    encryptionDisabled?: pulumi.Input<boolean>;
    /**
     * The location of the source code from git or s3.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to use in storing build artifacts. If `type` is set to `S3`, then valid values for this parameter are: `BUILD_ID` or `NONE`.
     */
    namespaceType?: pulumi.Input<string>;
    /**
     * If set to true, a name specified in the build spec file overrides the artifact name.
     */
    overrideArtifactName?: pulumi.Input<boolean>;
    /**
     * The type of build output artifact to create. If `type` is set to `S3`, valid values for this parameter are: `NONE` or `ZIP`
     */
    packaging?: pulumi.Input<string>;
    /**
     * If `type` is set to `S3`, this is the path to the output artifact
     */
    path?: pulumi.Input<string>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectSecondarySource {
    /**
     * Information about the authorization settings for AWS CodeBuild to access the source code to be built. Auth blocks are documented below.
     */
    auths?: pulumi.Input<pulumi.Input<inputApi.codebuild.ProjectSecondarySourceAuth>[]>;
    /**
     * The build spec declaration to use for this build project's related builds.
     */
    buildspec?: pulumi.Input<string>;
    /**
     * Truncate git history to this many commits.
     */
    gitCloneDepth?: pulumi.Input<number>;
    /**
     * Ignore SSL warnings when connecting to source control.
     */
    insecureSsl?: pulumi.Input<boolean>;
    /**
     * The location of the source code from git or s3.
     */
    location?: pulumi.Input<string>;
    /**
     * Set to `true` to report the status of a build's start and finish to your source provider. This option is only valid when your source provider is `GITHUB`, `BITBUCKET`, or `GITHUB_ENTERPRISE`.
     */
    reportBuildStatus?: pulumi.Input<boolean>;
    /**
     * The source identifier. Source data will be put inside a folder named as this parameter inside AWS CodeBuild source directory
     */
    sourceIdentifier: pulumi.Input<string>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectSecondarySourceAuth {
    /**
     * The resource value that applies to the specified authorization type.
     */
    resource?: pulumi.Input<string>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectSource {
    /**
     * Information about the authorization settings for AWS CodeBuild to access the source code to be built. Auth blocks are documented below.
     */
    auths?: pulumi.Input<pulumi.Input<inputApi.codebuild.ProjectSourceAuth>[]>;
    /**
     * The build spec declaration to use for this build project's related builds.
     */
    buildspec?: pulumi.Input<string>;
    /**
     * Truncate git history to this many commits.
     */
    gitCloneDepth?: pulumi.Input<number>;
    /**
     * Ignore SSL warnings when connecting to source control.
     */
    insecureSsl?: pulumi.Input<boolean>;
    /**
     * The location of the source code from git or s3.
     */
    location?: pulumi.Input<string>;
    /**
     * Set to `true` to report the status of a build's start and finish to your source provider. This option is only valid when your source provider is `GITHUB`, `BITBUCKET`, or `GITHUB_ENTERPRISE`.
     */
    reportBuildStatus?: pulumi.Input<boolean>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectSourceAuth {
    /**
     * The resource value that applies to the specified authorization type.
     */
    resource?: pulumi.Input<string>;
    /**
     * The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
     */
    type: pulumi.Input<string>;
}

export interface ProjectVpcConfig {
    /**
     * The security group IDs to assign to running builds.
     */
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The subnet IDs within which to run builds.
     */
    subnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the VPC within which to run builds.
     */
    vpcId: pulumi.Input<string>;
}
