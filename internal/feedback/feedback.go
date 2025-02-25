package feedback

import "fmt"

type Feedback struct {
	Text     string `json:"text"`
	Feedback string `json:"feedback"`
}

func GenerateFeedback(text string) (*Feedback, error) {
	// Generate feedback (this is just an example logic)
	feedbackText := fmt.Sprintf("Generated feedback for: %s", text)
	feedback := &Feedback{
		Text:     text,
		Feedback: feedbackText,
	}
	return feedback, nil
}
