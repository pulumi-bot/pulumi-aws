// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh service mesh resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_mesh.html.markdown.
type Mesh struct {
	s *pulumi.ResourceState
}

// NewMesh registers a new resource with the given unique name, arguments, and options.
func NewMesh(ctx *pulumi.Context,
	name string, args *MeshArgs, opts ...pulumi.ResourceOpt) (*Mesh, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["spec"] = nil
		inputs["tags"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["spec"] = args.Spec
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["createdDate"] = nil
	inputs["lastUpdatedDate"] = nil
	s, err := ctx.RegisterResource("aws:appmesh/mesh:Mesh", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Mesh{s: s}, nil
}

// GetMesh gets an existing Mesh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMesh(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MeshState, opts ...pulumi.ResourceOpt) (*Mesh, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["createdDate"] = state.CreatedDate
		inputs["lastUpdatedDate"] = state.LastUpdatedDate
		inputs["name"] = state.Name
		inputs["spec"] = state.Spec
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:appmesh/mesh:Mesh", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Mesh{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Mesh) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Mesh) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the service mesh.
func (r *Mesh) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The creation date of the service mesh.
func (r *Mesh) CreatedDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createdDate"])
}

// The last update date of the service mesh.
func (r *Mesh) LastUpdatedDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastUpdatedDate"])
}

// The name to use for the service mesh.
func (r *Mesh) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The service mesh specification to apply.
func (r *Mesh) Spec() *pulumi.Output {
	return r.s.State["spec"]
}

// A mapping of tags to assign to the resource.
func (r *Mesh) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Mesh resources.
type MeshState struct {
	// The ARN of the service mesh.
	Arn interface{}
	// The creation date of the service mesh.
	CreatedDate interface{}
	// The last update date of the service mesh.
	LastUpdatedDate interface{}
	// The name to use for the service mesh.
	Name interface{}
	// The service mesh specification to apply.
	Spec interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Mesh resource.
type MeshArgs struct {
	// The name to use for the service mesh.
	Name interface{}
	// The service mesh specification to apply.
	Spec interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
