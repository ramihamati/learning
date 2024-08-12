use crate::lexer::{InputReader, TokenLinePosition, TokenLocalisation, TokenMatcherAmpAmp, TokenMatcherBreak, TokenMatcherPipePipe, TokenState};
use crate::lexer::{TokenMatcherCloseBrace, TokenMatcherContinue, TokenMatcherCommentLine, TokenMatcherCommentMultiLine, };
use crate::lexer::{TokenMatcherCloseParen, TokenMatcherEqual, TokenMatcherStruct, TokenMatcherVar};
use crate::lexer::{TokenMatcherComma, TokenMatcherGreaterEqual, TokenMatcherLessEqual, TokenMatcherSemiColon, TokenMatcherStar};
use crate::lexer::{TokenMatcherEqualEqual, TokenMatcherBang};
use crate::lexer::{TokenMatcherOpenBrace, TokenMatcherBangEqual};
use crate::lexer::{TokenMatcherOpenParen, TokenMatcherFn};
use crate::lexer::Token;
use crate::lexer::{TokenMatcher};
use crate::lexer::TokenType;
use crate::lexer::TokenMatcherPlus;
use crate::lexer::TokenMatcherMinus;
use crate::lexer::TokenMatcherIf;
use crate::lexer::TokenMatcherElse;
use crate::lexer::TokenMatcherLoop;

pub struct Lexer<'a> {
    tokens: Vec<Token>,
    reader: InputReader<'a>,
    matchers: Vec<Box<dyn TokenMatcher<'a>>>,
}


impl<'a> Lexer<'a> {
    pub(crate) fn new(input: &'a str) -> Self {
        let mut matchers: Vec<Box<dyn TokenMatcher>> = vec![];

        /* order as priority matters
             == should be recognized before = therefore == matcher comes before the = matcher
        */
        // double length fixed characters
        matchers.push(Box::new(TokenMatcherEqualEqual {}));
        matchers.push(Box::new(TokenMatcherAmpAmp {}));
        matchers.push(Box::new(TokenMatcherPipePipe {}));
        matchers.push(Box::new(TokenMatcherGreaterEqual {}));
        matchers.push(Box::new(TokenMatcherLessEqual {}));
        matchers.push(Box::new(TokenMatcherCommentMultiLine {}));
        matchers.push(Box::new(TokenMatcherBangEqual {}));
        matchers.push(Box::new(TokenMatcherFn {}));

        // fixed length single characters
        matchers.push(Box::new(TokenMatcherOpenParen {}));
        matchers.push(Box::new(TokenMatcherCloseParen {}));
        matchers.push(Box::new(TokenMatcherOpenBrace {}));
        matchers.push(Box::new(TokenMatcherCloseBrace {}));
        matchers.push(Box::new(TokenMatcherComma {}));
        matchers.push(Box::new(TokenMatcherPlus {}));
        matchers.push(Box::new(TokenMatcherMinus {}));
        matchers.push(Box::new(TokenMatcherEqual {}));
        matchers.push(Box::new(TokenMatcherSemiColon {}));
        matchers.push(Box::new(TokenMatcherStar {}));
        matchers.push(Box::new(TokenMatcherBang {}));

        // reserved keywords bound by identifiers
        matchers.push(Box::new(TokenMatcherCommentLine {}));
        matchers.push(Box::new(TokenMatcherIf {}));
        matchers.push(Box::new(TokenMatcherElse {}));
        matchers.push(Box::new(TokenMatcherLoop {}));
        matchers.push(Box::new(TokenMatcherStruct {}));
        matchers.push(Box::new(TokenMatcherVar {}));
        matchers.push(Box::new(TokenMatcherContinue {}));
        matchers.push(Box::new(TokenMatcherBreak{}));

        Lexer {
            tokens: vec![],
            reader: InputReader::new(input),
            matchers,
        }
    }

    pub fn get_next_token(self: &mut Self) -> Option<Token> {
        for matcher in &self.matchers {
            let res = matcher.create(&mut self.reader);
            match res {
                None => {
                    continue;
                }
                Some(token) => {
                    return Some(token);
                }
            }
        }

        return None;
    }

    pub fn scan_tokens(self: &mut Self) -> Result<Vec<Token>, String> {
        // let mut errors = vec![];
        let mut unidentifierd = String::new();

        while self.reader.can_advance() {
            self.reader.forward_if_new_line();
            let next_token = self.get_next_token();
            match next_token {
                Some(token) => {
                    self.tokens.push(token);
                }
                None => {
                    self.reader.advance(1);
                    unidentifierd = format!("{}{}", unidentifierd, self.reader.collect());
                    self.reader.forward();
                }
            }

            // errors.push("failed to match");
            // self.reader.scanner_advance(1);
            // self.reader.scanner_forward();
        }

        self.add_eof();
        // if errors.len() > 0 {
        //     let mut joined = "".to_string();
        //     for err in errors {
        //         joined.push_str(&err);
        //         joined.push_str("\n");
        //     }
        //
        //     return Err(joined);
        // }
        println!("unidentified {} {}", unidentifierd.len(), unidentifierd);
        Ok(self.tokens.clone())
    }

    fn add_eof(self: &mut Self) {
        self.tokens.push(Token {
            position:  TokenLocalisation {
                start : TokenLinePosition{
                    position: self.reader.line_current,
                    line_number: self.reader.line
                },
                end : TokenLinePosition{
                    position: self.reader.line_current,
                    line_number: self.reader.line
                }
            },
            literal_value: None,
            token_type: TokenType::EOF,
            lexeme: "".to_string(),
            state: TokenState::Ok
        });
    }
}