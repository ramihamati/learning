use std::fmt::{Display, Formatter};

#[derive(Debug, PartialEq, Clone)]
pub enum LiteralValue{
    Int8Value(i8),
    Int32Value(i32),
    Int64Value(i64),
    Float32Value(f32),
    Float64Value(f64),
    StringValue(String),
    IdentifierValue(String)
}

impl  Display for LiteralValue{
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        match self{
            LiteralValue::Float32Value(value) =>{
                write!(f, "{}", value)
            },
            LiteralValue::Int8Value(value) => {
                write!(f, "{}", value)
            },
            LiteralValue::Int32Value(value) => {
                write!(f, "{}", value)
            },
            LiteralValue::Int64Value(value) => {
                write!(f, "{}", value)
            }
            LiteralValue::Float64Value(value) => {
                write!(f, "{}", value)
            },
            LiteralValue::StringValue(value) => {
                write!(f, "{}", value)
            },
            LiteralValue::IdentifierValue(value) => {
                write!(f, "{}", value)
            }
        }
    }
}