use actix_web::{web, App, HttpServer, HttpRequest, HttpResponse, Responder};
use actix_session::{Session, CookieSession};
use base64::{encode, decode};

const FLAG: &str = "Flag{cookie_tampering_success}";

async fn index(req: HttpRequest) -> impl Responder {
    let cookie = req.cookie("user");

    if let Some(user_cookie) = cookie {
        let decoded = decode(user_cookie.value()).unwrap_or_else(|_| b"guest".to_vec());
        let user = String::from_utf8(decoded).unwrap_or_else(|_| "guest".to_string());

        if user == "admin" {
            return HttpResponse::Ok().body(format!("Welcome, Admin! 🎉 Here is your flag: {}", FLAG));
        } else {
            return HttpResponse::Ok().body(format!("Hello, {}. You are not an admin! 🍪", user));
        }
    }

    let encoded_guest = encode("guest");
    HttpResponse::Ok()
        .cookie(actix_web::cookie::Cookie::new("user", encoded_guest))
        .body("Cookie set! Refresh and check your cookies.")
}

async fn login_page() -> impl Responder {
    let html = r#"
        <html>
            <head><title>Login</title></head>
            <body>
                <h2>Fake Login</h2>
                <form action="/set-cookie" method="post">
                    <label>Username: <input type="text" name="username"></label><br>
                    <label>Password: <input type="password" name="password"></label><br>
                    <input type="submit" value="Login">
                </form>
            </body>
        </html>
    "#;
    HttpResponse::Ok().content_type("text/html").body(html)
}

async fn set_cookie(form: web::Form<std::collections::HashMap<String, String>>) -> impl Responder {
    let username = form.get("username").unwrap_or(&"guest".to_string()).clone();
    let encoded = encode(username);

    HttpResponse::Found()
        .append_header(("Location", "/"))
        .cookie(actix_web::cookie::Cookie::new("user", encoded))
        .finish()
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .wrap(CookieSession::signed(&[0; 32]).secure(false))
            .route("/", web::get().to(index))
            .route("/login", web::get().to(login_page))
            .route("/set-cookie", web::post().to(set_cookie))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
