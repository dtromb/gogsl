package numint

type IntegrationRule int32
const (
	GSL_INTEG_GAUSS15  IntegrationRule 	= 1
	GSL_INTEG_GAUSS21  IntegrationRule 	= 2
	GSL_INTEG_GAUSS31  IntegrationRule 	= 3
	GSL_INTEG_GAUSS41  IntegrationRule 	= 4
	GSL_INTEG_GAUSS51  IntegrationRule 	= 5
	GSL_INTEG_GAUSS61  IntegrationRule 	= 6
)