package caesar

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		msg string
		rot int
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{"decoding", args{"Fqjc'b vh wjvn?", 9}, "What's my name?"},
		{"decoding long", args{"pm ol ohk hufaopun jvumpkluaphs av zhf, ol dyval pa pu jpwoly," +
			" aoha pz, if zv johunpun aol vykly vm aol slaalyz vm aol hswohila, aoha uva h dvyk jvbsk il thkl vba.", 7},
			"if he had anything confidential to say, he wrote it in cipher, that is, by so " +
				"changing the order of the letters of the alphabet, that not a word could be made out."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.msg, tt.args.rot); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		msg string
		rot int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hello world", args{"hello world", 7}, "olssv dvysk"},
		{"HELLO WORLD", args{"HELLO WORLD", 7}, "OLSSV DVYSK"},
		{"special characters", args{`;[]=+-_\'\"`, 9}, `;[]=+-_\'\"`},
		{"numbers", args{"0123456789", 13}, "0123456789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.msg, tt.args.rot); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
