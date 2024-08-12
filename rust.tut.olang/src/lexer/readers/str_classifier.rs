pub struct StrClassifier {

}

impl StrClassifier {
    pub fn is_identifier_char(c: char) -> bool{
        return c.is_alphanumeric() || c == '_';
    }

}