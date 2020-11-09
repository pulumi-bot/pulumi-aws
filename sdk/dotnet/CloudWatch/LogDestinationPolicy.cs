// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides a CloudWatch Logs destination policy resource.
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
    ///         var testDestination = new Aws.CloudWatch.LogDestination("testDestination", new Aws.CloudWatch.LogDestinationArgs
    ///         {
    ///             RoleArn = aws_iam_role.Iam_for_cloudwatch.Arn,
    ///             TargetArn = aws_kinesis_stream.Kinesis_for_cloudwatch.Arn,
    ///         });
    ///         var testDestinationPolicyPolicyDocument = testDestination.Arn.Apply(arn =&gt; Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Effect = "Allow",
    ///                     Principals = 
    ///                     {
    ///                         new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                         {
    ///                             Type = "AWS",
    ///                             Identifiers = 
    ///                             {
    ///                                 "123456789012",
    ///                             },
    ///                         },
    ///                     },
    ///                     Actions = 
    ///                     {
    ///                         "logs:PutSubscriptionFilter",
    ///                     },
    ///                     Resources = 
    ///                     {
    ///                         arn,
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var testDestinationPolicyLogDestinationPolicy = new Aws.CloudWatch.LogDestinationPolicy("testDestinationPolicyLogDestinationPolicy", new Aws.CloudWatch.LogDestinationPolicyArgs
    ///         {
    ///             DestinationName = testDestination.Name,
    ///             AccessPolicy = testDestinationPolicyPolicyDocument.Apply(testDestinationPolicyPolicyDocument =&gt; testDestinationPolicyPolicyDocument.Json),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudWatch Logs destination policies can be imported using the `destination_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy test_destination_policy test_destination
    /// ```
    /// </summary>
    public partial class LogDestinationPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Output("accessPolicy")]
        public Output<string> AccessPolicy { get; private set; } = null!;

        /// <summary>
        /// A name for the subscription filter
        /// </summary>
        [Output("destinationName")]
        public Output<string> DestinationName { get; private set; } = null!;


        /// <summary>
        /// Create a LogDestinationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogDestinationPolicy(string name, LogDestinationPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy", name, args ?? new LogDestinationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogDestinationPolicy(string name, Input<string> id, LogDestinationPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogDestinationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogDestinationPolicy Get(string name, Input<string> id, LogDestinationPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LogDestinationPolicy(name, id, state, options);
        }
    }

    public sealed class LogDestinationPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("accessPolicy", required: true)]
        public Input<string> AccessPolicy { get; set; } = null!;

        /// <summary>
        /// A name for the subscription filter
        /// </summary>
        [Input("destinationName", required: true)]
        public Input<string> DestinationName { get; set; } = null!;

        public LogDestinationPolicyArgs()
        {
        }
    }

    public sealed class LogDestinationPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("accessPolicy")]
        public Input<string>? AccessPolicy { get; set; }

        /// <summary>
        /// A name for the subscription filter
        /// </summary>
        [Input("destinationName")]
        public Input<string>? DestinationName { get; set; }

        public LogDestinationPolicyState()
        {
        }
    }
}
