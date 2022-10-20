package microsoft

import (
	"math/rand"
)

type GenericFlow = func(string, string, string) (string, string, error)

// API arguments: id, secret, refresh (string)
type GraphFlow struct {
	flow []GenericFlow
}

/*
Returns:

	newRefresh (string):	The latest refresh token.
	success(int):			The number of succeeded flow APIs.
	err (error):			The first error of flow APIs or others.
*/
func (gf *GraphFlow) Run(id, secret, refresh string) (string, int, error) {
	/*
		This is a starting point, like an application of `reduce()` in some
		languages.
		TODO: Can be replaced by native functional-programming functions if any.
	*/
	var (
		token string = refresh
		err   error  = nil
	)

	for i, Api := range gf.flow {
		token, _, err = Api(id, secret, token)
		if err != nil {
			return token, i, err
		}
	}

	return token, len(gf.flow), nil
}

type NamedGraphFlow struct {
	name string
	flow GraphFlow
}

type GraphFlowPool struct {
	pool  []NamedGraphFlow
	_rand rand.Rand
}

func (p *GraphFlowPool) SetSeed(seed int64) {
	p._rand.Seed(seed)
}

func (p *GraphFlowPool) GetFlow() *GraphFlow {
	return &p.pool[p._rand.Intn(len(p.pool))].flow
}
