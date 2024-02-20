package test

import (
	"testing"

	"authen.agnoshealth.com/core/password"
	"github.com/stretchr/testify/assert"
)

func TestPasswordTooShort(t *testing.T) {

  rawpwd := "Ab1"
  pwd := password.NewPassword(rawpwd)

  assert.Equal(t, 3, pwd.GetMinSteps(), "expecting result to be 3")
}

func TestPasswordTooLong(t *testing.T) {
  rawpwd := "Ab1Ab1Ab1Ab1Ab1Ab1Ab1"
  pwd := password.NewPassword(rawpwd)

  assert.Equal(t, 2, pwd.GetMinSteps(), "expecting result to be 2")
}

func TestPasswordMissingCase(t *testing.T) {
  rawpwd := "AA4455"
  pwd := password.NewPassword(rawpwd)
  assert.Equal(t, 1, pwd.GetMinSteps(), "expecting result to be 1")
}

func TestPasswordRepeated(t *testing.T) {
  rawpwd := "AAA4445566666"
  pwd := password.NewPassword(rawpwd)
  assert.Equal(t, 3, pwd.GetMinSteps())
}

func TestMixedCase(t *testing.T ) {
  // case short and missing case
  rawpwd := "AAVV2"
  pwd := password.NewPassword(rawpwd)
  assert.Equal(t, 1, pwd.GetMinSteps(), "expecting result to be 1")

  // case short and missing 2 case
  rawpwd = "AAVVA"
  pwd = password.NewPassword(rawpwd)
  assert.Equal(t, 2, pwd.GetMinSteps(), "expecting result to be 2")

  rawpwd = "AADDDDDDDD"
  pwd = password.NewPassword(rawpwd)
  assert.Equal(t, 2, pwd.GetMinSteps(), "expecting result to be 2")

  rawpwd = "AADDDDDDD2"
  pwd = password.NewPassword(rawpwd)
  assert.Equal(t, 2, pwd.GetMinSteps(), "expecting result to be 2")

  rawpwd = "AADDDDDDD2AADDDDDDD2AADDDDDDD2"
  // remove 11 and replace 1
  pwd = password.NewPassword(rawpwd)
  assert.Equal(t, 12, pwd.GetMinSteps(), "expecting result to be 12")
}

