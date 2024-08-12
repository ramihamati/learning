use std::string::ToString;
use crate::lexer::readers::SymbolLookupResult;

pub struct InputReader<'a> {
    input: &'a str,
    input_as_bytes: &'a [u8],
    pub line_current: usize, // line position of current cursor of token scan
    pub line_start: usize, // line position of start of token scan
    pub scanner_current: usize, // position in line
    pub line: usize, // line number
    pub scanner_start: usize, // scanner start
}
const EOL : char = '\n';


impl<'a> InputReader<'a> {
    pub fn new(input: &'a str) -> InputReader {
        InputReader {
            input,
            input_as_bytes: input.as_bytes(),
            line_current: 0, // position in current line
            line_start: 0, // position in current line
            scanner_current: 0, // current position per document. does not include last char
            scanner_start: 0, // where has the token started
            line: 0, // current line
        }
    }

    pub fn collect(&self) -> &str {
        /*
            let code = "hello"
            code[5..6] -> panic
            code[5..5] -> ""
            code[4..5] -> o
         */
        return &self.input[self.scanner_start..self.scanner_current];
    }

    pub fn advance(self: &mut Self, char_count: usize) -> bool {
        // if code.len = 5
        // if current = 5 & start = 4 we have a valid [4..5]
        // this means that line_current + char_count can be self.input.len()
        if (self.scanner_start + char_count) > self.input.len() {
            return false;
        }

        self.line_current += char_count;
        self.scanner_current += char_count;

        return true;
    }

    pub fn advance_until_symbol_found(self: &mut Self, symbol : &str) -> SymbolLookupResult
    {
        // code = "/*"
        // scanner_start is 2 (= len())
        // scanner_start + 1 is 3
        // code[2..3] would panic
        let symbol_length = symbol.len();

        if (self.scanner_current + symbol_length) > self.input.len() {
            return SymbolLookupResult::EndOfFileReached;
        }

        // code = "a/*a"
        // scanner_current is 3 (repr a)
        loop{
            // current us code[3..4] => a (last char)
            let current =self.input[self.scanner_current..self.scanner_current + symbol_length].to_string();
            if  current == symbol {
                return SymbolLookupResult::Found;
            }
            if self.advance_if_new_line() == false {
                self.scanner_current += 1;
                self.line_current += 1;
            }

            // code[4..5] would panic
            if (self.scanner_current + symbol_length) > self.input.len() {
                return SymbolLookupResult::EndOfFileReached;
            }
        }
    }

    pub fn advance_until_end_of_line(self: &mut Self) -> bool
    {
        let end_of_line_symbol = EOL.to_string();

        // code = "//"
        // scanner_start is 2 (= len())
        // scanner_start + 1 is 3
        // code[2..3] would panic
        if (self.scanner_current + 1) > self.input.len() {
            return false;
        }

        // code = "a//a"
        // scanner_current is 3 (repr a)
        loop{
            // current us code[3..4] => a (last char)
            let current =self.input[self.scanner_current..self.scanner_current + 1].to_string();
            if  current == end_of_line_symbol {
                break;
            }
            self.scanner_current += 1;
            self.line_current += 1;
            // code[4..5] would panic
            if (self.scanner_current + 1) > self.input.len() {
                break;
            }
        }

        return true;
    }
    pub fn peek_one(self: &Self) -> Option<char>{
        /*
         let code = "hello"
         if current == len then [current .. current+1] means [len .. len + 1]
         code[5..5+1] -> panic
         code[4..4+1] -> "o"
      */
        return if self.scanner_current + 1 >= self.input.len() {
            // if 4 + 1 = 5 and code.len() = 5 we have nothing to peek
            None
        } else {
            let peeked = &self.input[self.scanner_current..self.scanner_current + 1];
            peeked.chars().next()
        }
    }

    pub fn revert_advance(self: &mut Self){
        self.scanner_current = self.scanner_start;
        self.line_current = self.line_start;
    }

    pub(crate) fn forward(self: &mut Self) {
        self.scanner_start = self.scanner_current;
        self.line_start = self.line_current;
    }

    pub fn can_advance(self: &Self) -> bool {
        // if code.len = 5
        // if current = 5 it means that a next advance would be 5 + 1
        //    which would result in current = 6 which is outside of bounds
        //
        self.scanner_current < self.input.len()
    }

    pub fn advance_if_new_line(self: &mut Self) -> bool {
        match self.peek_one(){
            Some(ch) =>{
                if ch == EOL {
                    self.advance(1);
                    // resetting after advance to have the correct line_start and line_current
                    self.line += 1;
                    self.line_start = 0;
                    self.line_current = 0;
                    true
                }
                else{
                    false
                }
            }
            None => {
                false
            }
        }
    }

    pub fn forward_if_new_line(self: &mut Self) -> bool {
        match self.peek_one(){
            Some(ch) =>{
                if ch == EOL {
                    self.advance(1);
                    self.forward();
                    // resetting after advance to have the correct line_start and line_current
                    self.line += 1;
                    self.line_start = 0;
                    self.line_current = 0;
                    true
                }
                else{
                    false
                }
            }
            None => {
                false
            }
        }
    }
}