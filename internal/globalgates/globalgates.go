// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package globalgates // import "go.opentelemetry.io/collector/internal/globalgates"

import "go.opentelemetry.io/collector/featuregate"

var UseUnifiedEnvVarExpansionRules = featuregate.GlobalRegistry().MustRegister("confmap.unifyEnvVarExpansion",
	featuregate.StageStable,
	featuregate.WithRegisterFromVersion("v0.103.0"),
	featuregate.WithRegisterToVersion("v0.109.0"),
	featuregate.WithRegisterDescription("`${FOO}` will now be expanded as if it was `${env:FOO}` and no longer expands $ENV syntax. See https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/rfcs/env-vars.md for more details. When this feature gate is stable, expandconverter will be removed."))

var NoopTracerProvider = featuregate.GlobalRegistry().MustRegister("service.noopTracerProvider",
	featuregate.StageAlpha,
	featuregate.WithRegisterFromVersion("v0.107.0"),
	featuregate.WithRegisterToVersion("v0.109.0"),
	featuregate.WithRegisterDescription("Sets a Noop OpenTelemetry TracerProvider to reduce memory allocations. This featuregate is incompatible with the zPages extension."))
