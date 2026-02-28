package dna
import (
    "errors"
)

// 1. Define DNA as a string type
type DNA string

// 2. Define Histogram as a map from nucleotide to count
type Histogram map[byte]int

// 3. Add Counts method to DNA type
func (d DNA) Counts() (Histogram, error) {
    // 4. Initialize histogram with all nucleotides set to 0
    h := Histogram{
        'A': 0,
        'C': 0,
        'G': 0,
        'T': 0,
    }
    
    // 5. Iterate through each character in DNA
    for i := 0; i < len(d); i++ {
        nucleotide := d[i]
        
        // 6. Check if valid nucleotide
        switch nucleotide {
        case 'A', 'C', 'G', 'T':
            h[nucleotide]++  // Increment count
        default:
            return nil, errors.New("invalid nucleotide")  // Error on invalid
        }
    }
    
    // 7. Return histogram and nil error (success)
    return h, nil
}
