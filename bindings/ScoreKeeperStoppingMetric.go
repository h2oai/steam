package bindings

type ScoreKeeperStoppingMetric string

const (
	ScoreKeeperStoppingMetric_NONE_         ScoreKeeperStoppingMetric = "_NONE_"
	ScoreKeeperStoppingMetricAUC            ScoreKeeperStoppingMetric = "AUC"
	ScoreKeeperStoppingMetricr2             ScoreKeeperStoppingMetric = "r2"
	ScoreKeeperStoppingMetricdeviance       ScoreKeeperStoppingMetric = "deviance"
	ScoreKeeperStoppingMetricAUTO           ScoreKeeperStoppingMetric = "AUTO"
	ScoreKeeperStoppingMetriclogloss        ScoreKeeperStoppingMetric = "logloss"
	ScoreKeeperStoppingMetriclift_top_group ScoreKeeperStoppingMetric = "lift_top_group"
	ScoreKeeperStoppingMetricMSE            ScoreKeeperStoppingMetric = "MSE"
)
