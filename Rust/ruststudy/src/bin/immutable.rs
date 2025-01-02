fn main() {
    let mut s = String::from("Hello");
    let r1 = &s;
    let r2 = &s;
    let r3 = &mut s;  // No error at present

    // println!("{}, {}, {}", r1, r2, r3);  
    // But when we try to use r1, r2, r3 at the same time, we will get an error
    // in this case, if we just use r1 and r2 or just r3, we will not get an error.
    // Also, if we r1 and r2 before `let rs = &mut s`, we will not get an error.
    // Conclusion: We can have either one mutable reference or multiple immutable references, but not both.

    println!("{}", r3); 
}