// use core::num;
use std::io;
use rand::Rng;
use std::cmp::Ordering;
// use crate::front_of_house::hosting::add_to_waitlist;
fn main() {
    println!("Guess the number!");
    let secret_number:u32 = rand::thread_rng().gen_range(1..=100);
    println!("The secret number is: {secret_number}");
    
    loop {

        println!("Please input your guess.");
        let mut guess:String = String::new();
        io::stdin().read_line(&mut  guess).expect("Failed to read line");
        let guess: u32  = match guess.trim().parse() {  // parse: Convert string to number
            Ok(num) => num,
            Err(_) => continue,
        };
        println!("You guessed: {}", guess);

        // Compare the guess number with the secret number
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }

    let mut s = String::from("Hello");
    s.push_str(", Wold!");
    println!("{}", s);

    let x = 5;
    let y = x;
    println!("x = {}, y = {}", x, y);
}
