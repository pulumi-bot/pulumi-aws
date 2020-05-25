// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DirectoryService
{
    /// <summary>
    /// Provides a Log subscription for AWS Directory Service that pushes logs to cloudwatch.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
    ///         {
    ///             RetentionInDays = 14,
    ///         });
    ///         var ad_log_policyPolicyDocument = exampleLogGroup.Arn.Apply(arn =&gt; Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
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
    ///                     Effect = "Allow",
    ///                     Principals = 
    ///                     {
    ///                         new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                         {
    ///                             Identifiers = 
    ///                             {
    ///                                 "ds.amazonaws.com",
    ///                             },
    ///                             Type = "Service",
    ///                         },
    ///                     },
    ///                     Resources = 
    ///                     {
    ///                         arn,
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var ad-log-policyLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("ad-log-policyLogResourcePolicy", new Aws.CloudWatch.LogResourcePolicyArgs
    ///         {
    ///             PolicyDocument = ad-log-policyPolicyDocument.Apply(ad_log_policyPolicyDocument =&gt; ad_log_policyPolicyDocument.Json),
    ///             PolicyName = "ad-log-policy",
    ///         });
    ///         var exampleLogService = new Aws.DirectoryService.LogService("exampleLogService", new Aws.DirectoryService.LogServiceArgs
    ///         {
    ///             DirectoryId = aws_directory_service_directory.Example.Id,
    ///             LogGroupName = exampleLogGroup.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class LogService : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of directory.
        /// </summary>
        [Output("directoryId")]
        public Output<string> DirectoryId { get; private set; } = null!;

        /// <summary>
        /// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
        /// </summary>
        [Output("logGroupName")]
        public Output<string> LogGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a LogService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogService(string name, LogServiceArgs args, CustomResourceOptions? options = null)
            : base("aws:directoryservice/logService:LogService", name, args ?? new LogServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogService(string name, Input<string> id, LogServiceState? state = null, CustomResourceOptions? options = null)
            : base("aws:directoryservice/logService:LogService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogService Get(string name, Input<string> id, LogServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new LogService(name, id, state, options);
        }
    }

    public sealed class LogServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of directory.
        /// </summary>
        [Input("directoryId", required: true)]
        public Input<string> DirectoryId { get; set; } = null!;

        /// <summary>
        /// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
        /// </summary>
        [Input("logGroupName", required: true)]
        public Input<string> LogGroupName { get; set; } = null!;

        public LogServiceArgs()
        {
        }
    }

    public sealed class LogServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of directory.
        /// </summary>
        [Input("directoryId")]
        public Input<string>? DirectoryId { get; set; }

        /// <summary>
        /// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
        /// </summary>
        [Input("logGroupName")]
        public Input<string>? LogGroupName { get; set; }

        public LogServiceState()
        {
        }
    }
}
