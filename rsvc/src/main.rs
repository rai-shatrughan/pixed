use axum::{
    routing::{get, put, post, get_service},
    extract::Extension,
    Router,
    http::StatusCode,
    response::IntoResponse,

};
use std::{io};
use tower_http::{trace::TraceLayer};

mod util;
mod handler;
mod model;

#[tokio::main]
async fn main() {

    if std::env::var_os("RUST_LOG").is_none() {
        std::env::set_var(
            "RUST_LOG",
            "example_tracing_aka_logging=error,tower_http=error",
        )
    }

    tracing_subscriber::fmt::init();
    util::kv::put(String::from("sr"), String::from("sr-value"));
    println!("{}",util::kv::get(String::from("sr")));
    println!("{}",util::kv::get(String::from("s2")));
    let kafka_producer = util::kafka::KafkaProducer::init();
    let service = tower_http::services::ServeDir::new(util::conf::Props::init().web_dir);
    let app = Router::new()
                .route("/api/v1/am", get(handler::asset::hello))
                .route("/api/v1/timeseries/:asset_id", put(handler::timeseries::ts_put))
                .route("/api/mindconnect/v3/exchange/:asset_id", post(handler::exchange::post_multi_ts))
                .layer(Extension(kafka_producer))
                .layer(TraceLayer::new_for_http())
                .fallback(get_service(service).handle_error(handle_error));

    println!("*** serving web_dir {} - at address {}", util::conf::Props::init().web_dir, util::conf::Props::init().web_address);

    axum::Server::bind(&util::conf::Props::init().web_address.parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handle_error(_err: io::Error) -> impl IntoResponse {
    (StatusCode::INTERNAL_SERVER_ERROR, "Error Serving Static Server")
}