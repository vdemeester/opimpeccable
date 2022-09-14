/*
Copyright 2020 The OpenShift Pipelines Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"

	// The set of controllers this controller process runs.
	"github.com/tektoncd/operator/pkg/reconciler/openshift/openshiftplatform"
	"github.com/tektoncd/operator/pkg/reconciler/platform"
	"github.com/vdemeester/opimpeccable/pkg/reconciler/simpledeployment"

	// This defines the shared main for injected controllers.
	installer "github.com/tektoncd/operator/pkg/reconciler/shared/tektoninstallerset"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"
)

func main() {
	pConfig := platform.NewConfigFromFlags()
	p := openshiftplatform.NewOpenShiftPlatform(pConfig)
	controllers := []injection.ControllerConstructor{
		simpledeployment.NewController,
	}
	for _, c := range p.AllSupportedControllers() {
		controllers = append(controllers, c.ControllerConstructor)
	}
	cfg := injection.ParseAndGetRESTConfigOrDie()
	cfg.QPS = 50
	ctx, _ := injection.EnableInjectionOrDie(signals.NewContext(), cfg)
	ctx = contextWithPlatformName(ctx, "openshift")
	installer.InitTektonInstallerSetClient(ctx)
	sharedmain.MainWithConfig(ctx,
		"controller",
		cfg,
		controllers...,
	)
	// sharedmain.Main("controller", controllers...)
}

// contextWithPlatformName  adds platform name to a given context
func contextWithPlatformName(ctx context.Context, pName string) context.Context {
	ctx = context.WithValue(ctx, PlatformNameKey{}, pName)
	return ctx
}

// PlatformNameKey is defines a 'key' for adding platform name to an instance of context.Context
type PlatformNameKey struct{}
