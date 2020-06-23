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
    /// Provides a resource to create an association between a VPC endpoint and a subnet.
    /// 
    /// &gt; **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** This provider provides
    /// both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
    /// and a single `subnet_id`) and a VPC Endpoint resource with a `subnet_ids`
    /// attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
    /// Association resource. Doing so will cause a conflict of associations and will overwrite the association.
    /// 
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
    ///         var snEc2 = new Aws.Ec2.VpcEndpointSubnetAssociation("snEc2", new Aws.Ec2.VpcEndpointSubnetAssociationArgs
    ///         {
    ///             SubnetId = aws_subnet.Sn.Id,
    ///             VpcEndpointId = aws_vpc_endpoint.Ec2.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class VpcEndpointSubnetAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the subnet to be associated with the VPC endpoint.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC endpoint with which the subnet will be associated.
        /// </summary>
        [Output("vpcEndpointId")]
        public Output<string> VpcEndpointId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointSubnetAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointSubnetAssociation(string name, VpcEndpointSubnetAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation", name, args ?? new VpcEndpointSubnetAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointSubnetAssociation(string name, Input<string> id, VpcEndpointSubnetAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointSubnetAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointSubnetAssociation Get(string name, Input<string> id, VpcEndpointSubnetAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointSubnetAssociation(name, id, state, options);
        }
    }

    public sealed class VpcEndpointSubnetAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the subnet to be associated with the VPC endpoint.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// The ID of the VPC endpoint with which the subnet will be associated.
        /// </summary>
        [Input("vpcEndpointId", required: true)]
        public Input<string> VpcEndpointId { get; set; } = null!;

        public VpcEndpointSubnetAssociationArgs()
        {
        }
    }

    public sealed class VpcEndpointSubnetAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the subnet to be associated with the VPC endpoint.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The ID of the VPC endpoint with which the subnet will be associated.
        /// </summary>
        [Input("vpcEndpointId")]
        public Input<string>? VpcEndpointId { get; set; }

        public VpcEndpointSubnetAssociationState()
        {
        }
    }
}
