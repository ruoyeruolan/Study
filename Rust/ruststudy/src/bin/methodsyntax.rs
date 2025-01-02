#[derive (Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {  // Methods with more than one parameter
        self.width > other.width && self.height > other.height
    }

    fn square(size: u32) -> Self {  // related to the struct itself, not an instance of the struct
        Self {
            width: size,
            height: size,
        }
    }
    // In common Rust terminology, the first parameter of a method is always self and if the method does not take self as a parameter, it is called an associated function.
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };


    
    println!("The area of the rectangle is {} square pixels.", rect1.area());

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect2 hold rect1? {}", rect2.can_hold(&rect1));

    let sq = Rectangle::square(3);
    println!("The area of the square is {} square pixels.", sq.area());
}