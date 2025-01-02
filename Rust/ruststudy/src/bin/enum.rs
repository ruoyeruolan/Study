enum IpAddrKind {  // a custom data type
    V4, 
    V6,
}

struct IpAddr {
    kind: IpAddrKind,
    address: String,
}

enum IpAddr_ {
    V4(String),
    V6(String),
}

fn main() {
    // let four = IpAddrKind::V4;
    // let six = IpAddrKind::V6;

    // route(four);
    // Common way to define an instance of the IpAddr struct
    let home  = IpAddr {
        kind: IpAddrKind::V4,
        address: String::from("127.0.0.1"),
    };
    
    let loopback = IpAddr {
        kind: IpAddrKind::V6,
        address: String::from("::1"),
    };

    println!("The home address is: {}", home.address);
    println!("The loopback address is: {}", loopback.address);

    // Another way to define an instance of the IpAddr struct with enums
    let home_ = IpAddr_::V4(String::from("127.0.0.1"));
    let loopback_ = IpAddr_::V6(String::from("::1"));
    // println!("The home address is: {#?}", home_.address);

}

fn route(ip_kind: IpAddrKind) {}