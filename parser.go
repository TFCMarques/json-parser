package main

const MAX_DEPTH = 20

func parseValue(tokens []Token, pos *int, depth int) bool {
	if depth >= MAX_DEPTH {
		return false
	}

	switch tokens[*pos].Type {
	case STRING, NUMBER, BOOLEAN, NULL:
		*pos++
		return true
	case LEFT_BRACE:
		return parseObject(tokens, pos, depth+1)
	case LEFT_BRACKET:
		return parseArray(tokens, pos, depth+1)
	default:
		return false
	}
}

func parseObject(tokens []Token, pos *int, depth int) bool {
	if tokens[*pos].Type != LEFT_BRACE || depth >= MAX_DEPTH {
		return false
	}

	*pos++

	for tokens[*pos].Type != RIGHT_BRACE {
		if tokens[*pos].Type != STRING {
			return false
		}

		*pos++

		if tokens[*pos].Type != COLON {
			return false
		}

		*pos++

		if !parseValue(tokens, pos, depth) {
			return false
		}

		if tokens[*pos].Type == COMMA {
			if tokens[*pos+1].Type == RIGHT_BRACE {
				return false
			}

			*pos++
		} else if tokens[*pos].Type != RIGHT_BRACE {
			return false
		}
	}

	*pos++

	return true
}

func parseArray(tokens []Token, pos *int, depth int) bool {
	if tokens[*pos].Type != LEFT_BRACKET || depth >= MAX_DEPTH {
		return false
	}

	*pos++

	for tokens[*pos].Type != RIGHT_BRACKET {
		if !parseValue(tokens, pos, depth) {
			return false
		}

		if tokens[*pos].Type == COMMA {
			*pos++
		} else if tokens[*pos].Type != RIGHT_BRACKET {
			return false
		}
	}

	*pos++

	return true
}

func parser(tokens []Token) string {
	pos := 0

	if tokens[0].Type == LEFT_BRACE {
		if parseObject(tokens, &pos, 1) && tokens[pos].Type == EOF {
			return "Valid JSON"
		}
	} else if tokens[0].Type == LEFT_BRACKET {
		if parseArray(tokens, &pos, 1) && tokens[pos].Type == EOF {
			return "Valid JSON"
		}
	}

	return "Invalid JSON"
}
