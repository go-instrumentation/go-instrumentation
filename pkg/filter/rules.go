package filter

type Rule []Filter

func (r Rule) Allow(targetObject Object) (allow bool) {
	for _, f := range r {
		if !f.Allow(targetObject) {
			return false
		}
	}
	return true
}

var (
	RuleDenyGoInstrumentationFamily = Rule{
		Contains{
			AllowList: nil,
			DenyList: []string{
				"github.com/go-instrumentation",
			},
		},
	}
	RuleDenyPbDotGo = Rule{
		Regex{
			AllowList: nil,
			DenyList: []string{
				":.*\\.pb\\.go::",
			},
		},
	}
	RuleDenyProtobuf = Rule{
		Contains{
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
			AllowList: nil,
			DenyList: []string{
				"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc",
				"github.com/moby/buildkit/client",
			},
		},
	}
	RuleDenyInternal = Rule{
		Contains{
			AllowList: nil,
			DenyList: []string{
				"internal",
			},
		},
	}
	RuleDenyGolang = Rule{
		Prefix{
			AllowList: nil,
			DenyList: []string{
				"runtime*",
			},
		},
		Contains{
			AllowList: nil,
			DenyList: []string{
				"golang.org",
			},
		},
	}
	RuleDenyJaeger = Rule{
		Contains{
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
