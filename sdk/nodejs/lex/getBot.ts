// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about a specific Amazon Lex Bot.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const orderFlowersBot = pulumi.output(aws.lex.getBot({
 *     name: "OrderFlowers",
 *     version: "$LATEST",
 * }, { async: true }));
 * ```
 */
export function getBot(args: GetBotArgs, opts?: pulumi.InvokeOptions): Promise<GetBotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:lex/getBot:getBot", {
        "name": args.name,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getBot.
 */
export interface GetBotArgs {
    /**
     * The name of the bot. The name is case sensitive.
     */
    readonly name: string;
    /**
     * The version or alias of the bot.
     */
    readonly version?: string;
}

/**
 * A collection of values returned by getBot.
 */
export interface GetBotResult {
    readonly arn: string;
    readonly checksum: string;
    /**
     * Specifies if this Amazon Lex Bot is related to a website, program, or other application that is directed or targeted, in whole or in part, to children under age 13 and subject to COPPA.
     */
    readonly childDirected: boolean;
    readonly createdDate: string;
    /**
     * A description of the bot.
     */
    readonly description: string;
    /**
     * When set to true user utterances are sent to Amazon Comprehend for sentiment analysis.
     */
    readonly detectSentiment: boolean;
    /**
     * Set to true if natural language understanding improvements are enabled.
     */
    readonly enableModelImprovements: boolean;
    readonly failureReason: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The maximum time in seconds that Amazon Lex retains the data gathered in a conversation.
     */
    readonly idleSessionTtlInSeconds: number;
    readonly lastUpdatedDate: string;
    /**
     * Specifies the target locale for the bot. Any intent used in the bot must be compatible with the locale of the bot.
     */
    readonly locale: string;
    /**
     * The name of the bot, case sensitive.
     */
    readonly name: string;
    /**
     * The threshold where Amazon Lex will insert the AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning alternative intents in a PostContent or PostText response. AMAZON.FallbackIntent and AMAZON.KendraSearchIntent are only inserted if they are configured for the bot.
     */
    readonly nluIntentConfidenceThreshold: number;
    readonly status: string;
    readonly version?: string;
    /**
     * The Amazon Polly voice ID that the Amazon Lex Bot uses for voice interactions with the user.
     */
    readonly voiceId: string;
}
