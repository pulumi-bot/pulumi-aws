// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class CatalogTableStorageDescriptorArgs : Pulumi.ResourceArgs
    {
        [Input("bucketColumns")]
        private InputList<string>? _bucketColumns;
        public InputList<string> BucketColumns
        {
            get => _bucketColumns ?? (_bucketColumns = new InputList<string>());
            set => _bucketColumns = value;
        }

        [Input("columns")]
        private InputList<Inputs.CatalogTableStorageDescriptorColumnArgs>? _columns;
        public InputList<Inputs.CatalogTableStorageDescriptorColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.CatalogTableStorageDescriptorColumnArgs>());
            set => _columns = value;
        }

        [Input("compressed")]
        public Input<bool>? Compressed { get; set; }

        [Input("inputFormat")]
        public Input<string>? InputFormat { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("numberOfBuckets")]
        public Input<int>? NumberOfBuckets { get; set; }

        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("serDeInfo")]
        public Input<Inputs.CatalogTableStorageDescriptorSerDeInfoArgs>? SerDeInfo { get; set; }

        [Input("skewedInfo")]
        public Input<Inputs.CatalogTableStorageDescriptorSkewedInfoArgs>? SkewedInfo { get; set; }

        [Input("sortColumns")]
        private InputList<Inputs.CatalogTableStorageDescriptorSortColumnArgs>? _sortColumns;
        public InputList<Inputs.CatalogTableStorageDescriptorSortColumnArgs> SortColumns
        {
            get => _sortColumns ?? (_sortColumns = new InputList<Inputs.CatalogTableStorageDescriptorSortColumnArgs>());
            set => _sortColumns = value;
        }

        [Input("storedAsSubDirectories")]
        public Input<bool>? StoredAsSubDirectories { get; set; }

        public CatalogTableStorageDescriptorArgs()
        {
        }
    }
}
