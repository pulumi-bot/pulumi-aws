// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    public static class GetScript
    {
        /// <summary>
        /// Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetScriptResult> InvokeAsync(GetScriptArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScriptResult>("aws:glue/getScript:getScript", args ?? new GetScriptArgs(), options.WithVersion());

        public static Output<GetScriptResult> Apply(GetScriptApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.DagEdges.Box(),
                args.DagNodes.Box(),
                args.Language.Box()
            ).Apply(a => {
                    var args = new GetScriptArgs();
                    a[0].Set(args, nameof(args.DagEdges));
                    a[1].Set(args, nameof(args.DagNodes));
                    a[2].Set(args, nameof(args.Language));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetScriptArgs : Pulumi.InvokeArgs
    {
        [Input("dagEdges", required: true)]
        private List<Inputs.GetScriptDagEdgeArgs>? _dagEdges;

        /// <summary>
        /// A list of the edges in the DAG. Defined below.
        /// </summary>
        public List<Inputs.GetScriptDagEdgeArgs> DagEdges
        {
            get => _dagEdges ?? (_dagEdges = new List<Inputs.GetScriptDagEdgeArgs>());
            set => _dagEdges = value;
        }

        [Input("dagNodes", required: true)]
        private List<Inputs.GetScriptDagNodeArgs>? _dagNodes;

        /// <summary>
        /// A list of the nodes in the DAG. Defined below.
        /// </summary>
        public List<Inputs.GetScriptDagNodeArgs> DagNodes
        {
            get => _dagNodes ?? (_dagNodes = new List<Inputs.GetScriptDagNodeArgs>());
            set => _dagNodes = value;
        }

        /// <summary>
        /// The programming language of the resulting code from the DAG. Defaults to `PYTHON`. Valid values are `PYTHON` and `SCALA`.
        /// </summary>
        [Input("language")]
        public string? Language { get; set; }

        public GetScriptArgs()
        {
        }
    }

    public sealed class GetScriptApplyArgs
    {
        [Input("dagEdges", required: true)]
        private InputList<Inputs.GetScriptDagEdgeArgs>? _dagEdges;

        /// <summary>
        /// A list of the edges in the DAG. Defined below.
        /// </summary>
        public InputList<Inputs.GetScriptDagEdgeArgs> DagEdges
        {
            get => _dagEdges ?? (_dagEdges = new InputList<Inputs.GetScriptDagEdgeArgs>());
            set => _dagEdges = value;
        }

        [Input("dagNodes", required: true)]
        private InputList<Inputs.GetScriptDagNodeArgs>? _dagNodes;

        /// <summary>
        /// A list of the nodes in the DAG. Defined below.
        /// </summary>
        public InputList<Inputs.GetScriptDagNodeArgs> DagNodes
        {
            get => _dagNodes ?? (_dagNodes = new InputList<Inputs.GetScriptDagNodeArgs>());
            set => _dagNodes = value;
        }

        /// <summary>
        /// The programming language of the resulting code from the DAG. Defaults to `PYTHON`. Valid values are `PYTHON` and `SCALA`.
        /// </summary>
        [Input("language")]
        public Input<string>? Language { get; set; }

        public GetScriptApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetScriptResult
    {
        public readonly ImmutableArray<Outputs.GetScriptDagEdgeResult> DagEdges;
        public readonly ImmutableArray<Outputs.GetScriptDagNodeResult> DagNodes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Language;
        /// <summary>
        /// The Python script generated from the DAG when the `language` argument is set to `PYTHON`.
        /// </summary>
        public readonly string PythonScript;
        /// <summary>
        /// The Scala code generated from the DAG when the `language` argument is set to `SCALA`.
        /// </summary>
        public readonly string ScalaCode;

        [OutputConstructor]
        private GetScriptResult(
            ImmutableArray<Outputs.GetScriptDagEdgeResult> dagEdges,

            ImmutableArray<Outputs.GetScriptDagNodeResult> dagNodes,

            string id,

            string? language,

            string pythonScript,

            string scalaCode)
        {
            DagEdges = dagEdges;
            DagNodes = dagNodes;
            Id = id;
            Language = language;
            PythonScript = pythonScript;
            ScalaCode = scalaCode;
        }
    }
}
