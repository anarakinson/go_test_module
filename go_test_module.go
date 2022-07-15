package utils

func Contains(array []string, target string) bool {
    for _, el := range(array) {
        if (el == target) {
            return true
        }
    }
    return false
}
