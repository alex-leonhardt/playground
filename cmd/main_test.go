// Package main ...
//
// TestMain will execute tests wrapped into a func that will push a gauge value to prometheus pushgateway URL
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/alex-leonhardt/playground/pkg/metrics"
)

// TestMain runs a dummy the tests..
func TestMain(t *testing.T) {
	fmt.Println("playground")

	t.Run(
		"MyTest", metrics.WrappedTestWithGaugeMetric(t, testMyTest),
	)

	fmt.Println("done")

}

func testMyTest(t *testing.T) {

	t.Log("sleeping start")
	time.Sleep(2 * time.Second)
	t.Log("sleeping done")

}
