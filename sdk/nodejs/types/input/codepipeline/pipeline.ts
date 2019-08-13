// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface PipelineArtifactStore {
    encryptionKey?: pulumi.Input<inputApi.codepipeline.PipelineArtifactStoreEncryptionKey>;
    location: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface PipelineArtifactStoreEncryptionKey {
    /**
     * The codepipeline ID.
     */
    id: pulumi.Input<string>;
    type: pulumi.Input<string>;
}

export interface PipelineStage {
    actions: pulumi.Input<pulumi.Input<inputApi.codepipeline.PipelineStageAction>[]>;
    /**
     * The name of the pipeline.
     */
    name: pulumi.Input<string>;
}

export interface PipelineStageAction {
    category: pulumi.Input<string>;
    configuration?: pulumi.Input<{[key: string]: any}>;
    inputArtifacts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the pipeline.
     */
    name: pulumi.Input<string>;
    outputArtifacts?: pulumi.Input<pulumi.Input<string>[]>;
    owner: pulumi.Input<string>;
    provider: pulumi.Input<string>;
    /**
     * A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
     */
    roleArn?: pulumi.Input<string>;
    runOrder?: pulumi.Input<number>;
    version: pulumi.Input<string>;
}
