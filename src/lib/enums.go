package lib

const (
  EOF = -1
  NULL = 0
  AND = 1
  ARRAY = 2
  BEGIN = 3
  DIV = 4
  DO = 5
  ELSE = 6
  END = 7
  FUNCTION = 8
  IF = 9
  INTEGER = 10
  MOD = 11
  NOT = 12
  NUM = 13
  OF = 14
  OR = 15
  PROCEDURE = 16
  PROGRAM = 17
  REAL = 18
  THEN = 19
  VAR = 20
  WHILE = 21
  ID = 22
  LONG_REAL = 23
  INT = 24
  RELOP = 25
  ADDOP = 26
  MULOP = 27
  ASSIGNOP = 28
  OPEN_BRACKET = 29
  CLOSE_BRACKET = 30
  OPEN_PAREN = 31
  CLOSE_PAREN = 32
  SEMICOLON = 33
  COLON = 34
  COMMA = 35
  PERIOD = 36
  RANGE = 37
  LEXERR = 38
  EQ = 39
  NEQ = 40
  LT = 41
  LEQ = 42
  GEQ = 43
  GT = 44
  PLUS = 45
  MINUS = 46
  ASTERISK = 47
  SLASH = 48
  UNRECOGNIZED_SYMBOL = 49
  ID_TOO_LONG = 50
  XX_TOO_LONG = 51
  YY_TOO_LONG = 52
  ZZ_TOO_LONG = 53
)

func Annotate(c int) string {
  switch c {
    case -1: return "($)"
    case 0: return "(NULL)"
    case 1: return "(AND)"
    case 2: return "(ARRAY)"
    case 3: return "(BEGIN)"
    case 4: return "(DIV)"
    case 5: return "(DO)"
    case 6: return "(ELSE)"
    case 7: return "(END)"
    case 8: return "(FUNCTION)"
    case 9: return "(IF)"
    case 10: return "(INTEGER)"
    case 11: return "(MOD)"
    case 12: return "(NOT)"
    case 13: return "(NUM)"
    case 14: return "(OF)"
    case 15: return "(OR)"
    case 16: return "(PROCEDURE)"
    case 17: return "(PROGRAM)"
    case 18: return "(REAL)"
    case 19: return "(THEN)"
    case 20: return "(VAR)"
    case 21: return "(WHILE)"
    case 22: return "(ID)"
    case 23: return "(LONG_REAL)"
    case 24: return "(INT)"
    case 25: return "(RELOP)"
    case 26: return "(ADDOP)"
    case 27: return "(MULOP)"
    case 28: return "(ASSINGOP)"
    case 29: return "(OPEN_BRACKET)"
    case 30: return "(CLOSE_BRACKET)"
    case 31: return "(OPEN_PAREN)"
    case 32: return "(CLOSE_PAREN)"
    case 33: return "(SEMICOLON)"
    case 34: return "(COLON)"
    case 35: return "(COMMA)"
    case 36: return "(PERIOD)"
    case 37: return "(RANGE)"
    case 38: return "(LEXERR)"
    case 39: return "(EQ)"
    case 40: return "(NEQ)"
    case 41: return "(LT)"
    case 42: return "(LEQ)"
    case 43: return "(GEQ)"
    case 44: return "(GT)"
    case 45: return "(PLUS)"
    case 46: return "(MINUS)"
    case 47: return "(ASTERISK)"
    case 48: return "(SLASH)"
  }
  return "()"
}