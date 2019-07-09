// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh virtual service resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_service.html.markdown.
type VirtualService struct {
	s *pulumi.ResourceState
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOpt) (*VirtualService, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["meshName"] = nil
		inputs["name"] = nil
		inputs["spec"] = nil
	} else {
		inputs["meshName"] = args.MeshName
		inputs["name"] = args.Name
		inputs["spec"] = args.Spec
	}
	inputs["arn"] = nil
	inputs["createdDate"] = nil
	inputs["lastUpdatedDate"] = nil
	s, err := ctx.RegisterResource("aws:appmesh/virtualService:VirtualService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualService{s: s}, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualServiceState, opts ...pulumi.ResourceOpt) (*VirtualService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["createdDate"] = state.CreatedDate
		inputs["lastUpdatedDate"] = state.LastUpdatedDate
		inputs["meshName"] = state.MeshName
		inputs["name"] = state.Name
		inputs["spec"] = state.Spec
	}
	s, err := ctx.ReadResource("aws:appmesh/virtualService:VirtualService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualService) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualService) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the virtual service.
func (r *VirtualService) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The creation date of the virtual service.
func (r *VirtualService) CreatedDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createdDate"])
}

// The last update date of the virtual service.
func (r *VirtualService) LastUpdatedDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastUpdatedDate"])
}

// The name of the service mesh in which to create the virtual service.
func (r *VirtualService) MeshName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["meshName"])
}

// The name to use for the virtual service.
func (r *VirtualService) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The virtual service specification to apply.
func (r *VirtualService) Spec() *pulumi.Output {
	return r.s.State["spec"]
}

// Input properties used for looking up and filtering VirtualService resources.
type VirtualServiceState struct {
	// The ARN of the virtual service.
	Arn interface{}
	// The creation date of the virtual service.
	CreatedDate interface{}
	// The last update date of the virtual service.
	LastUpdatedDate interface{}
	// The name of the service mesh in which to create the virtual service.
	MeshName interface{}
	// The name to use for the virtual service.
	Name interface{}
	// The virtual service specification to apply.
	Spec interface{}
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service.
	MeshName interface{}
	// The name to use for the virtual service.
	Name interface{}
	// The virtual service specification to apply.
	Spec interface{}
}
