use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::str_classifier::StrClassifier;
use crate::lexer::TokenType;

pub struct TokenMatcherLoop {}

impl<'a> TokenMatcher<'a> for TokenMatcherLoop {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol_bounded_by(
            reader,
            "loop",
            TokenType::Loop,
            |next_char| {
                !StrClassifier::is_identifier_char(next_char)
            }
        )
    }


}