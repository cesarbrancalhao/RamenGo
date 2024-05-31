package filter

import "strconv"

func FindOrderRecipeId(b string, p string) string {
	broth := map[string]int { "1": 0, "2": 3, "3": 6 }
	protein := map[string]int { "1": 1, "2": 2, "3": 3 }
	
	responseId := strconv.Itoa(broth[b] + protein[p])
	return responseId
}