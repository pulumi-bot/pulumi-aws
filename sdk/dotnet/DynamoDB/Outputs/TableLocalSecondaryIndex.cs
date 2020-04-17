// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DynamoDB.Outputs
{

    [OutputType]
    public sealed class TableLocalSecondaryIndex
    {
        /// <summary>
        /// The name of the table, this needs to be unique
        /// within a region.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Only required with `INCLUDE` as a
        /// projection type; a list of attributes to project into the index. These
        /// do not need to be defined as attributes on the table.
        /// </summary>
        public readonly ImmutableArray<string> NonKeyAttributes;
        /// <summary>
        /// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
        /// where `ALL` projects every attribute into the index, `KEYS_ONLY`
        /// projects just the hash and range key into the index, and `INCLUDE`
        /// projects only the keys specified in the _non_key_attributes_
        /// parameter.
        /// </summary>
        public readonly string ProjectionType;
        /// <summary>
        /// The attribute to use as the range (sort) key. Must also be defined as an `attribute`, see below.
        /// </summary>
        public readonly string RangeKey;

        [OutputConstructor]
        private TableLocalSecondaryIndex(
            string name,

            ImmutableArray<string> nonKeyAttributes,

            string projectionType,

            string rangeKey)
        {
            Name = name;
            NonKeyAttributes = nonKeyAttributes;
            ProjectionType = projectionType;
            RangeKey = rangeKey;
        }
    }
}
