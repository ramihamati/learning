use crate::lexer::Token;
use crate::lexer::InputReader;
use crate::lexer::{TokenMatcher};
use crate::lexer::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherCloseBrace {}

impl<'a> TokenMatcher<'a> for TokenMatcherCloseBrace {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_character(
            reader,
            '}',
            TokenType::CloseBrace)
    }
}