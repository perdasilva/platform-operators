package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ActiveBundleDeployment = map[string]string{
	"":     "ActiveBundleDeployment references a BundleDeployment resource.",
	"name": "name is the metadata.name of the referenced BundleDeployment object.",
}

func (ActiveBundleDeployment) SwaggerDoc() map[string]string {
	return map_ActiveBundleDeployment
}

var map_Package = map[string]string{
	"":     "Package contains fields to configure which OLM package this PlatformOperator will install",
	"name": "name contains the desired OLM-based Operator package name that is defined in an existing CatalogSource resource in the cluster.\n\nThis configured package will be managed with the cluster's lifecycle. In the current implementation, it will be retrieving this name from a list of supported operators out of the catalogs included with OpenShift.",
}

func (Package) SwaggerDoc() map[string]string {
	return map_Package
}

var map_PlatformOperator = map[string]string{
	"": "PlatformOperator is the Schema for the PlatformOperators API.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (PlatformOperator) SwaggerDoc() map[string]string {
	return map_PlatformOperator
}

var map_PlatformOperatorList = map[string]string{
	"": "PlatformOperatorList contains a list of PlatformOperators\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (PlatformOperatorList) SwaggerDoc() map[string]string {
	return map_PlatformOperatorList
}

var map_PlatformOperatorSpec = map[string]string{
	"":        "PlatformOperatorSpec defines the desired state of PlatformOperator.",
	"package": "package contains the desired package and its configuration for this PlatformOperator.",
}

func (PlatformOperatorSpec) SwaggerDoc() map[string]string {
	return map_PlatformOperatorSpec
}

var map_PlatformOperatorStatus = map[string]string{
	"":                       "PlatformOperatorStatus defines the observed state of PlatformOperator",
	"conditions":             "conditions represent the latest available observations of a platform operator's current state.",
	"activeBundleDeployment": "activeBundleDeployment is the reference to the BundleDeployment resource that's being managed by this PO resource. If this field is not populated in the status then it means the PlatformOperator has either not been installed yet or is failing to install.",
}

func (PlatformOperatorStatus) SwaggerDoc() map[string]string {
	return map_PlatformOperatorStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
