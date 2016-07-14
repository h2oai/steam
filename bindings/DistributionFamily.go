package bindings

type DistributionFamily string

const (
	DistributionFamily_NONE_      DistributionFamily = "_NONE_"
	DistributionFamilymultinomial DistributionFamily = "multinomial"
	DistributionFamilyAUTO        DistributionFamily = "AUTO"
	DistributionFamilyhuber       DistributionFamily = "huber"
	DistributionFamilyquantile    DistributionFamily = "quantile"
	DistributionFamilygaussian    DistributionFamily = "gaussian"
	DistributionFamilytweedie     DistributionFamily = "tweedie"
	DistributionFamilypoisson     DistributionFamily = "poisson"
	DistributionFamilybernoulli   DistributionFamily = "bernoulli"
	DistributionFamilylaplace     DistributionFamily = "laplace"
	DistributionFamilygamma       DistributionFamily = "gamma"
)
