// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Glue Classifier resource.
// 
// > **NOTE:** It is only valid to create one type of classifier (grok, JSON, or XML). Changing classifier types will recreate the classifier.
type Classifier struct {
	s *pulumi.ResourceState
}

// NewClassifier registers a new resource with the given unique name, arguments, and options.
func NewClassifier(ctx *pulumi.Context,
	name string, args *ClassifierArgs, opts ...pulumi.ResourceOpt) (*Classifier, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["grokClassifier"] = nil
		inputs["jsonClassifier"] = nil
		inputs["name"] = nil
		inputs["xmlClassifier"] = nil
	} else {
		inputs["grokClassifier"] = args.GrokClassifier
		inputs["jsonClassifier"] = args.JsonClassifier
		inputs["name"] = args.Name
		inputs["xmlClassifier"] = args.XmlClassifier
	}
	s, err := ctx.RegisterResource("aws:glue/classifier:Classifier", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Classifier{s: s}, nil
}

// GetClassifier gets an existing Classifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClassifier(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClassifierState, opts ...pulumi.ResourceOpt) (*Classifier, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["grokClassifier"] = state.GrokClassifier
		inputs["jsonClassifier"] = state.JsonClassifier
		inputs["name"] = state.Name
		inputs["xmlClassifier"] = state.XmlClassifier
	}
	s, err := ctx.ReadResource("aws:glue/classifier:Classifier", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Classifier{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Classifier) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Classifier) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A classifier that uses grok patterns. Defined below.
func (r *Classifier) GrokClassifier() *pulumi.Output {
	return r.s.State["grokClassifier"]
}

// A classifier for JSON content. Defined below.
func (r *Classifier) JsonClassifier() *pulumi.Output {
	return r.s.State["jsonClassifier"]
}

// The name of the classifier.
func (r *Classifier) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A classifier for XML content. Defined below.
func (r *Classifier) XmlClassifier() *pulumi.Output {
	return r.s.State["xmlClassifier"]
}

// Input properties used for looking up and filtering Classifier resources.
type ClassifierState struct {
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier interface{}
	// A classifier for JSON content. Defined below.
	JsonClassifier interface{}
	// The name of the classifier.
	Name interface{}
	// A classifier for XML content. Defined below.
	XmlClassifier interface{}
}

// The set of arguments for constructing a Classifier resource.
type ClassifierArgs struct {
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier interface{}
	// A classifier for JSON content. Defined below.
	JsonClassifier interface{}
	// The name of the classifier.
	Name interface{}
	// A classifier for XML content. Defined below.
	XmlClassifier interface{}
}
