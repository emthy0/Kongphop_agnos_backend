package password

import (
	"unicode"
)

type password struct {
	plainTextPassword string
}
type LengthCheckResult int8

func NewPassword(plainTextPassword string) password {
	return password{
		plainTextPassword: plainTextPassword,
	}
}

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

const (
	PasswordTooShort LengthCheckResult = iota
	PasswordTooLong
	PasswordGoodLength
)

func (p *password) length() int {
	return len(p.plainTextPassword)
}

func (p *password) findRepeatingCharSequence() []int {
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

func (p *password) checkLengthRule() LengthCheckResult {
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
func (p *password) containsLowerCase() bool {
	for _, char := range p.plainTextPassword {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func (p *password) containsUpperCase() bool {
	for _, char := range p.plainTextPassword {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func (p *password) containsDigit() bool {
	for _, char := range p.plainTextPassword {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func (p *password) checkCase() int {
	return B2i(!p.containsDigit()) + B2i(!p.containsLowerCase()) + B2i(!p.containsUpperCase())
}

func (p *password) GetMinSteps() int {
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
