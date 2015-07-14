// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package execLogger

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestIndentation(c *C) {
	c.Assert("    * hello world \n", Equals, indent("* hello world \n", 4))
	c.Assert("* hello world", Equals, indent("* hello world", 0))
	c.Assert("   \n    \n    * hello world \n    \n", Equals, indent("\n \n * hello world \n \n", 3))
	c.Assert("  * first\n   *second\n   *third\n", Equals, indent("* first\n *second\n *third\n", 2))
}
