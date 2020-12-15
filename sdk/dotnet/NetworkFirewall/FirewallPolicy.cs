// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.NetworkFirewall
{
    /// <summary>
    /// Provides an AWS Network Firewall Firewall Policy Resource
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
    ///         var example = new Aws.NetworkFirewall.FirewallPolicy("example", new Aws.NetworkFirewall.FirewallPolicyArgs
    ///         {
    ///             FirewallPolicy = new Aws.NetworkFirewall.Inputs.FirewallPolicyFirewallPolicyArgs
    ///             {
    ///                 StatelessDefaultActions = 
    ///                 {
    ///                     "aws:pass",
    ///                 },
    ///                 StatelessFragmentDefaultActions = 
    ///                 {
    ///                     "aws:drop",
    ///                 },
    ///                 StatelessRuleGroupReferences = 
    ///                 {
    ///                     new Aws.NetworkFirewall.Inputs.FirewallPolicyFirewallPolicyStatelessRuleGroupReferenceArgs
    ///                     {
    ///                         Priority = 1,
    ///                         ResourceArn = aws_networkfirewall_rule_group.Example.Arn,
    ///                     },
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "Tag1", "Value1" },
    ///                 { "Tag2", "Value2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Policy with a Custom Action for Stateless Inspection
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var test = new Aws.NetworkFirewall.FirewallPolicy("test", new Aws.NetworkFirewall.FirewallPolicyArgs
    ///         {
    ///             FirewallPolicy = new Aws.NetworkFirewall.Inputs.FirewallPolicyFirewallPolicyArgs
    ///             {
    ///                 StatelessCustomActions = 
    ///                 {
    ///                     new Aws.NetworkFirewall.Inputs.FirewallPolicyFirewallPolicyStatelessCustomActionArgs
    ///                     {
    ///                         ActionDefinition = new Aws.NetworkFirewall.Inputs.FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinitionArgs
    ///                         {
    ///                             PublishMetricAction = new Aws.NetworkFirewall.Inputs.FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinitionPublishMetricActionArgs
    ///                             {
    ///                                 Dimension = 
    ///                                 {
    ///                                     
    ///                                     {
    ///                                         { "value", "1" },
    ///                                     },
    ///                                 },
    ///                             },
    ///                         },
    ///                         ActionName = "ExampleCustomAction",
    ///                     },
    ///                 },
    ///                 StatelessDefaultActions = 
    ///                 {
    ///                     "aws:pass",
    ///                     "ExampleCustomAction",
    ///                 },
    ///                 StatelessFragmentDefaultActions = 
    ///                 {
    ///                     "aws:drop",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Network Firewall Policies can be imported using their `ARN`.
    /// 
    /// ```sh
    ///  $ pulumi import aws:networkfirewall/firewallPolicy:FirewallPolicy example arn:aws:network-firewall:us-west-1:123456789012:firewall-policy/example
    /// ```
    /// </summary>
    [AwsResourceType("aws:networkfirewall/firewallPolicy:FirewallPolicy")]
    public partial class FirewallPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the firewall policy.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A friendly description of the firewall policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
        /// </summary>
        [Output("firewallPolicy")]
        public Output<Outputs.FirewallPolicyFirewallPolicy> FirewallPolicyConfiguration { get; private set; } = null!;

        /// <summary>
        /// A friendly name of the firewall policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of key:value pairs to associate with the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A string token used when updating a firewall policy.
        /// </summary>
        [Output("updateToken")]
        public Output<string> UpdateToken { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallPolicy(string name, FirewallPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:networkfirewall/firewallPolicy:FirewallPolicy", name, args ?? new FirewallPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallPolicy(string name, Input<string> id, FirewallPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:networkfirewall/firewallPolicy:FirewallPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallPolicy Get(string name, Input<string> id, FirewallPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallPolicy(name, id, state, options);
        }
    }

    public sealed class FirewallPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A friendly description of the firewall policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
        /// </summary>
        [Input("firewallPolicy", required: true)]
        public Input<Inputs.FirewallPolicyFirewallPolicyArgs> FirewallPolicyConfiguration { get; set; } = null!;

        /// <summary>
        /// A friendly name of the firewall policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// An array of key:value pairs to associate with the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FirewallPolicyArgs()
        {
        }
    }

    public sealed class FirewallPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the firewall policy.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A friendly description of the firewall policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
        /// </summary>
        [Input("firewallPolicy")]
        public Input<Inputs.FirewallPolicyFirewallPolicyGetArgs>? FirewallPolicyConfiguration { get; set; }

        /// <summary>
        /// A friendly name of the firewall policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// An array of key:value pairs to associate with the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A string token used when updating a firewall policy.
        /// </summary>
        [Input("updateToken")]
        public Input<string>? UpdateToken { get; set; }

        public FirewallPolicyState()
        {
        }
    }
}
