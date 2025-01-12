struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn main() {
    let user1 = User {
        email: String::from("someone@example.com"),
        username: String::from("someusername123"),
        active: true,
        sign_in_count: 1,
    }; // This is an instance of User struct and it is immutable.

    let user2 = User{
        active: false,
        ..user1 // This is a struct update syntax. It is used to update the value of the struct.
    };  // This is an instance of User struct and it is immutable. Also, user1 is not valid anymore.besuase it is moved to user2.

    // If the data is stack only, it will be copied default and the original data will still be valid.
    // If the data is heap, it will be moved to the new variable and the original data will not be valid anymore.
}
