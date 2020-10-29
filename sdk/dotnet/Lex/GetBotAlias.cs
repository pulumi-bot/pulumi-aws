// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lex
{
    public static class GetBotAlias
    {
        /// <summary>
        /// Provides details about a specific Amazon Lex Bot Alias.
        /// </summary>
        public static Task<GetBotAliasResult> InvokeAsync(GetBotAliasArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBotAliasResult>("aws:lex/getBotAlias:getBotAlias", args ?? new GetBotAliasArgs(), options.WithVersion());
    }


    public sealed class GetBotAliasArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the bot.
        /// </summary>
        [Input("botName", required: true)]
        public string BotName { get; set; } = null!;

        /// <summary>
        /// The name of the bot alias. The name is case sensitive.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBotAliasArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBotAliasResult
    {
        /// <summary>
        /// The ARN of the bot alias.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The name of the bot.
        /// </summary>
        public readonly string BotName;
        /// <summary>
        /// The version of the bot that the alias points to.
        /// </summary>
        public readonly string BotVersion;
        /// <summary>
        /// Checksum of the bot alias.
        /// </summary>
        public readonly string Checksum;
        /// <summary>
        /// The date that the bot alias was created.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// A description of the alias.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The date that the bot alias was updated. When you create a resource, the creation date and the last updated date are the same.
        /// </summary>
        public readonly string LastUpdatedDate;
        /// <summary>
        /// The name of the alias. The name is not case sensitive.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetBotAliasResult(
            string arn,

            string botName,

            string botVersion,

            string checksum,

            string createdDate,

            string description,

            string id,

            string lastUpdatedDate,

            string name)
        {
            Arn = arn;
            BotName = botName;
            BotVersion = botVersion;
            Checksum = checksum;
            CreatedDate = createdDate;
            Description = description;
            Id = id;
            LastUpdatedDate = lastUpdatedDate;
            Name = name;
        }
    }
}
