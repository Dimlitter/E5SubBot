/*
GraphFlow defines some task flows with sequences of Graph APIs to simulate real-
life tasks, i.e. reading emails, uploading files to OneDrive, checking user
status, etc.
*/

package microsoft

import (
	"time"
)

func GetBuiltinPool() GraphFlowPool {
	p := GraphFlowPool{}
	p.setSeed(time.Now().UnixNano())
	p.pool = []NamedGraphFlow{
		{
			"Get my emails",
			GraphFlow{[]GenericFlow{
				GetOutlookMails,
			}},
		},
		{
			"Get my profiles",
			GraphFlow{[]GenericFlow{
				GetUserInfo,
			}},
		},
	}

	return p
}
