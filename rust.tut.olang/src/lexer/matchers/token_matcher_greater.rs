use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherGreater {
}

impl<'a> TokenMatcher<'a> for TokenMatcherGreater {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol(
            reader,
            ">",
            TokenType::Greater,
        )
    }
}