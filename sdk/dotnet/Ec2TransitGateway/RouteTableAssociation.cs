// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2TransitGateway
{
    /// <summary>
    /// Manages an EC2 Transit Gateway Route Table association.
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
    ///         var example = new Aws.Ec2TransitGateway.RouteTableAssociation("example", new Aws.Ec2TransitGateway.RouteTableAssociationArgs
    ///         {
    ///             TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
    ///             TransitGatewayRouteTableId = aws_ec2_transit_gateway_route_table.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_ec2_transit_gateway_route_table_association` can be imported by using the EC2 Transit Gateway Route Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation example tgw-rtb-12345678_tgw-attach-87654321
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation")]
    public partial class RouteTableAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Type of the resource
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway Attachment.
        /// </summary>
        [Output("transitGatewayAttachmentId")]
        public Output<string> TransitGatewayAttachmentId { get; private set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway Route Table.
        /// </summary>
        [Output("transitGatewayRouteTableId")]
        public Output<string> TransitGatewayRouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteTableAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteTableAssociation(string name, RouteTableAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation", name, args ?? new RouteTableAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteTableAssociation(string name, Input<string> id, RouteTableAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteTableAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteTableAssociation Get(string name, Input<string> id, RouteTableAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteTableAssociation(name, id, state, options);
        }
    }

    public sealed class RouteTableAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of EC2 Transit Gateway Attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId", required: true)]
        public Input<string> TransitGatewayAttachmentId { get; set; } = null!;

        /// <summary>
        /// Identifier of EC2 Transit Gateway Route Table.
        /// </summary>
        [Input("transitGatewayRouteTableId", required: true)]
        public Input<string> TransitGatewayRouteTableId { get; set; } = null!;

        public RouteTableAssociationArgs()
        {
        }
    }

    public sealed class RouteTableAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Type of the resource
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway Attachment.
        /// </summary>
        [Input("transitGatewayAttachmentId")]
        public Input<string>? TransitGatewayAttachmentId { get; set; }

        /// <summary>
        /// Identifier of EC2 Transit Gateway Route Table.
        /// </summary>
        [Input("transitGatewayRouteTableId")]
        public Input<string>? TransitGatewayRouteTableId { get; set; }

        public RouteTableAssociationState()
        {
        }
    }
}
