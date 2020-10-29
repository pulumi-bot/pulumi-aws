// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lex
{
    public static class GetIntent
    {
        /// <summary>
        /// Provides details about a specific Amazon Lex Intent.
        /// </summary>
        public static Task<GetIntentResult> InvokeAsync(GetIntentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntentResult>("aws:lex/getIntent:getIntent", args ?? new GetIntentArgs(), options.WithVersion());
    }


    public sealed class GetIntentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the intent. The name is case sensitive.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The version of the intent.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetIntentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntentResult
    {
        /// <summary>
        /// The ARN of the Lex intent.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Checksum identifying the version of the intent that was created. The checksum is not
        /// included as an argument because the resource will add it automatically when updating the intent.
        /// </summary>
        public readonly string Checksum;
        /// <summary>
        /// The date when the intent version was created.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// A description of the intent.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The date when the $LATEST version of this intent was updated.
        /// </summary>
        public readonly string LastUpdatedDate;
        /// <summary>
        /// The name of the intent, not case sensitive.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A unique identifier for the built-in intent to base this
        /// intent on. To find the signature for an intent, see
        /// [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
        /// in the Alexa Skills Kit.
        /// </summary>
        public readonly string ParentIntentSignature;
        /// <summary>
        /// The version of the bot.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetIntentResult(
            string arn,

            string checksum,

            string createdDate,

            string description,

            string id,

            string lastUpdatedDate,

            string name,

            string parentIntentSignature,

            string? version)
        {
            Arn = arn;
            Checksum = checksum;
            CreatedDate = createdDate;
            Description = description;
            Id = id;
            LastUpdatedDate = lastUpdatedDate;
            Name = name;
            ParentIntentSignature = parentIntentSignature;
            Version = version;
        }
    }
}
