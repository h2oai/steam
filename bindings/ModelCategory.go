package bindings

type ModelCategory string

const (
	ModelCategory_NONE_       ModelCategory = "NONE"
	ModelCategoryClustering   ModelCategory = "Clustering"
	ModelCategoryAutoEncoder  ModelCategory = "AutoEncoder"
	ModelCategoryDimReduction ModelCategory = "DimReduction"
	ModelCategoryMultinomial  ModelCategory = "Multinomial"
	ModelCategoryUnknown      ModelCategory = "Unknown"
	ModelCategoryBinomial     ModelCategory = "Binomial"
	ModelCategoryRegression   ModelCategory = "Regression"
)
