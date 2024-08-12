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
    Func,           // fn (OK)
    Nil,            // null
    Return,         // return
    This,           // self,
    True,           // true,
    False,          // false
    Var,            // (OK) var
    Continue,       // (OK) continue
    Break,          // (OK) break
    Public,         // public
    ScopeReturn,    // <-
    // other
    Plus,           // (OK) +
    Minus,          // (OK) -
    Star,           // *
    Slash,          // /
    OpenParen,      // (OK) (
    CloseParen,     // (OK) )
    OpenBrace,      // (OK) {
    CloseBrace,     // (OK) }
    SemiColon,      // (OK) ;
    Comma,          //,
    Dot,            // .
    AmpAmp,         // (OK) &&
    PipePipe,       // (OK) ||

    Bang,               // (OK) !
    BangEqual,          // (OK) !=
    EqualEqual,         // (OK) ==
    Equal,              // (OK) =
    Greater,            // >
    GreaterEqual,       // (OK) >=
    Less,               // <
    LessEqual,          // (OK) <=
    CommentSingleLine,  // (OK) //
    CommentMultiLine,   // /* */
    EOF,
    Illegal,
}

impl Display for TokenType {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self)
    }
}