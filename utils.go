package utils

func Contains(array []string, target string) bool {
    for _, el := range(array) {
        if (el == target) {
            return true
        }
    }
    return false
}

func ContainsInt(array []int, target int) bool {
    for _, el := range(array) {
        if (el == target) {
            return true
        }
    }
    return false
}
