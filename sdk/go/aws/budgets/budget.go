// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package budgets

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a budgets budget resource. Budgets use the cost visualisation provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.
type Budget struct {
	s *pulumi.ResourceState
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOpt) (*Budget, error) {
	if args == nil || args.BudgetType == nil {
		return nil, errors.New("missing required argument 'BudgetType'")
	}
	if args == nil || args.LimitAmount == nil {
		return nil, errors.New("missing required argument 'LimitAmount'")
	}
	if args == nil || args.LimitUnit == nil {
		return nil, errors.New("missing required argument 'LimitUnit'")
	}
	if args == nil || args.TimePeriodStart == nil {
		return nil, errors.New("missing required argument 'TimePeriodStart'")
	}
	if args == nil || args.TimeUnit == nil {
		return nil, errors.New("missing required argument 'TimeUnit'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountId"] = nil
		inputs["budgetType"] = nil
		inputs["costFilters"] = nil
		inputs["costTypes"] = nil
		inputs["limitAmount"] = nil
		inputs["limitUnit"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["timePeriodEnd"] = nil
		inputs["timePeriodStart"] = nil
		inputs["timeUnit"] = nil
	} else {
		inputs["accountId"] = args.AccountId
		inputs["budgetType"] = args.BudgetType
		inputs["costFilters"] = args.CostFilters
		inputs["costTypes"] = args.CostTypes
		inputs["limitAmount"] = args.LimitAmount
		inputs["limitUnit"] = args.LimitUnit
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["timePeriodEnd"] = args.TimePeriodEnd
		inputs["timePeriodStart"] = args.TimePeriodStart
		inputs["timeUnit"] = args.TimeUnit
	}
	s, err := ctx.RegisterResource("aws:budgets/budget:Budget", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Budget{s: s}, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BudgetState, opts ...pulumi.ResourceOpt) (*Budget, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountId"] = state.AccountId
		inputs["budgetType"] = state.BudgetType
		inputs["costFilters"] = state.CostFilters
		inputs["costTypes"] = state.CostTypes
		inputs["limitAmount"] = state.LimitAmount
		inputs["limitUnit"] = state.LimitUnit
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["timePeriodEnd"] = state.TimePeriodEnd
		inputs["timePeriodStart"] = state.TimePeriodStart
		inputs["timeUnit"] = state.TimeUnit
	}
	s, err := ctx.ReadResource("aws:budgets/budget:Budget", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Budget{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Budget) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Budget) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the target account for budget. Will use current user's account_id by default if omitted.
func (r *Budget) AccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accountId"])
}

// Whether this budget tracks monetary cost or usage.
func (r *Budget) BudgetType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["budgetType"])
}

// Map of CostFilters key/value pairs to apply to the budget.
func (r *Budget) CostFilters() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["costFilters"])
}

// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
func (r *Budget) CostTypes() *pulumi.Output {
	return r.s.State["costTypes"]
}

// The amount of cost or usage being measured for a budget.
func (r *Budget) LimitAmount() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["limitAmount"])
}

// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
func (r *Budget) LimitUnit() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["limitUnit"])
}

// The name of a budget. Unique within accounts.
func (r *Budget) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The prefix of the name of a budget. Unique within accounts.
func (r *Budget) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
func (r *Budget) TimePeriodEnd() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["timePeriodEnd"])
}

// The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
func (r *Budget) TimePeriodStart() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["timePeriodStart"])
}

// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
func (r *Budget) TimeUnit() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["timeUnit"])
}

// Input properties used for looking up and filtering Budget resources.
type BudgetState struct {
	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountId interface{}
	// Whether this budget tracks monetary cost or usage.
	BudgetType interface{}
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters interface{}
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
	CostTypes interface{}
	// The amount of cost or usage being measured for a budget.
	LimitAmount interface{}
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit interface{}
	// The name of a budget. Unique within accounts.
	Name interface{}
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix interface{}
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd interface{}
	// The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart interface{}
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
	TimeUnit interface{}
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountId interface{}
	// Whether this budget tracks monetary cost or usage.
	BudgetType interface{}
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters interface{}
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
	CostTypes interface{}
	// The amount of cost or usage being measured for a budget.
	LimitAmount interface{}
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit interface{}
	// The name of a budget. Unique within accounts.
	Name interface{}
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix interface{}
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd interface{}
	// The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart interface{}
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
	TimeUnit interface{}
}
