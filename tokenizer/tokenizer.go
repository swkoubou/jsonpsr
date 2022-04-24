package tokenizer

import (
	"log"
	"unicode"
)

type Tokenizer struct {
	input []rune
	pos   int
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		input: []rune{},
		pos:   0,
	}
}

// 現在地の文字取得
func (t *Tokenizer) curt() rune {
	return t.input[t.pos]
}

// 一つ前の文字取得
func (t *Tokenizer) prev() rune {
	return t.input[t.pos-1]
}

// 一つ先の文字取得
func (t *Tokenizer) next() rune {
	return t.input[t.pos+1]
}

// 一つ先に進む
func (t *Tokenizer) goNext() {
	t.pos++
}

// curtとgoNextが混ざってる
func (t *Tokenizer) consume() rune {
	c := t.curt()
	t.goNext()
	return c
}

// 終端じゃないか
func (t *Tokenizer) isEof() bool {
	return t.pos >= len(t.input)
}

// string
func (t *Tokenizer) consumeString() Token {
	// 分岐してから何もいじっていないので、
	// t.curtは「"」
	// ここから次の「"」までのデータを取得したい。

	// 文字列の初めの「"」を消す。(Kind.STRINGというデータがあれば、文字列だということがわかるので。)
	s := t.pos
	t.goNext()
	str := ""
	// 「for t.curt != '"'」で回さないように注意。
	// (「\"」のようにエスケープされている場合は、終了の合図ではなく、文字列の中身として扱うため。)
	// => 文字列が終わっていないのに、ループが終了する可能性があるため。
	for !t.isEof() {
		c := t.curt()
		if c == '"' {
			// もし、現在地が、0じゃなくて、前が\だったら、エスケープされているので、終了しない。
			// 0じゃなくて =>
			//   prev()に、posをチェックする機能がないので、0でやると、-1を参照し、エラーとなる。
			// (!! 仕様書では、エスケープできる文字が限られているが、簡略化のため、全部通す。 !!)
			if t.pos != 0 && t.prev() == '\\' {
				// 文字列
				str += string(c)
				t.goNext()
				continue
			} else {
				// 終了
				break
			}
		}
		str += string(c)
		t.goNext()
	}

	// " を消費してあげて、終了
	t.goNext()
	e := t.pos

	return Token{
		STRING,
		str,
		s,
		e,
	}
}

// number
func (t *Tokenizer) consumeNumber() Token {
	num := ""
	s := t.pos
	for !t.isEof() {
		c := t.curt()
		if c == '-' || (48 <= c && c <= 57) || c == '.' {
			num += string(c)
		} else {
			// 消費の対象じゃなかったら触れずに終了
			break
		}
		t.goNext()
	}
	e := t.pos
	return Token{
		NUMBER,
		num,
		s,
		e,
	}
}

// keyword
func (t *Tokenizer) consumeKeyword() Token {
	// 使用できる文字は、アルファベットのみ
	keyword := ""
	s := t.pos
	for !t.isEof() {
		c := t.curt()
		if (65 <= c && c <= 90) || (97 <= c && c <= 122) {
			keyword += string(c)
		} else {
			// 消費の対象じゃなかったら触れずに終了
			break
		}
		t.goNext()
	}
	e := t.pos
	return Token{
		KEYWORD,
		keyword,
		s,
		e,
	}
}

// whitespace
func (t *Tokenizer) consumeWhiteSpace() Token {
	white := ""
	s := t.pos
	for !t.isEof() {
		c := t.curt()
		if unicode.IsSpace(c) {
			white += string(c)
		} else {
			// 消費の対象じゃなかったら触れずに終了
			break
		}
		t.goNext()
	}
	e := t.pos
	return Token{
		WHITESPACE,
		white,
		s,
		e,
	}
}

func (t *Tokenizer) Tokenize(in string) ([]Token, error) {
	t.input = []rune(in)
	t.pos = 0

	// 最終的に返却する入れ物(result)と、ループの中で得たものを入れる(tok)を宣言
	var result []Token
	var tok Token

	for !t.isEof() {
		switch {
		case t.curt() == '{':
			s := t.pos
			c := t.consume()
			e := t.pos
			tok = Token{LCUB, string(c), s, e}
		case t.curt() == '}':
			s := t.pos
			c := t.consume()
			e := t.pos
			tok = Token{RCUB, string(c), s, e}
		case t.curt() == '[':
			s := t.pos
			c := t.consume()
			e := t.pos
			tok = Token{LSQB, string(c), s, e}
		case t.curt() == ']':
			s := t.pos
			c := t.consume()
			e := t.pos
			tok = Token{RSQB, string(c), s, e}
		case t.curt() == ':':
			s := t.pos
			c := t.consume()
			e := t.pos
			tok = Token{COLON, string(c), s, e}
		case t.curt() == ',':
			s := t.pos
			c := t.consume()
			e := t.pos
			tok = Token{COMMA, string(c), s, e}
		case t.curt() == '"':
			// "abc"のように、「"」から始まるものは、基本的に文字列、string。
			tok = t.consumeString()
		case t.curt() == '-' || (48 <= t.curt() && t.curt() <= 57):
			// - or 0-9
			// 「-」、数字から始まるものは、数字、number。(「.」、「+」から始まることはない。)
			tok = t.consumeNumber()
		case (65 <= t.curt() && t.curt() <= 90) || (97 <= t.curt() && t.curt() <= 122):
			// A-Z or a-z
			// trueやfalse、nullのように、突然アルファベットから始まるものは、予約語、keyword。
			tok = t.consumeKeyword()
		default:
			// 改行や空白?
			if unicode.IsSpace(t.curt()) {
				tok = t.consumeWhiteSpace()
			} else {
				// その他は予想していないので、エラーを返す。
				log.Fatalf("`%v` : unexpected (pos=%d) => `%v`", string(t.input), t.pos, string(t.curt()))
			}
		}
		// 中身の入ったtokを結果に追加してあげる。
		if tok.Kind != WHITESPACE {
			result = append(result, tok)
		}
	}
	return result, nil
}
