use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherAmpAmp {
}

impl<'a> TokenMatcher<'a> for TokenMatcherAmpAmp {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol(
            reader,
            "&&",
            TokenType::AmpAmp
        )
    }
}