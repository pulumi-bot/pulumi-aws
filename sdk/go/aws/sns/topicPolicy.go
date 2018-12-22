// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SNS topic policy resource
// 
// > **NOTE:** If a Principal is specified as just an AWS account ID rather than an ARN, AWS silently converts it to the ARN for the root user, causing future terraform plans to differ. To avoid this problem, just specify the full ARN, e.g. `arn:aws:iam::123456789012:root`
type TopicPolicy struct {
	s *pulumi.ResourceState
}

// NewTopicPolicy registers a new resource with the given unique name, arguments, and options.
func NewTopicPolicy(ctx *pulumi.Context,
	name string, args *TopicPolicyArgs, opts ...pulumi.ResourceOpt) (*TopicPolicy, error) {
	if args == nil || args.Arn == nil {
		return nil, errors.New("missing required argument 'Arn'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["arn"] = nil
		inputs["policy"] = nil
	} else {
		inputs["arn"] = args.Arn
		inputs["policy"] = args.Policy
	}
	s, err := ctx.RegisterResource("aws:sns/topicPolicy:TopicPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TopicPolicy{s: s}, nil
}

// GetTopicPolicy gets an existing TopicPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TopicPolicyState, opts ...pulumi.ResourceOpt) (*TopicPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["policy"] = state.Policy
	}
	s, err := ctx.ReadResource("aws:sns/topicPolicy:TopicPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TopicPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TopicPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TopicPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the SNS topic
func (r *TopicPolicy) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The fully-formed AWS policy as JSON. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
func (r *TopicPolicy) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// Input properties used for looking up and filtering TopicPolicy resources.
type TopicPolicyState struct {
	// The ARN of the SNS topic
	Arn interface{}
	// The fully-formed AWS policy as JSON. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
	Policy interface{}
}

// The set of arguments for constructing a TopicPolicy resource.
type TopicPolicyArgs struct {
	// The ARN of the SNS topic
	Arn interface{}
	// The fully-formed AWS policy as JSON. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
	Policy interface{}
}
