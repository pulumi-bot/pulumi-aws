// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Outputs
{

    [OutputType]
    public sealed class CatalogTableStorageDescriptorSortColumn
    {
        public readonly string Column;
        public readonly int SortOrder;

        [OutputConstructor]
        private CatalogTableStorageDescriptorSortColumn(
            string column,

            int sortOrder)
        {
            Column = column;
            SortOrder = sortOrder;
        }
    }
}
