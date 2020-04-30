// Package proverb provides proverbs generator.
package proverb

// Proverb returns rhyme strings from 'For Want of a Nail'.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	proverb := make([]string, 0, len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		str := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		proverb = append(proverb, str)
	}

	proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
	return proverb
}
