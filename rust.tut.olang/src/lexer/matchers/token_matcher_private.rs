use crate::lexer::InputReader;
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::str_classifier::StrClassifier;
use crate::lexer::TokenType;

pub struct TokenMatcherPrivate {}

impl<'a> TokenMatcher<'a> for TokenMatcherPrivate {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol_bounded_by(
            reader,
            "private",
            TokenType::Private,
            |next_char| {
                !StrClassifier::is_identifier_char(next_char)
            }
        )
    }


}