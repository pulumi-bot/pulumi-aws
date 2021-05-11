// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Elastic Transcoder preset resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bar = new aws.elastictranscoder.Preset("bar", {
 *     audio: {
 *         audioPackingMode: "SingleTrack",
 *         bitRate: "96",
 *         channels: "2",
 *         codec: "AAC",
 *         sampleRate: "44100",
 *     },
 *     audioCodecOptions: {
 *         profile: "AAC-LC",
 *     },
 *     container: "mp4",
 *     description: "Sample Preset",
 *     thumbnails: {
 *         format: "png",
 *         interval: "120",
 *         maxHeight: "auto",
 *         maxWidth: "auto",
 *         paddingPolicy: "Pad",
 *         sizingPolicy: "Fit",
 *     },
 *     video: {
 *         bitRate: "1600",
 *         codec: "H.264",
 *         displayAspectRatio: "16:9",
 *         fixedGop: "false",
 *         frameRate: "auto",
 *         keyframesMaxDist: "240",
 *         maxFrameRate: "60",
 *         maxHeight: "auto",
 *         maxWidth: "auto",
 *         paddingPolicy: "Pad",
 *         sizingPolicy: "Fit",
 *     },
 *     videoCodecOptions: {
 *         ColorSpaceConversionMode: "None",
 *         InterlacedMode: "Progressive",
 *         Level: "2.2",
 *         MaxReferenceFrames: 3,
 *         Profile: "main",
 *     },
 *     videoWatermarks: [{
 *         horizontalAlign: "Right",
 *         horizontalOffset: "10px",
 *         id: "Test",
 *         maxHeight: "20%",
 *         maxWidth: "20%",
 *         opacity: "55.5",
 *         sizingPolicy: "ShrinkToFit",
 *         target: "Content",
 *         verticalAlign: "Bottom",
 *         verticalOffset: "10px",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Elastic Transcoder presets can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:elastictranscoder/preset:Preset basic_preset 1407981661351-cttk8b
 * ```
 */
export class Preset extends pulumi.CustomResource {
    /**
     * Get an existing Preset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PresetState, opts?: pulumi.CustomResourceOptions): Preset {
        return new Preset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elastictranscoder/preset:Preset';

    /**
     * Returns true if the given object is an instance of Preset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Preset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Preset.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the Elastic Transcoder Preset.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Audio parameters object (documented below).
     */
    public readonly audio!: pulumi.Output<outputs.elastictranscoder.PresetAudio | undefined>;
    /**
     * Codec options for the audio parameters (documented below)
     */
    public readonly audioCodecOptions!: pulumi.Output<outputs.elastictranscoder.PresetAudioCodecOptions | undefined>;
    /**
     * The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
     */
    public readonly container!: pulumi.Output<string>;
    /**
     * A description of the preset (maximum 255 characters)
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the preset. (maximum 40 characters)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Thumbnail parameters object (documented below)
     */
    public readonly thumbnails!: pulumi.Output<outputs.elastictranscoder.PresetThumbnails | undefined>;
    public readonly type!: pulumi.Output<string>;
    /**
     * Video parameters object (documented below)
     */
    public readonly video!: pulumi.Output<outputs.elastictranscoder.PresetVideo | undefined>;
    /**
     * Codec options for the video parameters
     */
    public readonly videoCodecOptions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Watermark parameters for the video parameters (documented below)
     */
    public readonly videoWatermarks!: pulumi.Output<outputs.elastictranscoder.PresetVideoWatermark[] | undefined>;

    /**
     * Create a Preset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PresetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PresetArgs | PresetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PresetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["audio"] = state ? state.audio : undefined;
            inputs["audioCodecOptions"] = state ? state.audioCodecOptions : undefined;
            inputs["container"] = state ? state.container : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["thumbnails"] = state ? state.thumbnails : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["video"] = state ? state.video : undefined;
            inputs["videoCodecOptions"] = state ? state.videoCodecOptions : undefined;
            inputs["videoWatermarks"] = state ? state.videoWatermarks : undefined;
        } else {
            const args = argsOrState as PresetArgs | undefined;
            if ((!args || args.container === undefined) && !opts.urn) {
                throw new Error("Missing required property 'container'");
            }
            inputs["audio"] = args ? args.audio : undefined;
            inputs["audioCodecOptions"] = args ? args.audioCodecOptions : undefined;
            inputs["container"] = args ? args.container : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["thumbnails"] = args ? args.thumbnails : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["video"] = args ? args.video : undefined;
            inputs["videoCodecOptions"] = args ? args.videoCodecOptions : undefined;
            inputs["videoWatermarks"] = args ? args.videoWatermarks : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Preset.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Preset resources.
 */
export interface PresetState {
    /**
     * Amazon Resource Name (ARN) of the Elastic Transcoder Preset.
     */
    arn?: pulumi.Input<string>;
    /**
     * Audio parameters object (documented below).
     */
    audio?: pulumi.Input<inputs.elastictranscoder.PresetAudio>;
    /**
     * Codec options for the audio parameters (documented below)
     */
    audioCodecOptions?: pulumi.Input<inputs.elastictranscoder.PresetAudioCodecOptions>;
    /**
     * The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
     */
    container?: pulumi.Input<string>;
    /**
     * A description of the preset (maximum 255 characters)
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the preset. (maximum 40 characters)
     */
    name?: pulumi.Input<string>;
    /**
     * Thumbnail parameters object (documented below)
     */
    thumbnails?: pulumi.Input<inputs.elastictranscoder.PresetThumbnails>;
    type?: pulumi.Input<string>;
    /**
     * Video parameters object (documented below)
     */
    video?: pulumi.Input<inputs.elastictranscoder.PresetVideo>;
    /**
     * Codec options for the video parameters
     */
    videoCodecOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Watermark parameters for the video parameters (documented below)
     */
    videoWatermarks?: pulumi.Input<pulumi.Input<inputs.elastictranscoder.PresetVideoWatermark>[]>;
}

/**
 * The set of arguments for constructing a Preset resource.
 */
export interface PresetArgs {
    /**
     * Audio parameters object (documented below).
     */
    audio?: pulumi.Input<inputs.elastictranscoder.PresetAudio>;
    /**
     * Codec options for the audio parameters (documented below)
     */
    audioCodecOptions?: pulumi.Input<inputs.elastictranscoder.PresetAudioCodecOptions>;
    /**
     * The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
     */
    container: pulumi.Input<string>;
    /**
     * A description of the preset (maximum 255 characters)
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the preset. (maximum 40 characters)
     */
    name?: pulumi.Input<string>;
    /**
     * Thumbnail parameters object (documented below)
     */
    thumbnails?: pulumi.Input<inputs.elastictranscoder.PresetThumbnails>;
    type?: pulumi.Input<string>;
    /**
     * Video parameters object (documented below)
     */
    video?: pulumi.Input<inputs.elastictranscoder.PresetVideo>;
    /**
     * Codec options for the video parameters
     */
    videoCodecOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Watermark parameters for the video parameters (documented below)
     */
    videoWatermarks?: pulumi.Input<pulumi.Input<inputs.elastictranscoder.PresetVideoWatermark>[]>;
}
