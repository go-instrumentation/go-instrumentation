package filter

type Rule struct {
	Base
	Rule []Filter
}

var GlobalRule Rule

func (r Rule) Allow(targetObject Object) (allow bool) {
	check(r, targetObject)
	for _, f := range r.Rule {
		if !f.Allow(targetObject) {
			debug(f, f.String(), targetObject, allow)
			return
		}
	}
	allow = true
	debug(r, r.String(), targetObject, allow)
	return
}

var (
	RuleDenyGoInstrumentationFamily = Rule{
		Base: Base{
			Name: "RuleDenyGoInstrumentationFamily",
		},
		Rule: []Filter{
			Contains{
				Base: Base{
					Name: "RuleDenyGoInstrumentationFamily",
				},
				AllowList: nil,
				DenyList: []string{
					"github.com/go-instrumentation",
				},
			},
		},
	}
	RuleDenyPbDotGo = Rule{
		Base: Base{Name: "RuleDenyPbDotGo"},
		Rule: []Filter{
			Regex{
				Base: Base{
					Name: "RuleDenyPbDotGo",
				},
				AllowList: nil,
				DenyList: []string{
					":.*\\.pb\\.go::",
				},
			},
		},
	}
	RuleDenyProtobuf = Rule{
		Base: Base{Name: "RuleDenyProtobuf"},
		Rule: []Filter{
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
		},
	}
	RuleOnlyMain = Rule{
		Base: Base{Name: "RuleOnlyMain"},
		Rule: []Filter{
			Prefix{
				Base: Base{
					Name: "RuleOnlyMain",
				},
				AllowList: []string{
					"main",
				},
				DenyList: nil,
			},
		},
	}
	// RuleDenyAlreadyUseJaeger
	// if you are using vendor mode, the jaeger in the vendor may be difference with $GOSRC/go_instrumentation/jaeger
	RuleDenyAlreadyUseJaeger = Rule{
		Base: Base{Name: "RuleDenyAlreadyUseJaeger"},
		Rule: []Filter{
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
		},
	}
	RuleDenyInternal = Rule{
		Base: Base{Name: "RuleDenyInternal"},
		Rule: []Filter{Contains{
			Base: Base{
				Name: "RuleDenyInternal",
			},
			AllowList: nil,
			DenyList: []string{
				"internal",
			},
		},
		},
	}
	RuleDenyGolang = Rule{
		Base: Base{Name: "RuleDenyGolang"},
		Rule: []Filter{
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
		},
	}
	RuleDenyJaeger = Rule{
		Base: Base{Name: "RuleDenyJaeger"},
		Rule: []Filter{
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
		},
	}
	RuleDenyTooManyDetails = Rule{
		Base: Base{Name: "RuleDenyTooManyDetails"},
		Rule: []Filter{
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
		},
	}
)
