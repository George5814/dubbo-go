package auth

import (
	"github.com/apache/dubbo-go/common/constant"
	"github.com/apache/dubbo-go/common/extension"
	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/filter"
	"github.com/apache/dubbo-go/protocol"
)

type ProviderAuthFilter struct {
}

func init() {
	extension.SetFilter(constant.PROVIDER_AUTH_FILTER, getProviderAuthFilter)
}

func (paf *ProviderAuthFilter) Invoke(invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	logger.Infof("invoking providerAuth filter.")
	url := invoker.GetUrl()

	err := doAuthWork(&url, func(authenticator filter.Authenticator) error {
		return authenticator.Authenticate(invocation, &url)
	})
	if err != nil {
		logger.Infof("auth the request: %v occur exception, cause: %s", invocation, err.Error())
		return &protocol.RPCResult{
			Err: err,
		}
	}

	return invoker.Invoke(invocation)
}

func (paf *ProviderAuthFilter) OnResponse(result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	return result
}
func getProviderAuthFilter() filter.Filter {
	return &ProviderAuthFilter{}
}
