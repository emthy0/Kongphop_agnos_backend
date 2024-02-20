package domain

import (
	"context"
	"unicode"

	"github.com/gin-gonic/gin"
)

const (
	PasswordTooShort LengthCheckResult = iota
	PasswordTooLong
	PasswordGoodLength
)
type (
 Password struct {
	plainTextPassword string
}
 LengthCheckResult int8

PasswordService interface {
  GetMinStep(ctx context.Context, pwd string) (int,error)
  
}
PasswordHandler interface {
  GetMinStep() gin.HandlerFunc
}

)

func NewPassword(plainTextPassword string) Password {
	return Password{
		plainTextPassword: plainTextPassword,
	}
}

func (p *Password) length() int {
	return len(p.plainTextPassword)
}

func (p *Password) findRepeatingCharSequence() []int {
	var repeatingCounts []int
	input := p.plainTextPassword
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count >= 3 {
				repeatingCounts = append(repeatingCounts, count)
			}
			count = 1
		}
	}
	if count > 3 {
		repeatingCounts = append(repeatingCounts, count)
	}
	return repeatingCounts
}

func (p *Password) checkLengthRule() LengthCheckResult {
	length := p.length()
	if length >= 20 {
		return PasswordTooLong
	}
	if length < 6 {
		return PasswordTooShort
	} else {
		return PasswordGoodLength
	}
}
func (p *Password) containsLowerCase() bool {
	for _, char := range p.plainTextPassword {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func (p *Password) containsUpperCase() bool {
	for _, char := range p.plainTextPassword {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func (p *Password) containsDigit() bool {
	for _, char := range p.plainTextPassword {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func (p *Password) checkCase() int {
	return B2i(!p.containsDigit()) + B2i(!p.containsLowerCase()) + B2i(!p.containsUpperCase())
}

func (p *Password) GetMinSteps() int {
	fixRep := 0
	repeateSeq := p.findRepeatingCharSequence()
	for _, s := range repeateSeq {
		fixRep += s / 3
	}

	fixLength := 0

	switch p.checkLengthRule() {
	case PasswordTooLong:
		fixLength = p.length() - 19
	case PasswordTooShort:
		fixLength = 6 - p.length()
	}

	fixCase := p.checkCase()

	if p.checkLengthRule() == PasswordTooLong && fixCase > 0 {
		if fixLength > fixRep {
			return fixLength + fixCase
		} 
		return max(fixRep - fixLength ,fixCase) + fixLength
		
	} 
	return max(fixRep, fixLength, fixCase)
}
