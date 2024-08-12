use std::collections::HashMap;
use std::fmt::{Debug};

#[derive(Debug)]
enum BidOrAsk{
    Bid,
    Ask
}
#[derive(Debug)]
struct OrderBook{
    asks: HashMap<Price, Limit>,
    bids : HashMap<Price,  Limit>
}

impl OrderBook{
    fn new() -> OrderBook{
        OrderBook{
            asks : HashMap::new(),
            bids : HashMap::new()
        }
    }

    fn add_order(&mut self, price: f64, order: Order){
        match order.bid_or_ask{
            BidOrAsk::Bid =>{
                
            },
            BidOrAsk::Ask =>{

            }
        }
    }
}

#[derive(Debug)]
struct Price{
    integral : u64,
    fractional : u64,
    scalar: u64
}

impl Price{
    fn new(price: f64) -> Price{
        let scalar = 100000;
        let integral = price as u64;
        let fractional = ((price % 1.0) * scalar as f64) as u64;

        Price{
            scalar,
            integral,
            fractional
        }
    }
}

#[derive(Debug)]
struct Limit{
    price: Price,
    orders: Vec<Order>
}

impl Limit{
    fn new(price : f64) -> Limit{
        Limit{
            price : Price::new(price),
            orders: Vec::new()
        }
    }

    fn add_order(&mut self, order: Order){
        self.orders.push(order)
    }
}

#[derive(Debug)]
struct Order{
    size : f64,
    bid_or_ask: BidOrAsk
}

impl Order{
    fn new(bid_or_ask: BidOrAsk, size: f64) -> Order{
        return Order{
            bid_or_ask,
            size
        }
    }
}

fn main() {
    let price = Price::new(50.5);
    let mut limit = Limit::new(40.5);

    limit.add_order(Order{
        size: 20.3,
        bid_or_ask: BidOrAsk::Bid
    });

    println!("{:?}", price);
    println!("{:?}", limit);
    println!("Hello, world!");
}
