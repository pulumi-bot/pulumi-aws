// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway Method Settings, e.g. logging or monitoring.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_method_settings.html.markdown.
type MethodSettings struct {
	s *pulumi.ResourceState
}

// NewMethodSettings registers a new resource with the given unique name, arguments, and options.
func NewMethodSettings(ctx *pulumi.Context,
	name string, args *MethodSettingsArgs, opts ...pulumi.ResourceOpt) (*MethodSettings, error) {
	if args == nil || args.MethodPath == nil {
		return nil, errors.New("missing required argument 'MethodPath'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	if args == nil || args.StageName == nil {
		return nil, errors.New("missing required argument 'StageName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["methodPath"] = nil
		inputs["restApi"] = nil
		inputs["settings"] = nil
		inputs["stageName"] = nil
	} else {
		inputs["methodPath"] = args.MethodPath
		inputs["restApi"] = args.RestApi
		inputs["settings"] = args.Settings
		inputs["stageName"] = args.StageName
	}
	s, err := ctx.RegisterResource("aws:apigateway/methodSettings:MethodSettings", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MethodSettings{s: s}, nil
}

// GetMethodSettings gets an existing MethodSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethodSettings(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MethodSettingsState, opts ...pulumi.ResourceOpt) (*MethodSettings, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["methodPath"] = state.MethodPath
		inputs["restApi"] = state.RestApi
		inputs["settings"] = state.Settings
		inputs["stageName"] = state.StageName
	}
	s, err := ctx.ReadResource("aws:apigateway/methodSettings:MethodSettings", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MethodSettings{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MethodSettings) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MethodSettings) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
func (r *MethodSettings) MethodPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["methodPath"])
}

// The ID of the REST API
func (r *MethodSettings) RestApi() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["restApi"])
}

// The settings block, see below.
func (r *MethodSettings) Settings() *pulumi.Output {
	return r.s.State["settings"]
}

// The name of the stage
func (r *MethodSettings) StageName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["stageName"])
}

// Input properties used for looking up and filtering MethodSettings resources.
type MethodSettingsState struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath interface{}
	// The ID of the REST API
	RestApi interface{}
	// The settings block, see below.
	Settings interface{}
	// The name of the stage
	StageName interface{}
}

// The set of arguments for constructing a MethodSettings resource.
type MethodSettingsArgs struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath interface{}
	// The ID of the REST API
	RestApi interface{}
	// The settings block, see below.
	Settings interface{}
	// The name of the stage
	StageName interface{}
}
