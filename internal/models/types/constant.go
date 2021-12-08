package types

const (
	PoolClientTimeoutSeconds            = 5
	PoolTransportMaxIdleConns           = 100
	PoolTransportMaxIdleConnsPerHost    = 2
	PoolTransportIdleConnTimeoutSeconds = 90

	BreakerPartnerApiTimeout = 10
	BreakerExtApiTimeout     = 5
	BreakerTimeout           = 5
	BreakerDeliveryTimeout   = 10
	BreakerErrorThreshold    = 10
	BreakerSuccessThreshold  = 1

	PR_APPROVED    = "APPROVED"
	PR_COMMENTED   = "COMMENTED"
	PR_REQ_CHANGES = "REQUEST_CHANGES"
)
