use crate::lexer::{InputReader, LiteralValue};
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherCommentLine {}

impl<'a> TokenMatcher<'a> for TokenMatcherCommentLine {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
         TokenMatcherHelper::match_symbol_extract_literal(
            reader,
            "//",
            TokenType::CommentSingleLine,
            |reader|{
                return if reader.advance_until_end_of_line() {
                    let literal_str = reader.collect().to_string();
                    let literal_value = LiteralValue::StringValue(String::from(literal_str));
                    Some(literal_value)
                } else {
                    None
                }
            }
        )
    }
}