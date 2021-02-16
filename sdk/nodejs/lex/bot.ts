// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Amazon Lex Bot resource. For more information see
 * [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const orderFlowersBot = new aws.lex.Bot("order_flowers_bot", {
 *     abortStatement: {
 *         messages: [{
 *             content: "Sorry, I am not able to assist at this time",
 *             contentType: "PlainText",
 *         }],
 *     },
 *     childDirected: false,
 *     clarificationPrompt: {
 *         maxAttempts: 2,
 *         messages: [{
 *             content: "I didn't understand you, what would you like to do?",
 *             contentType: "PlainText",
 *         }],
 *     },
 *     createVersion: false,
 *     description: "Bot to order flowers on the behalf of a user",
 *     idleSessionTtlInSeconds: 600,
 *     intents: [{
 *         intentName: "OrderFlowers",
 *         intentVersion: "1",
 *     }],
 *     locale: "en-US",
 *     name: "OrderFlowers",
 *     processBehavior: "BUILD",
 *     voiceId: "Salli",
 * });
 * ```
 *
 * ## Import
 *
 * Bots can be imported using their name.
 *
 * ```sh
 *  $ pulumi import aws:lex/bot:Bot order_flowers_bot OrderFlowers
 * ```
 */
export class Bot extends pulumi.CustomResource {
    /**
     * Get an existing Bot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BotState, opts?: pulumi.CustomResourceOptions): Bot {
        return new Bot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lex/bot:Bot';

    /**
     * Returns true if the given object is an instance of Bot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bot.__pulumiType;
    }

    /**
     * The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
     */
    public readonly abortStatement!: pulumi.Output<outputs.lex.BotAbortStatement>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Checksum identifying the version of the bot that was created. The checksum is not
     * included as an argument because the resource will add it automatically when updating the bot.
     */
    public /*out*/ readonly checksum!: pulumi.Output<string>;
    /**
     * By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
     */
    public readonly childDirected!: pulumi.Output<boolean>;
    /**
     * The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
     */
    public readonly clarificationPrompt!: pulumi.Output<outputs.lex.BotClarificationPrompt | undefined>;
    /**
     * Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
     */
    public readonly createVersion!: pulumi.Output<boolean | undefined>;
    /**
     * The date when the bot version was created.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * A description of the bot. Must be less than or equal to 200 characters in length.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
     */
    public readonly detectSentiment!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
     */
    public readonly enableModelImprovements!: pulumi.Output<boolean | undefined>;
    /**
     * If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
     */
    public /*out*/ readonly failureReason!: pulumi.Output<string>;
    /**
     * The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
     */
    public readonly idleSessionTtlInSeconds!: pulumi.Output<number | undefined>;
    /**
     * A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 100 Intent objects.
     */
    public readonly intents!: pulumi.Output<outputs.lex.BotIntent[]>;
    /**
     * The date when the $LATEST version of this bot was updated.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
     */
    public readonly locale!: pulumi.Output<string | undefined>;
    /**
     * The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
     */
    public readonly nluIntentConfidenceThreshold!: pulumi.Output<number | undefined>;
    /**
     * If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
     */
    public readonly processBehavior!: pulumi.Output<string | undefined>;
    /**
     * When you send a request to create or update a bot, Amazon Lex sets the status response
     * element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
     * build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
     * failureReason response element.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The version of the bot.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
     */
    public readonly voiceId!: pulumi.Output<string>;

    /**
     * Create a Bot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BotArgs | BotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BotState | undefined;
            inputs["abortStatement"] = state ? state.abortStatement : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["checksum"] = state ? state.checksum : undefined;
            inputs["childDirected"] = state ? state.childDirected : undefined;
            inputs["clarificationPrompt"] = state ? state.clarificationPrompt : undefined;
            inputs["createVersion"] = state ? state.createVersion : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["detectSentiment"] = state ? state.detectSentiment : undefined;
            inputs["enableModelImprovements"] = state ? state.enableModelImprovements : undefined;
            inputs["failureReason"] = state ? state.failureReason : undefined;
            inputs["idleSessionTtlInSeconds"] = state ? state.idleSessionTtlInSeconds : undefined;
            inputs["intents"] = state ? state.intents : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["locale"] = state ? state.locale : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nluIntentConfidenceThreshold"] = state ? state.nluIntentConfidenceThreshold : undefined;
            inputs["processBehavior"] = state ? state.processBehavior : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["voiceId"] = state ? state.voiceId : undefined;
        } else {
            const args = argsOrState as BotArgs | undefined;
            if ((!args || args.abortStatement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'abortStatement'");
            }
            if ((!args || args.childDirected === undefined) && !opts.urn) {
                throw new Error("Missing required property 'childDirected'");
            }
            if ((!args || args.intents === undefined) && !opts.urn) {
                throw new Error("Missing required property 'intents'");
            }
            inputs["abortStatement"] = args ? args.abortStatement : undefined;
            inputs["childDirected"] = args ? args.childDirected : undefined;
            inputs["clarificationPrompt"] = args ? args.clarificationPrompt : undefined;
            inputs["createVersion"] = args ? args.createVersion : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["detectSentiment"] = args ? args.detectSentiment : undefined;
            inputs["enableModelImprovements"] = args ? args.enableModelImprovements : undefined;
            inputs["idleSessionTtlInSeconds"] = args ? args.idleSessionTtlInSeconds : undefined;
            inputs["intents"] = args ? args.intents : undefined;
            inputs["locale"] = args ? args.locale : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nluIntentConfidenceThreshold"] = args ? args.nluIntentConfidenceThreshold : undefined;
            inputs["processBehavior"] = args ? args.processBehavior : undefined;
            inputs["voiceId"] = args ? args.voiceId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["checksum"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["failureReason"] = undefined /*out*/;
            inputs["lastUpdatedDate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Bot.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bot resources.
 */
export interface BotState {
    /**
     * The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
     */
    readonly abortStatement?: pulumi.Input<inputs.lex.BotAbortStatement>;
    readonly arn?: pulumi.Input<string>;
    /**
     * Checksum identifying the version of the bot that was created. The checksum is not
     * included as an argument because the resource will add it automatically when updating the bot.
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
     */
    readonly childDirected?: pulumi.Input<boolean>;
    /**
     * The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
     */
    readonly clarificationPrompt?: pulumi.Input<inputs.lex.BotClarificationPrompt>;
    /**
     * Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
     */
    readonly createVersion?: pulumi.Input<boolean>;
    /**
     * The date when the bot version was created.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * A description of the bot. Must be less than or equal to 200 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
     */
    readonly detectSentiment?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
     */
    readonly enableModelImprovements?: pulumi.Input<boolean>;
    /**
     * If status is FAILED, Amazon Lex provides the reason that it failed to build the bot.
     */
    readonly failureReason?: pulumi.Input<string>;
    /**
     * The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
     */
    readonly idleSessionTtlInSeconds?: pulumi.Input<number>;
    /**
     * A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 100 Intent objects.
     */
    readonly intents?: pulumi.Input<pulumi.Input<inputs.lex.BotIntent>[]>;
    /**
     * The date when the $LATEST version of this bot was updated.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
     */
    readonly locale?: pulumi.Input<string>;
    /**
     * The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
     */
    readonly nluIntentConfidenceThreshold?: pulumi.Input<number>;
    /**
     * If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
     */
    readonly processBehavior?: pulumi.Input<string>;
    /**
     * When you send a request to create or update a bot, Amazon Lex sets the status response
     * element to BUILDING. After Amazon Lex builds the bot, it sets status to READY. If Amazon Lex can't
     * build the bot, it sets status to FAILED. Amazon Lex returns the reason for the failure in the
     * failureReason response element.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The version of the bot.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
     */
    readonly voiceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Bot resource.
 */
export interface BotArgs {
    /**
     * The message that Amazon Lex uses to abort a conversation. Attributes are documented under statement.
     */
    readonly abortStatement: pulumi.Input<inputs.lex.BotAbortStatement>;
    /**
     * By specifying true, you confirm that your use of Amazon Lex is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA. For more information see the [Amazon Lex FAQ](https://aws.amazon.com/lex/faqs#data-security) and the [Amazon Lex PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-childDirected).
     */
    readonly childDirected: pulumi.Input<boolean>;
    /**
     * The message that Amazon Lex uses when it doesn't understand the user's request. Attributes are documented under prompt.
     */
    readonly clarificationPrompt?: pulumi.Input<inputs.lex.BotClarificationPrompt>;
    /**
     * Determines if a new bot version is created when the initial resource is created and on each update. Defaults to `false`.
     */
    readonly createVersion?: pulumi.Input<boolean>;
    /**
     * A description of the bot. Must be less than or equal to 200 characters in length.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When set to true user utterances are sent to Amazon Comprehend for sentiment analysis. If you don't specify detectSentiment, the default is `false`.
     */
    readonly detectSentiment?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable access to natural language understanding improvements. When you set the `enableModelImprovements` parameter to true you can use the `nluIntentConfidenceThreshold` parameter to configure confidence scores. For more information, see [Confidence Scores](https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html). You can only set the `enableModelImprovements` parameter in certain Regions. If you set the parameter to true, your bot has access to accuracy improvements. For more information see the [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-enableModelImprovements).
     */
    readonly enableModelImprovements?: pulumi.Input<boolean>;
    /**
     * The maximum time in seconds that Amazon Lex retains the data gathered in a conversation. Default is `300`. Must be a number between 60 and 86400 (inclusive).
     */
    readonly idleSessionTtlInSeconds?: pulumi.Input<number>;
    /**
     * A set of Intent objects. Each intent represents a command that a user can express. Attributes are documented under intent. Can have up to 100 Intent objects.
     */
    readonly intents: pulumi.Input<pulumi.Input<inputs.lex.BotIntent>[]>;
    /**
     * Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot. For available locales, see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-locale). Default is `en-US`.
     */
    readonly locale?: pulumi.Input<string>;
    /**
     * The name of the bot that you want to create, case sensitive. Must be between 2 and 50 characters in length.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Determines the threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot. For more information see [Amazon Lex Bot PutBot API Docs](https://docs.aws.amazon.com/lex/latest/dg/API_PutBot.html#lex-PutBot-request-nluIntentConfidenceThreshold) This value requires `enableModelImprovements` to be set to `true` and the default is `0`. Must be a float between 0 and 1.
     */
    readonly nluIntentConfidenceThreshold?: pulumi.Input<number>;
    /**
     * If you set the `processBehavior` element to `BUILD`, Amazon Lex builds the bot so that it can be run. If you set the element to `SAVE` Amazon Lex saves the bot, but doesn't build it. Default is `SAVE`.
     */
    readonly processBehavior?: pulumi.Input<string>;
    /**
     * The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions with the user. The locale configured for the voice must match the locale of the bot. For more information, see [Available Voices](http://docs.aws.amazon.com/polly/latest/dg/voicelist.html) in the Amazon Polly Developer Guide.
     */
    readonly voiceId?: pulumi.Input<string>;
}
