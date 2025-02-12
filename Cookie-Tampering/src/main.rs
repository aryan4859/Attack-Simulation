use actix_web::{web, App, HttpServer, HttpResponse, HttpRequest, Responder};
use actix_session::{Session, SessionMiddleware, storage::CookieSessionStore};
use actix_web::cookie::Key;
use base64::{engine::general_purpose, Engine as _};

async fn index(req: HttpRequest, session: Session) -> impl Responder {
    let user_cookie = session.get::<String>("user").unwrap_or(Some("guest".to_string()));

    let username = user_cookie.unwrap_or("guest".to_string());
    let decoded = general_purpose::STANDARD.decode(&username).unwrap_or_else(|_| b"guest".to_vec());

    if decoded == b"admin" {
        HttpResponse::Ok().body("Welcome, Admin! 🎉 Here is your flag: Flag{cookie_tampering_success}")
    } else {
        HttpResponse::Ok().body("Welcome, Guest! Try modifying your cookie to become an admin.")
    }
}

async fn login(session: Session) -> impl Responder {
    let encoded_guest = general_purpose::STANDARD.encode("guest");
    session.insert("user", encoded_guest).unwrap();
    HttpResponse::Ok().body("Logged in as Guest. Try modifying your cookie!")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let secret_key = Key::generate(); // Required for session encryption

    HttpServer::new(move || {
        App::new()
            .wrap(SessionMiddleware::new(
                CookieSessionStore::default(),
                secret_key.clone(),
            ))
            .route("/", web::get().to(index))
            .route("/login", web::get().to(login))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
