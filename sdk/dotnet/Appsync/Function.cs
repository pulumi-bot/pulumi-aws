// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync
{
    /// <summary>
    /// Provides an AppSync Function.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_function.html.markdown.
    /// </summary>
    public partial class Function : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the associated AppSync API.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Function object.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Function DataSource name.
        /// </summary>
        [Output("dataSource")]
        public Output<string> DataSource { get; private set; } = null!;

        /// <summary>
        /// The Function description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A unique ID representing the Function object.
        /// </summary>
        [Output("functionId")]
        public Output<string> FunctionId { get; private set; } = null!;

        /// <summary>
        /// The version of the request mapping template. Currently the supported value is `2018-05-29`.
        /// </summary>
        [Output("functionVersion")]
        public Output<string?> FunctionVersion { get; private set; } = null!;

        /// <summary>
        /// The Function name. The function name does not have to be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        /// </summary>
        [Output("requestMappingTemplate")]
        public Output<string> RequestMappingTemplate { get; private set; } = null!;

        /// <summary>
        /// The Function response mapping template.
        /// </summary>
        [Output("responseMappingTemplate")]
        public Output<string> ResponseMappingTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("aws:appsync/function:Function", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("aws:appsync/function:Function", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the associated AppSync API.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The Function DataSource name.
        /// </summary>
        [Input("dataSource", required: true)]
        public Input<string> DataSource { get; set; } = null!;

        /// <summary>
        /// The Function description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the request mapping template. Currently the supported value is `2018-05-29`.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? FunctionVersion { get; set; }

        /// <summary>
        /// The Function name. The function name does not have to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        /// </summary>
        [Input("requestMappingTemplate", required: true)]
        public Input<string> RequestMappingTemplate { get; set; } = null!;

        /// <summary>
        /// The Function response mapping template.
        /// </summary>
        [Input("responseMappingTemplate", required: true)]
        public Input<string> ResponseMappingTemplate { get; set; } = null!;

        public FunctionArgs()
        {
        }
    }

    public sealed class FunctionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the associated AppSync API.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The ARN of the Function object.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Function DataSource name.
        /// </summary>
        [Input("dataSource")]
        public Input<string>? DataSource { get; set; }

        /// <summary>
        /// The Function description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique ID representing the Function object.
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        /// <summary>
        /// The version of the request mapping template. Currently the supported value is `2018-05-29`.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? FunctionVersion { get; set; }

        /// <summary>
        /// The Function name. The function name does not have to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        /// </summary>
        [Input("requestMappingTemplate")]
        public Input<string>? RequestMappingTemplate { get; set; }

        /// <summary>
        /// The Function response mapping template.
        /// </summary>
        [Input("responseMappingTemplate")]
        public Input<string>? ResponseMappingTemplate { get; set; }

        public FunctionState()
        {
        }
    }
}
