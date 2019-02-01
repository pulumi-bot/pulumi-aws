// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a load balancer policy to an ELB backend server.
type LoadBalancerBackendServerPolicy struct {
	s *pulumi.ResourceState
}

// NewLoadBalancerBackendServerPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerBackendServerPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendServerPolicyArgs, opts ...pulumi.ResourceOpt) (*LoadBalancerBackendServerPolicy, error) {
	if args == nil || args.InstancePort == nil {
		return nil, errors.New("missing required argument 'InstancePort'")
	}
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["instancePort"] = nil
		inputs["loadBalancerName"] = nil
		inputs["policyNames"] = nil
	} else {
		inputs["instancePort"] = args.InstancePort
		inputs["loadBalancerName"] = args.LoadBalancerName
		inputs["policyNames"] = args.PolicyNames
	}
	s, err := ctx.RegisterResource("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerBackendServerPolicy{s: s}, nil
}

// GetLoadBalancerBackendServerPolicy gets an existing LoadBalancerBackendServerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerBackendServerPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LoadBalancerBackendServerPolicyState, opts ...pulumi.ResourceOpt) (*LoadBalancerBackendServerPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["instancePort"] = state.InstancePort
		inputs["loadBalancerName"] = state.LoadBalancerName
		inputs["policyNames"] = state.PolicyNames
	}
	s, err := ctx.ReadResource("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerBackendServerPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LoadBalancerBackendServerPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LoadBalancerBackendServerPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The instance port to apply the policy to.
func (r *LoadBalancerBackendServerPolicy) InstancePort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["instancePort"])
}

// The load balancer to attach the policy to.
func (r *LoadBalancerBackendServerPolicy) LoadBalancerName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["loadBalancerName"])
}

// List of Policy Names to apply to the backend server.
func (r *LoadBalancerBackendServerPolicy) PolicyNames() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policyNames"])
}

// Input properties used for looking up and filtering LoadBalancerBackendServerPolicy resources.
type LoadBalancerBackendServerPolicyState struct {
	// The instance port to apply the policy to.
	InstancePort interface{}
	// The load balancer to attach the policy to.
	LoadBalancerName interface{}
	// List of Policy Names to apply to the backend server.
	PolicyNames interface{}
}

// The set of arguments for constructing a LoadBalancerBackendServerPolicy resource.
type LoadBalancerBackendServerPolicyArgs struct {
	// The instance port to apply the policy to.
	InstancePort interface{}
	// The load balancer to attach the policy to.
	LoadBalancerName interface{}
	// List of Policy Names to apply to the backend server.
	PolicyNames interface{}
}
