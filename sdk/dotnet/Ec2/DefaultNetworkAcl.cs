// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to manage the default AWS Network ACL. VPC Only.
    /// 
    /// Each VPC created in AWS comes with a Default Network ACL that can be managed, but not
    /// destroyed. **This is an advanced resource**, and has special caveats to be aware
    /// of when using it. Please read this document in its entirety before using this
    /// resource.
    /// 
    /// The `aws.ec2.DefaultNetworkAcl` behaves differently from normal resources, in that
    /// this provider does not _create_ this resource, but instead attempts to "adopt" it
    /// into management. We can do this because each VPC created has a Default Network
    /// ACL that cannot be destroyed, and is created with a known set of default rules.
    /// 
    /// When this provider first adopts the Default Network ACL, it **immediately removes all
    /// rules in the ACL**. It then proceeds to create any rules specified in the
    /// configuration. This step is required so that only the rules specified in the
    /// configuration are created.
    /// 
    /// This resource treats its inline rules as absolute; only the rules defined
    /// inline are created, and any additions/removals external to this resource will
    /// result in diffs being shown. For these reasons, this resource is incompatible with the
    /// `aws.ec2.NetworkAclRule` resource.
    /// 
    /// For more information about Network ACLs, see the AWS Documentation on
    /// [Network ACLs][aws-network-acls].
    /// 
    /// ## Basic Example Usage, with default rules
    /// 
    /// The following config gives the Default Network ACL the same rules that AWS
    /// includes, but pulls the resource under management by this provider. This means that
    /// any ACL rules added or changed will be detected as drift.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mainvpc = new Aws.Ec2.Vpc("mainvpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.1.0.0/16",
    ///         });
    ///         var @default = new Aws.Ec2.DefaultNetworkAcl("default", new Aws.Ec2.DefaultNetworkAclArgs
    ///         {
    ///             DefaultNetworkAclId = mainvpc.DefaultNetworkAclId,
    ///             Ingress = 
    ///             {
    ///                 new Aws.Ec2.Inputs.DefaultNetworkAclIngressArgs
    ///                 {
    ///                     Protocol = "-1",
    ///                     RuleNo = 100,
    ///                     Action = "allow",
    ///                     CidrBlock = mainvpc.CidrBlock,
    ///                     FromPort = 0,
    ///                     ToPort = 0,
    ///                 },
    ///             },
    ///             Egress = 
    ///             {
    ///                 new Aws.Ec2.Inputs.DefaultNetworkAclEgressArgs
    ///                 {
    ///                     Protocol = "-1",
    ///                     RuleNo = 100,
    ///                     Action = "allow",
    ///                     CidrBlock = "0.0.0.0/0",
    ///                     FromPort = 0,
    ///                     ToPort = 0,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Example config to deny all Egress traffic, allowing Ingress
    /// 
    /// The following denies all Egress traffic by omitting any `egress` rules, while
    /// including the default `ingress` rule to allow all traffic.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mainvpc = new Aws.Ec2.Vpc("mainvpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.1.0.0/16",
    ///         });
    ///         var @default = new Aws.Ec2.DefaultNetworkAcl("default", new Aws.Ec2.DefaultNetworkAclArgs
    ///         {
    ///             DefaultNetworkAclId = mainvpc.DefaultNetworkAclId,
    ///             Ingress = 
    ///             {
    ///                 new Aws.Ec2.Inputs.DefaultNetworkAclIngressArgs
    ///                 {
    ///                     Protocol = "-1",
    ///                     RuleNo = 100,
    ///                     Action = "allow",
    ///                     CidrBlock = mainvpc.CidrBlock,
    ///                     FromPort = 0,
    ///                     ToPort = 0,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Example config to deny all traffic to any Subnet in the Default Network ACL
    /// 
    /// This config denies all traffic in the Default ACL. This can be useful if you
    /// want a locked down default to force all resources in the VPC to assign a
    /// non-default ACL.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mainvpc = new Aws.Ec2.Vpc("mainvpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.1.0.0/16",
    ///         });
    ///         var @default = new Aws.Ec2.DefaultNetworkAcl("default", new Aws.Ec2.DefaultNetworkAclArgs
    ///         {
    ///             DefaultNetworkAclId = mainvpc.DefaultNetworkAclId,
    ///         });
    ///         // no rules defined, deny all traffic in this ACL
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Default Network ACLs can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2/defaultNetworkAcl:DefaultNetworkAcl sample acl-7aaabd18
    /// ```
    /// </summary>
    public partial class DefaultNetworkAcl : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Default Network ACL
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Network ACL ID to manage. This
        /// attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
        /// </summary>
        [Output("defaultNetworkAclId")]
        public Output<string> DefaultNetworkAclId { get; private set; } = null!;

        /// <summary>
        /// Specifies an egress rule. Parameters defined below.
        /// </summary>
        [Output("egress")]
        public Output<ImmutableArray<Outputs.DefaultNetworkAclEgress>> Egress { get; private set; } = null!;

        /// <summary>
        /// Specifies an ingress rule. Parameters defined below.
        /// </summary>
        [Output("ingress")]
        public Output<ImmutableArray<Outputs.DefaultNetworkAclIngress>> Ingress { get; private set; } = null!;

        /// <summary>
        /// The ID of the AWS account that owns the Default Network ACL
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// A list of Subnet IDs to apply the ACL to. See the
        /// notes below on managing Subnets in the Default Network ACL
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated VPC
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultNetworkAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultNetworkAcl(string name, DefaultNetworkAclArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, args ?? new DefaultNetworkAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultNetworkAcl(string name, Input<string> id, DefaultNetworkAclState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultNetworkAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultNetworkAcl Get(string name, Input<string> id, DefaultNetworkAclState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultNetworkAcl(name, id, state, options);
        }
    }

    public sealed class DefaultNetworkAclArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Network ACL ID to manage. This
        /// attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
        /// </summary>
        [Input("defaultNetworkAclId", required: true)]
        public Input<string> DefaultNetworkAclId { get; set; } = null!;

        [Input("egress")]
        private InputList<Inputs.DefaultNetworkAclEgressArgs>? _egress;

        /// <summary>
        /// Specifies an egress rule. Parameters defined below.
        /// </summary>
        public InputList<Inputs.DefaultNetworkAclEgressArgs> Egress
        {
            get => _egress ?? (_egress = new InputList<Inputs.DefaultNetworkAclEgressArgs>());
            set => _egress = value;
        }

        [Input("ingress")]
        private InputList<Inputs.DefaultNetworkAclIngressArgs>? _ingress;

        /// <summary>
        /// Specifies an ingress rule. Parameters defined below.
        /// </summary>
        public InputList<Inputs.DefaultNetworkAclIngressArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Inputs.DefaultNetworkAclIngressArgs>());
            set => _ingress = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of Subnet IDs to apply the ACL to. See the
        /// notes below on managing Subnets in the Default Network ACL
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DefaultNetworkAclArgs()
        {
        }
    }

    public sealed class DefaultNetworkAclState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Default Network ACL
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Network ACL ID to manage. This
        /// attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
        /// </summary>
        [Input("defaultNetworkAclId")]
        public Input<string>? DefaultNetworkAclId { get; set; }

        [Input("egress")]
        private InputList<Inputs.DefaultNetworkAclEgressGetArgs>? _egress;

        /// <summary>
        /// Specifies an egress rule. Parameters defined below.
        /// </summary>
        public InputList<Inputs.DefaultNetworkAclEgressGetArgs> Egress
        {
            get => _egress ?? (_egress = new InputList<Inputs.DefaultNetworkAclEgressGetArgs>());
            set => _egress = value;
        }

        [Input("ingress")]
        private InputList<Inputs.DefaultNetworkAclIngressGetArgs>? _ingress;

        /// <summary>
        /// Specifies an ingress rule. Parameters defined below.
        /// </summary>
        public InputList<Inputs.DefaultNetworkAclIngressGetArgs> Ingress
        {
            get => _ingress ?? (_ingress = new InputList<Inputs.DefaultNetworkAclIngressGetArgs>());
            set => _ingress = value;
        }

        /// <summary>
        /// The ID of the AWS account that owns the Default Network ACL
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of Subnet IDs to apply the ACL to. See the
        /// notes below on managing Subnets in the Default Network ACL
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the associated VPC
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public DefaultNetworkAclState()
        {
        }
    }
}
