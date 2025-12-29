package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
    if n<=0 {
        return 0, errors.New("Only Positive numbers are allowed")
    }
    step:=0
    current:= n

    for current!=1{
        if current%2 == 0 {
            current = current / 2
        } else {
            current = 3 * current + 1
        }
        step++
    }
    return step, nil
}
