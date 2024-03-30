// @opencdk8s/cdk8s-argocd-resources
package opencdk8scdk8sargocdresources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-argocd-resources-go/opencdk8scdk8sargocdresources/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/opencdk8s/cdk8s-argocd-resources-go/opencdk8scdk8sargocdresources/internal"
	"github.com/opencdk8s/cdk8s-argocd-resources-go/opencdk8scdk8sargocdresources/k8s"
)

// Experimental.
type ApplicationDestination struct {
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	Namespace *string `json:"namespace"`
	// Experimental.
	Server *string `json:"server"`
}

// Experimental.
type ApplicationDirectory struct {
	// Experimental.
	Recurse *bool `json:"recurse"`
}

// Experimental.
type ApplicationPlugin struct {
	// Experimental.
	Env *[]*k8s.EnvVar `json:"env"`
	// Experimental.
	Name *string `json:"name"`
}

// Experimental.
type ApplicationSource struct {
	// Experimental.
	Chart *string `json:"chart"`
	// Experimental.
	Directory *ApplicationDirectory `json:"directory"`
	// Experimental.
	Helm *HelmOptions `json:"helm"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Plugin *ApplicationPlugin `json:"plugin"`
	// Experimental.
	Ref *string `json:"ref"`
	// Experimental.
	RepoURL *string `json:"repoURL"`
	// Experimental.
	TargetRevision *string `json:"targetRevision"`
}

// Experimental.
type ApplicationSyncPolicy struct {
	// Experimental.
	Automated *SyncPolicyAutomated `json:"automated"`
	// Experimental.
	Retry *SyncRetry `json:"retry"`
	// Experimental.
	SyncOptions *[]*string `json:"syncOptions"`
}

// Experimental.
type ArgoCdApplication interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ArgoCdApplication
type jsiiProxy_ArgoCdApplication struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ArgoCdApplication) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdApplication) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdApplication) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdApplication) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdApplication) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewArgoCdApplication(scope constructs.Construct, id *string, props *ArgoCdApplicationProps) ArgoCdApplication {
	_init_.Initialize()

	j := jsiiProxy_ArgoCdApplication{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewArgoCdApplication_Override(a ArgoCdApplication, scope constructs.Construct, id *string, props *ArgoCdApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdApplication",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ArgoCdApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for an ingress object. https://github.com/kubernetes-sigs/aws-load-balancer-controller.
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func ArgoCdApplication_Manifest(props *ArgoCdApplicationProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdApplication",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
// Experimental.
func ArgoCdApplication_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdApplication",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ArgoCdApplication_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdApplication",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
// Experimental.
func (a *jsiiProxy_ArgoCdApplication) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
// Experimental.
func (a *jsiiProxy_ArgoCdApplication) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
// Experimental.
func (a *jsiiProxy_ArgoCdApplication) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ArgoCdApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ArgoCdApplicationProps struct {
	// Experimental.
	Metadata *k8s.ObjectMeta `json:"metadata"`
	// Experimental.
	Spec *ArgoCdApplicationSpec `json:"spec"`
}

// Experimental.
type ArgoCdApplicationSpec struct {
	// Experimental.
	Destination *ApplicationDestination `json:"destination"`
	// Experimental.
	IgnoreDifferences *[]*ResourceIgnoreDifferences `json:"ignoreDifferences"`
	// Experimental.
	Project *string `json:"project"`
	// Experimental.
	Source *ApplicationSource `json:"source"`
	// Experimental.
	Sources *[]*ApplicationSource `json:"sources"`
	// Experimental.
	SyncPolicy *ApplicationSyncPolicy `json:"syncPolicy"`
}

// Experimental.
type ArgoCdProject interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ArgoCdProject
type jsiiProxy_ArgoCdProject struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ArgoCdProject) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdProject) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdProject) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdProject) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdProject) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoCdProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewArgoCdProject(scope constructs.Construct, id *string, props *ArgoCdProjectProps) ArgoCdProject {
	_init_.Initialize()

	j := jsiiProxy_ArgoCdProject{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewArgoCdProject_Override(a ArgoCdProject, scope constructs.Construct, id *string, props *ArgoCdProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdProject",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ArgoCdProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for an ingress object. https://github.com/kubernetes-sigs/aws-load-balancer-controller.
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func ArgoCdProject_Manifest(props *ArgoCdProjectProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdProject",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
// Experimental.
func ArgoCdProject_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdProject",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ArgoCdProject_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-argocd-resources.ArgoCdProject",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
// Experimental.
func (a *jsiiProxy_ArgoCdProject) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
// Experimental.
func (a *jsiiProxy_ArgoCdProject) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
// Experimental.
func (a *jsiiProxy_ArgoCdProject) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ArgoCdProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ArgoCdProjectProps struct {
	// Experimental.
	Metadata *k8s.ObjectMeta `json:"metadata"`
	// Experimental.
	Spec *ArgoCdProjectSpec `json:"spec"`
}

// Experimental.
type ArgoCdProjectSpec struct {
	// Experimental.
	ClusterResourceWhiteList *[]*ResourceRef `json:"clusterResourceWhiteList"`
	// Experimental.
	Description *string `json:"description"`
	// Experimental.
	Destination *[]*ApplicationDestination `json:"destination"`
	// Experimental.
	NamespaceResourceBlacklist *[]*ResourceRef `json:"namespaceResourceBlacklist"`
	// Experimental.
	NamespaceResourceWhitelist *[]*ResourceRef `json:"namespaceResourceWhitelist"`
	// Experimental.
	Roles *[]*ProjectRoles `json:"roles"`
	// Experimental.
	SourceRepos *[]*string `json:"sourceRepos"`
}

// Experimental.
type HelmOptions struct {
	// Experimental.
	Force *bool `json:"force"`
	// Experimental.
	HelmOptions *[]*string `json:"helmOptions"`
	// Experimental.
	HelmVersion *string `json:"helmVersion"`
	// Experimental.
	Install *bool `json:"install"`
	// Experimental.
	Lint *bool `json:"lint"`
	// Experimental.
	ReleaseName *string `json:"releaseName"`
	// Experimental.
	Repo *string `json:"repo"`
	// Experimental.
	TargetRevision *string `json:"targetRevision"`
	// Experimental.
	Timeout *string `json:"timeout"`
	// Experimental.
	Upgrade *bool `json:"upgrade"`
	// Experimental.
	ValueFiles *[]*string `json:"valueFiles"`
	// Experimental.
	Values *map[string]*string `json:"values"`
	// Experimental.
	ValuesFrom *[]*HelmValuesFromSource `json:"valuesFrom"`
	// Experimental.
	Verify *bool `json:"verify"`
	// Experimental.
	Version *string `json:"version"`
	// Experimental.
	Wait *bool `json:"wait"`
}

// Experimental.
type HelmValuesFromSource struct {
	// Experimental.
	Group *string `json:"group"`
	// Experimental.
	JqPathExpressions *[]*string `json:"jqPathExpressions"`
	// Experimental.
	JsonPointers *[]*string `json:"jsonPointers"`
	// Experimental.
	Kind *string `json:"kind"`
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	Namespace *string `json:"namespace"`
	// Experimental.
	Values *map[string]*string `json:"values"`
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type ProjectRoles struct {
	// Experimental.
	Description *string `json:"description"`
	// Experimental.
	Groups *[]*string `json:"groups"`
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	Policies *[]*string `json:"policies"`
}

// Experimental.
type ResourceIgnoreDifferences struct {
	// Experimental.
	Group *string `json:"group"`
	// Experimental.
	JqPathExpressions *[]*string `json:"jqPathExpressions"`
	// Experimental.
	JsonPointers *[]*string `json:"jsonPointers"`
	// Experimental.
	Kind *string `json:"kind"`
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	Namespace *string `json:"namespace"`
	// Experimental.
	Server *string `json:"server"`
}

// Experimental.
type ResourceRef struct {
	// Experimental.
	Group *string `json:"group"`
	// Experimental.
	Kind *string `json:"kind"`
}

// Experimental.
type RetryBackoff struct {
	// Experimental.
	Duration *string `json:"duration"`
	// Experimental.
	Factor *float64 `json:"factor"`
	// Experimental.
	MaxDuration *string `json:"maxDuration"`
}

// Experimental.
type SyncPolicyAutomated struct {
	// Experimental.
	AllowEmpty *bool `json:"allowEmpty"`
	// Experimental.
	Prune *bool `json:"prune"`
	// Experimental.
	SelfHeal *bool `json:"selfHeal"`
}

// Experimental.
type SyncRetry struct {
	// Experimental.
	Backoff *RetryBackoff `json:"backoff"`
	// Experimental.
	Limit *float64 `json:"limit"`
}

