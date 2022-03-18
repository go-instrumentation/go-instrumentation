package filter

import "fmt"

type Rule []Filter

var GlobalRule Rule

func (r Rule) GetName() string {
	var names []string
	for _, rule := range r {
		names = append(names, rule.GetName())
	}
	return fmt.Sprintf("%q", names)
}

func (r Rule) Allow(targetObject Object) (allow bool) {
	for _, f := range r {
		if !f.Allow(targetObject) {
			debug(f, "", targetObject, allow)
			return
		}
	}
	allow = true
	debug(r, "", targetObject, allow)
	return
}

var (
	RuleDenyGoInstrumentationFamily = Rule{
		Contains{
			Base: Base{
				Name: "RuleDenyGoInstrumentationFamily",
			},
			AllowList: nil,
			DenyList: []string{
				"github.com/go-instrumentation",
			},
		},
	}
	RuleDenyPbDotGo = Rule{
		Regex{
			Base: Base{
				Name: "RuleDenyPbDotGo",
			},
			AllowList: nil,
			DenyList: []string{
				":.*\\.pb\\.go::",
			},
		},
	}
	RuleDenyProtobuf = Rule{
		Contains{
			Base: Base{
				Name: "RuleDenyProtobuf",
			},
			AllowList: nil,
			DenyList: []string{
				"github.com/golang/protobuf",
				"github.com/gogo/protobuf",
			},
		},
		RuleDenyPbDotGo,
	}
	RuleOnlyMain = Rule{
		Prefix{
			Base: Base{
				Name: "RuleOnlyMain",
			},
			AllowList: []string{
				"main",
			},
			DenyList: nil,
		},
	}
	// RuleDenyAlreadyUseJaeger
	// if you are using vendor mode, the jaeger in the vendor may be difference with $GOSRC/go_instrumentation/jaeger
	RuleDenyAlreadyUseJaeger = Rule{
		Contains{
			Base: Base{
				Name: "RuleDenyAlreadyUseJaeger",
			},
			AllowList: nil,
			DenyList: []string{
				"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc",
				"github.com/moby/buildkit/client",
			},
		},
	}
	RuleDenyInternal = Rule{
		Contains{
			Base: Base{
				Name: "RuleDenyInternal",
			},
			AllowList: nil,
			DenyList: []string{
				"internal",
			},
		},
	}
	RuleDenyGolang = Rule{
		Prefix{
			Base: Base{
				Name: "RuleDenyGolang",
			},
			AllowList: nil,
			DenyList: []string{
				"runtime*",
			},
		},
		Contains{
			Base: Base{
				Name: "RuleDenyGolang",
			},
			AllowList: nil,
			DenyList: []string{
				"golang.org",
			},
		},
	}
	RuleDenyJaeger = Rule{
		Contains{
			Base: Base{
				Name: "RuleDenyJaeger",
			},
			AllowList: nil,
			DenyList: []string{
				"github.com/opentracing/opentracing-go",
				"github.com/uber/jaeger-client-go",
				"github.com/uber/jaeger-lib",
			},
		},
	}
	RuleDenyTooManyDetails = Rule{
		GoRootFilter,
		RuleDenyProtobuf,
		RuleDenyGolang,
		RuleDenyJaeger,
		Contains{
			Base: Base{
				Name: "RuleDenyTooManyDetails",
			},
			AllowList: nil,
			DenyList: []string{
				"github.com/checkpoint-restore/go-criu",
				"github.com/urfave/cli",
				"github.com/sirupsen/logrus",
				"github.com/pkg/errors",
				"go.opencensus.io",
				"github.com/davecgh/go-spew/spew",
			},
		},
	}
)
