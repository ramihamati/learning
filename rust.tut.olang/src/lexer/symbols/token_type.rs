use std::fmt::{Display, Formatter};

#[derive(Debug, PartialEq, Clone, Copy)]
pub enum TokenType {
    // literals
    Identifier,
    String,
    Number,

    // Keywords
    If,             // (OK) if
    Else,           // (OK) else
    Loop,           // (OK) while
    Struct,         // (OK) struct
    Func,           // (OK) fn
    Return,         // (OK) return
    This,           // (OK) self
    True,           // (OK) true,
    False,          // (OK) false
    Var,            // (OK) var
    Continue,       // (OK) continue
    Break,          // (OK) break
    Public,         // (OK) public
    Private,        // (OK) private
    LeftFlow,       // (OK) <-
    RightFlow,      // (OK) ->
    // other
    Plus,           // (OK) +
    Minus,          // (OK) -
    Star,           // (OK) *
    ForwardSlash,   // (OK) /
    OpenParen,      // (OK) (
    CloseParen,     // (OK) )
    OpenBrace,      // (OK) {
    CloseBrace,     // (OK) }
    SemiColon,      // (OK) ;
    Comma,          // (OK) ,
    Dot,            // (OK) .
    AmpAmp,         // (OK) &&
    PipePipe,       // (OK) ||

    Bang,               // (OK) !
    BangEqual,          // (OK) !=
    EqualEqual,         // (OK) ==
    Equal,              // (OK) =
    Greater,            // (OK) >
    GreaterEqual,       // (OK) >=
    Less,               // (OK) <
    LessEqual,          // (OK) <=
    CommentSingleLine,  // (OK) //
    CommentMultiLine,   // (OK) /* */
    EOF,
    Illegal,
}

impl Display for TokenType {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self)
    }
}