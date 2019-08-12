// Copyright (c) 2016 shawn1m. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

// Package core implements the essential features.
package core

import (
	"github.com/import-yuefeng/smartDNS/core/config"
	"github.com/import-yuefeng/smartDNS/core/inbound"
	"github.com/import-yuefeng/smartDNS/core/outbound"
)

// InitServer func Initiate the server with config file
func InitServer(configFilePath string) {
	conf := config.NewConfig(configFilePath)
	//New dispatcher without RemoteClientBundle, RemoteClientBundle must be initiated when server is running
	dispatcher := outbound.Dispatcher{

		DNSFilter:          conf.DNSFilter,
		DNSBunch:           conf.DNSBunch,
		RedirectIPv6Record: conf.IPv6UseAlternativeDNS,
		MinimumTTL:         conf.MinimumTTL,
		DomainTTLMap:       conf.DomainTTLMap,
		Hosts:              conf.Hosts,
		Cache:              conf.Cache,
	}
	s := inbound.NewServer(conf.BindAddress, conf.DebugHTTPAddress, dispatcher, conf.RejectQType)

	s.Run()

}