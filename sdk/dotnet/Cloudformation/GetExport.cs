// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFormation
{
    public static partial class Invokes
    {
        /// <summary>
        /// The CloudFormation Export data source allows access to stack
        /// exports specified in the [Output](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) section of the Cloudformation Template using the optional Export Property.
        /// 
        ///  &gt; Note: If you are trying to use a value from a Cloudformation Stack in the same deployment please use normal interpolation or Cloudformation Outputs. 
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudformation_export.html.markdown.
        /// </summary>
        [Obsolete("Use GetExport.InvokeAsync() instead")]
        public static Task<GetExportResult> GetExport(GetExportArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExportResult>("aws:cloudformation/getExport:getExport", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetExport
    {
        /// <summary>
        /// The CloudFormation Export data source allows access to stack
        /// exports specified in the [Output](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) section of the Cloudformation Template using the optional Export Property.
        /// 
        ///  &gt; Note: If you are trying to use a value from a Cloudformation Stack in the same deployment please use normal interpolation or Cloudformation Outputs. 
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudformation_export.html.markdown.
        /// </summary>
        public static Task<GetExportResult> InvokeAsync(GetExportArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExportResult>("aws:cloudformation/getExport:getExport", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetExportArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the export as it appears in the console or from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetExportArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetExportResult
    {
        /// <summary>
        /// The exporting_stack_id (AWS ARNs) equivalent `ExportingStackId` from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html) 
        /// </summary>
        public readonly string ExportingStackId;
        public readonly string Name;
        /// <summary>
        /// The value from Cloudformation export identified by the export name found from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetExportResult(
            string exportingStackId,
            string name,
            string value,
            string id)
        {
            ExportingStackId = exportingStackId;
            Name = name;
            Value = value;
            Id = id;
        }
    }
}
