
// use axum::{
//     extract::{self, BodyStream, Extension, Json, Path},
//     http::{Request, StatusCode},
//     response::IntoResponse,
// };

// use hyper::body;
// use hyper::{Body, Request};
// use rdkafka::producer::{FutureProducer, FutureRecord};
// use serde_json::json;
// use std::time::Duration;

// pub async fn post_ts(
//     Extension(kafka_producer): Extension<FutureProducer>,
//     Path(asset_id): Path<String>,
//     request: Request<Body>
// ) -> impl IntoResponse {
//     // let payload_string = format!("{:?}", &payload);

//     let bytes = body::to_bytes(request.into_body()).await?;
    
//     kafka_producer
//     .send(
//         FutureRecord::to("ts")
//             .key(&asset_id)
//             .payload(&payload_string),
//         Duration::from_secs(0),
//     )
//     .await
//     .expect("Failed to put msg into Kafka");


//     (StatusCode::OK, Json(json!({ "TS": "OK" })))
// }
