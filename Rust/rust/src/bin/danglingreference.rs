fn main() {
    // let reference_to_nothing = dangle();
    let reference = no_dangle();
}

fn dangle() -> &String {
    let s = String::from("hello");
    &s

    // `s` is created in the function and it will be deallocated when the function ends.
    // So, the return of `&s` will be a dangling reference and it will cause an error.
}

fn no_dangle() -> String {
    let s = String::from("hello");
    s
}