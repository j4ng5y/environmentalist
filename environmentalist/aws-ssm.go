package environmentalist

// SSM is a data structure that holds relevent data for accessing AWS SSM
type SSM struct {
	Credentials string
}

// NewSSM returns a new pointer to a new instance of SSM
//
// Arguments:
//     None
//
// Returns:
//     (*SSM): A pointer to a new instance of SSM
func NewSSM() *SSM {
	var S SSM

	return &S
}
