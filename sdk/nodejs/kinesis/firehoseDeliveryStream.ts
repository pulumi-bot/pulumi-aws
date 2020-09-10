// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class FirehoseDeliveryStream extends pulumi.CustomResource {
    /**
     * Get an existing FirehoseDeliveryStream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirehoseDeliveryStreamState, opts?: pulumi.CustomResourceOptions): FirehoseDeliveryStream {
        return new FirehoseDeliveryStream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream';

    /**
     * Returns true if the given object is an instance of FirehoseDeliveryStream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirehoseDeliveryStream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirehoseDeliveryStream.__pulumiType;
    }

    public readonly arn!: pulumi.Output<string>;
    public readonly destination!: pulumi.Output<string>;
    public readonly destinationId!: pulumi.Output<string>;
    public readonly elasticsearchConfiguration!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamElasticsearchConfiguration | undefined>;
    public readonly extendedS3Configuration!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamExtendedS3Configuration | undefined>;
    public readonly kinesisSourceConfiguration!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamKinesisSourceConfiguration | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly redshiftConfiguration!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamRedshiftConfiguration | undefined>;
    public readonly s3Configuration!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamS3Configuration | undefined>;
    public readonly serverSideEncryption!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamServerSideEncryption | undefined>;
    public readonly splunkConfiguration!: pulumi.Output<outputs.kinesis.FirehoseDeliveryStreamSplunkConfiguration | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly versionId!: pulumi.Output<string>;

    /**
     * Create a FirehoseDeliveryStream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirehoseDeliveryStreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirehoseDeliveryStreamArgs | FirehoseDeliveryStreamState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirehoseDeliveryStreamState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["destinationId"] = state ? state.destinationId : undefined;
            inputs["elasticsearchConfiguration"] = state ? state.elasticsearchConfiguration : undefined;
            inputs["extendedS3Configuration"] = state ? state.extendedS3Configuration : undefined;
            inputs["kinesisSourceConfiguration"] = state ? state.kinesisSourceConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["redshiftConfiguration"] = state ? state.redshiftConfiguration : undefined;
            inputs["s3Configuration"] = state ? state.s3Configuration : undefined;
            inputs["serverSideEncryption"] = state ? state.serverSideEncryption : undefined;
            inputs["splunkConfiguration"] = state ? state.splunkConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
        } else {
            const args = argsOrState as FirehoseDeliveryStreamArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["destinationId"] = args ? args.destinationId : undefined;
            inputs["elasticsearchConfiguration"] = args ? args.elasticsearchConfiguration : undefined;
            inputs["extendedS3Configuration"] = args ? args.extendedS3Configuration : undefined;
            inputs["kinesisSourceConfiguration"] = args ? args.kinesisSourceConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["redshiftConfiguration"] = args ? args.redshiftConfiguration : undefined;
            inputs["s3Configuration"] = args ? args.s3Configuration : undefined;
            inputs["serverSideEncryption"] = args ? args.serverSideEncryption : undefined;
            inputs["splunkConfiguration"] = args ? args.splunkConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FirehoseDeliveryStream.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirehoseDeliveryStream resources.
 */
export interface FirehoseDeliveryStreamState {
    readonly arn?: pulumi.Input<string>;
    readonly destination?: pulumi.Input<string>;
    readonly destinationId?: pulumi.Input<string>;
    readonly elasticsearchConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamElasticsearchConfiguration>;
    readonly extendedS3Configuration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamExtendedS3Configuration>;
    readonly kinesisSourceConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamKinesisSourceConfiguration>;
    readonly name?: pulumi.Input<string>;
    readonly redshiftConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamRedshiftConfiguration>;
    readonly s3Configuration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamS3Configuration>;
    readonly serverSideEncryption?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamServerSideEncryption>;
    readonly splunkConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamSplunkConfiguration>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly versionId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirehoseDeliveryStream resource.
 */
export interface FirehoseDeliveryStreamArgs {
    readonly arn?: pulumi.Input<string>;
    readonly destination: pulumi.Input<string>;
    readonly destinationId?: pulumi.Input<string>;
    readonly elasticsearchConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamElasticsearchConfiguration>;
    readonly extendedS3Configuration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamExtendedS3Configuration>;
    readonly kinesisSourceConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamKinesisSourceConfiguration>;
    readonly name?: pulumi.Input<string>;
    readonly redshiftConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamRedshiftConfiguration>;
    readonly s3Configuration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamS3Configuration>;
    readonly serverSideEncryption?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamServerSideEncryption>;
    readonly splunkConfiguration?: pulumi.Input<inputs.kinesis.FirehoseDeliveryStreamSplunkConfiguration>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly versionId?: pulumi.Input<string>;
}
