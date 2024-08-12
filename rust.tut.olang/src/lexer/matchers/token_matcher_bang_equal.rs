use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherBangEqual {
}

impl<'a> TokenMatcher<'a> for TokenMatcherBangEqual {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol(
            reader,
            "!=",
            TokenType::BangEqual
        )
    }
}