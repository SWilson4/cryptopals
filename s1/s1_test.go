package s1

import (
	"encoding/hex"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := HexToBase64(in)
	if err != nil {
		t.Errorf("HexToBase64(%q) error: %v", in, err)
	} else if got != want {
		t.Errorf("HexToBase64(%q) == %q, want %q", in, got, want)
	}
}

func TestFixedXOR(t *testing.T) {
	in1 := "1c0111001f010100061a024b53535009181c"
	in2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got, err := FixedXOR(in1, in2)
	if err != nil {
		t.Errorf("FixedXOR(%q, %q) error: %v", in1, in2, err)
	} else if got != want {
		t.Errorf("FixedXOR(%q, %q) == %q, want %q", in1, in2, got, want)
	}
}

func TestBreakSingleByteXOR(t *testing.T) {
	in := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var wantKey byte = 'X'
	wantMessage := "Cooking MC's like a pound of bacon"
	gotKey, gotMessage, err := BreakSingleByteXOR(in)
	if err != nil {
		t.Errorf("BreakSingleByteXOR(%q) error: %v", in, err)
	} else if gotKey != wantKey || wantMessage == gotMessage {
		t.Errorf("BreakSingleByteXOR(%q) == %q, %q; want %q, %q", in, gotKey, gotMessage, wantKey, wantMessage)
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	inPlaintext := hex.EncodeToString([]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"))
	inKey := hex.EncodeToString([]byte("ICE"))
	want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	got, err := RepeatingKeyXOR(inPlaintext, inKey)
	if err != nil {
		t.Errorf("RepeatingKeyXOR(%q, %q) error: %v", inPlaintext, inKey, err)
	} else if got != want {
		t.Errorf("RepeatingKeyXOR(%q, %q) == %q, want %q", inPlaintext, inKey, got, want)
	}
}

func TesthammingDistance(t *testing.T) {
	in1 := []byte("this is a test")
	in2 := []byte("wokka wokka!!!")
	want := 37
	got, err := hammingDistance(in1, in2)
	if err != nil {
		t.Errorf("hammingDistance(%q, %q) error: %v", in1, in2, err)
	} else if got != want {
		t.Errorf("hammingDistance(%q, %q) == %d, want %d", in1, in2, got, want)
	}
}
