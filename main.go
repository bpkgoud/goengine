// Package goengine_dev is a set of packages that provide many tools for event sourcing.
//
// goengine_dev contains the following packages:
//
// The aggregate package provides all components needed to create aggregates and record onto them.
//
// TODO write something nice about every sub package
package goengine_dev

// blank imports help docs.
import (
	// aggregate package
	_ "github.com/hellofresh/goengine/aggregate"
	// messaging package
	_ "github.com/hellofresh/goengine/messaging"
	// mocks package
	_ "github.com/hellofresh/goengine/mocks"
)