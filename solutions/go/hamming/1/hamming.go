package hamming
import "errors"

// Distance calculates the Hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
    // Check if strands are of equal length
    if len(a)!= len(b) {
        return 0, errors.New("strands must be of equal length")
    } 
// Count differences
    distance := 0
    for i:= 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance++
        }
    }
    return distance, nil
}
