// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WorkLink
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// Basic usage:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.WorkLink.Fleet("example", new Aws.WorkLink.FleetArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Network Configuration Usage:
    /// 
    /// ```csharp
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.WorkLink.Fleet("example", new Aws.WorkLink.FleetArgs
    ///         {
    ///             Network = new Aws.WorkLink.Inputs.FleetNetworkArgs
    ///             {
    ///                 SecurityGroupIds = 
    ///                 {
    ///                     aws_security_group.Test.Id,
    ///                 },
    ///                 SubnetIds = 
    ///                 {
    ///                     aws_subnet.Test.Select(__item =&gt; __item.Id).ToList(),
    ///                 },
    ///                 VpcId = aws_vpc.Test.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Identity Provider Configuration Usage:
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var test = new Aws.WorkLink.Fleet("test", new Aws.WorkLink.FleetArgs
    ///         {
    ///             IdentityProvider = new Aws.WorkLink.Inputs.FleetIdentityProviderArgs
    ///             {
    ///                 SamlMetadata = File.ReadAllText("saml-metadata.xml"),
    ///                 Type = "SAML",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Fleet : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the created WorkLink Fleet.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Amazon Kinesis data stream that receives the audit events.
        /// </summary>
        [Output("auditStreamArn")]
        public Output<string?> AuditStreamArn { get; private set; } = null!;

        /// <summary>
        /// The identifier used by users to sign in to the Amazon WorkLink app.
        /// </summary>
        [Output("companyCode")]
        public Output<string> CompanyCode { get; private set; } = null!;

        /// <summary>
        /// The time that the fleet was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        /// </summary>
        [Output("deviceCaCertificate")]
        public Output<string?> DeviceCaCertificate { get; private set; } = null!;

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        /// </summary>
        [Output("identityProvider")]
        public Output<Outputs.FleetIdentityProvider?> IdentityProvider { get; private set; } = null!;

        /// <summary>
        /// The time that the fleet was last updated.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        /// </summary>
        [Output("network")]
        public Output<Outputs.FleetNetwork?> Network { get; private set; } = null!;

        /// <summary>
        /// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        /// </summary>
        [Output("optimizeForEndUserLocation")]
        public Output<bool?> OptimizeForEndUserLocation { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:worklink/fleet:Fleet", name, args ?? new FleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
            : base("aws:worklink/fleet:Fleet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fleet Get(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
        {
            return new Fleet(name, id, state, options);
        }
    }

    public sealed class FleetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon Kinesis data stream that receives the audit events.
        /// </summary>
        [Input("auditStreamArn")]
        public Input<string>? AuditStreamArn { get; set; }

        /// <summary>
        /// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        /// </summary>
        [Input("deviceCaCertificate")]
        public Input<string>? DeviceCaCertificate { get; set; }

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        /// </summary>
        [Input("identityProvider")]
        public Input<Inputs.FleetIdentityProviderArgs>? IdentityProvider { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        /// </summary>
        [Input("network")]
        public Input<Inputs.FleetNetworkArgs>? Network { get; set; }

        /// <summary>
        /// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        /// </summary>
        [Input("optimizeForEndUserLocation")]
        public Input<bool>? OptimizeForEndUserLocation { get; set; }

        public FleetArgs()
        {
        }
    }

    public sealed class FleetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the created WorkLink Fleet.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The ARN of the Amazon Kinesis data stream that receives the audit events.
        /// </summary>
        [Input("auditStreamArn")]
        public Input<string>? AuditStreamArn { get; set; }

        /// <summary>
        /// The identifier used by users to sign in to the Amazon WorkLink app.
        /// </summary>
        [Input("companyCode")]
        public Input<string>? CompanyCode { get; set; }

        /// <summary>
        /// The time that the fleet was created.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        /// </summary>
        [Input("deviceCaCertificate")]
        public Input<string>? DeviceCaCertificate { get; set; }

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        /// </summary>
        [Input("identityProvider")]
        public Input<Inputs.FleetIdentityProviderGetArgs>? IdentityProvider { get; set; }

        /// <summary>
        /// The time that the fleet was last updated.
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        /// </summary>
        [Input("network")]
        public Input<Inputs.FleetNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        /// </summary>
        [Input("optimizeForEndUserLocation")]
        public Input<bool>? OptimizeForEndUserLocation { get; set; }

        public FleetState()
        {
        }
    }
}
