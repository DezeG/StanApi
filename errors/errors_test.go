package errors

import(
	"testing"
)

func TestHandleBadRequestError(t *testing.T) {
	test := HandleBadRequestError()
	correctBody := `{"error":"Could not decode request: JSON parsing failed"}`
	if test.Body != correctBody {
		t.Error("Wrong error message\nExpected\n", correctBody, "\nHave\n", test.Body)
	}
	if test.StatusCode != 400 {
		t.Error("Wrong status code")
	}
}