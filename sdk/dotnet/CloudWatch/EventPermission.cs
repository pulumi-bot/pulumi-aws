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
    /// Provides a resource to create a CloudWatch Events permission to support cross-account events in the current account default event bus.
    /// 
    /// ## Example Usage
    /// 
    /// ### Account Access
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var devAccountAccess = new Aws.CloudWatch.EventPermission("devAccountAccess", new Aws.CloudWatch.EventPermissionArgs
    ///         {
    ///             Principal = "123456789012",
    ///             StatementId = "DevAccountAccess",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Organization Access
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var organizationAccess = new Aws.CloudWatch.EventPermission("organizationAccess", new Aws.CloudWatch.EventPermissionArgs
    ///         {
    ///             Condition = new Aws.CloudWatch.Inputs.EventPermissionConditionArgs
    ///             {
    ///                 Key = "aws:PrincipalOrgID",
    ///                 Type = "StringEquals",
    ///                 Value = aws_organizations_organization.Example.Id,
    ///             },
    ///             Principal = "*",
    ///             StatementId = "OrganizationAccess",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class EventPermission : Pulumi.CustomResource
    {
        /// <summary>
        /// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.EventPermissionCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// An identifier string for the external account that you are granting permissions to.
        /// </summary>
        [Output("statementId")]
        public Output<string> StatementId { get; private set; } = null!;


        /// <summary>
        /// Create a EventPermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventPermission(string name, EventPermissionArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventPermission:EventPermission", name, args ?? new EventPermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventPermission(string name, Input<string> id, EventPermissionState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventPermission:EventPermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventPermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventPermission Get(string name, Input<string> id, EventPermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new EventPermission(name, id, state, options);
        }
    }

    public sealed class EventPermissionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.EventPermissionConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// An identifier string for the external account that you are granting permissions to.
        /// </summary>
        [Input("statementId", required: true)]
        public Input<string> StatementId { get; set; } = null!;

        public EventPermissionArgs()
        {
        }
    }

    public sealed class EventPermissionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.EventPermissionConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// An identifier string for the external account that you are granting permissions to.
        /// </summary>
        [Input("statementId")]
        public Input<string>? StatementId { get; set; }

        public EventPermissionState()
        {
        }
    }
}
