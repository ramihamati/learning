use crate::lexer::{InputReader, Token, TokenMatcher, TokenMatcherHelper, TokenType};

pub struct TokenMatcherFn{
}

impl<'a> TokenMatcher<'a> for TokenMatcherFn{
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        TokenMatcherHelper::match_symbol(
            reader,
            "fn",
            TokenType::Func
        )
    }
}