// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Budgets.Outputs
{

    [OutputType]
    public sealed class BudgetCostTypes
    {
        public readonly bool? IncludeCredit;
        public readonly bool? IncludeDiscount;
        public readonly bool? IncludeOtherSubscription;
        public readonly bool? IncludeRecurring;
        public readonly bool? IncludeRefund;
        public readonly bool? IncludeSubscription;
        public readonly bool? IncludeSupport;
        public readonly bool? IncludeTax;
        public readonly bool? IncludeUpfront;
        public readonly bool? UseAmortized;
        public readonly bool? UseBlended;

        [OutputConstructor]
        private BudgetCostTypes(
            bool? includeCredit,

            bool? includeDiscount,

            bool? includeOtherSubscription,

            bool? includeRecurring,

            bool? includeRefund,

            bool? includeSubscription,

            bool? includeSupport,

            bool? includeTax,

            bool? includeUpfront,

            bool? useAmortized,

            bool? useBlended)
        {
            IncludeCredit = includeCredit;
            IncludeDiscount = includeDiscount;
            IncludeOtherSubscription = includeOtherSubscription;
            IncludeRecurring = includeRecurring;
            IncludeRefund = includeRefund;
            IncludeSubscription = includeSubscription;
            IncludeSupport = includeSupport;
            IncludeTax = includeTax;
            IncludeUpfront = includeUpfront;
            UseAmortized = useAmortized;
            UseBlended = useBlended;
        }
    }
}
