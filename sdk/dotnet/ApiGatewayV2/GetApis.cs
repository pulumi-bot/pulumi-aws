// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    public static class GetApis
    {
        /// <summary>
        /// Provides details about multiple Amazon API Gateway Version 2 APIs.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.ApiGatewayV2.GetApis.InvokeAsync(new Aws.ApiGatewayV2.GetApisArgs
        ///         {
        ///             ProtocolType = "HTTP",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApisResult> InvokeAsync(GetApisArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApisResult>("aws:apigatewayv2/getApis:getApis", args ?? new GetApisArgs(), options.WithVersion());

        public static Output<GetApisResult> Apply(GetApisApplyArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetApisApplyArgs();
            return Pulumi.Output.All(
                args.Name.Box(),
                args.ProtocolType.Box(),
                args.Tags.Box()
            ).Apply(a => {
                    var args = new GetApisArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.ProtocolType));
                    a[2].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetApisArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The API name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The API protocol.
        /// </summary>
        [Input("protocolType")]
        public string? ProtocolType { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired APIs.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetApisArgs()
        {
        }
    }

    public sealed class GetApisApplyArgs
    {
        /// <summary>
        /// The API name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The API protocol.
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired APIs.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetApisApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApisResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of API identifiers.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly string? Name;
        public readonly string? ProtocolType;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetApisResult(
            string id,

            ImmutableArray<string> ids,

            string? name,

            string? protocolType,

            ImmutableDictionary<string, string>? tags)
        {
            Id = id;
            Ids = ids;
            Name = name;
            ProtocolType = protocolType;
            Tags = tags;
        }
    }
}
