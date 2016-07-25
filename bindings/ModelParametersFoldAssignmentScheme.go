package bindings

type ModelParametersFoldAssignmentScheme string

const (
	ModelParametersFoldAssignmentScheme_NONE_     ModelParametersFoldAssignmentScheme = "_NONE_"
	ModelParametersFoldAssignmentSchemeAUTO       ModelParametersFoldAssignmentScheme = "AUTO"
	ModelParametersFoldAssignmentSchemeRandom     ModelParametersFoldAssignmentScheme = "Random"
	ModelParametersFoldAssignmentSchemeModulo     ModelParametersFoldAssignmentScheme = "Modulo"
	ModelParametersFoldAssignmentSchemeStratified ModelParametersFoldAssignmentScheme = "Stratified"
)
