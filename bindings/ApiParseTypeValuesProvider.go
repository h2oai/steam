package bindings

type ApiParseTypeValuesProvider string

const (
	ApiParseTypeValuesProvider_NONE_   ApiParseTypeValuesProvider = "_NONE_"
	ApiParseTypeValuesProviderGUESS    ApiParseTypeValuesProvider = "GUESS"
	ApiParseTypeValuesProviderSVMLight ApiParseTypeValuesProvider = "SVMLight"
	ApiParseTypeValuesProviderAVRO     ApiParseTypeValuesProvider = "AVRO"
	ApiParseTypeValuesProviderCSV      ApiParseTypeValuesProvider = "CSV"
	ApiParseTypeValuesProviderXLS      ApiParseTypeValuesProvider = "XLS"
	ApiParseTypeValuesProviderARFF     ApiParseTypeValuesProvider = "ARFF"
)
