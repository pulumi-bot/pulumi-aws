// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/glue_script.html.markdown.
        /// </summary>
        [Obsolete("Use GetScript.InvokeAsync() instead")]
        public static Task<GetScriptResult> GetScript(GetScriptArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScriptResult>("aws:glue/getScript:getScript", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetScript
    {
        /// <summary>
        /// Use this data source to generate a Glue script from a Directed Acyclic Graph (DAG).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/glue_script.html.markdown.
        /// </summary>
        public static Task<GetScriptResult> InvokeAsync(GetScriptArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetScriptResult>("aws:glue/getScript:getScript", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetScriptArgs : Pulumi.InvokeArgs
    {
        [Input("dagEdges", required: true)]
        private List<Inputs.GetScriptDagEdgesArgs>? _dagEdges;

        /// <summary>
        /// A list of the edges in the DAG. Defined below.
        /// </summary>
        public List<Inputs.GetScriptDagEdgesArgs> DagEdges
        {
            get => _dagEdges ?? (_dagEdges = new List<Inputs.GetScriptDagEdgesArgs>());
            set => _dagEdges = value;
        }

        [Input("dagNodes", required: true)]
        private List<Inputs.GetScriptDagNodesArgs>? _dagNodes;

        /// <summary>
        /// A list of the nodes in the DAG. Defined below.
        /// </summary>
        public List<Inputs.GetScriptDagNodesArgs> DagNodes
        {
            get => _dagNodes ?? (_dagNodes = new List<Inputs.GetScriptDagNodesArgs>());
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

    [OutputType]
    public sealed class GetScriptResult
    {
        public readonly ImmutableArray<Outputs.GetScriptDagEdgesResult> DagEdges;
        public readonly ImmutableArray<Outputs.GetScriptDagNodesResult> DagNodes;
        public readonly string? Language;
        /// <summary>
        /// The Python script generated from the DAG when the `language` argument is set to `PYTHON`.
        /// </summary>
        public readonly string PythonScript;
        /// <summary>
        /// The Scala code generated from the DAG when the `language` argument is set to `SCALA`.
        /// </summary>
        public readonly string ScalaCode;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetScriptResult(
            ImmutableArray<Outputs.GetScriptDagEdgesResult> dagEdges,
            ImmutableArray<Outputs.GetScriptDagNodesResult> dagNodes,
            string? language,
            string pythonScript,
            string scalaCode,
            string id)
        {
            DagEdges = dagEdges;
            DagNodes = dagNodes;
            Language = language;
            PythonScript = pythonScript;
            ScalaCode = scalaCode;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetScriptDagEdgesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the node at which the edge starts.
        /// </summary>
        [Input("source", required: true)]
        public string Source { get; set; } = null!;

        /// <summary>
        /// The ID of the node at which the edge ends.
        /// </summary>
        [Input("target", required: true)]
        public string Target { get; set; } = null!;

        /// <summary>
        /// The target of the edge.
        /// </summary>
        [Input("targetParameter")]
        public string? TargetParameter { get; set; }

        public GetScriptDagEdgesArgs()
        {
        }
    }

    public sealed class GetScriptDagNodesArgs : Pulumi.InvokeArgs
    {
        [Input("args", required: true)]
        private List<GetScriptDagNodesArgsArgs>? _args;

        /// <summary>
        /// Nested configuration an argument or property of a node. Defined below.
        /// </summary>
        public List<GetScriptDagNodesArgsArgs> Args
        {
            get => _args ?? (_args = new List<GetScriptDagNodesArgsArgs>());
            set => _args = value;
        }

        /// <summary>
        /// A node identifier that is unique within the node's graph.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The line number of the node.
        /// </summary>
        [Input("lineNumber")]
        public int? LineNumber { get; set; }

        /// <summary>
        /// The type of node this is.
        /// </summary>
        [Input("nodeType", required: true)]
        public string NodeType { get; set; } = null!;

        public GetScriptDagNodesArgs()
        {
        }
    }

    public sealed class GetScriptDagNodesArgsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the argument or property.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Boolean if the value is used as a parameter. Defaults to `false`.
        /// </summary>
        [Input("param")]
        public bool? Param { get; set; }

        /// <summary>
        /// The value of the argument or property.
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetScriptDagNodesArgsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetScriptDagEdgesResult
    {
        /// <summary>
        /// The ID of the node at which the edge starts.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The ID of the node at which the edge ends.
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// The target of the edge.
        /// </summary>
        public readonly string? TargetParameter;

        [OutputConstructor]
        private GetScriptDagEdgesResult(
            string source,
            string target,
            string? targetParameter)
        {
            Source = source;
            Target = target;
            TargetParameter = targetParameter;
        }
    }

    [OutputType]
    public sealed class GetScriptDagNodesArgsResult
    {
        /// <summary>
        /// The name of the argument or property.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Boolean if the value is used as a parameter. Defaults to `false`.
        /// </summary>
        public readonly bool? Param;
        /// <summary>
        /// The value of the argument or property.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetScriptDagNodesArgsResult(
            string name,
            bool? param,
            string value)
        {
            Name = name;
            Param = param;
            Value = value;
        }
    }

    [OutputType]
    public sealed class GetScriptDagNodesResult
    {
        /// <summary>
        /// Nested configuration an argument or property of a node. Defined below.
        /// </summary>
        public readonly ImmutableArray<GetScriptDagNodesArgsResult> Args;
        /// <summary>
        /// A node identifier that is unique within the node's graph.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The line number of the node.
        /// </summary>
        public readonly int? LineNumber;
        /// <summary>
        /// The type of node this is.
        /// </summary>
        public readonly string NodeType;

        [OutputConstructor]
        private GetScriptDagNodesResult(
            ImmutableArray<GetScriptDagNodesArgsResult> args,
            string id,
            int? lineNumber,
            string nodeType)
        {
            Args = args;
            Id = id;
            LineNumber = lineNumber;
            NodeType = nodeType;
        }
    }
    }
}
