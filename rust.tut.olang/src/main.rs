use crate::lexer::Lexer;
mod lexer;

fn main() {
    let code = "!=!fn";

    let mut lexer = Lexer::new(code);
    let tokens = lexer.scan_tokens();

    match tokens{
        Ok(tokens) => {
            println!("tokens {:#?}", tokens)
        },
        Err(e) => {
            println!("errors: \n{}", e)
        }
    }
    //
    // let a = "hello";
    // let slice = &a[4..5];
    //
    // println!("{}", slice);
}
