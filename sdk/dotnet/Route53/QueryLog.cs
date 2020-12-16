// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    /// <summary>
    /// Provides a Route53 query logging configuration resource.
    /// 
    /// &gt; **NOTE:** There are restrictions on the configuration of query logging. Notably,
    /// the CloudWatch log group must be in the `us-east-1` region,
    /// a permissive CloudWatch log resource policy must be in place, and
    /// the Route53 hosted zone must be public.
    /// See [Configuring Logging for DNS Queries](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html?console_help=true#query-logs-configuring) for additional details.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Example CloudWatch log group in us-east-1
    ///         var us_east_1 = new Aws.Provider("us-east-1", new Aws.ProviderArgs
    ///         {
    ///             Region = "us-east-1",
    ///         });
    ///         var awsRoute53ExampleCom = new Aws.CloudWatch.LogGroup("awsRoute53ExampleCom", new Aws.CloudWatch.LogGroupArgs
    ///         {
    ///             RetentionInDays = 30,
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Us_east_1,
    ///         });
    ///         // Example CloudWatch log resource policy to allow Route53 to write logs
    ///         // to any log group under /aws/route53/*
    ///         var route53_query_logging_policyPolicyDocument = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Actions = 
    ///                     {
    ///                         "logs:CreateLogStream",
    ///                         "logs:PutLogEvents",
    ///                     },
    ///                     Resources = 
    ///                     {
    ///                         "arn:aws:logs:*:*:log-group:/aws/route53/*",
    ///                     },
    ///                     Principals = 
    ///                     {
    ///                         new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                         {
    ///                             Identifiers = 
    ///                             {
    ///                                 "route53.amazonaws.com",
    ///                             },
    ///                             Type = "Service",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var route53_query_logging_policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("route53-query-logging-policyLogResourcePolicy", new Aws.CloudWatch.LogResourcePolicyArgs
    ///         {
    ///             PolicyDocument = route53_query_logging_policyPolicyDocument.Apply(route53_query_logging_policyPolicyDocument =&gt; route53_query_logging_policyPolicyDocument.Json),
    ///             PolicyName = "route53-query-logging-policy",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = aws.Us_east_1,
    ///         });
    ///         // Example Route53 zone with query logging
    ///         var exampleComZone = new Aws.Route53.Zone("exampleComZone", new Aws.Route53.ZoneArgs
    ///         {
    ///         });
    ///         var exampleComQueryLog = new Aws.Route53.QueryLog("exampleComQueryLog", new Aws.Route53.QueryLogArgs
    ///         {
    ///             CloudwatchLogGroupArn = awsRoute53ExampleCom.Arn,
    ///             ZoneId = exampleComZone.ZoneId,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 route53_query_logging_policyLogResourcePolicy,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Route53 query logging configurations can be imported using their ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:route53/queryLog:QueryLog example_com xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
    /// ```
    /// </summary>
    [AwsResourceType("aws:route53/queryLog:QueryLog")]
    public partial class QueryLog : Pulumi.CustomResource
    {
        /// <summary>
        /// CloudWatch log group ARN to send query logs.
        /// </summary>
        [Output("cloudwatchLogGroupArn")]
        public Output<string> CloudwatchLogGroupArn { get; private set; } = null!;

        /// <summary>
        /// Route53 hosted zone ID to enable query logs.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a QueryLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueryLog(string name, QueryLogArgs args, CustomResourceOptions? options = null)
            : base("aws:route53/queryLog:QueryLog", name, args ?? new QueryLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QueryLog(string name, Input<string> id, QueryLogState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/queryLog:QueryLog", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing QueryLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueryLog Get(string name, Input<string> id, QueryLogState? state = null, CustomResourceOptions? options = null)
        {
            return new QueryLog(name, id, state, options);
        }
    }

    public sealed class QueryLogArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CloudWatch log group ARN to send query logs.
        /// </summary>
        [Input("cloudwatchLogGroupArn", required: true)]
        public Input<string> CloudwatchLogGroupArn { get; set; } = null!;

        /// <summary>
        /// Route53 hosted zone ID to enable query logs.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public QueryLogArgs()
        {
        }
    }

    public sealed class QueryLogState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CloudWatch log group ARN to send query logs.
        /// </summary>
        [Input("cloudwatchLogGroupArn")]
        public Input<string>? CloudwatchLogGroupArn { get; set; }

        /// <summary>
        /// Route53 hosted zone ID to enable query logs.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public QueryLogState()
        {
        }
    }
}
