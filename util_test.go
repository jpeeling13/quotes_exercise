package quotes_exercise

import (
	"testing"
)

func TestFavs(t *testing.T){
	quotes := favs()
	if quotes[0] != "\"To be or not to be, that is the question.\" - Hamlet"{
		t.Errorf("quotes[0] = %s; want \"To be or not to be, that is the question.\" - Hamlet", quotes[0])
	}

	if len(quotes) != 5{
		t.Errorf("len(quotes) = %d; want 5", len(quotes))
	}
}
