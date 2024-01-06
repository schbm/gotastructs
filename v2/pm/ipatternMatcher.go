package pm

type PatternMatcher interface {
	Match(string, string) int
}
