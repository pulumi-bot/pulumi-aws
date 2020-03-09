// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB
{
    /// <summary>
    /// Provides a DynamoDB table item resource
    /// 
    /// &gt; **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
    ///   You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dynamodb_table_item.html.markdown.
    /// </summary>
    public partial class TableItem : Pulumi.CustomResource
    {
        /// <summary>
        /// Hash key to use for lookups and identification of the item
        /// </summary>
        [Output("hashKey")]
        public Output<string> HashKey { get; private set; } = null!;

        /// <summary>
        /// JSON representation of a map of attribute name/value pairs, one for each attribute.
        /// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        /// </summary>
        [Output("item")]
        public Output<string> Item { get; private set; } = null!;

        /// <summary>
        /// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        /// </summary>
        [Output("rangeKey")]
        public Output<string?> RangeKey { get; private set; } = null!;

        /// <summary>
        /// The name of the table to contain the item.
        /// </summary>
        [Output("tableName")]
        public Output<string> TableName { get; private set; } = null!;


        /// <summary>
        /// Create a TableItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TableItem(string name, TableItemArgs args, CustomResourceOptions? options = null)
            : base("aws:dynamodb/tableItem:TableItem", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private TableItem(string name, Input<string> id, TableItemState? state = null, CustomResourceOptions? options = null)
            : base("aws:dynamodb/tableItem:TableItem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TableItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TableItem Get(string name, Input<string> id, TableItemState? state = null, CustomResourceOptions? options = null)
        {
            return new TableItem(name, id, state, options);
        }
    }

    public sealed class TableItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hash key to use for lookups and identification of the item
        /// </summary>
        [Input("hashKey", required: true)]
        public Input<string> HashKey { get; set; } = null!;

        /// <summary>
        /// JSON representation of a map of attribute name/value pairs, one for each attribute.
        /// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        /// </summary>
        [Input("item", required: true)]
        public Input<string> Item { get; set; } = null!;

        /// <summary>
        /// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        /// </summary>
        [Input("rangeKey")]
        public Input<string>? RangeKey { get; set; }

        /// <summary>
        /// The name of the table to contain the item.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public TableItemArgs()
        {
        }
    }

    public sealed class TableItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hash key to use for lookups and identification of the item
        /// </summary>
        [Input("hashKey")]
        public Input<string>? HashKey { get; set; }

        /// <summary>
        /// JSON representation of a map of attribute name/value pairs, one for each attribute.
        /// Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        /// </summary>
        [Input("item")]
        public Input<string>? Item { get; set; }

        /// <summary>
        /// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        /// </summary>
        [Input("rangeKey")]
        public Input<string>? RangeKey { get; set; }

        /// <summary>
        /// The name of the table to contain the item.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public TableItemState()
        {
        }
    }
}
