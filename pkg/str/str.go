package str


// Plural 转为复数 user -> users

func Plural(word string) string {
    return pluralize.NewClient().Plural(word)
}

// Singular 转为单数 users -> user
func Singular(word string) string {
    return pluralize.NewClient().Singular(word)
}

