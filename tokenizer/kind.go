package tokenizer

type Kind int

const (
	ILLEGAL Kind = iota
	LCUB         // {
	RCUB         // }
	LSQB         // [
	RSQB         // ]
	DQUO         // "
	COLN         // :
	KEY          //
	VAL
)
