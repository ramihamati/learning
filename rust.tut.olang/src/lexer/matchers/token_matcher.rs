use crate::lexer::InputReader;
use crate::lexer::Token;

pub trait TokenMatcher<'a> {
    fn create(&self, reader: &mut InputReader) -> Option<Token>;
}