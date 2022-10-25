use axum::{
    extract::{Extension, Json, Path},
    http::{Request, StatusCode},
    response::IntoResponse, 
    body::Body,
};

use rdkafka::producer::{FutureProducer, FutureRecord};
use serde_json::json;
use std::time::Duration;

pub async fn post_ts(
    Extension(kafka_producer): Extension<FutureProducer>,
    Path(asset_id): Path<String>,
    req: Request<Body>
) -> impl IntoResponse {
    // let payload_string = format!("{:?}", &payload);

    let body_bytes = match hyper::body::to_bytes(req.into_body()).await {
        Ok(v) => v,
        Err(e) => panic!("Invalid UTF-8 sequence: {}", e),
    };

    let payload_string = format!("{:?}", body_bytes);

    kafka_producer
    .send(
        FutureRecord::to("ts")
            .key(&asset_id)
            .payload(&payload_string),
        Duration::from_secs(0),
    )
    .await
    .expect("Failed to put msg into Kafka");


    (StatusCode::OK, Json(json!({ "TS": "OK" })))
}

pub async fn post_multi_ts(
    Extension(kafka_producer): Extension<FutureProducer>,
    Path(asset_id): Path<String>,
    req: Request<Body>
) -> impl IntoResponse {
    // let payload_string = format!("{:?}", &payload);

    let body_bytes = match hyper::body::to_bytes(req.into_body()).await {
        Ok(v) => v,
        Err(e) => panic!("Invalid UTF-8 sequence: {}", e),
    };

    let payload_string = format!("{:?}", body_bytes);
    let part2_string: Vec<&str> = payload_string.split("Content-Type: application/json").collect();
    let json_string: Vec<&str> = part2_string[1].split("--").collect();

    println!("{}",json_string[0]);
    
    kafka_producer
    .send(
        FutureRecord::to("ts")
            .key(&asset_id)
            .payload(json_string[0]),
        Duration::from_secs(0),
    )
    .await
    .expect("Failed to put msg into Kafka");


    (StatusCode::OK, Json(json!({ "TS": "OK" })))
}