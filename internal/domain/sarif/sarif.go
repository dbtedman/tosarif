package sarif

const (
	CurrentVersion = "2.1.0"
	CurrentSchema  = "https://json.schemastore.org/sarif-2.1.0.json"
)

type Root struct {
	Properties PropertyBag `json:"properties"`
	Runs       []Run       `json:"runs"`
	Schema     string      `json:"$schema"`
	Version    string      `json:"version"`
}

type PropertyBag struct {
	Tags []string `json:"tags"`
}

type Run struct {
	Tool Tool `json:"tool"` // required
	// invocations
	// conversion
	// language
	// versionControlProvenance
	// originalUriBaseIds
	// artifacts
	// logicalLocations
	// graphs
	// results
	// automationDetails
	// runAggregates
	// baselineGuid
	// redactionTokens
	// defaultEncoding
	// defaultSourceLanguage
	// newlineSequences
	// columnKind
	// externalPropertyFileReferences
	// threadFlowLocations
	// taxonomies
	// addresses
	// translations
	// policies
	// webRequests
	// webResponses
	// specialLocations
	Properties PropertyBag `json:"properties"`
}

type Tool struct {
	Driver     ToolComponent   `json:"driver"`
	Extensions []ToolComponent `json:"extensions"`
	Properties PropertyBag     `json:"properties"`
}

type ToolComponent struct {
	Name string `json:"name"` // required
	// guid
	// name
	// organization
	// product
	// productSuite
	// shortDescription
	// fullDescription
	// fullName
	// version
	// semanticVersion
	// dottedQuadFileVersion
	// releaseDateUtc
	// downloadUri
	// informationUri
	// globalMessageStrings
	// notifications
	// rules
	// taxa
	// locations
	// language
	// contents
	// isComprehensive
	// localizedDataSemanticVersion
	// minimumRequiredLocalizedDataSemanticVersion
	// associatedComponent
	// translationMetadata
	// supportedTaxonomies
	Properties PropertyBag `json:"properties"`
}
