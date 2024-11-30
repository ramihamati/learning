use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherRightFlow {
}

impl<'a> TokenMatcher<'a> for TokenMatcherRightFlow {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol(
            reader,
            "->",
            TokenType::RightFlow,
        )
    }
}