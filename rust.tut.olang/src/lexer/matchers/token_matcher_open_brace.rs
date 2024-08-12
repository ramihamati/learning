use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherOpenBrace {
}

impl<'a> TokenMatcher<'a> for TokenMatcherOpenBrace {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_character(
            reader,
            '{',
            TokenType::OpenBrace)
    }
}