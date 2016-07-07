package bindings

type ModelCategory int

const (
	ModelCategory_NONE_ ModelCategory = iota
	ModelCategoryClustering
	ModelCategoryAutoEncoder
	ModelCategoryDimReduction
	ModelCategoryMultinomial
	ModelCategoryUnknown
	ModelCategoryBinomial
	ModelCategoryRegression
)
