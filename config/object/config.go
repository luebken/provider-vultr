/*
Copyright 2021 Upbound Inc.
*/

package object

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("vultr_object_storage", func(r *ujconfig.Resource) {
		// And other overrides.
	})
}
