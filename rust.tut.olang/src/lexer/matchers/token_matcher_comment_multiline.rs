use crate::lexer::{InputReader, LiteralValue, SymbolLookupResult, TokenState};
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::matchers::token_matcher_helper::TokenMatcherHelper;
use crate::lexer::TokenType;

pub struct TokenMatcherCommentMultiLine {}

impl<'a> TokenMatcher<'a> for TokenMatcherCommentMultiLine {
    fn create(&self, reader: &mut InputReader) -> Option<Token> {
        let mut has_ending = false;

        let result = TokenMatcherHelper::match_symbol_extract_literal(
            reader,
            "/*",
            TokenType::CommentMultiLine,
            |reader|{
                match reader.advance_until_symbol_found("*/")
                {
                    SymbolLookupResult::Found => {
                        let rest_of_line = reader.collect().to_string();
                        let literal_value = LiteralValue::StringValue(String::from(rest_of_line));
                        has_ending = true;
                        Some(literal_value)
                    }
                    SymbolLookupResult::EndOfFileReached => {
                        let rest_of_line = reader.collect().to_string();
                        let literal_value = LiteralValue::StringValue(String::from(rest_of_line));
                        has_ending = false;
                        Some(literal_value)
                    }
                }
            }
        );

        match result{
            Some(value)  =>{
                reader.advance(2);
                reader.forward();
                if has_ending {
                    Some(value)
                }else {
                    let mutation = Token{
                        state : TokenState::MissingEndToken,
                        literal_value: value.literal_value,
                        lexeme: value.lexeme,
                        position: value.position.clone(),
                        token_type: value.token_type
                    };

                    Some(mutation)
                }
            }
            None =>{
                None
            }
        }

    }
}