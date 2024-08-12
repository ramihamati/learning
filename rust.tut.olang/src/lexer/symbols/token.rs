use std::fmt::{Display, Formatter};
use crate::lexer::LiteralValue;
use crate::lexer::TokenType;

#[derive(Debug, Clone)]
pub enum TokenState{
    Ok,
    MissingEndToken
}

#[derive(Debug, Clone)]
pub struct TokenLinePosition {
    pub line_number: usize,
    pub position: usize,
}

#[derive(Debug, Clone)]
pub struct TokenLocalisation {
    pub start: TokenLinePosition,
    pub end: TokenLinePosition,
}

#[derive(Debug, Clone)]
pub struct Token {
    pub token_type: TokenType,
    pub lexeme: String,
    pub literal_value: Option<LiteralValue>,
    pub position: TokenLocalisation,
    pub state: TokenState,
}

impl Token {
    pub fn to_string(self: &Self) -> String {
        format!("{} {} {:?}", self.token_type, self.lexeme, self.literal_value)
    }
}

impl Display for Token{
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "\
            {{ token_type: {}, \
               literal: {:?}, \
               line_start: {{\
                   number: {},\
                   position : {}\
               }}, \
               line_end : {{\
                    number: {},\
                    position: {}\
               }}\
             }}",
               self.token_type, self.literal_value,
               self.position.start.line_number,
               self.position.start.position,
               self.position.end.line_number,
               self.position.end.position,

        )
    }
}