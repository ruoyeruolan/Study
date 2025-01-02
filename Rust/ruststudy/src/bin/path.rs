use ruststudy::eat_at_restaurant;
use std::collections::HashMap;

fn main() {
    eat_at_restaurant();

    let mut map = HashMap::new();
    map.insert(1, 2);
    println!("{:?}", map);
}