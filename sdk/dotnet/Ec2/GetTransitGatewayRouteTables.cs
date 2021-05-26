// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetTransitGatewayRouteTables
    {
        /// <summary>
        /// Provides information for multiple EC2 Transit Gateway Route Tables, such as their identifiers.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following shows outputing all Transit Gateway Route Table Ids.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleTransitGatewayRouteTables = Output.Create(Aws.Ec2.GetTransitGatewayRouteTables.InvokeAsync());
        ///         this.Example = data.Aws_ec2_transit_gateway_route_table.Example.Ids;
        ///     }
        /// 
        ///     [Output("example")]
        ///     public Output&lt;string&gt; Example { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTransitGatewayRouteTablesResult> InvokeAsync(GetTransitGatewayRouteTablesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayRouteTablesResult>("aws:ec2/getTransitGatewayRouteTables:getTransitGatewayRouteTables", args ?? new GetTransitGatewayRouteTablesArgs(), options.WithVersion());

        public static Output<GetTransitGatewayRouteTablesResult> Apply(GetTransitGatewayRouteTablesApplyArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetTransitGatewayRouteTablesApplyArgs();
            return Pulumi.Output.All(
                args.Filters.Box(),
                args.Tags.Box()
            ).Apply(a => {
                    var args = new GetTransitGatewayRouteTablesArgs();
                    a[0].Set(args, nameof(args.Filters));
                    a[1].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetTransitGatewayRouteTablesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetTransitGatewayRouteTablesFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetTransitGatewayRouteTablesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetTransitGatewayRouteTablesFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match
        /// a pair on the desired transit gateway route table.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetTransitGatewayRouteTablesArgs()
        {
        }
    }

    public sealed class GetTransitGatewayRouteTablesApplyArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetTransitGatewayRouteTablesFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public InputList<Inputs.GetTransitGatewayRouteTablesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetTransitGatewayRouteTablesFilterArgs>());
            set => _filters = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match
        /// a pair on the desired transit gateway route table.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetTransitGatewayRouteTablesApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransitGatewayRouteTablesResult
    {
        public readonly ImmutableArray<Outputs.GetTransitGatewayRouteTablesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of Transit Gateway Route Table identifiers.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetTransitGatewayRouteTablesResult(
            ImmutableArray<Outputs.GetTransitGatewayRouteTablesFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableDictionary<string, string> tags)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Tags = tags;
        }
    }
}
