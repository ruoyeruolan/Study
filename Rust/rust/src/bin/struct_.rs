struct Renctangle {
    width: u32,
    height: u32,
}

fn main() {
    let width = 30;
    let height = 50;

    println!("The area of the rectangle is {} square pixels.", area(width, height));

    let rect1  = Renctangle {
        width: 30,
        height: 50,
    };
    println!("The area of the rectangle is {} square pixels.", area_(&rect1));
    // println!("rect1 is {rect1:?}");  // This will not work if the struct does not implement Debug trait, you can add #[derive(Debug)] to the struct definition to make it work.
}

fn area(width: u32, height: u32) -> u32 {
    width * height
}

fn area_(rect: &Renctangle) -> u32 {
    rect.width * rect.height
}
