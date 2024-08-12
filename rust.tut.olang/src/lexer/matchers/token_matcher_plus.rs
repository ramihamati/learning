use crate::lexer::readers::input_reader::InputReader;
use crate::lexer::symbols::token::Token;
use crate::lexer::matchers::token_matcher::{ TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::symbols::token_type::TokenType;

pub struct TokenMatcherPlus {
}

impl<'a> TokenMatcher<'a> for TokenMatcherPlus {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_character(
            reader,
            '+',
            TokenType::Plus)
     }
}