package filter

type Rule []Filter

func (r Rule) Match(pkg, functionName string) (result bool) {
	for _, f := range r {
		if !f.Match(pkg, functionName) {
			return false
		}
	}
	return true
}

var (
	RulePassProtobuf = Rule{
		Contains{
			AllowList: nil,
			DenyList: []string{
				"github.com/golang/protobuf",
				"github.com/gogo/protobuf",
			},
		},
	}
	RuleOnlyMain = Rule{
		Prefix{
			AllowList: []string{
				"main",
			},
			DenyList: nil,
		},
	}
	// RulePassAlreadyUseJaeger
	// if you are using vendor mode, the jaeger in the vendor may be difference with $GOSRC/go_instrumentation/jaeger
	RulePassAlreadyUseJaeger = Rule{
		Contains{
			AllowList: nil,
			DenyList: []string{
				"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc",
				"github.com/moby/buildkit/client",
			},
		},
	}
	RulePassInternal = Rule{
		Contains{
			AllowList: nil,
			DenyList: []string{
				"internal",
			},
		},
	}
	RulePassGolang = Rule{
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
	RulePassJaeger = Rule{
		Contains{
			AllowList: nil,
			DenyList: []string{
				"github.com/opentracing/opentracing-go",
				"github.com/uber/jaeger-client-go",
				"github.com/uber/jaeger-lib",
			},
		},
	}
	RulePassTooManyDetails = Rule{
		GoRootFilter,
		RulePassProtobuf,
		RulePassJaeger,
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
