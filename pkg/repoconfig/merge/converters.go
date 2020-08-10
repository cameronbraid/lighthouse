package merge

import (
	"github.com/jenkins-x/lighthouse/pkg/apis/lighthouse/v1alpha1"
	"github.com/jenkins-x/lighthouse/pkg/config"
	"github.com/jenkins-x/lighthouse/pkg/plugins"
)

// ToPresubmit converts the repo config to the lighthouse config resource
func ToPresubmit(r v1alpha1.Presubmit) config.Presubmit {
	return config.Presubmit{
		JobBase:             ToJobBase(r.JobBase),
		AlwaysRun:           r.AlwaysRun,
		Optional:            r.Optional,
		Trigger:             r.Trigger,
		RerunCommand:        r.RerunCommand,
		Brancher:            ToBrancher(r.Brancher),
		RegexpChangeMatcher: ToRegexpChangeMatcher(r.RegexpChangeMatcher),
		Reporter:            ToReporter(r.Reporter),
	}
}

// ToPostsubmit converts the repo config to the lighthouse config resource
func ToPostsubmit(r v1alpha1.Postsubmit) config.Postsubmit {
	return config.Postsubmit{
		JobBase:             ToJobBase(r.JobBase),
		RegexpChangeMatcher: ToRegexpChangeMatcher(r.RegexpChangeMatcher),
		Brancher:            ToBrancher(r.Brancher),
		Reporter:            ToReporter(r.Reporter),
	}
}

// ToJobBase converts the repo config to the lighthouse config resource
func ToJobBase(r v1alpha1.JobBase) config.JobBase {
	return config.JobBase{
		Name:            r.Name,
		Labels:          r.Labels,
		Annotations:     r.Annotations,
		MaxConcurrency:  r.MaxConcurrency,
		Agent:           r.Agent,
		Cluster:         r.Cluster,
		Namespace:       r.Namespace,
		ErrorOnEviction: r.ErrorOnEviction,
		SourcePath:      r.SourcePath,
		Spec:            r.Spec,
		PipelineRunSpec: r.PipelineRunSpec,
		UtilityConfig:   ToUtilityConfig(r.UtilityConfig),
	}
}

// ToBrancher converts the repo config to the lighthouse config resource
func ToBrancher(r v1alpha1.Brancher) config.Brancher {
	return config.Brancher{
		SkipBranches: r.SkipBranches,
		Branches:     r.Branches,
	}
}

// ToReporter converts the repo config to the lighthouse config resource
func ToReporter(r v1alpha1.Reporter) config.Reporter {
	return config.Reporter{
		Context:    r.Context,
		SkipReport: r.SkipReport,
	}
}

// ToRegexpChangeMatcher converts the repo config to the lighthouse config resource
func ToRegexpChangeMatcher(r v1alpha1.RegexpChangeMatcher) config.RegexpChangeMatcher {
	return config.RegexpChangeMatcher{RunIfChanged: r.RunIfChanged}
}

// ToUtilityConfig converts the repo config to the lighthouse config resource
func ToUtilityConfig(r v1alpha1.UtilityConfig) config.UtilityConfig {
	return config.UtilityConfig{
		Decorate:       r.Decorate,
		PathAlias:      r.PathAlias,
		CloneURI:       r.CloneURI,
		SkipSubmodules: r.SkipSubmodules,
		CloneDepth:     r.CloneDepth,
	}
}

// ToApprove converts the repo config to the lighthouse config resource
func ToApprove(r *v1alpha1.Approve, repoKey string) plugins.Approve {
	return plugins.Approve{
		Repos:               []string{repoKey},
		IssueRequired:       r.IssueRequired,
		RequireSelfApproval: r.RequireSelfApproval,
		LgtmActsAsApprove:   r.LgtmActsAsApprove,
		IgnoreReviewState:   r.IgnoreReviewState,
	}
}
