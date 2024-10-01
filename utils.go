package hang

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}
	return true
}

// Variables globales
var diff int
var tableauMot []string
var motATrouver []string
var motJeu string
var lettresEssayees []string
var essaisRestants = 10
var nbessais int
