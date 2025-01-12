fn main() {
    let mut s = String::from("hello world");
    let word = first_word(&s);  // This will get the first word of the string.
    println!("{}", word);
    s.clear();

    // println!("{}", word); // This will cause an error because the mutabale and immutable reference are not allowed at the same time.
}

fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }
    &s[..]
}